# Net-Cat Project Guide

> **Before you start:** Run `man nc` in your terminal and use `nc` to connect to a real server. You cannot recreate something you have never seen.

---

## Objectives

By completing this project you will learn:

1. **TCP Networking** — How a server listens for and accepts client connections over TCP
2. **Concurrency** — Using goroutines to handle many clients at the same time without blocking
3. **Mutual Exclusion** — Protecting shared data from race conditions using `sync.Mutex`
4. **Broadcast Architecture** — Designing a system that fans one message out to many recipients
5. **Connection Lifecycle** — Handling join, message, disconnect events cleanly
6. **Buffered I/O** — Reading line-by-line from a network stream using `bufio`
7. **Time Formatting** — Stamping messages with the correct timestamp format

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Go Basics
- Structs and methods
- Slices — appending and removing elements
- Goroutines (`go func()`)
- Error handling

### 2. TCP and Networking
- What a TCP connection is
- The difference between a server (listener) and a client (connector)
- What a port number is

### 3. Concurrency
- `go` keyword — launching a goroutine
- `sync.Mutex` — `Lock()` and `Unlock()`
- What a race condition is and why it is dangerous

### 4. I/O
- `bufio.NewReader` and `ReadString('\n')`
- `strings.TrimSpace`
- `net.Conn` — the interface for reading and writing over a network connection

**If any of these are unfamiliar, read about them before writing any code.**

- https://pkg.go.dev/net
- https://pkg.go.dev/sync
- https://pkg.go.dev/bufio
- Search: **"golang goroutines explained"**
- Search: **"what is a race condition"**

---

## Project Structure

```
net-cat/
├── main.go
├── server.go
├── client.go
└── go.mod
```

---

## Milestone 1 — Read the Port Argument

**Goal:**
```
go run .              → Listening on the port :8989
go run . 2525         → Listening on the port :2525
go run . 2525 extra   → [USAGE]: ./TCPChat $port
```

**Questions to answer:**
- How do you read command-line arguments in Go?
- What is the default port when no argument is given?
- What condition triggers the usage message?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Set default port to "8989"

    // 2. If exactly one argument was given: use it as the port

    // 3. If more than one argument was given: print usage and return

    // 4. Print "Listening on the port :PORT"

    // 5. Call startServer(port) — built in the next milestone
}
```

**Verify:**
- All three cases above produce the correct output before you write any server code.

---

## Milestone 2 — Start the TCP Server

**Goal:** The server starts, listens on the given port, and accepts connections. Each connection is handled in its own goroutine.

**Questions to answer:**
- What function in the `net` package starts a TCP listener?
- What does `listener.Accept()` do when no client is connected?
- Why must each connection be handled in a separate goroutine?

**Code Placeholder:**
```go
// server.go

func startServer(port string) {
    // 1. Call net.Listen to bind the server to the port
    //    Handle the error — if the port is in use, the server must exit cleanly

    // 2. Defer closing the listener

    // 3. Loop forever:
    //    - Accept the next incoming connection
    //    - If accept fails, log the error and continue (don't crash)
    //    - Launch handleClient(conn) in a new goroutine
}
```

**Resources:**
- https://pkg.go.dev/net#Listen
- https://gobyexample.com/tcp-server

**Verify:**
- Start the server and run `nc localhost 8989` in another terminal
- The connection should succeed and hang — no crash on either side

---

## Milestone 3 — Client Struct and Shared State

**Goal:** Define what a connected client looks like and create the shared variables all goroutines will access.

**Questions to answer:**
- What information does your program need to track for each connected client?
- Why do `clients` and `messages` need to be protected by a mutex?
- What is the maximum number of simultaneous connections allowed?

**Code Placeholder:**
```go
// client.go

type Client struct {
    // The network connection for this client
    // The client's chosen name
}
```

```go
// server.go

var (
    // Slice of all currently connected clients
    // Slice of all past messages (for history)
    // A sync.Mutex to protect the above two
    // A constant for maximum connections (10)
)
```

**Resources:**
- Search: **"golang sync Mutex example"**
- https://gobyexample.com/mutexes

---

## Milestone 4 — Welcome Banner and Name Prompt

**Goal:** When a client connects, send the Linux logo and prompt them for a name. Reject empty names and re-prompt.

```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
        ...
[ENTER YOUR NAME]:
```

**Questions to answer:**
- How do you send text to a client over a `net.Conn`?
- What type does `conn.Write` expect?
- How do you read what the client types back, line by line?
- What condition should cause the name prompt to repeat?

**Code Placeholder:**
```go
// client.go

var logo = `Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    ` + "`" + `    ` + "`" + `.       | ` + "`" + `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     ` + "`" + `-'       ` + "`" + `--' `

func welcomeClient(conn net.Conn) {
    // Send the logo + "[ENTER YOUR NAME]: " to the client
}

func getClientName(conn net.Conn) string {
    // Create a buffered reader for the connection

    // Loop:
    //   Read a line from the client
    //   If there is a read error, return empty string
    //   Trim whitespace from the line
    //   If the name is non-empty, return it
    //   Otherwise, re-send the prompt and loop again
}
```

**Verify:**
- `nc localhost 8989` shows the full logo and the name prompt
- Pressing Enter without a name re-prompts
- Typing a name and pressing Enter is accepted

---

## Milestone 5 — Connection Limit and Full Join Sequence

**Goal:** Reject clients beyond the maximum. For accepted clients, run the full join flow in the right order.

**Questions to answer:**
- Why must you check the connection count before adding the client to the slice?
- What is the correct order: send history before or after announcing the join to others?
- Why should the client be added to the slice before receiving history but the announce go out after?

**Code Placeholder:**
```go
// server.go

