# ASCII-Art-Justify Prerequisites Quiz
## syscall · TIOCGWINSZ · Alignment Math · Justify Algorithm · Refactoring

**Time Limit:** 55 minutes
**Total Questions:** 30
**Passing Score:** 23/30 (77%)

> ✅ Pass → You are ready to start ASCII-Art-Justify
> ⚠️ Score 23–26 → Study the justify section (Q15–Q23) carefully before starting
> ❌ Score below 23 → Work through the alignment math on paper first

---

## SECTION 1: syscall AND TIOCGWINSZ (7 Questions)

### Q1: This project uses only the standard library to get the terminal width. Which two standard packages are needed?

**A)** `fmt` and `os`
**B)** `syscall` and `unsafe`
**C)** `io` and `bufio`
**D)** `runtime` and `os`

<details><summary>Answer</summary>

**B) `syscall` and `unsafe`**

`syscall` gives you `syscall.Syscall`, `syscall.SYS_IOCTL`, and `syscall.TIOCGWINSZ`.
`unsafe` gives you `unsafe.Pointer`, which is required to pass the address of your `winsize` struct to the kernel.

No third-party packages. No `golang.org/x/term`.

</details>

---

### Q2: What is `TIOCGWINSZ`?

**A)** A Go function that returns terminal dimensions
**B)** An ioctl request code that tells the kernel to fill a `winsize` struct with terminal dimensions
**C)** A file in `/proc` that contains terminal width
**D)** A constant defined in `golang.org/x/term`

<details><summary>Answer</summary>

**B) An ioctl request code that tells the kernel to fill a `winsize` struct with terminal dimensions**

`TIOCGWINSZ` stands for "Terminal I/O Control Get WINdow SiZe". It is a constant (`0x5413` on Linux) passed to the `ioctl` syscall. The kernel reads the terminal dimensions from the driver and writes them into the struct you provide.

</details>

---

### Q3: What does the kernel's `winsize` struct look like? How must you define it in Go?

**A)** Any struct with a `Width int` field
**B)** A struct with fields `Row uint16`, `Col uint16`, `Xpixel uint16`, `Ypixel uint16` in that exact order
**C)** A struct with a single `uint64` field that encodes both width and height
**D)** You don't define it — the `syscall` package exports it

<details><summary>Answer</summary>

**B) A struct with fields `Row uint16`, `Col uint16`, `Xpixel uint16`, `Ypixel uint16` in that exact order**

```go
type winsize struct {
    Row    uint16
    Col    uint16
    Xpixel uint16
    Ypixel uint16
}
```

The kernel fills this struct by writing raw bytes into memory. If your field order is wrong, `Col` will contain `Row`'s value and vice versa. The layout must exactly match the C definition from `<sys/ioctl.h>`.

</details>

---

### Q4: Why do you need `unsafe.Pointer` to pass your `winsize` struct to `syscall.Syscall`?

**A)** `unsafe.Pointer` makes the struct faster to access
**B)** `syscall.Syscall` takes `uintptr` arguments, and `unsafe.Pointer` converts a Go pointer to a raw memory address that can be cast to `uintptr`
**C)** `unsafe.Pointer` is required by Go's garbage collector for all structs
**D)** The kernel cannot accept typed pointers

<details><summary>Answer</summary>

**B) `syscall.Syscall` takes `uintptr` arguments, and `unsafe.Pointer` converts a Go pointer to a raw memory address**

```go
uintptr(unsafe.Pointer(&ws))
```

`syscall.Syscall` passes arguments directly to the kernel as raw integers. Go pointers cannot be cast directly to `uintptr` safely — `unsafe.Pointer` is the correct bridge. The kernel writes the result directly into the memory your struct occupies.

</details>

---

### Q5: What are the three arguments you pass to `syscall.Syscall` for this ioctl call?

**A)** `(SYS_READ, Stdin, &ws)`
**B)** `(SYS_IOCTL, Stdout, TIOCGWINSZ, &ws)` — but that is 4 args, not 3
**C)** `(SYS_IOCTL, uintptr(Stdout), uintptr(TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))`
**D)** `(SYS_IOCTL, uintptr(Stdin), uintptr(TIOCGWINSZ))`

<details><summary>Answer</summary>

**C) `(SYS_IOCTL, uintptr(Stdout), uintptr(TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))`**

