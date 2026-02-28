# 🔥 Go Quiz — TCP Networking, Concurrency & Mutex

## BLOCK 1 — TCP Listener & Accept Loop

### Problem 1: Starting a Listener ⭐
```go
ln, err := net.Listen("tcp", ":8080")
if err != nil {
    log.Fatal(err)
}
fmt.Println("listening...")
```
**Question:** What does `":8080"` mean? What does `net.Listen` return?

**Key Concept:** `":"` before the port = listen on all interfaces. Always check the error before using the listener!

---

### Problem 2: The Accept Loop ⭐
```go
for {
    conn, err := ln.Accept()
    if err != nil {
        log.Println(err)
        continue
    }
    handleConn(conn)
}
```
**Question:** What is the problem with this loop? What happens when a second client connects while the first is being handled?

**Key Concept:** An accept loop without `go` is sequential — always use `go handleConn(conn)` for concurrent clients!

---

### Problem 3: Accept Returns What ⭐
```go
conn, err := ln.Accept()
fmt.Printf("%T\n", conn)
fmt.Println(conn.RemoteAddr())
```
**Question:** What type is `conn`? What does `RemoteAddr()` return?

**Key Concept:** `net.Conn` is an interface — it has `Read`, `Write`, `Close`, `RemoteAddr` and more!

---

### Problem 4: Closing the Listener ⭐⭐
```go
ln, _ := net.Listen("tcp", ":9000")
defer ln.Close()

for {
    conn, err := ln.Accept()
    if err != nil {
        fmt.Println("accept error:", err)
        return
    }
    go handleConn(conn)
}
```
**Question:** What happens when `ln.Close()` is called (e.g. on program exit)? Does the running accept loop notice?

**Key Concept:** Closing a listener unblocks `Accept()` with an error — use this for graceful shutdown!

---

### Problem 5: Port Already in Use ⭐⭐
```go
ln1, err1 := net.Listen("tcp", ":8080")
ln2, err2 := net.Listen("tcp", ":8080")
fmt.Println(err1)
fmt.Println(err2)
```
**Question:** What does each line print?

**Key Concept:** Only ONE listener can bind to a port at a time — the second attempt always fails!

---

## BLOCK 2 — Goroutines & go Keyword

### Problem 6: Launching a Goroutine ⭐
```go
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    go greet("Alice")
    go greet("Bob")
    time.Sleep(100 * time.Millisecond)
}
```
**Question:** Is the output always `Hello, Alice` then `Hello, Bob`? Why or why not?

**Key Concept:** Goroutines have no guaranteed order — never rely on goroutine execution sequence!

---

### Problem 7: Goroutine vs Function Call ⭐
```go
func task() {
    fmt.Println("running")
}

func main() {
    task()       // Line A
    go task()    // Line B
    fmt.Println("done")
}
```
**Question:** What is guaranteed to print? What might not print?

**Key Concept:** `go func()` is fire-and-forget — main won't wait for it unless you synchronize!

---

### Problem 8: Goroutine Per Connection ⭐⭐
```go
for {
    conn, err := ln.Accept()
    if err != nil {
        return
    }
    go func() {
        handleClient(conn)
    }()
}
```
**Question:** Is there a bug in this goroutine closure? What is the classic closure variable trap here?

**Key Concept:** Pass loop variables as goroutine parameters to avoid closure capture bugs in older Go!

---

### Problem 9: Goroutine Leak ⭐⭐
```go
func handleClient(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        fmt.Print(msg)
    }
}
```
**Question:** When does this goroutine end? What happens if you never close the connection?

**Key Concept:** Always close `conn` to unblock goroutines waiting on it — leaking goroutines are silent killers!

---

### Problem 10: How Many Goroutines? ⭐⭐
```go
func main() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go handleClient(conn)
    }
}
```
**Question:** If 500 clients connect simultaneously, how many goroutines are running (excluding main)?

**Key Concept:** One goroutine per connection is idiomatic in Go — goroutines are cheap, not OS threads!

---

## BLOCK 3 — sync.Mutex Lock/Unlock

### Problem 11: Basic Mutex Usage ⭐
```go
var mu sync.Mutex
count := 0

mu.Lock()
count++
mu.Unlock()
```
**Question:** What does `Lock()` do if another goroutine already holds the lock? What does `Unlock()` do?

