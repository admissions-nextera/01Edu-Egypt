# 🔥 Go File I/O, Strings & Parsing — Medium to Hard Quiz
## Master the Pipeline! 💪

---

## BLOCK 1 — File Operations (os package)

### Problem 1: os.Args Indexing
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(len(os.Args))
    fmt.Println(os.Args[0])
}
```
**Run as:** `go run main.go input.txt output.txt`

**Question:** What does each line print?

**Answer:**
- `3`
- The path to the compiled binary (e.g. `/tmp/go-build.../main` or `main`)

**Explanation:**
- `os.Args[0]` is always the **program name/path** — not the first user argument
- `os.Args[1]` = `"input.txt"`, `os.Args[2]` = `"output.txt"`
- `len(os.Args)` = 3 (program + 2 arguments)
- Accessing `os.Args[3]` would panic with index out of range!

**Key Concept:** `os.Args[0]` is the binary itself — user args start at index 1!

---

### Problem 2: Args Validation Trap
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: program <input> <output>")
        os.Exit(1)
    }
    input := os.Args[1]
    output := os.Args[2]
    fmt.Println(input, output)
}
```
**Run as:** `go run main.go`

**Question:** What happens? What's the exit code?

**Answer:**
- Prints: `Usage: program <input> <output>`
- Exits with code `1` (non-zero = error)
- The lines after `os.Exit(1)` are **never reached**

**Explanation:**
- `os.Exit(1)` terminates immediately — no deferred functions run!
- Exit code `0` = success, any non-zero = failure
- Checking `len(os.Args)` before indexing prevents panics
- This is the standard guard pattern for CLI tools

**Key Concept:** Always validate `len(os.Args)` before indexing — and `os.Exit` skips deferred calls!

---

### Problem 3: ReadFile Error Handling
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("nonexistent.txt")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    fmt.Println(string(data))
}
```
**Question:** What type is `data`? What does `string(data)` do? What prints when file doesn't exist?

**Answer:**
- `data` is `[]byte` (a byte slice)
- `string(data)` converts the byte slice to a UTF-8 string
- Prints something like: `Error: open nonexistent.txt: no such file or directory`

**Explanation:**
- `os.ReadFile` returns `([]byte, error)` — the entire file content as bytes
- Always check `err != nil` before using `data`
- Converting `[]byte` → `string` is a zero-allocation operation in this context
- The error message includes the OS-level description

**Key Concept:** `os.ReadFile` returns `[]byte` — always check the error before using the data!

---

### Problem 4: WriteFile Permissions
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    content := "Hello, World!\n"
    err := os.WriteFile("output.txt", []byte(content), 0644)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    fmt.Println("Written successfully")
}
```
**Question:** What does `0644` mean? What does `[]byte(content)` do? Does WriteFile append or overwrite?

**Answer:**
- `0644` is a Unix file permission: owner can read/write (6), group can read (4), others can read (4)
- `[]byte(content)` converts the string to a byte slice (required by WriteFile)
- `WriteFile` **overwrites** the entire file — it does NOT append!

**Explanation:**
- File permissions use octal notation (prefix `0`)
- `0644` is standard for readable files: `-rw-r--r--`
- `os.WriteFile` creates the file if it doesn't exist, or **truncates and overwrites** if it does
- To append, you'd use `os.OpenFile` with `os.O_APPEND` flag

**Key Concept:** `os.WriteFile` always overwrites — use `os.OpenFile` with `O_APPEND` to append!

---

### Problem 5: ReadFile + WriteFile Round Trip
```go
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    os.WriteFile("test.txt", []byte("hello world\n"), 0644)
    
    data, _ := os.ReadFile("test.txt")
    result := strings.ToUpper(string(data))
    os.WriteFile("test.txt", []byte(result), 0644)
    
    data2, _ := os.ReadFile("test.txt")
    fmt.Print(string(data2))
}
```
**Question:** What gets printed?

**Answer:** `HELLO WORLD`

**Explanation:**
- Write `"hello world\n"` → read it back → uppercase → write again → read and print
- `strings.ToUpper` uppercases the entire string including the `\n` (whitespace is unaffected)
- `fmt.Print` (not `Println`) — the `\n` in the content handles the newline
- `_` discards the error — fine for examples, bad in production!

**Key Concept:** Ignoring errors with `_` is a code smell — always handle them in real programs!

---

## BLOCK 2 — String Operations

