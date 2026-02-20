# 🎯 ASCII-Reverse Prerequisites Quiz
## Reverse Engineering · Slice Comparison · Block Parsing · Pattern Matching Without Regex

**Time Limit:** 50 minutes  
**Total Questions:** 28  
**Passing Score:** 22/28 (78%)

> ✅ Pass → You're ready to start ASCII-Reverse  
> ⚠️ This project requires thinking backwards. If you score 22–24, spend extra time on Section 3 before starting.

---

## 📋 SECTION 1: THE REVERSE CONCEPT (5 Questions)

### Q1: ASCII-Art takes text and produces art. ASCII-Reverse does the opposite. Given a file containing rendered ASCII art, what must your program output?

**A)** The banner file that was used  
**B)** The original text string that was rendered  
**C)** A list of ASCII code values  
**D)** The art file reformatted  

<details><summary>💡 Answer</summary>

**B) The original text string that was rendered**

If `standard.txt` was used to render `"Hello"` and the result was saved, `--reverse=file.txt` must output `Hello`. You're going from art → text, instead of text → art.

</details>

---

### Q2: You already have `getCharLines(banner, c)` which returns the 8 art lines for character `c`. How do you use this function to implement the reverse?

**A)** You don't — you need a completely different function  
**B)** You call it for every printable ASCII character (32–126) and compare each result to the block from the file; a match tells you which character it is  
**C)** You call it in reverse order (126 down to 32)  
**D)** You pass the block from the file directly to it  

<details><summary>💡 Answer</summary>

**B) You call it for every printable ASCII character (32–126) and compare each result to the block from the file**

This is a brute-force lookup: for each block in the art file, try all 95 printable characters until one's art matches. Since there's a 1-to-1 mapping (each character has unique art), exactly one will match.

</details>

---

### Q3: Why is it important that you already know the banner file format from ASCII-Art?

**A)** It isn't — the reverse project doesn't use banner files  
**B)** You need to load the same banner to get reference art for comparison. Without knowing the format, you can't load or use it.  
**C)** You only need to know the banner for standard.txt  
**D)** Banner files are different in the reverse project  

<details><summary>💡 Answer</summary>

**B) You need to load the same banner to get reference art for comparison**

The art file was created using one specific banner. To reverse it, you load that same banner and use it as your reference. `getCharLines(banner, 'A')` gives you the expected art for `'A'` in that banner — you compare this against each block in the art file.

</details>

---

### Q4: What determines how many characters are in the art file?

**A)** The number of lines in the file divided by 8  
**B)** Each character in the original string produced one 8-line block (plus blank-line separators for `\n`)  
**C)** The file size in bytes  
**D)** The number of printable ASCII characters (95)  

<details><summary>💡 Answer</summary>

**B) Each character in the original string produced one 8-line block**

If the original was `"Hi"`, the art file has two 8-line blocks. If the original was `"Hi\nBye"`, the file has blocks for `H`, `i`, a blank separator (for `\n`), `B`, `y`, `e`. Group the file into 8-line chunks and you get one chunk per character.

</details>

---

### Q5: The art file for `"hello"` has 5 words × 8 rows = 40 lines. But the file may have more or fewer lines due to trailing newlines. Which approach handles this most robustly?

**A)** Assume exactly 40 lines and crash otherwise  
**B)** Split into lines, filter out empty trailing lines from the split, then group into chunks  
**C)** Count lines and divide by 8, rounding up  
**D)** Read lines until EOF, grouping every 8 into one block  

<details><summary>💡 Answer</summary>

**B) Split into lines, filter out empty trailing lines from the split, then group into chunks**

`strings.Split` on a file ending with `\n` creates an empty final element. If you group without filtering, you get a partial empty block. Filter `""` from the end (not from the middle — blank lines in the middle are meaningful separators for `\n` characters).

</details>

---

## 📋 SECTION 2: SLICE COMPARISON (6 Questions)

### Q6: How do you compare two `[]string` slices for equality in Go?

**A)** `slice1 == slice2` (built-in equality)  
**B)** `reflect.DeepEqual(slice1, slice2)`  
**C)** Write a loop comparing element by element  
**D)** Both B and C work  

<details><summary>💡 Answer</summary>

**D) Both B and C work**

