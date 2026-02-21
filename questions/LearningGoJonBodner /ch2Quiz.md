# 📘 Learning Go — Chapter 2 Quiz
## Predeclared Types and Declarations

**Time Limit:** 35 minutes  
**Total Questions:** 25  
**Passing Score:** 20/25 (80%)

> This quiz covers: built-in types, zero values, `var` vs `:=`, type conversions, constants, and typed vs untyped constants.

---

## 📋 SECTION 1: BUILT-IN TYPES & ZERO VALUES (7 Questions)

### Q1: What is the zero value of each type below?

```go
var a int
var b string
var c bool
var d float64
```

**A)** `0`, `""`, `false`, `0.0`  
**B)** `nil`, `nil`, `nil`, `nil`  
**C)** `0`, `nil`, `false`, `0.0`  
**D)** Uninitialized — accessing them is undefined behavior  

<details><summary>💡 Answer</summary>

**A) `0`, `""`, `false`, `0.0`**

Go guarantees every variable is initialized to its zero value if no explicit value is provided. There is no "uninitialized memory" in Go:
- Numeric types: `0`
- `string`: `""` (empty string)
- `bool`: `false`
- Pointers, slices, maps, channels, functions: `nil`

This is one of Go's most important safety properties.

</details>

---

### Q2: What integer type should you use when you just need "a whole number" and have no special requirements?

**A)** `int32` — most portable  
**B)** `int64` — always large enough  
**C)** `int` — platform-sized integer, correct default choice  
**D)** `uint` — unsigned is faster  

<details><summary>💡 Answer</summary>

**C) `int` — the default integer type**

`int` is either 32 or 64 bits depending on the platform (64-bit on modern systems). Use `int` as the default. Use specific sizes (`int32`, `int64`, `uint8`) only when you need interoperability with binary protocols, specific memory layouts, or the standard library requires it (e.g. `byte` = `uint8`).

</details>

---

### Q3: What is the difference between `float32` and `float64`?

**A)** `float32` holds integers; `float64` holds decimals  
**B)** `float64` has more precision (15-17 decimal digits vs 6-9 for `float32`) — use `float64` unless you have a specific reason not to  
**C)** `float32` is faster on all modern hardware  
**D)** There is no difference in Go  

<details><summary>💡 Answer</summary>

**B) `float64` has more precision — it's the default choice**

```go
var f float64 = 3.14159265358979323846  // use float64
```

The book recommends `float64` as the default. `float32` has about 7 significant decimal digits of precision; `float64` has about 15. Floating-point comparison quirks apply to both — never use `==` for float comparison; use a tolerance.

</details>

---

### Q4: What is the output?
```go
var x int = 10
var y float64 = 3.14
fmt.Println(x + int(y))
```

**A)** `13.14`  
**B)** `13`  
**C)** Compile error — can't convert float64 to int  
**D)** `13.0`  

<details><summary>💡 Answer</summary>

**B) `13`**

`int(3.14)` truncates toward zero → `3`. Then `10 + 3 = 13`. Go requires explicit type conversions — there are NO implicit conversions between numeric types. `x + y` without the conversion would be a compile error: "mismatched types int and float64."

</details>

---

### Q5: What is `byte` in Go?

**A)** A separate type with special behavior  
**B)** An alias for `uint8` — the type used for individual bytes of data  
**C)** An alias for `int8`  
**D)** A pointer to a character  

<details><summary>💡 Answer</summary>

**B) An alias for `uint8`**

`byte` and `uint8` are completely interchangeable — they are the same type. Similarly, `rune` is an alias for `int32` and represents a Unicode code point. These aliases exist for readability: `[]byte` clearly signals "raw bytes," while `[]uint8` is less obvious.

</details>

---

### Q6: What does this print?
```go
var s string = "hello"
fmt.Println(len(s))
```

And if `s = "héllo"` (with an accent)?

**A)** `5` for both — `len` counts characters  
**B)** `5` for "hello", `6` for "héllo" — `len` counts bytes, not characters; `é` is 2 bytes in UTF-8  
**C)** `5` for "hello", `5` for "héllo" — Go normalizes Unicode  
**D)** Compile error — non-ASCII characters are not allowed in strings  

<details><summary>💡 Answer</summary>

**B) `5` for "hello", `6` for "héllo"**

`len(s)` returns the number of **bytes**, not characters. Go strings are sequences of bytes. UTF-8 encodes most accented characters as 2 bytes. To count Unicode characters (runes): `utf8.RuneCountInString(s)` or `len([]rune(s))`.

</details>

---

### Q7: Is this valid? What does it print?
```go
var b bool
if b {
    fmt.Println("true")
} else {
    fmt.Println("false")
}
```

