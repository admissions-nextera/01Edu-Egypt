# Go-Reloaded Project Guide

> **Rule before you start:** If you are stuck, search first. Every resource link in this guide points to where the answer lives. Do not paste code from AI — you will not understand it under pressure, and you will not learn the skill.

---

## What You Are Building

A command-line tool that takes a text file, applies a set of transformations to its content, and writes the result to a new file. The transformations include case changes, number base conversions, punctuation fixes, article corrections, and quote formatting.

---

## Before You Write a Single Line

Read these two things:

- https://pkg.go.dev/os#ReadFile
- https://gobyexample.com/command-line-arguments

Answer these to yourself before moving on:
- How does your program know which files to use? Where does that information come from?
- What type does `os.ReadFile` return? What do you need to do to work with it as text?

---

## Phase 1 — Project Setup

### Checkpoint 1.1 — Structure and Module

```
go-reloaded/
├── main.go
└── go.mod
```

```bash
go mod init go-reloaded
```

---

### Checkpoint 1.2 — Read the Arguments

Your program must be called like this:

```bash
go run . input.txt output.txt
```

Any other number of arguments should print an error and exit.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != __ {
        fmt.Println("Usage: go run . <input_file> <output_file>")
        return
    }

    inputFile  := os.Args[__]
    outputFile := os.Args[__]

    fmt.Println(inputFile, outputFile) // temporary, remove later
}
```

**Verify before moving on:**
- `go run . sample.txt result.txt` prints both filenames
- `go run . sample.txt` prints the usage message

---

### Checkpoint 1.3 — Read and Write Files

Write two small functions. Do not implement any logic yet — just read and write:

```go
func readFile(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func writeFile(filename string, content string) error {
    return os.WriteFile(filename, []byte(content), ____)
}
```

What permissions value goes in `WriteFile`'s third argument? Search: **"golang os.WriteFile file permissions"**

**Verify before moving on:**
- Create a `sample.txt` with a few words
- Read it, print its content, write it to `result.txt`
- Open `result.txt` and confirm the content matches

**Resources:**
- https://pkg.go.dev/os#ReadFile
- https://pkg.go.dev/os#WriteFile

---

## Phase 2 — Number Conversions

### Checkpoint 2.1 — Concept

Your tool must replace words that precede `(hex)` or `(bin)` with their decimal equivalent.

Example: `"42 (hex)"` → `"66"` because 42 in hexadecimal is 66 in decimal.

The function you need is `strconv.ParseInt`. Before writing anything, read its signature:

https://pkg.go.dev/strconv#ParseInt

Answer these questions:
- What does the second argument (base) control?
- What does the third argument (bitSize) mean?
- What does it return? How do you convert the result back to a string?

---

### Checkpoint 2.2 — Write the Conversion Functions

```go
import "strconv"

func hexToDecimal(s string) string {
    n, err := strconv.ParseInt(s, __, 64)
    if err != nil {
        return s // return original if invalid
    }
    return strconv.Itoa(int(n))
}

func binToDecimal(s string) string {
    n, err := strconv.ParseInt(s, __, 64)
    if err != nil {
        return s
    }
    return strconv.Itoa(int(n))
}
```

Fill in the correct base values.

**Verify before moving on (manual test — just call the functions from main temporarily):**

| Input | Function | Expected |
|---|---|---|
| `"1E"` | hexToDecimal | `"30"` |
| `"FF"` | hexToDecimal | `"255"` |
| `"10"` | binToDecimal | `"2"` |
| `"1010"` | binToDecimal | `"10"` |

---

## Phase 3 — Case Conversion

### Checkpoint 3.1 — Three Functions

```go
import "strings"

func toUpper(word string) string {
    return __________
}

func toLower(word string) string {
    return __________
}

func capitalize(word string) string {
    if word == "" {
        return word
    }
    // Make first letter uppercase, rest lowercase
    return strings.ToUpper(word[:1]) + strings.ToLower(word[__:])
}
```

Docs: https://pkg.go.dev/strings — look up `ToUpper`, `ToLower`.

**Verify before moving on:**

| Input | Function | Expected |
|---|---|---|
| `"hello"` | toUpper | `"HELLO"` |
| `"WORLD"` | toLower | `"world"` |
| `"brooklyn"` | capitalize | `"Brooklyn"` |
| `"BROOKLYN"` | capitalize | `"Brooklyn"` |

---

## Phase 4 — Processing Modifiers

### Checkpoint 4.1 — Concept: How Modifiers Work

Modifiers are special tokens that appear after a word and tell your program to transform the previous word (or several previous words).

Examples:
- `"hello (up)"` → `"HELLO"`
- `"hello world (low, 2)"` → `"hello world"` with both words lowercased... wait, re-read that. Both words lowercased — which words?

The modifier `(up, 2)` means: apply uppercase to the **2 previous words** in the result slice.

Before writing code, map out the logic on paper:
1. You are iterating through a slice of words
2. You encounter `"(up,"`
3. You look at the next token: `"2)"`
4. You go back 2 positions in your result and uppercase each word
5. You do NOT add the modifier tokens to the result

---

### Checkpoint 4.2 — The Processing Function

```go
func processModifiers(words []string) []string {
    result := []string{}

    for i := 0; i < len(words); i++ {
        word := words[i]

        switch word {
        case "(hex)":
            if len(result) > 0 {
                result[len(result)-1] = hexToDecimal(result[len(result)-1])
            }

        case "(bin)":
            // same pattern as hex — fill this in
            __________

        case "(up)":
            if len(result) > 0 {
                result[len(result)-1] = toUpper(result[len(result)-1])
            }

        case "(low)":
            // fill in
            __________

        case "(cap)":
            // fill in
            __________

        case "(up,":
            // next token is the count, e.g. "2)"
            // parse the number, apply toUpper to that many previous words
            __________

        case "(low,":
            __________

        case "(cap,":
            __________

        default:
            result = append(result, word)
        }
    }

    return result
}
```

For the numbered modifiers, you need to:
1. Read `words[i+1]` to get the count (like `"2)"`)
2. Strip the `)` from it
3. Parse it as an integer with `strconv.Atoi`
4. Loop backward through `result` and apply the transformation

Search: **"golang strconv.Atoi"** and **"golang strings.TrimSuffix"**

**Verify before moving on:**

| Input string | Expected output |
|---|---|
| `"hello (up)"` | `"HELLO"` |
| `"world (low)"` | `"world"` |
| `"brooklyn (cap)"` | `"Brooklyn"` |
| `"42 (hex)"` | `"66"` |
| `"10 (bin)"` | `"2"` |
| `"so exciting (up, 2)"` | `"SO EXCITING"` |
| `"the age of wisdom (cap, 4)"` | `"The Age Of Wisdom"` |

---

## Phase 5 — Punctuation Spacing

### Checkpoint 5.1 — What Needs Fixing

Punctuation marks (`. , ! ? : ;`) should have no space before them and exactly one space after.

Also handle groups: `...` and `!?` and `!!` behave as a single unit.

Input:  `"Hello ,world !How are you ?"`
Output: `"Hello, world! How are you?"`

---

### Checkpoint 5.2 — Write the Function

Think through the algorithm before writing code. For each word:
- If the word IS a punctuation token → attach it to the previous word (no space between them)
- If the word STARTS with punctuation → attach the punctuation part to the previous word, keep the rest as a new word

Write the function signature and logic yourself. This is intentionally left as a blank. Use this as your guide:

```go
func fixPunctuation(words []string) []string {
    punctuation := []string{".", ",", "!", "?", ":", ";", "...", "!?", "!!"}
    result := []string{}

    for i := 0; i < len(words); i++ {
        // Check if current word is in punctuation list
        // If yes: attach to last word in result, do not add as new element
        // If no: add normally
        __________
    }

    return result
}
```

Helper: write a small `isPunctuation(s string) bool` function that checks if a string is in the punctuation list.

**Verify before moving on:**

| Input | Expected |
|---|---|
| `"there , and"` | `"there, and"` |
| `"BAMM !!"` | `"BAMM!!"` |
| `"thinking ..."` | `"thinking..."` |
| `"Hello , world ! How are you ?"` | `"Hello, world! How are you?"` |

---

## Phase 6 — Single Quote Handling

### Checkpoint 6.1 — What Needs Fixing

Single quotes wrap words. They should attach directly to the word(s) inside, with no spaces.

Input:  `"' awesome '"` → `"'awesome'"`
Input:  `"' I am happy '"` → `"'I am happy'"`

---

### Checkpoint 6.2 — Write the Function

The algorithm:
1. Find the first standalone `'`
2. Find the matching closing `'`
3. Everything between them stays as words
4. Attach the opening `'` to the first word after it (no space)
5. Attach the closing `'` to the last word before it (no space)

