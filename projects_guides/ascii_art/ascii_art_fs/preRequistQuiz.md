# 🎯 ASCII-Art-FS Prerequisites Quiz
## Multi-Arg Parsing · Argument Validation · Backward Compatibility · File Selection

**Time Limit:** 30 minutes  
**Total Questions:** 20  
**Passing Score:** 16/20 (80%)

> ✅ Pass → You're ready to start ASCII-Art-FS  
> ⚠️ Also Required → ASCII-Art must be fully passing before you start

---

## 📋 SECTION 1: os.Args & ARGUMENT COUNTING (6 Questions)

### Q1: Your program is run as `go run . "Hello There!" shadow`. What does `len(os.Args)` equal?

**A)** 2  
**B)** 3  
**C)** 4  
**D)** 5  

<details><summary>💡 Answer</summary>

**B) 3**

`os.Args` = `[program_path, "Hello There!", "shadow"]` — 3 elements. The shell treats `"Hello There!"` as a single argument because of the quotes. `len(os.Args) == 3` means: program name + 2 user arguments.

</details>

---

### Q2: What is the output?
```go
// Run as: go run . "hello"
fmt.Println(len(os.Args))
fmt.Println(os.Args[1])
```

**A)** `1` and `hello`  
**B)** `2` and `hello`  
**C)** `2` and `"hello"`  
**D)** `0` and error  

<details><summary>💡 Answer</summary>

**B) `2` and `hello`**

`os.Args[0]` is the program path, `os.Args[1]` is `"hello"` (without quotes — the shell strips them). `len(os.Args) == 2` means exactly one user argument was given.

</details>

---

### Q3: ASCII-Art-FS supports two valid call forms: one argument or two arguments. Which condition correctly guards the main logic?

**A)**
```go
if len(os.Args) == 1 { ... }
```
**B)**
```go
if len(os.Args) < 2 || len(os.Args) > 3 { printUsage(); return }
```
**C)**
```go
if len(os.Args) > 2 { printUsage(); return }
```
**D)**
```go
if os.Args == nil { printUsage(); return }
```

<details><summary>💡 Answer</summary>

**B)**
```go
if len(os.Args) < 2 || len(os.Args) > 3 { printUsage(); return }
```

Valid: `len == 2` (1 user arg) or `len == 3` (2 user args). Everything else is invalid. Option A only catches zero args. Option C wrongly rejects the valid 2-arg form.

</details>

---

### Q4: After validating argument count, you need to set a default banner. Which pattern is correct?

**A)**
```go
banner := os.Args[2]
```
**B)**
```go
banner := "standard"
if len(os.Args) == 3 {
    banner = os.Args[2]
}
```
**C)**
```go
if len(os.Args) == 2 {
    banner := "standard"
}
```
**D)**
```go
banner := os.Args[2] || "standard"
```

<details><summary>💡 Answer</summary>

**B)**

This is the correct "default value with optional override" pattern. Declare with the default first, then override if the argument was provided. Option A panics when only 1 arg is given. Go doesn't have `||` for default values like Python or JS.

</details>

---

### Q5: What happens at runtime if you access `os.Args[2]` when only one user argument was given (i.e., `len(os.Args) == 2`)?

**A)** Returns an empty string  
**B)** Returns `nil`  
**C)** Runtime panic: index out of range  
**D)** Returns `os.Args[1]`  

<details><summary>💡 Answer</summary>

**C) Runtime panic: index out of range**

Go slices don't have safe out-of-bounds access. Always check `len(os.Args)` before accessing any index. This is one of the most common bugs in argument-parsing code.

</details>

---

### Q6: What does the condition `len(os.Args) == 3` mean in the context of this program?

**A)** The user passed 3 arguments  
**B)** The user passed 2 arguments (the string and the banner name)  
**C)** The user passed 1 argument  
**D)** The program name counts as argument 1 and there are 2 user args  

<details><summary>💡 Answer</summary>

**D) The program name counts as argument 1 and there are 2 user args**

`os.Args[0]` = program, `os.Args[1]` = string, `os.Args[2]` = banner name. Always remember: user arguments start at index 1, so `len(os.Args) == 3` means 2 user-provided arguments.

</details>

---

## 📋 SECTION 2: FILE SELECTION & VALIDATION (6 Questions)

### Q7: You need to map banner names to filenames. Which is the most appropriate data structure?

**A)** A slice of strings  
**B)** A `map[string]string`  
**C)** Three separate `if` statements  
**D)** A switch statement  

<details><summary>💡 Answer</summary>

**Both B and D are good answers.** The question rewards knowing why.

