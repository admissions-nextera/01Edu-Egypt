# 🎯 Go-Reloaded Prerequisites Quiz
## File I/O · Strings · Slices · strconv · Loops

**Time Limit:** 40 minutes  
**Total Questions:** 20  
**Passing Score:** 16/20 (80%)

> ✅ Pass → You're ready to start Go-Reloaded  
> ❌ Fail → Review the flagged topics before starting

---

## 📋 SECTION 1: GO BASICS & SLICES (5 Questions)

### Q1: What is the output?
```go
s := []string{"a", "b", "c", "d"}
fmt.Println(s[1:3])
```

**A)** `[a b c]`  
**B)** `[b c]`  
**C)** `[b c d]`  
**D)** `[a b]`  

<details><summary>💡 Answer</summary>

**B) `[b c]`**

Slice syntax `s[low:high]` includes index `low` up to (but **not** including) `high`.

```go
s := []string{"a", "b", "c", "d"}
// indices:       0    1    2    3
fmt.Println(s[1:3]) // index 1 and 2 → [b c]
```

</details>

---

### Q2: What is the output?
```go
s := []string{"go", "is", "fun"}
s = append(s[:1], s[2:]...)
fmt.Println(s)
```

**A)** `[go fun]`  
**B)** `[go is]`  
**C)** `[is fun]`  
**D)** `[go is fun]`  

<details><summary>💡 Answer</summary>

**A) `[go fun]`**

This is the classic "delete element at index 1" pattern:
- `s[:1]` → `["go"]`
- `s[2:]` → `["fun"]`
- `append` joins them → `["go", "fun"]`

The `...` unpacks the second slice as individual arguments to `append`.

</details>

---

### Q3: Which loop correctly iterates over a slice AND gives you both index and value?

**A)**
```go
for s := range words { fmt.Println(s) }
```
**B)**
```go
for i, v := range words { fmt.Println(i, v) }
```
**C)**
```go
for i := range words { fmt.Println(i, words) }
```
**D)**
```go
for _, words := range i { fmt.Println(words) }
```

<details><summary>💡 Answer</summary>

**B)**
```go
for i, v := range words { fmt.Println(i, v) }
```

`range` on a slice returns `(index, value)`. Use `_` to discard either one if not needed.

</details>

---

### Q4: What is the output?
```go
func add(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("b cannot be zero")
    }
    return a + b, nil
}

result, err := add(3, 0)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println(result)
}
```

**A)** `3`  
**B)** `0`  
**C)** `Error: b cannot be zero`  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**C) `Error: b cannot be zero`**

`b == 0` is true, so the function returns `0` and an error. Since `err != nil`, the error branch runs.

</details>

---

### Q5: You have a slice `result` and want to modify the **last element**. Which is correct?

**A)** `result[len(result)] = "X"`  
**B)** `result[len(result)-1] = "X"`  
**C)** `result[-1] = "X"`  
**D)** `result[last] = "X"`  

<details><summary>💡 Answer</summary>

**B) `result[len(result)-1] = "X"`**

Go slices are zero-indexed. A slice of length `n` has valid indices `0` through `n-1`. Go does **not** support negative indexing like Python. Always guard with `len(result) > 0` before doing this to avoid a panic.

</details>

---

## 📋 SECTION 2: FILE I/O & os PACKAGE (4 Questions)

### Q6: What does `os.Args` contain when you run `go run . input.txt output.txt`?

**A)** `["input.txt", "output.txt"]`  
**B)** `["go", "run", ".", "input.txt", "output.txt"]`  
**C)** `[".", "input.txt", "output.txt"]`  
**D)** `[os.Args[0], "input.txt", "output.txt"]` where `os.Args[0]` is the program path  

<details><summary>💡 Answer</summary>

**D) `[os.Args[0], "input.txt", "output.txt"]`**

`os.Args[0]` is always the program name/path. Your arguments start at index **1**:
```go
inputFile  := os.Args[1] // "input.txt"
outputFile := os.Args[2] // "output.txt"
```
So to ensure the user passed exactly 2 arguments: `len(os.Args) != 3`

</details>

---

### Q7: What type does `os.ReadFile("file.txt")` return?

**A)** `string, error`  
**B)** `[]byte, error`  
**C)** `io.Reader, error`  
**D)** `*os.File, error`  

<details><summary>💡 Answer</summary>

**B) `[]byte, error`**

`os.ReadFile` returns raw bytes. To work with it as text you must convert:
```go
data, err := os.ReadFile("file.txt")
text := string(data) // convert []byte → string
```

</details>

---

### Q8: What is the correct call to write a string `content` to `"output.txt"` with standard read/write permissions?

