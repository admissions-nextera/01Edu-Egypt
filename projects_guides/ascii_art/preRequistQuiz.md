# 🎯 ASCII-Art Prerequisites Quiz
## Bytes vs Runes · File I/O · String Building · Error Handling · os.Args

**Time Limit:** 45 minutes  
**Total Questions:** 28  
**Passing Score:** 22/28 (78%)

> Questions are tagged: 🟢 Easy · 🟡 Medium · 🔴 Hard  
> All topics are general — no specific project knowledge required.

---

## 📋 SECTION 1: BYTES, RUNES, AND STRINGS (9 Questions)

### Q1 🟢 — What is the difference between a `byte` and a `rune` in Go?

**A)** They are identical  
**B)** `byte` is an alias for `uint8` (one raw byte); `rune` is an alias for `int32` (one Unicode code point, which may span 1–4 bytes in UTF-8)  
**C)** `byte` holds characters; `rune` holds numbers  
**D)** `rune` is only used for Russian characters  

<details><summary>💡 Answer</summary>

**B) `byte` = `uint8` (one byte); `rune` = `int32` (one Unicode character)**

```go
var b byte = 'A'      // 65 — one byte
var r rune = 'é'      // 233 — one Unicode code point, but 2 bytes in UTF-8

s := "héllo"
fmt.Println(len(s))            // 6 — bytes
fmt.Println(len([]rune(s)))    // 5 — characters
```

This distinction is critical for text processing. If you treat multi-byte characters as bytes, you'll corrupt them. Always decide: are you working with bytes or characters?

</details>

---

### Q2 🟢 — What does iterating over a string with `for i, r := range s` give you?

**A)** `i` = character index, `r` = byte value  
**B)** `i` = byte offset of the rune's start, `r` = the rune (Unicode code point) at that position  
**C)** `i` = character index, `r` = rune  
**D)** Same as iterating `[]byte(s)`  

<details><summary>💡 Answer</summary>

**B) `i` = byte offset, `r` = rune value**

```go
s := "héllo"
for i, r := range s {
    fmt.Printf("byte offset %d: %c (%d)\n", i, r, r)
}
// byte offset 0: h (104)
// byte offset 1: é (233)   ← é is 2 bytes, so next offset is 3
// byte offset 3: l (108)
// byte offset 4: l (108)
// byte offset 5: o (111)

// To iterate with character index:
for i, r := range []rune(s) {
    fmt.Printf("char index %d: %c\n", i, r)
}
```

`for range` over a string automatically handles UTF-8 decoding. Use it when you care about characters. Use `for i := 0; i < len(s); i++` when you need raw bytes.

</details>

---

### Q3 🟢 — What is the zero value of a `string` in Go? Can you index into it?

**A)** `nil` — indexing panics  
**B)** `""` (empty string) — `len("") == 0`; indexing an empty string panics because there are no bytes  
**C)** `" "` (a space)  
**D)** `"\x00"` (null byte)  

<details><summary>💡 Answer</summary>

**B) `""` — empty string, length 0**

```go
var s string          // ""
fmt.Println(s == "") // true
fmt.Println(len(s))  // 0
// s[0]               // panic: index out of range

// Always check before indexing:
if len(s) > 0 {
    fmt.Println(s[0])
}
```

Unlike C, Go strings are not null-terminated and have no `\x00` zero value. The zero value is a genuinely empty sequence of bytes.

</details>

---

### Q4 🟡 — What does `s[i]` return for a string `s` — a `byte` or a `rune`?

**A)** A `rune` — Go handles Unicode automatically  
**B)** A `byte` (`uint8`) — the raw byte at position `i`, regardless of whether it's part of a multi-byte character  
**C)** A `string` of length 1  
**D)** Depends on the string's content  

<details><summary>💡 Answer</summary>

**B) A `byte` — the raw byte at that index**

```go
s := "héllo"
fmt.Printf("%T %v\n", s[0], s[0])  // uint8 104 (h)
fmt.Printf("%T %v\n", s[1], s[1])  // uint8 195 (first byte of é in UTF-8)
// s[1] is NOT 'é' — it's the raw first byte of é's two-byte encoding

// To get the character at position 1 safely:
r := []rune(s)[1]   // 'é' — correct
```