**A)** Compile error — `b` is uninitialized  
**B)** Runtime panic  
**C)** `false` — `b` is zero-valued to `false`  
**D)** Undefined behavior  

<details><summary>💡 Answer</summary>

**C) `false`**

This is perfectly valid. `var b bool` initializes `b` to `false` (the zero value for `bool`). The `if` evaluates `false`, so the `else` branch runs. No error, no panic, no undefined behavior — this is Go's zero value guarantee in action.

</details>

---

## 📋 SECTION 2: VAR, :=, AND DECLARATIONS (7 Questions)

### Q8: What is the difference between `var x int = 10` and `x := 10`?

**A)** `:=` is for package-level declarations; `var` is for function-level  
**B)** `:=` declares AND assigns in one step (only inside functions); `var` can be used at package level or when you want to separate declaration from assignment  
**C)** `var` is slower  
**D)** They are identical  

<details><summary>💡 Answer</summary>

**B) `:=` is shorthand inside functions; `var` works everywhere**

```go
var x int = 10  // explicit type — anywhere
var y = 10      // type inferred — anywhere  
z := 10         // shorthand — only inside functions
```

`:=` cannot be used at package level. When the type is not obvious from the value (e.g. `var x float64 = 10`), use `var` with the explicit type.

</details>

---

### Q9: What is wrong with this code?
```go
func main() {
    x := 10
    x := 20
    fmt.Println(x)
}
```

**A)** Nothing — `x` is just reassigned  
**B)** Compile error — `:=` cannot redeclare a variable already declared in the same scope  
**C)** `x` will be `10` because the second `:=` is ignored  
**D)** Compile error — can't assign `20` to an `int`  

<details><summary>💡 Answer</summary>

**B) Compile error — `x` already declared in this scope**

`:=` declares a NEW variable. If `x` already exists in the same scope, use `=` (assignment) instead:
```go
x := 10
x = 20      // correct: assignment
```

Exception: `:=` can redeclare if at least one variable on the left side is new:
```go
x, err := doSomething()
y, err := doOther()  // ok — y is new; err is reassigned
```

</details>

---

### Q10: What type does Go infer for `x := 10`? For `y := 10.5`?

**A)** `int32` and `float32`  
**B)** `int` and `float64`  
**C)** `int64` and `float64`  
**D)** Depends on the platform  

<details><summary>💡 Answer</summary>

**B) `int` and `float64`**

Go's type inference uses the default types for untyped constants:
- Integer literals → `int`
- Floating-point literals → `float64`
- String literals → `string`
- Boolean literals → `bool`

If you need a different type, declare it explicitly: `var x int64 = 10`.

</details>

---

### Q11: Is this valid Go?
```go
var (
    x    int    = 10
    name string = "Alice"
    flag bool
)
```

**A)** No — `var` blocks are not valid Go syntax  
**B)** Yes — `var` blocks declare multiple variables cleanly, with `flag` getting the zero value `false`  
**C)** No — all variables in a `var` block must have the same type  
**D)** No — `var` blocks can only appear at package level  

<details><summary>💡 Answer</summary>

**B) Yes — valid, idiomatic Go for multiple declarations**

`var` blocks are idiomatic for declaring related variables together. `flag` is not assigned, so it gets the zero value `false`. `var` blocks can appear both at package level and inside functions.

</details>

---

### Q12: When should you use `var x int` instead of `x := 0`?

**A)** Never — `:=` is always better  
**B)** When you want to explicitly document the type, when the zero value is the correct initial value, or when declaring at package level  
**C)** When performance is critical  
**D)** When the type is `float64`  

<details><summary>💡 Answer</summary>

**B) When type documentation matters, zero value is correct, or at package level**

```go
var count int      // clearly: an int starting at zero
var buf bytes.Buffer  // zero value is a valid, usable Buffer

// vs:
count := 0         // type is inferred but int is obvious here
```

The book emphasizes: use `var` when the zero value is intentionally the starting value, making the intent clear. Use `:=` when you're immediately assigning a meaningful value.

</details>

---

### Q13: What does the blank identifier `_` do on the left side of an assignment?

**A)** Creates a variable named `_`  
**B)** Discards the value — tells Go "I know there's a value here but I don't need it"  
**C)** Causes a compile error  
**D)** Assigns to the previous value of `_`  

<details><summary>💡 Answer</summary>

**B) Discards the value**

```go
x, _ := strconv.Atoi("42")   // discard the error
for _, v := range slice { }   // discard the index

// _ can appear multiple times:
_, err := fmt.Println("hello")
```

