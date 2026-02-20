# 🎯 Net-Cat Prerequisites Quiz
## TCP Networking · Goroutines · Mutex · bufio · Broadcast Architecture · Connection Lifecycle

**Time Limit:** 55 minutes  
**Total Questions:** 30  
**Passing Score:** 24/30 (80%)

> ✅ Pass → You're ready to start Net-Cat  
> ⚠️ This project has no web browser, no HTTP, no templates — just raw TCP and concurrent Go. A lower score here means real pain. If you score below 20, stop and study first.

---

## 📋 SECTION 1: TCP NETWORKING (6 Questions)

### Q1: What is the difference between a TCP server and a TCP client?

**A)** Servers are faster than clients  
**B)** A server **listens** on a port and **accepts** incoming connections; a client **dials** out to a server's address and port  
**C)** A client can only receive; a server can only send  
**D)** TCP clients require a browser  

<details><summary>💡 Answer</summary>

**B) Server listens + accepts; client dials**

```
Server: net.Listen("tcp", ":8989") → listener.Accept() → net.Conn
Client: net.Dial("tcp", "localhost:8989") → net.Conn
```

`nc localhost 8989` is the client. Your Go program is the server. Both sides get a `net.Conn` after the connection is established — the interface is symmetric for reading and writing.

</details>

---

### Q2: What does `net.Listen("tcp", ":8989")` return, and what happens next?

**A)** A `net.Conn` — you can immediately read messages from it  
**B)** A `net.Listener` — you must call `Accept()` in a loop to receive individual connections  
**C)** A `[]net.Conn` — all connections that will ever arrive  
**D)** Nothing useful — you must use `http.ListenAndServe` instead  

<details><summary>💡 Answer</summary>

**B) A `net.Listener` — then loop calling `Accept()`**

```go
ln, err := net.Listen("tcp", ":8989")
if err != nil { log.Fatal(err) }
defer ln.Close()

for {
    conn, err := ln.Accept()  // blocks until a client connects
    if err != nil { log.Println(err); continue }
    go handleClient(conn)     // handle in goroutine, don't block
}
```

`Accept()` blocks until a client connects. Each call returns one `net.Conn`. The loop keeps accepting new connections indefinitely.

</details>

---

### Q3: What does `listener.Accept()` do when there are no clients connected?

**A)** Returns an error immediately  
**B)** Returns `nil, nil`  
**C)** Blocks (sleeps) until a new client connects  
**D)** Panics  

<details><summary>💡 Answer</summary>

**C) Blocks until a new client connects**

`Accept()` is a blocking call — it parks the goroutine until a connection arrives. This is correct behavior. Your accept loop blocks here doing nothing until someone runs `nc localhost 8989`, at which point it returns immediately with the new connection.

</details>

---

### Q4: Why must each client connection be handled in a separate goroutine?

**A)** TCP requires one goroutine per connection  
**B)** Without goroutines, each `handleClient` call would block the accept loop — new clients would be unable to connect until the current one disconnects  
**C)** Goroutines are faster  
**D)** `net.Conn` can only be used inside a goroutine  

<details><summary>💡 Answer</summary>

**B) Without goroutines, the accept loop would block**

```go
// WRONG — while client A is connected, no new clients can connect:
conn, _ := ln.Accept()
handleClient(conn)  // blocks here until A disconnects

// CORRECT — accept loop keeps running:
conn, _ := ln.Accept()
go handleClient(conn)  // non-blocking; loop continues immediately
```

This is the single most important design decision in the project. Every client must get its own goroutine.

</details>

---

### Q5: How do you write a string to a `net.Conn`?

**A)** `conn.Print("message")`  
**B)** `conn.Write([]byte("message"))`  
**C)** `conn.Send("message")`  
**D)** `fmt.Println(conn, "message")`  

<details><summary>💡 Answer</summary>

**B) `conn.Write([]byte("message"))`**

`net.Conn` implements `io.ReadWriter`. `Write` expects `[]byte`. Common patterns:

```go
conn.Write([]byte("Welcome!\n"))
fmt.Fprintf(conn, "[%s]: %s\n", name, message)  // also works
```

`fmt.Fprintf(conn, ...)` is convenient because it formats and writes in one call — it uses `conn`'s `Write` method internally.

</details>

---