Indexing a string gives bytes, NOT characters. For multi-byte characters, `s[1]` gives half of `'é'`. This is the #1 source of Unicode bugs in Go string processing.

</details>

---

### Q5 🟡 — What does this print?

```go
s := "hello"
fmt.Println(string(s[0]))
fmt.Println(string([]byte{s[0], s[1]}))
```

**A)** `h` then `he`  
**B)** `104` then `[104 101]`  
**C)** Compile error  
**D)** `h` then `he` — only works for ASCII  

<details><summary>💡 Answer</summary>

**A) `h` then `he`**

- `s[0]` = `byte(104)` = `'h'`
- `string(byte(104))` = `"h"` (converts a single byte to its UTF-8 string)
- `[]byte{s[0], s[1]}` = `[]byte{104, 101}`
- `string([]byte{104, 101})` = `"he"` (converts a byte slice to string)

Converting a `[]byte` to `string` interprets the bytes as UTF-8. This is the correct way to build strings from individual bytes.

</details>

---

### Q6 🟡 — How do you convert between `string` and `[]byte`?

**A)** You can't — they are incompatible types  
**B)** `[]byte(s)` converts string to byte slice; `string(b)` converts byte slice to string — both create copies  
**C)** Use `fmt.Sprintf`  
**D)** Use `unsafe.Pointer`  

<details><summary>💡 Answer</summary>

**B) Explicit conversions — both create copies**

```go
s := "hello"
b := []byte(s)   // copy of the bytes: [104 101 108 108 111]
b[0] = 'H'       // modifies the copy, NOT s
s2 := string(b)  // "Hello" — new string from modified bytes

// Round-trip: string → []byte → modify → string
s3 := string(append([]byte(s), '!'))  // "hello!"
```

Both conversions copy the data — Go strings are immutable. This is safe but has a cost. For performance-critical code that does many conversions, restructure to work with `[]byte` throughout.

</details>

---

### Q7 🟡 — What does `strings.Split("hello\nworld\n", "\n")` return?

**A)** `["hello", "world"]`  
**B)** `["hello", "world", ""]` — trailing delimiter creates an empty last element  
**C)** `["hello\n", "world\n"]`  
**D)** `["hello", "world", "\n"]`  

<details><summary>💡 Answer</summary>

**B) `["hello", "world", ""]` — trailing newline creates an empty string**

```go
parts := strings.Split("hello\nworld\n", "\n")
// ["hello", "world", ""]  ← empty string at end!

// To avoid the empty trailing element:
parts = strings.Split(strings.TrimRight("hello\nworld\n", "\n"), "\n")
// ["hello", "world"]

// Or filter:
lines := strings.Split(content, "\n")
for _, line := range lines {
    if line == "" { continue }
    // process line
}
```

Files often end with a newline, producing an empty final element after splitting. Always handle this when processing file content split by lines.

</details>

---

### Q8 🔴 — What is the output?

```go
s := "hello"
b := []byte(s)
b[0] = 'H'
fmt.Println(s)
fmt.Println(string(b))
```

**A)** `Hello` then `Hello`  
**B)** `hello` then `Hello` — `[]byte(s)` is a copy; `s` is unchanged  
**C)** `Hello` then `Hello` — `b` references `s`'s memory  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) `hello` then `Hello`**

`[]byte(s)` makes a copy of the string's bytes. Modifying `b[0]` changes the copy, not the original string. Go strings are immutable — you can never modify a string in place. To "modify" a string, convert to `[]byte`, change it, then convert back to `string`.

</details>

---

### Q9 🔴 — A file contains ASCII art where each line is exactly 8 characters. How do you safely access the character at column `col` of row `row`?

**A)** `content[row][col]`  
**B)** Split by `"\n"`, then `lines[row][col]` — gives a `byte`, which is fine for pure ASCII  
**C)** `[]rune(content)[row*8+col]`  
**D)** `content[row*8+col]`  

<details><summary>💡 Answer</summary>

**B) Split into lines, then index — gives a `byte`, correct for pure ASCII**

