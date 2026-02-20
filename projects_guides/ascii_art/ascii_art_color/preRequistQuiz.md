# 🎯 ASCII-Art-Color Prerequisites Quiz
## ANSI Escape Codes · Substring Search · Boolean Masking · Selective Rendering

**Time Limit:** 40 minutes  
**Total Questions:** 25  
**Passing Score:** 20/25 (80%)

> ✅ Pass → You're ready to start ASCII-Art-Color  
> ⚠️ Also Required → ASCII-Art, ASCII-Art-FS, and ASCII-Art-Output must be fully passing

---

## 📋 SECTION 1: ANSI ESCAPE CODES (7 Questions)

### Q1: What is an ANSI escape sequence?

**A)** A special file format for storing colors  
**B)** A sequence of characters starting with `\033[` that terminals interpret as formatting instructions rather than printing them  
**C)** A Go package for terminal colors  
**D)** A CSS-style color specification  

<details><summary>💡 Answer</summary>

**B) A sequence of characters starting with `\033[` that terminals interpret as formatting instructions rather than printing them**

`\033` is the ESC character (ASCII 27, octal 033). When a terminal sees `\033[31m`, it switches foreground color to red instead of printing those characters. This is why color output looks garbled in plain text files — the escape sequences aren't being interpreted.

</details>

---

### Q2: What is the ANSI escape code to set the foreground color to red?

**A)** `\033[0m`  
**B)** `\033[31m`  
**C)** `\033[41m`  
**D)** `\033[1;31m`  

<details><summary>💡 Answer</summary>

**B) `\033[31m`**

Standard foreground color codes: `30`=black, `31`=red, `32`=green, `33`=yellow, `34`=blue, `35`=magenta, `36`=cyan, `37`=white. `41m` would be a red **background** (40–47 are background colors). `1;31m` is bold red.

</details>

---

### Q3: Why is `\033[0m` critical after a colored section?

**A)** It makes the text bold  
**B)** It resets all formatting — without it, every character printed afterwards remains colored  
**C)** It moves the cursor to the next line  
**D)** It's not needed in modern terminals  

<details><summary>💡 Answer</summary>

**B) It resets all formatting — without it, every character printed afterwards remains colored**

Color state persists in the terminal until explicitly reset. If you color `'H'` red and don't reset, `'e'`, `'l'`, `'l'`, `'o'` and everything after will also be red — including your shell prompt.

```go
// Correct pattern:
colored := colorCode + artLine + resetCode
```

</details>

---

### Q4: What is the output of this Go code in a color-capable terminal?
```go
fmt.Println("\033[32mHello\033[0m World")
```

**A)** `\033[32mHello\033[0m World` (printed literally)  
**B)** `Hello World` with "Hello" in green and "World" in default color  
**C)** `Hello World` all in green  
**D)** A compile error  

<details><summary>💡 Answer</summary>

**B) `Hello World` with "Hello" in green and "World" in default color**

The terminal interprets `\033[32m` as "switch to green", renders `Hello`, then `\033[0m` resets, and `World` prints in the default color. In a plain file or non-ANSI terminal, you'd see the raw escape characters.

</details>

---

### Q5: You want to support the color `"red"`. Which Go data structure maps color names to ANSI codes?

**A)** `[]string{"red", "\033[31m"}`  
**B)** `map[string]string{"red": "\033[31m", "green": "\033[32m"}`  
**C)** `type Color struct{ name, code string }`  
**D)** `const red = 31`  

<details><summary>💡 Answer</summary>

**B) `map[string]string{"red": "\033[31m", "green": "\033[32m"}`**

A map gives O(1) lookup and makes unknown color detection trivial with the two-value lookup: `code, ok := colorMap[name]`. If `!ok`, the color name is invalid.

</details>

---

### Q6: What should your program do if the user passes `--color=purple` and your map doesn't contain `"purple"`?

**A)** Default to red  
**B)** Print with no color  
**C)** Print the usage message and exit  
**D)** Panic  

<details><summary>💡 Answer</summary>

**C) Print the usage message and exit**

An unrecognized color is invalid input. The user should be told what the valid options are. Silently ignoring or defaulting to another color would be confusing behavior.

</details>

---

### Q7: The RGB escape sequence format is `\033[38;2;R;G;Bm`. What does `\033[38;2;255;0;0m` produce?

**A)** A blue foreground  
**B)** A red foreground using full RGB  
**C)** A red background  
**D)** Bold text  

<details><summary>💡 Answer</summary>

**B) A red foreground using full RGB**

