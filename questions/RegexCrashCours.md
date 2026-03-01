# 🎓 JavaScript RegEx — Complete Crash Course
## From Zero to Fluent 💪

> Read every section. Run every example. Do every exercise. Then do the quizzes.

---

## CHAPTER 1 — What Even IS a Regular Expression?

A regular expression (RegEx) is a **pattern** used to search inside a string.

Think of it like a very powerful "find" feature — instead of searching for exact text like `"hello"`, you can search for patterns like "any word that starts with a capital letter" or "any sequence of 3 digits."

```js
// Without RegEx — only finds exact text
"hello world".includes("hello")  // true

// With RegEx — finds a PATTERN
/hello/.test("hello world")      // true
/h.llo/.test("hello world")      // true  ← the dot matches ANY character
/h.llo/.test("hxllo world")      // true  ← 'x' matches the dot too!
```

A RegEx is written between two forward slashes: `/pattern/`

---

## CHAPTER 2 — Your First Pattern

The simplest pattern is just plain text:

```js
const pattern = /hi/;

pattern.test("say hi there");   // true  — 'hi' found somewhere
pattern.test("HIGH FIVE");      // false — RegEx is case-sensitive by default
pattern.test("this");           // true  — 'hi' is inside 't-hi-s'!
```

⚠️ RegEx searches ANYWHERE in the string unless you tell it otherwise (anchors — coming soon).

---

## CHAPTER 3 — Flags

Flags go AFTER the closing slash and change how the pattern behaves.

| Flag | Name | What it does |
|------|------|--------------|
| `g`  | global | Find ALL matches, not just the first |
| `i`  | case-insensitive | `A` matches `a` |
| `m`  | multiline | `^` and `$` apply to each line, not whole string |

```js
// No flags — finds first match only, case-sensitive
"aAbBaA".match(/a/)    // ['a'] — only first 'a'

// g flag — finds ALL matches
"aAbBaA".match(/a/g)   // ['a', 'a'] — both lowercase a's

// i flag — case-insensitive
"aAbBaA".match(/a/i)   // ['a'] — first match, any case

// gi flags — all matches, any case
"aAbBaA".match(/a/gi)  // ['a', 'A', 'a', 'A'] — all four
```

### The m (multiline) flag

```js
const text = "line one\nline two\nline three";

// Without m — ^ matches only start of whole string
text.match(/^line/g)    // ['line'] — only first line

// With m — ^ matches start of EACH line
text.match(/^line/gm)   // ['line', 'line', 'line'] — all three!
```

### 🏋️ Mini Exercise 1
```js
const str = "Cat cat CAT";
// Q1: Match all occurrences of 'cat' regardless of case
// Q2: How many matches do you get?

// Answer:
str.match(/cat/gi)  // ['Cat', 'cat', 'CAT'] — 3 matches
```

---

## CHAPTER 4 — Anchors: ^ and $

Anchors don't match characters — they match **positions** in the string.

| Anchor | Matches |
|--------|---------|
| `^`    | Start of string (or start of line with `m` flag) |
| `$`    | End of string (or end of line with `m` flag) |

```js
// ^ = must be at the START
/^hi/.test("hi there")    // true  — 'hi' is at the start
/^hi/.test("say hi")      // false — 'hi' is NOT at the start
/^hi/.test("higher")      // true  — starts with 'hi'

// $ = must be at the END
/hi$/.test("say hi")      // true  — 'hi' is at the end
/hi$/.test("hi there")    // false — 'hi' is NOT at the end
/hi$/.test("alihi")       // true  — ends with 'hi'

// ^ AND $ together = must be the ENTIRE string
/^hi$/.test("hi")         // true  — the whole string is exactly 'hi'
/^hi$/.test("hi there")   // false — more than just 'hi'
/^hi$/.test("say hi")     // false — more than just 'hi'
```

Think of it like this:
- `^` = "the very beginning"
- `$` = "the very end"
- Together = "nothing before AND nothing after"

### 🏋️ Mini Exercise 2
```js
// Q1: Write a pattern that matches strings STARTING with 'Error'
// Q2: Write a pattern that matches strings ENDING with a period '.'
// Q3: Write a pattern that matches ONLY the string 'ok'

// Answers:
/^Error/
/\.$/      // dot needs escaping — see Chapter 6
/^ok$/
```

---

