# ASCII-Art-Justify Project Guide

> **Before you start:** This project builds on ascii-art, ascii-art-fs, and ascii-art-output. All three must be working. Read about how text alignment works in any word processor — center, left, right, and justify — before writing any code.
>
> **Standard library only.** This project uses only packages from https://pkg.go.dev/std. You will use `syscall` and `unsafe` directly to read the terminal size. A popular third-party alternative is `golang.org/x/term` — you will learn what it does and why students use it, but you will not use it here.

---

## Objectives

By completing this project you will learn:

1. **Terminal Width Detection** — Reading the current terminal size using a raw syscall from Go's standard library
2. **syscall and unsafe** — How Go talks directly to the Linux kernel without a third-party wrapper
3. **Text Alignment Algorithms** — Implementing center, left, right, and justify alignment mathematically
4. **Dynamic Layout** — Adapting output to the available space instead of printing fixed-width content
5. **Justify Algorithm** — Distributing extra space evenly between words — the hardest alignment to implement
6. **Refactoring** — Changing a function's return type without breaking existing behavior

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art-Output (Completed)
- `render` returns a string
- Banner loading and rendering pipeline working

### 2. The Terminal Size Syscall
- What `TIOCGWINSZ` is — the ioctl request code that asks the kernel for terminal dimensions
- What `syscall.Syscall` does and how to pass a pointer to a struct through it
- What `unsafe.Pointer` is and why it is needed here
- Search: **"linux TIOCGWINSZ ioctl terminal size"**
- Search: **"golang syscall SYS_IOCTL TIOCGWINSZ"**
- Search: **"golang unsafe Pointer syscall example"**

> **Third-party awareness:** `golang.org/x/term` is a package maintained by the Go team that wraps exactly this syscall into a clean `term.GetSize(fd)` function. It is not in the standard library — it requires `go get golang.org/x/term`. In real projects you would likely use it. Here, you will implement what it wraps yourself, so you understand what is underneath.

### 3. Alignment Math
- What it means to center a string of width W in a space of width T
- What right-alignment means in terms of padding
- What justify alignment is and how it differs from the others
- Draw each alignment on paper before writing any code

### 4. Go standard packages you will use
- `syscall` — `syscall.Syscall`, `syscall.SYS_IOCTL`, `syscall.TIOCGWINSZ`
- `unsafe` — `unsafe.Pointer`, `unsafe.Sizeof`
- `strings` — split words, join rows, repeat spaces
- `os` — `os.Stderr`, check stdout file descriptor

**If any of these are unfamiliar, read about them before writing any code.**

---

## Project Structure

```
ascii-art-justify/
├── main.go
├── banner.go       ← existing render logic
├── align.go        ← new: terminal width + all alignment functions
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Parse the `--align` Flag

**Goal:**
```
go run . --align=center "hello" standard
go run . --align=right "hello" shadow
go run . --align=left "Hello There" standard
go run . --align=justify "how are you" shadow
go run . "hello"                         → left alignment by default
go run . --align=unknown "hello"         → usage message
```

**Questions to answer before writing anything:**
- After stripping `--align=center` from the args, what arguments remain?
- What are the four valid alignment types? What happens for anything else?
- What is the default when no `--align` flag is given?
- Your program also supports `--output=` from ascii-art-output. How do you strip both flags before reading the string and banner arguments?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Scan os.Args for "--align=..."
    //    Extract the value after "=", remove the arg from the slice
    //    Default to "left" if not found
    //    If value is not one of: left, right, center, justify → print usage and return

    // 2. Also scan for "--output=..." (from ascii-art-output)
    //    Extract filename, remove the arg

    // 3. Parse remaining args for [STRING] and optional [BANNER]
    //    Same logic as ascii-art-fs

    // 4. Get terminal width (Milestone 2)

    // 5. Load banner, render with alignment applied

    // 6. Write to file or stdout
}
```

**Usage message:**
```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

**Verify:** All six examples above produce the correct behavior before you write any alignment logic.

---

## Milestone 2 — Get the Terminal Width Using `syscall`

**Goal:** At runtime, detect how wide the terminal window currently is — using only the standard library.

**Questions to answer:**
- What struct does the kernel fill when you call `TIOCGWINSZ`? What fields does it have?
- How do you define that struct in Go so its memory layout matches what the kernel expects?
- What three arguments does `syscall.Syscall` take? What do you pass for each one?
- Why do you need `unsafe.Pointer` to pass your struct to the kernel?
- What should you return if the syscall fails (for example, when stdout is redirected to a file)?

> **Third-party comparison:** `golang.org/x/term.GetSize(fd int)` does exactly this internally. If you read its source at https://cs.opensource.google/go/x/term, you will see it calls the same `syscall.Syscall` with the same `TIOCGWINSZ`. The difference is it handles multiple operating systems (Linux, macOS, Windows). Your implementation only needs to work on Linux.

**Code Placeholder:**
```go
// align.go

import (
    "syscall"
    "unsafe"
)