```go
lines := strings.Split(content, "\n")
if row >= len(lines) || col >= len(lines[row]) {
    // bounds check
}
ch := lines[row][col]   // byte — safe for pure ASCII (0-127)
```

For pure ASCII content (values 0–127), one byte = one character, so `lines[row][col]` works correctly. If the content could contain UTF-8 multi-byte characters, you'd need to convert to `[]rune` first. Always bounds-check before indexing.

</details>

---

## 📋 SECTION 2: FILE READING AND os.Args (7 Questions)

### Q10 🟢 — What is the idiomatic way to read an entire file as a string in Go?

**A)** `os.ReadFile(name)` then `string(data)`  
**B)** `ioutil.ReadFile(name)` (deprecated since Go 1.16)  
**C)** Both A and B work; A is the modern approach  
**D)** `bufio.ReadAll(name)`  

<details><summary>💡 Answer</summary>

**C) `os.ReadFile` is the modern way; `ioutil.ReadFile` is deprecated**

```go
data, err := os.ReadFile("banner.txt")
if err != nil {
    return fmt.Errorf("could not read file: %w", err)
}
content := string(data)   // convert to string for text processing
lines := strings.Split(content, "\n")
```

`ioutil.ReadFile` still works but is deprecated since Go 1.16. Use `os.ReadFile` in all new code.

</details>

---

### Q11 🟢 — How do you check if a file exists before opening it?

**A)** `os.FileExists("file.txt")`  
**B)** `os.Exists("file.txt")`  
**C)** Try to open it and check the error: `_, err := os.Stat("file.txt"); os.IsNotExist(err)`  
**D)** `file.Exists()`  

<details><summary>💡 Answer</summary>

**C) `os.Stat` + `os.IsNotExist`**

```go
_, err := os.Stat("banner.txt")
if os.IsNotExist(err) {
    fmt.Println("file not found")
    return
}
if err != nil {
    fmt.Println("other error:", err)
    return
}
// file exists — proceed to open
```

There's no `os.FileExists` in Go. The idiomatic pattern is to attempt the operation and handle the error, rather than pre-checking. For reads, just `os.ReadFile` and check the error directly — the error will tell you if the file doesn't exist.

</details>

---

### Q12 🟡 — Your program requires exactly one command-line argument. How do you validate this?

**A)**
```go
if os.Args[0] == "" { ... }
```
**B)**
```go
if len(os.Args) != 2 {
    fmt.Fprintln(os.Stderr, "Usage: program <argument>")
    os.Exit(1)
}
```
**C)**
```go
if os.Args == nil { ... }
```
**D)**
```go
if os.Args[1] == "" { ... }
```

<details><summary>💡 Answer</summary>

**B) Check `len(os.Args) != 2`**

```go
// os.Args[0] = program name (always present)
// os.Args[1] = first user argument
// len(os.Args) == 1 means no user arguments

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "Usage: program <text>")
        os.Exit(1)
    }
    input := os.Args[1]
    // ...
}
```

Always write to `os.Stderr` for error/usage messages, not `os.Stdout`. Exit with code 1 (or non-zero) to signal failure to the shell.

</details>

---

### Q13 🟡 — What does `os.Exit(1)` do and when should you use it?

**A)** Returns `1` from `main`  
**B)** Terminates the program immediately with exit code 1, bypassing all deferred functions — use for unrecoverable errors or invalid arguments  
**C)** Same as `return` in `main`  
**D)** Sends signal 1 to the process  

<details><summary>💡 Answer</summary>

**B) Terminates immediately — deferred functions do NOT run**

```go
func main() {
    defer fmt.Println("this will NOT print")
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "Usage: program <text>")
        os.Exit(1)  // terminates immediately
    }
}
```

Exit code conventions: `0` = success, non-zero = failure. `os.Exit(1)` is for "invalid usage or arguments." `os.Exit(2)` is sometimes used for "can't open file." Programs that exit with 0 on failure make scripting and CI difficult.

</details>

---

### Q14 🟡 — A file is read into `content`. After `strings.Split(content, "\n")`, the last element is sometimes `""`. How do you handle this robustly?