### Problem 6: Fields vs Split
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "  hello   world   go  "
    
    a := strings.Fields(s)
    b := strings.Split(s, " ")
    
    fmt.Println(len(a))
    fmt.Println(len(b))
    fmt.Printf("%q\n", a)
    fmt.Printf("%q\n", b[0])
}
```
**Question:** What does each line print?

**Answer:**
- `3` — Fields splits on ANY whitespace and ignores leading/trailing/multiple spaces
- `23` — Split splits on every single `" "` character, including the empty strings between spaces
- `["hello" "world" "go"]`
- `""` — first element is the empty string before the leading space

**Explanation:**
- `strings.Fields` is smart: trims whitespace, splits on any whitespace run
- `strings.Split(s, " ")` is literal: splits at every single space character
- `"  hello"` split by `" "` gives `["", "", "hello", ...]`
- For tokenizing user input or files, **always prefer `strings.Fields`**

**Key Concept:** `Fields` ignores extra whitespace — `Split` is literal and creates empty strings!

---

### Problem 7: TrimSpace Scope
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "  hello   world  "
    
    fmt.Printf("%q\n", strings.TrimSpace(s))
    fmt.Printf("%q\n", strings.Trim(s, " "))
    fmt.Printf("%q\n", strings.TrimLeft(s, " "))
    fmt.Printf("%q\n", strings.TrimRight(s, " "))
}
```
**Question:** What does each line print?

**Answer:**
- `"hello   world"` — trims leading AND trailing whitespace (tabs, newlines too)
- `"hello   world"` — same result here (trims spaces from both ends)
- `"hello   world  "` — trims leading only
- `"  hello   world"` — trims trailing only

**Explanation:**
- `TrimSpace` handles all Unicode whitespace: spaces, tabs `\t`, newlines `\n`, etc.
- `Trim(s, cutset)` trims any characters in the cutset from both ends
- `TrimLeft` / `TrimRight` trim from one side only
- The internal spaces `"   "` between words are NOT touched by any of these

**Key Concept:** `TrimSpace` trims all whitespace — internal spaces are never touched!

---

