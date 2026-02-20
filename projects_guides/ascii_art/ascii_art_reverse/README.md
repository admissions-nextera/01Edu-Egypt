# ASCII-Reverse Project Guide

> **Before you start:** This project is the inverse of ascii-art. Instead of text → art, you go art → text. Study the banner file format again carefully before writing anything — you need to understand it in both directions.

---

## Objectives

By completing this project you will learn:

1. **Reverse Engineering a Format** — Reading structured data and working backwards to its source
2. **Pattern Matching Without Regex** — Comparing slices of strings to find which character they correspond to
3. **File Parsing** — Reading a multi-line art file and reconstructing its logical structure
4. **Algorithm Design** — Thinking through a problem from the output backwards to the input

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art (Completed)
- The banner file format and the formula `(ASCII value - 32) * 9 + 1`
- How each character occupies exactly 8 lines in the banner
- `loadBanner`, `getCharLines` — you will reuse these

### 2. File Reading
- `os.ReadFile` — reading a file into memory
- `strings.Split` — splitting content into lines

### 3. Slice Comparison
- How to compare two `[]string` slices for equality in Go
- Search: **"golang compare two string slices"**

### 4. Flag Parsing
- `strings.HasPrefix`, `strings.TrimPrefix` — extracting `--reverse=filename`

---

## Project Structure

```
ascii-reverse/
├── main.go
├── banner.go       ← reuse loadBanner and getCharLines from ascii-art
├── reverse.go      ← new file for the reverse logic
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Parse the `--reverse` Flag

**Goal:**
```
go run . --reverse=file.txt    → reads file.txt, prints the original string
go run .                       → usage message
go run . --reverse             → usage message (no filename given)
```

**Questions to answer:**
- How do you extract the filename from `--reverse=file.txt`?
- What is the valid argument count for this program?
- Should your program still run normally with a single `[STRING]` argument (no flag)?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Check if any argument starts with "--reverse="
    //    If yes: extract the filename, call reverse logic, return
    //    If no: proceed with normal ascii-art rendering (single STRING arg)

    // 2. If "--reverse" is present but has no "=filename" part: print usage and return

    // 3. Handle wrong argument counts with usage message
}
```

**Usage message:**
```
Usage: go run . [OPTION]
EX: go run . --reverse=<fileName>
```

---

## Milestone 2 — Read and Structure the Input File

**Goal:** Read the art file and split it into 8-line blocks, one block per character (or newline).

**Questions to answer:**
- The art file has 8 lines per character. How do you group lines into chunks of 8?
- What does an empty line (or a line with only a newline) mean in the art file? Look at the spec — a `$` on its own line means a blank line in the output, which corresponds to `\n` in the original string.
- How do you handle the newline at the end of the file — does it produce an extra empty block?

**Code Placeholder:**
```go
// reverse.go

func readArtFile(filename string) ([][]string, error) {
    // Read the file

    // Split into lines
    // Note: handle "\r\n" line endings if present

    // Group every 8 lines into one block
    // Each block represents one character or one \n in the original string

    // Return the slice of blocks
}
```

**Verify:** Read the example `file.txt` from the spec and print how many blocks you get. For `"hello"` that is 5 characters = 5 blocks.

---

## Milestone 3 — Match a Block to a Character

**Goal:** Given one 8-line block from the art file and the loaded banner, find which character it represents.

**Questions to answer:**
- You already have `getCharLines(banner, c)` which returns the 8 art lines for any character. How do you use it in reverse — comparing a block to every possible character until you find a match?
- Which characters should you search through? (Printable ASCII: 32 to 126.)
- What should you return if no character matches?

**Code Placeholder:**
```go
// reverse.go

func blockToChar(block []string, banner []string) (rune, bool) {
    // Loop through every printable ASCII character (rune 32 to 126):
    //   Get that character's 8 art lines using getCharLines
    //   Compare them to the block
    //   If they match: return the rune and true

    // If nothing matched: return 0 and false
}
```

**Resources:**
- Search: **"golang compare string slices reflect.DeepEqual"**
- Or: write a simple loop that compares line by line — no need for reflection

**Verify:** Load `standard.txt`, extract the art lines for `'h'` manually using `getCharLines`, then pass them to `blockToChar` and confirm it returns `'h'`.

---

## Milestone 4 — Reconstruct the Original String

**Goal:** Process all blocks from the art file, convert each to a character, and print the result.

**Questions to answer:**
- How do you represent a line break (`\n`) in the art file? What does that block look like compared to a character block?
- What should you do when `blockToChar` returns false (no match found)?

**Code Placeholder:**
```go
// reverse.go

func reverseArt(filename string, banner []string) (string, error) {
    // 1. Read and structure the art file into blocks

    // 2. For each block:
    //    Check if it is a newline block (all 8 lines are empty or just spaces)
    //    If yes: append "\n" to the result
    //    If no: call blockToChar to find the matching character
    //          Append the character to the result

    // 3. Return the result string
}
```

**Verify:**
```bash
# Create file.txt with the art for "hello" (copy from the spec)
go run . --reverse=file.txt
# Should print: hello
```

---

## Milestone 5 — Handle Multiple Lines in the Input

**Goal:** An art file can contain multiple rows of text separated by blank lines (8 lines of empty content). The output should reconstruct the `\n` between them.

**Questions to answer:**
- In the art file, what does a blank line group (8 empty lines) look like compared to the space character art?
- How do you tell the difference between the art for `' '` (space character) and a blank separator?

**Verify:**
```bash
# Create a file with the art for "Hello\nThere"
go run . --reverse=file.txt
# Should print:
# Hello
# There
```

---

## Milestone 6 — Detect the Banner Automatically

**Goal:** The reverse program should try to determine which banner was used to create the art file, rather than always assuming `standard`.

**Questions to answer:**
- How do you detect which banner was used? What would you compare?
- Is it possible to detect this automatically, or should you ask the user to specify?
- If you require the user to specify, how do you add a banner argument to the `--reverse` workflow?

This milestone is open-ended. Think through your approach and implement what makes the most sense to you.

---

## Debugging Checklist

- Does `blockToChar` never find a match? Print one block from your file and the result of `getCharLines(banner, 'h')` side by side — are they identical or do they differ by whitespace or trailing characters?
- Are you getting extra empty characters at the end? Check whether your block-grouping logic creates an empty block from the trailing newline in the file.
- Does a `\n` in the original string not get reconstructed? Make sure you detect blank separator blocks separately from the space character art.
- Are you loading the wrong banner? The art was created with a specific banner — if you load the wrong one, nothing will match.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Read the art file and args | https://pkg.go.dev/os |
| `strings` | Split lines, trim whitespace, compare | https://pkg.go.dev/strings |
| `fmt` | Print the result | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] `--reverse=file.txt` prints the original string
- [ ] Works for single-line art files
- [ ] Works for multi-line art files (with `\n` in original)
- [ ] Space character is reconstructed correctly
- [ ] Special characters (`!`, `?`, `{`, etc.) are reconstructed correctly
- [ ] Invalid flag format prints usage message
- [ ] Program still runs normally with a single `[STRING]` argument
- [ ] No crash when file does not exist — meaningful error message
- [ ] All previous ascii-art test cases still pass