**A)** `os.WriteFile("output.txt", content, 0644)`  
**B)** `os.WriteFile("output.txt", []byte(content), 0644)`  
**C)** `os.WriteFile("output.txt", []byte(content), "rw")`  
**D)** `os.Write("output.txt", content)`  

<details><summary>💡 Answer</summary>

**B) `os.WriteFile("output.txt", []byte(content), 0644)`**

- You must convert `string` → `[]byte` explicitly.
- `0644` is an octal permission value (owner read/write, others read-only).
- `0644` is the standard choice for text files.

</details>

---

### Q9: What happens if you call `os.ReadFile` on a file that does not exist?

**A)** It returns an empty `[]byte` and `nil` error  
**B)** The program panics automatically  
**C)** It returns `nil` and a non-nil `error`  
**D)** It creates the file and returns empty bytes  

<details><summary>💡 Answer</summary>

**C) It returns `nil` and a non-nil `error`**

Go never panics for expected I/O failures. You must always check:
```go
data, err := os.ReadFile("missing.txt")
if err != nil {
    fmt.Println("Error:", err)
    return
}
```

</details>

---

## 📋 SECTION 3: STRINGS PACKAGE (5 Questions)

### Q10: What is the difference between `strings.Fields` and `strings.Split(s, " ")`?

**A)** No difference  
**B)** `strings.Fields` splits on any whitespace and ignores leading/trailing spaces; `strings.Split` splits only on the exact separator  
**C)** `strings.Split` is faster  
**D)** `strings.Fields` returns a map, `strings.Split` returns a slice  

<details><summary>💡 Answer</summary>

**B) `strings.Fields` splits on any whitespace and ignores leading/trailing spaces**

```go
s := "  hello   world  "
strings.Fields(s)       // ["hello", "world"]  ← 2 elements
strings.Split(s, " ")   // ["", "", "hello", "", "", "world", "", ""] ← messy!
```

For text processing, **always prefer `strings.Fields`** unless you need to preserve exact spacing.

</details>

---

### Q11: What is the output?
```go
s := "hello world"
words := strings.Fields(s)
words[0] = strings.ToUpper(words[0])
fmt.Println(strings.Join(words, " "))
```

**A)** `HELLO WORLD`  
**B)** `HELLO world`  
**C)** `hello world`  
**D)** `Hello World`  

<details><summary>💡 Answer</summary>

**B) `HELLO world`**

Only `words[0]` is uppercased. `strings.Join` puts the words back together with a single space separator.

</details>

---

### Q12: Which function correctly checks if a word equals `"(up)"` AND removes it from the result?

**A)**
```go
if word == "(up)" {
    result = append(result, word)
}
```
**B)**
```go
if strings.Contains(word, "up") {
    continue
}
```
**C)**
```go
if word == "(up)" {
    // do NOT append word to result
    result[len(result)-1] = strings.ToUpper(result[len(result)-1])
}
```
**D)**
```go
if word == "(up)" {
    result = result[:len(result)-1]
}
```

<details><summary>💡 Answer</summary>

**C)**

The correct approach: detect the modifier token, apply the transformation to the **last element already in result**, and **don't append the modifier itself**. Option A wrongly appends it. Option B only skips it without transforming. Option D removes the previous word instead of transforming it.

</details>

---

### Q13: What is the output?
```go
s := "(cap, 3)"
s = strings.TrimPrefix(s, "(cap, ")
s = strings.TrimSuffix(s, ")")
fmt.Println(s)
```

**A)** `(cap, 3)`  
**B)** `cap, 3`  
**C)** `3`  
**D)** ` 3`  

<details><summary>💡 Answer</summary>

**C) `3`**

`TrimPrefix` removes the leading `"(cap, "` and `TrimSuffix` removes the trailing `")"`, leaving just `"3"`. This is the pattern for extracting the count from numbered modifiers.

</details>

---

### Q14: You want to capitalize a word — first letter uppercase, rest lowercase. Which is correct?

**A)**
```go
strings.ToUpper(word[:1]) + word[1:]
```
**B)**
```go
strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
```
**C)**
```go
strings.Title(word)
```
**D)**
```go
strings.ToUpper(word)
```

<details><summary>💡 Answer</summary>

**B)**
```go
strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
```

Option A leaves the rest of the word's original casing (e.g. `"hELLO"` → `"HELLO"` instead of `"Hello"`). Option C (`strings.Title`) is deprecated. Option D uppercases the entire word.

</details>

---

## 📋 SECTION 4: strconv & NUMBER CONVERSION (4 Questions)

### Q15: What does `strconv.ParseInt("FF", 16, 64)` return?

**A)** Error — `"FF"` is not a valid number  
**B)** `15`  
**C)** `255`  
**D)** `16`  

<details><summary>💡 Answer</summary>

**C) `255`**

