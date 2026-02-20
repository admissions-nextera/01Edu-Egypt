# 🎯 ASCII-Art-Output Prerequisites Quiz
## Flag Parsing · File Writing · Return Values vs Print · strings.HasPrefix/TrimPrefix

**Time Limit:** 30 minutes  
**Total Questions:** 20  
**Passing Score:** 16/20 (80%)

> ✅ Pass → You're ready to start ASCII-Art-Output  
> ⚠️ Also Required → ASCII-Art and ASCII-Art-FS must be fully passing

---

## 📋 SECTION 1: FLAG DETECTION & PARSING (6 Questions)

### Q1: Your program receives `--output=banner.txt` as an argument. Which expression correctly checks for this flag prefix?

**A)** `os.Args[1] == "--output"`  
**B)** `strings.HasPrefix(os.Args[1], "--output=")`  
**C)** `strings.Contains(os.Args[1], "output")`  
**D)** `os.Args[1][0:8] == "--output"`  

<details><summary>💡 Answer</summary>

**B) `strings.HasPrefix(os.Args[1], "--output=")`**

Option A would only match `--output` exactly (no `=` or filename). Option C would match `--output=` but also match any argument that happens to contain the word "output". Option D skips the `=` sign. `HasPrefix` with the full `"--output="` is the precise and correct check.

</details>

---

### Q2: After detecting `--output=banner.txt`, how do you extract just `"banner.txt"`?

**A)** `strings.Split(arg, "=")[1]`  
**B)** `strings.TrimPrefix(arg, "--output=")`  
**C)** `arg[9:]`  
**D)** Both A and B are correct; B is more robust  

<details><summary>💡 Answer</summary>

**D) Both A and B are correct; B is more robust**

`strings.Split` works but if the filename itself contains `=` (e.g., `--output=my=file.txt`), `[1]` only gets `"my"`. `TrimPrefix` removes exactly the prefix `"--output="` and returns everything after it, handling all filenames correctly.

```go
filename := strings.TrimPrefix(arg, "--output=")
```

</details>

---

### Q3: What is the output?
```go
arg := "--output=result.txt"
fmt.Println(strings.TrimPrefix(arg, "--output="))
fmt.Println(strings.TrimPrefix(arg, "--color="))
```

**A)** `result.txt` and `result.txt`  
**B)** `result.txt` and `--output=result.txt`  
**C)** `result.txt` and `""`  
**D)** `result.txt` and error  

<details><summary>💡 Answer</summary>

**B) `result.txt` and `--output=result.txt`**

`TrimPrefix` only removes the prefix if it matches exactly. If the string does NOT start with the given prefix, it returns the original string unchanged. No error is thrown.

</details>

---

### Q4: You need to scan ALL `os.Args` for a `--output=` flag (it might not be at a fixed position). Which approach is correct?

**A)**
```go
if strings.HasPrefix(os.Args[1], "--output=") { ... }
```
**B)**
```go
for i, arg := range os.Args {
    if strings.HasPrefix(arg, "--output=") {
        outputFile = strings.TrimPrefix(arg, "--output=")
        os.Args = append(os.Args[:i], os.Args[i+1:]...)
        break
    }
}
```
**C)**
```go
outputFile = os.Args[len(os.Args)-1]
```
**D)**
```go
outputFile = strings.Replace(os.Args[1], "--output=", "", 1)
```

<details><summary>💡 Answer</summary>

**B)**

This correctly: finds the flag anywhere in args, extracts the value, AND removes the flag from `os.Args` so remaining args can be parsed normally as `[STRING] [BANNER]`. Option A only checks position 1. Option C blindly takes the last arg. Option D doesn't check if the flag is actually present.

</details>

---

### Q5: After removing the `--output=` flag from `os.Args`, the remaining valid combinations are:

**A)** Only `[program, STRING]`  
**B)** `[program, STRING]` and `[program, STRING, BANNER]`  
**C)** Only `[program, STRING, BANNER]`  
**D)** `[program]`, `[program, STRING]`, or `[program, STRING, BANNER]`  

<details><summary>💡 Answer</summary>

**B) `[program, STRING]` and `[program, STRING, BANNER]`**

The flag is separate. After removing it, the same argument rules as ascii-art-fs apply: 1 or 2 user arguments remain. No arguments (`[program]` alone) is still invalid.

</details>

---

### Q6: What should happen if the user passes `--output` (without `=filename`)?

**A)** Default to a file named `"output"`  
**B)** Default to `"output.txt"`  
**C)** Treat it as the STRING to render  
**D)** Print usage message and exit  

<details><summary>💡 Answer</summary>

**D) Print usage message and exit**

`"--output"` without `=filename` is a malformed flag. It should not be silently treated as input text or given a default filename. Print the usage message with an example of the correct format and exit.

</details>

---

## 📋 SECTION 2: RETURNING STRINGS VS PRINTING (5 Questions)

