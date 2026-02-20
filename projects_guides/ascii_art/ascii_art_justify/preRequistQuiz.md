# 🎯 ASCII-Art-Justify Prerequisites Quiz
## Terminal Size · Alignment Math · Justify Algorithm · Refactoring to Return Rows

**Time Limit:** 50 minutes  
**Total Questions:** 28  
**Passing Score:** 22/28 (78%)

> ✅ Pass → You're ready to start ASCII-Art-Justify  
> ⚠️ This is the hardest project in the series. If you score between 22–25, study the justify section carefully before starting.

---

## 📋 SECTION 1: TERMINAL WIDTH & syscall (5 Questions)

### Q1: How do you get the current terminal width in Go at runtime?

**A)** `fmt.TerminalWidth()`  
**B)** Query the terminal dimensions via a syscall (e.g., `TIOCGWINSZ`) using `golang.org/x/term` or the `syscall` package  
**C)** Read it from `os.Environ()`  
**D)** It's always 80 columns  

<details><summary>💡 Answer</summary>

**B) Query the terminal dimensions via a syscall**

The terminal size is a kernel-level property. In Go you can use:
```go
import "golang.org/x/term"
width, height, err := term.GetSize(int(os.Stdout.Fd()))
```
Or directly with `syscall.TIOCGWINSZ`. The key point: terminal width is read at **runtime**, not hardcoded. If the query fails, fall back to a default like `80`.

</details>

---

### Q2: Why must terminal width be read at runtime rather than hardcoded to 80?

**A)** 80 is not a valid terminal width  
**B)** Different users have different terminal sizes, and the same user may resize their window  
**C)** The Go compiler requires it  
**D)** `os.Stdout` doesn't support widths  

<details><summary>💡 Answer</summary>

**B) Different users have different terminal sizes, and the same user may resize their window**

Hardcoding 80 would cause wrong alignment on terminals wider or narrower than 80. The spec requires adapting to the actual terminal. If you run the program in a 200-column terminal, the alignment should use 200 columns.

</details>

---

### Q3: `term.GetSize` can fail (e.g., when stdout is redirected to a file). What should your program do?

**A)** Crash with a helpful error message  
**B)** Return 0 and proceed  
**C)** Fall back to a sensible default width like 80  
**D)** Ask the user to enter their terminal width  

<details><summary>💡 Answer</summary>

**C) Fall back to a sensible default width like 80**

```go
width, _, err := term.GetSize(int(os.Stdout.Fd()))
if err != nil {
    width = 80 // sensible default
}
```

Redirected output (`go run . --align=center "hello" > out.txt`) has no terminal — failing gracefully with a default is correct behavior.

</details>

---

### Q4: The `golang.org/x/term` package is NOT in the standard library. What must you do before using it?

**A)** Nothing — it's available automatically  
**B)** Add it to your `go.mod` with `go get golang.org/x/term`  
**C)** Copy its source code into your project  
**D)** Use `import "terminal"` instead  

<details><summary>💡 Answer</summary>

**B) Add it to your `go.mod` with `go get golang.org/x/term`**

The `golang.org/x/` packages are "extended" standard library maintained by the Go team but versioned separately. Run `go get golang.org/x/term` in your project directory to add it to `go.mod` and `go.sum`.

</details>

---

### Q5: After resizing your terminal window, you run the program again. Which statement is true?

**A)** The program uses the new size because it reads width at startup  
**B)** The program still uses the old size because terminal size is cached  
**C)** The program crashes  
**D)** The program uses the default 80 because resizing invalidates the syscall  

<details><summary>💡 Answer</summary>

**A) The program uses the new size because it reads width at startup**

Each time the program runs, it calls `term.GetSize` fresh. The terminal size at that moment is used. If you resize between runs, the next run sees the new size. If you resize during a run, the program doesn't know — it read the size once at startup.

</details>

---

## 📋 SECTION 2: ALIGNMENT MATH (8 Questions)

### Q6: Terminal width is 80 columns. Your ASCII art is 30 columns wide. How many spaces do you add BEFORE each row for right alignment?

**A)** 30  
**B)** 25  
**C)** 50  
**D)** 80  

<details><summary>💡 Answer</summary>

**C) 50**

`padding = termWidth - artWidth = 80 - 30 = 50`. Prepend 50 spaces to each row so the art ends exactly at column 80.

</details>

---

### Q7: Terminal width is 80. Art is 30 wide. How many spaces do you add BEFORE each row for center alignment?

**A)** 50  
**B)** 25  
**C)** 15  
**D)** 40  

<details><summary>💡 Answer</summary>

**B) 25**

`padding = (termWidth - artWidth) / 2 = (80 - 30) / 2 = 50 / 2 = 25`. Integer division truncates — if the result is odd, the left side gets the smaller amount.

</details>

---

