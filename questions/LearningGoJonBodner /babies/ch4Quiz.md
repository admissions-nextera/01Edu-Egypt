# 📘 Learning Go — Chapter 4 Quiz
## Blocks, Shadows, and Control Structures

**Time Limit:** 35 minutes  
**Total Questions:** 24  
**Passing Score:** 19/24 (79%)

> This quiz covers: variable shadowing, `if`/`else`, all forms of `for`, `switch`, `goto`, and the blank identifier.

---

## 📋 SECTION 1: BLOCKS AND SHADOWING (5 Questions)

### Q1: What is variable shadowing?

**A)** Declaring a variable with type `shadow`  
**B)** Declaring a variable with the same name in an inner scope — the inner variable hides the outer one within its block  
**C)** Using a variable before declaring it  
**D)** An error Go prevents at compile time  

<details><summary>💡 Answer</summary>

**B) Inner declaration hides the outer variable within its scope**

```go
x := 10
if true {
    x := 20         // shadows outer x — new variable!
    fmt.Println(x)  // 20
}
fmt.Println(x)      // 10 — outer x unchanged
```

Shadowing is a common, hard-to-spot bug. The book highlights this because `:=` inside `if` or `for` blocks can accidentally create a new variable instead of assigning to the outer one. Use `go vet` or `shadow` linter to catch unintentional shadowing.

</details>

---

### Q2: What is the "universe block" in Go?

**A)** The block inside `package main`  
**B)** The outermost scope containing all predeclared identifiers: `true`, `false`, `nil`, `len`, `cap`, `make`, `new`, `append`, `int`, `string`, etc.  
**C)** The block containing imported packages  
**D)** A special block for global variables  

<details><summary>💡 Answer</summary>

**B) The outermost scope with all built-in identifiers**

This means you CAN shadow built-in identifiers:
```go
true := false          // shadows universe-block `true`
fmt.Println(true)      // false — dangerous!

len := 5               // shadows the built-in len function
```

The Go compiler allows this (the universe block is shadowed by package block). It's extremely bad practice. The `shadow` linter catches it.

</details>

---

### Q3: What does this code print?
```go
x := 1
if x > 0 {
    x := x + 1   // note: := not =
    fmt.Println(x)
}
fmt.Println(x)
```

**A)** `2` then `2`  
**B)** `2` then `1` — inner `:=` creates a new `x`; outer `x` is unchanged  
**C)** Compile error  
**D)** `1` then `1`  

<details><summary>💡 Answer</summary>

**B) `2` then `1`**

`x := x + 1` inside the `if` block:
1. Evaluates `x + 1` using the outer `x` (1 + 1 = 2)
2. Creates a NEW inner `x` via `:=`, assigns it `2`
3. The outer `x` is still `1` after the block

This is the canonical shadowing bug. If you intended to modify the outer `x`, use `x = x + 1` (assignment, not declaration).

</details>

---

### Q4: At what scope level do `if`, `for`, and `switch` blocks create their own scope?

**A)** They don't — they share the surrounding function's scope  
**B)** They create their own block scope — variables declared inside are not visible outside  
**C)** Only `for` creates a scope; `if` and `switch` do not  
**D)** Only named blocks create scopes  

<details><summary>💡 Answer</summary>

**B) All three create their own block scope**

```go
if true {
    x := 10       // x exists only here
}
fmt.Println(x)    // compile error: x undefined

for i := 0; i < 3; i++ {
    // i exists only in this loop
}
fmt.Println(i)    // compile error: i undefined
```

`for` has two scopes: the init statement scope (contains `i`) and the loop body scope.

</details>

---

### Q5: What is the `if` init statement and when is it useful?

**A)** A special first statement that runs before the `if` condition, scoped to the `if`/`else` block  
**B)** The first `if` in a chain  
**C)** An `if` without an `else`  
**D)** Initializing variables before the function containing the `if`  

<details><summary>💡 Answer</summary>

**A) A statement scoped to the `if`/`else` block**

```go
if err := doSomething(); err != nil {
    fmt.Println("error:", err)
    return
}
// err is not in scope here

if val, ok := m["key"]; ok {
    fmt.Println("found:", val)
}
// val and ok are not in scope here
```

This is idiomatic Go for error handling and map lookups — the variable's scope is exactly as wide as needed.

</details>

---

## 📋 SECTION 2: FOR LOOPS (6 Questions)

### Q6: Go has only one looping keyword. What are its four usage forms?