**A)** It never happens  
**B)** Use `strings.TrimRight(content, "\n")` before splitting, or check `if line == "" { continue }` when processing  
**C)** Use `strings.SplitAfter` instead  
**D)** Always remove the first element  

<details><summary>💡 Answer</summary>

**B) Trim trailing newline before splitting, or skip empty lines**

```go
// Option 1: trim before split (cleanest):
lines := strings.Split(strings.TrimRight(content, "\n"), "\n")

// Option 2: filter after split:
allLines := strings.Split(content, "\n")
var lines []string
for _, l := range allLines {
    if l != "" {
        lines = append(lines, l)
    }
}

// Option 3: use strings.Fields for whitespace-separated tokens
```

Text files almost always end with `\n`. Failing to handle this produces an extra empty string that corrupts array indexing. This is one of the most common bugs in file-reading programs.

</details>

---

### Q15 🔴 — What does this code do, and is it correct?

```go
data, err := os.ReadFile("banner.txt")
if err != nil {
    fmt.Println("Error:", err)
}
lines := strings.Split(string(data), "\n")
```

**A)** Correct — reads the file and splits into lines  
**B)** Bug — if `ReadFile` fails, `data` is `nil` and `string(nil)` panics  
**C)** Bug — if `ReadFile` fails, `data` is `nil` but `string(nil)` = `""` so `Split` gives `[""]` — the error is printed but execution silently continues with an empty/wrong result  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**C) Bug — error is printed but execution continues incorrectly**

```go
// WRONG — continues after error:
data, err := os.ReadFile("banner.txt")
if err != nil {
    fmt.Println("Error:", err)
    // missing return! falls through with data == nil
}
lines := strings.Split(string(data), "\n")  // lines == [""] — wrong

// CORRECT — return on error:
data, err := os.ReadFile("banner.txt")
if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
    os.Exit(1)  // or return err if inside a function
}
lines := strings.Split(string(data), "\n")
```

`string(nil)` = `""` in Go, so it doesn't panic — but continuing after an error silently produces wrong results. Always `return` or `os.Exit` after handling an error.

</details>

---

### Q16 🔴 — A banner file stores each character as 8 lines of 8 bytes. How do you extract the 8 lines for character `'A'` given the file has all printable ASCII (32–126)?

**A)**
```go
index := 'A'
lines[index*8 : index*8+8]
```
**B)**
```go
index := int('A') - 32  // 65-32=33
charLines := lines[index*8 : index*8+8]
```
**C)**
```go
index := int('A')
charLines := lines[index : index+8]
```
**D)**
```go
charLines := lines['A']
```

<details><summary>💡 Answer</summary>

**B) Subtract the ASCII offset of the first printable character (space = 32)**

```go
// Printable ASCII starts at space (32), ends at tilde (126)
// 'A' = 65, offset = 65 - 32 = 33
// In an 8-line-per-character file:
// char 32 (space) → lines 0-7
// char 33 ('!')   → lines 8-15
// char 65 ('A')   → lines 33*8 to 33*8+7 = lines 264-271

char := 'A'
index := int(char) - 32
startLine := index * 8
charLines := lines[startLine : startLine+8]
```

The offset (`- 32`) is the key insight — the file doesn't store characters 0–31 (control characters), so character indexing starts at ASCII 32 (space). Getting this arithmetic wrong is the most common bug in this type of program.

</details>

---

## 📋 SECTION 3: STRING BUILDING AND OUTPUT (7 Questions)

### Q17 🟢 — What is the fastest way to build a result string from many small pieces in a loop?

**A)** `result += piece` on each iteration  
**B)** `var b strings.Builder; b.WriteString(piece); result := b.String()`  
**C)** `fmt.Sprintf("%s%s", result, piece)`  
**D)** `append([]byte(result), piece...)` then convert at the end  

<details><summary>💡 Answer</summary>

**B) `strings.Builder`**

```go
var b strings.Builder
for i := 0; i < 1000; i++ {
    b.WriteString("line ")
    b.WriteString(strconv.Itoa(i))
    b.WriteByte('\n')
}
result := b.String()
```