## CHAPTER 5 — Character Sets: [ ]

A character set `[abc]` matches **any ONE character** from the list inside the brackets.

```js
/[aeiou]/.test("hello")     // true  — 'e' matches
/[aeiou]/.test("gym")       // false — no vowels
/[aeiou]/g                  // matches each vowel individually

"hello".match(/[aeiou]/g)   // ['e', 'o'] — both vowels
```

### Ranges inside [ ]

```js
[a-z]    // any lowercase letter a through z
[A-Z]    // any uppercase letter A through Z
[0-9]    // any digit 0 through 9
[a-zA-Z] // any letter, upper or lower
[a-zA-Z0-9] // any letter or digit
```

```js
"h3ll0".match(/[a-z]/g)     // ['h', 'l', 'l'] — only letters
"h3ll0".match(/[0-9]/g)     // ['3', '0'] — only digits
```

### Negated sets: [^ ]

Put `^` INSIDE `[ ]` to mean "anything EXCEPT these":

```js
[^aeiou]   // any character that is NOT a vowel
[^0-9]     // any character that is NOT a digit
```

```js
"hello".match(/[^aeiou]/g)   // ['h', 'l', 'l'] — non-vowels
```

⚠️ `^` inside `[ ]` means NOT. `^` outside `[ ]` means start of string. Same symbol, different meanings!

### 🏋️ Mini Exercise 3
```js
// Q1: Match all vowels in "javascript is awesome"
// Q2: Match all NON-vowels in "hi"
// Q3: Match all digits in "abc123def456"

// Answers:
"javascript is awesome".match(/[aeiou]/gi)   // ['a', 'a', 'i', 'i', 'a', 'e', 'o', 'e']
"hi".match(/[^aeiou]/g)                      // ['h']
"abc123def456".match(/[0-9]/g)               // ['1','2','3','4','5','6']
```

---

## CHAPTER 6 — Special Characters & Escaping

Some characters have special meaning in RegEx. To match them literally, put `\` before them.

| Character | Meaning in RegEx | To match literally |
|-----------|------------------|--------------------|
| `.`       | Any character    | `\.`               |
| `*`       | Quantifier       | `\*`               |
| `+`       | Quantifier       | `\+`               |
| `?`       | Quantifier       | `\?`               |
| `(`       | Group start      | `\(`               |
| `)`       | Group end        | `\)`               |
| `[`       | Set start        | `\[`               |
| `]`       | Set end          | `\]`               |
| `^`       | Anchor / Negate  | `\^`               |
| `$`       | Anchor           | `\$`               |
| `\`       | Escape char      | `\\`               |

```js
// . matches ANY character (including dot!)
/3.14/.test("3x14")    // true  — x matches the dot!
/3.14/.test("3.14")    // true  — dot also matches dot

// \. matches ONLY a literal dot
/3\.14/.test("3x14")   // false — x is not a dot
/3\.14/.test("3.14")   // true  — literal dot
```

---

## CHAPTER 7 — Shorthand Character Classes

These are shortcuts for common character sets:

| Shorthand | Equivalent     | Matches |
|-----------|----------------|---------|
| `\d`      | `[0-9]`        | Any digit |
| `\D`      | `[^0-9]`       | Any NON-digit |
| `\w`      | `[a-zA-Z0-9_]` | Any word character (letter, digit, underscore) |
| `\W`      | `[^a-zA-Z0-9_]`| Any NON-word character |
| `\s`      | `[ \t\n\r]`    | Any whitespace (space, tab, newline) |
| `\S`      | `[^ \t\n\r]`   | Any NON-whitespace |
| `.`       | (almost anything) | Any character EXCEPT newline |

```js
"hello world 123".match(/\d/g)   // ['1','2','3']
"hello world 123".match(/\D/g)   // ['h','e','l','l','o',' ','w','o','r','l','d',' ']
"hello world 123".match(/\w/g)   // ['h','e','l','l','o','w','o','r','l','d','1','2','3']
"hello world 123".match(/\s/g)   // [' ', ' ']
"hello world 123".match(/\W/g)   // [' ', ' ']
```

### Combining shorthands

```js
/\w\s\d/   // word char + whitespace + digit — e.g. "a 3" or "z 7"