func handleClient(conn net.Conn) {
    // 1. Lock the mutex and check if len(clients) >= maxClients
    //    If full: send a message to the client, close the connection, unlock, return
    //    Unlock immediately after the check

    // 2. Send the welcome banner

    // 3. Get the client's name — loop until non-empty
    //    If the connection drops during this step, close and return

    // 4. Create the Client struct

    // 5. Lock → append to clients slice → unlock

    // 6. Send message history to the new client

    // 7. Broadcast "NAME has joined our chat..." to everyone else

    // 8. Start the message loop — call listenToClient(client)
}
```

---

## Milestone 6 — Message History and Broadcast

**Goal:** New clients receive all past messages on join. Every message sent is delivered to all other connected clients.

**Questions to answer:**
- Why should saving to history and broadcasting happen inside the same mutex lock?
- Should the sender receive their own broadcast? Why or why not?
- When sending history to a new client, why do you need to lock the mutex?

**Code Placeholder:**
```go
// server.go

func sendHistory(conn net.Conn) {
    // Lock the mutex
    // Send each message in the messages slice to conn
    // Unlock
}

func broadcast(message string, sender *Client) {
    // Lock the mutex

    // Append message to the messages slice

    // Loop through clients:
    //   If the client is not the sender, write the message to their connection

    // Unlock
}
```

---

## Milestone 7 — The Message Loop

**Goal:** After joining, each client continuously sends messages. Messages are formatted with a timestamp and the client's name. Empty messages are ignored. When a client disconnects, the rest are notified.

**Questions to answer:**
- What format string produces `2020-01-20 16:03:43` in Go?
- How do you detect that a client has disconnected?
- What should happen to the rest of the chat when one client disconnects?

**Code Placeholder:**
```go
// client.go

func listenToClient(client *Client) {
    // Create a buffered reader for client.conn

    // Loop forever:
    //   Read a line from the client
    //   If there is an error (disconnect):
    //     Remove the client from the clients slice
    //     Broadcast "NAME has left our chat..." with no sender
    //     Return

    //   Trim the message
    //   If the message is empty: re-send the prompt and continue

    //   Format: "[TIMESTAMP][NAME]:MESSAGE\n"
    //   Broadcast the formatted message (sender = client, so they don't get it back)
    //   Re-send the prompt to the client themselves
}

func removeClient(client *Client) {
    // Lock the mutex
    // Find the client in the clients slice by comparing pointers
    // Remove it using the append(s[:i], s[i+1:]...) pattern
    // Close the connection
    // Unlock
}
```

**Resources:**
- Search: **"golang time format 2006"**
- Search: **"golang remove element from slice"**

**Verify:**
- Connect two clients A and B
- A sends a message → B sees `[timestamp][A]:message`
- A sees their own prompt reprinted after sending
- Empty messages do not appear for anyone
- B disconnects → A sees "B has left our chat..."
- A remains connected

---

## Milestone 8 — Race Condition Check

**Goal:** The server passes the Go race detector with no warnings.

**Questions to answer:**
- Which shared variables could be accessed from multiple goroutines simultaneously?
- Is every access to `clients` and `messages` inside a mutex lock?

**Verify:**
```bash
go run -race . 2525
```
Connect several clients and send messages rapidly. Fix every race the detector reports before submission.

**Resource:** Search: **"golang race detector"**

---

## Debugging Checklist

Before asking for help, go through this:

- Does the server freeze completely? You likely have a deadlock — a function that locks the mutex is being called from code that already holds the lock. Never call `broadcast` while holding `mu`.
- Does `go run -race .` report a race? Find every unprotected access to `clients` or `messages`.
- Does a new client not see history? Check that `sendHistory` is called before `listenToClient` but after the client is added to the slice.
- Do messages appear on the wrong line or overwrite input? Print `\r` before a broadcast to the client, then reprint their prompt after.
- Does a dead client stay in the list? Check that `removeClient` is called on every error path in `listenToClient`.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `net` | Listen, accept, read, write over TCP | https://pkg.go.dev/net |
| `bufio` | Read lines from a network stream | https://pkg.go.dev/bufio |
| `sync` | Mutex for shared state | https://pkg.go.dev/sync |
| `time` | Format message timestamps | https://pkg.go.dev/time |
| `strings` | TrimSpace to clean user input | https://pkg.go.dev/strings |
| `fmt` | Format messages and prompts | https://pkg.go.dev/fmt |
| `log` | Log server-side errors | https://pkg.go.dev/log |
| `os` | Read command-line arguments | https://pkg.go.dev/os |

---

## Submission Checklist

- [ ] Default port 8989 when no argument given
- [ ] Usage message when more than one argument given
- [ ] Linux logo and name prompt sent to connecting clients
- [ ] Empty names rejected and re-prompted
- [ ] Maximum 10 connections enforced
- [ ] New client receives full message history on join
- [ ] All clients notified when someone joins
- [ ] All clients notified when someone leaves
- [ ] Empty messages not broadcast
- [ ] Messages formatted as `[timestamp][name]:message`
- [ ] Remaining clients unaffected when one disconnects
- [ ] Mutex protects all shared state
- [ ] Goroutines used for each client connection
- [ ] `go run -race .` reports no races
- [ ] Errors handled on server and client side
- [ ] Unit tests written and passing
