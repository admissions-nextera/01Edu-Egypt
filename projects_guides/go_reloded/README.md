# Go-Reloaded Project Guide

> **Before you start:** Read the project spec carefully. Run the provided test cases manually so you know exactly what the output should look like before writing any code.

---

## Objectives

By completing this project you will learn:

1. **File I/O** — Reading from and writing to files using Go's `os` package
2. **String Manipulation** — Splitting, joining, trimming, and transforming text
3. **Number Base Conversion** — Converting hexadecimal and binary strings to decimal
4. **Text Parsing** — Iterating over words and identifying special tokens
5. **Pipeline Design** — Chaining multiple transformation passes over the same data
6. **Edge Case Handling** — Writing code that does not crash on unexpected input

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Go Basics
- Functions, return values, error handling
- Slices — appending, indexing, slicing (`s[:i]`, `s[i+1:]`)
- `for` loops with index and `range`

### 2. File Operations
- `os.ReadFile` — read an entire file into memory
- `os.WriteFile` — write a string to a file
- `os.Args` — read command-line arguments

### 3. String Operations
- `strings.Fields` vs `strings.Split` — know the difference
- `strings.ToUpper`, `strings.ToLower`
- `strings.TrimSpace`, `strings.HasPrefix`, `strings.TrimPrefix`
- `strings.Join`

### 4. Number Parsing
- `strconv.ParseInt` — converting a string to an integer with a given base
- `strconv.Atoi` — converting a string to a decimal integer
- `strconv.Itoa` — converting an integer back to a string

**If any of these are unfamiliar, read about them before writing any code.**

- https://pkg.go.dev/strings
- https://pkg.go.dev/strconv
- https://pkg.go.dev/os

---

## Project Structure

```
go-reloaded/
├── main.go
└── go.mod
```

---

## Milestone 1 — Read Input, Write Output

**Goal:**
```
go run . input.txt output.txt
```
Reads `input.txt`, does nothing to it yet, and writes the exact same content to `output.txt`.

**Questions to answer before writing anything:**
- How many command-line arguments does your program expect? What happens if the user provides the wrong number?
- What type does `os.ReadFile` return, and what do you need to do before you can work with it as text?
- What does the third argument of `os.WriteFile` control?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Check that exactly 2 arguments were provided (input and output file)
    //    If not, print usage message and return

    // 2. Read the input file into a string

    // 3. (Transformations will go here in later milestones)

    // 4. Write the result string to the output file
}

func readFile(filename string) (string, error) {
    // Read the file
    // Convert the result to a string and return it
}

func writeFile(filename string, content string) error {
    // Write the content to the file with appropriate permissions
}
```

**Verify:**
- Create a `sample.txt` with a few words
- Run `go run . sample.txt result.txt`
- Open `result.txt` — it must be identical to `sample.txt`

---

## Milestone 2 — Number Conversions: `(hex)` and `(bin)`

**Goal:**
```
input:    42 (hex) and 10 (bin)
output:   66 and 2
```
The word before `(hex)` is treated as a hexadecimal number and replaced with its decimal value. Same for `(bin)`.

**Questions to answer:**
- What function converts a string like `"42"` from base 16 to a decimal integer?
- What base value do you pass for hexadecimal? For binary?
- What should your function return if the input is not a valid number?

**Code Placeholder:**
```go
func hexToDecimal(s string) string {
    // Parse s as a base-16 integer
    // If parsing fails, return s unchanged
    // Convert the result back to a decimal string and return it
}

