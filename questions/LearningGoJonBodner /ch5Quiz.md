# 📘 Learning Go — Chapter 5 Quiz
## Functions

**Time Limit:** 40 minutes  
**Total Questions:** 25  
**Passing Score:** 20/25 (80%)

> This quiz covers: multiple return values, named returns, variadic functions, `defer`, closures, function types, and higher-order functions.

---

## 📋 SECTION 1: FUNCTION BASICS & MULTIPLE RETURNS (7 Questions)

### Q1: What is the idiomatic Go way to signal that a function can fail?

**A)** Return `-1` or `nil` as an error sentinel  
**B)** Return a second value of type `error` — the last return value by convention  
**C)** Panic  
**D)** Use a global error variable  

<details><summary>💡 Answer</summary>

**B) Return `(result, error)` — the last value is `error`**

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
```

This is the cornerstone of Go error handling. Errors are values, returned explicitly, checked by the caller.

</details>

---

### Q2: Is it valid to ignore all return values of a function?

**A)** No — all return values must be assigned  
**B)** Yes — `doSomething()` discards all return values; use `_` to discard specific values  
**C)** Only for functions that return `error`  
**D)** Only for void functions  

<details><summary>💡 Answer</summary>

**B) Yes — calling without assignment discards all returns**

```go
doSomething()              // discards all return values
result, _ := mayFail()    // discard the error (be careful!)
_, err := getAndIgnore()  // discard the value, keep the error
```

Discarding errors (`_, _ = ...`) is a code smell but syntactically valid. The book emphasizes that ignoring errors is one of Go's biggest anti-patterns.

</details>

---

### Q3: What are named return values?

**A)** Return values with the name of the function  
**B)** Variables declared in the function signature as return values — they are pre-initialized to their zero values and can be returned with a bare `return`  
**C)** Return values that must be named by the caller  
**D)** A syntax error  

<details><summary>💡 Answer</summary>

**B) Named return variables — pre-initialized, usable with bare `return`**

```go
func minMax(nums []int) (min, max int) {
    min, max = nums[0], nums[0]
    for _, v := range nums[1:] {
        if v < min { min = v }
        if v > max { max = v }
    }
    return  // bare return — returns current values of min and max
}
```

The book recommends using named returns sparingly — mainly to document what a function returns. Bare `return` is discouraged in all but the simplest functions as it hides what's being returned.

</details>

---

### Q4: What does a bare `return` do in a function with named return values?

**A)** Returns zero values for all return types  
**B)** Returns the current values of the named return variables  
**C)** Compile error  
**D)** Returns nothing — only valid for void functions  

<details><summary>💡 Answer</summary>

**B) Returns the current values of the named return variables**

```go
func f() (x int) {
    x = 10
    return      // returns 10
}

func g() (x int) {
    x = 10
    x++
    return      // returns 11
}
```

Bare returns can be confusing in longer functions — the reader must track the named variables' current values. The book advises using bare returns only in very short functions.

</details>

---

### Q5: What is the syntax for a variadic function?

**A)** `func f(args []int)`  
**B)** `func f(args ...int)` — the `...` makes the last parameter variadic; callers can pass zero or more `int` values  
**C)** `func f(int...)`  
**D)** `func f(*int)`  

<details><summary>💡 Answer</summary>

**B) `func f(args ...int)` — variadic with `...` before the type**

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums { total += n }
    return total
}

sum(1, 2, 3)        // pass individual values
nums := []int{1, 2, 3}
sum(nums...)        // spread a slice — must add ... to call
```

Inside the function, `nums` is a `[]int`. The `...` spread operator when calling is required — you cannot pass a slice directly to a variadic function without spreading it.

</details>

---

### Q6: What is the output?
```go
func modify(s ...int) {
    s[0] = 99
}

nums := []int{1, 2, 3}
modify(nums...)
fmt.Println(nums)
```

**A)** `[1 2 3]` — the variadic parameter is a copy  
**B)** `[99 2 3]` — the variadic slice shares the underlying array when spread with `...`  
**C)** Compile error  
**D)** Panic  

<details><summary>💡 Answer</summary>

**B) `[99 2 3]` — spreading shares the underlying array**

When you call `f(slice...)`, the variadic parameter and the original slice share the same backing array. This is a subtle but important behavior. If you want isolation, copy the slice first: `modify(append([]int{}, nums...)...)`.

</details>

---

### Q7: Can Go functions return functions?

**A)** No — functions can only return value types  
**B)** Yes — functions are first-class values; they can be passed as arguments, assigned to variables, and returned from functions  
**C)** Only anonymous functions  
**D)** Only if the return type is `interface{}`  

