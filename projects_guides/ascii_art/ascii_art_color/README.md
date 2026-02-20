# ASCII-Art-Color Project Guide

> **Before you start:** This project builds on ascii-art, ascii-art-fs, and ascii-art-output. All three must be working. Run your test cases before you start.

---

## Objectives

By completing this project you will learn:

1. **ANSI Escape Codes** — How terminals use special character sequences to display colors
2. **Substring Matching** — Finding all occurrences of a pattern within a string
3. **Selective Rendering** — Applying a transformation only to specific characters based on their position
4. **Color Systems** — How RGB, HSL, and ANSI color codes work and how to convert between them
5. **Flag Parsing** — Extracting a color value and an optional substring from a `--color=value` flag

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art-Output (Completed)
- `render` returns a string
- Flag parsing with `strings.HasPrefix` and `strings.TrimPrefix`

### 2. ANSI Escape Codes
- What an ANSI escape sequence is and how a terminal interprets it
- The format for setting a foreground color and resetting it
- Search: **"ANSI escape codes colors terminal"**
- Search: **"golang ANSI color codes"**

### 3. String Searching
- How to find every position where a substring appears in a string
- `strings.Index` vs `strings.Contains` — what is the difference?
- Search: **"golang find all occurrences of substring"**

**Read before starting:**
- Search: **"ANSI color codes list"**
- https://pkg.go.dev/strings

---

## Project Structure