`syscall.Syscall` always takes exactly 3 arguments after the syscall number (using `Syscall6` for 6). Here:
- arg1 = file descriptor (stdout = 1)
- arg2 = request code (TIOCGWINSZ)
- arg3 = pointer to the struct for the kernel to fill

```go
ret, _, _ := syscall.Syscall(
    syscall.SYS_IOCTL,
    uintptr(syscall.Stdout),
    uintptr(syscall.TIOCGWINSZ),
    uintptr(unsafe.Pointer(&ws)),
)
```

</details>

---

### Q6: `syscall.Syscall` returns three values: `(r1, r2 uintptr, err syscall.Errno)`. How do you check if the call failed?

**A)** Check if `r1 == 0`
**B)** Check if `r1 == ^uintptr(0)` (all bits set, equivalent to -1 in C) or if `err != 0`
**C)** Check if `r2 != 0`
**D)** Syscalls never fail

<details><summary>Answer</summary>

**B) Check if `r1 == ^uintptr(0)` or `err != 0`**

In C, `ioctl` returns -1 on failure. In Go's `syscall` package, -1 is represented as `^uintptr(0)`. The simpler idiom is:

```go
ret, _, errno := syscall.Syscall(...)
if errno != 0 {
    return 80 // fallback
}
```

If the call fails (e.g., stdout is redirected to a file and has no terminal), fall back to 80.

</details>

---

### Q7: A popular third-party package, `golang.org/x/term`, offers `term.GetSize(fd int) (width, height int, err error)`. How does it differ from what you are building?

**A)** It does something completely different — it has nothing to do with TIOCGWINSZ
**B)** It wraps the same `syscall.Syscall` + `TIOCGWINSZ` logic internally, but provides a clean cross-platform API that also works on macOS and Windows
**C)** It reads terminal width from an environment variable instead of a syscall
**D)** It is faster because it caches the result

<details><summary>Answer</summary>

**B) It wraps the same `syscall.Syscall` + `TIOCGWINSZ` logic internally, but provides a clean cross-platform API**

`golang.org/x/term` is maintained by the Go team and lives at `https://cs.opensource.google/go/x/term`. If you read its source, you will find the exact same `winsize` struct and `SYS_IOCTL` call for Linux. For macOS and Windows it uses different platform-specific APIs. This project forbids it so you learn what is underneath — but in real projects, `golang.org/x/term` is the right tool.

</details>

---

## SECTION 2: ALIGNMENT MATH (8 Questions)

### Q8: Terminal width is 80. Art is 30 columns wide. How many spaces before each row for right alignment?

**A)** 30
**B)** 80
**C)** 50
**D)** 25

<details><summary>Answer</summary>

**C) 50**

`padding = termWidth - artWidth = 80 - 30 = 50`. Prepend 50 spaces so the art ends flush at column 80.

</details>

---

### Q9: Terminal width is 80. Art is 30 wide. Spaces before each row for center alignment?

**A)** 50
**B)** 25
**C)** 15
**D)** 40

<details><summary>Answer</summary>

**B) 25**

`padding = (80 - 30) / 2 = 25`. Integer division. The art is centered with 25 spaces on the left (and 25 implied on the right).

</details>

---

### Q10: Terminal is 80, art is 31 wide. Center padding = `(80 - 31) / 2 = 49 / 2 = 24` (integer division). The art shifts slightly left. Is this correct?

**A)** No — you must always center exactly
**B)** Yes — perfect center is impossible when the gap is odd; truncation to the left is standard behavior
**C)** Add 25 spaces and truncate the right side
**D)** Use floating-point division and round

<details><summary>Answer</summary>

**B) Yes — truncation to the left is standard behavior**

When `(termWidth - artWidth)` is odd, perfect centering is mathematically impossible. Integer division truncates, shifting the art slightly left. This is the standard convention in all terminal tools and text editors.

</details>

---

### Q11: What does `renderedWidth(banner, "Hi")` return if `'H'` is 6 columns wide and `'i'` is 3 columns wide?

**A)** 10 (9 + 1 space gap)
**B)** 9
**C)** 6
**D)** 18

<details><summary>Answer</summary>

**B) 9**

Characters in ASCII art are placed directly adjacent — no automatic gap. `6 + 3 = 9`. The width of any rendered string is the sum of its individual character widths.

