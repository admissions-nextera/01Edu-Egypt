# 🎯 Go-Reloaded Readiness Assessment

## Slices, File I/O, Strings, & Number Conversions

**Time Limit:** 45 minutes

**Total Questions:** 20

**Passing Score:** 16/20 (80%)

---

## 📋 SECTION 1: GO BASICS & SLICES (Questions 1-5)

### Q1: What is the resulting slice when slicing `s := []string{"a", "b", "c", "d", "e"}` with `s[1:4]`?

**A)** `["a", "b", "c"]`

**B)** `["b", "c", "d"]`

**C)** `["b", "c", "d", "e"]`

**D)** `["c", "d", "e"]`

<details><summary>💡 Answer</summary>

**B) `["b", "c", "d"]**`

Slicing `s[low:high]` includes elements from `low` up to, but not including, `high`. Index 1 is `"b"`, 2 is `"c"`, 3 is `"d"`. Index 4 is excluded.

</details>

---

### Q2: How do you safely remove the word at index `i` from a slice of strings named `words` while keeping the original order?

**A)** `words = append(words[:i], words[i+1:]...)`

**B)** `words = words[:i] + words[i+1:]`

**C)** `remove(words, i)`

**D)** `words = append(words[:i], words[i:])`

<details><summary>💡 Answer</summary>

**A) `words = append(words[:i], words[i+1:]...)**`

This appends all elements *after* index `i` to the slice of elements *before* index `i`, effectively skipping the element at `i`. The `...` unpacks the second slice.

</details>

---

### Q3: What happens if `words` has a length of 2, and you try to access `words[2]`?

**A)** It returns an empty string `""`

**B)** It returns `nil`

**C)** The program panics (crashes) with "index out of range"

**D)** It wraps around and returns `words[0]`

<details><summary>💡 Answer</summary>

**C) The program panics (crashes) with "index out of range"**

In Go, arrays and slices are zero-indexed. A slice of length 2 only has indices 0 and 1. Attempting to access index 2 will cause a runtime panic.

</details>

---

### Q4: When iterating with `for i, word := range words`, what do `i` and `word` represent?

**A)** `i` is the current element, `word` is the next element

**B)** `i` is the index (integer), `word` is a copy of the value (string)

**C)** `i` is a pointer, `word` is the value

**D)** `i` is the value, `word` is the index

<details><summary>💡 Answer</summary>

**B) `i` is the index (integer), `word` is a copy of the value (string)**

The `range` keyword returns the index first, and a *copy* of the element second.

</details>

---

### Q5: You want to change the last word in a slice `result`. Which code correctly updates the slice?

**A)** `result[len(result)] = "new"`

**B)** `result[-1] = "new"`

**C)** `result[len(result)-1] = "new"`

**D)** `last(result) = "new"`

<details><summary>💡 Answer</summary>

**C) `result[len(result)-1] = "new"**`

Because slices are zero-indexed, the last element is always at `len(result) - 1`. Go does not support negative indexing like Python (so `-1` is invalid), and `len(result)` would be out of bounds.

</details>

---

## 📋 SECTION 2: FILE OPERATIONS (`os`) (Questions 6-10)

### Q6: If the user runs `go run . input.txt output.txt`, what is the value of `os.Args[1]`?

**A)** `go`

**B)** `.`

**C)** `input.txt`

**D)** `output.txt`

<details><summary>💡 Answer</summary>

**C) `input.txt**`

`os.Args` contains the command-line arguments. `os.Args[0]` is the path to the executable, `os.Args[1]` is the first passed argument (`input.txt`), and `os.Args[2]` is the second (`output.txt`).

</details>

---

### Q7: What are the return types of `os.ReadFile("input.txt")`?

**A)** `(string, error)`

**B)** `([]byte, error)`

**C)** `(*os.File, error)`

**D)** `([]string, error)`

<details><summary>💡 Answer</summary>

**B) `([]byte, error)**`

`os.ReadFile` returns the file's contents as a byte slice `[]byte`. You must cast it to a string using `string(data)` to work with it as text.

</details>

---

### Q8: What is the standard way to check if an error occurred during a file read in Go?

**A)** `try { os.ReadFile(...) } catch (e) { ... }`

**B)** `if err == false { ... }`

**C)** `if err != nil { ... }`

**D)** `if os.HasError() { ... }`

<details><summary>💡 Answer</summary>

**C) `if err != nil { ... }**`

Go does not use try/catch blocks. Errors are returned as standard values, and you must explicitly check if the error variable is not `nil`.

</details>

---

### Q9: When using `os.WriteFile`, what does the third argument represent? (e.g., `os.WriteFile("out.txt", data, 0644)`)

**A)** The file size limit

**B)** The file mode / permissions (e.g., read/write access)

**C)** The encoding format (e.g., UTF-8)

**D)** The line-ending format

<details><summary>💡 Answer</summary>

**B) The file mode / permissions**

`0644` is a standard Unix permission octal indicating the owner can read/write, and others can read.

</details>

---

### Q10: How do you gracefully exit a program and print a message if the user doesn't provide exactly 2 arguments (input and output files)?

**A)** `fmt.Println("Error"); os.Exit(1)`

**B)** `panic("Error")`

**C)** `return "Error"`

**D)** `os.Stop("Error")`

<details><summary>💡 Answer</summary>

**A) `fmt.Println("Error"); os.Exit(1)**` (or returning from `main`)

