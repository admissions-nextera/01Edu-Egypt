# 📘 Learning Go — Chapter 1 Quiz
## Setting Up Your Go Environment

**Time Limit:** 25 minutes  
**Total Questions:** 20  
**Passing Score:** 16/20 (80%)

> This quiz covers: installing Go, the Go toolchain, `go run` vs `go build`, `go fmt`, `go vet`, modules, and the workspace layout.

---

## 📋 SECTION 1: THE GO TOOLCHAIN (7 Questions)

### Q1: What is the difference between `go run` and `go build`?

**A)** `go run` is for scripts; `go build` is for packages  
**B)** `go run` compiles and immediately executes without saving a binary; `go build` compiles and saves a binary to disk  
**C)** `go build` is faster  
**D)** `go run` only works on `main.go`  

<details><summary>💡 Answer</summary>

**B) `go run` compiles + executes immediately; `go build` saves a binary**

```bash
go run main.go       # compile + run, no binary saved
go build .           # produces ./myprogram (or myprogram.exe on Windows)
./myprogram          # run the saved binary
```

Use `go run` during development for quick iteration. Use `go build` when you want to distribute or deploy an executable.

</details>

---

### Q2: What does `go fmt` do?

**A)** Formats output to the terminal  
**B)** Automatically reformats your Go source code to match the official Go style — indentation, spacing, brace placement  
**C)** Checks for formatting errors without changing files  
**D)** Compresses the source code  

<details><summary>💡 Answer</summary>

**B) Automatically reformats source code to the official Go style**

```bash
go fmt ./...   # format all files in the current module
```

In Go, formatting is not a style preference — it is mandatory. `go fmt` (or `gofmt`) produces the one canonical format. Tabs for indentation, specific spacing rules. Most editors run it on save. PRs with unformatted code are rejected by convention.

</details>

---

### Q3: What does `go vet` do?

**A)** Checks your code compiles correctly  
**B)** Installs veterinary software  
**C)** Runs static analysis to find likely bugs: mismatched `Printf` format strings, unreachable code, incorrect mutex usage, and similar  
**D)** Removes unused variables  

<details><summary>💡 Answer</summary>

**C) Static analysis for likely bugs**

```bash
go vet ./...   # vet all packages in the module
```

`go vet` catches things the compiler won't: `fmt.Printf("%d", "hello")` — wrong type for format verb. It doesn't catch all bugs but it finds common, real mistakes. Run it as part of your standard workflow alongside `go fmt`.

</details>

---

### Q4: What is `go build ./...` doing — what does `./...` mean?

**A)** Builds all files named `...`  
**B)** A pattern meaning "the current directory and all subdirectories recursively" — builds every package in the module  
**C)** Builds only the main package  
**D)** Downloads dependencies  

<details><summary>💡 Answer</summary>

**B) `./...` = current directory and all subdirectories recursively**

This pattern is used throughout Go tooling:
```bash
go build ./...   # build everything
go test ./...    # test everything
go fmt ./...     # format everything
go vet ./...     # vet everything
```

It's the standard way to apply a command to an entire module at once.

</details>

---

### Q5: What does `go get` do in a module-based project?

**A)** Downloads the Go runtime  
**B)** Adds or updates a dependency in `go.mod` and downloads it to the module cache  
**C)** Gets the current Go version  
**D)** Fetches environment variables  

<details><summary>💡 Answer</summary>

**B) Adds/updates a dependency in `go.mod` and downloads it**

```bash
go get github.com/some/package@v1.2.3
```

This updates `go.mod` with the new dependency and its version, and updates `go.sum` with the cryptographic checksum. The package is downloaded to the module cache (`$GOPATH/pkg/mod`).

</details>

---

### Q6: What command do you run after cloning a Go project to install its dependencies?

**A)** `go install`  
**B)** `go get ./...`  
**C)** `go mod download` or just `go build ./...` (which downloads automatically)  
**D)** `npm install`  

<details><summary>💡 Answer</summary>

**C) `go mod download` or `go build`**

Go modules are self-describing via `go.mod` and `go.sum`. Running `go build` or `go test` automatically downloads any missing dependencies. `go mod download` pre-downloads them explicitly without building. No separate install step like npm or pip.

</details>

---

### Q7: What is the purpose of `go mod tidy`?

**A)** Sorts the imports alphabetically  
**B)** Removes unused dependencies from `go.mod` and `go.sum` and adds any missing ones  
**C)** Tidies the source code formatting  
**D)** Removes compiled binaries  

<details><summary>💡 Answer</summary>

