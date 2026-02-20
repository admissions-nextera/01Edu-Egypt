# ASCII-Art-Justify Project Guide

> **Before you start:** This project builds on ascii-art, ascii-art-fs, and ascii-art-output. All three must be working. Additionally, read about how text alignment works in any word processor — center, left, right, and justify — before writing any code.

---

## Objectives

By completing this project you will learn:

1. **Terminal Width Detection** — Reading the current terminal size at runtime
2. **Text Alignment Algorithms** — Implementing center, left, right, and justify alignment mathematically
3. **Dynamic Layout** — Adapting output to the available space instead of printing fixed-width content
4. **Justify Algorithm** — Distributing extra space evenly between words — the hardest alignment to implement

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art-Output (Completed)
- `render` returns a string
- Banner loading and rendering pipeline

### 2. Terminal Size
- How to get the width of the current terminal window in Go
- Search: **"golang get terminal width size"**
- Search: **"golang syscall terminal size TIOCGWINSZ"**

### 3. Alignment Math
- What does it mean to center a string of width W in a space of width T?
- What does right-alignment mean in terms of padding?
- What is justify alignment — how does it differ from the others?

**Think through the math before writing any code.** Draw it on paper first.

---

## Project Structure

```
ascii-art-justify/
├── main.go
├── banner.go
├── align.go        ← new file for alignment logic
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
go run . "hello"                              → left alignment by default
go run . --align=unknown "hello"              → usage message
```

**Questions to answer before writing anything:**
- After parsing `--align=center`, what arguments remain?
- What are the four valid alignment types? What happens for anything else?
- What is the default alignment when no `--align` flag is given?
- How do you combine this flag with the existing `--output` flag from ascii-art-output?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Scan args for "--align=..."
    //    Extract the alignment type, remove from args
    //    Default to "left" if not present
    //    If value is not one of: center, left, right, justify → print usage and return

    // 2. Also handle "--output=..." if present (from ascii-art-output)

    // 3. Parse remaining args for [STRING] and optional [BANNER]

    // 4. Get terminal width

    // 5. Load banner and render with alignment applied

    // 6. Output to file or stdout
}
```

**Usage message:**
```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

---

## Milestone 2 — Get the Terminal Width

**Goal:** At runtime, detect how wide the terminal window currently is.

**Questions to answer:**
- Which Go package gives you access to terminal dimensions?
- What syscall is used to query the terminal size?
- What should you fall back to if the terminal size cannot be determined?

**Code Placeholder:**
```go
// align.go

func getTerminalWidth() int {
    // Query the terminal dimensions using a syscall
    // Return the number of columns
    // If the query fails, return a sensible default (e.g. 80)
}
```

**Resources:**
- Search: **"golang terminal width TIOCGWINSZ"**
- Search: **"golang sys unix Winsize"**

**Verify:**
- Print the terminal width before rendering
- Resize your terminal window and run again — the number should change

---

## Milestone 3 — Calculate the Width of Rendered Text

**Goal:** Before aligning, you need to know how wide the rendered ASCII art is (in characters).

**Questions to answer:**
- Each character in the banner is some number of columns wide. How do you find out how wide a specific character is?
- The width of a rendered string is the sum of the widths of all its characters. How do you calculate this without rendering first?
- Do all characters in a banner have the same width? Test with `'i'` vs `'W'`.

**Code Placeholder:**
```go
// align.go

func renderedWidth(banner []string, text string) int {
    // For each rune in text:
    //   Get that character's 8 art lines
    //   The width of the character = length of any one of those lines
    //   Add it to the total width
    // Return the total
}
```

**Verify:** Calculate the expected width of `"hello"` in `standard.txt` manually by examining the banner file, then confirm your function returns the same number.

---

## Milestone 4 — Left and Right Alignment

**Goal:**
- Left: no padding (this is already how your render works)
- Right: add enough spaces before each row so the art ends at the terminal's right edge

**Questions to answer:**
- If the terminal is T columns wide and your art is W columns wide, how many spaces do you add before each row for right alignment?
- Where exactly do you add the padding — before the entire output or before each individual row?

**Code Placeholder:**
```go
// align.go

func alignRight(renderedRows []string, termWidth int) []string {
    // For each row in renderedRows:
    //   Calculate padding = termWidth - len(row)
    //   If padding < 0, set to 0 (text wider than terminal)
    //   Prepend that many spaces to the row
    // Return the padded rows
}
```

