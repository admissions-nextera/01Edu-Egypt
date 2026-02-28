# 🔥 Go Quiz — Runes, ASCII, Strings & Files
---

## BLOCK 1 — Bytes vs Runes

### Problem 1: len() on a String ⭐
```go
s := "café"
fmt.Println(len(s))
fmt.Println(len([]rune(s)))
```
**Question:** What do the two lines print, and why are they different?

**Key Concept:** `len(s)` = bytes. `len([]rune(s))` = characters. They differ for non-ASCII!

---

### Problem 2: Ranging Over a String ⭐⭐
```go
s := "A€B"
for i, r := range s {
    fmt.Printf("i=%d r=%d\n", i, r)
}
```
**Question:** What are the index values printed? Are they 0, 1, 2?

**Key Concept:** `range` over a string gives **byte offsets** as indices — not sequential character numbers!

---

### Problem 3: String Index Returns What? ⭐
```go
s := "Hello"
fmt.Printf("%T %v\n", s[0], s[0])
r := []rune(s)
fmt.Printf("%T %v\n", r[0], r[0])
```
**Question:** What type and value does each line print?

**Key Concept:** `s[i]` = `uint8` byte. `[]rune(s)[i]` = `int32` rune. For non-ASCII they diverge!

---

### Problem 4: Modifying a Rune Slice ⭐⭐
```go
s := "hello"
r := []rune(s)
r[0] = 'H'
fmt.Println(s)
fmt.Println(string(r))
```
**Question:** What does each line print?

**Key Concept:** `[]rune(s)` makes a copy — modifying it never touches the original string!

---

### Problem 5: Counting Characters Safely ⭐⭐
```go
words := []string{"hello", "café", "日本語"}
for _, w := range words {
    fmt.Println(len(w), len([]rune(w)))
}
```
**Question:** What does each line print?

**Key Concept:** ASCII strings: bytes = runes. Non-ASCII strings: bytes > runes. Always use `[]rune` for character counting!

---

### Problem 6: Rune Is Just int32 ⭐
```go
var r rune = 'A'
fmt.Println(r)
fmt.Println(r + 1)
fmt.Printf("%c\n", r+1)
fmt.Printf("%T\n", r)
```
**Question:** What does each line print?

**Key Concept:** Runes are `int32` — character arithmetic is just integer arithmetic!

---

## BLOCK 2 — ASCII Table & Arithmetic

### Problem 7: Printable ASCII Range ⭐
```go
tests := []rune{31, 32, 65, 126, 127}
for _, r := range tests {
    fmt.Println(r, r >= 32 && r <= 126)
}
```
**Question:** What does each line print?

**Key Concept:** Printable ASCII = codes **32 to 126 inclusive**. 127 (DEL) is NOT printable!

---

### Problem 8: Space is 32 ⭐
```go
fmt.Println(' ' == 32)
fmt.Println(int(' '))
fmt.Printf("%c\n", 32)
```
**Question:** What does each line print?

**Key Concept:** Space = ASCII 32. It is the starting anchor for all printable ASCII lookups!

---

### Problem 9: Case Arithmetic ⭐⭐
```go
fmt.Println('a' - 'A')
fmt.Println('z' - 'a')
fmt.Printf("%c\n", 'a'-32)
fmt.Printf("%c\n", 'A'+32)
```
**Question:** What does each line print?

**Key Concept:** `'a' - 'A' = 32` — uppercase↔lowercase is just ±32!

---

### Problem 11: Digit Arithmetic ⭐⭐
```go
fmt.Println('0')
fmt.Println('9' - '0')
fmt.Println(int('5' - '0'))
fmt.Printf("%c\n", '0'+7)
```
**Question:** What does each line print?

**Key Concept:** `c - '0'` converts a digit character to its integer value — no strconv needed!

---

### Problem 12: Validate Before Lookup ⭐⭐⭐
```go
func isValid(r rune) bool {
    return r >= 32 && r <= 126
}
for _, r := range []rune{'A', 'é', '\n', '~', 0} {
    fmt.Printf("%q valid=%v\n", r, isValid(r))
}
```
**Question:** What does each line print?

**Key Concept:** Always validate `r >= 32 && r <= 126` before using a rune as an ASCII table index!

---

## BLOCK 3 — strings.Split / Fields / Join

### Problem 13: Split vs Fields ⭐
```go
s := "  hello   world  "
fmt.Println(len(strings.Split(s, " ")))
fmt.Println(len(strings.Fields(s)))
```
**Question:** What does each line print?

**Key Concept:** `Split(" ")` is literal — use `Fields` when you want to ignore extra whitespace!