### Q7: Your current `renderLine` function uses `fmt.Println` to print each row. Why must you change this for ASCII-Art-Output?

**A)** `fmt.Println` is too slow  
**B)** You need to collect the rendered string and then decide whether to write it to a file or to stdout — you can't do that if it's already printed  
**C)** `fmt.Println` doesn't support multi-line output  
**D)** The file writing API requires a specific format `fmt.Println` can't produce  

<details><summary>💡 Answer</summary>

**B) You need to collect the rendered string and then decide whether to write it to a file or to stdout**

This is the key architectural change. The render function must **return** a string. The caller decides what to do with it:
```go
output := render(banner, input)
if outputFile != "" {
    os.WriteFile(outputFile, []byte(output), 0644)
} else {
    fmt.Print(output)
}
```

</details>

---

### Q8: What is the output?
```go
func greet(name string) string {
    return "Hello, " + name + "!"
}

func main() {
    greet("World")
}
```

**A)** `Hello, World!`  
**B)** Nothing — the return value is discarded  
**C)** Compile error — unused return value  
**D)** `Hello, World!\n`  

<details><summary>💡 Answer</summary>

**B) Nothing — the return value is discarded**

Go does not error on discarded return values (unlike errors). The function runs but its output goes nowhere. To see it you must `fmt.Println(greet("World"))` or assign it to a variable.

</details>

---

### Q9: You're building a multi-line string in a loop. Which is more efficient and why?

```go
// Option A
result := ""
for i := 0; i < 8; i++ {
    result += "row " + strconv.Itoa(i) + "\n"
}

// Option B
var sb strings.Builder
for i := 0; i < 8; i++ {
    sb.WriteString("row " + strconv.Itoa(i) + "\n")
}
result := sb.String()
```

**A)** Option A — fewer lines of code  
**B)** Option B — avoids repeated string allocation on each `+=`  
**C)** Both are identical in performance  
**D)** Option A — `strings.Builder` has a bug with newlines  

<details><summary>💡 Answer</summary>

**B) Option B — avoids repeated string allocation on each `+=`**

Each `+=` creates a new string and copies all previous content. For 8 iterations this is minor, but for large inputs it becomes O(n²). `strings.Builder` grows its internal buffer exponentially, making it O(n) overall.

</details>

---

### Q10: Your `render` function returns a `string`. The last character of the string is a `\n`. When you write this to a file with `os.WriteFile`, the file will:

**A)** Have a trailing newline at the end  
**B)** Have the `\n` removed automatically  
**C)** Cause an error  
**D)** Have `\n` converted to `\r\n` on Windows  

<details><summary>💡 Answer</summary>

**A) Have a trailing newline at the end**

`os.WriteFile` writes bytes exactly as given. If your string ends with `\n`, the file ends with a newline. Check the spec to see if the expected output has a trailing newline or not — your file must match exactly (use `cat -e` to see `$` at end of each line).

</details>

---

### Q11: After refactoring `renderLine` to return a string, what is the correct signature?

**A)** `func renderLine(banner []string, text string)`  
**B)** `func renderLine(banner []string, text string) string`  
**C)** `func renderLine(banner []string, text string) (string, error)`  
**D)** `func renderLine(text string) string`  

<details><summary>💡 Answer</summary>

**B) `func renderLine(banner []string, text string) string`**

No error return needed — rendering a valid string with a valid banner can't fail. The function takes the loaded banner data and the text to render, and returns the 8-row string. If an invalid character appears, you can either skip it or leave a blank space.

</details>

---

## 📋 SECTION 3: FILE WRITING (5 Questions)

### Q12: Which `os.WriteFile` call is correct for writing the string `output` to `"banner.txt"`?

**A)** `os.WriteFile("banner.txt", output, 0644)`  
**B)** `os.WriteFile("banner.txt", []byte(output), 0644)`  
**C)** `os.WriteFile("banner.txt", output, "rw")`  
**D)** `os.Write("banner.txt", []byte(output))`  

<details><summary>💡 Answer</summary>

**B) `os.WriteFile("banner.txt", []byte(output), 0644)`**

`os.WriteFile` takes `[]byte`, not `string`. The permission `0644` is octal (owner read/write, group/others read). `os.Write` doesn't exist as a top-level function — that's a method on `*os.File`.

</details>

---

### Q13: What does `os.WriteFile` do if the file already exists?

**A)** Returns an error — use `os.Create` instead  
**B)** Appends to the existing file  
**C)** Truncates and overwrites the file  
**D)** Creates a backup of the old file  

<details><summary>💡 Answer</summary>

**C) Truncates and overwrites the file**

`os.WriteFile` creates the file if it doesn't exist, and overwrites it completely if it does. This is the desired behavior for generating output files. If you need to append, use `os.OpenFile` with `os.O_APPEND`.

</details>