// winsize mirrors the kernel's winsize struct from <sys/ioctl.h>
// The field order and types must match exactly
type winsize struct {
    // Row    uint16  ← number of terminal rows
    // Col    uint16  ← number of terminal columns
    // Xpixel uint16  ← pixel width (unused here)
    // Ypixel uint16  ← pixel height (unused here)
}

func getTerminalWidth() int {
    // 1. Create a zeroed winsize struct

    // 2. Call syscall.Syscall with:
    //    - syscall.SYS_IOCTL as the syscall number
    //    - uintptr(syscall.Stdout) as the file descriptor
    //    - uintptr(syscall.TIOCGWINSZ) as the request
    //    - uintptr(unsafe.Pointer(&ws)) as the pointer to your struct

    // 3. Check the return value — if it indicates failure, return 80

    // 4. Return int(ws.Col)
}
```

**Resources:**
- `man ioctl_tty` — search for `TIOCGWINSZ` and the `winsize` struct
- https://pkg.go.dev/syscall — search for `SYS_IOCTL`, `TIOCGWINSZ`, `Syscall`
- https://pkg.go.dev/unsafe — read `Pointer`
- Search: **"golang syscall TIOCGWINSZ winsize struct"**

**Verify:**
```bash
go run . --align=left "hello"   # prints terminal width before rendering
# Resize your terminal and run again — the number should change
```

---

## Milestone 3 — Calculate the Width of Rendered Text

**Goal:** Before aligning, know how wide the ASCII art will be in columns — without rendering it first.

**Questions to answer:**
- Each character in the banner occupies some number of columns. How do you measure it?
- Does every character in `standard.txt` have the same width? Test `'i'` versus `'W'` to find out.
- The total rendered width is the sum of individual character widths. How do you compute this for a multi-character string?

**Code Placeholder:**
```go
// align.go

func renderedWidth(banner []string, text string) int {
    // For each rune in text:
    //   Call getCharLines to get the 8 art lines for that character
    //   The character's column width = len of any one of those 8 lines
    //   Add it to a running total
    // Return the total
}
```

**Verify:** Calculate the expected width of `"hello"` in `standard.txt` by hand (open the file and count). Confirm your function returns the same number.

---

## Milestone 4 — Left and Right Alignment

**Goal:**
- Left: no padding (already how your renderer works)
- Right: each row is padded with spaces on the left so the art ends at the terminal's right edge

**Questions to answer:**
- If the terminal is T columns wide and the art is W columns wide, how many spaces go before each row?
- What do you do if the padding calculation produces a negative number?
- Do you add padding before the complete output, or before each individual row? Why does it matter?

**Code Placeholder:**
```go
// align.go

func alignRight(rows []string, termWidth int) []string {
    // For each row:
    //   padding = termWidth - len(row)
    //   If padding < 0: set to 0
    //   Prepend padding spaces to the row
    // Return the padded rows
}
```

**Verify:**
```bash
go run . --align=right "hello" shadow | cat -e
```
The art should be pushed against the right edge. Compare column positions against the spec example.

---

## Milestone 5 — Center Alignment

**Goal:** Each row of ASCII art is horizontally centered within the terminal width.

**Questions to answer:**
- If the terminal is T wide and the art is W wide, how many spaces go before each row?
- When `(T - W)` is odd, integer division truncates. Which side gets the smaller amount — left or right?
- Is this acceptable? What do standard text editors do in this situation?

**Code Placeholder:**
```go
// align.go

func alignCenter(rows []string, termWidth int) []string {
    // For each row:
    //   padding = (termWidth - len(row)) / 2
    //   If padding < 0: set to 0
    //   Prepend padding spaces to the row
    // Return the padded rows
}
```

**Verify:**
```bash
go run . --align=center "hello" standard | cat -e
```
Compare position against the spec example carefully.

---

## Milestone 6 — Justify Alignment

**Goal:** Multiple words are spread so the first starts at the left edge and the last ends at the right edge. Extra space is distributed as evenly as possible between words.

This is the most complex milestone. Think through the algorithm on paper before writing any code.

**Questions to answer:**
- How do you split the input string into words?
- What happens when there is only one word? (There are no gaps to distribute — fall back to left alignment.)
- You need the rendered width of each word separately. How do you get it?
- `totalSpace = termWidth - sum of all word widths`. You have `gaps = len(words) - 1` gaps. How do you split `totalSpace` across `gaps` gaps as evenly as possible?
- When `totalSpace` does not divide evenly into `gaps`, which gaps get an extra space?
- Each word is 8 rows tall. How do you assemble the final output row by row — interleaving word rows with the calculated gap spaces?

**Code Placeholder:**
```go
// align.go