---

### Problem 14: Split Trailing Element ⭐⭐
```go
s := "a,b,c,"
parts := strings.Split(s, ",")
fmt.Println(len(parts))
fmt.Printf("%q\n", parts[len(parts)-1])
```
**Question:** What does each line print?

**Key Concept:** A trailing delimiter always creates a trailing `""` element — always check for it!

---

### Problem 15: Split on Multi-Char Delimiter ⭐⭐
```go
s := "hello\\nworld\\ngo"
parts := strings.Split(s, "\\n")
fmt.Println(len(parts))
fmt.Println(parts[1])
```
**Question:** What does each line print? What is `"\\n"` here?

**Key Concept:** `"\\n"` = literal `\n` text (2 chars). `"\n"` = actual newline byte (1 char). Completely different!

---

### Problem 16: Join Behavior ⭐
```go
words := []string{"go", "is", "fun"}
fmt.Println(strings.Join(words, " "))
fmt.Println(strings.Join(words, ""))
fmt.Println(strings.Join([]string{"solo"}, "-"))
fmt.Println(strings.Join([]string{}, "-"))
```
**Question:** What does each line print?

**Key Concept:** `Join` puts separator BETWEEN elements — never before first or after last!

---

### Problem 17: Fields Preserves Nothing ⭐⭐⭐
```go
row := "##  ##"
a := strings.Fields(row)
b := strings.Split(row, " ")
fmt.Println(a)
fmt.Println(b)
fmt.Println(strings.Join(a, " ") == row)
fmt.Println(strings.Join(b, " ") == row)
```
**Question:** What does each line print?

**Key Concept:** `Fields` destroys multiple spaces — never use it when spaces are meaningful data!

---

### Problem 18: Splitting an Empty String ⭐⭐
```go
fmt.Println(len(strings.Split("", ",")))
fmt.Println(len(strings.Fields("")))
fmt.Printf("%q\n", strings.Split("", ",")[0])
```
**Question:** What does each line print?

**Key Concept:** `Split("")` = `[""]` (length 1). `Fields("")` = `[]` (length 0). Different behaviours on empty input!

---

## BLOCK 4 — strings.Builder

### Problem 19: Builder Accumulates ⭐
```go
var sb strings.Builder
sb.WriteString("hello")
fmt.Println(sb.String())
sb.WriteString(" world")
fmt.Println(sb.String())
```
**Question:** What does each line print?

**Key Concept:** `Builder.String()` reads the content WITHOUT clearing — it always accumulates!

---

### Problem 20: Reset Between Uses ⭐⭐
```go
var sb strings.Builder
for _, w := range []string{"a", "b", "c"} {
    sb.WriteString(w)
    fmt.Println(sb.String())
    sb.Reset()
}
```
**Question:** What does each line print?

**Key Concept:** `Reset()` empties the builder — without it, output would be `a`, `ab`, `abc`!

---

### Problem 21: Builder vs += ⭐⭐
```go
var sb strings.Builder
for i := 0; i < 4; i++ {
    sb.WriteString("go")
}
fmt.Println(sb.String())
fmt.Println(sb.Len())
```
**Question:** What does each line print? Why prefer Builder over `+=`?

**Key Concept:** Builder = O(n). String `+=` in a loop = O(n²). Always use Builder for building strings in loops!

---

### Problem 22: WriteRune vs WriteByte vs WriteString ⭐⭐
```go
var sb strings.Builder
sb.WriteRune('H')
sb.WriteByte(105)
sb.WriteString("!")
fmt.Println(sb.String())
fmt.Println(sb.Len())
```
**Question:** What does each line print?

**Key Concept:** `WriteRune` handles Unicode. `WriteByte` writes one byte. `WriteString` writes many. All write to the same buffer!

---

### Problem 23: Len Counts Bytes ⭐⭐⭐
```go
var sb strings.Builder
sb.WriteRune('€')
fmt.Println(sb.Len())
sb.WriteRune('A')
fmt.Println(sb.Len())
```
**Question:** What does each line print?

**Key Concept:** `Builder.Len()` counts **bytes**, not runes — same as `len(s)` on a string!

---

## BLOCK 5 — os.ReadFile / WriteFile

### Problem 24: ReadFile Return Type ⭐
```go
data, err := os.ReadFile("file.txt")
fmt.Printf("%T\n", data)
fmt.Printf("%T\n", err)
fmt.Println(string(data))
```
Assume `file.txt` contains `hello`.

**Question:** What does each line print?

**Key Concept:** `os.ReadFile` returns `([]byte, error)` — always check the error before using the data!

