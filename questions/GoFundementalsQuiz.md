# 🎯 Go Fundamentals Quiz
## Data Types, Zero Values, Loops & Conditions

**Time Limit:** 45 minutes  
**Total Questions:** 50  
**Passing Score:** 40/50 (80%)

---

## 📋 SECTION 1: QUOTES & STRINGS (10 Questions)

### Q1: What's the difference between single quotes `' '` and double quotes `" "`?

**A)** No difference  
**B)** Single quotes for rune/char, double quotes for string  
**C)** Single quotes for string, double quotes for rune  
**D)** Both can be used interchangeably  

<details><summary>💡 Answer</summary>

**B) Single quotes for rune/char, double quotes for string**

```go
ch := 'A'      // rune (single character)
str := "Hello" // string (multiple characters)
```

</details>

---

### Q2: What's the output?
```go
fmt.Println('A')
```

**A)** A  
**B)** 65  
**C)** 'A'  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) 65**

`'A'` is a rune (int32), prints its ASCII value: 65

To print the character:
```go
fmt.Println(string('A'))  // Outputs: A
```

</details>

---

### Q3: What's the output?
```go
fmt.Println("A")
```

**A)** A  
**B)** 65  
**C)** "A"  
**D)** Error  

<details><summary>💡 Answer</summary>

**A) A**

Double quotes create a string, prints the character directly.

</details>

---

### Q4: Which is valid?
```go
a := 'Hello'
b := "Hello"
c := 'H'
d := "H"
```

**A)** a, b  
**B)** b, c, d  
**C)** c, d  
**D)** All invalid  

<details><summary>💡 Answer</summary>

**B) b, c, d**

```go
// a := 'Hello'  // ❌ Error: single quotes only for single character
b := "Hello"     // ✅ String
c := 'H'         // ✅ Rune (single character)
d := "H"         // ✅ String (one character)
```

</details>

---

### Q5: What's the type of `'A'`?

**A)** string  
**B)** byte  
**C)** rune (int32)  
**D)** char  

<details><summary>💡 Answer</summary>

**C) rune (int32)**

```go
ch := 'A'
fmt.Printf("%T", ch)  // Output: int32
```

</details>

---

### Q6: What's the output?
```go
fmt.Println('A' + 1)
```

**A)** A1  
**B)** B  
**C)** 66  
**D)** Error  

<details><summary>💡 Answer</summary>

**C) 66**

'A' = 65, adding 1 gives 66

To get 'B':
```go
fmt.Println(string('A' + 1))  // Output: B
```

</details>

---

### Q7: Can strings be modified directly?
```go
s := "Hello"
s[0] = 'h'
```

**A)** Yes  
**B)** No, strings are immutable  
**C)** Only if using var  
**D)** Only uppercase letters  

<details><summary>💡 Answer</summary>

**B) No, strings are immutable**

```go
// ❌ Error: cannot assign to s[0]
s := "Hello"
s[0] = 'h'

// ✅ Correct way:
s := "Hello"
s = "hello"  // Reassign entire string

// OR convert to []byte:
bytes := []byte(s)
bytes[0] = 'h'
s = string(bytes)
```

</details>

---

### Q8: What's the output?
```go
s := "Go"
fmt.Println(s[0])
```

**A)** G  
**B)** 71  
**C)** "G"  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) 71**

String indexing returns a byte (uint8), which prints as ASCII value.

To print character:
```go
fmt.Println(string(s[0]))  // Output: G
```

</details>

---

### Q9: What's the difference between `""` and `''`?

**A)** No difference  
**B)** `""` is empty string, `''` is invalid  
**C)** `""` is string, `''` is empty rune  
**D)** Both are strings  

<details><summary>💡 Answer</summary>

**B) `""` is empty string, `''` is invalid**

```go
s := ""   // ✅ Empty string (valid)
// c := ''  // ❌ Error: empty character literal
c := ' '  // ✅ Space character (valid)
```

</details>

---

### Q10: What's the output?
```go
fmt.Println("Hello" + "World")
```

**A)** HelloWorld  
**B)** Hello World  
**C)** Error  
**D)** Hello+World  

<details><summary>💡 Answer</summary>

**A) HelloWorld**

String concatenation using `+` operator.

</details>

---

## 📋 SECTION 2: ASCII CODE (5 Questions)

### Q11: What's the ASCII value of '0'?

**A)** 0  
**B)** 48  
**C)** 65  
**D)** 97  