### Q6: How do you read input from a client line by line over a `net.Conn`?

**A)** `conn.ReadLine()`  
**B)** `bufio.NewReader(conn).ReadString('\n')`  
**C)** `io.ReadAll(conn)`  
**D)** `conn.Scan()`  

<details><summary>💡 Answer</summary>

**B) `bufio.NewReader(conn).ReadString('\n')`**

```go
reader := bufio.NewReader(conn)
for {
    line, err := reader.ReadString('\n')
    if err != nil {
        // client disconnected
        return
    }
    line = strings.TrimSpace(line)
    // process line
}
```

`ReadString('\n')` reads bytes until it hits a newline, returning the complete line including the delimiter. `io.ReadAll` would block until the connection closes — unusable for interactive chat.

</details>

---

## 📋 SECTION 2: GOROUTINES & MUTEX (8 Questions)

### Q7: Two goroutines both read and write to the same `[]Client` slice without any protection. What can go wrong?

**A)** Nothing — Go slices handle concurrent access automatically  
**B)** A race condition: one goroutine might read stale data, corrupt the slice header, or cause a panic — the outcome is undefined  
**C)** The second write is automatically queued  
**D)** The program slows down but stays correct  

<details><summary>💡 Answer</summary>

**B) Race condition — undefined behavior, potential panic**

Go slices are not thread-safe. Concurrent writes can corrupt the slice's internal length/capacity/pointer fields, causing panics or silent data corruption. The Go race detector (`-race`) will catch this and report exactly where it happens.

</details>

---

### Q8: What is the correct pattern for protecting a shared variable with `sync.Mutex`?

**A)**
```go
mu.Lock()
data = append(data, item)
// (forget to unlock — deadlock risk)
```
**B)**
```go
mu.Lock()
defer mu.Unlock()
data = append(data, item)
```
**C)**
```go
go mu.Lock()
data = append(data, item)
mu.Unlock()
```
**D)**
```go
mu.Lock(); mu.Unlock()
data = append(data, item)
```

<details><summary>💡 Answer</summary>

**B) `mu.Lock()` then `defer mu.Unlock()`**

`defer mu.Unlock()` placed immediately after `Lock()` guarantees the unlock runs when the function returns — even if it panics or returns early. Option D locks and immediately unlocks before touching the data — useless. Never forget to unlock.

</details>

---

### Q9: What is a deadlock and how does it happen with a mutex in this project?

**A)** When the server runs out of memory  
**B)** When goroutine A holds the lock and calls a function that also tries to lock — it waits forever for itself  
**C)** When two clients connect at the same time  
**D)** When a goroutine exits without calling `Done()` on a WaitGroup  

<details><summary>💡 Answer</summary>

**B) A goroutine tries to lock a mutex it already holds**

```go
// DEADLOCK:
func broadcast(...) {
    mu.Lock()
    defer mu.Unlock()
    // ...
    sendHistory(conn)  // sendHistory also calls mu.Lock() — DEADLOCK
}

// FIX: don't call locking functions from inside a lock:
func broadcast(...) {
    mu.Lock()
    defer mu.Unlock()
    // send history inline, not via a function that re-locks
}
```

`sync.Mutex` is NOT reentrant. If goroutine A holds the lock and tries to lock again, it waits forever. Go detects full deadlocks and panics: "all goroutines are asleep - deadlock!"

</details>

---

### Q10: Your `broadcast` function must: (1) append the message to history, (2) send it to all other clients. Must these happen inside the same mutex lock? Why?

**A)** No — history and client list can be locked separately  
**B)** Yes — locking once for both prevents a message from being added to history but not delivered (or delivered but not added to history) if another goroutine modifies the list between two separate locks  
**C)** Neither needs a lock  
**D)** Only appending to history needs a lock  

<details><summary>💡 Answer</summary>

**B) One lock for both — atomicity**

If you used two separate locks (lock → append → unlock → lock → send → unlock), a new client could join between the two unlocks and miss the message in history OR receive it twice. A single lock makes the "append + send" atomic from the perspective of all other goroutines.

</details>

---

### Q11: Should the sender receive their own broadcast? What does the spec say, and how do you implement it?

**A)** Yes — everyone including the sender receives the message  
**B)** No — skip the sender when iterating the clients slice: `if client != sender { ... }`  
**C)** The sender disconnects after sending  
**D)** The spec doesn't specify  