**B) Removes unused deps and adds missing ones**

Run `go mod tidy` after adding or removing imports. It keeps `go.mod` and `go.sum` in sync with what your code actually uses. Without it, `go.mod` can accumulate stale entries that slow builds and confuse readers.

</details>

---

## 📋 SECTION 2: MODULES AND WORKSPACE (6 Questions)

### Q8: What is a Go module?

**A)** A single `.go` file  
**B)** A collection of packages that are versioned together, defined by a `go.mod` file at the root  
**C)** A function that can be imported  
**D)** A directory that contains only one package  

<details><summary>💡 Answer</summary>

**B) A collection of packages versioned together, defined by `go.mod`**

A module is the unit of dependency management in Go (since Go 1.11). The `go.mod` file declares the module path (e.g. `module github.com/user/myapp`) and the Go version. Everything inside the module directory is part of the module.

</details>

---

### Q9: What does `go mod init github.com/user/myapp` do?

**A)** Clones the repo from GitHub  
**B)** Creates a `go.mod` file declaring the module path `github.com/user/myapp`  
**C)** Initializes a Git repository  
**D)** Installs the module globally  

<details><summary>💡 Answer</summary>

**B) Creates `go.mod` with the module path**

```
module github.com/user/myapp

go 1.21
```

The module path is used for imports. If you write `package util` in `util/util.go`, other packages in the same module import it as `github.com/user/myapp/util`. The module path doesn't need to match a real URL during development, but it must for published modules.

</details>

---

### Q10: What is `go.sum` and why should you commit it to version control?

**A)** A file listing all functions in the module  
**B)** A file containing cryptographic checksums of every dependency version — committing it ensures everyone uses identical dependency code  
**C)** A summary of test results  
**D)** An auto-generated file that should be in `.gitignore`  

<details><summary>💡 Answer</summary>

**B) Cryptographic checksums of dependencies — commit it**

`go.sum` contains SHA-256 hashes of every dependency zip and `go.mod` file. When anyone runs `go build`, Go verifies the downloaded code matches these checksums — preventing supply chain attacks. Always commit `go.sum`. Never gitignore it.

</details>

---

### Q11: Where does Go store downloaded module dependencies?

**A)** Inside the project directory in a `vendor/` folder (always)  
**B)** In `$GOPATH/pkg/mod` — the module cache shared across all projects on the machine  
**C)** In `/usr/local/go`  
**D)** In the current directory  

<details><summary>💡 Answer</summary>

**B) `$GOPATH/pkg/mod` — the shared module cache**

Downloaded modules are cached at `$GOPATH/pkg/mod` (default: `~/go/pkg/mod`). This cache is shared — if two projects use the same dependency version, it's downloaded only once. The cache is immutable (read-only) to prevent accidental modification.

</details>

---

### Q12: What is `GOPATH`?

**A)** The path to the Go compiler  
**B)** A workspace directory (default `~/go`) that stores the module cache, installed binaries, and (historically) source code before modules were introduced  
**C)** The path to the current project  
**D)** An environment variable that must point to the project root  

<details><summary>💡 Answer</summary>

**B) A workspace directory storing cache and installed binaries**

In pre-module Go, all code had to live under `$GOPATH/src`. With modules (Go 1.11+), you can put your project anywhere. Today `GOPATH` mainly matters for:
- `$GOPATH/pkg/mod` — module cache
- `$GOPATH/bin` — installed binaries (`go install` puts them here)

Add `$GOPATH/bin` to your `PATH` to run installed tools.

</details>

---

### Q13: What is the `go install` command used for?

**A)** Installs Go itself  
**B)** Compiles and installs a package's binary to `$GOPATH/bin` — used to install command-line tools written in Go  
**C)** Same as `go build` but faster  
**D)** Installs a module as a dependency  

<details><summary>💡 Answer</summary>

**B) Compiles and installs binary to `$GOPATH/bin`**

```bash
go install github.com/some/tool@latest
```

After this, `tool` is available as a command in your terminal (if `$GOPATH/bin` is in `PATH`). This is how Go CLI tools are distributed — `gopls`, `staticcheck`, `golangci-lint`, etc.

</details>

---

## 📋 SECTION 3: GO PROGRAM STRUCTURE (4 Questions)

### Q14: Every executable Go program must have what?

**A)** A file named `main.go`  
**B)** A `package main` declaration and a `func main()` function  
**C)** A `go.mod` file in the same directory  
**D)** An `init()` function  

<details><summary>💡 Answer</summary>

**B) `package main` + `func main()`**

