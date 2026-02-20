# ASCII-Art-FS Project Guide

> **Before you start:** This project builds directly on the ascii-art project. You must have that working before you start here. If you have not completed it yet, go back and finish it first.

---

## Objectives

By completing this project you will learn:

1. **Argument Parsing** — Handling an optional second argument that changes program behavior
2. **File Selection Logic** — Choosing which banner file to load based on user input
3. **Input Validation** — Returning a clean usage message for any incorrect argument format
4. **Backward Compatibility** — Making a program still work with its original single-argument form after adding new features

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art Project (Completed)
- `loadBanner`, `getCharLines`, `renderLine`, `render` — all working correctly
- The banner file format and the start-line formula

### 2. Argument Handling
- `os.Args` — reading multiple arguments
- `strings.HasSuffix` — checking file extensions or names
- How to distinguish between `go run . "hello"` and `go run . "hello" standard`

**If your ascii-art project is not passing all test cases, fix it before continuing.**

---

## Project Structure

```
ascii-art-fs/
├── main.go
├── banner.go       ← your existing banner loading and rendering logic
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Accept an Optional Banner Argument

**Goal:**
```
go run . "hello" standard       → renders using standard.txt
go run . "Hello There!" shadow  → renders using shadow.txt
go run . "hello"                → renders using standard.txt (default)
go run . "hello" unknown        → prints usage message
go run .                        → prints usage message
go run . too many args here     → prints usage message
```

**Questions to answer before writing anything:**
- How many valid argument combinations does your program now accept?
- How do you map the string `"standard"` to the filename `"standard.txt"`?
- What are the three valid banner names? What should happen for anything else?
- What condition triggers the usage message?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Check argument count — valid cases are 1 argument or 2 arguments
    //    Any other count: print usage and return

    // 2. Set the input string from the first argument

    // 3. Set the banner name — default to "standard" if only 1 argument given
    //    Use the second argument if provided

    // 4. Resolve the banner name to a filename
    //    Valid names: "standard", "shadow", "thinkertoy"
    //    Anything else: print usage and return

    // 5. Load the banner file
    //    Handle the error

    // 6. Render and print
}
```

**Usage message to print:**
```
Usage: go run . [STRING] [BANNER]
EX: go run . something standard
```

**Verify:**
```bash
go run . "hello" standard | cat -e
go run . "Hello There!" shadow | cat -e
go run . "Hello There!" thinkertoy | cat -e
go run . "hello"                          # should use standard
go run . "hello" unknown                  # should print usage
go run .                                  # should print usage
```
Compare each against the spec output exactly.

---

## Milestone 2 — Backward Compatibility

**Goal:** Every test case from the original ascii-art project still passes without modification.

**Questions to answer:**
- Does your new argument parsing break any of the original single-argument cases?
- Does `go run . ""` still print nothing?
- Does `go run . "Hello\nThere"` still render correctly?

**Verify:** Run your original ascii-art test cases with the new code. None of them should behave differently.

---

## Debugging Checklist

- Does the program crash when only 1 argument is given? Make sure your banner name defaulting logic runs before any access to `os.Args[2]`.
- Does `"shadow"` map to the correct filename? Print the resolved filename before loading to verify.
- Does an unknown banner name print the usage message and exit cleanly without a file-not-found panic?

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Read args, read banner file | https://pkg.go.dev/os |
| `strings` | Map banner name to filename | https://pkg.go.dev/strings |
| `fmt` | Print usage message | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] `go run . "hello" standard` renders correctly
- [ ] `go run . "hello" shadow` renders correctly
- [ ] `go run . "hello" thinkertoy` renders correctly
- [ ] `go run . "hello"` defaults to standard without error
- [ ] Unknown banner name prints usage message and exits
- [ ] Wrong argument count prints usage message and exits
- [ ] All original ascii-art test cases still pass
- [ ] Output verified with `cat -e` — no trailing spaces, correct line endings