"a 3 and b 44".match(/\w\s\d/g)  // ['a 3'] — b 44 fails because 44 is two digits
```

### 🏋️ Mini Exercise 4
```js
// Q1: Match all words (sequences of word characters) in "hello world 123"
// Q2: What does \w\s\d match in "x 5 and y 20"?
// Q3: What does \D+ match in "abc123def"?

// Answers:
"hello world 123".match(/\w+/g)      // ['hello', 'world', '123']  (+ means one or more — next chapter)
"x 5 and y 20".match(/\w\s\d/g)     // ['x 5']  — y 20 fails ('20' is two digits)
"abc123def".match(/\D+/g)            // ['abc', 'def']
```

---

## CHAPTER 8 — Quantifiers

Quantifiers say **how many times** the previous element must appear.

| Quantifier | Meaning |
|------------|---------|
| `*`        | 0 or more |
| `+`        | 1 or more |
| `?`        | 0 or 1 (optional) |
| `{n}`      | Exactly n times |
| `{n,}`     | At least n times |
| `{n,m}`    | Between n and m times (inclusive) |

```js
/\d+/    // one or more digits
/\d*/    // zero or more digits
/\d?/    // zero or one digit
/\d{3}/  // exactly 3 digits
/\d{2,4}/ // 2, 3, or 4 digits
```

```js
"color colour".match(/colou?r/g)  // ['color', 'colour'] — u is optional!

"aaa".match(/a{2}/g)   // ['aa'] — matches first two a's
"aaa".match(/a{2,3}/g) // ['aaa'] — matches all three (greedy — takes as many as possible)

"phone: 123-456-7890".match(/\d{3}-\d{3}-\d{4}/)  // ['123-456-7890']
```

### Quantifiers apply to the thing immediately before them

```js
/ab+/    // a followed by ONE OR MORE b's
/[ab]+/  // one or more a's or b's (any combination)
/(ab)+/  // one or more repetitions of the GROUP "ab"
```

### 🏋️ Mini Exercise 5
```js
// Q1: Match one or more digits in "price: 42 dollars"
// Q2: Match 'colour' or 'color' using ?
// Q3: Match exactly 3 letters followed by exactly 2 digits

// Answers:
"price: 42 dollars".match(/\d+/)    // ['42']
/colou?r/                            // matches both spellings
/[a-zA-Z]{3}\d{2}/                  // e.g. matches 'abc12'
```

---

## CHAPTER 9 — Greedy vs Lazy

By default, quantifiers are **greedy** — they match as MUCH as possible.
Add `?` after a quantifier to make it **lazy** — match as LITTLE as possible.

```js
const str = "<b>bold</b> and <i>italic</i>";

// Greedy — matches from first < to LAST >
str.match(/<.+>/)    // ['<b>bold</b> and <i>italic</i>'] — too much!

// Lazy — matches from < to the NEXT >
str.match(/<.+?>/)   // ['<b>'] — stops at first possible >
str.match(/<.+?>/g)  // ['<b>', '</b>', '<i>', '</i>'] — all tags
```

| Greedy | Lazy | Meaning |
|--------|------|---------|
| `*`    | `*?` | 0 or more, as few as possible |
| `+`    | `+?` | 1 or more, as few as possible |
| `?`    | `??` | 0 or 1, preferring 0 |
| `{n,m}`| `{n,m}?` | n to m, as few as possible |

```js
"aXbXc".match(/a.+c/)    // ['aXbXc'] — greedy, takes everything
"aXbXc".match(/a.+?c/)   // ['aXbXc'] — lazy, but still needs to reach c

"aXcXc".match(/a.+c/)    // ['aXcXc'] — greedy, takes to LAST c
"aXcXc".match(/a.+?c/)   // ['aXc']   — lazy, stops at FIRST c
```

### 🏋️ Mini Exercise 6
```js
const html = '<a href="http://example.com">click here</a>';

// Q1: Extract just the URL using lazy matching
// Q2: What would greedy matching give you instead?

// Answers:
html.match(/href="(.+?)"/)   // captures 'http://example.com' (lazy stops at first ")
html.match(/href="(.+)"/)    // greedy — if multiple quotes exist it would overshoot
```

---

## CHAPTER 10 — Capturing Groups: ( )

Parentheses `( )` create a **capturing group** — they:
1. Group part of the pattern together (for quantifiers)
2. **Capture** the matched text so you can extract it

```js
// Group for quantifiers
/(ab)+/.test("ababab")   // true — matches 'ab' repeated