<details><summary>💡 Answer</summary>

**B) Yes — functions are first-class values**

```go
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

add5 := makeAdder(5)
fmt.Println(add5(3))  // 8
```

Functions in Go have types: `func(int) int` is the type of a function taking one `int` and returning one `int`. This enables higher-order functions, closures, and functional patterns.

</details>

---

## 📋 SECTION 2: DEFER (6 Questions)

### Q8: What does `defer` do?

**A)** Delays execution of a function for 1 second  
**B)** Schedules a function call to run when the surrounding function returns — in LIFO (last in, first out) order if multiple defers are present  
**C)** Makes a function call asynchronous  
**D)** Prevents a function from being called more than once  

<details><summary>💡 Answer</summary>

**B) Schedules a call for when the surrounding function returns — LIFO order**

```go
func cleanup() {
    f, _ := os.Open("file.txt")
    defer f.Close()  // guaranteed to run when cleanup() returns

    defer fmt.Println("third")   // runs first (LIFO)
    defer fmt.Println("second")  // runs second
    defer fmt.Println("first")   // runs third
}
// Output: first, second, third
```

`defer` is used for cleanup (close files, release locks, close HTTP responses). LIFO order means nested resources are cleaned up in reverse acquisition order.

</details>

---

### Q9: When are the arguments to a `defer`red function evaluated?

**A)** When the surrounding function returns  
**B)** When the `defer` statement is reached — the arguments are captured immediately  
**C)** At program exit  
**D)** It depends on the argument type  

<details><summary>💡 Answer</summary>

**B) Arguments are evaluated when `defer` is reached, not when it executes**

```go
x := 10
defer fmt.Println(x)  // x is captured as 10 RIGHT NOW
x = 20
// When function returns: prints 10, not 20
```

This catches people off guard. The function reference and arguments are locked in when the `defer` statement executes. Use a closure to capture the "current" value at execution time:
```go
defer func() { fmt.Println(x) }()  // prints 20 (captures x by reference)
```

</details>

---

### Q10: Does `defer` run if a function panics?

**A)** No — panic bypasses `defer`  
**B)** Yes — `defer` runs during panic unwinding, which is how `recover()` can catch panics  
**C)** Only if the defer was set before the panic  
**D)** Only for `recover()` calls  

<details><summary>💡 Answer</summary>

**B) Yes — `defer` runs during panic unwinding**

```go
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    return a / b, nil  // panics if b == 0
}
```

This is essential: `defer` + `recover()` is the pattern for catching panics. Defers run in LIFO order during panic unwinding, before the program (potentially) terminates.

</details>

---

### Q11: What is the output?
```go
func f() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
}
```

**A)** `0 1 2`  
**B)** `2 1 0` — LIFO; i is captured at defer time  
**C)** `0 0 0`  
**D)** `2 2 2`  

<details><summary>💡 Answer</summary>

**B) `2 1 0` — LIFO order, each `i` captured at time of defer**

Three `defer` calls are registered with `i` values 0, 1, 2 respectively (arguments evaluated immediately). When `f()` returns, they execute in LIFO order: defer with `i=2` runs first, then `i=1`, then `i=0`.

</details>

---

### Q12: Can `defer` modify named return values?

**A)** No — deferred functions run after the return value is set  
**B)** Yes — a deferred closure can modify named return variables, which changes the actual returned value  
**C)** Only if using `recover`  
**D)** Only for error returns  

<details><summary>💡 Answer</summary>

**B) Yes — deferred closures can modify named return variables**

```go
func double(x int) (result int) {
    defer func() {
        result *= 2  // modifies the named return variable
    }()
    result = x
    return  // bare return — result is x; then defer doubles it
}
fmt.Println(double(5))  // 10
```

This is an advanced but legitimate pattern — often used to modify the `err` return value in cleanup logic.

</details>

---

### Q13: When does a `defer`red function in `main()` run?

**A)** When `main()` returns normally  
**B)** When `os.Exit()` is called  
**C)** Both A and B  
**D)** Neither — `main`'s defers are a special case  

<details><summary>💡 Answer</summary>

**A) Only when `main()` returns normally — NOT when `os.Exit()` is called**

`os.Exit()` terminates the program immediately — defers DO NOT run. This is one of the reasons to avoid `os.Exit()` in library code. If you need cleanup on exit, ensure it happens before `os.Exit()` is called, or use `defer` only for cleanup paths that don't go through `os.Exit()`. `log.Fatal()` calls `os.Exit(1)` — defers also don't run.

</details>

---

## 📋 SECTION 3: CLOSURES (6 Questions)

### Q14: What is a closure?