<details><summary>💡 Answer</summary>

**B) Skip the sender — compare client pointers**

```go
for _, c := range clients {
    if c != sender {  // pointer comparison
        c.conn.Write([]byte(message))
    }
}
```

The sender sees their own message via the prompt reprint, not via broadcast. If you send the message back to the sender, they'd see it twice — once as a broadcast and once via their reprinted prompt.

</details>

---

### Q12: What is `sync.Mutex` zero value? Do you need to initialize it before using it?

**A)** `nil` — you must call `sync.NewMutex()`  
**B)** An unlocked mutex — you can declare `var mu sync.Mutex` and use it immediately without initialization  
**C)** A locked mutex — you must call `mu.Unlock()` before first use  
**D)** It has no zero value — must be heap allocated  

<details><summary>💡 Answer</summary>

**B) Zero value is an unlocked mutex — no initialization needed**

```go
var mu sync.Mutex  // ready to use immediately
// or as a struct field:
type Server struct {
    mu sync.Mutex  // zero value = unlocked, correct
}
```

This is a deliberate Go design choice. Never copy a `sync.Mutex` after first use — pass by pointer.

</details>

---

### Q13: How do you remove a specific client from a `[]*Client` slice when you know its pointer?

**A)** `clients = clients[1:]` — removes the first  
**B)** Find the index by comparing pointers, then `clients = append(clients[:i], clients[i+1:]...)`  
**C)** `delete(clients, client)` — Go built-in  
**D)** Set the element to `nil`  

<details><summary>💡 Answer</summary>

**B) Find index by pointer comparison, then splice**

```go
func removeClient(client *Client) {
    mu.Lock()
    defer mu.Unlock()
    for i, c := range clients {
        if c == client {
            clients = append(clients[:i], clients[i+1:]...)
            break
        }
    }
    client.conn.Close()
}
```

`append(s[:i], s[i+1:]...)` removes the element at index `i` without leaving a gap. Compare by pointer (`c == client`) since two clients could theoretically have the same name.

</details>

---

### Q14: `go run -race . 2525` starts the server. Two clients connect and chat. The race detector reports a race on `clients`. What does this mean and how do you fix it?

**A)** The race detector is wrong — ignore it  
**B)** Two goroutines are accessing `clients` concurrently without the mutex — find every read/write to `clients` that isn't inside `mu.Lock()/mu.Unlock()`  
**C)** Use a different data structure  
**D)** Add `time.Sleep` before accessing `clients`  

<details><summary>💡 Answer</summary>

**B) Find every unprotected access to `clients`**

The race detector reports the exact file and line numbers. Common mistakes:
- Reading `len(clients)` without a lock (connection limit check)
- Iterating `clients` in `broadcast` without a lock
- Appending to `clients` in `handleClient` without a lock

Every read AND write to shared variables must be inside a mutex lock.

</details>

---

## 📋 SECTION 3: bufio & I/O PATTERNS (5 Questions)

### Q15: What is `bufio.NewReader` and why is it needed for reading from a `net.Conn`?

**A)** It compresses the data  
**B)** It wraps an `io.Reader` with an internal buffer — allows reading line-by-line efficiently without making one syscall per byte  
**C)** It encrypts the connection  
**D)** It's required to use `net.Conn` at all  

<details><summary>💡 Answer</summary>

**B) Wraps a reader with a buffer — enables efficient line reading**

`net.Conn.Read()` reads whatever bytes are available. Reading one byte at a time to find a newline would be 80+ syscalls for a short message. `bufio.NewReader` buffers a larger chunk (4096 bytes by default) and `ReadString('\n')` scans it for the newline character without extra syscalls.

</details>

---

### Q16: `reader.ReadString('\n')` returns `line, err`. When `err != nil`, what does `line` contain?

**A)** Always empty  
**B)** Any bytes that were read before the error — partial data before the connection dropped  
**C)** The error message  
**D)** `nil`  

<details><summary>💡 Answer</summary>

**B) Partial data that was read before the error**

When a client disconnects mid-message, `ReadString` may return a partial line AND a non-nil error. Always check `err` FIRST. If there's an error, treat it as a disconnect regardless of what `line` contains — don't process partial messages.

```go
line, err := reader.ReadString('\n')
if err != nil {
    // client disconnected — remove and return
    return
}
```

