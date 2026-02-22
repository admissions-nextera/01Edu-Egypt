# Go Interfaces: Complete Crash Course 🚀

For someone learning interfaces for the **first time ever**!

---

## What is an Interface? (Simple Explanation)

Think of an interface as a **contract** or **promise**:

> "If you can do these specific things, I don't care what you are - you're good to go!"

**Real-world analogy:**
- A **power outlet** is like an interface
- It doesn't care if you plug in a phone, laptop, or lamp
- As long as your device has the right plug (implements the interface), it works!

---

## Part 1: The Problem Without Interfaces

### Scenario: You have different animals

```go
package main

import "fmt"

type Dog struct {
    Name string
}

type Cat struct {
    Name string
}

// Each animal makes a sound differently
func (d Dog) Bark() {
    fmt.Println(d.Name, "says: Woof!")
}

func (c Cat) Meow() {
    fmt.Println(c.Name, "says: Meow!")
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    
    dog.Bark()  // Works
    cat.Meow()  // Works
    
    // But what if we want to make ALL animals make a sound?
    // We'd need separate functions for each!
}
```

**Problem:** How do we write ONE function that works with ALL animals?

---

## Part 2: Interfaces to the Rescue!

### Step 1: Define the Interface (The Contract)

```go
// This interface says:
// "Anyone who has a MakeSound() method is an Animal"
type Animal interface {
    MakeSound()  // This is the contract - you MUST have this method
}
```

**Read it as:** "To be considered an Animal, you must be able to MakeSound()"

---

### Step 2: Implement the Interface (Fulfill the Contract)

```go
type Dog struct {
    Name string
}

type Cat struct {
    Name string
}

type Cow struct {
    Name string
}

// Dog implements Animal interface by having MakeSound() method
func (d Dog) MakeSound() {
    fmt.Println(d.Name, "says: Woof!")
}

// Cat implements Animal interface by having MakeSound() method
func (c Cat) MakeSound() {
    fmt.Println(c.Name, "says: Meow!")
}

// Cow implements Animal interface by having MakeSound() method
func (c Cow) MakeSound() {
    fmt.Println(c.Name, "says: Moo!")
}
```

**Key Point:** You don't need to write `implements Animal` - Go figures it out automatically! If you have the required methods, you implement the interface!

---

### Step 3: Use the Interface

```go
// This function accepts ANY Animal!
func MakeAnimalSpeak(a Animal) {
    a.MakeSound()  // We can call this because ALL Animals have MakeSound()
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    cow := Cow{Name: "Bessie"}
    
    // All of these work!
    MakeAnimalSpeak(dog)  // Buddy says: Woof!
    MakeAnimalSpeak(cat)  // Whiskers says: Meow!
    MakeAnimalSpeak(cow)  // Bessie says: Moo!
}
```

**Magic!** ✨ One function works with ALL animals!

---

## Part 3: The Rules of Interfaces

### Rule 1: Interfaces Define Behavior, Not Data

```go
// ✅ GOOD - defines what something can DO
type Speaker interface {
    Speak() string
}

// ❌ BAD - interfaces don't have fields
type BadInterface interface {
    Name string  // This won't compile!
    Speak() string
}
```

### Rule 2: Implicit Implementation

```go
type Writer interface {
    Write(data []byte) (int, error)
}

type FileWriter struct {
    filename string
}

// This automatically implements Writer interface!
// No need to declare it explicitly
func (f FileWriter) Write(data []byte) (int, error) {
    // Write to file
    return len(data), nil
}
```

### Rule 3: Empty Interface Accepts ANYTHING

```go
// interface{} or any (Go 1.18+) accepts ANY type
func PrintAnything(thing interface{}) {
    fmt.Println(thing)
}

func main() {
    PrintAnything(42)           // int
    PrintAnything("hello")      // string
    PrintAnything(true)         // bool
    PrintAnything([]int{1,2,3}) // slice
}
```

---

## Part 4: Real-World Examples

### Example 1: Shapes (Classic)

```go
package main

import (
    "fmt"
    "math"
)

// Interface: Any shape must be able to calculate area
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// This function works with ANY shape!
func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 10, Height: 5}
    
    PrintArea(circle)     // Area: 78.54
    PrintArea(rectangle)  // Area: 50.00
}
```

---

### Example 2: Payment Processing (Practical)

```go
package main

import "fmt"

// Any payment method must be able to process payments
type PaymentMethod interface {
    ProcessPayment(amount float64) bool
}

type CreditCard struct {
    Number string
    CVV    string
}

type PayPal struct {
    Email string
}

type Bitcoin struct {
    WalletAddress string
}

func (cc CreditCard) ProcessPayment(amount float64) bool {
    fmt.Printf("Charging $%.2f to credit card %s\n", amount, cc.Number)
    return true
}

func (pp PayPal) ProcessPayment(amount float64) bool {
    fmt.Printf("Charging $%.2f to PayPal account %s\n", amount, pp.Email)
    return true
}

func (btc Bitcoin) ProcessPayment(amount float64) bool {
    fmt.Printf("Sending $%.2f worth of Bitcoin to %s\n", amount, btc.WalletAddress)
    return true
}

// Checkout function works with ANY payment method!
func Checkout(amount float64, method PaymentMethod) {
    if method.ProcessPayment(amount) {
        fmt.Println("Payment successful!")
    } else {
        fmt.Println("Payment failed!")
    }
}

func main() {
    creditCard := CreditCard{Number: "1234-5678", CVV: "123"}
    paypal := PayPal{Email: "user@example.com"}
    bitcoin := Bitcoin{WalletAddress: "1A1zP1..."}
    
    Checkout(100.00, creditCard)
    Checkout(50.00, paypal)
    Checkout(25.00, bitcoin)
}
```