**A)** `for`, `while`, `do-while`, `foreach`  
**B)** Complete `for` (C-style), condition-only (`while`-style), infinite `for`, `for-range`  
**C)** `for`, `for-each`, `for-ever`, `for-range`  
**D)** Only one form — `for i := 0; i < n; i++`  

<details><summary>💡 Answer</summary>

**B) Four forms of `for`**

```go
// 1. Complete (C-style):
for i := 0; i < 10; i++ { }

// 2. Condition-only (while-style):
for x > 0 { x-- }

// 3. Infinite:
for { }  // use break to exit

// 4. For-range:
for i, v := range slice { }
for k, v := range m { }
for i, r := range str { }  // i=byte index, r=rune
```

Go replaced `while` and `do-while` with these `for` variants.

</details>

---

### Q7: What is the output?
```go
for i := 0; i < 3; i++ {
    if i == 1 {
        continue
    }
    fmt.Println(i)
}
```

**A)** `0 1 2`  
**B)** `0 2`  
**C)** `0 1`  
**D)** `1 2`  

<details><summary>💡 Answer</summary>

**B) `0` then `2`**

`continue` skips the rest of the current loop body and goes to the next iteration. When `i == 1`, `fmt.Println(1)` is skipped. The loop continues with `i = 2`.

</details>

---

### Q8: How do you iterate over the runes (Unicode characters) of a string?

**A)** `for i := 0; i < len(s); i++ { c := s[i] }` — iterates bytes  
**B)** `for i, r := range s { }` — `i` is the byte index of the rune's start, `r` is the rune value  
**C)** `for _, r := range []rune(s) { }` — also valid but converts first  
**D)** Both B and C work; B is more efficient for ASCII-heavy strings  

<details><summary>💡 Answer</summary>

**D) Both B and C work — B is preferred**

```go
s := "héllo"
for i, r := range s {
    fmt.Printf("byte %d: %c\n", i, r)
}
// byte 0: h
// byte 1: é  (occupies bytes 1-2 in UTF-8)
// byte 3: l
// byte 4: l
// byte 5: o
```

Option A iterates bytes — for multi-byte runes like 'é', you'd get individual bytes (garbage characters). Option B handles UTF-8 correctly. The byte index `i` can skip by 2+ for multi-byte characters.

</details>

---

### Q9: What does `break` with a label do?

**A)** Breaks out of the labeled function  
**B)** Breaks out of the outer loop or switch identified by the label, not just the innermost one  
**C)** A syntax error  
**D)** Breaks and jumps to the label  

<details><summary>💡 Answer</summary>

**B) Breaks out of the labeled construct (outer loop/switch)**

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer  // exits the OUTER loop
        }
    }
}
```

Without the label, `break` only exits the inner loop. Labels are the idiomatic Go way to break out of nested loops — an alternative to using a boolean flag.

</details>

---

### Q10: Is there a `do-while` loop in Go?

**A)** Yes — `do { } while (condition)`  
**B)** No — simulate with an infinite `for` and `break`:
```go
for {
    // body
    if !condition { break }
}
```
**C)** Yes — `for ; ; { }` with a trailing condition  
**D)** Use `repeat`  

<details><summary>💡 Answer</summary>

**B) No `do-while` — simulate with infinite `for` + `break`**

```go
for {
    input := getInput()
    process(input)
    if input == "quit" {
        break
    }
}
```

The book explicitly covers this pattern. Go has no `do-while` or `until` keyword.

</details>

---

### Q11: What is the output?
```go
s := []int{1, 2, 3}
for _, v := range s {
    s = append(s, v*10)
}
fmt.Println(s)
```

**A)** `[1 2 3 10 20 30]` — appends as it goes  
**B)** `[1 2 3 10 20 30]` — but might infinite-loop  
**C)** `[1 2 3 10 20 30]` — `range` evaluates the slice header once at start; the loop iterates exactly 3 times  
**D)** Panic — cannot modify a slice during range iteration  

<details><summary>💡 Answer</summary>

**C) `[1 2 3 10 20 30]` — `range` captures slice header at loop start**

`range` evaluates the slice (captures len=3) at the beginning. Even though `append` grows `s`, the loop iterates exactly the original 3 times. This is safe and produces the expected output. For maps, elements added during range iteration may or may not be seen.

</details>

---

## 📋 SECTION 3: SWITCH (7 Questions)

### Q12: Does a Go `switch` case require a `break` statement?

**A)** Yes — same as C  
**B)** No — Go cases do NOT fall through by default; `break` is implicit. Use `fallthrough` to explicitly fall through.  
**C)** It depends on the type  
**D)** `break` is required after the last case only  

<details><summary>💡 Answer</summary>

**B) No — cases don't fall through by default; `break` is implicit**

```go
switch x {
case 1:
    fmt.Println("one")
    // implicit break — no fall-through to case 2
case 2:
    fmt.Println("two")
}