func alignJustify(banner []string, words []string, termWidth int) []string {
    // 1. If len(words) <= 1: render the single word and return left-aligned rows

    // 2. Render each word into its 8-row block using renderLine
    //    Store as [][]string — one []string (8 rows) per word

    // 3. Calculate total word width: sum of len(wordRows[i][0]) for each word

    // 4. Calculate spacing:
    //    gaps      = len(words) - 1
    //    totalSpace = termWidth - totalWordWidth
    //    baseSpace  = totalSpace / gaps
    //    extraSpaces = totalSpace % gaps
    //    → first extraSpaces gaps get (baseSpace + 1) spaces
    //    → remaining gaps get baseSpace spaces

    // 5. Build output row by row (rows 0 to 7):
    //    For each row:
    //      Start with wordRows[0][row]
    //      For each subsequent word i:
    //        Append the gap for position i (baseSpace or baseSpace+1 spaces)
    //        Append wordRows[i][row]
    //    Collect the 8 assembled rows

    // 6. Return the 8 rows
}
```

**Resources:**
- Search: **"text justification algorithm spaces between words"**
- `strings.Repeat(" ", n)` — produces n spaces

**Verify:**
```bash
go run . --align=justify "how are you" shadow | cat -e
```
First word starts at column 1. Last word ends at the terminal's right edge. Space between words is as even as possible.

---

## Milestone 7 — Refactor `renderLine` to Return Rows

**Goal:** Justify needs individual rows from `renderLine`, not a joined string. Refactor so the return type is `[]string`.

**Questions to answer:**
- After the refactor, how does the non-justify path still produce the same output? (Hint: `strings.Join`)
- Does changing `renderLine`'s return type break any existing test cases? How do you check?
- What is the correct pipeline order for right alignment? `renderLine → alignRight → join → print` — or something else?

**Code Placeholder:**
```go
// banner.go

func renderLine(banner []string, text string) []string {
    // Build and return a slice of exactly 8 strings (one per art row)
    // Do NOT print — return the rows to the caller
}

func render(banner []string, input string) string {
    // Split input on "\\n"
    // For each part:
    //   If empty: add a blank line to output
    //   Otherwise: call renderLine, join the 8 rows with "\n", add to output
    // Return the full output string
}
```

**Verify:** After refactoring, run all existing ascii-art test cases. Every one must produce identical output.

---

## Debugging Checklist

- Does `syscall.Syscall` always return -1? Make sure you are passing `uintptr(syscall.Stdout)` (value 1), not `uintptr(os.Stdout.Fd())` — both work but the second requires the import of `os`. Check that `TIOCGWINSZ` is the right constant on Linux (`0x5413`).
- Does the terminal width come back as 0? Your `winsize` struct field order may not match the kernel's layout. `Row` must come before `Col` — the kernel fills them in that order.
- Is right alignment off by a few columns? You are measuring padding with a hardcoded width instead of using `len(row)` on the actual rendered row. Measure the rendered row, not a calculated estimate.
- Is justify alignment producing uneven spacing? Check your `extraSpaces` logic — the first `extraSpaces` gaps get one extra space, the rest get `baseSpace`. Print each gap size to verify before assembling rows.
- Does the justify output look correct for 3 words but wrong for 2? A two-word input has exactly 1 gap — all the space goes there. Trace through your formula manually for that case.
- Does refactoring `renderLine` break previous test cases? Make sure `render` calls `strings.Join(renderLine(...), "\n") + "\n"` instead of printing directly.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `syscall` | `Syscall`, `SYS_IOCTL`, `TIOCGWINSZ`, `Stdout` | https://pkg.go.dev/syscall |
| `unsafe` | `Pointer` — pass struct address to kernel | https://pkg.go.dev/unsafe |
| `strings` | Split words, join rows, repeat spaces | https://pkg.go.dev/strings |
| `os` | Args, write output to file or stdout | https://pkg.go.dev/os |
| `fmt` | Print usage message and errors | https://pkg.go.dev/fmt |

> **Third-party packages you will NOT use here but should know exist:**
>
> | Package | What it does | Why students use it |
> |---|---|---|
> | `golang.org/x/term` | Wraps `TIOCGWINSZ` into `term.GetSize(fd)` | Cleaner API, cross-platform (Linux, macOS, Windows) |
> | `github.com/mattn/go-isatty` | Detects if stdout is a real terminal | Useful before calling `GetSize` |
>
> These are not allowed in this project, but they exist precisely to solve the problem you are solving here. Reading their source code after you finish is a good exercise.

---

## Submission Checklist

- [ ] `--align=left` produces left-aligned output (same as default behavior)
- [ ] `--align=right` pushes art flush to the right edge of the terminal
- [ ] `--align=center` centers art horizontally in the terminal
- [ ] `--align=justify` spreads words from left edge to right edge with even spacing
- [ ] Single-word justify falls back to left alignment without crashing
- [ ] Terminal width read using `syscall.Syscall` + `TIOCGWINSZ` — no third-party packages
- [ ] Falls back to width 80 when the syscall fails (redirected output)
- [ ] Invalid alignment type prints usage message and exits cleanly
- [ ] Works with `shadow` and `thinkertoy` banners
- [ ] Compatible with `--output` flag from ascii-art-output
- [ ] All previous ascii-art test cases still pass after `renderLine` refactor
- [ ] Only standard Go packages used (`go.mod` has no external dependencies)
