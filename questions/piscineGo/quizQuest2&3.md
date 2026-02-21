# 🎓 Piscine Go - Mastery Quiz
## Quest 02-03 Final Assessment

**Time Limit:** 45 minutes  
**Passing Score:** 28/35 (80%)  
**Instructions:** No code execution allowed - predict output mentally!

---

## 📋 SECTION 1: OUTPUT PREDICTION (10 Questions)

### Q1: What's the output?
```go
package main
import "github.com/01-edu/z01"

func main() {
    for i := '0'; i <= '2'; i++ {
        z01.PrintRune(i)
    }
    z01.PrintRune('\n')
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `012`

**Explanation:** Loop from '0' to '2' printing each digit character.
</details>

---

### Q2: What's the output?
```go
func main() {
    x := 5
    p := &x
    *p = 10
    fmt.Println(x)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `10`

**Explanation:** Pointer `p` points to `x`, modifying `*p` changes `x`.
</details>

---

### Q3: What's the output?
```go
func Test(n int) {
    n = 20
}

func main() {
    x := 10
    Test(x)
    fmt.Println(x)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `10`

**Explanation:** Go passes by value - `Test` gets a copy of `x`.
</details>

---

### Q4: What's the output?
```go
func Swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    x, y := 3, 7
    Swap(&x, &y)
    fmt.Println(x, y)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `7 3`

**Explanation:** Swap exchanges values through pointers.
</details>

---

### Q5: What's the output?
```go
func main() {
    s := "Go"
    for i := 0; i < len(s); i++ {
        fmt.Print(string(s[i]))
    }
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `Go`

**Explanation:** Iterates through string, printing each character.
</details>

---

### Q6: What's the output?
```go
func StrLen(s string) int {
    count := 0
    for range s {
        count++
    }
    return count
}

func main() {
    fmt.Println(StrLen("Hello"))
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `5`

**Explanation:** Counts each character in "Hello".
</details>

---

### Q7: What's the output?
```go
func DivMod(a, b int, div, mod *int) {
    *div = a / b
    *mod = a % b
}

func main() {
    var d, m int
    DivMod(17, 5, &d, &m)
    fmt.Println(d, m)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `3 2`

**Explanation:** 17 / 5 = 3, 17 % 5 = 2
</details>

---

### Q8: What's the output?
```go
func main() {
    for i := 0; i < 3; i++ {
        for j := i + 1; j < 4; j++ {
            fmt.Print(i, j, " ")
        }
    }
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `01 02 03 12 13 23 `

**Explanation:** Nested loop prints all combinations where i < j.
</details>

---

### Q9: What's the output?
```go
func StrRev(s string) string {
    bytes := []byte(s)
    for i := 0; i < len(bytes)/2; i++ {
        j := len(bytes) - 1 - i
        bytes[i], bytes[j] = bytes[j], bytes[i]
    }
    return string(bytes)
}

func main() {
    fmt.Println(StrRev("abc"))
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `cba`

**Explanation:** Two-pointer swap reverses the string.
</details>

---

### Q10: What's the output?
```go
func BasicAtoi(s string) int {
    result := 0
    for _, ch := range s {
        result = result*10 + int(ch-'0')
    }
    return result
}

func main() {
    fmt.Println(BasicAtoi("123"))
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `123`

**Explanation:** Converts string digits to integer: 0*10+1=1, 1*10+2=12, 12*10+3=123
</details>

---

## 📋 SECTION 2: CODE COMPLETION (10 Questions)

### Q11: Complete PrintAlphabet
```go
func PrintAlphabet() {
    for ch := ___; ch <= ___; ch++ {
        z01.PrintRune(ch)
    }
    z01.PrintRune('\n')
}
```

<details><summary>💡 Solution</summary>
```go
for ch := 'a'; ch <= 'z'; ch++ {
```
</details>

---

### Q12: Complete PointOne
```go
func PointOne(n ___) {
    ___ = 1
}
```

<details><summary>💡 Solution</summary>
```go
func PointOne(n *int) {
    *n = 1
}
```
</details>

---

### Q13: Complete UltimatePointOne
```go
func UltimatePointOne(n ______) {
    ______ = 1
}
```

<details><summary>💡 Solution</summary>
```go
func UltimatePointOne(n ***int) {
    ***n = 1
}
```
</details>

---

### Q14: Complete Swap
```go
func Swap(a, b *int) {
    ___, ___ = ___, ___
}
```

<details><summary>💡 Solution</summary>
```go
*a, *b = *b, *a
```
</details>

---

### Q15: Complete UltimateDivMod
```go
func UltimateDivMod(a, b *int) {
    div := ___
    mod := ___
    *a = ___
    *b = ___
}
```

<details><summary>💡 Solution</summary>
```go
div := *a / *b
mod := *a % *b
*a = div
*b = mod
```
</details>

---

### Q16: Complete PrintStr
```go
func PrintStr(s string) {
    for ___ := range s {
        z01.PrintRune(___)
    }
}
```

<details><summary>💡 Solution</summary>
```go
for _, ch := range s {
    z01.PrintRune(ch)
}
```
</details>

---

### Q17: Complete IsNegative
```go
func IsNegative(n int) {
    if ___ {
        z01.PrintRune('T')
    } else {
        z01.PrintRune('F')
    }
    z01.PrintRune('\n')
}
```

<details><summary>💡 Solution</summary>
```go
if n < 0 {
```
</details>

---

### Q18: Complete Atoi with sign handling
```go
func Atoi(s string) int {
    sign := 1
    start := 0
    
    if s[0] == '-' {
        sign = ___
        start = ___
    }
    
    result := 0
    for i := start; i < len(s); i++ {
        result = result*___ + int(s[i]-___)
    }
    return result * ___
}
```

<details><summary>💡 Solution</summary>
```go
sign = -1
start = 1
result = result*10 + int(s[i]-'0')
return result * sign
```
</details>

---

### Q19: Complete PrintComb (3 digits)
```go
func PrintComb() {
    for i := '0'; i <= ___; i++ {
        for j := ___; j <= ___; j++ {
            for k := ___; k <= ___; k++ {
                z01.PrintRune(i)
                z01.PrintRune(j)
                z01.PrintRune(k)
                if !(i == ___ && j == ___ && k == ___) {
                    z01.PrintRune(',')
                    z01.PrintRune(' ')
                }
            }
        }
    }
}
```

<details><summary>💡 Solution</summary>
```go
for i := '0'; i <= '7'; i++ {
    for j := i + 1; j <= '8'; j++ {
        for k := j + 1; k <= '9'; k++ {
            // ...
            if !(i == '7' && j == '8' && k == '9') {
```
</details>

---

### Q20: Complete SortIntegerTable
```go
func SortIntegerTable(table []int) {
    for i := 0; i < len(table); i++ {
        for j := ___; j < len(table); j++ {
            if table[i] ___ table[j] {
                table[i], table[j] = ___, ___
            }
        }
    }
}
```

<details><summary>💡 Solution</summary>
```go
for j := i + 1; j < len(table); j++ {
    if table[i] > table[j] {
        table[i], table[j] = table[j], table[i]
```
</details>

---

## 📋 SECTION 3: BUG FIXING (5 Questions)

### Q21: Fix this code
```go
// BUG: Doesn't reverse!
func StrRev(s string) string {
    for i := 0; i < len(s); i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
    return s
}
```

<details><summary>💡 Solution</summary>
**Bug:** Strings are immutable in Go!

**Fix:**
```go
func StrRev(s string) string {
    bytes := []byte(s)  // Convert to mutable slice
    for i := 0; i < len(bytes)/2; i++ {
        bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
    }
    return string(bytes)
}
```
</details>

---

### Q22: Fix this code
```go
// BUG: Doesn't modify x!
func SetToTen(n int) {
    n = 10
}

func main() {
    x := 5
    SetToTen(x)
    fmt.Println(x)  // Still 5!
}
```

<details><summary>💡 Solution</summary>
**Bug:** Passing by value!

**Fix:**
```go
func SetToTen(n *int) {
    *n = 10
}

func main() {
    x := 5
    SetToTen(&x)
    fmt.Println(x)  // Now 10!
}
```
</details>

---

### Q23: Fix this code
```go
// BUG: Returns wrong value!
func Atoi(s string) int {
    result := 0
    for _, ch := range s {
        result = result*10 + int(ch)  // Wrong!
    }
    return result
}
```

<details><summary>💡 Solution</summary>
**Bug:** Not converting character to digit!

**Fix:**
```go
result = result*10 + int(ch - '0')
// or
result = result*10 + int(ch) - int('0')
```
</details>

---

### Q24: Fix this code
```go
// BUG: Infinite loop!
func PrintDigits() {
    for i := '0'; i <= '9'; {
        z01.PrintRune(i)
    }
}
```

<details><summary>💡 Solution</summary>
**Bug:** Missing increment!

**Fix:**
```go
for i := '0'; i <= '9'; i++ {
    z01.PrintRune(i)
}
```
</details>

---

### Q25: Fix this code
```go
// BUG: Compile error!
func main() {
    s := "hello"
    fmt.Println(s[0])  // Want to print 'h'
}
```

<details><summary>💡 Solution</summary>
**Bug:** `s[0]` is a byte (uint8), prints as number!

**Fix:**
```go
fmt.Println(string(s[0]))  // Convert to string
// or
z01.PrintRune(rune(s[0]))  // Convert to rune
```
</details>

---

## 📋 SECTION 4: CONCEPT QUESTIONS (10 Questions)

### Q26: Multiple Choice
What does `&x` return?
- A) Value of x
- B) Address of x
- C) Type of x
- D) Copy of x

