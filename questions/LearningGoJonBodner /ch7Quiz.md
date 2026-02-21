# 📘 Learning Go — Chapter 7 Quiz
## Types, Methods, and Interfaces

**Questions:** 28 | **Time:** 40 minutes | **Passing Score:** 22/28 (79%)

---

## TYPES & METHODS

### Q1: What does this code do?
```go
type Celsius float64
type Fahrenheit float64
```

**A)** Creates aliases for `float64` — they are interchangeable  
**B)** Creates distinct named types based on `float64` — they are NOT interchangeable without explicit conversion  
**C)** Wraps `float64` in a struct  
**D)** Creates interface types  

<details><summary>💡 Answer</summary>

**B) Distinct named types — not interchangeable**

```go
var c Celsius = 100
var f Fahrenheit = Fahrenheit(c)  // explicit conversion required
// var f Fahrenheit = c            // COMPILE ERROR
```

Named types allow you to add domain-specific meaning and methods to existing types. `Celsius` and `Fahrenheit` are different types even though both are backed by `float64`.

</details>

---

### Q2: What is the method set of a type `T` vs `*T`?

**A)** Identical — Go handles this automatically  
**B)** `T`'s method set contains only methods with value receivers; `*T`'s method set contains methods with both value AND pointer receivers  
**C)** `*T` contains only pointer receiver methods  
**D)** `T`'s method set contains all methods regardless of receiver type  

<details><summary>💡 Answer</summary>

**B) T = value receivers only; *T = value receivers + pointer receivers**

```go
type Counter struct { n int }
func (c Counter) Value() int  { return c.n }   // value receiver
func (c *Counter) Inc()       { c.n++ }         // pointer receiver

var c Counter
c.Value()   // OK — c has value receiver methods
c.Inc()     // OK — Go auto-takes address: (&c).Inc()

var p *Counter = &Counter{}
p.Value()   // OK — *Counter has both
p.Inc()     // OK
```

This matters for interface satisfaction.

</details>

---

### Q3: Can you define methods on types from other packages?

**A)** Yes — you can add methods to any type  
**B)** No — you can only define methods on types declared in the same package  
**C)** Yes, but only value receivers  
**D)** Only for exported types  

<details><summary>💡 Answer</summary>

**B) Methods only on types in the same package**

```go
// COMPILE ERROR — int is from builtin, not your package:
func (i int) Double() int { return i * 2 }

// Workaround — create your own type:
type MyInt int
func (i MyInt) Double() MyInt { return i * 2 }
```

This prevents external packages from adding surprising behavior to types they don't own.

</details>

---

## INTERFACES

### Q4: How does a type "implement" an interface in Go?

**A)** By declaring `implements InterfaceName` in the type definition  
**B)** Implicitly — if a type has all the methods in the interface's method set, it satisfies the interface  
**C)** By embedding the interface  
**D)** By registering with the compiler  

<details><summary>💡 Answer</summary>

**B) Implicit — no declaration needed**

```go
type Stringer interface { String() string }

type Person struct { Name string }
func (p Person) String() string { return p.Name }

var s Stringer = Person{"Alice"}  // Person implements Stringer automatically
```

This is structural typing ("duck typing" but checked at compile time). No `implements` keyword. This allows retrofitting: a type from another package can satisfy your interface without modification.

</details>

---

### Q5: What is the zero value of an interface variable?

**A)** An empty struct  
**B)** `nil` — a nil interface holds no type and no value  
**C)** An empty interface value  
**D)** `0`  

<details><summary>💡 Answer</summary>

**B) `nil`**

```go
var s Stringer  // s == nil
s.String()      // PANIC — nil interface
```

A nil interface has no concrete type. Any method call on a nil interface panics. Always check `if s != nil` before calling methods on interface variables that might be nil.

</details>

---

### Q6: What is the infamous "nil interface" trap?
```go
type MyError struct{ msg string }
func (e *MyError) Error() string { return e.msg }

func getError() error {
    var err *MyError = nil
    return err
}

e := getError()
fmt.Println(e == nil)
```