</details>

---

### Q17: Why do you call `strings.TrimSpace(line)` after `ReadString('\n')`?

**A)** To make the string lowercase  
**B)** `ReadString` includes the `\n` delimiter in the returned string — `TrimSpace` removes the trailing newline (and any `\r` from Windows clients)  
**C)** To check for empty messages  
**D)** Both B and C  

<details><summary>💡 Answer</summary>

**D) Both — removes the newline delimiter AND enables the empty check**

`ReadString('\n')` returns `"hello\n"` or `"hello\r\n"` (Windows). After `TrimSpace`: `"hello"`. Then you can check `if line == "" { /* ignore */ }`. Without trimming, `"hello\n" == ""` is always false — empty messages would never be caught.

</details>

---

### Q18: A client using Windows `nc` sends `"hello\r\n"`. After `ReadString('\n')`, what do you get, and what does `TrimSpace` do?

**A)** `"hello"` — `ReadString` strips both `\r` and `\n`  
**B)** `"hello\r\n"` → `TrimSpace` → `"hello"` — `TrimSpace` removes all leading/trailing whitespace including `\r`  
**C)** An error — Windows line endings are not supported  
**D)** `"hello\r"` and `TrimSpace` doesn't remove `\r`  

<details><summary>💡 Answer</summary>

**B) `"hello\r\n"` → `TrimSpace` → `"hello"`**

`strings.TrimSpace` removes `\t`, `\n`, `\r`, `\v`, `\f`, and space from both ends. This handles cross-platform differences. If you only trim `\n` (e.g. `strings.TrimRight(line, "\n")`), Windows clients would have a dangling `\r` in their messages — appearing as strange characters in terminals.

</details>

---

### Q19: After broadcasting a message to all other clients, you need to "reprint the prompt" for the sender. What is `\r` and why is it needed?

**A)** A carriage return — moves the cursor to the start of the current line without advancing to the next line  
**B)** A newline — same as `\n`  
**C)** A tab character  
**D)** Not needed — terminals handle this automatically  

<details><summary>💡 Answer</summary>

**A) Carriage return — moves cursor to start of line without newline**

When a broadcast arrives while the user is typing, it appears on the current input line, corrupting the display. Sending `\r` before the broadcast moves the cursor to the start of the current line, overwriting any partial input. Then reprint `[NAME]: ` after to give them a fresh prompt.

```go
fmt.Fprintf(client.conn, "\r%s\n[%s]: ", message, client.name)
```

</details>

---

## 📋 SECTION 4: CONNECTION LIFECYCLE (5 Questions)

### Q20: What is the correct order for the join sequence?

**A)** Add to clients → get name → send history → announce join → start message loop  
**B)** Get name → add to clients → send history → announce join → start message loop  
**C)** Announce join → get name → add to clients → send history → start message loop  
**D)** Start message loop → get name → add to clients → send history → announce join  

<details><summary>💡 Answer</summary>

**B) Get name → add to clients → send history → announce join → start message loop**

The order matters:
1. **Get name first** — you need the name for all subsequent steps
2. **Add to clients** — must happen before sending history (so new messages during history delivery are handled) but before announcing (so the announcement can include them)
3. **Send history** — new client sees all past messages before fresh ones
4. **Announce join** — now everyone (including the new client) knows they joined
5. **Start message loop** — now accept their messages

</details>

---

### Q21: A client connects but refuses to enter a name — they just keep pressing Enter. What should your server do?

**A)** Accept the client with an empty name  
**B)** Disconnect after 3 retries  
**C)** Keep re-sending the `[ENTER YOUR NAME]:` prompt and never proceed until a non-empty name is given  
**D)** Use "Anonymous" as the default name  

<details><summary>💡 Answer</summary>

**C) Keep re-prompting until a non-empty name is given**

```go
for {
    conn.Write([]byte("[ENTER YOUR NAME]: "))
    name, err := reader.ReadString('\n')
    if err != nil { return "" }  // client disconnected
    name = strings.TrimSpace(name)
    if name != "" { return name }
    // loop: re-prompt
}
```

Empty names would cause formatting issues and confusion. Always validate before proceeding.

</details>

---

### Q22: What should happen when the 11th client tries to connect (max is 10)?