**Benefit:** Add new payment methods WITHOUT changing the `Checkout` function!

---

### Example 3: File Operations (Web Dev!)

```go
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

// io.Reader is an interface in Go's standard library
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

func ReadAndPrint(r io.Reader) {
    data := make([]byte, 100)
    n, _ := r.Read(data)
    fmt.Println(string(data[:n]))
}

func main() {
    // Read from file
    file, _ := os.Open("data.txt")
    ReadAndPrint(file)
    
    // Read from string
    stringReader := strings.NewReader("Hello from string!")
    ReadAndPrint(stringReader)
    
    // Read from HTTP response, network, etc.
    // ALL work because they implement io.Reader!
}
```

---

## Part 5: Multiple Interfaces

A type can implement MULTIPLE interfaces!

```go
type ReadWriter interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

type File struct {
    name string
}

// File implements ReadWriter
func (f File) Read(data []byte) (int, error) {
    // Read logic
    return 0, nil
}

func (f File) Write(data []byte) (int, error) {
    // Write logic
    return 0, nil
}

// File also implements Closer
func (f File) Close() error {
    // Close logic
    return nil
}

func main() {
    f := File{name: "test.txt"}
    
    var rw ReadWriter = f  // Works!
    var c Closer = f       // Also works!
}
```

---

## Part 6: Interface Composition

Interfaces can be built from other interfaces!

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Composed interface - combines all three!
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// To implement ReadWriteCloser, you need ALL three methods:
// Read(), Write(), Close()
```

---

## Part 7: Type Assertions (Getting the Real Type Back)

```go
func main() {
    var animal Animal = Dog{Name: "Buddy"}
    
    // Type assertion: "I know this is really a Dog!"
    dog, ok := animal.(Dog)
    if ok {
        fmt.Println("It's a dog named:", dog.Name)
    }
    
    // Type switch: Check multiple types
    switch v := animal.(type) {
    case Dog:
        fmt.Println("Woof! I'm", v.Name)
    case Cat:
        fmt.Println("Meow! I'm", v.Name)
    default:
        fmt.Println("Unknown animal")
    }
}
```

---

## Part 8: Common Go Interfaces (You'll Use These!)

### 1. `fmt.Stringer` - Custom Print Format

```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p)  // Automatically calls String() method!
    // Output: Alice (30 years old)
}
```

### 2. `error` - Go's Error Handling

```go
// error is an interface!
type error interface {
    Error() string
}

type MyError struct {
    Code    int
    Message string
}

func (e MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return MyError{Code: 404, Message: "Not Found"}
}
```

### 3. `io.Reader` and `io.Writer` - File/Network Operations

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Used EVERYWHERE in Go:
// - Files
// - HTTP requests/responses
// - Network connections
// - Buffers
```

---

## Part 9: When to Use Interfaces?

### ✅ USE interfaces when:

1. **Multiple types do the same thing differently**
   ```go
   type Database interface {
       Save(data interface{}) error
   }
   // MySQL, PostgreSQL, MongoDB all implement differently
   ```

2. **You want flexible, testable code**
   ```go
   type EmailSender interface {
       Send(to, subject, body string) error
   }
   // Use real sender in production, mock in tests
   ```

3. **Working with standard library**
   ```go
   func ProcessData(r io.Reader) {
       // Works with files, HTTP, strings, anything!
   }
   ```

### ❌ DON'T use interfaces when:

1. **Only one implementation exists (and will likely stay that way)**
2. **You're adding unnecessary abstraction**
3. **It makes code harder to understand**

**Go proverb:** *"The bigger the interface, the weaker the abstraction"*

Keep interfaces **small and focused**!

---

## Part 10: Practice Exercises

### Exercise 1: Vehicle Interface
```go
// Create Vehicle interface with methods:
// - Start()
// - Stop()
// - GetSpeed() int

// Implement for Car, Bike, Truck
// Write a function that works with any Vehicle
```

### Exercise 2: Storage Interface
```go
// Create Storage interface:
// - Save(key, value string) error
// - Get(key string) (string, error)
// - Delete(key string) error

// Implement for:
// - MemoryStorage (map)
// - FileStorage (write to files)
```

### Exercise 3: Notification Interface
```go
// Create Notifier interface:
// - Send(message string) error

// Implement for:
// - EmailNotifier
// - SMSNotifier
// - PushNotifier

// Write function that sends to all notifiers
```

---

## Quick Reference Card

```go
// Define interface
type MyInterface interface {
    Method1() string
    Method2(int) error
}

// Implement interface (implicit)
type MyType struct {}

func (m MyType) Method1() string {
    return "hello"
}

func (m MyType) Method2(n int) error {
    return nil
}

// Use interface
func UseInterface(m MyInterface) {
    m.Method1()
    m.Method2(42)
}

// Type assertion
concrete := myInterface.(MyType)

// Type switch
switch v := myInterface.(type) {
case MyType:
    // use v as MyType
default:
    // unknown type
}
```

---

## Key Takeaways

1. **Interfaces define behavior** (methods), not data (fields)
2. **Implementation is implicit** - no keywords needed
3. **Interfaces enable polymorphism** - write code that works with many types
4. **Keep interfaces small** - prefer many small interfaces over large ones
5. **Accept interfaces, return structs** - common Go idiom

---