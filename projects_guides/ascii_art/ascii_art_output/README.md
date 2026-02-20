# ASCII-Art-Output Project Guide

> **Before you start:** This project builds on ascii-art and ascii-art-fs. Both must be working before you start here.

---

## Objectives

By completing this project you will learn:

1. **Output Redirection** — Writing program output to a file instead of the terminal
2. **Flag Parsing** — Detecting and extracting a value from a `--flag=value` formatted argument
3. **Conditional Output** — Making your program write to either stdout or a file depending on arguments
4. **File Writing** — Creating and writing text files with Go's `os` package

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art-FS (Completed)
- Multi-argument parsing with a default banner
- `render` function that produces a string or prints directly

### 2. File Writing
- `os.WriteFile` or `os.Create` + writing — which is better here and why?
- File permission values

### 3. Flag Format
- `strings.HasPrefix` — detect `--output=`
- `strings.TrimPrefix` — extract the filename after `=`

**Read before starting:**
- https://pkg.go.dev/os#WriteFile
- Search: **"golang write string to file"**

---

## Project Structure

```
ascii-art-output/
├── main.go
├── banner.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Detect the `--output` Flag

**Goal:**
```
go run . --output=banner.txt "hello" standard   → creates banner.txt, prints nothing to terminal
go run . "hello" standard                        → prints to terminal as before
go run . --output=banner.txt "hello"             → creates banner.txt using default banner
```

**Questions to answer before writing anything:**
- After parsing `--output=banner.txt`, what arguments remain for the string and banner?
- How do you shift your argument positions when the flag is present?
- If `--output` is present but in the wrong format (no `=`, no filename), what should happen?

**Code Placeholder:**
```go
// main.go

func main() {
    // 1. Scan os.Args for an argument starting with "--output="
    //    If found:
    //      Extract the filename after "="
    //      Remove it from the args slice so the remaining args are clean
    //    If not found:
    //      outputFile = "" (means print to stdout)

    // 2. Parse the remaining args for [STRING] and optional [BANNER]
    //    Same logic as ascii-art-fs

    // 3. Load banner and render to a string (not print yet)

    // 4. If outputFile is set:
    //      Write the rendered string to the file
    //    Else:
    //      Print to stdout
}
```

**Usage message to print for any invalid format:**
```
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard
```

**Verify:**
```bash
go run . --output=banner.txt "hello" standard
cat -e banner.txt       # should match ascii art for "hello" exactly

go run . --output=banner.txt "Hello There!" shadow
cat -e banner.txt

go run . "hello"        # should still print to terminal, no file created
```

---

## Milestone 2 — Your `render` Function Returns a String

**Goal:** Instead of printing directly in `renderLine`, collect the output and return it as a string so the caller decides where to send it.

**Questions to answer:**
- Does your current `renderLine` print with `fmt.Println`? If yes, how do you change it to build and return a string instead?
- What is `strings.Builder` and why is it better than `+=` for building multiline output?

**Code Placeholder:**
```go
// banner.go

func render(banner []string, input string) string {
    // Split input on "\\n"
    // For each part:
    //   If empty: add a newline to the result
    //   Otherwise: call renderLine and add its output to the result
    // Return the complete result string
}

func renderLine(banner []string, text string) string {
    // Build and return the 8-row output as a single string
    // Do NOT print — return the string to the caller
}
```

**Verify:** After this change, all previous test cases still produce identical output when printed to stdout.

---

## Debugging Checklist

- Does `cat -e banner.txt` show a trailing newline at the end of the file? Check whether the spec output has one or not and match it exactly.
- Do the previous test cases break after changing `render` to return a string? Make sure you are printing the returned string in the non-output case.
- Does the file get created even when rendering fails or the input is invalid? Handle errors before creating the file.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Write output file | https://pkg.go.dev/os |
| `strings` | Detect and strip the flag, build output | https://pkg.go.dev/strings |
| `fmt` | Print to stdout when no file flag | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] `--output=banner.txt` creates the file with correct content
- [ ] File content matches terminal output exactly (verified with `cat -e`)
- [ ] Without `--output`, program prints to terminal as before
- [ ] Works with and without a banner argument
- [ ] Invalid flag format prints usage message
- [ ] All previous ascii-art and ascii-art-fs test cases still pass