```go
func fixQuotes(words []string) []string {
    // Hint: iterate with a flag that tracks whether you are
    // "inside" a quote pair or not
    __________
}
```

**Verify before moving on:**

| Input | Expected |
|---|---|
| `"say ' hello ' to"` | `"say 'hello' to"` |
| `"' I am fine '"` | `"'I am fine'"` |

---

## Phase 7 — A/An Correction

### Checkpoint 7.1 — The Rule

The article "a" should become "an" when the next word starts with a vowel sound. For this project, treat words starting with `a e i o u h` (upper or lower case) as vowel-starting.

```go
func fixArticles(words []string) []string {
    vowels := "aeiouAEIOU"

    result := make([]string, len(words))
    copy(result, words)

    for i := 0; i < len(result)-1; i++ {
        if result[i] == "a" || result[i] == "A" {
            nextWord := result[i+1]
            if __________ {
                // preserve the original case of "a" or "A"
                if result[i] == "A" {
                    result[i] = "An"
                } else {
                    result[i] = "an"
                }
            }
        }
    }

    return result
}
```

Fill in the blank: how do you check if `nextWord` starts with a vowel?

Search: **"golang strings.ContainsRune"** or **"golang check first character of string"**

**Verify before moving on:**

| Input | Expected |
|---|---|
| `"a apple"` | `"an apple"` |
| `"A apple"` | `"An apple"` |
| `"a hour"` | `"an hour"` |
| `"a banana"` | `"a banana"` |
| `"a day"` | `"a day"` |

