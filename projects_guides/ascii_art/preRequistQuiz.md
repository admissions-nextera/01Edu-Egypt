# 🎯 ASCII-Art Prerequisites Quiz
## File I/O · Runes vs Bytes · strings.Split · ASCII Table · String Building

**Time Limit:** 45 minutes  
**Total Questions:** 25  
**Passing Score:** 20/25 (80%)

> ✅ Pass → You're ready to start ASCII-Art  
> ❌ Fail → Review flagged topics before starting

---

## 📋 SECTION 1: RUNES VS BYTES (6 Questions)

### Q1: What is the output?
```go
s := "Hello"
fmt.Println(len(s))
```

**A)** 5  
**B)** 4  
**C)** 10  
**D)** 8  

<details><summary>💡 Answer</summary>

**A) 5**

`len(s)` returns the number of **bytes**, not characters. For pure ASCII strings these are the same. The trap comes with Unicode — `len("é")` is `2` (two bytes) even though it's one character.

</details>

---

### Q2: What is the output?
```go
s := "Hello"
for i, ch := range s {
    if i == 1 {
        fmt.Printf("%T %v\n", ch, ch)
    }
}
```

**A)** `byte 101`  
**B)** `rune 101`  
**C)** `int32 101`  
**D)** Both B and C — rune is an alias for int32  

<details><summary>💡 Answer</summary>

**D) Both B and C — rune is an alias for int32**

`range` over a string yields `(int index, rune value)`. `rune` is defined as `type rune = int32`. `'e'` has ASCII value `101`.

```go
// These are identical:
var a rune  = 'e'
var b int32 = 'e'
```

</details>

---

### Q3: You need to get the ASCII decimal value of the character `'A'`. Which expression gives you `65`?

**A)** `len('A')`  
**B)** `int('A')`  
**C)** `string('A')`  
**D)** `byte("A")`  

<details><summary>💡 Answer</summary>

**B) `int('A')`**

A rune is already an integer — casting it to `int` just makes the type explicit. `'A'` = 65 in ASCII.

```go
fmt.Println(int('A'))    // 65
fmt.Println('A' - 32)   // 33 (same math the formula uses)
```

</details>

---

### Q4: What is the output?
```go
s := "Go!"
for _, ch := range s {
    fmt.Println(ch)
}
```

**A)** `G o !`  
**B)** `71 111 33`  
**C)** `G`, `o`, `!` — each on its own line  
**D)** `71`, `111`, `33` — each on its own line  

<details><summary>💡 Answer</summary>

**D) `71`, `111`, `33` — each on its own line**

`range` yields runes (integers). `fmt.Println` on an integer prints its decimal value. To print the character itself: `fmt.Println(string(ch))`.

</details>

---

### Q5: What is the output?
```go
ch := 'A'
fmt.Println(string(ch + 1))
```

**A)** `B`  
**B)** `66`  
**C)** `AB`  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**A) `B`**

`'A'` is `65`. `65 + 1 = 66`. `string(66)` converts the code point to the character `'B'`. This is exactly the arithmetic you use to walk through the ASCII table.

</details>

---

### Q6: The printable ASCII characters run from code 32 to code 126. How many printable characters are there?

**A)** 94  
**B)** 95  
**C)** 96  
**D)** 126  

<details><summary>💡 Answer</summary>

**B) 95**

`126 - 32 + 1 = 95`. The `+1` is because both ends are **inclusive**. This matters for your bounds checking — any character outside `32–126` is not in the banner file.

</details>

---

## 📋 SECTION 2: strings PACKAGE & STRING BUILDING (6 Questions)

### Q7: What does `strings.Split("a\nb\nc", "\n")` return?

**A)** `["a", "b", "c"]`  
**B)** `["a\n", "b\n", "c"]`  
**C)** `["a", "\n", "b", "\n", "c"]`  
**D)** `["a b c"]`  

<details><summary>💡 Answer</summary>