`_` is a write-only variable — you can assign to it but never read from it. It silences the "declared and not used" error for values you intentionally ignore.

</details>

---

### Q14: Is this valid? What does it print?
```go
x := 10
{
    x := 20
    fmt.Println(x)
}
fmt.Println(x)
```

**A)** Compile error — can't declare `x` twice  
**B)** `20` then `10` — the inner `x` shadows the outer `x` within the block  
**C)** `20` then `20` — both refer to the same variable  
**D)** `10` then `10` — the inner declaration is ignored  

<details><summary>💡 Answer</summary>

**B) `20` then `10` — shadowing**

The inner `{` `}` creates a new block scope. `:=` inside that block creates a NEW `x` that shadows the outer one. When the inner block ends, the inner `x` is gone. The outer `x` is still `10`. This is a common source of bugs — the book emphasizes watching for unintentional shadowing.

</details>

---

## 📋 SECTION 3: CONSTANTS (5 Questions)

### Q15: What is the difference between typed and untyped constants in Go?

**A)** Untyped constants don't exist — all constants have types  
**B)** Typed constants have a specific type; untyped constants have a "default type" that is assigned when the constant is used in an expression — they can be used with different numeric types without explicit conversion  
**C)** Untyped constants are slower  
**D)** Typed constants can only be used with `const`  

<details><summary>💡 Answer</summary>

**B) Untyped constants have a default type used contextually**

```go
const x = 10      // untyped integer constant — can be used as int, int64, float64, etc.
const y int = 10  // typed — only usable as int

var a int32 = x   // OK — x adapts
var b int32 = y   // compile error — y is int, not int32
```

Untyped constants are one of Go's subtle but important features — they make mathematical expressions across types natural.

</details>

---

### Q16: What is `iota` used for?

**A)** A built-in function for generating random numbers  
**B)** A predeclared identifier in `const` blocks that starts at `0` and increments by `1` for each constant — used to create enumerated values  
**C)** A keyword for infinite loops  
**D)** A package for I/O  

<details><summary>💡 Answer</summary>

**B) Increments by 1 for each constant in a `const` block — used for enumerations**

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
)