`==` on slices is a compile error (slices aren't comparable). `reflect.DeepEqual` works but adds an import and has overhead. A manual loop is simple, avoids reflection, and is easy to understand:

```go
func slicesEqual(a, b []string) bool {
    if len(a) != len(b) { return false }
    for i := range a {
        if a[i] != b[i] { return false }
    }
    return true
}
```

</details>

---

### Q7: What is the output?
```go
a := []string{"hello", "world"}
b := []string{"hello", "world"}
fmt.Println(a == b)
```

**A)** `true`  
**B)** `false`  
**C)** Compile error — slices cannot be compared with `==`  
**D)** `0`  

<details><summary>💡 Answer</summary>

**C) Compile error — slices cannot be compared with `==`**

In Go, `==` works on comparable types (int, string, struct with no slice fields, etc.). Slices are NOT comparable with `==`. Use `reflect.DeepEqual` or a manual loop.

</details>

---

### Q8: You call `reflect.DeepEqual(block, expected)` where `block` has trailing `\r` on some lines (Windows line endings). What happens?

**A)** It still works — `DeepEqual` ignores whitespace  
**B)** Returns `false` — `"hello\r"` != `"hello"`  
**C)** Panics  
**D)** Returns `true` — `\r` is invisible  

<details><summary>💡 Answer</summary>

**B) Returns `false` — `"hello\r"` != `"hello"`**

`reflect.DeepEqual` compares byte-by-byte. A trailing `\r` makes the strings different. This is why you must `strings.TrimRight(line, "\r")` when loading the art file on Windows, or you'll never match any character.

</details>

---

### Q9: Your `blockToChar` function iterates from ASCII 32 to 126. For a block matching `'A'`, on average how many comparisons are needed before finding the match?

**A)** 1 — it always finds it immediately  
**B)** ~33 — `'A'` is ASCII 65, which is `(65-32)/2 = 16.5` characters from the start... wait, let me recalculate: it's the 34th character (65-32+1 = 34 iterations needed)  
**C)** 95 — always checks all characters  
**D)** The position depends on the search order  

<details><summary>💡 Answer</summary>

**B) 34 — `'A'` is the 34th printable ASCII character (ASCII 65 - 32 = 33 characters after space, so 34th)**

For `'A'` (ASCII 65): you iterate through space (32), `!` (33), `"` (34)... all the way up to `'A'` (65). That's `65 - 32 + 1 = 34` characters checked. For `'~'` (126) it's 95. This brute-force approach is completely fine — it runs once per character in the art file.

</details>

---

### Q10: `blockToChar` returns `(rune, bool)`. What does the `bool` represent and when is it `false`?

**A)** Whether the character is uppercase  
**B)** Whether a matching character was found — `false` means no printable ASCII character's art matched the block  
**C)** Whether the banner file was loaded correctly  
**D)** Whether the block has 8 lines  

<details><summary>💡 Answer</summary>

**B) Whether a matching character was found**

`false` means the block didn't match any of the 95 printable characters — the art file might be corrupted, use a different banner than expected, or contain custom content. Your program should handle this gracefully (skip, print a warning, or use a placeholder).

</details>

---

### Q11: When you iterate with `for c := 32; c <= 126; c++` to find a match, what type do you need to pass to `getCharLines`?

**A)** `int`  
**B)** `byte`  
**C)** `rune`  
**D)** `string`  

<details><summary>💡 Answer</summary>

**C) `rune`**

`getCharLines` takes a `rune`. You must cast: `getCharLines(banner, rune(c))`. Since all printable ASCII values fit in int32 (rune), this cast is safe.

```go
for c := 32; c <= 126; c++ {
    expected := getCharLines(banner, rune(c))
    if slicesEqual(block, expected) {
        return rune(c), true
    }
}
```

</details>

---

## 📋 SECTION 3: FILE PARSING INTO BLOCKS (7 Questions)

### Q12: The art file contains the rendering of `"Hi"` — two characters, each 8 rows tall. How many lines does the file have (assuming no `\n` in original)?

**A)** 8  
**B)** 16  
**C)** 16 + 1 trailing newline = 17 lines when split  
**D)** 18 (9 lines per character)  

<details><summary>💡 Answer</summary>

**C) 16 + 1 trailing newline = 17 elements when split on `"\n"`**