**Key Concept:** `Lock` blocks until free. `Unlock` releases and wakes a waiter. Never need to initialize a Mutex!

---

### Problem 12: defer mu.Unlock() ⭐
```go
var mu sync.Mutex

func safeUpdate(data map[string]int, key string) {
    mu.Lock()
    defer mu.Unlock()
    data[key]++
}
```
**Question:** Why is `defer mu.Unlock()` better than calling `mu.Unlock()` manually at the end?

**Key Concept:** Always `defer mu.Unlock()` right after `mu.Lock()` — it protects against panics and early returns!

---

### Problem 13: Double Lock Deadlock ⭐⭐
```go
var mu sync.Mutex

func main() {
    mu.Lock()
    fmt.Println("first lock")
    mu.Lock()
    fmt.Println("second lock")
    mu.Unlock()
    mu.Unlock()
}
```
**Question:** What happens when this runs?

**Key Concept:** `sync.Mutex` is NOT reentrant — a goroutine locking it twice deadlocks itself!

---

### Problem 14: Mutex Protects Which Data? ⭐⭐
```go
type Server struct {
    mu      sync.Mutex
    clients []net.Conn
}

func (s *Server) Add(conn net.Conn) {
    s.mu.Lock()
    s.clients = append(s.clients, conn)
    s.mu.Unlock()
}

func (s *Server) Count() int {
    return len(s.clients)  // ← no lock!
}
```
**Question:** Is `Count()` safe to call from multiple goroutines? Why or why not?

**Key Concept:** EVERY access to shared data — reads AND writes — must be protected by the mutex!

---

### Problem 15: Unlock Before Long Operation ⭐⭐⭐
```go
func (s *Server) broadcast(msg string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for _, conn := range s.clients {
        conn.Write([]byte(msg))  // ← network write inside lock!
    }
}
```
**Question:** What is the performance problem with holding the mutex during `conn.Write`?

**Key Concept:** Never do slow I/O while holding a mutex — copy shared data, release the lock, then operate!

---

## BLOCK 4 — Race Conditions

### Problem 16: What is a Race Condition? ⭐
```go
count := 0
go func() { count++ }()
go func() { count++ }()
time.Sleep(time.Millisecond)
fmt.Println(count)
```
**Question:** What are the possible values of `count`? Why is this dangerous?

**Key Concept:** Unsynchronized concurrent reads+writes = race condition = undefined, unpredictable results!

---

### Problem 17: Detecting Races ⭐
```go
// Run with: go run -race main.go
```
**Question:** What does Go's race detector do? Does it fix the race or just find it?

**Key Concept:** `go run -race` detects races at runtime — always use it during development of concurrent code!

---

### Problem 18: Map Race ⭐⭐
```go
m := map[string]int{}
go func() { m["a"] = 1 }()
go func() { m["b"] = 2 }()
time.Sleep(time.Millisecond)
```
**Question:** Is this safe? What makes maps special in Go regarding concurrency?

**Key Concept:** Go maps panic on concurrent writes — always protect maps with a mutex!

---

### Problem 19: The Slice Race ⭐⭐
```go
type Server struct {
    mu      sync.Mutex
    clients []net.Conn
}

func (s *Server) Remove(conn net.Conn) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for i, c := range s.clients {
        if c == conn {
            s.clients = append(s.clients[:i], s.clients[i+1:]...)
            return
        }
    }
}
```
**Question:** Is this safe? What would happen if you removed the mutex here?

**Key Concept:** Slice append modifies the header AND backing array — always mutex-protect concurrent slice modifications!

---

### Problem 20: Read + Write Race ⭐⭐⭐
```go
var mu sync.Mutex
clients := []net.Conn{}

// Goroutine A — writer
go func() {
    mu.Lock()
    clients = append(clients, newConn)
    mu.Unlock()
}()

// Goroutine B — reader (NO lock!)
go func() {
    fmt.Println(len(clients))
}()
```
**Question:** Is Goroutine B safe? Does reading require a lock if only one goroutine writes?

**Key Concept:** Reads are NOT free in concurrent code — if there's any writer, ALL accesses need a lock!

---