---

### Q14: Should you create the output file BEFORE or AFTER rendering the ASCII art?

**A)** Before — create the file first, then fill it  
**B)** After — render first, then write the result  
**C)** Doesn't matter  
**D)** At the same time using goroutines  

<details><summary>💡 Answer</summary>

**B) After — render first, then write the result**

If rendering fails or the input is invalid, you don't want to create an empty file. Always produce the content first, validate it, then write to disk. This prevents partial or corrupt output files.

</details>

---

### Q15: The file permission `0644` means:

**A)** The file is hidden  
**B)** Owner can read and write; group and others can only read  
**C)** Everyone can read and write  
**D)** Only the owner can read  

<details><summary>💡 Answer</summary>

**B) Owner can read and write; group and others can only read**

`0644` in octal: `6` = read+write (4+2), `4` = read only. This is the standard permission for text files. Never use `0777` (everyone has execute permission) for text output files.

</details>

---

### Q16: You write to a file and then run `cat -e banner.txt`. What does the `$` symbol at the end of each line indicate?

**A)** An error in the file  
**B)** The exact position of the line ending — every line must end with `$` (a newline)  
**C)** The file is empty  
**D)** A special character was written  

<details><summary>💡 Answer</summary>

**B) The exact position of the line ending — every line must end with `$` (a newline)**

`cat -e` shows `$` at each `\n`. This is how you verify: no trailing spaces (a space before `$`), correct number of newlines, and no extra blank lines. Always compare your file output against the spec using `cat -e`.

</details>

---

## 📋 SECTION 4: INTEGRATION & CONDITIONAL OUTPUT (4 Questions)

### Q17: When `--output` is NOT provided, your program should:

**A)** Create a default file named `output.txt`  
**B)** Print to stdout exactly as ascii-art-fs did  
**C)** Print to stderr  
**D)** Print a message saying "no output file specified"  

<details><summary>💡 Answer</summary>

**B) Print to stdout exactly as ascii-art-fs did**

Backward compatibility. Without the flag, behavior is identical to before. The flag is purely additive — it redirects output that would otherwise go to the terminal.

</details>

---

### Q18: What is the correct conditional output pattern?

```go
output := render(banner, input)
```

**A)**
```go
if outputFile != "" {
    os.WriteFile(outputFile, []byte(output), 0644)
} else {
    fmt.Print(output)
}
```
**B)**
```go
fmt.Print(output)
os.WriteFile(outputFile, []byte(output), 0644)
```
**C)**
```go
if outputFile == "" {
    os.WriteFile("stdout", []byte(output), 0644)
} else {
    fmt.Print(output)
}
```
**D)**
```go
os.WriteFile(outputFile, []byte(output), 0644)
fmt.Print(output)
```

<details><summary>💡 Answer</summary>

**A)**

Only write to file OR to stdout — not both. Option B always does both. Option D also does both. Option C has the condition backwards.

</details>

---

### Q19: Should you handle the error returned by `os.WriteFile`?

**A)** No — file writes rarely fail  
**B)** Only in production  
**C)** Yes — always check and report the error  
**D)** Only if the disk might be full  

<details><summary>💡 Answer</summary>

**C) Yes — always check and report the error**

```go
if err := os.WriteFile(outputFile, []byte(output), 0644); err != nil {
    fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
    os.Exit(1)
}
```

File writes can fail for many reasons: disk full, no permissions, invalid path. Silent failure is worse than a clear error message.

</details>

---

### Q20: What is the difference between `fmt.Print(output)` and `fmt.Println(output)` when `output` already ends with `\n`?

**A)** No difference  
**B)** `fmt.Println` adds an extra `\n`, resulting in a blank line at the end  
**C)** `fmt.Print` adds a `\n`, `fmt.Println` does not  
**D)** `fmt.Println` is faster  

<details><summary>💡 Answer</summary>

**B) `fmt.Println` adds an extra `\n`, resulting in a blank line at the end**

If `output` ends with `\n` (which it should after the last rendered row), using `fmt.Println` adds another `\n` — producing a spurious blank line at the end. Use `fmt.Print(output)` when your string already contains its own newlines.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 18–20 ✅ | **Ready.** Start ASCII-Art-Output. |
| 16–17 ✅ | **Ready with review.** Fix the questions you missed. |
| 12–15 ⚠️ | **Study first.** Focus on flag parsing and the print-vs-return pattern. |
| Below 12 ❌ | **Not ready.** Review string manipulation, `os.WriteFile`, and function return values. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q6 | `strings.HasPrefix`, `strings.TrimPrefix`, flag parsing patterns |
| Q7–Q11 | Returning strings from functions, `strings.Builder`, `fmt.Print` vs `Println` |
| Q12–Q16 | `os.WriteFile`, file permissions, `cat -e` verification |
| Q17–Q20 | Conditional output, error handling, backward compatibility |