**A)** A function that closes a file  
**B)** A function that captures and retains access to variables from the scope in which it was defined — even after that scope has exited  
**C)** A function with no parameters  
**D)** An anonymous function called immediately  

<details><summary>💡 Answer</summary>

**B) A function that captures variables from its surrounding scope**

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++    // count is captured — lives beyond makeCounter's return
        return count
    }
}

c := makeCounter()
fmt.Println(c())  // 1
fmt.Println(c())  // 2
fmt.Println(c())  // 3
```

`count` is captured by reference. Each call to `c()` reads and modifies the same `count` variable.

</details>

---

### Q15: What is the classic closure-in-loop bug?

**A)** Loops and closures can't be used together  
**B)** When a closure captures a loop variable by reference, all closures share the SAME variable — which has the loop's final value by the time they execute  
**C)** Closures in loops cause infinite loops  
**D)** The loop index is always 0 inside a closure  

<details><summary>💡 Answer</summary>

**B) All closures share the same loop variable — they all see the final value**

```go
// BUG (pre-Go 1.22):
funcs := make([]func(), 3)
for i := 0; i < 3; i++ {
    funcs[i] = func() { fmt.Println(i) }  // captures i by reference
}
// All print 3 (final value of i)

// FIX (pre-1.22): copy the variable
for i := 0; i < 3; i++ {
    i := i  // new i in inner scope — captures THIS copy
    funcs[i] = func() { fmt.Println(i) }
}
// Prints 0, 1, 2

// In Go 1.22+: loop variable semantics changed — each iteration gets its own copy
```

This bug has burned countless developers. The `i := i` fix is the classic idiom.

</details>

---

### Q16: What does this print in Go 1.21 and earlier?
```go
funcs := []func(){}
for _, v := range []int{1, 2, 3} {
    funcs = append(funcs, func() { fmt.Println(v) })
}
for _, f := range funcs { f() }
```

**A)** `1 2 3`  
**B)** `3 3 3` — all closures capture the same `v`; final value is 3  
**C)** `1 1 1`  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) `3 3 3` in Go 1.21 and earlier; `1 2 3` in Go 1.22+**

In Go 1.21 and earlier: `v` is a single variable in the loop's scope; all closures capture the same `v`. By the time the closures run, `v` = 3 (the last assigned value).

In Go 1.22+: loop variables have per-iteration semantics by default, so each closure captures its own `v`. The book covers the pre-1.22 behavior as the important thing to understand.

</details>

---

### Q17: When is a function value `nil`?

**A)** Never — function values are always initialized  
**B)** When declared as `var f func()` (nil) or when a function returns `nil` for a function type  
**C)** After it is called once  
**D)** Only for imported functions  

<details><summary>💡 Answer</summary>

**B) A declared-but-unassigned function variable is `nil`**

```go
var f func(int) int     // nil
f(5)                     // panic: call of nil function value

// Always check before calling:
if f != nil {
    f(5)
}
```

Calling a nil function panics. This is equivalent to calling a method on a nil pointer — always validate before calling function values you received from outside code.

</details>

---

### Q18: What is an immediately invoked function expression (IIFE) in Go?

**A)** A function that calls itself (recursion)  
**B)** An anonymous function defined and called immediately: `func() { ... }()`  
**C)** A built-in Go feature  
**D)** A function invoked via reflection  

<details><summary>💡 Answer</summary>

**B) Anonymous function defined and called immediately**

```go
result := func(x, y int) int {
    return x + y
}(3, 4)  // called immediately with arguments 3 and 4
fmt.Println(result)  // 7

// Common pattern with defer for complex initialization:
cleanup := func() {
    // setup code
}()
```

IIFEs are useful for initializing complex values in one expression, or for creating a new scope inside a function.

</details>

---

### Q19: What is the type of `func(x, y int) int`?

**A)** `Function`  
**B)** `func(int, int) int` — function type includes parameter types and return types  
**C)** `func`  
**D)** `callable`  

<details><summary>💡 Answer</summary>

**B) `func(int, int) int`**

```go
type MathFunc func(int, int) int   // named function type

var add MathFunc = func(x, y int) int { return x + y }

// Higher-order function:
func apply(f func(int, int) int, a, b int) int {
    return f(a, b)
}
```

Parameter names are optional in type declarations. `func(x, y int) int` and `func(int, int) int` are the same type.

</details>

---

## 📋 SECTION 4: ADVANCED PATTERNS (6 Questions)

### Q20: How do you write a function that accepts a function as a parameter?

**A)** `func apply(f Function, x int) int`  
**B)** `func apply(f func(int) int, x int) int`  
**C)** `func apply(f interface{}, x int) int`  
**D)** Functions cannot be parameters  

<details><summary>💡 Answer</summary>

**B) Specify the full function type in the parameter**

```go
func apply(f func(int) int, x int) int {
    return f(x)
}