**A)** The server crashes  
**B)** Send a message like "Chat is full" to the 11th client, then close their connection — all existing clients are unaffected  
**C)** Disconnect the oldest client to make room  
**D)** Queue the client until a spot opens  

<details><summary>💡 Answer</summary>

**B) Send rejection message, close connection, leave existing clients alone**

```go
mu.Lock()
if len(clients) >= maxClients {
    mu.Unlock()
    conn.Write([]byte("Chat is full. Try again later.\n"))
    conn.Close()
    return
}
mu.Unlock()
```

Note: unlock BEFORE closing — don't hold the mutex across I/O operations. The check-and-reject must be atomic with respect to other goroutines adding clients.

</details>

---

### Q23: Client B disconnects while client A is in the middle of `broadcast()` sending to B. What happens to `conn.Write([]byte(msg))` for B?

**A)** The server panics  
**B)** `conn.Write` returns a non-nil error — you can ignore it or log it; the error does NOT affect other clients in the loop  
**C)** The server blocks waiting for B to reconnect  
**D)** The write succeeds silently  

<details><summary>💡 Answer</summary>

**B) `conn.Write` returns an error — safe to ignore in broadcast**

Writing to a closed connection returns an error (typically "broken pipe" or "connection reset"). In `broadcast`, you're iterating all clients — one failed write should NOT stop the loop or crash. Log it if you want, but continue to the next client. The actual cleanup (removing B from the list) is handled when B's own `listenToClient` goroutine detects the disconnect.

</details>

---

### Q24: Your `removeClient` function is called. What are the two things it must do?

**A)** Close the connection only  
**B)** Remove the client from the `clients` slice (inside a mutex lock) AND close the connection  
**C)** Remove from slice, close connection, AND send a farewell message  
**D)** Only remove from the slice — the connection closes itself  

<details><summary>💡 Answer</summary>

**B) Remove from slice (with mutex) AND close the connection**

The broadcast (farewell message) happens in `listenToClient` BEFORE calling `removeClient` — by the time you remove the client, the message is already sent. `removeClient` itself only handles cleanup: remove the pointer from the slice and close the TCP connection.

</details>

---

## 📋 SECTION 5: TIME, FORMATTING & MISC (4 Questions)

### Q25: How do you format the current time as `2020-01-20 16:03:43` in Go?

**A)** `time.Now().Format("YYYY-MM-DD HH:MM:SS")`  
**B)** `time.Now().Format("2006-01-02 15:04:05")`  
**C)** `time.Now().Format("%Y-%m-%d %H:%M:%S")`  
**D)** `fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", ...)`  

<details><summary>💡 Answer</summary>

**B) `time.Now().Format("2006-01-02 15:04:05")`**

Go uses a reference time of `Mon Jan 2 15:04:05 MST 2006` as its format template — not `YYYY/MM/DD`. The numbers are fixed: `2006` = year, `01` = month, `02` = day, `15` = hour (24h), `04` = minute, `05` = second. This is one of Go's most infamous quirks — memorize the reference time.

</details>

---

### Q26: A correctly formatted message looks like `[2020-01-20 16:03:43][Alice]:Hello world`. What is the Go format string that produces this?

**A)** `fmt.Sprintf("[%s][%s]:%s\n", timestamp, name, message)`  
**B)** `"[" + timestamp + "][" + name + "]:" + message + "\n"`  
**C)** `fmt.Sprintf("[%s][%s]: %s\n", timestamp, name, message)` (with a space before message)  
**D)** Both A and B produce identical output  

<details><summary>💡 Answer</summary>

**A) `fmt.Sprintf("[%s][%s]:%s\n", timestamp, name, message)`**

Note carefully: `[timestamp][name]:message` — the colon is immediately followed by the message with NO space. Check the spec's exact format. Option C adds a space after `:` which is wrong. Both A and B produce the same string — `fmt.Sprintf` is just cleaner.

</details>

---

### Q27: Your `os.Args` parsing: `go run . 2525 extra` should print the usage message. `go run . 2525` should use port 2525. `go run .` should use 8989. How do you implement this?

**A)**
```go
if len(os.Args) == 1 { port = "8989" }
if len(os.Args) == 2 { port = os.Args[1] }
if len(os.Args) > 2 { fmt.Println("[USAGE]: ./TCPChat $port"); return }
```
**B)**
```go
port = os.Args[1]
```
**C)**
```go
if len(os.Args) != 2 { port = "8989" }
```
**D)**
```go
port = os.Args[0]
```