### Problem 8: HasPrefix + TrimPrefix Chain
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    words := []string{"(hello)", "world", "(go)", "lang"}
    
    for _, w := range words {
        if strings.HasPrefix(w, "(") {
            inner := strings.TrimPrefix(w, "(")
            inner = strings.TrimSuffix(inner, ")")
            fmt.Println(strings.ToUpper(inner))
        } else {
            fmt.Println(w)
        }
    }
}
```
**Question:** What gets printed?

**Answer:**
```
HELLO
world
GO
lang
```

**Explanation:**
- For words starting with `(`: strip the `(` prefix, strip the `)` suffix, uppercase
- `TrimPrefix` only removes the prefix if it exists — safe to call even if absent
- `TrimSuffix` works the same way for the end
- For words without `(`: print as-is

**Key Concept:** `TrimPrefix`/`TrimSuffix` are safe — they do nothing if the fix isn't present!

---

### Problem 9: strings.Join Behavior
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    words := []string{"the", "quick", "brown", "fox"}
    
    fmt.Println(strings.Join(words, " "))
    fmt.Println(strings.Join(words, ", "))
    fmt.Println(strings.Join(words, ""))
    fmt.Println(strings.Join([]string{}, " "))
    fmt.Println(strings.Join([]string{"solo"}, " "))
}
```
**Question:** What does each line print?

**Answer:**
- `the quick brown fox`
- `the, quick, brown, fox`
- `thequickbrownfox`
- `` (empty string)
- `solo`

**Explanation:**
- `Join` places the separator **between** elements — not before the first or after the last
- Empty slice → empty string (no panic)
- Single element → just the element, separator never used
- This is the inverse of `strings.Fields`/`strings.Split`

**Key Concept:** `Join` places separator BETWEEN elements — empty/single slices are handled safely!

---

### Problem 10: String Building in a Loop
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    words := []string{"hello", "world", "go"}
    
    // Method A
    result1 := ""
    for _, w := range words {
        result1 += w + " "
    }
    
    // Method B
    result2 := strings.Join(words, " ")
    
    fmt.Printf("%q\n", strings.TrimSpace(result1))
    fmt.Printf("%q\n", result2)
    fmt.Println(result1 == result2)
}
```
**Question:** What does each line print?

**Answer:**
- `"hello world go"`
- `"hello world go"`
- `false` — result1 has a trailing space before TrimSpace, but after trim they match... wait!

**Corrected Answer:**
- `"hello world go"` — trimmed
- `"hello world go"` — no trailing space
- `false` — because result1 (before trim) has a trailing space, they are different strings in memory, BUT after `TrimSpace` the printed values look the same. The `==` comparison compares `result1` (with trailing space) to `result2` (without) → `false`

**Explanation:**
- `result1` = `"hello world go "` (trailing space)
- `result2` = `"hello world go"` (no trailing space)
- The `Printf` trims result1 for display but the `==` compares the original
- Method A with `+=` creates a new string each iteration (O(n²)) — `strings.Join` is more efficient

**Key Concept:** `+=` in a loop is inefficient — prefer `strings.Join` or `strings.Builder`!

---

## BLOCK 3 — Number Base Conversion

### Problem 11: strconv.ParseInt Basics
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    n1, err1 := strconv.ParseInt("ff", 16, 64)
    n2, err2 := strconv.ParseInt("1010", 2, 64)
    n3, err3 := strconv.ParseInt("42", 10, 64)
    n4, err4 := strconv.ParseInt("xyz", 16, 64)
    
    fmt.Println(n1, err1)
    fmt.Println(n2, err2)
    fmt.Println(n3, err3)
    fmt.Println(n4, err4)
}
```
**Question:** What does each line print?

**Answer:**
- `255 <nil>` — `ff` in hex = 255
- `10 <nil>` — `1010` in binary = 8+2 = 10
- `42 <nil>`
- `0 strconv.ParseInt: parsing "xyz": invalid syntax` — x, y, z are not valid hex!

**Explanation:**
- `ParseInt(s, base, bitSize)` — base is 2/8/10/16, bitSize is 32 or 64
- Valid hex digits: `0-9`, `a-f`, `A-F` — `x`, `y`, `z` are invalid
- On error, the returned int is `0` (not garbage)
- Always check the error before using the result!

**Key Concept:** `ParseInt` returns `0` on error — check `err != nil` before using the value!

---

### Problem 12: Hex to Decimal Conversion
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func hexToDec(s string) (int64, error) {
    s = strings.TrimPrefix(strings.ToLower(s), "0x")
    return strconv.ParseInt(s, 16, 64)
}

func main() {
    values := []string{"ff", "0xFF", "0XFF", "FF", "10", "GG"}
    
    for _, v := range values {
        result, err := hexToDec(v)
        if err != nil {
            fmt.Printf("%s -> ERROR\n", v)
        } else {
            fmt.Printf("%s -> %d\n", v, result)
        }
    }
}
```
**Question:** What does each iteration print?

**Answer:**
```
ff -> 255
0xFF -> 255
0XFF -> 255
FF -> 255
10 -> 16
GG -> ERROR
```

**Explanation:**
- `strings.ToLower` normalizes `"0XFF"` → `"0xff"`, `"FF"` → `"ff"`
- `TrimPrefix(s, "0x")` removes the `0x` prefix if present — safe if not present
- `"10"` in hex = 16 in decimal (not 10!)
- `"GG"` fails — `G` is not a valid hex digit (only 0-9, a-f)

**Key Concept:** `"10"` in hex = 16 decimal — always know your base when parsing!

---

### Problem 13: Binary to Decimal
```go
package main

import (
    "fmt"
    "strconv"
)

func binToDec(s string) (int64, error) {
    return strconv.ParseInt(s, 2, 64)
}

func main() {
    tests := []string{"101", "1111", "10000000", "2", ""}
    
    for _, t := range tests {
        result, err := binToDec(t)
        if err != nil {
            fmt.Printf("%q -> ERROR\n", t)
        } else {
            fmt.Printf("%q -> %d\n", t, result)
        }
    }
}
```
**Question:** What does each line print?

**Answer:**
```
"101" -> 5
"1111" -> 15
"10000000" -> 128
"2" -> ERROR
"" -> ERROR
```

**Explanation:**
- `101` binary = 4+0+1 = 5
- `1111` binary = 8+4+2+1 = 15
- `10000000` = 2^7 = 128
- `"2"` fails — binary only allows `0` and `1`
- `""` fails — empty string is invalid for ParseInt

**Key Concept:** Binary only uses `0` and `1` — any other digit causes a parse error!

---

### Problem 14: strconv.Atoi vs ParseInt
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    a, err1 := strconv.Atoi("42")
    b, err2 := strconv.Atoi("3.14")
    c, err3 := strconv.Atoi("-17")
    
    fmt.Println(a, err1)
    fmt.Println(b, err2)
    fmt.Println(c, err3)
    
    s := strconv.Itoa(255)
    fmt.Println(s, len(s))
}
```
**Question:** What does each line print?

**Answer:**
- `42 <nil>`
- `0 strconv.Atoi: parsing "3.14": invalid syntax` — Atoi is integers only!
- `-17 <nil>` — Atoi handles negative numbers
- `255 3` — "255" is 3 characters long

**Explanation:**
- `Atoi` = "ASCII to integer" — base 10 only, no floats, no other bases
- `Atoi` is shorthand for `ParseInt(s, 10, 0)` — it returns `int` (not int64)
- `Itoa` = "integer to ASCII" — converts int to base-10 string
- For hex/binary parsing you MUST use `ParseInt` with the correct base

**Key Concept:** `Atoi` is base-10 integers only — for other bases use `ParseInt`!

---

### Problem 15: ParseInt Bit Size Overflow
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Max value for int8 is 127
    a, err1 := strconv.ParseInt("127", 10, 8)
    b, err2 := strconv.ParseInt("128", 10, 8)
    c, err3 := strconv.ParseInt("128", 10, 64)
    
    fmt.Println(a, err1)
    fmt.Println(b, err2)
    fmt.Println(c, err3)
}
```
**Question:** What does each line print?

**Answer:**
- `127 <nil>`
- `127 strconv.ParseInt: parsing "128": value out of range` — 128 overflows int8, returns max value!
- `128 <nil>` — int64 can hold 128 fine

**Explanation:**
- The bitSize (8, 16, 32, 64) limits the range of the parsed value
- bitSize 8 = range [-128, 127] — 128 overflows
- On overflow: returns the clamped max value (127) AND an error
- Always use bitSize 64 unless you have a specific reason for smaller sizes

**Key Concept:** On overflow, `ParseInt` returns the clamped max value AND an error — check both!

---

## BLOCK 4 — Text Parsing & Token Detection

### Problem 16: Iterating Words
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "hello (hex)ff world (bin)1010 end"
    words := strings.Fields(text)
    
    for i, w := range words {
        fmt.Printf("[%d] %q\n", i, w)
    }
}
```
**Question:** What gets printed?