**Verify:**
```bash
go run . --align=right "hello" shadow | cat -e
```
The art should be pushed to the right side. Compare position against the spec example carefully.

---

## Milestone 5 — Center Alignment

**Goal:** Each row of the ASCII art is horizontally centered in the terminal.

**Questions to answer:**
- If the terminal is T wide and the art is W wide, how many spaces go before each row?
- What do you do with the remainder when the number is odd?

**Code Placeholder:**
```go
// align.go

func alignCenter(renderedRows []string, termWidth int) []string {
    // For each row:
    //   Calculate padding = (termWidth - len(row)) / 2
    //   If padding < 0, set to 0
    //   Prepend that many spaces to the row
    // Return the padded rows
}
```

**Verify:**
```bash
go run . --align=center "hello" standard | cat -e
```
Compare against the spec example. The art should appear in the middle of the terminal.

---

## Milestone 6 — Justify Alignment

**Goal:** When there are multiple words, spread them so the first word starts at the left edge and the last word ends at the right edge, with extra spaces distributed evenly between words.

This is the most complex alignment. Think carefully before coding.

**Questions to answer:**
- How many words does the input have? How do you split by words?
- Each word becomes a block of rendered art. How wide is each word's block?
- The total gap to distribute = `termWidth - sum of all word widths`. How do you split this gap evenly across the spaces between words?
- What do you do when the gap does not divide evenly? Which gaps get an extra space?
- What happens when there is only one word? (It should align left.)
- Each word is 8 rows tall. How do you interleave the padding row by row?

**Code Placeholder:**
```go
// align.go

func alignJustify(banner []string, words []string, termWidth int) []string {
    // 1. If only one word, render it left-aligned and return

    // 2. Render each word into its 8-row block separately

    // 3. Calculate total width used by all words (sum of each word block's width)

    // 4. Calculate total space to distribute between words
    //    gaps = len(words) - 1
    //    totalSpace = termWidth - totalWordWidth
    //    baseSpace = totalSpace / gaps
    //    extraSpaces = totalSpace % gaps  (first `extraSpaces` gaps get one extra space)

    // 5. For each row (0 to 7):
    //    Build the row by joining word rows with the calculated spacing between each pair
    //    Return the 8 combined rows
}
```

**Verify:**
```bash
go run . --align=justify "how are you" shadow | cat -e
```
The first word should start at column 1, the last word should end at the terminal's right edge, spacing between should be distributed as evenly as possible.

---

## Milestone 7 — Refactor render to Return Rows

**Goal:** Your current `render` function returns a joined string. For alignment you need individual rows. Refactor so alignment can work row by row.

**Questions to answer:**
- Should `renderLine` return `[]string` (one element per row) instead of a single string?
- How does this change the way `render` calls `renderLine`?
- Will your previous test cases still pass after the refactor?

**Verify:** After refactoring, all previous test cases with `--align=left` (default) still produce identical output.

---

## Debugging Checklist

- Does the alignment look correct in your terminal but not in the spec example? Make sure you are using the real terminal width, not a hardcoded value.
- Is right alignment off by a few columns? Double-check your `renderedWidth` function — measure with `len(row)` on an actual rendered row, not a calculation.
- Is justify alignment creating uneven spacing? Check your `extraSpaces` logic — the first few gaps should have one more space than the rest.
- Does the output break when you resize the terminal mid-test? That is expected — the terminal width is read once at program start.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Read args | https://pkg.go.dev/os |
| `strings` | Parse flags, split words, build rows | https://pkg.go.dev/strings |
| `fmt` | Print output | https://pkg.go.dev/fmt |
| `syscall` or `golang.org/x/term` | Get terminal width | https://pkg.go.dev/syscall |

---

## Submission Checklist

- [ ] `--align=left` produces left-aligned output (same as default)
- [ ] `--align=right` pushes art to the right edge of the terminal
- [ ] `--align=center` centers art horizontally
- [ ] `--align=justify` spreads words edge to edge with even spacing
- [ ] Single-word justify falls back to left alignment
- [ ] Terminal width is read at runtime — output adapts when window is resized
- [ ] Invalid alignment type prints usage message and exits
- [ ] Works with shadow and thinkertoy banners
- [ ] Compatible with `--output` flag from ascii-art-output
- [ ] All previous ascii-art test cases still pass