### Q8: Terminal width is 80. Art is 31 wide. Center padding = `(80 - 31) / 2 = 49 / 2 = 24` (integer division). The art is visually shifted slightly left. Is this acceptable?

**A)** No — you must always center exactly, which is impossible when odd  
**B)** Yes — when centering an odd-width element in an even-width space, perfect center is impossible; truncation to the left is the standard behavior  
**C)** You should add 25 spaces and truncate the right side  
**D)** You should use float division and round  

<details><summary>💡 Answer</summary>

**B) Yes — truncation to the left is the standard behavior**

Perfect center is mathematically impossible when `(termWidth - artWidth)` is odd. The standard approach is integer division (truncating the fraction), which shifts the art very slightly to the left. This is what word processors and text editors do too.

</details>

---

### Q9: What is the width (in columns) of the rendered ASCII art for a single character in `standard.txt`? How do you find it programmatically?

**A)** Always 5 — all characters are 5 columns wide  
**B)** Use `len(lines[startLine])` — the length of any one of the character's 8 art lines  
**C)** Divide the total file width by 95 characters  
**D)** Use a fixed constant defined in the banner file  

<details><summary>💡 Answer</summary>

**B) Use `len(lines[startLine])` — the length of any one of the character's 8 art lines**

Each character's width equals the length of its art row (all 8 rows are the same width for a given character). Open `standard.txt` and look: different characters have different widths (e.g., `'!'` is narrower than `'W'`). Always measure, never assume.

</details>

---

### Q10: The text `"Hi"` has characters `'H'` (6 cols) and `'i'` (3 cols). What is the total rendered width?

**A)** 9  
**B)** 10 (9 + 1 space between)  
**C)** 6  
**D)** 18 (doubled for 8 rows)  

<details><summary>💡 Answer</summary>

**A) 9**

Characters are placed directly adjacent with no gap between them. `6 + 3 = 9`. The total rendered width is the sum of each individual character's column width.

</details>

---

### Q11: For left alignment, how much padding do you add?

**A)** `termWidth - artWidth`  
**B)** `(termWidth - artWidth) / 2`  
**C)** 0 — no padding needed  
**D)** `artWidth`  

<details><summary>💡 Answer</summary>

**C) 0 — no padding needed**

Left alignment is the default — art starts at column 1 with no leading spaces. This is what your render function already does. If `--align=left` is specified (or it's the default), no padding logic is needed.

</details>

---

### Q12: Art is wider than the terminal. What should your padding calculation do?

**A)** Crash with an error  
**B)** Set padding to 0 and render the art without truncation  
**C)** Truncate the art to fit  
**D)** Wrap the art to the next line  

<details><summary>💡 Answer</summary>

**B) Set padding to 0 and render the art without truncation**

Always guard: `if padding < 0 { padding = 0 }`. Never add negative spaces. Never truncate art — just render it starting at column 1 and let it overflow. The spec doesn't require truncation handling.

</details>

---

### Q13: Your `renderedWidth` function must measure the width of a text string in ASCII art columns WITHOUT rendering it. What is the correct logic?

**A)** `len(text) * 5` — assume 5 columns per character  
**B)** For each rune in `text`, get one of its 8 art lines and add `len(artLine)` to the total  
**C)** Count the columns in the banner file header  
**D)** `len(text) * 9` (9 lines per character)  

<details><summary>💡 Answer</summary>

**B) For each rune in `text`, get one of its 8 art lines and add `len(artLine)` to the total**

```go
func renderedWidth(banner []string, text string) int {
    total := 0
    for _, ch := range text {
        artLines := getCharLines(banner, ch)
        total += len(artLines[0])  // width = length of any row
    }
    return total
}
```

This correctly handles variable-width characters and works for all three banner files.

</details>

---

## 📋 SECTION 3: THE JUSTIFY ALGORITHM (9 Questions)

*This section is deliberately difficult. Justify is the hardest alignment to implement.*

### Q14: What does "justify" alignment mean for ASCII art?

**A)** Each word is individually centered  
**B)** The first word starts at the left edge, the last word ends at the right edge, and the remaining space is distributed evenly between words  
**C)** All text is aligned to the right edge  
**D)** Words are spread randomly across the line  

<details><summary>💡 Answer</summary>

**B) The first word starts at the left edge, the last word ends at the right edge, and remaining space is distributed evenly between words**

This is the same as "full justification" in word processors. The key constraint: left edge and right edge are both filled, with the gaps between words absorbing all the extra space.

</details>

---

### Q15: Terminal width is 100. You have 3 words with rendered widths: 20, 15, 25. Total word width = 60. How many spaces must you distribute between the words?

**A)** 20  
**B)** 40  
**C)** 25  
**D)** 100  

<details><summary>💡 Answer</summary>

**B) 40**