## BLOCK 5 — Broadcast Architecture

### Problem 21: What is a Broadcast? ⭐
**Question:** In a chat server, what does "broadcast" mean? How is it different from sending to one client?

**Key Concept:** Broadcast = fan-out — one message in, N messages out (one per connected client)!

---

### Problem 22: Basic Broadcast Function ⭐⭐
```go
func broadcast(clients []net.Conn, msg string) {
    for _, conn := range clients {
        conn.Write([]byte(msg))
    }
}
```
**Question:** What happens if one client's `Write` blocks (e.g. slow network)? How does it affect other clients?

**Key Concept:** Synchronous broadcast blocks on slow clients — in production, write to each client concurrently!

---

### Problem 23: Skipping the Sender ⭐⭐
```go
func (s *Server) broadcast(sender net.Conn, msg string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for _, conn := range s.clients {
        if conn != sender {
            conn.Write([]byte(msg))
        }
    }
}
```
**Question:** What does `conn != sender` accomplish? When would you skip this check?

**Key Concept:** `conn != sender` skips the originator — compare `net.Conn` interfaces directly with `==`!

---

### Problem 24: Adding and Removing Clients ⭐⭐
```go
func (s *Server) join(conn net.Conn) {
    s.mu.Lock()
    s.clients = append(s.clients, conn)
    s.mu.Unlock()
}

func (s *Server) leave(conn net.Conn) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for i, c := range s.clients {
        if c == conn {
            s.clients = append(s.clients[:i], s.clients[i+1:]...)
            return
        }
    }
}
```
**Question:** In `leave`, why is it safe to `return` immediately after removing? Does the range loop need to continue?

**Key Concept:** Remove-by-value with early return is safe when each item appears exactly once — don't continue iterating after removal!

---

### Problem 25: Broadcast on Disconnect ⭐⭐⭐
```go
func (s *Server) handleClient(conn net.Conn) {
    s.join(conn)
    defer func() {
        s.leave(conn)
        conn.Close()
        s.broadcast(nil, conn.RemoteAddr().String()+" has left\n")
    }()

    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        s.broadcast(conn, msg)
    }
}
```
**Question:** What is the order of operations when a client disconnects? Is there a potential issue with the broadcast after `leave`?

**Key Concept:** Order in `defer` matters — remove first, then close, then notify others!

---

## BLOCK 6 — bufio.NewReader & ReadString

### Problem 26: Why bufio? ⭐
```go
// Without bufio:
buf := make([]byte, 1024)
n, err := conn.Read(buf)

// With bufio:
reader := bufio.NewReader(conn)
line, err := reader.ReadString('\n')
```
**Question:** What problem does `bufio.NewReader` solve that raw `conn.Read` doesn't?

**Key Concept:** `bufio.NewReader` + `ReadString('\n')` = read one complete line at a time, no manual scanning!

---

### Problem 27: ReadString Return Value ⭐
```go
reader := bufio.NewReader(conn)
line, err := reader.ReadString('\n')
fmt.Printf("%q\n", line)
```
Assume the client sent `"hello\n"`.

**Question:** What does `line` contain exactly? What is the value of `err` if read was successful?

**Key Concept:** `ReadString('\n')` includes the `'\n'` in the result — always `TrimSpace` before processing!

---

### Problem 28: TrimSpace After ReadString ⭐⭐
```go
reader := bufio.NewReader(conn)
msg, err := reader.ReadString('\n')
if err != nil {
    return
}
msg = strings.TrimSpace(msg)
fmt.Println(len(msg))
```
Client sends `"hello\r\n"` (Windows line ending).

**Question:** What does `len(msg)` print? What does `TrimSpace` remove?

**Key Concept:** Always `TrimSpace` after `ReadString` — it strips `\n`, `\r\n`, and leading/trailing spaces!

---

### Problem 29: ReadString on Disconnect ⭐⭐
```go
for {
    msg, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("client disconnected:", err)
        return
    }
    broadcast(msg)
}
```
**Question:** What error does `ReadString` return when the client closes the connection? Is `msg` usable when `err != nil`?

**Key Concept:** `io.EOF` from `ReadString` = client disconnected. Check error first — partial `msg` on error is unreliable!

---