// To fall through explicitly:
switch x {
case 1:
    fmt.Println("one")
    fallthrough    // falls into case 2
case 2:
    fmt.Println("two")
}
```

This is the opposite of C/Java, where fall-through is the default and `break` is required.

</details>

---

### Q13: What is the output?
```go
x := 2
switch x {
case 1, 2:
    fmt.Println("one or two")
case 2, 3:
    fmt.Println("two or three")
}
```

**A)** `one or two` then `two or three` — both cases match  
**B)** `one or two` — first matching case wins; no fall-through  
**C)** Compile error — `2` appears in two cases  
**D)** `two or three`  

<details><summary>💡 Answer</summary>

**B) `one or two` — first match wins**

Go switch evaluates cases top-to-bottom. The first case where any value matches is executed, then the switch exits. Duplicate values in different cases are actually a compile error in `switch` statements in newer Go versions, but the key concept is: no fall-through.

</details>

---

### Q14: How do you write a `switch` without a condition?

**A)** `switch { }` — replaces `if/else if` chains; each case has a boolean expression  
**B)** That's not valid Go  
**C)** `switch nil { }`  
**D)** `switch true { }`  

<details><summary>💡 Answer</summary>

**A and D are equivalent — `switch {}` defaults to comparing against `true`**

```go
// switch without condition — equivalent to switch true:
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive")
}
```

This is idiomatic Go for complex boolean conditions that would be messy as `if/else if`. Each case is an independent boolean expression.

</details>

---

### Q15: What is the `switch` init statement?

**A)** A `case` that initializes the switch variable  
**B)** A statement before the switch expression that creates a variable scoped to the switch block  
**C)** An error  
**D)** The `default` case  

<details><summary>💡 Answer</summary>

**B) A statement scoped to the switch block**

```go
switch x := getValue(); {
case x > 100:
    fmt.Println("large")
case x > 10:
    fmt.Println("medium")
default:
    fmt.Println("small")
}
// x not accessible here
```

Same pattern as `if` init statement — the variable is scoped to exactly where it's needed.

</details>

---

### Q16: What does this print?
```go
x := 1
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}
```

**A)** `one`  
**B)** `one` then `two` — `fallthrough` executes the next case's body regardless of its condition  
**C)** `one` then `two` then `three`  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) `one` then `two` — `fallthrough` executes the next body unconditionally**

`fallthrough` bypasses the next case's condition and runs its body directly. It only falls through one level — it does NOT cascade to `case 3`. `fallthrough` must be the last statement in a case body and cannot be used in the final case.

</details>

---

### Q17: When should you use `goto` in Go?

**A)** Never — `goto` causes compile errors  
**B)** Rarely — `goto` jumps to a labeled statement within the same function; most legitimate uses are breaking out of nested loops (though labeled `break` is preferred) or cleanup in error paths  
**C)** Whenever you need to jump between functions  
**D)** As an alternative to `return`  

<details><summary>💡 Answer</summary>

**B) Rarely — labeled `break` is almost always preferred**

`goto` exists in Go but is discouraged. Constraints: can only jump within a function, cannot jump over variable declarations, cannot jump into a block. The book mentions it exists but suggests using labeled `break`/`continue` instead. The one case `goto` is useful: cleanup in deeply nested error paths in generated or very low-level code.

</details>

---

### Q18: What is a type switch?

**A)** A `switch` on the type of a variable, used to handle different types stored in an interface  
**B)** A switch that only works with type assertions  
**C)** An error in Go  
**D)** A `switch` using reflect  

<details><summary>💡 Answer</summary>

**A) Switch on the dynamic type of an interface value**

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    default:
        fmt.Printf("unknown type: %T\n", v)
    }
}
```

Inside each case, `v` has the concrete type of that case. Type switches are how Go code inspects what's inside an interface value.

</details>

---

## 📋 SECTION 4: INTEGRATION (6 Questions)

### Q19: What does this print?
```go
for i := 0; i < 5; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Print(i, " ")
}
```

**A)** `0 2 4`  
**B)** `1 3`  
**C)** `0 1 2 3 4`  
**D)** `2 4`  

<details><summary>💡 Answer</summary>

**B) `1 3`**