**Answer:**
```
[0] "hello"
[1] "(hex)ff"
[2] "world"
[3] "(bin)1010"
[4] "end"
```

**Explanation:**
- `strings.Fields` splits on whitespace and returns a clean slice
- Tokens like `(hex)ff` are treated as single words — no space inside them
- Indexing with `i` gives us position information for look-ahead/look-behind
- This is the foundation of a simple token parser

**Key Concept:** `strings.Fields` preserves tokens like `(hex)ff` as single units!

---

### Problem 17: Look-Behind Token Parser
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    words := strings.Fields("hello (hex)ff world (bin)1010 end")
    result := []string{}
    
    for _, w := range words {
        switch {
        case strings.HasPrefix(w, "(hex)"):
            hex := strings.TrimPrefix(w, "(hex)")
            n, err := strconv.ParseInt(hex, 16, 64)
            if err == nil {
                result = append(result, strconv.Itoa(int(n)))
            }
        case strings.HasPrefix(w, "(bin)"):
            bin := strings.TrimPrefix(w, "(bin)")
            n, err := strconv.ParseInt(bin, 2, 64)
            if err == nil {
                result = append(result, strconv.Itoa(int(n)))
            }
        default:
            result = append(result, w)
        }
    }
    
    fmt.Println(strings.Join(result, " "))
}
```
**Question:** What gets printed?

**Answer:** `hello 255 world 10 end`

**Explanation:**
- `(hex)ff` → strip prefix → parse `"ff"` as hex → 255 → convert to string
- `(bin)1010` → strip prefix → parse `"1010"` as binary → 10 → convert to string
- Regular words pass through untouched
- `strings.Join` reassembles with spaces
- This is the classic **single-pass token substitution pipeline**

**Key Concept:** Strip prefix → parse → convert back to string — single-pass pipeline pattern!

---

### Problem 18: The Modifier Token Problem
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    words := strings.Fields("my name is (upper) alice and i love (lower) GOLANG")
    result := []string{}
    
    i := 0
    for i < len(words) {
        w := words[i]
        if w == "(upper)" && i+1 < len(words) {
            result = append(result, strings.ToUpper(words[i+1]))
            i += 2
        } else if w == "(lower)" && i+1 < len(words) {
            result = append(result, strings.ToLower(words[i+1]))
            i += 2
        } else {
            result = append(result, w)
            i++
        }
    }
    
    fmt.Println(strings.Join(result, " "))
}
```
**Question:** What gets printed?

**Answer:** `my name is ALICE and i love golang`

**Explanation:**
- When `(upper)` is found, it consumes the NEXT word and uppercases it: skip 2 indices
- When `(lower)` is found, it consumes the NEXT word and lowercases it: skip 2 indices
- Regular words are added as-is and advance by 1
- The `i+1 < len(words)` check prevents out-of-bounds panic if `(upper)` is the last word

**Key Concept:** Look-ahead parsing consumes multiple tokens per iteration — use `i += 2` not `range`!

---

### Problem 19: Look-Behind Token (Previous Word)
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    words := strings.Fields("hello world (upper) go (lower) LANG")
    result := []string{}
    
    for i, w := range words {
        if w == "(upper)" {
            if len(result) > 0 {
                result[len(result)-1] = strings.ToUpper(result[len(result)-1])
            }
        } else if w == "(lower)" {
            if len(result) > 0 {
                result[len(result)-1] = strings.ToLower(result[len(result)-1])
            }
        } else {
            result = append(result, w)
            _ = i
        }
    }
    
    fmt.Println(strings.Join(result, " "))
}
```
**Question:** What gets printed?

**Answer:** `hello WORLD go lang`

**Explanation:**
- `(upper)` modifies the **previous** word already in result
- `world` is added, then `(upper)` uppercases `result[len(result)-1]`
- `go` is added, then `(lower)` lowercases `result[len(result)-1]` — but `go` is already lowercase
- `LANG` is added as-is since it has no modifier after it

**Key Concept:** Look-behind modifies the LAST element of the result slice — `result[len(result)-1]`!

---

### Problem 20: Punctuation Token — (cap)
```go
package main

import (
    "fmt"
    "strings"
)

