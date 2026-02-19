# ASCII-Art Project Guide

> **Rule before you start:** If you are stuck, search first. Every resource link in this guide points to where the answer lives. Do not paste code from AI — you will not understand it under pressure, and you will not learn the skill.

---

## What You Are Building

A command-line program that takes a string and prints it in large ASCII art letters using one of three provided banner files. Each character in your input is represented by an 8-line tall ASCII drawing read from the banner file.

---

## Before You Write a Single Line

Download and open the three banner files: `standard.txt`, `shadow.txt`, `thinkertoy.txt`.

Stare at `standard.txt` for a few minutes and answer these questions to yourself:

- How many lines tall is each character?
- What separates one character from the next in the file?
- What is the ASCII code of the space character `' '`? Where does it appear in the file?
- If you wanted the letter `'A'`, how would you calculate which line in the file it starts on?

Do not move on until you can answer all four. The entire project depends on understanding the file format.

**Resources:**
- Search: **"ASCII table printable characters"**
- Search: **"golang read file line by line"**
- https://pkg.go.dev/os#ReadFile

---

## Phase 1 — Project Setup

### Checkpoint 1.1 — Structure

```
ascii-art/
├── main.go
├── go.mod
├── standard.txt
├── shadow.txt
└── thinkertoy.txt
```

```bash
go mod init ascii-art
```

---

### Checkpoint 1.2 — Read the Argument

Your program takes exactly one argument: the string to render.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != __ {
        fmt.Println("Usage: go run . <string>")
        return
    }

    input := os.Args[1]
    fmt.Println(input) // temporary
}
```

**Verify before moving on:**
- `go run . "Hello"` prints `Hello`
- `go run . ""` prints an empty line (or nothing) without crashing

---

## Phase 2 — Understanding the Banner File

### Checkpoint 2.1 — Concept: The File Format

Open `standard.txt` and read it carefully. Here is what you will find:

- The very first line is **empty**
- Then comes the space character `' '` — 8 lines of spaces
- Then another empty line
- Then the `'!'` character — 8 lines
- Then another empty line
- And so on through all printable ASCII characters from 32 (space) to 126 (`~`)

This means every character in the banner is exactly **9 lines** in the file: 8 lines of art + 1 empty separator line.

The file starts at ASCII 32 (space). So to find any character `c` in the file:

```
startLine = (ASCII value of c - 32) * 9 + 1
```

That formula gives you the line index (0-based) where the first of the 8 art lines begins.

Write that formula down. You will use it in code soon.

---

### Checkpoint 2.2 — Load the Banner File

Write a function that reads the banner file and returns all its lines as a slice of strings:

```go
import (
    "os"
    "strings"
)