`38;2` means "set foreground using 24-bit RGB". `R=255, G=0, B=0` is pure red. This allows any of 16 million colors instead of just the 8 named ANSI colors. Not all terminals support it, but most modern ones do.

</details>

---

## 📋 SECTION 2: SUBSTRING SEARCHING (6 Questions)

### Q8: What does `strings.Index("a king kitten have kit", "kit")` return?

**A)** `0`  
**B)** `7`  
**C)** `2`  
**D)** `true`  

<details><summary>💡 Answer</summary>

**B) `7`**

`strings.Index` returns the **byte index** of the first occurrence. Counting: `"a king "` is 7 characters (0–6), so `"kit"` starts at index `7`. It returns `-1` if not found.

</details>

---

### Q9: You need to find ALL positions where `"kit"` appears in `"a king kitten have kit"`. `strings.Index` only finds the first. How do you find all occurrences?

**A)** Call `strings.Index` in a loop, advancing the search start past each match  
**B)** Use `strings.Count` which returns all positions  
**C)** Use `strings.Fields` to split and search each word  
**D)** It's impossible with the standard library  

<details><summary>💡 Answer</summary>

**A) Call `strings.Index` in a loop, advancing the search start past each match**

```go
func findAll(s, sub string) []int {
    var positions []int
    offset := 0
    for {
        idx := strings.Index(s[offset:], sub)
        if idx == -1 { break }
        positions = append(positions, offset+idx)
        offset += idx + 1  // advance past current match
    }
    return positions
}
```

`strings.Count` counts occurrences but doesn't return positions. Option C splits by word — misses `"kit"` inside `"kitten"`.

</details>

---

### Q10: For input `"a king kitten have kit"` and substring `"kit"`, list ALL starting byte indices where `"kit"` appears.

**A)** `[7]` — only the word `"kit"` at the end  
**B)** `[7, 18]` — `"kitten"` starts at 7, standalone `"kit"` starts at 18  
**C)** `[7, 19]`  
**D)** `[8, 19]`  

<details><summary>💡 Answer</summary>

**B) `[7, 18]` — `"kitten"` starts at 7, standalone `"kit"` starts at 18**

Count the characters: `"a king "` = 7 chars → `"kitten"` starts at index 7. `"a king kitten have "` = 19 chars... let's count carefully:
- `a` = 0, ` ` = 1, `k` = 2, `i` = 3, `n` = 4, `g` = 5, ` ` = 6, `k` = 7 → `"kitten"` at 7 ✓
- `"a king kitten have "` → 7+6+1+4+1 = 19... `"kit"` at 19.

Both occurrences of `"kit"` must be colored — including when it's inside `"kitten"`. The spec explicitly requires this.

</details>

---

### Q11: `strings.Contains("kitten", "kit")` returns:

**A)** `false` — "kit" is a different word  
**B)** `true` — "kit" is a substring of "kitten"  
**C)** `3` — the number of characters matched  
**D)** `"ten"` — the remaining part  

<details><summary>💡 Answer</summary>

**B) `true`**

`strings.Contains` does pure substring matching — it doesn't care about word boundaries. `"kit"` is contained within `"kitten"`. This is exactly the behavior required by the spec (color all occurrences including inside words).

</details>

---

### Q12: What is the difference between `strings.Index` and `strings.Contains`?

**A)** No practical difference  
**B)** `strings.Index` returns the position (-1 if not found); `strings.Contains` returns a bool  
**C)** `strings.Contains` is case-insensitive  
**D)** `strings.Index` works with runes, `strings.Contains` works with bytes  

<details><summary>💡 Answer</summary>

**B) `strings.Index` returns the position (-1 if not found); `strings.Contains` returns a bool**

For finding all occurrences you need positions, so `strings.Index` is what you use in the search loop. `strings.Contains` is for a quick yes/no check.

</details>

---

### Q13: You create a `[]bool` mask of length `len(input)` where `true` means "this character should be colored." For input `"kit"` (3 characters) and substring `"kit"`, what does the mask look like?

**A)** `[false, false, false]`  
**B)** `[true, true, true]`  
**C)** `[true, false, false]`  
**D)** `[false, false, true]`  

<details><summary>💡 Answer</summary>

**B) `[true, true, true]`**

`"kit"` starts at index 0, length 3, so positions 0, 1, and 2 are all `true`. The mask marks every character position that belongs to at least one occurrence of the substring.

</details>

---

## 📋 SECTION 3: SELECTIVE RENDERING (6 Questions)

### Q14: Your render loop processes each character in the input text. At which point do you apply the ANSI color code?

**A)** Once at the start of the entire output  
**B)** Once per row (8 times per character)  
**C)** Before each art line of a colored character, with a reset after  
**D)** After the entire output is assembled  