func capitalize(s string) string {
    if len(s) == 0 {
        return s
    }
    return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func main() {
    words := strings.Fields("this is (cap) a test (cap) of capitalization")
    result := []string{}
    
    for _, w := range words {
        if w == "(cap)" {
            if len(result) > 0 {
                result[len(result)-1] = capitalize(result[len(result)-1])
            }
        } else {
            result = append(result, w)
        }
    }
    
    fmt.Println(strings.Join(result, " "))
}
```
**Question:** What gets printed?

**Answer:** `this is A test Of capitalization`

**Explanation:**
- `(cap)` uppercases the first letter of the previous word
- `"a"` → `capitalize("a")` = `"A"` (upper first char + lower rest = just `"A"`)
- `"test"` → is added normally, then `(cap)` → `"Test"` ... wait — let's re-read

**Corrected trace:**
- `this` → result: `["this"]`
- `is` → result: `["this","is"]`
- `(cap)` → capitalize `"is"` → `"Is"` → result: `["this","Is"]`
- `a` → result: `["this","Is","a"]`
- `test` → result: `["this","Is","a","test"]`
- `(cap)` → capitalize `"test"` → `"Test"` → result: `["this","Is","a","Test"]`
- `of` → result: `["this","Is","a","Test","of"]`
- `capitalization` → result: `["this","Is","a","Test","of","capitalization"]`

**Corrected Answer:** `this Is a Test of capitalization`

**Key Concept:** Trace token parsers step by step — look-behind always targets the LAST added word!

---

## BLOCK 5 — Pipeline Design

### Problem 21: Multi-Pass Pipeline
```go
package main

import (
    "fmt"
    "strings"
)

func passOne(words []string) []string {
    result := []string{}
    for _, w := range words {
        result = append(result, strings.TrimSpace(w))
    }
    return result
}

func passTwo(words []string) []string {
    result := []string{}
    for _, w := range words {
        if w != "" {
            result = append(result, w)
        }
    }
    return result
}

func passThree(words []string) string {
    return strings.Join(words, " ")
}

func main() {
    input := []string{"  hello  ", "", "  world  ", "", "  go  "}
    
    step1 := passOne(input)
    step2 := passTwo(step1)
    step3 := passThree(step2)
    
    fmt.Printf("%q\n", step3)
    fmt.Println(len(step1), len(step2))
}
```
**Question:** What does each line print?

**Answer:**
- `"hello world go"`
- `5 3` — passOne keeps all 5 (trimmed), passTwo removes the 2 empty strings

**Explanation:**
- Pass 1: trim each word → `["hello", "", "world", "", "go"]`
- Pass 2: filter empty strings → `["hello", "world", "go"]`
- Pass 3: join with spaces → `"hello world go"`
- Each pass has a single responsibility — this is pipeline design
- `len(step1) = 5` (same count, just trimmed), `len(step2) = 3` (empties removed)

**Key Concept:** Pipeline passes each have ONE job — trim, then filter, then join!

---

### Problem 22: Pipeline with Conversion Pass
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func convertTokens(words []string) []string {
    result := make([]string, 0, len(words))
    for _, w := range words {
        switch {
        case strings.HasPrefix(w, "(hex)"):
            n, err := strconv.ParseInt(strings.TrimPrefix(w, "(hex)"), 16, 64)
            if err == nil {
                result = append(result, strconv.Itoa(int(n)))
            } else {
                result = append(result, w)
            }
        case strings.HasPrefix(w, "(bin)"):
            n, err := strconv.ParseInt(strings.TrimPrefix(w, "(bin)"), 2, 64)
            if err == nil {
                result = append(result, strconv.Itoa(int(n)))
            } else {
                result = append(result, w)
            }
        default:
            result = append(result, w)
        }
    }
    return result
}

func main() {
    input := "val1 is (hex)1a val2 is (bin)111 invalid is (hex)zz"
    words := strings.Fields(input)
    converted := convertTokens(words)
    fmt.Println(strings.Join(converted, " "))
}
```
**Question:** What gets printed?

**Answer:** `val1 is 26 val2 is 7 invalid is (hex)zz`

**Explanation:**
- `(hex)1a` → `1a` hex = 16+10 = 26 ✓
- `(bin)111` → binary = 4+2+1 = 7 ✓
- `(hex)zz` → `z` is not a valid hex digit → error → keep original `(hex)zz`
- On parse error: preserve the original token (don't silently drop it)

**Key Concept:** On conversion failure, preserve the original token — don't silently discard data!

---

### Problem 23: Chained Modifier Pipeline
```go
package main

import (
    "fmt"
    "strings"
)

func applyModifiers(words []string) []string {
    result := []string{}
    for i := 0; i < len(words); i++ {
        w := words[i]
        switch w {
        case "(upper)":
            if i+1 < len(words) {
                result = append(result, strings.ToUpper(words[i+1]))
                i++
            }
        case "(lower)":
            if i+1 < len(words) {
                result = append(result, strings.ToLower(words[i+1]))
                i++
            }
        default:
            result = append(result, w)
        }
    }
    return result
}

func main() {
    input := "(upper) (lower) hello"
    words := strings.Fields(input)
    out := applyModifiers(words)
    fmt.Println(strings.Join(out, " "))
}
```
**Question:** What gets printed? This is a TRICKY edge case!

**Answer:** `(lower) hello` — wait, let's trace:
- `i=0`: `w = "(upper)"`, `words[1] = "(lower)"` → uppercase `"(lower)"` → `"(LOWER)"`, `i=2`
- `i=2`: `w = "hello"` → append `"hello"`
- Result: `["(LOWER)", "hello"]`

**Corrected Answer:** `(LOWER) hello`

**Explanation:**
- `(upper)` doesn't know the next token is another modifier — it blindly uppercases it
- `(LOWER)` is not recognized as a modifier token anymore (case-sensitive match!)
- This is the edge case of **stacked/chained modifiers** — real parsers need to handle this
- Solution: check if the next word is also a modifier before applying

**Key Concept:** Modifier tokens applied to other modifier tokens create unexpected results — guard against this!

---

### Problem 24: Pipeline Immutability
```go
package main

import (
    "fmt"
    "strings"
)

func process(words []string) []string {
    for i, w := range words {
        words[i] = strings.ToUpper(w)
    }
    return words
}

func main() {
    original := []string{"hello", "world"}
    processed := process(original)
    
    fmt.Println(original)
    fmt.Println(processed)
    fmt.Println(&original[0] == &processed[0])
}
```
**Question:** What does each line print?

**Answer:**
- `[HELLO WORLD]` — original is MODIFIED!
- `[HELLO WORLD]`
- `true` — same underlying array!

**Explanation:**
- Slices are reference types — `process` modifies the ORIGINAL slice in-place
- `processed` and `original` point to the same underlying array
- This violates pipeline immutability — each pass should create a new slice
- Fix: `result := make([]string, len(words))` and write to `result[i]`

**Key Concept:** Modifying a slice parameter mutates the caller's data — always create new slices in pipelines!

---

## BLOCK 6 — Edge Case Handling

### Problem 25: Empty Input Guards
```go
package main

import (
    "fmt"
    "strings"
)

func firstWord(s string) string {
    words := strings.Fields(s)
    if len(words) == 0 {
        return ""
    }
    return words[0]
}

func lastWord(s string) string {
    words := strings.Fields(s)
    return words[len(words)-1] // Bug?
}

func main() {
    fmt.Println(firstWord("hello world"))
    fmt.Println(firstWord(""))
    fmt.Println(firstWord("   "))
    
    fmt.Println(lastWord("hello world"))
    fmt.Println(lastWord(""))
}
```
**Question:** What prints and what panics?

**Answer:**
- `hello`
- `` (empty string)
- `` (empty string) — `Fields` of whitespace-only string is `[]`
- `world`
- ❌ **PANIC** — `lastWord("")` → `Fields("")` = `[]` → `[][−1]` → index out of range!

**Explanation:**
- `strings.Fields("")` and `strings.Fields("   ")` both return an empty slice `[]`
- `firstWord` correctly guards with `len(words) == 0`
- `lastWord` has no guard — accessing `words[len(words)-1]` when len is 0 means `words[-1]` → panic
- Always guard before indexing slices!

**Key Concept:** `strings.Fields` on empty/whitespace input returns an empty slice — always guard!

---

### Problem 26: Args Edge Cases
```go
package main

import (
    "fmt"
    "os"
)

func getArg(index int, fallback string) string {
    if index < len(os.Args) {
        return os.Args[index]
    }
    return fallback
}

func main() {
    input := getArg(1, "input.txt")
    output := getArg(2, "output.txt")
    fmt.Println(input, output)
}
```
**Run as:** `go run main.go myfile.txt`

**Question:** What gets printed?

**Answer:** `myfile.txt output.txt`

**Explanation:**
- `os.Args = ["main", "myfile.txt"]` — only 2 elements
- `getArg(1, "input.txt")` → index 1 exists → returns `"myfile.txt"`
- `getArg(2, "output.txt")` → index 2 does NOT exist → returns fallback `"output.txt"`
- This is a clean way to provide default values for optional CLI arguments

**Key Concept:** Bounds-check before accessing `os.Args` — provide sensible defaults for optional args!

---

### Problem 27: Handling Malformed Tokens
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func safeConvert(token string) string {
    if !strings.HasPrefix(token, "(hex)") {
        return token
    }
    
    hex := strings.TrimPrefix(token, "(hex)")
    
    if hex == "" {
        return token
    }
    
    n, err := strconv.ParseInt(hex, 16, 64)
    if err != nil {
        return token
    }
    
    return strconv.Itoa(int(n))
}

func main() {
    tests := []string{
        "(hex)ff",
        "(hex)",
        "(hex)zz",
        "(hex)0",
        "notahex",
    }
    for _, t := range tests {
        fmt.Printf("%q -> %q\n", t, safeConvert(t))
    }
}
```
**Question:** What does each line print?

**Answer:**
```
"(hex)ff" -> "255"
"(hex)" -> "(hex)"
"(hex)zz" -> "(hex)zz"
"(hex)0" -> "0"
"notahex" -> "notahex"
```

**Explanation:**
- `(hex)ff` → valid → converts to 255
- `(hex)` → empty string after prefix → return original token
- `(hex)zz` → invalid hex → ParseInt fails → return original token
- `(hex)0` → valid → 0 in hex = 0
- `"notahex"` → no prefix match → pass through
- Each guard handles a specific failure mode: no prefix, empty value, invalid chars

**Key Concept:** Layered guards — check prefix, check empty, check parse error — in that order!

---

### Problem 28: The Off-By-One in Slice Delete
```go
package main

import "fmt"

func removeIndex(s []string, i int) []string {
    return append(s[:i], s[i+1:]...)
}

func main() {
    words := []string{"a", "b", "c", "d", "e"}
    
    fmt.Println(removeIndex(words, 0))
    fmt.Println(removeIndex(words, 4))
    fmt.Println(removeIndex(words, 2))
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("PANIC:", r)
        }
    }()
    fmt.Println(removeIndex(words, 5))
}
```
**Question:** What does each line print?

**Answer:**
```
[b c d e]
[a b c d]
[a b d e]  <- wait, but words was modified by earlier calls!
```

**Corrected trace:** `words` is modified in place by the first `removeIndex` call (append modifies underlying array)!

**Actually:** Each call is on the same `words` slice which shares the same backing array. The first `removeIndex(words, 0)` returns `[b c d e]` but ALSO modifies `words`'s backing array. By the third call `words` may be `[b c d d e]` in the backing array.

**Simplified Answer for discussion:**
- `removeIndex(words, 5)` → `s[5:]` = `s[6]` on a 5-element slice → **PANIC**: index out of range
- The recover catches it and prints: `PANIC: runtime error: index out of range [5] with length 5`

**Key Concept:** Removing the last element (`i == len-1`) is valid; removing beyond bounds panics!

---

### Problem 29: Newline Handling in File Content
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Simulating file content read with os.ReadFile
    fileContent := "hello world\nfoo bar\nbaz\n"
    
    // Approach A
    linesA := strings.Split(fileContent, "\n")
    
    // Approach B
    linesB := strings.Fields(fileContent)
    
    fmt.Println(len(linesA))
    fmt.Println(len(linesB))
    fmt.Printf("%q\n", linesA[len(linesA)-1])
}
```
**Question:** What does each line print?

