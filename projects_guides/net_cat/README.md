# Net-Cat Project Guide

> **Rule before you start:** If you are stuck, search first. Every resource link in this guide points to where the answer lives. Do not paste code from AI — you will not understand it under pressure, and you will not learn the skill.

---

## What You Are Building

A group chat system over TCP. One server, many clients. When a client sends a message, everyone else sees it. When someone joins or leaves, everyone is notified. It works with the standard `nc` (netcat) command — no special client needed.

---

## Before You Write a Single Line

Run this in your terminal:

```bash
man nc
```

Read it. Then answer these questions to yourself before moving on:

- What does `nc` do at its core?
- What is the difference between TCP and UDP?
- What does it mean to "listen" on a port vs "connect" to one?

If you cannot answer these, search: **"TCP vs UDP explained"** and **"what is a socket"**.

---

## Phase 1 — The Skeleton

### Checkpoint 1.1 — Project Structure

Create this layout manually:

```
net-cat/
├── main.go
├── server.go
├── client.go
└── go.mod
```

```bash
go mod init net-cat
```

**Why separate files?** `main.go` is your entry point only. `server.go` handles listening and accepting connections. `client.go` handles everything about a single connected client. Keeping them separate makes the code readable and testable.

---

### Checkpoint 1.2 — Reading the Port Argument

The program must behave like this:

```
$ go run .
Listening on the port :8989

$ go run . 2525
Listening on the port :2525

$ go run . 2525 localhost
[USAGE]: ./TCPChat $port
```

Fill in the blanks:

```go
// main.go
package main

import (
    "fmt"
    "os"
)

func main() {
    port := "8989"

    if len(os.Args) == __ {
        port = __________
    } else if len(os.Args) __ 2 {
        fmt.Println("__________")
        return
    }

    fmt.Println("Listening on the port :" + port)
}
```

**Verify before moving on:**
- `go run .` prints `Listening on the port :8989`
- `go run . 2525` prints `Listening on the port :2525`
- `go run . 2525 localhost` prints the usage message and exits

**Resources:**
- Search: **"golang os.Args"**
- Docs: https://pkg.go.dev/os#pkg-variables

---

## Phase 2 — Starting the Server

### Checkpoint 2.1 — Concept: How TCP Servers Work

A TCP server does three things in a loop:

1. **Listen** — bind to a port and tell the OS "I want connections here"
2. **Accept** — wait (block) until a client connects, then get a `Conn` object
3. **Handle** — do something with that connection (in a separate goroutine so the loop can accept the next client immediately)

Read this before continuing:
- https://gobyexample.com/tcp-server
- Docs: https://pkg.go.dev/net#Listen

---

### Checkpoint 2.2 — Create the Listener

```go
// server.go
package main

import (
    "fmt"
    "log"
    "net"
)

func startServer(port string) {
    listener, err := net.Listen("__", ":"+port)
    if err != nil {
        log.Fatal(__)
    }
    defer listener.Close()

    fmt.Println("Listening on the port :" + port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(__)
            continue
        }
        go handleClient(conn) // we'll build this next
    }
}
```

Questions to answer before filling the blanks:
- What string goes in the first argument of `net.Listen()`? (TCP or UDP?)
- What does `defer` do here? What happens if you forget it?
- Why is `handleClient` launched with `go`? What breaks if you remove `go`?

**Verify before moving on:**
- Call `startServer("8989")` from `main()` and run the server
- Open a new terminal and run `nc localhost 8989`
- The connection should succeed and hang (nothing will happen yet — that is correct)

---

## Phase 3 — The Client

### Checkpoint 3.1 — Define a Client

Every connected client needs a name and a connection. Define a struct:

```go
// client.go
package main

import "net"

type Client struct {
    conn net.Conn
    name string
}
```

You also need shared state that all goroutines can access. Add this to `server.go`:

```go
import "sync"

var (
    clients    []*Client
    messages   []string
    mu         sync.Mutex
    maxClients = 10
)
```

**Why `sync.Mutex`?** When two goroutines (two clients) try to modify `clients` at the same time, data corruption happens. The mutex makes sure only one goroutine touches shared state at a time.

Read: https://gobyexample.com/mutexes

Answer before moving on: What is a race condition? Write a one-sentence answer to yourself.

---

### Checkpoint 3.2 — Welcome Banner

When a client connects, send this exact text before asking for their name:

```
Welcome to TCP-Chat!
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
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: 
```

Store this as a constant string. Then write a function:

```go
func welcomeClient(conn net.Conn) {
    conn.Write([]byte(__________))
}
```

**Note:** `conn.Write()` takes `[]byte`, not a string. Look up how to convert.

**Verify before moving on:**
- Connect with `nc localhost 8989`
- You should see the logo and the name prompt

---

### Checkpoint 3.3 — Reading the Client's Name

Now read what the client types. Use `bufio`:

```go
import (
    "bufio"
    "strings"
)

func getClientName(conn net.Conn) string {
    reader := bufio.NewReader(conn)
    for {
        name, err := reader.ReadString('\n')
        if err != nil {
            return ""
        }
        name = strings.TrimSpace(name)
        if __________ {
            return name
        }
        conn.Write([]byte("[ENTER YOUR NAME]: "))
    }
}
```

Fill in the blank: what condition keeps the loop going until a valid name is given?

**Why `bufio.NewReader`?** Network data arrives in chunks. `bufio` buffers the incoming bytes and lets you read line by line. Without it, you would have to manage raw byte slices yourself.

Read: https://pkg.go.dev/bufio#NewReader

**Verify before moving on:**
- Connect with `nc localhost 8989`
- Press Enter without typing a name — it should re-prompt
- Type a name and press Enter — the server should print it (add a temporary `log.Println` in your code to verify)

---

### Checkpoint 3.4 — Enforcing the Connection Limit

Before accepting any client past the limit:

```go
func handleClient(conn net.Conn) {
    mu.Lock()
    if len(clients) >= maxClients {
        conn.Write([]byte("Chat is full. Try again later.\n"))
        conn.Close()
        mu.Unlock()
        return
    }
    mu.Unlock()

    // continue with welcome, name, etc.
}
```

**Why unlock before `return`?** If you use `defer mu.Unlock()`, the unlock happens when the entire function returns — which may be much later. Here you want to unlock immediately after the check. Understand the difference.

**Verify before moving on:**
- Start the server
- Write a small shell loop: `for i in $(seq 1 11); do nc localhost 8989 & done`
- The 11th connection attempt should receive the "Chat is full" message

---

## Phase 4 — Joining the Chat

### Checkpoint 4.1 — Full handleClient Flow

Now put the full join sequence together. The order matters:

```go
func handleClient(conn net.Conn) {
    // 1. Check capacity — reject and return if full
    // 2. Send welcome banner
    // 3. Get name — loop until non-empty
    // 4. Create Client struct
    // 5. Lock → add to clients slice → unlock
    // 6. Send message history to the new client
    // 7. Broadcast "[name] has joined our chat..." to everyone else
    // 8. Start listening for their messages (next phase)
}
```

Do not move to step 5 until step 3 is complete. Think about why.

---

### Checkpoint 4.2 — Sending History

```go
func sendHistory(conn net.Conn) {
    mu.Lock()
    defer mu.Unlock()
    for _, msg := range messages {
        conn.Write([]byte(msg))
    }
}
```

**Why lock here?** Another goroutine could be appending to `messages` while you are ranging over it. That is a race condition.