**A)** `true` — `err` is nil  
**B)** `false` — the interface holds a non-nil type (`*MyError`) with a nil value  
**C)** Panic  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) `false` — non-nil interface with nil concrete value**

An interface value has two parts: (type, value). When `err` is returned, the interface gets type=`*MyError`, value=`nil`. Even though the value is nil, the interface is NOT nil because it has a concrete type. This is a famous Go gotcha. Fix: return `nil` directly (not a typed nil pointer) when you want a nil interface.

```go
func getError() error {
    // ...
    return nil  // correct — returns nil interface
}
```

</details>

---

### Q7: What is `interface{}` (or `any` in Go 1.18+)?

**A)** An error type  
**B)** An interface with no methods — every type satisfies it, so it can hold any value  
**C)** A generic type  
**D)** A special string type  

<details><summary>💡 Answer</summary>

**B) The empty interface — satisfied by every type**

```go
var x any = 5
x = "hello"
x = []int{1, 2, 3}
```

`any` (alias for `interface{}`) accepts any value. To use the stored value, you need a type assertion. The book warns: use `any` sparingly — it bypasses type safety.

</details>

---

### Q8: What is a type assertion and when does it panic?

**A)** `x.(T)` — asserts x holds a value of type T; panics if it doesn't  
**B)** `x.T` — field access on an interface  
**C)** `T(x)` — type conversion  
**D)** `assert(x, T)` — built-in function  

<details><summary>💡 Answer</summary>

**A) `x.(T)` — panics if x doesn't hold type T**

```go
var i any = "hello"
s := i.(string)     // OK — s is "hello"
n := i.(int)        // PANIC — i doesn't hold int

// Safe form — no panic:
s, ok := i.(string) // ok=true, s="hello"
n, ok := i.(int)    // ok=false, n=0
```

Always use the comma-ok form when the concrete type is not certain.

</details>

---

### Q9: What is a type switch?

**A)** A switch over type names as strings  
**B)** A `switch` using `x.(type)` — each case handles a different concrete type; a cleaner alternative to multiple type assertions  
**C)** A switch that only works with interface types  
**D)** A compile-time construct  

<details><summary>💡 Answer</summary>

**B) `switch x.(type)` — handles different concrete types**

```go
func describe(i any) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    case bool:
        fmt.Printf("bool: %t\n", v)
    default:
        fmt.Printf("unknown: %T\n", v)
    }
}
```

`v` in each case is the concrete type — no assertion needed. Type switches are the idiomatic way to handle multiple concrete types behind an interface.

</details>

---

### Q10: What is the `Stringer` interface from `fmt` and why does it matter?

**A)** An interface for string parsing  
**B)** `interface { String() string }` — if your type implements it, `fmt.Println` and related functions use your `String()` method for output  
**C)** A required interface for all Go types  
**D)** An interface in the `strings` package  

<details><summary>💡 Answer</summary>

**B) `String() string` — controls `fmt` output**

```go
type Temperature struct { Celsius float64 }

func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°C", t.Celsius)
}

t := Temperature{100}
fmt.Println(t)  // "100.0°C" — uses your String() method
```

Implementing `fmt.Stringer` is the Go way to control how your type appears when printed.

</details>

---

### Q11: What does the `error` interface look like?

**A)** `interface { Error() error }`  
**B)** `interface { Error() string }`  
**C)** `interface { Err() string }`  
**D)** `struct { Message string }`  

<details><summary>💡 Answer</summary>

**B) `interface { Error() string }`**

Any type with an `Error() string` method satisfies the `error` interface. This is why you can create custom error types:

```go
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}
```

</details>

---

### Q12: Can a type satisfy multiple interfaces?

**A)** No — a type can only satisfy one interface  
**B)** Yes — a type satisfies every interface whose method set is a subset of the type's method set  
**C)** Yes, but you must declare it explicitly  
**D)** Only if the interfaces are in the same package  

<details><summary>💡 Answer</summary>

**B) Yes — automatically satisfies all matching interfaces**

```go
type File struct{}
func (f *File) Read(b []byte) (int, error) { ... }
func (f *File) Write(b []byte) (int, error) { ... }
func (f *File) Close() error { ... }

// *File satisfies: io.Reader, io.Writer, io.Closer, io.ReadWriter, io.ReadCloser...
```