**Answer:**
- `4` — Split creates an EXTRA empty string after the trailing `\n`
- `6` — Fields splits on all whitespace including `\n`, gives individual words
- `""` — the trailing empty string after the final `\n`

**Explanation:**
- `strings.Split("a\nb\n", "\n")` = `["a", "b", ""]` — trailing newline creates trailing empty element
- This is a very common file-reading bug — always filter empty strings after splitting on `\n`
- `strings.Fields` treats `\n` as whitespace — good for word-level processing
- For line-level processing: `Split` then filter, or use `bufio.Scanner`

**Key Concept:** `Split` on `\n` creates a trailing empty string if file ends with newline — always filter it!

---

### Problem 30: Consecutive Modifier Edge Case
```go
package main

import (
    "fmt"
    "strings"
)

func process(words []string) []string {
    result := []string{}
    for _, w := range words {
        switch w {
        case "(upper)":
            if len(result) > 0 {
                result[len(result)-1] = strings.ToUpper(result[len(result)-1])
            }
        case "(lower)":
            if len(result) > 0 {
                result[len(result)-1] = strings.ToLower(result[len(result)-1])
            }
        default:
            result = append(result, w)
        }
    }
    return result
}

func main() {
    // Edge case: modifier at start (no previous word)
    a := process(strings.Fields("(upper) hello world"))
    
    // Edge case: two modifiers in a row
    b := process(strings.Fields("hello (upper) (lower) world"))
    
    // Edge case: modifier at end
    c := process(strings.Fields("hello world (upper)"))
    
    fmt.Println(strings.Join(a, " "))
    fmt.Println(strings.Join(b, " "))
    fmt.Println(strings.Join(c, " "))
}
```
**Question:** What does each line print?