// Capture for extraction
"2024-03-15".match(/(\d{4})-(\d{2})-(\d{2})/)
// Full result: ['2024-03-15', '2024', '03', '15']
// Index 0: full match
// Index 1: first group  (year)
// Index 2: second group (month)
// Index 3: third group  (day)
```

### Non-capturing group: (?: )

If you want to group without capturing, use `(?:...)`:

```js
"ababab".match(/(?:ab)+/)    // ['ababab'] — no extra captured groups
"ababab".match(/(ab)+/)      // ['ababab', 'ab'] — 'ab' captured (last repetition)
```

### Groups with replace()

```js
// Rearrange date format: YYYY-MM-DD → DD/MM/YYYY
"2024-03-15".replace(/(\d{4})-(\d{2})-(\d{2})/, "$3/$2/$1")
// Result: '15/03/2024'
// $1 = first group, $2 = second, $3 = third
```

### Groups with exec() or matchAll()

```js
const regex = /([A-Z]{3})(\d+\.\d+)/g;
const str = "USD12.31 and EUR9.99";

let match;
while ((match = regex.exec(str)) !== null) {
    console.log(match[0]);  // full: 'USD12.31'
    console.log(match[1]);  // group 1: 'USD'
    console.log(match[2]);  // group 2: '12.31'
}
```

### 🏋️ Mini Exercise 7
```js
// Q1: Extract the integer and decimal parts of "42.75"
// Q2: Use replace to swap first and last name: "John Smith" → "Smith, John"

// Answers:
"42.75".match(/(\d+)\.(\d+)/)
// ['42.75', '42', '75'] — index 1 = '42', index 2 = '75'

"John Smith".replace(/(\w+)\s(\w+)/, "$2, $1")
// 'Smith, John'
```

---

## CHAPTER 11 — Lookahead & Lookbehind

Lookarounds let you match something **based on what comes before or after it**, WITHOUT including that context in the match.

### Lookahead: (?=...) and (?!...)

```
X(?=Y)   — match X only if followed by Y     (positive lookahead)
X(?!Y)   — match X only if NOT followed by Y  (negative lookahead)
```

```js
// Match 'cat' only if followed by 's'
"cats and cat".match(/cat(?=s)/g)    // ['cat'] — only the first one

// Match 'cat' only if NOT followed by 's'
"cats and cat".match(/cat(?!s)/g)    // ['cat'] — only the second one
```

### Lookbehind: (?<=...) and (?<!...)

```
(?<=Y)X  — match X only if preceded by Y      (positive lookbehind)
(?<!Y)X  — match X only if NOT preceded by Y  (negative lookbehind)
```

```js
// Match digits only if preceded by '$'
"$100 and 200".match(/(?<=\$)\d+/g)    // ['100'] — only after $

// Match digits only if NOT preceded by '$'
"$100 and 200".match(/(?<!\$)\d+/g)    // ['200'] — only without $
```

### Key insight: lookarounds don't consume characters

```js
// Without lookahead — 'on' is consumed, lost in the split
"button".replace(/on/, "")    // 'butt'

// With lookahead — position matched but 'on' stays
"button".match(/butt(?=on)/)  // ['butt'] — matched 'butt', 'on' not consumed
```

### Word boundaries: \b

`\b` matches the **position** between a word character and a non-word character.

```js
"ion is in action".match(/\bion\b/g)    // ['ion'] — standalone word only
"ion is in action".match(/ion/g)         // ['ion', 'ion'] — both occurrences
```

### 🏋️ Mini Exercise 8
```js
// Q1: Match numbers only if followed by 'px'
// Q2: Match words ending in 'tion' using lookbehind
// Q3: Match 'ion' only when preceded by 't'

// Answers:
"12px and 34em".match(/\d+(?=px)/g)          // ['12']
"action nation".match(/\w+(?=tion)/g)         // hmm... extracts before 'tion'
"action caution".match(/(?<=t)ion/g)          // ['ion', 'ion'] — the ion parts
```

---

## CHAPTER 12 — JavaScript RegEx Methods

### method 1: .test() — returns true/false

```js
const regex = /\d+/;
regex.test("hello 42")    // true
regex.test("hello")       // false
```

Use when you just need to know IF a match exists.

---

### method 2: .match() — returns matches

```js
// Without g flag — returns first match with details
"hello 42 world 7".match(/\d+/)
// ['42', index: 6, input: 'hello 42 world 7', groups: undefined]