</details>

---

### Q12: How do you measure a single character's column width from the banner data?

**A)** Always 5 — all characters are 5 columns wide
**B)** `len(getCharLines(banner, ch)[0])` — the length of any one of its 8 art rows
**C)** Read a width table from the banner file header
**D)** `len(string(ch)) * 9`

<details><summary>Answer</summary>

**B) `len(getCharLines(banner, ch)[0])`**

All 8 rows of a character have the same width. The length of any one row equals the character's column width. Different characters have different widths (e.g. `'i'` is narrower than `'W'`) — always measure, never assume a fixed width.

</details>

---

### Q13: Art is wider than the terminal. What should `alignRight` do?

**A)** Crash with an error
**B)** Set padding to 0 and render the art without truncation
**C)** Truncate the art to fit
**D)** Wrap the art to the next line

<details><summary>Answer</summary>

**B) Set padding to 0 and render the art without truncation**

Always guard: `if padding < 0 { padding = 0 }`. The spec says "only text that fits the terminal size will be tested" — so you will not be tested with art wider than the terminal. But your code must not crash if it happens.

</details>

---

### Q14: For left alignment, how much padding do you add?

**A)** `termWidth - artWidth`
**B)** `(termWidth - artWidth) / 2`
**C)** 0 — no padding
**D)** `artWidth`

<details><summary>Answer</summary>

**C) 0 — no padding**

Left alignment is the default. Art starts at column 1 with no leading spaces. Your existing render function already does this — no changes needed for the left-align case.

</details>

---

### Q15: The correct pipeline order for right-aligned output is:

**A)** `alignRight → renderLine → join → print`
**B)** `renderLine → join → alignRight → print`
**C)** `renderLine → alignRight → join → print`
**D)** `join → renderLine → alignRight → print`

<details><summary>Answer</summary>

**C) `renderLine → alignRight → join → print`**

1. `renderLine` produces the 8 raw art rows
2. `alignRight` pads each row individually with leading spaces
3. `join` combines the 8 padded rows with `\n`
4. `print` outputs the result

Alignment must happen on individual rows before joining. If you join first, you cannot add per-row padding without parsing the string back apart.

</details>

---

## SECTION 3: THE JUSTIFY ALGORITHM (9 Questions)

### Q16: What does justify alignment do?

**A)** Each word is individually centered
**B)** The first word starts at the left edge, the last word ends at the right edge, remaining space distributed evenly between words
**C)** All words align to the right edge
**D)** Words are spaced equally from the center outward

<details><summary>Answer</summary>

**B) First word at left edge, last word at right edge, space distributed between words**

This is "full justification" — the same as text justification in word processors. Both edges are filled. The gaps between words absorb all the extra space.

</details>

---

### Q17: Terminal width 100. Three words with rendered widths 20, 15, 25. Total word width = 60. How much space must be distributed between words?

**A)** 20
**B)** 40
**C)** 25
**D)** 100

<details><summary>Answer</summary>

**B) 40**

`totalSpace = termWidth - totalWordWidth = 100 - 60 = 40`. Three words means 2 gaps between them. The 40 spaces fill those 2 gaps.

</details>

---

### Q18: From Q17 — 40 spaces across 2 gaps. How many spaces per gap?

**A)** 40 in the first gap, 0 in the second
**B)** 20 in each gap
**C)** 19 in first, 21 in second
**D)** Cannot be done evenly

<details><summary>Answer</summary>

**B) 20 in each gap**

`baseSpace = 40 / 2 = 20`. `extraSpaces = 40 % 2 = 0`. Perfect division — each gap gets exactly 20. Total: `20 + 20 + 15 + 20 + 25 = 100`. ✓

</details>

---

### Q19: Terminal width 100. Three words: 20, 15, 26 (total 61). Space = 39. Gaps = 2. `base = 39/2 = 19`, `extra = 39%2 = 1`. How are gaps assigned?

**A)** Gap 1 = 19, Gap 2 = 20
**B)** Gap 1 = 20, Gap 2 = 19
**C)** Gap 1 = 19, Gap 2 = 19 (1 space lost)
**D)** Gap 1 = 20, Gap 2 = 20 (1 space added)

<details><summary>Answer</summary>