**A) `["a", "b", "c"]`**

`strings.Split(s, sep)` removes the separator completely. The result has `len(s) + 1` elements where `len` is the number of separator occurrences. For 2 newlines: 3 parts.

</details>

---

### Q8: A file ends with a newline: `"a\nb\n"`. You call `strings.Split(content, "\n")`. How many elements does the result have?

**A)** 2  
**B)** 3 — the trailing newline creates an empty final element  
**C)** 1  
**D)** Depends on the OS  

<details><summary>💡 Answer</summary>

**B) 3 — the trailing newline creates an empty final element**

```go
parts := strings.Split("a\nb\n", "\n")
// ["a", "b", ""]  ← empty string at end
fmt.Println(len(parts)) // 3
```

This is a critical detail for loading banner files. The trailing empty element must be accounted for or you'll get off-by-one errors in your line indexing.

</details>

---

### Q9: Which approach builds a long string more efficiently inside a loop?

**A)**
```go
result := ""
for _, line := range lines {
    result += line + "\n"
}
```
**B)**
```go
var sb strings.Builder
for _, line := range lines {
    sb.WriteString(line + "\n")
}
result := sb.String()
```
**C)** Both are identical in performance  
**D)** `strings.Join(lines, "\n")`  

<details><summary>💡 Answer</summary>

**B) strings.Builder**

Option A uses `+=` which creates a new string allocation on every iteration — O(n²) total. `strings.Builder` amortizes allocations and is O(n). Option D is also fine but doesn't work when you need to build conditionally across a nested loop (which you do in row-by-row rendering).

</details>

---

### Q10: What is the output?
```go
lines := []string{"row1_h", "row1_i"}
fmt.Println(strings.Join(lines, " "))
```

**A)** `row1_h\nrow1_i`  
**B)** `row1_h row1_i`  
**C)** `["row1_h", "row1_i"]`  
**D)** `row1_hrow1_i`  

<details><summary>💡 Answer</summary>

**B) `row1_h row1_i`**

`strings.Join` places the separator **between** elements, not after the last one. For rendering one row of ASCII art you concatenate character row pieces — `strings.Join(pieces, "")` with an empty separator.

</details>

---

### Q11: You're on Windows and a file uses `\r\n` line endings. You split on `"\n"`. What problem occurs?

**A)** No problem — Go handles this automatically  
**B)** Each line will have a trailing `\r`, which will corrupt your art comparisons  
**C)** The file won't load  
**D)** Lines will be doubled  

<details><summary>💡 Answer</summary>

**B) Each line will have a trailing `\r`, which will corrupt your art comparisons**

After splitting on `"\n"`, each line becomes `"content\r"`. When you compare this to expected art lines or use `len()`, the `\r` adds an extra invisible character. Fix with:
```go
line = strings.TrimRight(line, "\r")
```

</details>

---

### Q12: What is the output?
```go
s := "hello"
fmt.Println(s[:3])
fmt.Println(s[3:])
```

**A)** `hel` and `lo`  
**B)** `hel` and `llo`  
**C)** `hell` and `lo`  
**D)** `hel` and `elo`  

<details><summary>💡 Answer</summary>

**A) `hel` and `lo`**

`s[:3]` = indices 0,1,2 → `"hel"`. `s[3:]` = indices 3,4 → `"lo"`. The split point `3` is the first index NOT included in `s[:3]` and the first index included in `s[3:]`.

</details>

---

## 📋 SECTION 3: FILE I/O (4 Questions)

### Q13: You read a banner file and get `data []byte`. Which line correctly converts it for string processing?

**A)** `text := data.String()`  
**B)** `text := fmt.Sprintf("%s", data)`  
**C)** `text := string(data)`  
**D)** `text := strconv.Itoa(data)`  

<details><summary>💡 Answer</summary>

**C) `text := string(data)`**

Direct cast from `[]byte` to `string`. Option B works but is unnecessarily verbose. Option A doesn't exist on a slice.