**Answer:**
- `hello world` — `(upper)` at start: `len(result) == 0` → skipped entirely
- `hello WORLD` — wait, trace: `hello` added → `(upper)` → `HELLO` → `(lower)` → `hello` → `world` added

**Corrected trace b:**
- `"hello"` → result: `["hello"]`
- `"(upper)"` → uppercase last → `["HELLO"]`
- `"(lower)"` → lowercase last → `["hello"]`
- `"world"` → result: `["hello", "world"]`

**Answer b:** `hello world` — the two modifiers cancel each other out!

**Answer c:** `hello WORLD` — `(upper)` at end uppercases `"world"`

**Final Answers:**
- `hello world` (modifier at start silently skipped)
- `hello world` (upper then lower cancel out)
- `hello WORLD` (modifier at end works normally)

**Key Concept:** Two consecutive modifiers cancel out — modifier at start with no previous word is silently skipped!

---

## 🎯 BONUS CHALLENGE: Full Pipeline Integration

### Problem 31: Complete Transform Pipeline
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func tokenize(content string) []string {
    return strings.Fields(content)
}

func convertNumbers(words []string) []string {
    result := make([]string, 0, len(words))
    for _, w := range words {
        switch {
        case strings.HasPrefix(w, "(hex)"):
            n, err := strconv.ParseInt(strings.TrimPrefix(w, "(hex)"), 16, 64)
            if err == nil {
                result = append(result, strconv.Itoa(int(n)))
            } else {
                result = append(result, w)
            }
        case strings.HasPrefix(w, "(bin)"):
            n, err := strconv.ParseInt(strings.TrimPrefix(w, "(bin)"), 2, 64)
            if err == nil {
                result = append(result, strconv.Itoa(int(n)))
            } else {
                result = append(result, w)
            }
        default:
            result = append(result, w)
        }
    }
    return result
}