**B) Gap 1 = 20, Gap 2 = 19**

The first `extraSpaces` gaps get `baseSpace + 1`. `extraSpaces = 1` → only gap 1 gets the extra.
Total: `20 + 15 + 20 + 26 + 19 = 100`. ✓

The first-gap-gets-extra convention is standard typographic behavior.

</details>

---

### Q20: What does `39 % 2` compute?

**A)** 19
**B)** 1
**C)** 0
**D)** 78

<details><summary>Answer</summary>

**B) 1**

`%` is the modulo (remainder) operator. `39 = 19 × 2 + 1`. The remainder is 1. This tells you how many gaps receive one extra space beyond `baseSpace`.

</details>

---

### Q21: Four words. `totalSpace = 41`, `gaps = 3`. Calculate `baseSpace` and `extraSpaces`.

**A)** `base = 13`, `extra = 2`
**B)** `base = 14`, `extra = 0`
**C)** `base = 13`, `extra = 3`
**D)** `base = 10`, `extra = 11`

<details><summary>Answer</summary>

**A) `base = 13`, `extra = 2`**

`41 / 3 = 13` (integer division). `41 % 3 = 2`.
Gap 1 = 14, Gap 2 = 14, Gap 3 = 13.
`14 + 14 + 13 = 41`. ✓

</details>

---

### Q22: Input is `"hello"` (one word) with `--align=justify`. What should happen?

**A)** Crash — justify needs at least 2 words
**B)** Center the word
**C)** Left-align the word — no gaps to distribute
**D)** Right-align the word

<details><summary>Answer</summary>

**C) Left-align the word**

`gaps = len(words) - 1 = 0`. You cannot divide by zero. The fallback for a single word is left alignment — the spec does not require centering or right-aligning it.

Always guard: `if len(words) <= 1 { return renderLeftAligned(...) }`.

</details>

---

### Q23: Why must justify be applied row by row, not to the complete joined string?

**A)** Go strings do not support multi-line operations
**B)** Each character is 8 rows tall. The gap between words is 8 rows of spaces stacked. You must insert the gap into each of the 8 rows individually — you cannot do this after joining.
**C)** `strings.Join` does not preserve spacing
**D)** Because `unsafe.Pointer` interferes with string operations

<details><summary>Answer</summary>

**B) Each character is 8 rows tall — the gap must be inserted into each row individually**

```
Row 0: [word1_row0][spaces][word2_row0][spaces][word3_row0]
Row 1: [word1_row1][spaces][word2_row1][spaces][word3_row1]
...
Row 7: [word1_row7][spaces][word2_row7][spaces][word3_row7]
```

You render each word into 8 rows, then build each output row by concatenating word rows with the calculated gap. Joining first and then trying to insert gaps would require re-parsing the string — which is error-prone and wrong.

</details>

---

### Q24: In Go's standard library, what function produces a string of N repeated spaces?

**A)** `fmt.Sprintf("%Ns", "")`
**B)** `strings.Repeat(" ", n)`
**C)** `strings.Pad(" ", n)`
**D)** `bytes.Fill(' ', n)`

<details><summary>Answer</summary>

**B) `strings.Repeat(" ", n)`**

```go
gap := strings.Repeat(" ", baseSpace)
// or for the extra-space gaps:
gap := strings.Repeat(" ", baseSpace+1)
```

`strings.Repeat` is in the standard library at https://pkg.go.dev/strings#Repeat.

</details>

---

## SECTION 4: REFACTORING renderLine (3 Questions)

### Q25: `renderLine` currently prints directly. After refactoring it to return `[]string`, how does the non-justify path still produce the same output?

**A)** It automatically produces the same output
**B)** The caller joins the 8 returned rows with `"\n"` and prints the result
**C)** Write a second version of `renderLine` that prints
**D)** Use `fmt.Println` directly on the returned `[]string`

<details><summary>Answer</summary>

**B) The caller joins the 8 returned rows with `"\n"` and prints the result**

```go
rows := renderLine(banner, text)
fmt.Println(strings.Join(rows, "\n"))
```

The refactor is backward-compatible as long as every non-justify caller does this join step. No behavior changes — only where the joining happens.

</details>

---

### Q26: After the refactor, what is the correct return type for `renderLine`?