<details><summary>💡 Answer</summary>

**B) 48**

Important ASCII values:
- '0' = 48
- 'A' = 65
- 'a' = 97

</details>

---

### Q12: Convert digit character '5' to integer 5:

**A)** `int('5')`  
**B)** `'5' - '0'`  
**C)** `'5' - 48`  
**D)** Both B and C  

<details><summary>💡 Answer</summary>

**D) Both B and C**

```go
digit := '5' - '0'  // 53 - 48 = 5
// or
digit := '5' - 48   // Same thing
```

</details>

---

### Q13: Convert lowercase 'a' to uppercase 'A':

**A)** `'a' - 32`  
**B)** `'a' + 32`  
**C)** `'a' - 'a' + 'A'`  
**D)** Both A and C  

<details><summary>💡 Answer</summary>

**D) Both A and C**

```go
upper := 'a' - 32           // 97 - 32 = 65 = 'A'
// or
upper := 'a' - 'a' + 'A'   // 97 - 97 + 65 = 65 = 'A'
```

</details>

---

### Q14: What's the output?
```go
fmt.Println('z' - 'a')
```

**A)** za  
**B)** 25  
**C)** 26  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) 25**

'z' = 122, 'a' = 97, difference = 25

</details>

---

### Q15: Check if character is digit:
```go
ch := '5'
if ??? {
    fmt.Println("digit")
}
```

**A)** `ch >= 0 && ch <= 9`  
**B)** `ch >= '0' && ch <= '9'`  
**C)** `ch >= 48 && ch <= 57`  
**D)** Both B and C  

<details><summary>💡 Answer</summary>

**D) Both B and C**

```go
// Method 1:
if ch >= '0' && ch <= '9' {
    fmt.Println("digit")
}

// Method 2:
if ch >= 48 && ch <= 57 {
    fmt.Println("digit")
}
```

</details>

---

## 📋 SECTION 3: ZERO VALUES (8 Questions)

### Q16: What's the zero value of int?

**A)** nil  
**B)** 0  
**C)** -1  
**D)** undefined  

<details><summary>💡 Answer</summary>

**B) 0**

```go
var x int
fmt.Println(x)  // Output: 0
```

</details>

---

### Q17: What's the zero value of string?

**A)** nil  
**B)** "0"  
**C)** ""  
**D)** " "  

<details><summary>💡 Answer</summary>

**C) ""** (empty string)

```go
var s string
fmt.Println(s)        // Output: (nothing)
fmt.Println(len(s))   // Output: 0
```

</details>

---

### Q18: What's the zero value of bool?

**A)** true  
**B)** false  
**C)** 0  
**D)** nil  

<details><summary>💡 Answer</summary>

**B) false**

```go
var b bool
fmt.Println(b)  // Output: false
```

</details>

---

### Q19: What's the zero value of float64?

**A)** 0  
**B)** 0.0  
**C)** nil  
**D)** NaN  

<details><summary>💡 Answer</summary>

**B) 0.0** (or just 0)

```go
var f float64
fmt.Println(f)  // Output: 0
```

</details>

---

### Q20: What's the zero value of pointer?

**A)** 0  
**B)** nil  
**C)** undefined  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) nil**

```go
var p *int
fmt.Println(p)       // Output: <nil>
fmt.Println(p == nil) // Output: true
```

</details>

---

### Q21: What gets printed?
```go
var a int
var b string
var c bool
fmt.Println(a, b, c)
```

**A)** nil nil nil  
**B)** 0 "" false  
**C)** 0 " " false  
**D)** undefined undefined undefined  

<details><summary>💡 Answer</summary>

**B) 0 "" false**

Each type has its zero value.

</details>

---

### Q22: What's the zero value of slice?

**A)** []  
**B)** nil  
**C)** empty slice  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) nil**

```go
var s []int
fmt.Println(s == nil)  // Output: true
fmt.Println(len(s))    // Output: 0
```

</details>

---

### Q23: What's the zero value of map?

**A)** {}  
**B)** nil  
**C)** empty map  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) nil**

```go
var m map[string]int
fmt.Println(m == nil)  // Output: true

// ⚠️ Can't write to nil map:
// m["key"] = 1  // panic!

// Must initialize:
m = make(map[string]int)
m["key"] = 1  // ✅ OK
```

</details>

---

## 📋 SECTION 4: DATA TYPES (12 Questions)

### Q24: Which can hold negative numbers?