---

### Problem 25: WriteFile Overwrites ⭐⭐
```go
os.WriteFile("out.txt", []byte("first"), 0644)
os.WriteFile("out.txt", []byte("second"), 0644)
data, _ := os.ReadFile("out.txt")
fmt.Println(string(data))
```
**Question:** What does the final line print?

**Key Concept:** `os.WriteFile` truncates and rewrites from scratch — it never appends!

---

### Problem 26: Missing File Error ⭐⭐
```go
data, err := os.ReadFile("missing.txt")
fmt.Println(err != nil)
fmt.Println(len(data))
```
**Question:** What does each line print?

**Key Concept:** On error, `ReadFile` returns `nil` data and a non-nil error — always check `err != nil` first!

---

### Problem 27: Convert []byte to String ⭐
```go
data := []byte{72, 101, 108, 108, 111}
fmt.Println(string(data))
fmt.Println(len(data))
fmt.Printf("%T\n", data[0])
```
**Question:** What does each line print?

**Key Concept:** `string([]byte)` interprets bytes as UTF-8 — the inverse of `[]byte(string)`!

---

### Problem 28: WriteFile Needs []byte ⭐⭐
```go
message := "saved!"
err := os.WriteFile("out.txt", message, 0644)
fmt.Println(err)
```
**Question:** Does this compile? What happens?

**Key Concept:** `WriteFile` signature is `(name string, data []byte, perm fs.FileMode)` — strings must be converted!

---

## 🎯 BONUS — Mixed Traps

### Problem 29: Three Traps in One ⭐⭐⭐
```go
s := "go€"
fmt.Println(len(s))
fmt.Println(s[2])
fmt.Printf("%c\n", s[2])
```
**Question:** What does each line print?

**Key Concept:** `s[i]` gives a raw byte — indexing mid-multibyte char gives garbage output. Use `range` or `[]rune`!

---

### Problem 30: Builder + Split + Join Pipeline ⭐⭐⭐
```go
input := "hello world go"
words := strings.Fields(input)
var sb strings.Builder
for i, w := range words {
    sb.WriteString(strings.ToUpper(w))
    if i < len(words)-1 {
        sb.WriteByte('-')
    }
}
fmt.Println(sb.String())
```
**Question:** What gets printed?

**Key Concept:** `i < len(words)-1` prevents a trailing separator — the classic "join with separator" pattern using Builder!

---

### Problem 31: ReadFile + Rune Validation ⭐⭐⭐
```go
data, _ := os.ReadFile("input.txt")
content := string(data)
for _, r := range content {
    if r < 32 || r > 126 {
        fmt.Printf("invalid: %d\n", r)
    }
}
```
Assume `input.txt` contains `"hi\n"`.

**Question:** What gets printed?

**Key Concept:** Newlines (`\n` = 10), tabs (`\t` = 9) are below 32 — they fail printable ASCII validation!

---

### Problem 32: Split Roundtrip ⭐⭐⭐
```go
original := "one\ntwo\nthree\n"
lines := strings.Split(original, "\n")
rebuilt := strings.Join(lines, "\n")
fmt.Println(original == rebuilt)
fmt.Println(len(lines))
```
**Question:** What does each line print?

**Key Concept:** `Split` + `Join` roundtrips — but `Split` on a trailing delimiter always adds an extra `""`!

---

## 🏆 Quick Reference Card

| Topic | Key Rule |
|-------|----------|
| `len(s)` | Bytes — not characters |
| `len([]rune(s))` | Characters (code points) |
| `s[i]` | `uint8` byte — NOT a rune |
| `range s` index | Byte offset — skips for multi-byte chars |
| `'A'` | `int32` = 65 |
| Printable ASCII | 32–126 inclusive |
| `'a' - 'A'` | 32 — case arithmetic |
| `(c-32)*h` | Maps char → table index |
| `Split(" ")` | Literal — creates empty strings |
| `Fields` | Smart — trims & collapses whitespace |
| `Fields` on font rows | ❌ Never — spaces are meaningful! |
| `Split` trailing `""` | Always appears after trailing delimiter |
| `"\\n"` vs `"\n"` | 2-char text vs 1-char newline byte |
| `Builder.String()` | Reads WITHOUT clearing |
| `Builder.Len()` | Counts bytes, not runes |
| `Builder` vs `+=` | O(n) vs O(n²) |
| `os.ReadFile` | Returns `([]byte, error)` |
| `os.WriteFile` | Always overwrites — never appends |
| `WriteFile` arg | Needs `[]byte`, not `string` |

**Go make those runes dance! 💪🔥**