2 characters × 8 rows = 16 rows. The file ends with `\n`, so `strings.Split` gives `["row1", ..., "row16", ""]` — 17 elements, with the last being empty. You must handle (ignore) this trailing empty element.

</details>

---

### Q13: What does an 8-line block look like in the art file for the newline character (`\n` in the original string)?

**A)** 8 lines, each containing only spaces  
**B)** 8 completely empty lines (zero length)  
**C)** A single line saying `"\n"`  
**D)** Nothing — newlines are skipped  

<details><summary>💡 Answer</summary>

**B) 8 completely empty lines (zero length)**

When the original string contains `\n`, the render function outputs 8 empty rows (no art, no spaces). In the art file this appears as 8 consecutive empty lines. Your reversal must detect this pattern and output `"\n"` for the reconstructed string.

Wait — actually it's more nuanced. Look at the render output: for an empty part (a `\n` in the input), `render` prints a single blank line (just `\n`). The art file would then have 1 blank line, not 8. Check the spec carefully, as the exact representation depends on the render implementation.

The safe approach: compare the block against the art for space `' '` too, and check if all lines are empty/spaces. If it's 8 empty-ish lines that don't match any character, it's a separator.

</details>

---

### Q14: You group lines into chunks of 8. The file has 17 lines (after split): lines 0–7 are block 1, lines 8–15 are block 2, line 16 is empty (trailing). What index-based slicing extracts block `n` (0-indexed)?

**A)** `lines[n*8 : n*8+8]`  
**B)** `lines[n*9 : n*9+8]`  
**C)** `lines[n+8 : n+16]`  
**D)** `lines[n*8+1 : n*8+9]`  

<details><summary>💡 Answer</summary>

**A) `lines[n*8 : n*8+8]`**

Block 0: `lines[0:8]`, block 1: `lines[8:16]`. No separator lines in the art file — it's 8 lines per character, packed together. (Unlike the banner file which has 1 separator, making it 9 lines per character.)

</details>

---

### Q15: How many blocks does a file with `N` clean lines (no trailing empty) contain?

**A)** `N`  
**B)** `N * 8`  
**C)** `N / 8`  
**D)** `N / 9`  

<details><summary>💡 Answer</summary>

**C) `N / 8`**

Each block is exactly 8 lines. `N` total lines → `N/8` blocks. For `N=16`: 2 blocks. For `N=40`: 5 blocks. Always ensure `N` is divisible by 8 before processing, or handle partial blocks as an error.

</details>

---

### Q16: What is the correct way to check if a block (8 lines) represents a blank line separator (`\n` in the original)?

**A)** Check if `len(block) == 0`  
**B)** Check if all 8 lines in the block are empty strings or strings of only spaces  
**C)** Compare the block to the art for `'\n'` (ASCII 10)  
**D)** Check if the first line of the block starts with `" "`  

<details><summary>💡 Answer</summary>

**B) Check if all 8 lines in the block are empty strings or strings of only spaces**

`'\n'` is NOT a printable ASCII character and is not in the banner file. So your `blockToChar` will never match it. You must check for the "blank block" pattern separately, before trying to match characters.

```go
func isBlankBlock(block []string) bool {
    for _, line := range block {
        if strings.TrimSpace(line) != "" { return false }
    }
    return true
}
```

</details>

---

### Q17: The art file was created with `shadow.txt`. You load `standard.txt` to reverse it. What happens?

**A)** It works — all banners have the same art  
**B)** No character blocks will match — the shadow art is different from standard art  
**C)** Only letters match, numbers don't  
**D)** It works for A-Z but not for special characters  

<details><summary>💡 Answer</summary>

**B) No character blocks will match — the shadow art is different from standard art**

Each banner has unique art for each character. If you try to match shadow art against standard reference art, nothing will match and `blockToChar` will return `false` for every block. The user must specify which banner was used (or you must auto-detect it).

</details>

---

### Q18: Your `readArtFile` function reads `file.txt` and returns `[][]string` — a slice of 8-line blocks. What is the return type of a single block?

**A)** `string`  
**B)** `[]string` — a slice of 8 strings  
**C)** `[8]string` — a fixed-size array of 8 strings  
**D)** `[][]byte`  

<details><summary>💡 Answer</summary>

**B) `[]string` — a slice of 8 strings**