**A)** int  
**B)** uint  
**C)** byte  
**D)** All of them  

<details><summary>💡 Answer</summary>

**A) int**

```go
var a int = -5     // ✅ OK
// var b uint = -5  // ❌ Error: uint is unsigned (positive only)
// var c byte = -5  // ❌ Error: byte is uint8
```

</details>

---

### Q25: What's the range of int8?

**A)** 0 to 255  
**B)** -127 to 127  
**C)** -128 to 127  
**D)** -255 to 255  

<details><summary>💡 Answer</summary>

**C) -128 to 127**

8 bits: 2^8 = 256 values
- Signed: -128 to 127
- Unsigned (uint8): 0 to 255

</details>

---

### Q26: What happens?
```go
var x int8 = 127
x = x + 1
fmt.Println(x)
```

**A)** 128  
**B)** -128  
**C)** Error  
**D)** 127  

<details><summary>💡 Answer</summary>

**B) -128**

Integer overflow wraps around:
127 + 1 = -128 (wraps to minimum value)

</details>

---

### Q27: What's the difference between rune and byte?

**A)** No difference  
**B)** rune is int32, byte is uint8  
**C)** rune for UTF-8, byte for ASCII  
**D)** Both B and C  

<details><summary>💡 Answer</summary>

**D) Both B and C**

```go
type rune = int32  // Unicode code point
type byte = uint8  // ASCII character

var r rune = '世'  // ✅ OK (Unicode)
// var b byte = '世'  // ❌ Error: value too large
```

</details>

---

### Q28: What's the output?
```go
var x uint = 5
var y int = -3
fmt.Println(x + y)
```

**A)** 2  
**B)** 8  
**C)** Error  
**D)** -2  

<details><summary>💡 Answer</summary>

**C) Error**

**Cannot mix uint and int without conversion:**

```go
// ❌ Error: invalid operation: mismatched types
var x uint = 5
var y int = -3
// fmt.Println(x + y)

// ✅ Fix: Convert types
fmt.Println(int(x) + y)    // Output: 2
// or
fmt.Println(x + uint(y))   // Would panic if y is negative!
```

</details>

---

### Q29: What's valid for float?

**A)** `var f float = 3.14`  
**B)** `var f float32 = 3.14`  
**C)** `var f float64 = 3.14`  
**D)** Both B and C  

<details><summary>💡 Answer</summary>

**D) Both B and C**

```go
// var f float = 3.14     // ❌ Error: no type called 'float'
var f1 float32 = 3.14     // ✅ OK
var f2 float64 = 3.14     // ✅ OK
f3 := 3.14                // ✅ OK (inferred as float64)
```

</details>

---

### Q30: What's complex number in Go?

**A)** Real + Imaginary parts  
**B)** complex64 and complex128  
**C)** Uses 'i' for imaginary  
**D)** All of the above  

<details><summary>💡 Answer</summary>

**D) All of the above**

```go
var c1 complex64 = 3 + 4i
var c2 complex128 = 5 + 2i

fmt.Println(real(c1))  // Output: 3
fmt.Println(imag(c1))  // Output: 4
```

</details>

---

### Q31: What's the default numeric type?

**A)** int for integers, float32 for floats  
**B)** int64 for integers, float64 for floats  
**C)** int for integers, float64 for floats  
**D)** int32 for integers, float32 for floats  

<details><summary>💡 Answer</summary>

**C) int for integers, float64 for floats**

```go
a := 10     // Type: int
b := 3.14   // Type: float64

fmt.Printf("%T %T", a, b)  // Output: int float64
```

</details>

---

### Q32: Convert float to int:
```go
f := 3.99
i := ???
```

**A)** `i := int(f)`  
**B)** `i := f`  
**C)** `i := (int)f`  
**D)** Automatic conversion  

<details><summary>💡 Answer</summary>

**A) i := int(f)**

```go
f := 3.99
i := int(f)
fmt.Println(i)  // Output: 3 (truncates, doesn't round!)
```

</details>

---

### Q33: What's true about bool?

**A)** Only values: true, false  
**B)** Zero value: false  
**C)** Can't convert from int  
**D)** All of the above  

<details><summary>💡 Answer</summary>

**D) All of the above**

```go
var b bool  // Zero value: false

b = true    // ✅ OK
b = false   // ✅ OK
// b = 1    // ❌ Error: cannot use int as bool
// b = "true"  // ❌ Error: cannot use string as bool
```

</details>

---

### Q34: Range of byte (uint8)?