This is why Go's small interface design is so powerful — types accumulate interface satisfaction organically.

</details>

---

### Q13: The book strongly recommends keeping interfaces small. What is the ideal size?

**A)** As many methods as needed  
**B)** One or two methods — small interfaces are more composable and easier to implement  
**C)** At least 5 methods to be useful  
**D)** Exactly 3 methods  

<details><summary>💡 Answer</summary>

**B) One or two methods**

`io.Reader` has 1 method. `io.Writer` has 1. `io.Closer` has 1. They compose into `io.ReadWriter`, `io.ReadCloser`, etc. The book quotes Rob Pike: "The bigger the interface, the weaker the abstraction." A 20-method interface is almost impossible to mock or satisfy with alternative implementations.

</details>

---

### Q14: Interfaces are defined by the CONSUMER not the PRODUCER. What does this mean?

**A)** You must define interfaces before writing the types that implement them  
**B)** Define the interface where it's used (as a function parameter), not where the concrete type is defined — this allows loose coupling and easier testing  
**C)** Interfaces must be in the `main` package  
**D)** Interfaces are automatically generated from struct methods  

<details><summary>💡 Answer</summary>

**B) Define interfaces at the point of use — loose coupling**

```go
// In your handler package (consumer):
type DataStore interface { Get(id int) (Item, error) }

// In your main or wire-up code:
handler := NewHandler(myConcreteDB)  // *DB satisfies DataStore automatically
```

This means: `*DB` doesn't import or know about `DataStore`. You can substitute any type that has `Get(int) (Item, error)`. This makes testing easy — just create a mock that has that method.

</details>

---

### Q15: What is embedding an interface in a struct?

**A)** Implementing all interface methods automatically  
**B)** Including an interface as a field — the struct's method set includes all the interface's methods; useful for mocking and wrapping  
**C)** Inheriting from the interface  
**D)** A compile error  

<details><summary>💡 Answer</summary>

**B) Interface field — promotes methods; useful for wrapping**

```go
type LoggingReader struct {
    io.Reader        // embedded interface
    log *log.Logger
}

func (lr LoggingReader) Read(b []byte) (int, error) {
    n, err := lr.Reader.Read(b)  // delegate to wrapped reader
    lr.log.Printf("read %d bytes", n)
    return n, err
}
```

This is the decorator pattern in Go — wrap an interface, override some methods, delegate the rest.

</details>

---

### Q16: What is the output?
```go
type Animal interface { Sound() string }
type Dog struct{}
func (d Dog) Sound() string { return "woof" }

func makeSound(a Animal) { fmt.Println(a.Sound()) }

makeSound(Dog{})
```

**A)** Compile error — must pass `&Dog{}`  
**B)** `woof`  
**C)** `Dog`  
**D)** Empty string  

<details><summary>💡 Answer</summary>

**B) `woof`**

`Dog` (not `*Dog`) satisfies `Animal` because `Sound()` has a value receiver. Value types with only value receivers can be passed directly as interfaces without taking their address.

</details>

---

### Q17: What is wrong with this code?
```go
type Writer interface { Write([]byte) (int, error) }

type MyWriter struct{}
func (w *MyWriter) Write(b []byte) (int, error) { return len(b), nil }

var wr Writer = MyWriter{}  // uses MyWriter, not *MyWriter
```

**A)** Nothing — both `MyWriter` and `*MyWriter` satisfy `Writer`  
**B)** Compile error — `Write` has a pointer receiver; only `*MyWriter` satisfies `Writer`, not `MyWriter`  
**C)** Runtime panic  
**D)** Works but is less efficient  

<details><summary>💡 Answer</summary>

**B) Compile error — value type doesn't have pointer receiver methods**

`MyWriter` only has the method set of value receivers. Since `Write` is defined on `*MyWriter`, only `*MyWriter` satisfies `Writer`. Fix: `var wr Writer = &MyWriter{}`.

This is one of the most common Go beginner errors. Remember: value type = value receiver methods only; pointer type = both.

