# 🎉 NextEra Fun Day — GoLang Coding Challenge
### 4-Part Team Competition

---

## 📌 Part 1: Individual Speed Debug

> **Goal:** Each student solves one question. The fastest students become **Team Leaders**.

---

### Question 1: Infinite Loop

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; {
        fmt.Println(i)
    }
}
```

❓ **What is the bug?**

<details>
<summary>Answer</summary>

❌ **Bug:** `i++` is missing inside the loop body — this causes an **infinite loop**!

✅ **Fix:** Add `i++` inside the loop:
```go
for i := 0; i < 5; {
    fmt.Println(i)
    i++
}
```
</details>

---

### Question 2: Nil Map Assignment

```go
package main

import "fmt"

func main() {
    var m map[string]int
    m["a"] = 1
    fmt.Println(m)
}
```

❓ **What is the bug?**

<details>
<summary>Answer</summary>

❌ **Bug:** The map is declared but **never initialized** — causes a panic: `assignment to entry in nil map`

✅ **Fix:** Initialize the map with `make`:
```go
m := make(map[string]int)
```
</details>

---

### Question 3: Index Out of Range

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    fmt.Println(s[3])
}
```

❓ **What is the bug?**

<details>
<summary>Answer</summary>

❌ **Bug:** Index `3` doesn't exist — the slice has indices `0, 1, 2` only.

✅ **Fix:** Use a valid index:
```go
fmt.Println(s[2]) // prints 3
```
</details>

---

### Question 4: Missing Return Statement

```go
package main

import "fmt"

func add(a, b int) int {
    sum := a + b
}

func main() {
    result := add(3, 5)
    fmt.Println(result)
}
```

❓ **What is the bug?**

<details>
<summary>Answer</summary>

❌ **Bug:** The function declares a return type of `int` but has **no return statement**.

✅ **Fix:** Add the return:
```go
func add(a, b int) int {
    sum := a + b
    return sum
}
```
</details>

---

### Question 5: Array vs Slice (Append)

```go
package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    arr = append(arr, 4)
    fmt.Println(arr)
}
```

❓ **What is the bug?**

<details>
<summary>Answer</summary>

❌ **Bug:** `append` only works on **slices**, not arrays. `[3]int` is a fixed-size array.

✅ **Fix:** Declare a slice instead:
```go
arr := []int{1, 2, 3}
arr = append(arr, 4)
```
</details>

---

## 📌 Part 2: Team vs Team Battle

> **Idea:** Each team prepares 3 questions to challenge another team.

### Rules

1. Each team prepares **3 questions** (debugging, syntax, or concepts)
2. Team A asks their first question to Team B
3. Team B has **2 minutes** to answer
4. ✅ Correct answer → **1 point for the answering team**
5. ❌ Wrong answer → **1 point for the asking team**
6. The team with the **most points wins**

### Question Type Examples

| Type | Example |
|------|---------|
| 🐛 Debugging | "There's a bug in this code — what is it?" |
| 🔤 Syntax | "What's the difference between `var` and `:=`?" |
| 📤 Output | "What will this code print?" |
| 💡 Concepts | "What's the difference between an array and a slice?" |

> 💡 **Tip:** Harder questions give your team a better chance at earning points when the other team gets it wrong!

---

## 📌 Part 3: Speed Questions — First to Ring Wins!

> **Idea:** The first team to ring in with a correct answer wins the point.

---

**Q1:** What is the **zero value** of an `int` in Go?

<details>
<summary>Answer</summary>

**`0`**
</details>

---

**Q2:** What does `:=` do in Go?

<details>
<summary>Answer</summary>

**Short variable declaration** — it declares and assigns a variable in one step.
</details>

---

**Q3:** What types of loops exist in Go, and what's the difference between them?

<details>
<summary>Answer</summary>

Go only has **`for` loops**. There are 3 styles:

```go
for i := 0; i < 10; i++ {}  // classic loop
for condition {}              // while-style loop
for {}                        // infinite loop
```
</details>

---

**Q4:** What does **recursion** do in Go?

<details>
<summary>Answer</summary>

A function that **calls itself** to solve a problem by breaking it into smaller sub-problems.
</details>

---

**Q5:** What is the difference between an **array** and a **slice** in Go?

<details>
<summary>Answer</summary>

| | Array | Slice |
|--|-------|-------|
| Size | Fixed | Dynamic |
| Declaration | `[3]int` | `[]int` |
</details>

---

**Q6:** What does `defer` do in Go?

<details>
<summary>Answer</summary>

**Delays** the execution of a function until the surrounding function returns. Commonly used for **cleanup** (e.g., closing files).
</details>

---

## 📌 Part 4: Build From Scratch

> ⏱ **Duration:** 30 minutes

### Scoring Rubric

| Criterion | Weight | Description |
|-----------|--------|-------------|
| ✅ Functionality | 50% | Code runs without errors and produces correct output |
| 🧹 Code Quality | 30% | Clean, readable code with proper naming and structure |
| 💡 Creativity | 20% | Creative approach or elegant solution |

---

### Problem 1: Print Odd Numbers

**Task:** Write a function that prints numbers from 1 to 10, skipping even numbers.

```go
func main() {
    // Your code here
}
```

---

### Problem 2: Sum of a Slice

**Task:** Write a function that takes a slice of integers and returns their sum.

```go
func sum(nums []int) int {
    // Your code here
}
```

---

### Problem 3: Palindrome Checker

**Task:** Write a function that takes a string and returns `true` if it's a palindrome, `false` otherwise.

```go
func isPalindrome(s string) bool {
    // Your code here
}
```

---

## 💬 Reflection Questions

*Take a moment to think — share with your team or the group!*

1. What's the most valuable thing you learned at NextEra that you couldn't have learned anywhere else?

2. How did the **peer-to-peer** learning system change things for you — is it really different?

3. How did **the Piscine** change your view of programming? What did you expect, and what did you find?

4. Was there a moment you felt like giving up? What made you keep going?

5. What's the longest bug you ever struggled with — and how did it feel when it finally worked?

6. What was the moment you said to yourself: *"I actually know how to code now"*?

---

*Good luck, have fun, and may the best team win! 🚀*