<details><summary>💡 Answer</summary>

**C) Before each art line of a colored character, with a reset after**

Each character spans 8 rows. For each of those 8 rows, if the character should be colored:
```go
rowContent = colorCode + charArtLine + resetCode
```
If not colored, just `charArtLine`. You wrap each row individually — not the whole character block at once — so the color boundaries are clean per row.

</details>

---

### Q15: What is the output if you apply color wrapping to the entire 8-row block of a character at once (instead of per row)?

**A)** Identical visual result  
**B)** The color may bleed across row boundaries, looking wrong in some terminals  
**C)** A compile error  
**D)** Faster rendering  

<details><summary>💡 Answer</summary>

**B) The color may bleed across row boundaries, looking wrong in some terminals**

Each rendered row is a separate `fmt.Println` call (or `\n` in the string). Wrapping the color around row 1 content and putting the reset at the end of the block (after row 8) works correctly in most terminals, but wrapping per-row is more explicit and guarantees correct behavior across different terminal emulators.

</details>

---

### Q16: When no substring is provided (`go run . --color=red "hello"`), the entire string should be colored. What is the simplest way to implement this?

**A)** Special-case the logic with an `if colorsAll` flag  
**B)** When substring is empty, `findOccurrences` returns a `[]bool` of all `true` values  
**C)** Wrap the entire output string in `colorCode + output + resetCode`  
**D)** Apply the color in `loadBanner`  

<details><summary>💡 Answer</summary>

**B) When substring is empty, `findOccurrences` returns a `[]bool` of all `true` values**

This reuses the same code path. Your render function doesn't need to know whether "all" or "some" characters are colored — it just reads the mask. When substring is `""`, mark all positions as `true`.

```go
if sub == "" {
    for i := range mask { mask[i] = true }
}
```

</details>

---

### Q17: You're rendering `"kitten"` with `"kit"` colored. The mask is `[true, true, true, false, false, false]`. In row 2 of the output, what is the correct way to assemble characters `'k'`, `'i'`, `'t'`, `'t'`, `'e'`, `'n'`?

**A)** `colorCode + "kit" + "ten" + resetCode`  
**B)** For each char: if `mask[i]` wrap with color; otherwise append raw  
**C)** Color the entire word, then decolor the last 3 chars  
**D)** Search for "kit" in the rendered output and wrap it  

<details><summary>💡 Answer</summary>

**B) For each char: if `mask[i]` wrap with color; otherwise append raw**

You process character by character. For each index `i`, check `mask[i]`:
- If `true`: `sb.WriteString(colorCode + charArtRow + resetCode)`
- If `false`: `sb.WriteString(charArtRow)`

This gives precise control over which characters are colored regardless of position.

</details>

---

### Q18: The input string is `"kit"` (3 chars). Your mask is `[]bool` with 3 elements. When you loop with `for i, ch := range input`, what are the valid mask indices?

**A)** 0, 1, 2 — but only if all characters are single-byte ASCII  
**B)** 0, 1, 2 always — because `range` gives you rune indices  
**C)** 0, 1, 2 for ASCII; would be different for multi-byte Unicode characters  
**D)** Depends on the input  

<details><summary>💡 Answer</summary>

**C) 0, 1, 2 for ASCII; would be different for multi-byte Unicode characters**

The mask has `len(input)` elements where `len` counts **bytes**. For pure ASCII (which the banner files support — printable ASCII only), each character is 1 byte, so byte index = character index. This is why the project works cleanly — but be aware this assumption breaks for Unicode input.

</details>

---

### Q19: Your program colors `"kit"` in `"a king kitten have kit"`. How many characters in the output will be colored?

**A)** 3 — only the standalone word `"kit"`  
**B)** 6 — `"kit"` in `"kitten"` and `"kit"` at the end  
**C)** 9 — the entire word `"kitten"` plus `"kit"`  
**D)** 22 — the entire string  

<details><summary>💡 Answer</summary>

**B) 6 — `"kit"` in `"kitten"` and `"kit"` at the end**

Only the 3 characters matching `"kit"` in each occurrence are colored. The `"ten"` in `"kitten"` is NOT colored — only the matched substring positions are marked in the mask.

</details>

---

## 📋 SECTION 4: PARSING & INTEGRATION (6 Questions)

### Q20: The call is `go run . --color=red kit "a king kitten have kit"`. What are the correct values after parsing?

**A)** `color="red"`, `substring="kit"`, `input="a king kitten have kit"`  
**B)** `color="red kit"`, `input="a king kitten have kit"`  
**C)** `color="red"`, `input="kit"`, `banner="a king kitten have kit"`  
**D)** `color="red"`, `input="kit a king kitten have kit"`  