<details><summary>💡 Solution</summary>
**B) Address of x**

`&` is the address-of operator.
</details>

---

### Q27: Multiple Choice
What does `*p` do when p is a pointer?
- A) Gets address
- B) Gets/sets value at address
- C) Creates pointer
- D) Deletes pointer

<details><summary>💡 Solution</summary>
**B) Gets/sets value at address**

`*` dereferences the pointer.
</details>

---

### Q28: Multiple Choice
Can you modify a string directly in Go?
- A) Yes, always
- B) No, strings are immutable
- C) Only with pointers
- D) Only uppercase letters

<details><summary>💡 Solution</summary>
**B) No, strings are immutable**

Must convert to `[]byte` to modify.
</details>

---

### Q29: Multiple Choice
What's the ASCII value of '0'?
- A) 0
- B) 48
- C) 65
- D) 97

<details><summary>💡 Solution</summary>
**B) 48**

'0' = 48, 'A' = 65, 'a' = 97
</details>

---

### Q30: Multiple Choice
How to convert lowercase 'a' to uppercase 'A'?
- A) ch + 32
- B) ch - 32
- C) ch * 32
- D) ch / 32

<details><summary>💡 Solution</summary>
**B) ch - 32**

'a' (97) - 32 = 'A' (65)
</details>