func loadBanner(filename string) ([]string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    lines := strings.Split(string(data), "\n")
    return lines, nil
}
```

**Important:** On some systems the file may use `\r\n` line endings. After splitting, you may need to clean each line. Search: **"golang strings.TrimRight carriage return"**

**Verify before moving on:**
- Call `loadBanner("standard.txt")` and print `len(lines)`
- Count the lines manually: there are 95 printable ASCII characters (32–126), each taking 9 lines, plus a leading empty line. What number do you expect? Does it match?

---

### Checkpoint 2.3 — Extract One Character

Write a function that takes the full lines slice and a character, and returns the 8 art lines for that character:

```go
func getCharLines(lines []string, c rune) []string {
    startLine := (int(c) - 32) * 9 + __

    result := []string{}
    for i := startLine; i < startLine+__; i++ {
        result = append(result, lines[i])
    }
    return result
}
```

Fill in the two blanks:
- The first blank: what offset gets you past the separator line to the first art line?
- The second blank: how many art lines does each character have?

**Verify before moving on (add temporary code in main to test):**
- Extract `'H'` and print its 8 lines — it should look like an H
- Extract `' '` (space) and print its 8 lines — should be 8 lines of spaces
- Extract `'!'` and print its 8 lines

---

## Phase 3 — Rendering a Single Line of Text

### Checkpoint 3.1 — Concept: Building the Output

You cannot print character by character. Each character is 8 lines tall. So for the word `"Hi"`:

```
line 1:  [first line of H] + [first line of i]
line 2:  [second line of H] + [second line of i]
...
line 8:  [eighth line of H] + [eighth line of i]
```

Your approach: loop through rows 0–7. For each row, loop through every character in the input string, take that row from its art lines, and concatenate them all. Then print that combined row.

---

### Checkpoint 3.2 — Render a String (No Newlines Yet)

```go
func renderLine(banner []string, text string) {
    for row := 0; row < 8; row++ {
        line := ""
        for _, c := range text {
            charLines := getCharLines(banner, c)
            line += charLines[__]
        }
        fmt.Println(line)
    }
}
```

Fill in the blank: which element of `charLines` do you want on each row?

**Verify before moving on:**
- `go run . "Hi"` should print H and i side by side across 8 rows
- `go run . "Hello"` should match the expected output shown in the project spec
- `go run . ""` should print nothing

---

## Phase 4 — Handling `\n` in the Input

### Checkpoint 4.1 — Concept: What `\n` Means in the Argument

When the user types `go run . "Hello\nThere"`, the shell passes the literal string `Hello\nThere` — two characters `\` and `n`, not a real newline.

Your program must detect these `\n` sequences and treat them as line breaks in the output.

The simplest approach: split the input on `"\\n"` (the two-character sequence) to get a slice of lines. Then render each line separately.

```go
parts := strings.Split(input, "\\n")
```

After splitting `"Hello\nThere"` you get `["Hello", "There"]`.
After splitting `"Hello\n\nThere"` you get `["Hello", "", "There"]`.

What should an empty string in the parts slice render as? Look at the expected output in the spec for `"Hello\n\nThere"` — there is a blank line between the two words. So an empty part prints one empty line.

---

### Checkpoint 4.2 — Handle the Split Parts

```go
func render(banner []string, input string) {
    parts := strings.Split(input, "\\n")

    for _, part := range parts {
        if part == "" {
            fmt.Println()
            continue
        }
        renderLine(banner, part)
    }
}
```

**Verify before moving on:**

| Input | Expected behavior |
|---|---|
| `"Hello\nThere"` | Hello rendered, then There rendered below it |
| `"Hello\n\nThere"` | Hello, blank line, There |
| `"\n"` | One blank line printed |
| `""` | Nothing printed |

Compare your output against the spec examples character by character using `| cat -e`. The `$` at the end of each line in that output marks the end of the line — every line must end exactly there with no trailing spaces.

---

## Phase 5 — Putting It All Together

### Checkpoint 5.1 — Final main.go

```go
func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . <string>")
        return
    }

    input := os.Args[1]

    banner, err := loadBanner("standard.txt")
    if err != nil {
        fmt.Println("Error loading banner:", err)
        return
    }

    render(banner, input)
}
```

**Verify before moving on — run every example from the spec:**

```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "{Hello There}" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
```

Each one must match the spec output exactly. If even one space is off, find out why before moving on.

---

## Phase 6 — Edge Cases

Before submission, think through each of these and test them:

**Special characters**
- `go run . "!@#$%"` — these are all valid printable ASCII characters. Do they render correctly?
- `go run . "{Hello There}"` — the `{` and `}` characters are in the banner file. Verify they appear.

**Numbers**
- `go run . "123"` — numbers are printable ASCII. Do they work automatically?

**Spaces**
- `go run . "Hello World"` — the space between words must render as the space character art (8 lines of spaces). Does it?

**Only newlines**
- `go run . "\n\n\n"` — should print three blank lines

**Long strings**
- `go run . "Hello There"` — multiple words with a space between them

For each one: predict the output before running it. If reality differs from your prediction, understand why.

---

## Phase 7 — Unit Tests

Write at least these tests:

```go
// main_test.go
package main

import "testing"