Checking `len(os.Args) != 3` and then printing usage and exiting is the standard way to handle missing arguments.

</details>

---

## 📋 SECTION 3: STRING MANIPULATION (`strings`) (Questions 11-15)

### Q11: What is the crucial difference between `strings.Fields(text)` and `strings.Split(text, " ")`?

**A)** `Fields` splits on commas, `Split` splits on spaces.

**B)** `Fields` treats multiple consecutive spaces/newlines as a single separator, while `Split` strictly splits on every single space, which can result in empty strings `""` in the slice.

**C)** `Fields` modifies the original string, `Split` returns a new one.

**D)** There is no difference.

<details><summary>💡 Answer</summary>

**B) `Fields` treats multiple consecutive spaces/newlines as a single separator...**

If `text := "hello   world"`, `strings.Fields` returns `["hello", "world"]`. `strings.Split` returns `["hello", "", "", "world"]`. `Fields` is much safer for parsing words.

</details>

---

### Q12: Which function combines a slice of strings `[]string{"Go", "is", "fun"}` into the single string `"Go is fun"`?

**A)** `strings.Concat(words, " ")`

**B)** `strings.Join(words, " ")`

**C)** `strings.Combine(words, " ")`

**D)** `words.ToString(" ")`

<details><summary>💡 Answer</summary>

**B) `strings.Join(words, " ")**`

`strings.Join` takes a slice and a separator string, and merges them together.

</details>

---

### Q13: If `word := "brooklyn"`, how do you capitalize ONLY the first letter to make it `"Brooklyn"`?

**A)** `strings.Capitalize(word)`

**B)** `strings.Title(word)` *(Note: deprecated in modern Go)* **C)** `strings.ToUpper(string(word[0])) + word[1:]`

**D)** `word.Capitalize()`

<details><summary>💡 Answer</summary>

**C) `strings.ToUpper(string(word[0])) + word[1:]**`

Because `strings.Title` is deprecated, the safest manual way is to uppercase the first character (index 0) and concatenate it with the rest of the string (`word[1:]`).

</details>

---

### Q14: Which function checks if a word starts with a specific substring, such as checking if a word starts with `"a"`?

**A)** `strings.HasPrefix(word, "a")`

**B)** `strings.StartsWith(word, "a")`

**C)** `strings.Contains(word, "a")`

**D)** `strings.Index(word, "a") == 1`

<details><summary>💡 Answer</summary>

**A) `strings.HasPrefix(word, "a")**`

`HasPrefix` is the correct Go standard library function to check the beginning of a string.

</details>

---

### Q15: What will `strings.TrimPrefix("(up,", "(")` return?

**A)** `"up)"`

**B)** `"up,"`

**C)** `" (up,"`

**D)** `"(up,"`

<details><summary>💡 Answer</summary>

**B) `"up,"**`

`TrimPrefix` removes the specified string from the beginning if it exists.

</details>

---

## 📋 SECTION 4: NUMBER PARSING (`strconv`) (Questions 16-20)

### Q16: How do you correctly convert the hexadecimal string `"1E"` into a base-10 decimal integer (like 30)?

**A)** `strconv.Atoi("1E")`

**B)** `strconv.ParseInt("1E", 10, 64)`

**C)** `strconv.ParseInt("1E", 16, 64)`

**D)** `strconv.HexToInt("1E")`

<details><summary>💡 Answer</summary>

**C) `strconv.ParseInt("1E", 16, 64)**`

`ParseInt` takes the string, the base of that string (16 for hex), and the bit size.

</details>

---

### Q17: How do you convert the binary string `"1010"` into a base-10 decimal integer?

**A)** `strconv.ParseInt("1010", 2, 64)`

**B)** `strconv.Atoi("1010", 2)`

**C)** `strconv.ParseBin("1010")`

**D)** `strconv.ParseInt("1010", 10, 64)`

<details><summary>💡 Answer</summary>

**A) `strconv.ParseInt("1010", 2, 64)**`

By passing `2` as the base, Go interprets the string as binary.

</details>

---

### Q18: You have a string extracted from a numbered modifier, like `"4"`. What is the easiest way to convert it to a standard `int`?

**A)** `strconv.ParseInt("4", 10, 32)`

**B)** `strconv.Atoi("4")`

**C)** `int("4")`

**D)** `strconv.ToInt("4")`

<details><summary>💡 Answer</summary>

**B) `strconv.Atoi("4")**`

`Atoi` stands for "ASCII to Integer". It is the most convenient function for converting standard base-10 strings into `int` types.

</details>

---

### Q19: You have an integer `66` that you need to put back into your text slice as a string. Which is the correct way?

**A)** `string(66)`

**B)** `strconv.FormatInt(66)`

**C)** `strconv.Itoa(66)`

**D)** `66.String()`

<details><summary>💡 Answer</summary>

**C) `strconv.Itoa(66)**`

`Itoa` stands for "Integer to ASCII". It takes an `int` and returns its string representation `"66"`.

</details>

---

### Q20: *The Trap:* What actually happens if you cast an integer to a string directly, like `string(65)`?

**A)** It returns `"65"`

**B)** It returns `"A"` (the ASCII/Unicode character for 65)

**C)** It won't compile

**D)** It returns `""`

<details><summary>💡 Answer</summary>

**B) It returns `"A"**`

In Go, directly casting an integer to a string interprets the integer as a Unicode code point (rune). This is a very common bug when developers forget to use `strconv.Itoa()`.

</details>