func binToDecimal(s string) string {
    // Parse s as a base-2 integer
    // If parsing fails, return s unchanged
    // Convert the result back to a decimal string and return it
}
```

**Resources:**
- Search: **"golang strconv ParseInt base"**
- https://pkg.go.dev/strconv#ParseInt

**Verify manually before moving on:**

| Input | Function | Expected |
|---|---|---|
| `"1E"` | hexToDecimal | `"30"` |
| `"FF"` | hexToDecimal | `"255"` |
| `"10"` | binToDecimal | `"2"` |
| `"1010"` | binToDecimal | `"10"` |

---

## Milestone 3 — Case Modifiers: `(up)`, `(low)`, `(cap)`

**Goal:**
```
input:    hello (up) WORLD (low) brooklyn (cap)
output:   HELLO world Brooklyn
```

**Questions to answer:**
- After splitting the text into words, how do you detect that a word is `(up)`, `(low)`, or `(cap)`?
- When you encounter a modifier, which word do you transform — and where is it relative to the modifier in your result slice?
- What does `capitalize` mean exactly? What should happen to the letters after the first one?

**Code Placeholder:**
```go
func processModifiers(words []string) []string {
    result := []string{}

    for i := 0; i < len(words); i++ {
        // Check if current word is a modifier: (hex), (bin), (up), (low), (cap)
        // If it is a modifier:
        //   - Apply the correct transformation to the last word in result
        //   - Do NOT add the modifier itself to result
        // If it is not a modifier:
        //   - Append it to result as-is
    }

    return result
}
```

**Verify:**

| Input | Expected |
|---|---|
| `"hello (up)"` | `"HELLO"` |
| `"WORLD (low)"` | `"world"` |
| `"brooklyn (cap)"` | `"Brooklyn"` |
| `"42 (hex)"` | `"66"` |
| `"10 (bin)"` | `"2"` |

---

## Milestone 4 — Numbered Modifiers: `(up, 2)`, `(cap, 4)`

**Goal:**
```
input:    this is great (up, 2)
output:   this IS GREAT
```
The number after the comma tells you how many previous words to transform.

**Questions to answer:**
- In your word slice, how does `(up,` appear as a token? What comes right after it?
- How do you extract the number from a token like `"2)"`?
- What should happen if the number is larger than the number of words already in result?

**Code Placeholder:**
```go
// Extend processModifiers to also handle:
// case "(up,":
    // 1. Read the next word from words (it contains the count, e.g. "2)")
    // 2. Strip the ")" and parse it as an integer
    // 3. Skip forward past that token (increment i)
    // 4. Apply toUpper to the last N words in result
    // 5. Guard against N being larger than len(result)

// Same pattern for "(low," and "(cap,"
```

**Resources:**
- Search: **"golang strconv Atoi"**
- Search: **"golang strings TrimSuffix"**

**Verify:**

| Input | Expected |
|---|---|
| `"this is great (up, 2)"` | `"this IS GREAT"` |
| `"the age of wisdom (cap, 4)"` | `"The Age Of Wisdom"` |

---

## Milestone 5 — Punctuation Spacing

**Goal:**
```
input:    Hello ,world !How are you ?
output:   Hello, world! How are you?
```
Punctuation marks (`. , ! ? : ;`) must have no space before them. Groups like `...`, `!!`, `!?` are treated as single tokens.

**Questions to answer:**
- How do you detect that a word is a punctuation token?
- Instead of adding a punctuation token as a new element, what do you do with it?
- What does "attach to the previous word" look like in terms of your result slice?

**Code Placeholder:**
```go
func fixPunctuation(words []string) []string {
    result := []string{}

    for _, word := range words {
        // If word is a punctuation token (. , ! ? : ; ... !! !?):
        //   Attach it to the last element of result (no space between)
        // Otherwise:
        //   Append it normally
    }

    return result
}
```

**Verify:**

| Input | Expected |
|---|---|
| `"there , and"` | `"there, and"` |
| `"BAMM !!"` | `"BAMM!!"` |
| `"thinking ..."` | `"thinking..."` |

---

## Milestone 6 — Single Quote Formatting

**Goal:**
```
input:    ' awesome '         →   'awesome'
input:    ' I am happy '      →   'I am happy'
```
Single quotes must attach directly to the words inside them — no spaces.

**Questions to answer:**
- How do you find the opening `'` and the matching closing `'` in your word slice?
- What does "attach the opening quote to the first word inside" look like in your result slice?
- How do you handle multiple quote pairs in the same line?

**Code Placeholder:**
```go
func fixQuotes(words []string) []string {
    // Iterate through words
    // Track whether you are currently inside an open quote or not
    // When you find a standalone ':
    //   If no quote is open: mark the NEXT word as starting with '
    //   If a quote is open: attach ' to the END of the previous word, close the quote
    // Return the cleaned result
}
```

**Verify:**

| Input | Expected |
|---|---|
| `"say ' hello ' to"` | `"say 'hello' to"` |
| `"' I am fine '"` | `"'I am fine'"` |

---

## Milestone 7 — Article Correction: `a` → `an`

**Goal:**
```
input:    a apple a day keeps a doctor away
output:   an apple a day keeps a doctor away
```
The article `a` (or `A`) must become `an` (or `An`) when the next word starts with a vowel or `h`.

**Questions to answer:**
- How do you check the first character of the next word?
- How do you preserve the original casing — `a` → `an`, `A` → `An`?
- What happens if `a` is the last word in the text with no word after it?

**Code Placeholder:**
```go
func fixArticles(words []string) []string {
    // Copy the words slice so you do not modify the original

    // Iterate through the copy (up to len-1 so you can always look ahead)
    // If current word is "a" or "A":
    //   Check if the next word starts with a vowel or 'h' (upper or lower)
    //   If yes: replace with "an" or "An" to match the original casing

    // Return the updated slice
}
```

**Verify:**

| Input | Expected |
|---|---|
| `"a apple"` | `"an apple"` |
| `"A apple"` | `"An apple"` |
| `"a hour"` | `"an hour"` |
| `"a banana"` | `"a banana"` |

---

## Milestone 8 — Connect the Pipeline

**Goal:**
```
input:    it (cap) was the best of times, it was the worst of times (up) , it was the age of foolishness (cap, 6) ...
```
All transformations run in sequence on the same text.

**Questions to answer:**
- What is the correct order to run the transformations? Does order matter?
- Why use `strings.Fields` to split instead of `strings.Split(text, " ")`?

**Code Placeholder:**
```go
func processText(text string) string {
    // 1. Split text into words using strings.Fields

    // 2. Run processModifiers (hex, bin, up, low, cap, numbered variants)

    // 3. Run fixPunctuation

    // 4. Run fixQuotes

    // 5. Run fixArticles

    // 6. Join words back into a string with single spaces between them
    //    Return the result
}
```

**Resources:**
- Search: **"golang strings Fields vs Split"**

**Verify:** Run every test case from the project spec and confirm each one matches exactly.

---

## Debugging Checklist

Before asking for help, go through this:

- Does your program crash with "index out of range"? Are you checking `len(result) > 0` before accessing the last element?
- Do modifiers still appear in the output? Are you skipping the `append` for modifier tokens?
- Is punctuation still surrounded by spaces? Is it being merged into the previous word, not added as a new element?
- Are articles not changing? Are you checking for both `"a"` and `"A"`?
- Does the output file differ by trailing whitespace? Check `strings.Fields` vs `strings.Split` behavior.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Read/write files, command-line args | https://pkg.go.dev/os |
| `strings` | Fields, Join, ToUpper, ToLower, TrimSpace | https://pkg.go.dev/strings |
| `strconv` | ParseInt, Atoi, Itoa | https://pkg.go.dev/strconv |
| `fmt` | Print usage and error messages | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] Exactly 2 arguments required — usage message otherwise
- [ ] Reads input file and handles read errors
- [ ] `(hex)` converts previous word from hex to decimal
- [ ] `(bin)` converts previous word from binary to decimal
- [ ] `(up)`, `(low)`, `(cap)` transform the previous word
- [ ] `(up, N)`, `(low, N)`, `(cap, N)` transform the previous N words
- [ ] Punctuation attached to previous word with no space before
- [ ] Punctuation groups `...`, `!!`, `!?` handled as single tokens
- [ ] Single quotes attached directly to enclosed words
- [ ] `a` corrected to `an` before vowel-starting and h-starting words
- [ ] Output written to file, write errors handled
- [ ] No crashes on edge cases (modifier with no previous word, empty file)
- [ ] All spec test cases pass