`continue` skips even numbers (`i%2 == 0` is true for 0, 2, 4). Only odd numbers (1, 3) reach `fmt.Print`.

</details>

---

### Q20: Can `if` and `switch` have an `else` and `default` clause respectively if those blocks always return or panic?

**A)** No — `else` and `default` are required for completeness  
**B)** Yes — if all other paths return, the final clause is implicitly the remaining path; the Go compiler understands this for control flow  
**C)** Only for `switch`  
**D)** Only for `if`  

<details><summary>💡 Answer</summary>

**B) Yes — Go's compiler tracks control flow; an explicit `else`/`default` is optional when other paths return**

```go
func sign(x int) string {
    if x > 0 {
        return "positive"
    } else if x < 0 {
        return "negative"
    }
    return "zero"   // no else needed — compiler knows we reach here only if both if conditions were false
}
```

This is a matter of style — the book shows both patterns.

</details>

---

### Q21: What is wrong with this shadowing pattern?
```go
err := doFirst()
if err != nil {
    return err
}
result, err := doSecond()  // err is reassigned here, ok?
```

**A)** Compile error — `err` already declared  
**B)** Nothing — this is correct; `:=` is valid here because `result` is new  
**C)** `err` is shadowed — the second `err` is a different variable  
**D)** `result` must be declared before `err`  

<details><summary>💡 Answer</summary>

**B) Correct — `:=` reassigns `err` because `result` is new**

When `:=` is used and at least one variable on the left is new, existing variables are reassigned (not redeclared). `result` is new, so `:=` is valid. `err` is reassigned to the new error value. This is the idiomatic Go error handling pattern.

</details>

---

### Q22: What does `for range` do with a channel?

**A)** Not supported  
**B)** Iterates over values received from the channel until the channel is closed  
**C)** Iterates a fixed number of times  
**D)** `range` works on channels only with the index form  

<details><summary>💡 Answer</summary>

**B) Iterates over channel values until the channel is closed**

```go
ch := make(chan int)
go func() {
    for i := 0; i < 5; i++ { ch <- i }
    close(ch)
}()

for v := range ch {   // receives 0, 1, 2, 3, 4; then loop exits when ch is closed
    fmt.Println(v)
}
```

This is the idiomatic pattern for consuming all values from a channel. Only one variable (the value) is used with channel range — there is no index.

</details>

---

### Q23: Why does Go not have a ternary operator (`x ? y : z`)?

**A)** It does — use `?:`  
**B)** The Go authors considered it and chose not to include it — it encourages complex expressions that hurt readability; use an `if/else` instead  
**C)** It was removed in Go 1.18  
**D)** Use `select` instead  

<details><summary>💡 Answer</summary>

**B) Deliberate design choice — use `if/else` instead**

```go
// No ternary in Go. Instead:
var max int
if a > b {
    max = a
} else {
    max = b
}

// Or using a function:
func max(a, b int) int {
    if a > b { return a }
    return b
}
```

The book explains this is intentional — Go prioritizes clarity over brevity.

</details>

---

### Q24: What is the `for` loop's `post` statement scope?

**A)** The post statement (`i++`) runs in the surrounding function scope  
**B)** The post statement runs in the same scope as the init statement — variables declared in `init` are visible in `post` and `condition`  
**C)** The post statement creates its own scope  
**D)** There is no post statement — use `continue` instead  

<details><summary>💡 Answer</summary>

**B) Init, condition, and post all share the same scope**

```go
for i := 0; i < 10; i++ {
    // i is in scope here (body scope, inner)
}
// i is NOT in scope here

// Init declares i; condition uses i; post uses i — all same scope:
for i, j := 0, 10; i < j; i, j = i+1, j-1 {
    fmt.Println(i, j)
}
```

The for loop's three clauses share one scope, which is the parent scope of the loop body.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 22–24 ✅ | **Excellent.** Control flow mastered — proceed to Chapter 5. |
| 19–21 ✅ | **Ready.** Review shadowing patterns and `for-range` behavior. |
| 14–18 ⚠️ | **Review first.** Shadowing and `switch` fall-through behavior need more attention. |
| Below 14 ❌ | **Re-read Chapter 4.** These patterns appear constantly in real Go code. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q5 | Shadowing, universe block, `if` init statement, block scopes |
| Q6–Q11 | `for` forms, `continue`, `break` with label, rune iteration, `range` snapshot |
| Q12–Q18 | `switch` no fall-through, blank `switch`, `fallthrough`, `goto`, type switch |
| Q19–Q24 | Shadowing in `:=`, channel `range`, no ternary, `for` scope |