---

### Q31: True/False
Passing a slice to a function allows modifying its elements.

<details><summary>💡 Solution</summary>
**TRUE**

Slices are reference types - modifications affect original.
</details>

---

### Q32: True/False
Passing an array to a function allows modifying its elements.

<details><summary>💡 Solution</summary>
**FALSE**

Arrays are value types - function gets a copy.
</details>

---

### Q33: Fill in the blank
To convert string "123" to integer, the formula for each digit is:
`result = result * ___ + int(ch - ___)`

<details><summary>💡 Solution</summary>
```go
result = result * 10 + int(ch - '0')
```
</details>

---

### Q34: Multiple Choice
What's the difference between `var s string` and `s := ""`?
- A) No difference
- B) := can only be used in functions
- C) var is faster
- D) := requires type annotation

<details><summary>💡 Solution</summary>
**B) := can only be used in functions**

Short declaration `:=` only works inside functions.
</details>

---

### Q35: Multiple Choice
What happens if you dereference a nil pointer?
- A) Returns 0
- B) Returns nil
- C) Panic (crash)
- D) Compile error

<details><summary>💡 Solution</summary>
**C) Panic (crash)**

Dereferencing nil pointer causes runtime panic.
</details>

---

## 🎯 SCORING GUIDE

**Count your correct answers:**