**Verify before moving on:**
- Connect client A, send a few messages (just type them — broadcast isn't done yet, but you can add messages to the slice manually for testing)
- Connect client B — they should see A's messages immediately on join

---

### Checkpoint 4.3 — Broadcasting

```go
func broadcast(message string, sender *Client) {
    mu.Lock()
    defer mu.Unlock()

    messages = append(messages, message)

    for _, c := range clients {
        if c != sender {
            c.conn.Write([]byte(message))
        }
    }
}
```

Notice that `broadcast` both **saves** the message to history and **sends** it to all clients. This happens inside the same lock. Think about why that is important.

**Verify before moving on:**
- Connect two clients A and B
- B should see the notification that A joined

---

## Phase 5 — The Message Loop

### Checkpoint 5.1 — Reading and Sending Messages

After a client joins, keep reading from their connection forever:

```go
func listenToClient(client *Client) {
    reader := bufio.NewReader(client.conn)

    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            // Client disconnected
            removeClient(client)
            broadcast(client.name+" has left our chat...\n", nil)
            return
        }

        message = strings.TrimSpace(message)
        if message == "" {
            // Show prompt again but do not broadcast
            client.conn.Write([]byte(prompt(client.name)))
            continue
        }

        timestamp := time.Now().Format("__________")
        formatted := fmt.Sprintf("[%s][%s]:%s\n", __, __, __)

        broadcast(formatted, client)
        client.conn.Write([]byte(prompt(client.name)))
    }
}
```

Fill in the blanks:
- What format string makes Go produce `2020-01-20 16:03:43`? This is the trickiest part of Go's `time` package. Search: **"golang time format reference time"**
- What do the three `__` in `Sprintf` refer to?

A helper to generate the prompt:

```go
func prompt(name string) string {
    return fmt.Sprintf("[%s][%s]:", time.Now().Format("__________"), name)
}
```

**Resources:**
- https://pkg.go.dev/time#Time.Format
- Search: **"golang time format 2006"** — read why Go uses that specific reference date

**Verify before moving on:**
- Connect two clients
- Client A sends a message
- Client B sees `[timestamp][A]:message`
- Client A sees their own prompt reprinted after sending
- Empty messages do not appear for anyone

---

### Checkpoint 5.2 — Removing a Client

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

This is a standard Go slice removal pattern. Make sure you understand it before moving on.

Search: **"golang remove element from slice"** and read the explanation, not just the code.

**Verify before moving on:**
- Connect clients A and B
- Kill client A with `Ctrl+C`
- Client B should see "A has left our chat..."
- Client B should remain connected

---

## Phase 6 — Wiring Everything Together

### Checkpoint 6.1 — Call listenToClient from handleClient

Add this as the last line of `handleClient`:

```go
listenToClient(client)
```

This call blocks — it runs the loop until the client disconnects. That is fine because `handleClient` is already running in its own goroutine.

---

### Checkpoint 6.2 — Final main.go

```go
func main() {
    port := "8989"
    if len(os.Args) == 2 {
        port = os.Args[1]
    } else if len(os.Args) > 2 {
        fmt.Println("[USAGE]: ./TCPChat $port")
        return
    }
    startServer(port)
}
```

---

## Phase 7 — Full System Test

Run through every scenario manually before submission:

| Scenario | Expected Result |
|---|---|
| `go run .` | Listens on :8989 |
| `go run . 2525` | Listens on :2525 |
| `go run . 2525 localhost` | Prints usage and exits |
| Client connects | Sees logo and name prompt |
| Client enters empty name | Re-prompted |
| Client A connects, sends messages, then B connects | B sees A's history |
| A sends a message | B sees it with timestamp and name |
| A sends an empty message | Nobody sees it |
| B disconnects | A sees "B has left our chat..." |
| A stays connected after B leaves | A is still in the chat |
| 11th client connects | Gets rejected |

---

## Phase 8 — The Race Detector

Run your server with:

```bash
go run -race . 2525
```

Then connect multiple clients and send messages rapidly. If you have unprotected shared state, Go will report the race here. Fix every race the detector finds before submission.

Search: **"golang race detector"** to understand the output format.

---

## Phase 9 — Unit Tests

The project requires test files. Write at least these:

- A test that starts a server on a random port and verifies it accepts a connection
- A test that verifies `removeClient` correctly removes a client from the slice
- A test that verifies `broadcast` does not send a message back to the sender

Read: https://go.dev/doc/tutorial/add-a-test

```go
// server_test.go
package main

import (
    "testing"
)

func TestRemoveClient(t *testing.T) {
    // Set up: create fake clients slice with 3 entries
    // Call removeClient on the middle one
    // Assert: slice now has 2 entries
    // Assert: the removed client is gone
}
```

---

## Debugging Reference

**Server freezes completely (deadlock)**
Cause: You called a function that locks the mutex from inside code that already holds the lock.
Fix: Never call `broadcast` while holding `mu`. Restructure the call order.

**Messages appear on the wrong line / overwrite user input**
Cause: A broadcast arrives while the user is mid-typing.
Fix: Before writing to a client, send `\r` to return the cursor to the start of the line, write the message, then reprint their prompt.

**A dead client stays in the list**
Cause: `ReadString` returned an error but you did not call `removeClient`.
Fix: Check every error return from `reader.ReadString`.

**`go run -race` reports a race**
Cause: A slice or variable is being read and written from two goroutines without a lock.
Fix: Wrap every access to `clients` and `messages` in `mu.Lock()` / `mu.Unlock()`.

---

## Bonus Challenges

These are not required but will deepen your understanding:

- Let clients type `/name NewName` to change their display name. Notify the group.
- Save all messages to a `logs.txt` file using `os.OpenFile`. Search: **"golang append to file"**
- Create named chat rooms. Clients type `/join roomname` to switch.

---

## Key Packages Used

| Package | What You Use It For | Docs |
|---|---|---|
| `net` | Listen, accept, read, write over TCP | https://pkg.go.dev/net |
| `bufio` | Read lines from a network stream | https://pkg.go.dev/bufio |
| `sync` | Mutex to protect shared state | https://pkg.go.dev/sync |
| `time` | Format message timestamps | https://pkg.go.dev/time |
| `strings` | TrimSpace to clean user input | https://pkg.go.dev/strings |
| `fmt` | Format messages | https://pkg.go.dev/fmt |
| `log` | Log server-side errors | https://pkg.go.dev/log |
| `os` | Read command-line arguments | https://pkg.go.dev/os |

---

## Submission Checklist

- [ ] Default port 8989 when no argument given
- [ ] Usage message when more than one argument given
- [ ] Linux logo sent to connecting clients
- [ ] Empty names rejected and re-prompted
- [ ] Maximum 10 connections enforced
- [ ] New client receives full message history on join
- [ ] All clients notified when someone joins
- [ ] All clients notified when someone leaves
- [ ] Empty messages not broadcast
- [ ] Messages formatted as `[timestamp][name]:message`
- [ ] Remaining clients unaffected when one disconnects
- [ ] Mutex used for all shared state
- [ ] Goroutines used for concurrent client handling
- [ ] Errors handled on server and client side
- [ ] `go run -race` reports no races
- [ ] Unit tests written and passing