<details><summary>💡 Answer</summary>

**A)**

`os.Args[0]` is always the program name. `os.Args[1]` is the first user argument. So:
- `go run .` → `len(os.Args) == 1` → use default `"8989"`
- `go run . 2525` → `len(os.Args) == 2` → `os.Args[1] = "2525"`
- `go run . 2525 extra` → `len(os.Args) == 3` → usage message

Option D uses `os.Args[0]` which is the binary name, not a port.

</details>

---

### Q28: A new client joins. The spec says they should receive all **past** messages. When exactly during the join sequence should this happen, and from where should the messages come?

**A)** Before getting their name, from a file on disk  
**B)** After getting their name and being added to the client list, from the in-memory `messages` slice (protected by mutex)  
**C)** After the first message is sent by any client  
**D)** Only if they ask for it  

<details><summary>💡 Answer</summary>

**B) After adding to client list, from the in-memory `messages` slice with mutex**

```go
func sendHistory(conn net.Conn) {
    mu.Lock()
    defer mu.Unlock()
    for _, msg := range messages {
        conn.Write([]byte(msg))
    }
}
```

History is sent before the join announcement — the new client sees past messages, then the "X has joined" message appears as the first fresh message they receive.

</details>

---

## 📋 SECTION 6: INTEGRATION & EDGE CASES (2 Questions)

### Q29: You have clients A and B chatting. Client C joins. In what order do things happen, and what does each client see?

**A)** C joins silently; A and B see nothing until C sends a message  
**B)** A and B see "C has joined our chat..." (join announcement); C sees the full message history then "C has joined our chat..." should NOT appear in C's history since it happens after  
**C)** C sees history, then the join announcement goes to A and B; C sees the announcement too since they're now in the client list  
**D)** Nothing — only the server logs the join  

<details><summary>💡 Answer</summary>

**C) History → announcement to A and B (and C, since they're now in the list)**

The sequence:
1. C is added to `clients` slice
2. History (past messages) is sent only to C
3. Broadcast "C has joined" — goes to everyone in `clients` **including C**

This is subtle: C sees their own join announcement because they were added before the broadcast. The spec may vary on whether C should see their own join message — read it carefully.

</details>

---

### Q30: Describe the complete sequence of events when Client B, who is actively chatting with Client A, closes their terminal.

**A)** The server detects nothing — B just stops sending  
**B)** 1) B's `bufio.ReadString` returns an error; 2) `removeClient(B)` is called — removes B from slice and closes connection; 3) Broadcast "B has left our chat..."; 4) A sees the message; A's goroutine continues running normally  
**C)** The server crashes  
**D)** A is also disconnected  

<details><summary>💡 Answer</summary>

**B) Read error → remove → broadcast leave message → A continues**

The exact order: `listenToClient(B)` gets a read error → broadcast the leave message (BEFORE removing B so the message has B's name) → call `removeClient(B)` → B's goroutine exits. A's separate goroutine is completely unaffected — the key benefit of the per-client goroutine architecture.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 28–30 ✅ | **Exceptional.** Deep networking + concurrency understanding — start immediately. |
| 24–27 ✅ | **Ready.** Review the sections you missed. Pay special attention to mutex deadlock patterns. |
| 18–23 ⚠️ | **Study first.** Goroutines and mutex are the core of this project — shaky foundations here mean race conditions and deadlocks you won't know how to debug. |
| Below 18 ❌ | **Not ready.** Work through goroutine and TCP examples on Go by Example before attempting this project. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q6 | `net.Listen`, `Accept`, `net.Conn.Write`, `bufio.NewReader`, client vs server roles |
| Q7–Q14 | `sync.Mutex`, deadlock, race conditions, `go run -race`, removing from slice |
| Q15–Q19 | `bufio.ReadString`, `TrimSpace`, `\r\n` handling, carriage return for prompt reprint |
| Q20–Q24 | Join sequence order, empty name rejection, connection limit, disconnect cleanup |
| Q25–Q28 | Go time format reference (`2006-01-02 15:04:05`), message format string, `os.Args`, history delivery |
| Q29–Q30 | Full join/disconnect lifecycle, goroutine isolation |