---

## Phase 8 — Putting It All Together

### Checkpoint 8.1 — The Pipeline

```go
func processText(text string) string {
    words := strings.Fields(text)

    words = processModifiers(words)
    words = fixPunctuation(words)
    words = fixQuotes(words)
    words = fixArticles(words)

    return strings.Join(words, " ")
}
```

Why `strings.Fields` instead of `strings.Split(text, " ")`? Search: **"golang strings.Fields vs strings.Split"** and understand the difference before moving on.

---

### Checkpoint 8.2 — Final main.go

```go
func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . <input_file> <output_file>")
        return
    }

    content, err := readFile(os.Args[1])
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    result := processText(content)

    err = writeFile(os.Args[2], result)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }
}
```

---

## Phase 9 — Full System Tests

Create these input files and run your program on each one. Compare the output to the expected result manually.

**Test 1 — Basic modifiers**
```
input:    hello (up) world (low) go (cap)
expected: HELLO world Go
```

**Test 2 — Number conversions**
```
input:    1E (hex) and 10 (bin) are numbers
expected: 30 and 2 are numbers
```

**Test 3 — Numbered modifiers**
```
input:    this is great (up, 2)
expected: this IS GREAT
```

**Test 4 — Punctuation**
```
input:    Hello ,world !How are you ?
expected: Hello, world! How are you?
```

**Test 5 — Quotes**
```
input:    He said : ' hello world '
expected: He said: 'hello world'
```

**Test 6 — Articles**
```
input:    a apple a day keeps a doctor away
expected: an apple a day keeps a doctor away
```

**Test 7 — The provided example (complex)**
```
input:    it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) ...
```
Work out the expected output yourself before running it. Then run it to verify.

---

## Phase 10 — Edge Cases

Before submission, manually test these:

- What happens if the input file is empty?
- What happens if a modifier appears at the start of the file with no previous word?
- What happens if `(up, 10)` is used but there are fewer than 10 previous words?
- What happens with multiple quote pairs in the same line?

Your code should not crash on any of these. Add guard conditions where needed.

---

## Debugging Reference

**Index out of range panic**
Cause: You accessed `words[i-1]` or `result[len(result)-1]` without checking if the index exists.
Fix: Always guard with `if len(result) > 0` or `if i > 0` before accessing previous elements.

**Modifiers appear in output**
Cause: Your switch/if block fell through to `result = append(result, word)`.
Fix: Make sure every modifier case does NOT call append, or return before reaching it.

**Punctuation still has spaces**
Cause: You are joining words with `" "` but the punctuation was not attached to the previous word.
Fix: The fixPunctuation function should merge punctuation into the previous element, not add it as a new element.

**Articles not corrected**
Cause: The comparison is case-sensitive and the word in the text is `"A"` not `"a"`.
Fix: Check for both `"a"` and `"A"`.

---

## Key Packages Used

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Read/write files, command-line args | https://pkg.go.dev/os |
| `strings` | Fields, Join, ToUpper, ToLower, TrimSpace | https://pkg.go.dev/strings |
| `strconv` | ParseInt (hex/bin), Atoi, Itoa | https://pkg.go.dev/strconv |
| `fmt` | Formatted output and error messages | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] Program accepts exactly two command-line arguments
- [ ] Prints usage message for wrong argument count
- [ ] Reads input file correctly and handles read errors
- [ ] `(hex)` converts previous word from hex to decimal
- [ ] `(bin)` converts previous word from binary to decimal
- [ ] `(up)` uppercases previous word
- [ ] `(low)` lowercases previous word
- [ ] `(cap)` capitalizes previous word
- [ ] `(up, N)` uppercases previous N words
- [ ] `(low, N)` lowercases previous N words
- [ ] `(cap, N)` capitalizes previous N words
- [ ] Punctuation spacing is fixed correctly
- [ ] Punctuation groups (`...`, `!?`, `!!`) are handled
- [ ] Single quotes attached with no spaces
- [ ] `a` corrected to `an` before vowel-starting words
- [ ] Writes output file correctly and handles write errors
- [ ] No crashes on edge cases (empty file, modifier with no previous word, etc.)
- [ ] All provided test cases pass