## BLOCK 7 — net.Conn Read/Write

### Problem 30: Writing to a Connection ⭐⭐
```go
func sendMessage(conn net.Conn, msg string) error {
    _, err := conn.Write([]byte(msg))
    return err
}
```
**Question:** What does `conn.Write` return? Why might you ignore the first return value but not the second?

**Key Concept:** Never ignore `Write` errors — a non-nil error means the connection is dead!

---

### Problem 31: Write Failure Handling in Broadcast ⭐⭐⭐
```go
func (s *Server) broadcast(sender net.Conn, msg string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for _, conn := range s.clients {
        if conn == sender {
            continue
        }
        if _, err := conn.Write([]byte(msg)); err != nil {
            conn.Close()
            // Can we call s.leave(conn) here?
        }
    }
}
```
**Question:** Why is calling `s.leave(conn)` inside `broadcast` while holding the lock dangerous?

**Key Concept:** Never call a function that locks the same mutex while already holding it — deadlock!

---

## BLOCK 8 — Time Formatting

### Problem 32: time.Now and Format ⭐
```go
t := time.Now()
fmt.Println(t.Format("2006-01-02 15:04:05"))
```
**Question:** What is special about `"2006-01-02 15:04:05"` in Go? What does it represent?

**Key Concept:** Go's reference time is `2006-01-02 15:04:05` — format by example, not `%Y%m%d`!

---

### Problem 33: Stamping a Message ⭐⭐
```go
func stamp(msg string) string {
    t := time.Now()
    return fmt.Sprintf("[%s] %s", t.Format("2006-01-02 15:04:05"), msg)
}
fmt.Println(stamp("hello"))
```
**Question:** What does the output look like? What does `Sprintf` return here?

**Key Concept:** Wrap `time.Now().Format(...)` in `Sprintf` to inject timestamps into any message string!

---

### Problem 34: Wrong Format String ⭐⭐
```go
t := time.Now()
fmt.Println(t.Format("2024-01-02 15:04:05"))
fmt.Println(t.Format("2006-01-02 15:04:05"))
```
**Question:** What is wrong with the first line? What does it actually print?

**Key Concept:** The reference year must be EXACTLY `2006` — using any other year breaks year formatting silently!

---

### Problem 35: Time Zones ⭐⭐⭐
```go
t := time.Now()
fmt.Println(t.Format("15:04:05 MST"))
fmt.Println(t.UTC().Format("15:04:05 MST"))
```
**Question:** What is the difference between these two lines? When would you use UTC for a chat server?

**Key Concept:** Use `t.UTC()` for multi-timezone systems — local time causes confusion when clients span the globe!

---

## 🏆 Quick Reference Card

| Topic | Key Rule |
|-------|----------|
| `net.Listen(":8080")` | Binds to all interfaces on port 8080 |
| `ln.Accept()` | Blocks until a client connects — returns `net.Conn` |
| Accept without `go` | Sequential — only one client at a time |
| `go handleClient(conn)` | One goroutine per client — idiomatic Go |
| Goroutine order | Never guaranteed — scheduler decides |
| `go` without sync | main may exit before goroutine runs |
| `mu.Lock()` | Blocks until lock is free |
| `defer mu.Unlock()` | Guarantees unlock even on panic |
| Double `Lock()` | Deadlock — mutex is NOT reentrant |
| Reads need locks too | If any goroutine writes, ALL must lock |
| Map concurrent write | Panics — always protect maps with mutex |
| Broadcast = fan-out | One message → write to every `net.Conn` |
| Skip sender | `conn != sender` in broadcast loop |
| Slow client in broadcast | Blocks entire loop — copy list, unlock, then write |
| `bufio.NewReader` | Buffers TCP stream for line-by-line reading |
| `ReadString('\n')` | Includes `'\n'` in result — always `TrimSpace`! |
| `io.EOF` | Client disconnected cleanly |
| `conn.Write` error | Never ignore — means connection is dead |
| Lock + nested lock | Deadlock — never call locking func while locked |
| `time.Format` | By example: `"2006-01-02 15:04:05"` |
| Reference year | Must be exactly `2006` — not `2024`! |
| `t.UTC()` | Use for multi-timezone chat servers |

**Go build that chat server! 💪🔥**