Use slices, not arrays. The full return type is `[][]string`: a slice of blocks, where each block is a `[]string` of 8 lines. Arrays (`[8]string`) are less idiomatic and harder to pass to `getCharLines` for comparison.

</details>

---

## 📋 SECTION 4: ALGORITHM DESIGN & TRICKY CASES (7 Questions)

### Q19: The original string was `"Hello World"`. The art file contains blocks for `H`, `e`, `l`, `l`, `o`, ` `, `W`, `o`, `r`, `l`, `d`. How many blocks?

**A)** 10  
**B)** 11  
**C)** 12  
**D)** 5 (counting words)  

<details><summary>💡 Answer</summary>

**B) 11**

Every character — including the space — produces a block. `"Hello World"` has 11 characters: 5 + 1 (space) + 5. Each produces an 8-line block. The total file has 11 × 8 = 88 lines.

</details>

---

### Q20: Is the art for the space character `' '` (ASCII 32) always 8 lines of pure spaces? How would you verify this?

**A)** Yes — by definition, space produces blank lines  
**B)** Open the banner file and look at the first character's art (line index 1–8)  
**C)** No — the space character has decorative art  
**D)** It depends on the banner  

<details><summary>💡 Answer</summary>

**B) Open the banner file and look at the first character's art (line index 1–8)**

The space character IS the first printable ASCII character (32). Its art starts at line index 1 in the banner file. Open `standard.txt` and look — you'll see 8 lines of spaces (all empty or all spaces). The exact content (empty vs. spaces of a specific width) matters for your comparison.

This is why you MUST use `getCharLines(banner, ' ')` for comparison rather than assuming what space art looks like.

</details>

---

### Q21: The original was `"Hello\nWorld"`. The art file has blocks for `H`, `e`, `l`, `l`, `o`, **[blank block]**, `W`, `o`, `r`, `l`, `d`. How does the blank block differ from the space block?

**A)** They're identical — you can't tell them apart  
**B)** The blank block is produced by the `\n` rendering (an empty line in the output), which looks different from the 8-line art for the space character  
**C)** The blank block has exactly 1 empty line; the space block has 8 lines  
**D)** The space block always has exactly 5 characters per line; the blank block has 0  

<details><summary>💡 Answer</summary>

**B) The blank block looks different from the space character art**

When `render` processes `\n`, it outputs a single blank line (not 8 rows of art). In the art file, this appears differently from the 8-row art block for `' '`. You'll need to actually check what your render function produces for `\n` to know exactly how to detect it during reversal.

This is why the spec says: "study the banner file and the output format carefully" — you must understand the exact bytes in the file.

</details>

---

### Q22: You've successfully reversed the art and got the characters `['H', 'e', 'l', 'l', 'o']`. How do you combine them into the output string?

**A)** `fmt.Println(chars)`  
**B)** Build a `strings.Builder` and write each rune; at the end call `.String()`  
**C)** `strings.Join(chars, "")`  
**D)** Both B is correct; C won't work because `chars` is `[]rune` not `[]string`  

<details><summary>💡 Answer</summary>

**D) B is correct; `strings.Join` won't work directly on `[]rune`**

`strings.Join` takes `[]string`. Your matched characters are `rune` values. Use `strings.Builder`:

```go
var sb strings.Builder
for _, r := range matched {
    sb.WriteRune(r)
}
fmt.Println(sb.String())
```

Or convert to string inline: `string(r)` for each rune.

</details>

---

### Q23: How should your program handle a block that doesn't match any character?

**A)** Panic with an error  
**B)** Skip the block silently  
**C)** Print a placeholder like `?` or print an error to stderr and continue  
**D)** Stop processing all remaining blocks  

<details><summary>💡 Answer</summary>

**C) Print a placeholder or print an error to stderr and continue**

A non-matching block might mean: wrong banner loaded, corrupted file, or custom art. Printing `?` as a placeholder and continuing is a reasonable approach. Alternatively, write a warning to stderr while still outputting what you can. Never silently skip or crash.

</details>

---

### Q24: For automatic banner detection, you could try loading each banner file and testing if the first block matches any character. What is the problem with this approach?

**A)** There's no problem — it works perfectly  
**B)** All three banners may match the same common characters — the space and some letters look similar across banners  
**C)** You'd need to check all blocks in the file against all three banners — expensive  
**D)** Banner files can't be loaded more than once  