`map[string]string` is clean and extensible:
```go
banners := map[string]string{
    "standard":   "standard.txt",
    "shadow":     "shadow.txt",
    "thinkertoy": "thinkertoy.txt",
}
filename, ok := banners[name]
if !ok { printUsage(); return }
```

A `switch` is also idiomatic for a small fixed set:
```go
switch name {
case "standard":   filename = "standard.txt"
case "shadow":     filename = "shadow.txt"
case "thinkertoy": filename = "thinkertoy.txt"
default:           printUsage(); return
}
```

The `map` approach makes it easier to check if a key exists with the two-value assignment `val, ok := m[key]`.

</details>

---

### Q8: What does the two-value map lookup `val, ok := m[key]` give you that `val := m[key]` does not?

**A)** The value converted to a different type  
**B)** `ok` is `true` if the key exists, `false` if it does not — without `ok`, a missing key silently returns the zero value  
**C)** A copy of the map  
**D)** The index of the key  

<details><summary>💡 Answer</summary>

**B) `ok` is `true` if the key exists, `false` if it does not**

```go
banners := map[string]string{"standard": "standard.txt"}
v, ok := banners["unknown"]
// ok == false, v == "" (zero value for string)
```

Without the `ok` check, `"unknown"` silently returns `""` and you'd try to load a file named `""`. Always use the two-value form when existence matters.

</details>

---

### Q9: What should your program do if the user provides `go run . "hello" unknown`?

**A)** Try to load `"unknown.txt"` and crash with a file error  
**B)** Default to `"standard.txt"`  
**C)** Print the usage message and exit cleanly  
**D)** Print an empty output  

<details><summary>💡 Answer</summary>

**C) Print the usage message and exit cleanly**

An unknown banner name is invalid input — treat it the same as a wrong argument count. Print the usage message and return. Never silently fall back to a default when the user explicitly passed an invalid value.

</details>

---

### Q10: `strings.HasSuffix("standard.txt", ".txt")` returns what?

**A)** `"standard"`  
**B)** `true`  
**C)** `false`  
**D)** `".txt"`  

<details><summary>💡 Answer</summary>

**B) `true`**

`strings.HasSuffix(s, suffix)` returns a `bool`. It does not return the string. If you want to strip the suffix: `strings.TrimSuffix("standard.txt", ".txt")` → `"standard"`.

</details>

---

### Q11: You want to check whether a banner name is one of the three valid options WITHOUT using a map. Which is the most readable correct approach?

**A)**
```go
if bannerName == "standard" || bannerName == "shadow" || bannerName == "thinkertoy" {
```
**B)**
```go
if strings.Contains("standard shadow thinkertoy", bannerName) {
```
**C)**
```go
if bannerName != "" {
```
**D)**
```go
validBanners := []string{"standard", "shadow", "thinkertoy"}
if slices.Contains(validBanners, bannerName) {
```

<details><summary>💡 Answer</summary>

**A or D are both valid.** Option B is a bug — `strings.Contains("standard shadow thinkertoy", "tan")` would return `true` for the invalid input `"tan"`.

</details>

---

### Q12: Your program already works for `go run . "hello"` (one argument). After adding two-argument support, which test confirms backward compatibility?

**A)** Run `go run . "hello" standard` and check it still works  
**B)** Run `go run . "hello"` and verify it produces identical output to before  
**C)** Check that `go run .` prints the usage message  
**D)** All of the above  

<details><summary>💡 Answer</summary>

**D) All of the above**

Backward compatibility means: everything that worked before still works. You must test the one-argument form, the two-argument form, and invalid inputs.

</details>

---

## 📋 SECTION 3: CODE ORGANIZATION & ERROR HANDLING (4 Questions)

### Q13: Your `loadBanner` function returns `([]string, error)`. What is the correct way to call it in main?

**A)**
```go
lines := loadBanner("standard.txt")
```
**B)**
```go
lines, _ := loadBanner("standard.txt")
```
**C)**
```go
lines, err := loadBanner("standard.txt")
if err != nil { fmt.Println(err); return }
```
**D)**
```go
lines, err := loadBanner("standard.txt")
fmt.Println(err)
```

<details><summary>💡 Answer</summary>

**C)**

Option A is a compile error — the function returns two values. Option B silently ignores errors (crashes will happen with no explanation). Option D prints errors but continues regardless. Always check and handle errors explicitly.

</details>

---

### Q14: Should your main function call a separate `printUsage()` function, or inline the usage message each time?

**A)** Inline it — simpler code  
**B)** Separate function — usage message is printed in multiple places, so a function avoids duplication  
**C)** Doesn't matter, both are identical  
**D)** Put the usage message in a global variable  