```go
package main

func main() {
    // entry point
}
```

The file doesn't have to be named `main.go` (any name works). But the package must be `main` and there must be a `main` function. Library packages use other package names and have no `main`.

</details>

---

### Q15: What is the `goimports` tool?

**A)** A replacement for `go get` that imports packages  
**B)** A formatter that does everything `gofmt` does AND automatically adds/removes import statements  
**C)** A tool for generating import documentation  
**D)** Part of the standard Go installation  

<details><summary>💡 Answer</summary>

**B) `gofmt` + automatic import management**

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

`goimports` reformats code AND figures out which packages need to be imported (or removed). Most Go developers configure their editor to run `goimports` on save. It's not in the standard toolchain but is universally used.

</details>

---

### Q16: You have a Go file with an import that isn't used. What happens when you try to compile it?

**A)** A warning — the program still compiles  
**B)** A compile error: "imported and not used"  
**C)** The import is silently removed  
**D)** The program compiles but runs slower  

<details><summary>💡 Answer</summary>

**B) Compile error — unused imports are not allowed**

Go enforces unused import elimination at compile time — unlike most languages where unused imports produce only a warning. This keeps code clean and prevents bloat. The fix: remove the import, or use `_` for side-effect imports.

</details>

---

### Q17: What does this import do?
```go
import _ "github.com/lib/pq"
```

**A)** Compile error — `_` is invalid in imports  
**B)** Imports the package for its side effects (e.g. `init()` function runs and registers a database driver) without making its exported names available  
**C)** Imports an anonymous package  
**D)** Ignores any errors from the package  

<details><summary>💡 Answer</summary>

**B) Side-effect import — runs `init()` without exposing names**

Database drivers, image decoders, and similar packages register themselves via `init()`. The blank identifier `_` tells Go "I know I'm not using any exported names — import only for the side effect." Without it, the `import` would be rejected as unused.

</details>

---

### Q18: What is `golangci-lint`?

**A)** The official Go linter  
**B)** A meta-linter that runs many linters simultaneously (including `go vet` and dozens more) with a single configuration file  
**C)** A tool for checking Go license compliance  
**D)** A formatter  

<details><summary>💡 Answer</summary>

**B) A meta-linter running many linters at once**

`golangci-lint` is the industry-standard linting solution for Go. It wraps `go vet`, `staticcheck`, `errcheck`, `gosimple`, and 50+ others. Configure which linters to run in `.golangci.yml`. Most professional Go teams use it in CI.

</details>

---

## 📋 SECTION 4: ENVIRONMENT & VERSIONS (2 Questions)

### Q19: How do you check which version of Go is installed?

**A)** `go --version`  
**B)** `go version`  
**C)** `go env GOVERSION`  
**D)** Both B and C work  

<details><summary>💡 Answer</summary>

**D) Both `go version` and `go env GOVERSION` work**

```bash
go version        # go version go1.21.0 linux/amd64
go env GOVERSION  # go1.21.0
```

`go env` prints all Go environment variables. `go env GOPATH`, `go env GOROOT`, `go env GOOS`, `go env GOARCH` are all useful.

</details>

---

### Q20: What does `GOOS=linux GOARCH=amd64 go build .` do?

**A)** Runs your program on Linux  
**B)** Cross-compiles your program for Linux on an AMD64 processor, regardless of what OS and architecture you're currently on  
**C)** An error — you can't change GOOS  
**D)** Installs Go on a Linux system  

<details><summary>💡 Answer</summary>

**B) Cross-compiles for Linux/AMD64 from any machine**

Go supports cross-compilation out of the box. Set `GOOS` (target OS) and `GOARCH` (target architecture) as environment variables before `go build`. No extra tools needed. This is one of Go's most practically useful features — build for Linux from a Mac with one command.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 18–20 ✅ | **Excellent.** Strong foundation — proceed to Chapter 2. |
| 16–17 ✅ | **Ready.** Review any missed questions before moving on. |
| 12–15 ⚠️ | **Review first.** Spend time with the Go toolchain docs and `go help`. |
| Below 12 ❌ | **Setup issues likely.** Work through the official Go installation and "Getting Started" guide. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q7 | `go run`, `go build`, `go fmt`, `go vet`, `go get`, `go mod tidy` |
| Q8–Q13 | Modules, `go.mod`, `go.sum`, GOPATH, module cache, `go install` |
| Q14–Q18 | `package main`, `func main`, unused imports, blank import `_`, linters |
| Q19–Q20 | `go version`, `go env`, cross-compilation |