<details><summary>💡 Answer</summary>

**B) Some characters might look similar or identical across different banners**

Auto-detection by trying to match is the right idea, but you need to pick a character that's distinctive in each banner. The space character art might look the same across all banners. Use a distinctive character like `'@'` or `'$'` that has noticeably different art in each banner. Or require the user to specify the banner.

</details>

---

### Q25: The flag is `--reverse=my-file.txt`. What is the precise extraction code?

**A)**
```go
filename := strings.TrimPrefix(os.Args[1], "--reverse")
```
**B)**
```go
filename := strings.TrimPrefix(os.Args[1], "--reverse=")
```
**C)**
```go
parts := strings.Split(os.Args[1], "=")
filename := parts[1]
```
**D)** B and C both work; B is more robust for filenames containing `=`  

<details><summary>💡 Answer</summary>

**D) B and C both work; B is more robust for filenames containing `=`**

Same reasoning as the `--output` flag in ascii-art-output: `TrimPrefix` removes exactly `"--reverse="` and leaves everything after, including any `=` in the filename itself. `strings.Split(arg, "=")[1]` would incorrectly truncate filenames containing `=`.

</details>

---

## 📋 SECTION 5: INTEGRATION (3 Questions)

### Q26: Your program supports both `go run . "hello"` (normal render) and `go run . --reverse=file.txt`. In `main`, what is the cleanest way to distinguish these two modes?

**A)** Check `len(os.Args)` only  
**B)** Check if any argument starts with `"--reverse="` — if yes, enter reverse mode; otherwise, normal render mode  
**C)** Always try reverse first, fall back to render  
**D)** Use a global boolean `isReverse`  

<details><summary>💡 Answer</summary>

**B) Check if any argument starts with `"--reverse="` — if yes, enter reverse mode; otherwise, normal render mode**

```go
for _, arg := range os.Args[1:] {
    if strings.HasPrefix(arg, "--reverse=") {
        filename := strings.TrimPrefix(arg, "--reverse=")
        reverseMode(filename)
        return
    }
}
// else: normal render mode
```

This cleanly separates the two code paths.

</details>

---

### Q27: You run `go run . --reverse=file.txt` but `file.txt` doesn't exist. What should happen?

**A)** Panic  
**B)** Print an empty string  
**C)** Print a meaningful error message and exit with a non-zero code  
**D)** Create an empty `file.txt`  

<details><summary>💡 Answer</summary>

**C) Print a meaningful error message and exit with a non-zero code**

```go
data, err := os.ReadFile(filename)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: cannot read file '%s': %v\n", filename, err)
    os.Exit(1)
}
```

Never crash (panic) when a file is missing — always give a clear, helpful error.

</details>

---

### Q28: After implementing ASCII-Reverse, you run your original `go run . "hello"` test. It must produce identical output to before. What does this require about your code structure?

**A)** Nothing — they're completely separate programs  
**B)** The reverse mode must be in a separate file that doesn't affect the main render path  
**C)** The original render path in `main` must be unchanged — the reverse flag check must come first and `return` early, leaving the normal path intact  
**D)** You must re-test all 3 banner files  

<details><summary>💡 Answer</summary>

**C) The reverse flag check must come first and `return` early**

The "enter reverse mode and return" pattern ensures the normal code path runs only when no `--reverse` flag is present. This is backward-compatible: all existing behavior is preserved, and the new flag adds a new entry point.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 26–28 ✅ | **Exceptional.** Strong algorithmic thinking — start immediately. |
| 22–25 ✅ | **Ready.** Review missed questions, especially Section 3 (block parsing). |
| 17–21 ⚠️ | **Study first.** Work through the algorithm on paper: draw the art file, group blocks, and trace through `blockToChar` manually. |
| Below 17 ❌ | **Not ready.** You need to solidify your understanding of the banner file format and slice operations before attempting the reverse direction. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q5 | Reverse concept: art → text, reusing `getCharLines`, grouping blocks |
| Q6–Q11 | Slice comparison without `==`, `reflect.DeepEqual`, rune casting |
| Q12–Q18 | File parsing into 8-line blocks, blank block detection, return types |
| Q19–Q25 | Space vs blank, `strings.Builder` for runes, non-matching blocks, auto-detection |
| Q26–Q28 | Mode switching in main, error handling, backward compatibility |