// With g flag — returns array of ALL matches (no details)
"hello 42 world 7".match(/\d+/g)
// ['42', '7']

// No match — returns null
"hello".match(/\d+/)      // null
"hello".match(/\d+/g)     // null
```

⚠️ Always guard against null: `str.match(/pattern/g) || []`

---

### method 3: .replace() — replace matches

```js
// Replace first match
"hello world".replace(/o/, "0")     // 'hell0 world'

// Replace ALL matches (use g flag)
"hello world".replace(/o/g, "0")    // 'hell0 w0rld'

// Replace with a function
"hello world".replace(/[aeiou]/g, (match) => match.toUpperCase())
// 'hEllO wOrld'

// Replace using capture groups
"John Smith".replace(/(\w+) (\w+)/, "$2 $1")   // 'Smith John'
```

---

### method 4: .exec() — stateful matching

```js
const regex = /\d+/g;
const str = "a1 b2 c3";

let match;
while ((match = regex.exec(str)) !== null) {
    console.log(match[0], "at index", match.index);
}
// '1' at index 1
// '2' at index 4
// '3' at index 7
```

Use `exec` in a loop for complex extraction — especially with capturing groups.

---

### method 5: .matchAll() — modern exec loop

```js
const regex = /(\d+)/g;
const str = "a1 b2 c3";

for (const match of str.matchAll(regex)) {
    console.log(match[0], match[1], match.index);
}
// '1' '1' 1
// '2' '2' 4
// '3' '3' 7
```

`matchAll` requires the `g` flag. Returns an iterator of full match objects.

---

### .split() with RegEx

```js
"one1two2three".split(/\d/)     // ['one', 'two', 'three']
"a,b;c d".split(/[,; ]/)        // ['a', 'b', 'c', 'd']
```

---

### Method Summary

| Method | Returns | Use for |
|--------|---------|---------|
| `.test(str)` | `boolean` | Does it match? |
| `str.match(regex)` | array or null | Get matches |
| `str.replace(regex, rep)` | string | Transform text |
| `regex.exec(str)` | array or null | Loop through matches |
| `str.matchAll(regex)` | iterator | Loop with groups |
| `str.split(regex)` | array | Split on pattern |

---

## CHAPTER 13 — URL Matching Pattern

URLs are a great real-world RegEx target. Let's break down a URL pattern:

```
https://example.com/path?key=value&key2=value2
```

```js
// Basic URL pattern
/https?:\/\/[^\s]+/

// Breaking it down:
// https?    — 'http' or 'https' (s is optional)
// :\/\/     — '://' (slashes need escaping)
// [^\s]+    — one or more non-whitespace characters (the rest of the URL)
```

```js
// Match query parameters: key=value pairs
/[?&](\w+)=(\w+)/g

// Count parameters — each & adds one more
"?a=1&b=2&c=3".match(/&/g)   // ['&', '&'] — 2 ampersands = 3 params total
```

---

## CHAPTER 14 — IP Address Matching

IPv4: four groups of 0-255 separated by dots.

```
192.168.0.1
```

The trick: each octet must be 0-255. RegEx for that:

```js
// 0-199: [01]?\d{1,2}   — 0, 1, 99, 150, etc.
// 200-249: 2[0-4]\d
// 250-255: 25[0-5]

const octet = /(?:25[0-5]|2[0-4]\d|[01]?\d{1,2})/;
// Combined into full IP:
const ip = /(?:25[0-5]|2[0-4]\d|[01]?\d{1,2})(?:\.(?:25[0-5]|2[0-4]\d|[01]?\d{1,2})){3}/;
```

---

## CHAPTER 15 — Common Patterns Cheatsheet

```js
// Email (simplified)
/[\w.+-]+@[\w-]+\.[a-zA-Z]{2,}/

// US phone number
/\d{3}[-.\s]\d{3}[-.\s]\d{4}/

// Hex color
/#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})/

// Any word
/\b\w+\b/g

// Line starting with whitespace
/^\s+/m

// Empty line
/^$/m