func applyCase(words []string) []string {
    result := []string{}
    for i := 0; i < len(words); i++ {
        w := words[i]
        switch w {
        case "(upper)":
            if i+1 < len(words) {
                result = append(result, strings.ToUpper(words[i+1]))
                i++
            }
        case "(lower)":
            if i+1 < len(words) {
                result = append(result, strings.ToLower(words[i+1]))
                i++
            }
        case "(cap)":
            if i+1 < len(words) {
                next := words[i+1]
                capped := strings.ToUpper(next[:1]) + strings.ToLower(next[1:])
                result = append(result, capped)
                i++
            }
        default:
            result = append(result, w)
        }
    }
    return result
}

func assemble(words []string) string {
    return strings.Join(words, " ")
}

func main() {
    input := "There are (hex)ff zombies and (bin)1010 (upper) heroes"
    
    step1 := tokenize(input)
    step2 := convertNumbers(step1)
    step3 := applyCase(step2)
    output := assemble(step3)
    
    fmt.Println(output)
}
```
**Question:** What gets printed? Trace every step.

**Answer:** `There are 255 zombies and 10 HEROES`

**Step-by-step trace:**
- Step 1 (tokenize): `["There", "are", "(hex)ff", "zombies", "and", "(bin)1010", "(upper)", "heroes"]`
- Step 2 (convertNumbers):
  - `(hex)ff` → 255 → `"255"`
  - `(bin)1010` → 10 → `"10"`
  - Result: `["There", "are", "255", "zombies", "and", "10", "(upper)", "heroes"]`
- Step 3 (applyCase):
  - `(upper)` consumes `heroes` → `"HEROES"`, skip 2
  - Result: `["There", "are", "255", "zombies", "and", "10", "HEROES"]`
- Step 4 (assemble): `"There are 255 zombies and 10 HEROES"`

**Key Concept:** Order of pipeline passes MATTERS — convert numbers first, THEN apply case modifiers!

---

### Problem 32: Pipeline Order Dependency
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    input := "(upper) (hex)ff"
    words := strings.Fields(input)
    
    // Pipeline A: case first, then convert
    // Pipeline B: convert first, then case
    
    // Simulate Pipeline A
    resultA := []string{}
    for i := 0; i < len(words); i++ {
        if words[i] == "(upper)" && i+1 < len(words) {
            resultA = append(resultA, strings.ToUpper(words[i+1]))
            i++
        } else {
            resultA = append(resultA, words[i])
        }
    }
    // Then convert
    for j, w := range resultA {
        if strings.HasPrefix(w, "(HEX)") {
            resultA[j] = "CONVERTED"
        }
    }
    
    fmt.Println("Pipeline A:", strings.Join(resultA, " "))
    
    // Simulate Pipeline B (convert first)
    fmt.Println("Pipeline B: 255")
}
```
**Question:** What does Pipeline A print? Why does Pipeline B give a different result?

**Answer:**
- Pipeline A: `(HEX)FF` — the `(upper)` uppercases the token `(hex)ff` to `(HEX)FF`, but then the conversion step looks for `(HEX)` prefix which doesn't match the conversion logic expecting lowercase `(hex)`
- Pipeline B: `255` — converting numbers first turns `(hex)ff` → `255`, then there's nothing for `(upper)` to uppercase (it was consumed... or it uppercases `255` → `255`)

**Explanation:**
- Pass order is critical — uppercasing `(hex)ff` to `(HEX)FF` breaks the hex detection if it's case-sensitive
- Always run **number conversion passes before case modifier passes**
- This is why pipeline design requires careful thought about dependencies between passes

**Key Concept:** Passes that transform tokens must run BEFORE passes that read those token formats!

---

## 🏆 You Are Now the Pipeline Monster!

### Quick Reference:
| Topic | Key Rule |
|-------|----------|
| `os.Args[0]` | Always the binary — user args start at `[1]` |
| `os.WriteFile` | Always overwrites — never appends |
| `strings.Fields` | Handles multiple/leading/trailing spaces |
| `strings.Split` | Literal split — creates empty strings |
| `TrimPrefix` | Safe — does nothing if prefix absent |
| `ParseInt(s,base,bits)` | Returns `0` AND error on failure |
| `Atoi` | Base-10 only — use ParseInt for hex/bin |
| Look-ahead | Use `i += 2` — must use index loop not range |
| Look-behind | `result[len(result)-1]` — guard with `len > 0` |
| Pipeline passes | Each pass does ONE thing — order matters |
| File + newlines | `Split("\n")` leaves trailing empty string |
| Slice modification | Passing slice to func can mutate caller's data |

**Master these and you'll handle any text processing pipeline in Go! 💪🔥**