`FF` in hexadecimal = `15×16 + 15 = 255`.  
`ParseInt(s, base, bitSize)`:
- `s` = the string to parse
- `base` = 16 for hex, 2 for binary, 10 for decimal
- `bitSize` = 64 is a safe default

</details>

---

### Q16: What is the output?
```go
n, err := strconv.ParseInt("1010", 2, 64)
if err != nil {
    fmt.Println("error")
} else {
    fmt.Println(n)
}
```

**A)** `1010`  
**B)** `error`  
**C)** `10`  
**D)** `2`  

<details><summary>💡 Answer</summary>

**C) `10`**

`"1010"` in binary = `1×8 + 0×4 + 1×2 + 0×1 = 10` in decimal.

</details>

---

### Q17: You parsed a number with `ParseInt` and got `int64`. You need to put it back into a string. Which is correct?

**A)** `string(n)`  
**B)** `strconv.Itoa(n)`  
**C)** `strconv.Itoa(int(n))`  
**D)** `fmt.Sprint(n)` only  

<details><summary>💡 Answer</summary>

**C) `strconv.Itoa(int(n))`**

`strconv.Itoa` accepts `int`, not `int64`, so you need the explicit cast `int(n)`. Option A (`string(n)`) converts the integer to a Unicode character — not the decimal string. `fmt.Sprint(n)` also works but `strconv.Itoa` is the idiomatic choice for simple integer-to-string conversion.

</details>

---

### Q18: What is the output?
```go
fmt.Println(string(65))
```

**A)** `65`  
**B)** `A`  
**C)** `"65"`  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) `A`**

`string(65)` converts the integer `65` to the Unicode character at code point 65, which is `'A'`. This is a **common trap** — if you want the string `"65"`, use `strconv.Itoa(65)`.

</details>

---

## 📋 SECTION 5: PUTTING IT ALL TOGETHER (2 Questions)

### Q19: You're processing words and encounter the token `"(up,"`. The next token is `"3)"`. What is the correct sequence of steps?

**A)** Uppercase the next 3 words coming in the input  
**B)** Strip `")"` from `"3)"`, parse it as int, uppercase the last 3 elements of `result`, skip the `"3)"` token  
**C)** Add `"(up,"` and `"3)"` to result, fix them later  
**D)** Uppercase the entire `result` slice  

<details><summary>💡 Answer</summary>

**B) Strip `")"` from `"3)"`, parse it as int, uppercase the last 3 elements of `result`, skip the `"3)"` token**

The numbered modifier pattern:
1. Detect `"(up,"` → read next word `"3)"`
2. `strings.TrimSuffix("3)", ")")` → `"3"`
3. `strconv.Atoi("3")` → `3`
4. Advance index `i` past `"3)"` so it's not processed again
5. Apply `strings.ToUpper` to `result[len(result)-3 : len(result)]`

</details>

---

### Q20: What is the output of this complete mini-pipeline?
```go
text := "  hello ,   world  "
words := strings.Fields(text)

result := []string{}
for _, w := range words {
    if w == "," {
        result[len(result)-1] = result[len(result)-1] + w
    } else {
        result = append(result, w)
    }
}

fmt.Println(strings.Join(result, " "))
```

**A)** `hello , world`  
**B)** `hello, world`  
**C)** `hello,world`  
**D)** `hello ,world`  

<details><summary>💡 Answer</summary>

**B) `hello, world`**

`strings.Fields` removes all extra whitespace, giving `["hello", ",", "world"]`.  
The loop attaches `","` directly to `"hello"` → `result = ["hello,", "world"]`.  
`strings.Join` with `" "` → `"hello, world"`.  

This is exactly the punctuation-fixing logic you will implement in Milestone 5.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 18–20 ✅ | **Strong start.** You're well prepared — begin the project. |
| 16–17 ✅ | **Ready.** Review any questions you missed, then start. |
| 12–15 ⚠️ | **Almost there.** Study your weak sections (noted below) before starting. |
| Below 12 ❌ | **Not ready yet.** Revisit Go fundamentals before attempting the project. |

---

## 🔍 Review Map — What to Study If You Missed It

| Question(s) Missed | Topic to Review |
|---|---|
| Q1, Q2, Q5 | Slice indexing and manipulation |
| Q3, Q4 | `range` loops and error handling |
| Q6, Q7, Q8, Q9 | `os` package — Args, ReadFile, WriteFile |
| Q10, Q11 | `strings.Fields` vs `strings.Split`, `strings.Join` |
| Q12, Q13, Q14 | String inspection and transformation |
| Q15, Q16, Q17, Q18 | `strconv.ParseInt`, `strconv.Itoa` — number conversion traps |
| Q19, Q20 | Pipeline logic — putting it all together |