**A)** `string`
**B)** `[]string`
**C)** `(string, error)`
**D)** `[][]string`

<details><summary>Answer</summary>

**B) `[]string`**

Returns exactly 8 strings — one per art row. The caller decides whether to join them (for normal output) or use them individually (for alignment operations).

</details>

---

### Q27: You refactor `renderLine` to return `[]string`. What is the first thing you should do before committing the change?

**A)** Run `go vet`
**B)** Run all previous test cases and confirm they still pass
**C)** Update `go.mod`
**D)** Rewrite `main.go` from scratch

<details><summary>Answer</summary>

**B) Run all previous test cases and confirm they still pass**

Refactoring changes internal structure without changing external behavior. The only way to confirm you have not broken anything is to run your existing test cases. If they all pass, the refactor is safe. This is why having test cases from the start matters.

</details>

---

## SECTION 5: STANDARD LIBRARY VS THIRD-PARTY (3 Questions)

### Q28: `golang.org/x/term` is maintained by the Go team. Why is it still not in the standard library?

**A)** It is deprecated and will be removed
**B)** It requires a paid license
**C)** The `golang.org/x/` packages are "extended" — maintained by the Go team but versioned and released independently from the standard library, allowing faster iteration and optional adoption
**D)** It only works on Linux

<details><summary>Answer</summary>

**C) The `golang.org/x/` packages are versioned and released independently**

The `golang.org/x/` namespace is the Go team's "experimental" or "extended" space. Packages there are production-quality but not subject to the same compatibility guarantee as the standard library. They can be updated faster, changed more freely, and adopted optionally. `golang.org/x/term`, `golang.org/x/net`, `golang.org/x/crypto` all follow this pattern.

</details>

---

### Q29: If you were building a real production tool (not a 01edu project), which would you use for terminal width? Why?

**A)** Always `syscall` directly — third-party packages are unreliable
**B)** `golang.org/x/term` — it handles Linux, macOS, and Windows with a clean one-line API, and is maintained by the Go team
**C)** Hardcode 80 — terminal width rarely matters
**D)** Read it from an environment variable

<details><summary>Answer</summary>

**B) `golang.org/x/term` — cleaner API, cross-platform, Go team maintained**

In a real tool, `golang.org/x/term.GetSize(int(os.Stdout.Fd()))` is three words and handles all platforms. Your `syscall` implementation is Linux-only and requires understanding `unsafe`. Both are correct — `syscall` teaches you what is underneath, `golang.org/x/term` is what you would ship. Knowing the difference makes you a better engineer.

</details>

---

### Q30: Your `go.mod` currently has no `require` section. After completing this project correctly (standard library only), should `go.mod` still have no `require` section?

**A)** No — `syscall` and `unsafe` require a module declaration
**B)** Yes — `syscall` and `unsafe` are part of the standard library and do not appear in `go.mod`
**C)** It depends on the Go version
**D)** You need to run `go mod tidy` to add them

<details><summary>Answer</summary>

**B) Yes — standard library packages never appear in `go.mod`**

`go.mod` only tracks external module dependencies. `syscall`, `unsafe`, `strings`, `os`, `fmt` are all part of the Go standard library and are always available — no `require` entry needed. If your `go.mod` has a `require` section after this project, you accidentally imported something external.

```
module ascii-art-justify

go 1.21
```

That is all your `go.mod` should contain.

</details>

---

## Score Interpretation

| Score | Result |
|---|---|
| 27–30 | **Exceptional.** Strong understanding of syscalls and the justify algorithm. Start immediately. |
| 23–26 | **Ready.** Review any missed questions before starting — especially the justify section (Q16–Q24). |
| 18–22 | **Study first.** Work through Q16–Q24 on paper with concrete examples before writing any code. |
| Below 18 | **Not ready.** Master the alignment math and the syscall section before starting. |

---

## Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q7 | `syscall.Syscall`, `TIOCGWINSZ`, `winsize` struct, `unsafe.Pointer` |
| Q8–Q15 | Padding formulas, `renderedWidth`, pipeline order |
| Q16–Q24 | Justify algorithm: space distribution, `%` operator, row-by-row assembly |
| Q25–Q27 | Refactoring `renderLine`, backward compatibility |
| Q28–Q30 | Standard library vs third-party, `go.mod`, `golang.org/x/term` awareness |