</details>

---

### Q14: After calling `os.ReadFile`, when should you check the error?

**A)** Only if you suspect the file might not exist  
**B)** Always — immediately after the call, before using the data  
**C)** At the end of the function  
**D)** Only in production builds  

<details><summary>💡 Answer</summary>

**B) Always — immediately after the call, before using the data**

```go
data, err := os.ReadFile("standard.txt")
if err != nil {
    fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
    os.Exit(1)
}
```

Using `data` when `err != nil` is undefined behavior — `data` may be `nil`, causing a panic downstream.

</details>

---

### Q15: `os.Args[0]` contains the program name. Your program is called with `go run . "hello"`. What is `os.Args[1]`?

**A)** `"go"`  
**B)** `"run"`  
**C)** `"hello"`  
**D)** `"."`  

<details><summary>💡 Answer</summary>

**C) `"hello"`**

When using `go run .`, the Go toolchain compiles and runs the binary. `os.Args[0]` is the compiled temp binary path. `os.Args[1]` is the first argument you passed — `"hello"`.

</details>

---

### Q16: How many arguments does the basic ASCII-Art program expect?

**A)** 0  
**B)** 1 — exactly the string to render  
**C)** 2 — input file and output file  
**D)** Any number  

<details><summary>💡 Answer</summary>

**B) 1 — exactly the string to render**

`len(os.Args)` must equal `2` (program name + 1 argument). If `len(os.Args) != 2`, print usage and return.

</details>

---

## 📋 SECTION 4: ASCII TABLE & THE FORMULA (5 Questions)

### Q17: What is the decimal ASCII value of the space character `' '`?

**A)** 0  
**B)** 32  
**C)** 48  
**D)** 64  

<details><summary>💡 Answer</summary>

**B) 32**

The ASCII table starts at 0 (null) and printable characters begin at 32 (space). This is the anchor for the banner file formula — space is the first character and appears at line index 1 in the banner file.

</details>

---

### Q18: In the banner file, each character occupies 8 art lines plus 1 blank separator line = 9 lines total. The space character starts at line index 1. At which line index does `'!'` (ASCII 33) start?

**A)** 1  
**B)** 9  
**C)** 10  
**D)** 11  

<details><summary>💡 Answer</summary>

**C) 10**

Formula: `startLine = (ASCII - 32) * 9 + 1`  
For `'!'` = ASCII 33: `(33 - 32) * 9 + 1 = 1 * 9 + 1 = 10`

</details>

---

### Q19: Using the formula `startLine = (ASCII - 32) * 9 + 1`, at which line index does `'A'` (ASCII 65) start?

**A)** 297  
**B)** 298  
**C)** 577  
**D)** 586  

<details><summary>💡 Answer</summary>

**B) 298**

`(65 - 32) * 9 + 1 = 33 * 9 + 1 = 297 + 1 = 298`

</details>

---

### Q20: Your `getCharLines` function takes a rune `c` and returns a `[]string` of 8 lines. Which slice expression extracts those 8 lines correctly?

```go
startLine := (int(c) - 32) * 9 + 1
```

**A)** `lines[startLine : startLine+9]`  
**B)** `lines[startLine : startLine+8]`  
**C)** `lines[startLine-1 : startLine+8]`  
**D)** `lines[startLine : startLine+7]`  

<details><summary>💡 Answer</summary>

**B) `lines[startLine : startLine+8]`**

You want exactly 8 lines: indices `startLine`, `startLine+1`, ..., `startLine+7`. In Go slice syntax `[low:high]`, `high` is exclusive, so you need `startLine+8`. The 9th line at `startLine+8` is the blank separator — you skip it.

</details>

---

### Q21: Your program receives the input `"Hello~"`. The `~` character is ASCII 126. Is it in the banner file?