<details><summary>💡 Answer</summary>

**B) Separate function**

In ASCII-Art-FS there are at least 3 places that print usage (wrong arg count, unknown banner name, and potentially file load failure). A `printUsage()` function means you change the message in one place.

</details>

---

### Q15: When loading the banner file fails (file not found), should your program:

**A)** Panic  
**B)** Print the error to stdout with `fmt.Println` and continue  
**C)** Print a meaningful error to stderr and exit with a non-zero code  
**D)** Return an empty banner and continue silently  

<details><summary>💡 Answer</summary>

**C) Print a meaningful error to stderr and exit with a non-zero code**

```go
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: could not load banner '%s': %v\n", filename, err)
    os.Exit(1)
}
```

Errors go to `stderr` so they don't mix with program output. `os.Exit(1)` signals failure to the shell.

</details>

---

### Q16: After this refactor, what is the correct minimum structure of `main`?

**A)** Validate args → load banner → render → print  
**B)** Print → load banner → validate args → render  
**C)** Load banner → validate args → render → print  
**D)** Render → validate args → load banner → print  

<details><summary>💡 Answer</summary>

**A) Validate args → load banner → render → print**

You must validate arguments first so you know which file to load. Load the file before rendering so you have data to render. Print last. The order matters — each step depends on the previous.

</details>

---

## 📋 SECTION 4: INTEGRATION THINKING (4 Questions)

### Q17: The spec says `go run . "hello" unknown` should print a usage message. The spec also says `go run . "hello"` should render with the standard banner. Which test distinguishes a correct implementation from a buggy one?

**A)** Running `go run . "hello" shadow` and checking it renders  
**B)** Running `go run . "hello" STANDARD` (uppercase) and checking whether it's accepted or rejected  
**C)** Running `go run .` and checking it prints usage  
**D)** Running `go run . "hello" standard shadow` and checking it prints usage  

<details><summary>💡 Answer</summary>

**B) Running `go run . "hello" STANDARD` (uppercase) and checking whether it's rejected**

The spec says valid names are `"standard"`, `"shadow"`, and `"thinkertoy"` (all lowercase). `"STANDARD"` should print usage. This tests whether your validation is case-sensitive. It's a subtle but real edge case.

</details>

---

### Q18: Your friend says: "I'll just load `standard.txt` always, and if the user passes a banner name I'll ignore it — it passes all the basic tests." What's wrong with this approach?

**A)** Nothing — the basic tests are what matters  
**B)** It will fail any test that passes `shadow` or `thinkertoy`  
**C)** `os.ReadFile` doesn't work that way  
**D)** It will fail because standard.txt requires two arguments  

<details><summary>💡 Answer</summary>

**B) It will fail any test that passes `shadow` or `thinkertoy`**

The project exists specifically to test multi-banner support. A solution that ignores the banner argument will produce wrong output for shadow and thinkertoy inputs — and also won't reject invalid banner names.

</details>

---

### Q19: What should `go run . ""` (empty string, no banner) produce?

**A)** Usage message — empty string is invalid  
**B)** Nothing — same as the original ascii-art behavior  
**C)** A blank line  
**D)** An error about empty input  

<details><summary>💡 Answer</summary>

**B) Nothing — same as the original ascii-art behavior**

Backward compatibility. The original ascii-art produced no output for an empty string. ASCII-Art-FS must match. The empty string is still a valid (1-argument) call — don't add extra validation that breaks previous behavior.

</details>

---

### Q20: You want to run all your original ASCII-Art test cases against the new ASCII-Art-FS code to check backward compatibility. The fastest way is:

**A)** Run each test manually one at a time  
**B)** Write a shell script or test file that runs all original cases and diffs the output  
**C)** Only test the cases you remember  
**D)** Assume it's fine if the new features work  

<details><summary>💡 Answer</summary>

**B) Write a shell script or test file that runs all original cases and diffs the output**

This is the professional approach. Automated regression testing catches issues you didn't think to check manually. Even a simple bash loop comparing outputs is far more reliable than manual spot-checking.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 18–20 ✅ | **Ready.** Start ASCII-Art-FS. |
| 16–17 ✅ | **Ready with review.** Fix the questions you missed. |
| 12–15 ⚠️ | **Study first.** Focus on argument parsing and validation patterns. |
| Below 12 ❌ | **Not ready.** Review `os.Args`, map lookups, and error handling. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q6 | `os.Args` indexing, length checks, default values |
| Q7–Q12 | Map lookups with `ok`, banner validation, backward compatibility |
| Q13–Q16 | Error handling patterns, function structure, exit codes |
| Q17–Q20 | Edge cases, regression testing, empty input handling |