```
ascii-art-color/
├── main.go
├── banner.go
├── color.go        ← new file for color logic
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Understand ANSI Color Codes

**This milestone has no code.** Do not skip it.

Before writing anything, answer these in your terminal:

```bash
echo -e "\033[31mThis is red\033[0m"
echo -e "\033[32mThis is green\033[0m"
echo -e "\033[34mThis is blue\033[0m"
```

**Questions to answer:**
- What does `\033[` mark the start of?
- What does `31` mean? What does `32` mean?
- What does `\033[0m` do and why is it important?
- What is the escape sequence for: red, green, blue, yellow, cyan, magenta, white?
- If you want to use RGB colors, what escape sequence format does your terminal support?

Write down the escape sequences. You will use them directly in your code.

---

## Milestone 2 — Parse the `--color` Flag

**Goal:**
```
go run . --color=red kit "a king kitten have kit"
```
The flag provides a color name. The second argument before the string is the substring to color.

**Questions to answer:**
- After stripping `--color=red`, what arguments remain?
- How do you tell apart the substring to color and the main string?
- What should happen if no substring is given? (The whole string should be colored.)
- What are the valid argument counts with the `--color` flag?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Scan args for "--color=..."
    //    If found: extract the color value, remove from args
    //    If not found: color = "" (no coloring)

    // 2. Parse remaining args for [substring] [STRING] [BANNER]
    //    Rules:
    //      If color is set and 2 remaining args: substring = args[0], string = args[1], banner = default
    //      If color is set and 3 remaining args: substring = args[0], string = args[1], banner = args[2]
    //      If color is set and 1 remaining arg: substring = "", string = args[0] (color whole string)
    //      Handle invalid counts with usage message

    // 3. Load banner, render with color applied, print result
}
```

**Usage message:**
```
Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <substring to be colored> "something"
```

---

## Milestone 3 — Map Color Names to ANSI Codes

**Goal:** Convert a color name like `"red"` into the correct ANSI escape sequence.

**Questions to answer:**
- What data structure is best for mapping color names to escape codes?
- What should happen when the user gives a color name your map does not contain?
- Will you support only named colors, or also hex codes like `#ff0000` or RGB like `rgb(255,0,0)`?

**Code Placeholder:**
```go
// color.go

func colorCode(name string) string {
    // Look up the name in a map of color name → ANSI escape sequence
    // Return the matching escape code, or empty string if not found
}

func resetCode() string {
    // Return the ANSI reset sequence
}
```

**Verify:**
```bash
go run . --color=red "hello"     # entire output should be red
go run . --color=green "hello"   # entire output should be green
```

---

## Milestone 4 — Identify Which Characters to Color

**Goal:** Given the full input string and a substring, find every position in the input where the substring appears. Color only those characters when rendering.

**Questions to answer:**
- How do you find all starting positions of a substring in a string (not just the first)?
- If the substring appears inside a word and as a standalone word, should both be colored? (Yes — see the spec example with `kit` inside `kitten`.)
- Your render function processes the input character by character. How do you know, at each character position, whether it is inside a colored substring?

**Code Placeholder:**
```go
// color.go

func findOccurrences(input string, sub string) []bool {
    // Return a slice of booleans, one per character in input
    // true  = this character is inside an occurrence of sub
    // false = this character should not be colored

    // Steps:
    // 1. Initialize a slice of false values with len(input) entries
    // 2. Find every start index where sub appears in input
    // 3. For each start index, mark positions [start, start+len(sub)) as true
    // 4. Return the slice
}
```

**Verify manually:** For input `"a king kitten have kit"` and substring `"kit"`:
- Which indices should be `true`?
- Mark them out by hand before running any code.

---

## Milestone 5 — Apply Color During Rendering

**Goal:** Pass color information into your render pipeline so colored characters are wrapped in ANSI codes and non-colored characters are not.

**Questions to answer:**
- Your `renderLine` currently loops over characters in the text. Where exactly do you add the ANSI color code — before which line? After which line?
- Remember: each character is 8 lines tall. Do you wrap each of the 8 lines separately or wrap the whole character block?
- What happens if you forget `\033[0m` (the reset) after a colored character?

**Code Placeholder:**
```go
// banner.go

func renderWithColor(banner []string, text string, coloredPositions []bool, color string) string {
    // For each row (0 to 7):
    //   For each character at index i in text:
    //     Get the character's art line for this row
    //     If coloredPositions[i] is true:
    //       Wrap the art line with color code before and reset code after
    //     Append to the row result
    //   Add the row result to the output
    // Return the full output string
}
```

**Verify:**
```bash
go run . --color=red kit "a king kitten have kit"
```
The letters `k`, `i`, `t` inside `kitten` and the word `kit` at the end should be red. Everything else should be the default terminal color.

---

## Milestone 6 — Whole String Coloring

**Goal:** When no substring is provided, the entire rendered output is colored.

**Questions to answer:**
- If substring is empty, what does `findOccurrences` return?
- Is it easier to pass a "color everything" flag, or to make `findOccurrences` return all-true when substring is empty?

**Verify:**
```bash
go run . --color=blue "hello"     # entire output is blue
go run . --color=blue "Hello\nThere"  # both lines are blue
```

---

## Debugging Checklist

- Does color bleed into characters that should not be colored? You are missing the reset code `\033[0m` after a colored section.
- Does the entire output turn colored when you only want part of it? Your `findOccurrences` is marking too many positions as true.
- Does the output look correct in your editor but wrong in the terminal? ANSI codes only work in a real terminal — always test with `go run .`, not by reading output files.
- Does `kit` inside `kitten` get colored? If not, your occurrence search is stopping after the first match. Look up how to find all occurrences, not just the first.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `strings` | Find substrings, parse flags | https://pkg.go.dev/strings |
| `fmt` | Print colored output | https://pkg.go.dev/fmt |
| `os` | Read args | https://pkg.go.dev/os |

---

## Submission Checklist

- [ ] `--color=red` colors the whole string when no substring given
- [ ] `--color=red kit "..."` colors only `kit` occurrences including inside words
- [ ] Color resets correctly after each colored section
- [ ] Works with shadow and thinkertoy banners
- [ ] Works with `\n` in the input
- [ ] Invalid flag format prints usage message
- [ ] All previous ascii-art test cases still pass without color flag
- [ ] At least 6 named colors supported