func TestGetCharLines(t *testing.T) {
    banner, _ := loadBanner("standard.txt")

    // Test: space character returns 8 lines
    lines := getCharLines(banner, ' ')
    if len(lines) != __ {
        t.Errorf("expected 8 lines, got %d", len(lines))
    }

    // Test: 'A' starts at the correct position in the file
    // What line index should 'A' start on? Calculate it manually first.
    aLines := getCharLines(banner, 'A')
    if len(aLines) != __ {
        t.Errorf("expected 8 lines for A, got %d", len(aLines))
    }
}

func TestLoadBanner(t *testing.T) {
    // Test that the file loads without error
    // Test that the number of lines is what you expect
    __________
}
```

Read: https://go.dev/doc/tutorial/add-a-test

---

## Phase 8 — Debugging Reference

**Output looks right but `cat -e` shows extra spaces at end of lines**

Cause: The banner file lines may have trailing spaces or `\r` characters.
Fix: When loading the banner, trim each line with `strings.TrimRight(line, "\r")`.
Search: **"golang strings TrimRight"**

**Characters appear shifted or wrong**

Cause: Your start line formula is off by one, or you are including the separator line in your 8 lines.
Fix: Print the raw start line index for `' '` (should be 1), `'!'` (should be 10), `'"'` (should be 19). Each is 9 apart. If they aren't, your formula is wrong.

**Empty input crashes the program**

Cause: You are trying to access `os.Args[1]` without checking the length first.
Fix: The argument count check at the top of `main` must happen before any access to `os.Args[1]`.

**`\n` in input not being treated as a line break**

Cause: You might be splitting on `"\n"` (a real newline) instead of `"\\n"` (the two-character sequence backslash-n).
Fix: Use `strings.Split(input, "\\n")`. Print the result of the split to verify you're getting separate strings.

**Program panics with "index out of range"**

Cause: A character in the input has an ASCII value outside 32–126 (e.g. a tab, a non-ASCII unicode character).
Fix: Before calling `getCharLines`, validate that the character is in range. Search: **"golang check if rune is printable ASCII"**

---

## Key Concepts Used

| Concept | Where to Learn It |
|---|---|
| Reading a file into a string | https://pkg.go.dev/os#ReadFile |
| Splitting a string into lines | https://pkg.go.dev/strings#Split |
| Iterating over a string as runes | Search: **"golang range string rune"** |
| Converting a rune to its ASCII integer value | Search: **"golang rune to int"** |
| String concatenation in a loop | Search: **"golang strings.Builder"** (better than `+=` in loops) |
| Command-line arguments | https://pkg.go.dev/os#pkg-variables |
| Writing tests | https://go.dev/doc/tutorial/add-a-test |

---

## A Note on `strings.Builder`

If your output is slow on long inputs, it is because you are using `line += charLines[row]` inside a loop. Every `+=` on a string creates a new string in memory. For short inputs this is fine. For long inputs, use `strings.Builder`:

```go
var sb strings.Builder
for _, c := range text {
    charLines := getCharLines(banner, c)
    sb.WriteString(charLines[row])
}
fmt.Println(sb.String())
```

Read: https://pkg.go.dev/strings#Builder

Understand why it is faster before deciding whether to use it.

---

## Submission Checklist

- [ ] `go run . ""` outputs nothing without crashing
- [ ] `go run . "\n"` outputs one blank line
- [ ] `go run . "Hello"` matches the spec output exactly (verified with `cat -e`)
- [ ] `go run . "Hello\nThere"` renders two separate lines of ASCII art
- [ ] `go run . "Hello\n\nThere"` has a blank line between the two words
- [ ] Numbers, spaces, and special characters render correctly
- [ ] `{`, `}` and other bracket characters render correctly
- [ ] No trailing spaces on any output line (verified with `cat -e`)
- [ ] No crashes on any valid printable ASCII input
- [ ] `loadBanner` handles file not found with a proper error message
- [ ] Unit tests written and passing
- [ ] Only standard Go packages used
- [ ] Code follows good practices (no global variables without reason, functions are small and focused)