Option A creates a new string on every iteration — O(n²) total allocations. `strings.Builder` uses a growing byte buffer internally — O(n) total. For a few concatenations, `+` is fine. For loops, always use `strings.Builder`.

</details>

---

### Q18 🟢 — How do you write to standard error in Go?

**A)** `fmt.Println("error")` — automatically goes to stderr  
**B)** `fmt.Fprintln(os.Stderr, "error")`  
**C)** `os.Stderr.Print("error")`  
**D)** Both B and C work  

<details><summary>💡 Answer</summary>

**D) Both `fmt.Fprintln(os.Stderr, ...)` and `fmt.Fprintf(os.Stderr, ...)` work**

```go
// Error/usage messages go to stderr:
fmt.Fprintln(os.Stderr, "Error: invalid input")
fmt.Fprintf(os.Stderr, "Usage: %s <text> [font]\n", os.Args[0])

// Normal output goes to stdout:
fmt.Println("result")
fmt.Fprintf(os.Stdout, "result: %s\n", output)
```

Separating stdout and stderr allows users to redirect them independently: `./program 2>errors.txt` captures errors separately, while `./program > output.txt` captures only the result. Programs that mix output and errors into stdout are harder to use in pipelines.

</details>

---

### Q19 🟡 — How does `fmt.Fprintf(w, format, args...)` differ from `fmt.Sprintf(format, args...)`?

**A)** `Fprintf` is faster  
**B)** `Fprintf` writes to an `io.Writer` (file, `os.Stdout`, `http.ResponseWriter`); `Sprintf` returns the formatted string without writing anywhere  
**C)** `Sprintf` is deprecated  
**D)** They produce different output  

<details><summary>💡 Answer</summary>

**B) `Fprintf` writes to a writer; `Sprintf` returns a string**

```go
// Fprintf — writes directly to a destination:
fmt.Fprintf(os.Stdout, "Hello, %s!\n", name)  // writes to stdout
fmt.Fprintf(os.Stderr, "Error: %v\n", err)    // writes to stderr
fmt.Fprintf(file, "Line: %d\n", n)            // writes to a file

// Sprintf — produces a string you can use later:
msg := fmt.Sprintf("Hello, %s!", name)
lines = append(lines, fmt.Sprintf("%d: %s", i, line))
```

Use `Fprintf` to avoid creating an intermediate string when you're writing directly to a destination. Use `Sprintf` when you need the string value itself.

</details>

---

### Q20 🟡 — What does `fmt.Print` vs `fmt.Println` vs `fmt.Printf` do differently?

**A)** They are identical  
**B)** `Print` outputs without newline or formatting; `Println` adds a newline and spaces between args; `Printf` uses format verbs (`%s`, `%d`) — no automatic newline  
**C)** `Print` is for integers; `Println` for strings; `Printf` for floats  
**D)** `Printf` always adds a newline  

<details><summary>💡 Answer</summary>

**B) No format vs newline vs format verbs**

```go
fmt.Print("hello", "world")      // helloworld  (no space, no newline)
fmt.Println("hello", "world")    // hello world\n (space between, newline at end)
fmt.Printf("%s %s\n", "hello", "world") // hello world\n (format verbs, explicit \n)

// Common verbs:
// %s — string, %d — integer, %f — float, %v — default format
// %q — quoted string, %T — type name, %p — pointer address
```

`Println` adds spaces between arguments automatically and always adds a trailing newline. `Printf` gives you full control but requires explicit `\n`. `Print` is rarely the right choice.

</details>

---

### Q21 🟡 — You want to join all rows of a 2D output (each row stored in a `[]string`) into a final string separated by newlines. What's the idiomatic approach?

**A)** Loop and concatenate with `+`  
**B)** `strings.Join(rows, "\n")`  
**C)** `fmt.Sprintf("%v", rows)`  
**D)** `bytes.Join`  

<details><summary>💡 Answer</summary>

**B) `strings.Join(rows, "\n")`**

```go
rows := []string{"##   ##", "#####", "##   ##"}

// Idiomatic — one allocation:
result := strings.Join(rows, "\n")
fmt.Println(result)

// Equivalent but less efficient:
var b strings.Builder
for i, row := range rows {
    if i > 0 { b.WriteByte('\n') }
    b.WriteString(row)
}
```