</details>

---

### Q18: When does the book say to use `any` (or `interface{}`)?

**A)** Whenever you don't know the type  
**B)** Sparingly — only when there's genuinely no better alternative; it gives up compile-time type safety  
**C)** For all function parameters  
**D)** Never — it's deprecated  

<details><summary>💡 Answer</summary>

**B) Sparingly — last resort**

The book says: if you find yourself reaching for `any`, ask whether a defined interface, generics (Chapter 8 in 2nd edition), or a concrete type would work instead. `any` loses type safety — you get runtime panics instead of compile errors. Legitimate uses: serialization/deserialization, very generic library code.

</details>

---

### Q19: What is an interface guard and why is it useful?

**A)** A runtime check that a type implements an interface  
**B)** A compile-time check: `var _ InterfaceName = (*MyType)(nil)` — if `*MyType` doesn't satisfy `InterfaceName`, the program won't compile  
**C)** A mutex that protects interface access  
**D)** A type that wraps an interface to add safety  

<details><summary>💡 Answer</summary>

**B) Compile-time interface satisfaction check**

```go
// Verify at compile time that *MyWriter implements io.Writer:
var _ io.Writer = (*MyWriter)(nil)
```

This line creates a nil `*MyWriter` and assigns it to a blank `io.Writer` variable. If the interface isn't satisfied, you get a compile error immediately — before any tests run. Put these at the top of files that define types meant to implement specific interfaces.

</details>

---

### Q20: What are "implicit interfaces" and how do they differ from Java's explicit `implements`?

**A)** They are the same concept  
**B)** In Go, a type satisfies an interface just by having the right methods — no `implements` declaration. In Java, you must explicitly declare `class Foo implements Bar`.  
**C)** Implicit means the compiler does it automatically without methods  
**D)** Only exported interfaces are implicit  

<details><summary>💡 Answer</summary>

**B) Structural typing vs nominal typing**

Go: structural — "if it walks like a duck..." Java: nominal — "I declare this IS a duck." The Go approach allows: existing types to satisfy new interfaces without modification, cross-package interface satisfaction, easy mocking without a mocking framework. The Java approach requires every implementor to know about the interface at definition time.

</details>

---

### Q21: What is the output?
```go
var s fmt.Stringer
fmt.Println(s == nil)
```

**A)** `false`  
**B)** `true`  
**C)** Compile error  
**D)** Panic  

<details><summary>💡 Answer</summary>

**B) `true`**

`var s fmt.Stringer` declares a nil interface — no concrete type, no value. A nil interface equals `nil`. Compare with Q6 where the interface is non-nil despite holding a nil concrete value.

</details>

---

### Q22: Why can't you use `==` to compare two interface values in general?

**A)** You can — all interface values support `==`  
**B)** You can compare interfaces only if the underlying concrete type is comparable. If the concrete type is a slice or map, the comparison panics at runtime  
**C)** Interface comparison is always a compile error  
**D)** Only nil interface comparison works  

<details><summary>💡 Answer</summary>

**B) Comparison depends on concrete type — slice/map underlying types panic**

```go
var a, b any = []int{1, 2}, []int{1, 2}
fmt.Println(a == b)  // RUNTIME PANIC — slice is not comparable
```

The compiler can't always know the concrete type at compile time. Runtime comparison panics if the underlying type is not comparable. Use `reflect.DeepEqual` for deep comparison of potentially non-comparable types.

</details>

---

### Q23: The book introduces the concept of "accept interfaces, return structs." What does this mean?

**A)** Functions should always return interfaces for flexibility  
**B)** Function parameters should be interfaces (for flexibility and testability); return types should be concrete types (so callers know exactly what they get)  
**C)** Structs should embed interfaces  
**D)** Return interfaces for polymorphism  

<details><summary>💡 Answer</summary>

**B) Parameters = interfaces; return types = concrete**

```go
// Good: accept interface, return concrete
func Process(r io.Reader) (*Result, error) { ... }

// Problematic: return interface — caller must type-assert to use concrete features
func NewReader() io.Reader { return &myReader{} }
```