**A)** No — only letters and numbers are in the banner  
**B)** Yes — all printable ASCII characters 32–126 are in the banner  
**C)** Maybe — depends on which banner file  
**D)** No — the banner only goes up to ASCII 122 (`z`)  

<details><summary>💡 Answer</summary>

**B) Yes — all printable ASCII characters 32–126 are in the banner**

All three banner files (standard, shadow, thinkertoy) contain every printable ASCII character from 32 (space) to 126 (`~`). Your program should render any of them without special-casing.

</details>

---

## 📋 SECTION 5: RENDERING LOGIC (4 Questions)

### Q22: You want to render `"Hi"` as ASCII art. Why can you NOT loop like this?

```go
for _, ch := range "Hi" {
    // print all 8 rows of ch, then move on
}
```

**A)** `range` doesn't work on strings  
**B)** Because each character is 8 rows tall — printing H fully before i would stack them vertically, not side by side  
**C)** `range` gives bytes, not runes  
**D)** You can — this is the correct approach  

<details><summary>💡 Answer</summary>

**B) Because each character is 8 rows tall — printing H fully before i would stack them vertically, not side by side**

ASCII art characters must be rendered **row by row across all characters simultaneously**:
```
// Correct:
for row := 0; row < 8; row++ {
    for _, ch := range text {
        // append row `row` of character `ch`
    }
    // print the assembled row
}
```

</details>

---

### Q23: The input is `"Hello\nWorld"` — but this comes from the command line, so what does your program actually receive?

**A)** A string with a real newline character in the middle  
**B)** The two characters `\` and `n` — a backslash followed by the letter n  
**C)** Two separate arguments: `"Hello"` and `"World"`  
**D)** An error, because shells don't pass newlines  

<details><summary>💡 Answer</summary>

**B) The two characters `\` and `n` — a backslash followed by the letter n**

When the user types `go run . "Hello\nWorld"` in the shell, the shell passes the literal characters `\` and `n` inside the string — it does NOT expand them into a real newline. You must split on the **two-character sequence** `"\\n"` in Go:

```go
parts := strings.Split(input, "\\n")
```

</details>

---

### Q24: For input `"A\n\nB"` split on `"\\n"`, how many parts do you get and what are they?

**A)** 2 parts: `"A"` and `"B"`  
**B)** 3 parts: `"A"`, `""`, `"B"`  
**C)** 3 parts: `"A"`, `"\n"`, `"B"`  
**D)** 4 parts: `"A"`, `""`, `""`, `"B"`  

<details><summary>💡 Answer</summary>

**B) 3 parts: `"A"`, `""`, `"B"`**

Two `\n` sequences produce 3 parts with one empty string in the middle. The empty string means "print a blank line" — your render logic must handle it without crashing (no call to renderLine).

</details>

---

### Q25: What should your program output for `go run . ""`?

**A)** A blank line  
**B)** Nothing — no output at all  
**C)** An error message  
**D)** 8 blank lines (one for each row)  

<details><summary>💡 Answer</summary>

**B) Nothing — no output at all**

An empty string has no characters to render and no `\n` sequences to process. The output is empty. This is an edge case you must handle — do not print anything if the input is empty.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 23–25 ✅ | **Excellent.** Strong foundation — start immediately. |
| 20–22 ✅ | **Ready.** Review missed questions, then start. |
| 15–19 ⚠️ | **Almost.** Study your weak sections carefully before starting. |
| Below 15 ❌ | **Not ready.** You'll get stuck on the formula and rendering logic. Review runes, ASCII table, and strings. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q6 | Runes vs bytes, iterating strings, ASCII arithmetic |
| Q7–Q12 | `strings.Split`, `strings.Builder`, `strings.Join`, `\r` handling |
| Q13–Q16 | `os.ReadFile`, error handling, `os.Args` |
| Q17–Q21 | ASCII table, the `(ASCII-32)*9+1` formula, valid character range |
| Q22–Q25 | Row-by-row rendering, `\n` in command-line input, edge cases |