`strings.Join` is the cleanest and most efficient way to build a newline-delimited string from a slice. It makes exactly one allocation for the final string.

</details>

---

### Q22 🔴 — What is the output of this code?

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Print("*")
    }
    fmt.Println()
}
```

**A)** `*********` on one line  
**B)** Three lines, each `***`  
**C)** `* * * * * * * * *` with spaces  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) Three lines of `***`**

```
***
***
***
```

`fmt.Print("*")` writes `*` without a newline. The inner loop writes three `*` in a row. `fmt.Println()` with no arguments writes just a newline, ending the row. This pattern — inner loop for columns, outer loop for rows — is fundamental to 2D text output.

</details>

---

### Q23 🔴 — How do you correctly align ASCII art characters when building multi-line output?

**A)** Concatenate each character's lines directly: `row += charLines[i]`  
**B)** Build the output row by row: for each output row, concatenate that row of every character — then join all output rows with `\n`  
**C)** Print each character completely before the next  
**D)** Use a 2D array and print column by column  

<details><summary>💡 Answer</summary>

**B) Build row by row — one output row collects that row from each character**

```go
// Each char has 8 art lines. Output has 8 rows.
// Row 0 of output = row 0 from 'H' + row 0 from 'i'
// Row 1 of output = row 1 from 'H' + row 1 from 'i'
// etc.

var result strings.Builder
for row := 0; row < 8; row++ {
    for _, ch := range inputText {
        index := int(ch) - 32
        result.WriteString(charLines[index][row])
    }
    result.WriteByte('\n')
}
```

This is the core algorithm. If you print each character fully before moving to the next, you get characters stacked vertically, not side by side. Building row-by-row is the only way to get horizontal layout.

</details>

---

## 📋 SECTION 4: ERROR HANDLING (5 Questions)

### Q24 🟢 — What is the idiomatic Go pattern for handling an error from a function call?

**A)** `try { } catch { }`  
**B)** Check `if err != nil` immediately after the call and handle or return  
**C)** Ignore errors for simplicity  
**D)** Use `panic` and `recover`  

<details><summary>💡 Answer</summary>

**B) `if err != nil` immediately after each call**

```go
data, err := os.ReadFile("file.txt")
if err != nil {
    return fmt.Errorf("reading file: %w", err)
}

n, err := strconv.Atoi(os.Args[1])
if err != nil {
    fmt.Fprintln(os.Stderr, "argument must be a number")
    os.Exit(1)
}
```

Go has no exceptions. Errors are explicit return values. The pattern is repetitive by design — it forces you to handle every error at the point it occurs. Never use `_` to silently discard errors in production code.

</details>

---

### Q25 🟡 — What does `fmt.Errorf("reading %s: %w", filename, err)` do differently from `fmt.Errorf("reading %s: %v", filename, err)`?

**A)** Nothing — `%w` and `%v` are identical  
**B)** `%w` wraps the error, allowing `errors.Is` and `errors.As` to inspect the original error; `%v` just formats it as a string, losing the original error type  
**C)** `%w` is for warnings; `%v` is for errors  
**D)** `%w` includes a stack trace  

<details><summary>💡 Answer</summary>

**B) `%w` wraps for introspection; `%v` just formats as string**

```go
// %w — wraps the error (Go 1.13+):
err := fmt.Errorf("reading %s: %w", filename, os.ErrNotExist)
errors.Is(err, os.ErrNotExist)  // true — can inspect the wrapped error

// %v — formats as string:
err2 := fmt.Errorf("reading %s: %v", filename, os.ErrNotExist)
errors.Is(err2, os.ErrNotExist) // false — original error info lost
```

Use `%w` when callers might need to check the error type with `errors.Is`. Use `%v` when you just want to include the error message in a log and don't need type checking.

</details>

---

### Q26 🟡 — What is the difference between `log.Fatal` and `fmt.Fprintln(os.Stderr, ...); os.Exit(1)`?

**A)** They are identical  
**B)** `log.Fatal` prefixes the message with a timestamp and calls `os.Exit(1)`; the manual approach gives you full control over the message format  
**C)** `log.Fatal` sends an email  
**D)** `log.Fatal` panics instead of exiting  

<details><summary>💡 Answer</summary>

**B) `log.Fatal` adds timestamp + date prefix, then exits**

```go
log.Fatal("cannot open file")
// Output: 2024/01/15 10:23:45 cannot open file
// Then: os.Exit(1)