double := func(x int) int { return x * 2 }
fmt.Println(apply(double, 5))  // 10
fmt.Println(apply(func(x int) int { return x + 1 }, 5))  // 6
```

</details>

---

### Q21: What is the "functional options" pattern used for?

**A)** Making functions optional  
**B)** A pattern where a function accepts variadic options (themselves functions) that configure a struct — allows extensible APIs without changing function signatures  
**C)** Optional return values  
**D)** Making closures optional  

<details><summary>💡 Answer</summary>

**B) Variadic option functions for configuring structs — extensible API design**

```go
type Server struct { host string; port int; timeout time.Duration }

type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.timeout = d }
}
func WithPort(p int) Option {
    return func(s *Server) { s.port = p }
}

func NewServer(host string, opts ...Option) *Server {
    s := &Server{host: host, port: 8080, timeout: 30 * time.Second}
    for _, opt := range opts { opt(s) }
    return s
}

srv := NewServer("localhost", WithPort(9090), WithTimeout(10*time.Second))
```

This is a very common pattern in Go libraries. It avoids large config structs and allows backward-compatible addition of new options.

</details>

---

### Q22: Can a Go function call itself recursively?

**A)** No — Go does not support recursion  
**B)** Yes — but Go does not guarantee tail-call optimization; deep recursion can stack-overflow  
**C)** Yes — Go optimizes all tail calls  
**D)** Only if the function is named  

<details><summary>💡 Answer</summary>

**B) Yes — recursion works, but Go does NOT optimize tail calls**

```go
func factorial(n int) int {
    if n <= 1 { return 1 }
    return n * factorial(n-1)  // recursive call
}
```

Go goroutines start with a small stack (8KB) that grows dynamically up to a limit (default 1GB). Deep recursion can hit this limit. For very deep recursion, consider an iterative approach or explicit stack.

</details>

---

### Q23: What is the `init` function?

**A)** The first function that must be called in `main`  
**B)** A special function that runs before `main`, used to set up package state — no arguments, no return values, called automatically  
**C)** Same as a constructor  
**D)** A function that initializes variables  

<details><summary>💡 Answer</summary>

**B) Special function that runs before `main` for package initialization**

```go
var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "...")
    if err != nil { log.Fatal(err) }
}
```

Rules: no parameters, no return values, cannot be called explicitly, multiple `init` functions are allowed in one file (all run). The book recommends using `init` sparingly — prefer explicit initialization in `main` or constructor functions.

</details>

---

### Q24: What is the difference between a function and a method in Go?

**A)** Methods have return values; functions don't  
**B)** A method is a function with a receiver — it is associated with a specific type  
**C)** Methods are faster  
**D)** Functions can't accept structs  

<details><summary>💡 Answer</summary>

**B) A method has a receiver — it is called on a specific type**

```go
// Function:
func add(x, y int) int { return x + y }

// Method (on type Counter):
type Counter struct { count int }
func (c *Counter) Increment() { c.count++ }
func (c Counter) Value() int { return c.count }

var c Counter
c.Increment()       // method call
add(1, 2)           // function call
```

Methods are covered in depth in Chapter 7.

</details>

---

### Q25: What does this print?
```go
func do() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("wrapped: %w", err)
        }
    }()
    return errors.New("original error")
}

err := do()
fmt.Println(err)
```

**A)** `original error`  
**B)** `wrapped: original error` — the deferred closure modifies the named return `err`  
**C)** `nil`  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) `wrapped: original error`**

The `return errors.New("original error")` sets the named return `err` to the error value. Then the deferred closure runs, sees `err != nil`, and wraps it. The final returned value is the wrapped error. This is the pattern used by libraries to add context to errors automatically.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 23–25 ✅ | **Excellent.** Functions and closures mastered — proceed to Chapter 6. |
| 20–22 ✅ | **Ready.** Review `defer` argument capture and the closure-in-loop bug. |
| 15–19 ⚠️ | **Review first.** Closures, `defer`, and function types are used everywhere. |
| Below 15 ❌ | **Re-read Chapter 5.** These patterns are fundamental to idiomatic Go. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q7 | Multiple returns, variadic `...`, function as first-class value, spreading slices |
| Q8–Q13 | `defer` LIFO, argument capture timing, `defer` + panic, `os.Exit` bypasses defer |
| Q14–Q19 | Closures, loop variable capture bug, nil function values, IIFEs, function types |
| Q20–Q25 | Higher-order functions, functional options, `init`, methods vs functions, defer + named returns |