**A)** -128 to 127  
**B)** 0 to 255  
**C)** 0 to 256  
**D)** -255 to 255  

<details><summary>💡 Answer</summary>

**B) 0 to 255**

```go
var b byte = 255   // ✅ OK (max value)
// var b byte = 256  // ❌ Error: overflow
// var b byte = -1   // ❌ Error: negative
```

</details>

---

### Q35: What's the output?
```go
fmt.Printf("%T", 'A')
fmt.Printf("%T", "A")
```

**A)** string string  
**B)** rune string  
**C)** int32 string  
**D)** char string  

<details><summary>💡 Answer</summary>

**C) int32 string**

```go
fmt.Printf("%T\n", 'A')   // Output: int32
fmt.Printf("%T\n", "A")   // Output: string
```

</details>

---

## 📋 SECTION 5: LOOPS (8 Questions)

### Q36: Valid for loop syntax?

**A)** `for i = 0; i < 10; i++`  
**B)** `for i := 0; i < 10; i++`  
**C)** `for(i := 0; i < 10; i++)`  
**D)** Both A and B  

<details><summary>💡 Answer</summary>

**B) for i := 0; i < 10; i++**

```go
// ✅ Correct:
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// ❌ Wrong: No parentheses in Go
// for(i := 0; i < 10; i++) {
```

</details>

---

### Q37: Infinite loop in Go:

**A)** `for { }`  
**B)** `while(true) { }`  
**C)** `loop { }`  
**D)** `for(;;) { }`  

<details><summary>💡 Answer</summary>

**A) for { }**

```go
// ✅ Go way:
for {
    fmt.Println("infinite")
}

// ❌ No 'while' in Go
// while(true) { }
```

</details>

---

### Q38: While loop equivalent in Go:

**A)** `for condition { }`  
**B)** `while condition { }`  
**C)** `do { } while(condition)`  
**D)** No equivalent  

<details><summary>💡 Answer</summary>

**A) for condition { }**

```go
i := 0
for i < 10 {  // Like while(i < 10)
    fmt.Println(i)
    i++
}
```

</details>

---

### Q39: What's the output?
```go
for i := 0; i < 3; i++ {
    fmt.Print(i)
}
```

**A)** 0 1 2  
**B)** 012  
**C)** 1 2 3  
**D)** 0 1 2 3  

<details><summary>💡 Answer</summary>

**B) 012**

Prints without spaces: 012

</details>

---

### Q40: Range over string:
```go
for i, v := range "Go" {
    fmt.Println(i, v)
}
```

**A)** Prints indices and bytes  
**B)** Prints indices and runes  
**C)** Error  
**D)** Prints only values  

<details><summary>💡 Answer</summary>

**B) Prints indices and runes**

```go
for i, v := range "Go" {
    fmt.Println(i, v)
}
// Output:
// 0 71  (G as rune)
// 1 111 (o as rune)
```

</details>

---

### Q41: Skip iteration in loop:

**A)** `skip`  
**B)** `continue`  
**C)** `next`  
**D)** `break`  

<details><summary>💡 Answer</summary>

**B) continue**

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue  // Skip 2
    }
    fmt.Println(i)
}
// Output: 0 1 3 4
```

</details>

---

### Q42: Exit loop early:

**A)** `exit`  
**B)** `return`  
**C)** `break`  
**D)** `stop`  

<details><summary>💡 Answer</summary>

**C) break**

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Exit loop
    }
    fmt.Println(i)
}
// Output: 0 1 2 3 4
```

</details>

---

### Q43: What's the output?
```go
i := 0
for i < 3 {
    fmt.Print(i)
    i++
}
```

**A)** 0 1 2  
**B)** 012  
**C)** Infinite loop  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) 012**

While-style loop in Go.

</details>

---

## 📋 SECTION 6: CONDITIONS (7 Questions)

### Q44: Valid if statement:

**A)** `if x > 5 { }`  
**B)** `if(x > 5) { }`  
**C)** `if x > 5 then { }`  
**D)** All valid  

<details><summary>💡 Answer</summary>

**A) if x > 5 { }**

```go
// ✅ Correct:
if x > 5 {
    fmt.Println("big")
}

// ❌ Wrong: No parentheses needed
// if(x > 5) {

// ❌ Wrong: No 'then' keyword
// if x > 5 then {
```

</details>

---

### Q45: If with initialization:

**A)** `if x := 10; x > 5 { }`  
**B)** `if x = 10; x > 5 { }`  
**C)** Both valid  
**D)** Neither valid  

<details><summary>💡 Answer</summary>

**A) if x := 10; x > 5 { }**

```go
if x := 10; x > 5 {
    fmt.Println(x)  // x available here
}
// x not available here!
```

</details>

---

### Q46: What's the output?
```go
x := 10
if x > 5 {
    fmt.Print("A")
} else {
    fmt.Print("B")
}
```

**A)** A  
**B)** B  
**C)** AB  
**D)** Nothing  

<details><summary>💡 Answer</summary>

**A) A**

10 > 5 is true, so executes if block.

</details>

---

### Q47: Switch without condition:

**A)** Invalid  
**B)** Same as `switch true`  
**C)** Error  
**D)** Infinite loop  

<details><summary>💡 Answer</summary>

**B) Same as switch true**

```go
x := 10
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
case x > 0:
    fmt.Println("positive")
}
```

</details>

---

### Q48: Switch fallthrough:

**A)** Automatic by default  
**B)** Must use `fallthrough` keyword  
**C)** Not possible  
**D)** Use `break` to fall through  

<details><summary>💡 Answer</summary>

**B) Must use fallthrough keyword**

```go
x := 1
switch x {
case 1:
    fmt.Println("one")
    fallthrough  // Explicitly fall through
case 2:
    fmt.Println("two")
}
// Output:
// one
// two
```

</details>

---

### Q49: Multiple conditions in switch case:

**A)** Not possible  
**B)** `case 1, 2, 3:`  
**C)** `case 1 || 2 || 3:`  
**D)** `case 1: case 2: case 3:`  

<details><summary>💡 Answer</summary>

**B) case 1, 2, 3:**

```go
x := 2
switch x {
case 1, 2, 3:
    fmt.Println("1, 2, or 3")
case 4, 5:
    fmt.Println("4 or 5")
}
```

</details>

---

### Q50: What's the output?
```go
x := 5
switch x {
case 5:
    fmt.Print("A")
case 10:
    fmt.Print("B")
default:
    fmt.Print("C")
}
```

**A)** A  
**B)** AB  
**C)** ABC  
**D)** C  

<details><summary>💡 Answer</summary>

**A) A**

Matches first case, no automatic fallthrough in Go.

</details>

---

## 🎯 ANSWER KEY

**Section 1 (Quotes & Strings):** 1-B, 2-B, 3-A, 4-B, 5-C, 6-C, 7-B, 8-B, 9-B, 10-A

**Section 2 (ASCII):** 11-B, 12-D, 13-D, 14-B, 15-D

**Section 3 (Zero Values):** 16-B, 17-C, 18-B, 19-B, 20-B, 21-B, 22-B, 23-B

**Section 4 (Data Types):** 24-A, 25-C, 26-B, 27-D, 28-C, 29-D, 30-D, 31-C, 32-A, 33-D, 34-B, 35-C

**Section 5 (Loops):** 36-B, 37-A, 38-A, 39-B, 40-B, 41-B, 42-C, 43-B

**Section 6 (Conditions):** 44-A, 45-A, 46-A, 47-B, 48-B, 49-B, 50-A

---

## 📊 SCORING

**45-50:** 🏆 Master - Perfect understanding!  
**40-44:** ✅ Pass - Ready to move forward  
**35-39:** ⚠️ Review - Practice more  
**Below 35:** 🔄 Study again - Need solid foundation

---

## 💡 QUICK REFERENCE

### **Quotes:**
- `' '` = rune (single character, int32)
- `" "` = string (text)

### **ASCII:**
- '0' = 48
- 'A' = 65
- 'a' = 97
- Difference: 32

### **Zero Values:**
- int: 0
- string: ""
- bool: false
- pointer/slice/map: nil

### **Types:**
- **Signed:** int, int8, int16, int32, int64
- **Unsigned:** uint, uint8/byte, uint16, uint32, uint64
- **Float:** float32, float64
- **Complex:** complex64, complex128
- **Char:** rune (int32), byte (uint8)
- **Bool:** true, false

### **Loops:**
```go
for i := 0; i < 10; i++ { }  // C-style
for condition { }             // While-style
for { }                       // Infinite
for i, v := range collection { }  // Range
```

### **Conditions:**
```go
if condition { }
if x := value; condition { }  // With init
switch value { case x: }
switch { case condition: }    // No value
```

**Good luck! 🚀**