`totalSpace = termWidth - totalWordWidth = 100 - 60 = 40`. This space must fill the gaps between words. With 3 words there are 2 gaps.

</details>

---

### Q16: From Q15: 40 spaces across 2 gaps. How many spaces per gap?

**A)** 40  
**B)** 20 each  
**C)** 19 in first gap, 21 in second  
**D)** 20 in first gap, 20 in second — but this doesn't fill the terminal if 40 is odd  

<details><summary>💡 Answer</summary>

**B) 20 each**

`baseSpace = totalSpace / gaps = 40 / 2 = 20`. `extraSpaces = 40 % 2 = 0`. Perfect division — each gap gets exactly 20 spaces. The art will span the full 100 columns: 20 + 20 + 15 + 20 + 25 = 100. ✓

</details>

---

### Q17: Terminal width is 100. Words are: 20, 15, 26 (total 61). Space to distribute = 39. Gaps = 2. `39 / 2 = 19` base, `39 % 2 = 1` extra. How are the gaps distributed?

**A)** Gap 1 = 19, Gap 2 = 20  
**B)** Gap 1 = 20, Gap 2 = 19  
**C)** Gap 1 = 19.5, Gap 2 = 19.5  
**D)** Gap 1 = 20, Gap 2 = 20 (one space lost)  

<details><summary>💡 Answer</summary>

**B) Gap 1 = 20, Gap 2 = 19**

The convention is: the first `extraSpaces` gaps get one extra space. `extraSpaces = 1`, so gap 1 gets `19 + 1 = 20` and gap 2 gets `19`. Total: `20 + 15 + 20 + 26 + 19 = 100`. ✓

The first gap getting extra space is the standard typographic convention.

</details>

---

### Q18: What does `39 % 2` compute?

**A)** 19  
**B)** 1  
**C)** 0  
**D)** 78  

<details><summary>💡 Answer</summary>

**B) 1**

`%` is the modulo operator — the remainder after division. `39 / 2 = 19` remainder `1`. This tells you how many gaps get an extra space.

</details>

---

### Q19: You have 4 words and need to justify. `totalSpace = 41`, `gaps = 3`. Calculate `baseSpace` and `extraSpaces`.

**A)** `base = 13`, `extra = 2`  
**B)** `base = 14`, `extra = 0`  
**C)** `base = 13`, `extra = 3`  
**D)** `base = 10`, `extra = 11`  

<details><summary>💡 Answer</summary>

**A) `base = 13`, `extra = 2`**

`41 / 3 = 13` (integer division). `41 % 3 = 2`. So: gap 1 = 14, gap 2 = 14, gap 3 = 13. Total space used: `14 + 14 + 13 = 41`. ✓

</details>

---

### Q20: What should happen when there is only one word and `--align=justify` is used?

**A)** Crash — justify requires at least 2 words  
**B)** Center the single word  
**C)** Left-align the single word (no gaps to distribute)  
**D)** Right-align the single word  

<details><summary>💡 Answer</summary>

**C) Left-align the single word**

With one word there are zero gaps, so `gaps = len(words) - 1 = 0`. You can't divide by zero. The fallback is left alignment — render the word starting from column 1.

Always guard: `if len(words) <= 1 { return leftAlign(rendered) }`

</details>

---

### Q21: Justify alignment must be applied ROW BY ROW (across the 8 rows). Why can't you apply it to the complete rendered string at once?

**A)** You can — it's easier that way  
**B)** Because each row of the output is a separate line. The gap between words must appear on every row, not just the first. You must insert the correct number of spaces into each of the 8 rows between the word blocks.  
**C)** Because `strings.Builder` doesn't support multi-line strings  
**D)** Because ANSI codes interfere with width calculation  

<details><summary>💡 Answer</summary>

**B) The gap between words must appear on every row**

ASCII art characters are 8 rows tall. The gap between words is 8 space-rows stacked. You must construct the output row by row, inserting the space-gap into each row:

```
Row 0: [word1_row0] + [spaces] + [word2_row0] + [spaces] + [word3_row0]
Row 1: [word1_row1] + [spaces] + [word2_row1] + [spaces] + [word3_row1]
...
Row 7: [word1_row7] + [spaces] + [word2_row7] + [spaces] + [word3_row7]
```

</details>

---

### Q22: How do you render each word separately for justify? Which existing function helps?

**A)** You can't — you must rewrite the render function from scratch  
**B)** Use `renderLine(banner, word)` — render each word independently into its 8-row block  
**C)** Use `strings.Split` on the final output  
**D)** Render all words together, then split on spaces  

<details><summary>💡 Answer</summary>

**B) Use `renderLine(banner, word)` — render each word independently into its 8-row block**

This is why the refactor to return `[]string` (one per row) from `renderLine` is important. You call `renderLine` for each word, getting a `[]string` of 8 rows per word. Then you assemble them horizontally row by row with the calculated gaps between them.