### 35/35 - 🏆 PERFECT MASTER
- Ready for advanced piscine immediately!
- Consider helping others

### 32-34 - 🔥 EXCELLENT
- Very strong understanding
- Ready to move forward
- Review missed concepts

### 28-31 - ✅ PASS
- Good grasp of fundamentals
- Ready for next level
- Practice weak areas

### 24-27 - ⚠️ BORDERLINE
- Need more practice
- Review Quest 02-03 again
- Try more exercises before advancing

### Below 24 - 🔄 NEED REVIEW
- Re-study Quest 02 & 03
- Practice all exercises again
- Take quiz again in 2-3 days

---

## 📊 TOPIC BREAKDOWN

**Check which areas need work:**

**Loops & Printing (Q1, Q8, Q11, Q19, Q24):**
- ☐ Master level (5/5)
- ☐ Good (4/5)
- ☐ Needs practice (3 or less)

**Pointers (Q2, Q4, Q7, Q12-15, Q21-22, Q26-27, Q35):**
- ☐ Master level (10/11)
- ☐ Good (8-9/11)
- ☐ Needs practice (7 or less)

**Strings (Q5-6, Q9-10, Q16, Q18, Q21, Q23, Q28, Q33):**
- ☐ Master level (9/10)
- ☐ Good (7-8/10)
- ☐ Needs practice (6 or less)

**Basic Concepts (Q3, Q17, Q29-32, Q34):**
- ☐ Master level (7/7)
- ☐ Good (5-6/7)
- ☐ Needs practice (4 or less)

**Code Quality (Q20, Q25):**
- ☐ Master level (2/2)
- ☐ Needs practice (less than 2)

---

## 💡 STUDY RECOMMENDATIONS

**Based on your score:**

### If you scored 28-35:
✅ **Move forward!**
- You're ready for next challenges
- Help teammates who scored lower
- Push yourself with harder problems

### If you scored 24-27:
📚 **Light review needed**
- Re-read problem descriptions for missed questions
- Practice similar exercises
- Take quiz again tomorrow

### If you scored below 24:
🔄 **Solid review required**
- Go through Quest 02-03 material again
- Practice each type of problem
- Work with a partner
- Retake in 3-5 days

---

## 🎓 NEXT STEPS

**After passing (28+):**

1. ✅ **Quest 04:** Functions & Recursion
2. ✅ **Quest 05:** Advanced String Operations
3. ✅ **More complex algorithms**

**Keep coding! Practice makes perfect! 💪**

---

## 📝 ANSWER SHEET

**Section 1 (Output Prediction):**
1. _____ 2. _____ 3. _____ 4. _____ 5. _____
6. _____ 7. _____ 8. _____ 9. _____ 10. _____

**Section 2 (Code Completion):**
11-20: Check your code against solutions

**Section 3 (Bug Fixing):**
21-25: Check your fixes

**Section 4 (Concepts):**
26. _____ 27. _____ 28. _____ 29. _____ 30. _____
31. _____ 32. _____ 33. _____ 34. _____ 35. _____

**Total Score: _____/35**

**Ready to move forward? _____ (Yes/No)**