<details><summary>💡 Answer</summary>

**A) `color="red"`, `substring="kit"`, `input="a king kitten have kit"`**

After stripping `--color=red` from args, remaining args are `["kit", "a king kitten have kit"]`. When a color flag is present AND 2 args remain: first is the substring, second is the full string to render.

</details>

---

### Q21: The call is `go run . --color=blue "hello"` (no substring). What are the correct values?

**A)** `color="blue"`, `substring=""`, `input="hello"`  
**B)** `color="blue"`, `substring="hello"`, `input=""`  
**C)** Error — substring is required  
**D)** `color="blue"`, `substring="hello"`, no input  

<details><summary>💡 Answer</summary>

**A) `color="blue"`, `substring=""`, `input="hello"`**

When color flag is present and only 1 arg remains after stripping the flag: that arg is the input string, and the substring is empty (meaning color the whole thing).

</details>

---

### Q22: What is the valid number of user arguments when `--color` flag is present? (Not counting `os.Args[0]`)

**A)** Always exactly 2: substring and string  
**B)** 1, 2, or 3: `[STRING]`, `[SUBSTRING STRING]`, or `[SUBSTRING STRING BANNER]`  
**C)** Always exactly 3: color, substring, string  
**D)** Only 1: the string  

<details><summary>💡 Answer</summary>

**B) 1, 2, or 3 remaining args after the flag is stripped: `[STRING]`, `[SUBSTRING STRING]`, or `[SUBSTRING STRING BANNER]`**

The flag itself is stripped. Then:
- 1 remaining: `STRING` (color whole thing, default banner)
- 2 remaining: `SUBSTRING STRING` (color substring, default banner)
- 3 remaining: `SUBSTRING STRING BANNER` (color substring, specific banner)

Anything else → usage message.

</details>

---

### Q23: Should `--color` and `--output` flags be compatible (usable together)?

**A)** No — they conflict  
**B)** Yes — `--color=red --output=file.txt "hello"` should write colored output to the file  
**C)** Only one flag can be used at a time  
**D)** The `--output` flag overrides `--color`  

<details><summary>💡 Answer</summary>

**B) Yes — they should be compatible**

Both flags modify how the output is handled — one adds color, one redirects destination. They operate independently. Your arg-scanning loop should be able to detect and strip both flags from `os.Args` before parsing the remaining args.

</details>

---

### Q24: After stripping all flags, your remaining `os.Args` slice is `["program", "hello", "shadow"]`. What does this mean?

**A)** Input is `"hello"`, banner is `"shadow"`, no substring  
**B)** Substring is `"hello"`, input is `"shadow"`  
**C)** This is invalid — 2 args without a color flag  
**D)** Input is `"hello shadow"`  

<details><summary>💡 Answer</summary>

**A) Input is `"hello"`, banner is `"shadow"`, no substring**

When there's no `--color` flag (or it's been processed and the substring was provided separately), the remaining args follow the ascii-art-fs rules: `[STRING]` or `[STRING BANNER]`. So `os.Args[1]="hello"`, `os.Args[2]="shadow"`.

</details>

---

### Q25: You've implemented coloring correctly for `standard.txt`. A tester runs `go run . --color=red "hello" shadow`. What must be true for this to work?

**A)** The color logic is banner-independent — it works the same regardless of which banner is loaded  
**B)** You need a separate color implementation for each banner  
**C)** Shadow doesn't support colors  
**D)** You need to reload the banner in the color module  

<details><summary>💡 Answer</summary>

**A) The color logic is banner-independent — it works the same regardless of which banner is loaded**

The coloring works at the character level: wrap individual character art rows with ANSI codes based on the mask. It doesn't matter which banner's art is used — the wrapping logic is identical. This is why separating concerns (banner loading vs. color rendering) leads to clean, reusable code.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 23–25 ✅ | **Excellent.** Start ASCII-Art-Color. |
| 20–22 ✅ | **Ready.** Review missed questions first. |
| 15–19 ⚠️ | **Study first.** Focus on ANSI codes and the boolean mask approach. |
| Below 15 ❌ | **Not ready.** Review ANSI escape sequences, `strings.Index`, and selective rendering logic. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q7 | ANSI escape code format, color codes, reset, RGB format |
| Q8–Q13 | `strings.Index`, finding all occurrences, boolean mask design |
| Q14–Q19 | Per-row color wrapping, mask-driven rendering, "color all" case |
| Q20–Q25 | Flag parsing with substring, multi-flag compatibility, banner independence |