// Repeated words (duplicate word detector)
/\b(\w+)\s+\1\b/   // \1 = backreference to group 1
```

### Backreferences: \1, \2, ...

```js
// Match repeated words
"the the cat sat".match(/\b(\w+)\s+\1\b/)
// ['the the', 'the'] — 'the' repeated twice
```

---

## CHAPTER 16 — Building Complex Patterns (The Process)

When facing a complex RegEx task, follow this process:

1. **Describe the pattern in English first**
2. **Break it into small pieces**
3. **Build piece by piece**
4. **Test each piece**
5. **Combine**

### Example: match `"letter space single-digit"` but not followed by a letter

English: "a word character, then a space, then a single digit, then not a word character"

```js
// Piece 1: a letter
/[a-zA-Z]/

// Piece 2: a space
/\s/

// Piece 3: single digit (not followed by another digit)
/\d(?!\d)/

// But also: digit not followed by a LETTER
/\d(?![a-zA-Z])/

// Combined:
/[a-zA-Z]\s\d(?![a-zA-Z])/g
```

---

## CHAPTER 17 — Mistakes Beginners Make

### 1. Forgetting the `g` flag for multiple matches

```js
"aaa".match(/a/)    // ['a'] — only first!
"aaa".match(/a/g)   // ['a', 'a', 'a'] — all!
```

### 2. Not escaping special characters

```js
"3.14".match(/3.14/)    // matches '3x14' too! Dot = any char
"3.14".match(/3\.14/)   // correct — literal dot
```

### 3. Confusing `^` inside vs outside `[]`

```js
/^[aeiou]/    // string starts with a vowel
/[^aeiou]/    // any character that is NOT a vowel
```

### 4. match() returns null (not empty array)

```js
// WRONG:
const matches = "hello".match(/\d/g);
console.log(matches.length);  // TypeError — matches is null!

// CORRECT:
const matches = "hello".match(/\d/g) || [];
console.log(matches.length);  // 0
```

### 5. `.` doesn't match newlines

```js
const str = "line1\nline2";
str.match(/.+/)    // ['line1'] — stops at newline!
str.match(/[\s\S]+/)   // matches everything including newlines
```

### 6. Greedy trap with HTML

```js
"<b>one</b><b>two</b>".match(/<b>.+<\/b>/)
// ['<b>one</b><b>two</b>'] — too greedy! Use .+? instead
```

---

## FINAL EXERCISE — Apply Everything

```js
// Given this text:
const text = `
USD12.50 EUR8.99 hello world
email: user@example.com
IP: 192.168.1.1 and 999.0.0.0
action caution hello nation
`;

// Challenge 1: Extract all prices (3-letter code + number.decimal)
text.match(/[A-Z]{3}\d+\.\d+/g)
// ['USD12.50', 'EUR8.99']

// Challenge 2: Extract words containing 'tion' — remove 'ion'
// (words like action, caution, nation)
text.match(/\b\w+(?=tion\b)/g)
// ['ac', 'cau', 'na']... need to include the 'tion' part in the word:
text.match(/\b(?<=t)ion\b/g)   // just the 'ion' parts after 't'

// Challenge 3: Validate the IP (only the valid one)
text.match(/\b(?:(?:25[0-5]|2[0-4]\d|[01]?\d{1,2})\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d{1,2})\b/g)
// ['192.168.1.1'] — 999.0.0.0 excluded
```

---

## 🏆 Complete Reference Card

```
ANCHORS          QUANTIFIERS        FLAGS
^   start        *    0 or more     g  global
$   end          +    1 or more     i  case-insensitive
\b  word edge    ?    0 or 1        m  multiline
                 {n}  exactly n
SETS             {n,} at least n
[abc]  any of   {n,m} n to m
[^abc] none of  *?   lazy
[a-z]  range    +?   lazy

SHORTHANDS       GROUPS             LOOKAROUNDS
\d  digit        (x)   capture      (?=x)  followed by
\D  non-digit    (?:x) no capture   (?!x)  not followed
\w  word char    \1    backref      (?<=x) preceded by
\W  non-word                        (?<!x) not preceded
\s  whitespace
\S  non-whitespace
.   any (no \n)

METHODS
.test(str)          → boolean
str.match(re)       → array | null
str.replace(re, x)  → string
re.exec(str)        → array | null (loop)
str.matchAll(re)    → iterator
str.split(re)       → array
```

**You now know everything you need. Do the quizzes! 💪🔥**
