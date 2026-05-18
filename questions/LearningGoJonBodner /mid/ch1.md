# Learning Go — Chapter 1: Questions
### Jon Bodner | Deep Mastery via First Principles

---

## Section 1: Compilation Model & Binary Execution

**Q1.**
`go build` produces a native binary. At the OS level — not the Go level — what does the kernel's loader do before handing execution to your program? What must exist in process memory before `main()` runs?

---

**Q2.**
Go produces a **statically linked** binary by default. What does static linking physically mean in terms of what is embedded in the binary vs. what is resolved at runtime? How does this differ from a dynamically linked binary (like a typical C program)?

---

**Q3.**
A dynamically linked program (e.g., a C program using glibc) requires the correct version of `libc.so` to exist on the host machine. Does a statically linked Go binary have this requirement? What does this mean operationally when deploying to a Docker container or a different Linux distro?

---

**Q4.**
Go embeds its own **runtime** into every compiled binary. At a physical level, what responsibilities does the Go runtime carry that would otherwise belong to the OS or an external VM? Name at least three.

---

**Q5.**
When you run `go run main.go`, is the resulting binary statically linked the same way as `go build`? What is the mechanical difference between the two commands at the filesystem and process level?

---

## Section 2: The Module System (`go mod`)

**Q6.**
Before Go modules (pre-1.11), Go used `GOPATH`. What was the fundamental problem with `GOPATH` in terms of **dependency versioning** — specifically, what happened if two projects on the same machine needed different versions of the same library?

---

**Q7.**
`go mod init` creates a `go.mod` file. What is the minimum information encoded in `go.mod` and why is each field necessary for **reproducible builds**?

---

**Q8.**
There are two files: `go.mod` and `go.sum`. They serve different purposes. What specific security or integrity guarantee does `go.sum` provide that `go.mod` cannot? What attack does `go.sum` protect against?

---

**Q9.**
When you run `go get` to add a dependency, Go uses **Minimum Version Selection (MVS)**. Most package managers (npm, pip) default to fetching the *latest* compatible version. Go deliberately does not. Why? What problem does MVS solve at a systems level?

---

**Q10.**
A Go module path looks like `github.com/someuser/someproject`. This looks like a URL but Go doesn't always fetch from it directly. What is the module path *actually* used for at the compiler and linker level? Is it a network address, a namespace, an identifier, or something else — and why does it matter that it's globally unique?

---

## Section 3: Program Structure & the `main` Package

**Q11.**
Every executable Go program requires `package main` and `func main()`. At the **OS and linker level**, what is `main()` in terms of process entry points? Is `main()` the *actual* first instruction the CPU executes when your binary runs? If not, what runs before it?

---

**Q12.**
Go has `package main` as a special package. What makes it structurally different from any other package (e.g., `package fmt`)? What does the linker do differently when it encounters `package main`?

---

**Q13.**
The `import` statement in Go is resolved entirely at **compile time**, not at runtime. What does this mean for unused imports? Why does the Go compiler treat an unused import as a **compile error** rather than a warning? Reason from the perspective of binary size and dependency graph integrity.

---

**Q14.**
Go's `fmt.Println` is in the `fmt` package. When your program calls `fmt.Println`, where does that function's machine code physically live at runtime? Is it in a separate `.so` file, loaded dynamically, or somewhere else? Trace the path from source to execution.

---

## Section 4: The Go Toolchain

**Q15.**
`go fmt` enforces a single, non-negotiable code style. This is unusual — most languages leave formatting to the developer. What is the **compiler and toolchain** argument for enforced formatting? Think in terms of AST normalization, diff noise in version control, and cognitive overhead.

---

**Q16.**
`go vet` is separate from `go build`. Both operate on Go source code. What is the fundamental distinction between what a **compiler** checks and what a **linter/vet tool** checks? Why can't `go build` catch everything `go vet` catches?

---

**Q17.**
`go build` compiles only the packages that have **changed** since the last build (incremental compilation via build cache). Where does Go store this cache, and at what granularity is the cache invalidated — per file, per package, per module? What triggers a cache miss?

---

**Q18.**
Cross-compilation in Go is done by setting `GOOS` and `GOARCH` environment variables before `go build`. No special toolchain or cross-compiler is needed. Why is this architecturally possible in Go but traditionally difficult in C? What property of Go's standard library and runtime makes this work?

---

## Section 5: Hand-Trace & Execution Proof

**Q19. — Hand Trace Required**
Given this program:

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Trace the following **without running it**:
- What does `go build` produce on a Linux/amd64 machine?
- What ELF sections exist in the output binary?
- When the OS executes it, what is the sequence of events from `execve()` syscall to the string `"Hello, World!"` appearing on stdout?
- What syscall ultimately writes the string to the terminal?

---

**Q20. — Contradiction Probe**
A classmate says:

> *"Go programs start faster than Java programs because Go doesn't have a runtime — it compiles straight to machine code with no overhead."*

Identify every factually incorrect claim in that statement. Be precise. Reference what Go's runtime actually does before `main()` is called.

---

*End of Chapter 1 Questions*