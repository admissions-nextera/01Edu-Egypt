# Learning Go — Chapter 2: Questions
### Jon Bodner | Predeclared Types, Variables & Constants

---

## Section 1: Memory Layout of Predeclared Types

**Q1.**
Go has `int8`, `int16`, `int32`, `int64` — and also plain `int`. At the hardware level, what determines the size of `int` on a given machine? Is `int` on a 64-bit Linux system the same type as `int64`? Are they interchangeable? Reason from the ABI and register width.

---

**Q2.**
Go's `uint` and `int` are the same size but differ in bit interpretation. At the CPU level, is there a physical difference between a signed and unsigned integer stored in a register? What actually changes — the bits, the operations, or the interpretation of overflow?

---

**Q3.**
`float32` and `float64` both represent real numbers, but they are not interchangeable. At the hardware level, what is the structural difference in how these two types encode a decimal value? What specific precision trade-off is made when you choose `float32` over `float64`?

---

**Q4.**
A Go `string` is described in the spec as an immutable sequence of bytes. But at the **runtime memory layout** level, a string variable is not just a pointer. What is the actual in-memory representation of a string variable in Go? How many words does it occupy on the stack, and what do they contain?

---

**Q5.**
Go's `bool` is 1 byte in memory, not 1 bit. Why? A single bit can represent true/false. What is the hardware and memory-alignment reason that a `bool` cannot be 1 bit in most architectures?

---

## Section 2: Zero Values

**Q6.**
Every declared variable in Go is initialized to its **zero value** if no explicit value is given. This is a deliberate design decision. In C, uninitialized local variables have **undefined behavior** — their value is whatever bytes happened to be on the stack. What OS and hardware mechanism does Go rely on to guarantee zero values, and what is the performance cost of this guarantee?

---

**Q7.**
The zero value for a `bool` is `false`, for numeric types is `0`, for `string` is `""`, and for pointers is `nil`. These are not arbitrary. For each one, explain what the zero value means at the **binary/memory level** — i.e., what pattern of bits is actually written into memory?

---

**Q8.**
Given this declaration: `var s string`
The zero value is `""`. You know a string is a two-word struct (pointer + length). What are the exact bit values stored in those two words for the zero string? What would happen at the machine level if you tried to dereference the data pointer of a zero-value string?

---

## Section 3: Variable Declaration — `var` vs `:=`

**Q9.**
Go offers four ways to declare a variable:
```
var x int = 10
var x = 10
var x int
x := 10
```
These are not stylistic alternatives — they have different semantics. For each form, state precisely what the compiler infers or enforces, and in which **scoping contexts** each is valid or invalid.

---

**Q10.**
`:=` (short variable declaration) can declare **multiple variables at once**, and at least one must be **new**. Consider:
```
x := 1
x, y := 2, 3
```
The second line does not fail even though `x` already exists. What exactly does the compiler do here? Is `x` re-declared, re-assigned, or something else? What is the rule at the symbol table level?

---

**Q11.**
`:=` is not allowed at **package level** (outside any function). Why not? This is not an arbitrary restriction. Reason from how the Go compiler processes package-level declarations vs. function-level declarations and what `:=`'s type inference requires.

---

**Q12.**
Consider:
```
var x int = 10
```
vs.
```
x := 10
```
Both result in `x` being an `int` with value `10`. Are these identical from the compiler's perspective? Is there any scenario where the generated machine code or memory layout differs between the two?

---

## Section 4: Explicit Type Conversion

**Q13.**
Go has **no implicit type conversion**. Every conversion between numeric types must be explicit: `int32(x)`, `float64(y)`, etc. Most other languages (C, Java, Python) perform implicit widening conversions. What is the **compiler design argument** for forcing explicitness? What class of bugs does this prevent?

---

**Q14.**
Hand-trace this conversion:
```
var x int32 = 1000
var y int8 = int8(x)
```
What is the **binary representation** of `1000` as `int32`? When truncated to 8 bits, what value does `y` hold? Show the bit-level work. What does this reveal about type conversion safety in Go?

---

**Q15.**
Converting between `float64` and `int` truncates, not rounds:
```
var f float64 = 3.99
var i int = int(f)
```
Hand-trace this. What is `i`? At the **CPU instruction level**, what operation performs this conversion — is it a rounding instruction or a truncation instruction? What does this imply when using integer conversion in financial calculations?

---

**Q16.**
You cannot use a `bool` as a numeric type in Go. `if 1` is a compile error. In C, any non-zero integer is truthy. What does Go's strict boolean type system prevent at the **compiler type-checking** level? What class of logic errors does this eliminate?

---

## Section 5: `const` and Untyped Constants

**Q17.**
Go constants are evaluated entirely at **compile time**. What does this mean for where constant values live in the compiled binary? Are constants stored in the `.data` section, `.rodata`, embedded in instructions, or something else? Can a constant ever occupy heap memory?

---

**Q18.**
Go has **typed** and **untyped** constants. An untyped constant like `const x = 10` has a **kind** (untyped integer) but no fixed type. A typed constant like `const x int32 = 10` is fixed. What is the practical consequence of this distinction when the constant is used in an expression? Trace through what happens when an untyped `10` is assigned to `float64` vs. `int8`.

---

**Q19.**
`iota` is a compile-time counter used in `const` blocks. Consider:
```
const (
    A = iota
    B
    C
)
```
What are the values of `A`, `B`, `C`? Now consider:
```
const (
    X = iota + 1
    Y
    Z
)
```
Hand-trace the values of `X`, `Y`, `Z`. At the compiler level, what is `iota` — is it a variable, a keyword, a macro, or something else? Where does its counter reset?

---

**Q20.**
Go does not have an `enum` keyword. The idiomatic Go equivalent uses `iota` in a `const` block with a named type. What problem does adding a **named type** (e.g., `type Direction int`) solve over a plain `iota` block? Think in terms of the **type checker** and function signatures.

---

## Section 6: Contradiction Probes & Hand-Traces

**Q21. — Contradiction Probe**
A classmate says:

> *"In Go, `int` and `int64` are both 64-bit on my machine, so I can pass an `int` variable to a function expecting `int64` without any conversion."*

Without running any code, identify the precise compiler rule this violates. What does Go's type system say about **named types** even when their underlying representation is identical?

---

**Q22. — Hand Trace Required**
Trace through this entire declaration block line by line. For each variable, state its **type**, **value**, and **memory size in bytes**:
```
var a int
var b float64 = 3.14
c := true
var d string = "Go"
const e = 42
```
Then answer: which of these live on the **stack**, which in **read-only memory**, and which (if any) could escape to the **heap**?

---

**Q23. — Contradiction Probe**
A classmate says:

> *"Go's zero values are just a safety convenience — the compiler inserts a memset to zero out variables, which adds overhead on every function call. Real systems code should use `var x int` only when needed to avoid this cost."*

Is this reasoning correct? What does the OS and hardware actually guarantee about freshly allocated stack frames and heap pages, and when is an explicit zeroing operation necessary vs. already handled by the system?

---

*End of Chapter 2 Questions*