// Manual — clean output, no timestamp:
fmt.Fprintln(os.Stderr, "Error: cannot open file")
os.Exit(1)
```

For user-facing CLI tools, the timestamp prefix from `log.Fatal` often looks out of place. Use `fmt.Fprintln(os.Stderr, ...)` for clean user-facing errors. Use `log` for server applications and debugging where timestamps are helpful.

</details>

---

### Q27 🔴 — What is the difference between `panic` and `os.Exit` for handling an unrecoverable error?

**A)** They are equivalent  
**B)** `panic` unwinds the stack and runs all `defer` functions before terminating; `os.Exit` terminates immediately without running defers. `panic` also prints a stack trace; `os.Exit` does not.  
**C)** `panic` can be caught with `recover`; `os.Exit` can't — both print a stack trace  
**D)** Use `panic` for runtime errors, `os.Exit` for logic errors  

<details><summary>💡 Answer</summary>

**B) `panic` runs defers + prints stack trace; `os.Exit` terminates immediately**

```go
// panic — defers run, stack trace printed:
defer cleanup()
panic("something impossible happened")
// cleanup() RUNS, then stack trace, then exit

// os.Exit — immediate, no defers:
defer cleanup()
os.Exit(1)
// cleanup() does NOT run
```

For programs: use `os.Exit(1)` for expected error conditions (wrong arguments, file not found). Use `panic` for programming errors that should never happen ("this code path should be unreachable"). Never `panic` for user input errors.

</details>

---

### Q28 🔴 — Is this error handling correct?

```go
result, _ := strconv.Atoi(os.Args[1])
lines := bannerLines[result*8 : result*8+8]
```

**A)** Yes — the underscore ignores the unneeded error  
**B)** No — if `os.Args[1]` is not a valid integer, `result` is `0` and the slice expression silently uses index 0, producing wrong output with no error  
**C)** Compile error — `_` can't be used here  
**D)** Yes — `strconv.Atoi` never fails for typical inputs  

<details><summary>💡 Answer</summary>

**B) Bug — silent wrong result when input is invalid**

```go
// WRONG:
result, _ := strconv.Atoi(os.Args[1])  // "abc" → result = 0, err ignored
lines := bannerLines[0:8]              // silently uses first character

// CORRECT:
result, err := strconv.Atoi(os.Args[1])
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: '%s' is not a valid number\n", os.Args[1])
    os.Exit(1)
}
```

`_` is sometimes appropriate (discarding a known-safe error), but discarding `strconv.Atoi`'s error and using the zero value is always wrong — the zero value `0` is a plausible real input, making the bug hard to detect.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 26–28 ✅ | **Excellent** — bytes, strings, and file I/O are solid. |
| 22–25 ✅ | **Ready** — review any sections you missed. |
| 17–21 ⚠️ | **Study first** — byte vs rune distinction and error handling need more attention. |
| Below 17 ❌ | **Not ready** — work through the Go Tour strings section and the `os`, `strings`, `strconv` package docs. |

---

## 🔍 Review Map

| Missed | Topic to Study |
|---|---|
| Q1–Q9 | `byte` vs `rune`, `for range` string, `string([]byte)`, `strings.Split` trailing empty, indexing gives bytes |
| Q10–Q16 | `os.ReadFile`, `os.Stat`, `os.Args` validation, `defer f.Close()`, trailing newline handling |
| Q17–Q23 | `strings.Builder`, `fmt.Fprintf` vs `Sprintf`, `fmt.Print/Println/Printf`, `strings.Join`, row-by-row output |
| Q24–Q28 | `if err != nil` pattern, `%w` vs `%v`, `log.Fatal` vs `os.Exit`, `panic` vs `os.Exit`, silent zero-value bug |