Returning concrete types gives callers access to the full API. Accepting interfaces makes the function reusable with any compatible type. This is one of the most important Go API design principles.

</details>

---

### Q24: What is embedding one interface in another?

**A)** A compile error  
**B)** Creating a composed interface whose method set is the union of all embedded interfaces  
**C)** An alias for the embedded interface  
**D)** Only valid in the standard library  

<details><summary>💡 Answer</summary>

**B) Union of method sets**

```go
type Reader interface { Read([]byte) (int, error) }
type Writer interface { Write([]byte) (int, error) }

type ReadWriter interface {
    Reader  // embedded — includes Read method
    Writer  // embedded — includes Write method
}
```

A type satisfies `ReadWriter` only if it implements both `Read` and `Write`. This is how the `io` package builds complex interfaces from simple ones.

</details>

---

### Q25: Can methods have types from any package as a receiver?

**A)** Yes — any type can have methods  
**B)** No — the receiver type must be defined in the same package as the method  
**C)** Only exported types can have methods  
**D)** Only struct types can have methods  

<details><summary>💡 Answer</summary>

**B) Same package only**

```go
// Your package defines MyInt:
type MyInt int
func (m MyInt) Double() MyInt { return m * 2 }  // OK

// CANNOT add methods to int from builtin package:
func (i int) Double() int { return i * 2 }  // COMPILE ERROR
```

You CAN create a new type based on `int` in your package and add methods to that.

</details>

---

### Q26: What is the difference between a concrete type and an interface type?

**A)** No difference — both can be instantiated  
**B)** Concrete types specify both structure and behavior; interface types specify only behavior (method signatures). You can create values of concrete types; interfaces are used as parameter/variable types.  
**C)** Interface types can hold any value  
**D)** Concrete types are faster  

<details><summary>💡 Answer</summary>

**B) Concrete = structure + behavior; interface = behavior only**

`struct Point{X, Y int}` is concrete — it has fields, can be instantiated, memory layout defined. `interface Stringer { String() string }` specifies only the contract. You can't do `Stringer{...}` — you use `Stringer` as a variable type and store concrete values in it.

</details>

---

### Q27: What is wrong with a function that returns `interface{}` for all its return types?

**A)** Nothing — maximum flexibility  
**B)** It forces all callers to type-assert every return value, giving up compile-time type safety. If the assertion is wrong, the caller gets a runtime panic.  
**C)** Compile error  
**D)** Performance issue  

<details><summary>💡 Answer</summary>

**B) Runtime type assertions replace compile-time checks**

```go
func getAge() interface{} { return 30 }  // bad

age := getAge().(int)   // could panic if implementation changes to return string
age := getAge().(string) // panics — always have to know the type anyway
```

If you always know the type, just return the concrete type. If the type genuinely varies, a typed interface with methods is better than `interface{}`.

</details>

---

### Q28: The book says "Go encourages small interfaces and discourages large ones." What is the practical benefit for testing?

**A)** Tests are smaller  
**B)** Small interfaces are easy to mock — implement a mock with just 1-2 methods instead of 20  
**C)** Only for unit testing  
**D)** Large interfaces can't be tested  

<details><summary>💡 Answer</summary>

**B) Small interfaces = easy mocking**

```go
type DataFetcher interface { Fetch(id int) (Data, error) }

// In tests — trivial to implement:
type mockFetcher struct{ data Data; err error }
func (m mockFetcher) Fetch(int) (Data, error) { return m.data, m.err }
```

A 20-method interface would require a 20-method mock. A 1-method interface requires a 1-method mock. The book emphasizes this as the key practical benefit of small interfaces.

</details>

---

## 📊 Score

| Score | Result |
|---|---|
| 26–28 ✅ | **Excellent.** Deep interface understanding — ready for errors and concurrency. |
| 22–25 ✅ | **Ready.** Review the nil interface trap and pointer receiver rules. |
| 17–21 ⚠️ | **Study the nil interface trap (Q6) and pointer vs value receiver interface satisfaction (Q17).** |
| Below 17 ❌ | **Reread Chapter 7 carefully — interfaces are the foundation of Go design.** |