</details>

---

## 📋 SECTION 4: REFACTORING TO RETURN ROWS (3 Questions)

### Q23: Currently `renderLine` prints directly. For justify, you need `renderLine` to return `[]string` (one element per row). What is the correct new signature?

**A)** `func renderLine(banner []string, text string) string`  
**B)** `func renderLine(banner []string, text string) []string`  
**C)** `func renderLine(banner []string, text string) ([]string, error)`  
**D)** `func renderLine(text string) []string`  

<details><summary>💡 Answer</summary>

**B) `func renderLine(banner []string, text string) []string`**

Returns exactly 8 strings (one per art row). The caller assembles them however it wants — join with `\n` for normal output, or use them individually for justify alignment.

</details>

---

### Q24: After refactoring `renderLine` to return `[]string`, how do you make the non-justify code path still work?

**A)** Write a second version of `renderLine` that prints  
**B)** The calling function joins the 8 rows with `\n` and adds a final `\n`  
**C)** It automatically works  
**D)** Call `fmt.Println` on the returned slice  

<details><summary>💡 Answer</summary>

**B) The calling function joins the 8 rows with `\n` and adds a final `\n`**

```go
rows := renderLine(banner, text)
output := strings.Join(rows, "\n") + "\n"
fmt.Print(output)
```

This restores the original behavior. The refactor is backward-compatible as long as you add this join step in the non-justify paths.

</details>

---

### Q25: After the refactor, you have `--align=right`. The rendering pipeline should be:

**A)** `renderLine → alignRight → join → print`  
**B)** `alignRight → renderLine → join → print`  
**C)** `join → alignRight → renderLine → print`  
**D)** `renderLine → join → alignRight → print`  

<details><summary>💡 Answer</summary>

**A) `renderLine → alignRight → join → print`**

1. Get the raw rendered rows from `renderLine`
2. Apply padding to each row with `alignRight`
3. Join the padded rows into a final string
4. Print

Alignment is applied to individual rows, not to the final string (that would require parsing the string back into rows).

</details>

---

## 📋 SECTION 5: EDGE CASES & INTEGRATION (3 Questions)

### Q26: Input is `"hello world"` with `--align=justify`. The words are `"hello"` and `"world"`. Is there a gap to distribute?

**A)** No — only 1 gap between 2 words, and it gets all the space  
**B)** Yes — 1 gap that receives all the available space (`termWidth - wordWidths`)  
**C)** No — 2 words means justify behaves like right alignment  
**D)** Yes — 2 gaps even with 2 words  

<details><summary>💡 Answer</summary>

**B) Yes — 1 gap that receives all the available space**

`gaps = len(words) - 1 = 2 - 1 = 1`. The single gap between `"hello"` and `"world"` receives `termWidth - renderedWidth("hello") - renderedWidth("world")` spaces. This pushes `"world"` flush to the right edge.

</details>

---

### Q27: `--align=justify` with input `"Hello\nWorld"` — should both lines be individually justified?

**A)** Yes — each line is justified independently within the terminal width  
**B)** No — only the first line is justified  
**C)** The `\n` causes an error in justify mode  
**D)** Both lines are merged and justified together  

<details><summary>💡 Answer</summary>

**A) Yes — each line is individually justified independently**

Each `\n`-separated part is processed independently. `"Hello"` (1 word) → left-aligned. `"World"` (1 word) → left-aligned. A blank part prints a blank line. Multi-word lines each get their own justify calculation.

</details>

---

### Q28: You've implemented all four alignments. The test runs `go run . "hello"` (no `--align` flag). What alignment must be used?

**A)** Right  
**B)** Center  
**C)** Left (default)  
**D)** Justify  

<details><summary>💡 Answer</summary>

**C) Left (default)**

The spec says `--align=left` is the default when no flag is given. Your program's behavior without the flag must be identical to with `--align=left`. And both must be identical to the original ascii-art/ascii-art-fs output — backward compatibility.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 26–28 ✅ | **Exceptional.** You understand the justify algorithm deeply — start immediately. |
| 22–25 ✅ | **Ready.** Review missed questions (especially the justify math) before starting. |
| 17–21 ⚠️ | **Study first.** The justify algorithm will block you — work through Q14–Q22 on paper with examples. |
| Below 17 ❌ | **Not ready.** Master alignment math and the row-by-row rendering model first. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q5 | `golang.org/x/term`, `term.GetSize`, syscall fallback |
| Q6–Q13 | Alignment math: padding formulas for left/right/center, `renderedWidth` |
| Q14–Q22 | Justify algorithm: space distribution, `%` operator, row-by-row assembly |
| Q23–Q25 | Refactoring `renderLine` to return `[]string`, alignment pipeline order |
| Q26–Q28 | Edge cases: single-word justify, multi-line justify, default alignment |