// More complex: bit flags
const (
    Read    = 1 << iota  // 1
    Write                // 2
    Execute              // 4
)
```

`iota` resets to 0 at the start of each `const` block.

</details>

---

### Q17: Can a constant hold a value that changes at runtime (e.g. the result of `time.Now()`)?

**A)** Yes — constants can hold runtime values  
**B)** No — constants must be computable at compile time: literals, arithmetic on literals, or values from other constants  
**C)** Only if the type is `interface{}`  
**D)** Only `int` and `string` constants can be computed  

<details><summary>💡 Answer</summary>

**B) No — constants must be compile-time values**

```go
const x = 10 * 2    // valid — pure arithmetic
const y = len("hi") // valid — some builtins are compile-time
const z = time.Now() // compile error — runtime function call
```

This is fundamentally different from `var` — variables can hold runtime values, constants cannot.

</details>

---

### Q18: What value does this print?
```go
const x = 1 / 2
fmt.Println(x)
```

**A)** `0.5`  
**B)** `0` — integer division  
**C)** Compile error  
**D)** `0` for `int`, `0.5` for `float64` — depends on context  

<details><summary>💡 Answer</summary>

**B) `0` — integer division because both literals are untyped integers**

Untyped constant arithmetic uses the type of the operands. `1` and `2` are untyped integer constants, so `1/2` is integer division = `0`. To get `0.5`: `const x = 1.0 / 2` or `const x = 1 / 2.0`.

</details>

---

### Q19: What is wrong with this code?
```go
const Pi = 3.14159
Pi = 3.14
```

**A)** Nothing — reassigning constants is fine  
**B)** Compile error — constants cannot be reassigned  
**C)** Runtime error  
**D)** `Pi` becomes `3.14` silently  

<details><summary>💡 Answer</summary>

**B) Compile error — constants are immutable**

"Cannot assign to Pi (untyped float constant)" — constants are fixed at declaration. This is the entire point: use `const` when the value must never change. If you need a value that changes, use `var`.

</details>

---

## 📋 SECTION 4: TYPE CONVERSIONS (4 Questions)

### Q20: Go has no implicit type conversions. What does this mean in practice?

**A)** You can never convert between types  
**B)** Every type conversion must be written explicitly in code — `int(x)`, `float64(y)`, `string(z)` — the compiler never automatically converts  
**C)** Only numeric types can be converted  
**D)** Conversions always truncate the value  

<details><summary>💡 Answer</summary>

**B) Every conversion must be explicit**

```go
var x int = 10
var y float64 = x        // compile error — implicit conversion not allowed
var z float64 = float64(x)  // correct — explicit conversion
```

In C/Java, `double d = someInt` works implicitly. Go requires `float64(someInt)`. This verbosity prevents accidental precision loss and makes code intent clear.

</details>

---

### Q21: What is the result of `int(-3.7)`?

**A)** `-4` — rounds to nearest  
**B)** `-3` — truncates toward zero  
**C)** Compile error  
**D)** Runtime panic  

<details><summary>💡 Answer</summary>

**B) `-3` — truncates toward zero (not toward negative infinity)**

`int()` truncates, not rounds. For positive floats: `int(3.7) == 3`. For negative floats: `int(-3.7) == -3` (toward zero, not `-4`). If you need rounding: `int(math.Round(x))`.

</details>

---

### Q22: What happens when you convert a `string` to `[]byte`?

**A)** Compile error  
**B)** Returns a copy of the string's underlying bytes — modifying the slice does NOT modify the original string  
**C)** Returns a pointer to the string's bytes — modifying the slice modifies the string  
**D)** Returns the ASCII codes of each character  

<details><summary>💡 Answer</summary>

**B) Returns a copy of the string's bytes — strings are immutable**

```go
s := "hello"
b := []byte(s)    // copy
b[0] = 'H'        // modifies b, NOT s
fmt.Println(s)    // "hello" — unchanged
fmt.Println(string(b))  // "Hello"
```

Strings in Go are immutable. Converting to `[]byte` always makes a copy. This is safe but has a performance cost — avoid unnecessary conversions in hot paths.

</details>

---

### Q23: What does converting an integer to a string do in Go?

```go
x := 65
s := string(x)
fmt.Println(s)
```

**A)** `"65"` — converts the number to its string representation  
**B)** `"A"` — converts the integer as a Unicode code point  
**C)** Compile error  
**D)** `"\x41"` — the hex representation  

<details><summary>💡 Answer</summary>

**B) `"A"` — treats the integer as a Unicode code point**

`string(65)` gives `"A"` (Unicode code point 65 = 'A'). To get `"65"`, use `strconv.Itoa(65)` or `fmt.Sprintf("%d", 65)`. This is a common, real bug. In Go 1.15+, `go vet` warns about this conversion.

</details>

---

### Q24: What is the output?
```go
var x int8 = 127
x++
fmt.Println(x)
```

**A)** `128`  
**B)** `-128` — integer overflow wraps around  
**C)** Compile error  
**D)** Runtime panic  

<details><summary>💡 Answer</summary>

**B) `-128` — integer overflow, wraps around**

`int8` ranges from -128 to 127. Adding 1 to 127 wraps to -128 (two's complement). Go does NOT panic on integer overflow — it silently wraps. This can be a source of bugs if you use fixed-size integer types near their limits. The compiler will catch overflow of **constant** expressions, but not runtime overflow.

</details>

---

### Q25: Is this valid Go?
```go
type Celsius float64
type Fahrenheit float64

func main() {
    var c Celsius = 100
    var f Fahrenheit = Fahrenheit(c)
    fmt.Println(f)
}
```

**A)** No — you can't define types based on `float64`  
**B)** No — `Celsius` and `Fahrenheit` are different types; you must convert  
**C)** Yes — `Celsius` and `Fahrenheit` both have the same underlying type `float64`, so the explicit conversion is valid  
**D)** Yes — but only if `Celsius` and `Fahrenheit` are in the same package  

<details><summary>💡 Answer</summary>

**C) Yes — explicit conversion between types with the same underlying type is valid**

Named types with the same underlying type can be explicitly converted between each other. `var f Fahrenheit = c` (without conversion) would be a compile error — they are distinct types. `Fahrenheit(c)` is explicit and valid. This is the correct way to implement type-safe units.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 23–25 ✅ | **Excellent.** Solid type foundations — proceed to Chapter 3. |
| 20–22 ✅ | **Ready.** Review missed questions on type conversions or constants. |
| 15–19 ⚠️ | **Review first.** Type system fundamentals are load-bearing for everything ahead. |
| Below 15 ❌ | **Re-read Chapter 2.** Zero values, `:=` vs `var`, and explicit conversions must be solid. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q7 | Zero values, `byte`/`rune`, `len` counts bytes, `int` as default, `float64` as default |
| Q8–Q14 | `var` vs `:=`, shadowing, blank identifier, `var` blocks |
| Q15–Q19 | Typed vs untyped constants, `iota`, compile-time only, integer division in constants |
| Q20–Q25 | Explicit conversions, truncation direction, `string(int)` trap, overflow wrapping, named types |