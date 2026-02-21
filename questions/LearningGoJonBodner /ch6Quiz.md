# 📘 Learning Go — Chapter 6 Quiz
## Pointers

**Questions:** 20 | **Time:** 28 minutes | **Passing Score:** 16/20 (80%)

---

### Q1: What is a pointer?

**A)** A type that points to another type in the package hierarchy  
**B)** A variable that holds the memory address of another variable  
**C)** A reference to a function  
**D)** A type alias  

<details><summary>💡 Answer</summary>

**B) A variable holding a memory address**

```go
x := 5
p := &x       // p holds the address of x; type: *int
fmt.Println(*p) // dereference: read the value at the address → 5
*p = 10         // write to the address → x is now 10
```

`&` = "address of". `*` = "dereference" (get/set the value at that address). `*int` = "pointer to int".

</details>

---

### Q2: What is the zero value of a pointer?

**A)** `0`  
**B)** An address pointing to zero memory  
**C)** `nil`  
**D)** An empty string  

<details><summary>💡 Answer</summary>

**C) `nil`**

A nil pointer holds no address. Dereferencing a nil pointer causes a runtime panic. Always check `if p != nil` before dereferencing a pointer that might be nil.

```go
var p *int
fmt.Println(p)   // <nil>
fmt.Println(*p)  // PANIC: runtime error: invalid memory address or nil pointer dereference
```

</details>

---

### Q3: What is the output?
```go
x := 5
p := &x
*p = 10
fmt.Println(x)
```

**A)** `5` — `*p = 10` only changes the pointer  
**B)** `10` — `*p = 10` writes through the pointer to `x`  
**C)** Compile error  
**D)** Address of x  

<details><summary>💡 Answer</summary>

**B) `10`**

`*p = 10` dereferences `p` and writes `10` to the memory location `p` points to — which is `x`. Pointers allow indirect mutation of variables.

</details>

---

### Q4: When should you use a pointer receiver vs a value receiver for a method?

**A)** Always pointer — it's faster  
**B)** Pointer when the method must modify the receiver, or when the receiver is large enough that copying is expensive; value for read-only operations on small types  
**C)** Value when the type has exported fields; pointer otherwise  
**D)** It makes no difference  

<details><summary>💡 Answer</summary>

**B) Pointer for mutation or large types; value for small read-only**

The book's rule: if ANY method on a type needs a pointer receiver, use pointer receivers for ALL methods on that type — consistency prevents confusion.

```go
func (c *Counter) Increment() { c.count++ }  // needs pointer — mutates
func (c Counter) Value() int  { return c.count }  // value is fine — read-only
// But consistency rule: prefer *Counter for both
```

</details>

---

### Q5: What is the difference between passing a struct by value and passing a pointer to it?

**A)** No difference — Go handles it transparently  
**B)** By value: a copy is made — the function cannot mutate the caller's struct. By pointer: the function can mutate the caller's data and avoids the copy overhead for large structs.  
**C)** Pointers are always slower  
**D)** Only structs with exported fields can be passed by pointer  

<details><summary>💡 Answer</summary>

**B) Value = copy (safe, no mutation); pointer = shared access (mutation, no copy)**

```go
func scaleValue(r Rectangle, f float64) {  // copy — caller's rect unchanged
    r.Width *= f
}

func scalePointer(r *Rectangle, f float64) {  // mutates caller's rect
    r.Width *= f
}
```

Choose based on whether mutation is intended and whether the copy cost matters.

</details>

---

### Q6: What does `new(int)` return?

**A)** `int` with value 0  
**B)** `*int` pointing to a zero-valued `int`  
**C)** `nil`  
**D)** An error if memory allocation fails  

<details><summary>💡 Answer</summary>

**B) `*int` pointing to a zeroed `int`**

```go
p := new(int)   // *int, *p == 0
*p = 5
```

`new(T)` is equivalent to `&T{}` for structs. The book says `new` is rarely used — `&T{}` is more idiomatic because it can also set initial field values.

</details>

---

### Q7: The book says Go does NOT have pointer arithmetic. What does this mean and why?

**A)** You cannot do math with pointer values (no `p++` to advance to the next element)  
**B)** Pointers cannot point to integers  
**C)** All pointer operations require the unsafe package  
**D)** Pointers are automatically garbage-collected  

<details><summary>💡 Answer</summary>

**A) No `p++` or arithmetic on pointer addresses**

In C, you can do `p++` to advance a pointer to the next array element. Go disallows this — it prevents entire classes of memory safety bugs (buffer overruns, out-of-bounds access). If you need pointer arithmetic, use `unsafe.Pointer` (rarely appropriate).

</details>

---

### Q8: What is the output?
```go
func increment(n *int) { *n++ }

x := 5
increment(&x)
fmt.Println(x)
```

**A)** `5` — `n` is a copy of `x`  
**B)** `6` — `*n++` modifies `x` through the pointer  
**C)** Compile error  
**D)** Address of x  

<details><summary>💡 Answer</summary>

**B) `6`**

`&x` passes the address of `x`. Inside `increment`, `*n++` dereferences and increments the value at that address — which is `x`. This is how functions mutate their caller's variables in Go.

</details>

---

### Q9: Can you take the address of a literal value like `&5`?

**A)** Yes — `p := &5` creates a pointer to 5  
**B)** No — you cannot take the address of a literal value in Go  
**C)** Only for string literals  
**D)** Only inside functions  

<details><summary>💡 Answer</summary>

**B) Cannot take address of a literal**

```go
p := &5     // COMPILE ERROR
p := &true  // COMPILE ERROR
```

Literals don't have addressable memory locations. A common workaround when you need a `*int` with value 5:

```go
func intPtr(i int) *int { return &i }  // helper
p := intPtr(5)
```

Or: `x := 5; p := &x`

</details>

---

### Q10: What is the output?
```go
type Point struct{ X, Y int }

func moveRight(p *Point) { p.X++ }

pt := Point{1, 2}
moveRight(&pt)
fmt.Println(pt)
```

**A)** `{1 2}` — the copy is moved  
**B)** `{2 2}` — `p.X++` modifies the original through the pointer  
**C)** Compile error  
**D)** `&{2 2}`  

<details><summary>💡 Answer</summary>

**B) `{2 2}`**

`p.X++` is shorthand for `(*p).X++` — Go automatically dereferences the pointer when accessing struct fields. This is syntactic sugar: you can write `p.X` instead of `(*p).X` when `p` is a pointer to a struct.

</details>

---

### Q11: What is the difference between `*p` and `**p`?

**A)** No difference  
**B)** `*p` dereferences once (pointer to T); `**p` dereferences twice (pointer to pointer to T)  
**C)** `**p` is not valid Go  
**D)** `*p` is for reading; `**p` is for writing  

<details><summary>💡 Answer</summary>

**B) `*p` = one level; `**p` = two levels**

```go
x := 5
p := &x   // *int
pp := &p  // **int

fmt.Println(*p)   // 5 — dereference once
fmt.Println(**pp) // 5 — dereference twice
**pp = 10         // modifies x through two pointer levels
```

Pointers to pointers are uncommon but appear in some patterns (e.g., linked lists, tree nodes).

</details>

---

### Q12: What does the book mean by "pointer semantics" vs "value semantics"?

**A)** Pointer semantics means the type is stored as a pointer; value semantics means stored as a value — the key difference is whether copying creates a share or an independent copy  
**B)** Pointer semantics is slower  
**C)** Only applies to built-in types  
**D)** Pointer semantics is required for concurrency  

<details><summary>💡 Answer</summary>

**A) Semantics determines sharing behavior**

Value semantics: copying creates an independent copy (int, struct). Changes don't affect the original. Pointer semantics: copying the pointer shares the data (map, slice header, *T). Changes through any copy affect all. The book says: understand which semantics your type uses — it determines behavior in function calls and assignments.

</details>

---

### Q13: A function returns a pointer to a local variable. Is this safe in Go?

**A)** No — local variables are destroyed when the function returns; the pointer becomes dangling  
**B)** Yes — Go's garbage collector detects this and moves the variable to the heap so it remains valid  
**C)** Only for primitive types  
**D)** Only if the variable is declared with `new`  

<details><summary>💡 Answer</summary>

**B) Safe — Go's escape analysis moves the variable to the heap**

```go
func newInt() *int {
    x := 5      // Go detects this escapes to heap
    return &x   // safe — x lives until GC collects it
}
```

"Escape analysis" is how Go decides: if a variable's address is returned or stored somewhere that outlives the function, Go allocates it on the heap. No dangling pointers like in C. The book discusses this — it's a key safety feature.

</details>

---

### Q14: What is the danger of passing a `nil` pointer to a function?

**A)** The function receives a zero-value struct  
**B)** Any dereference of the pointer inside the function causes a runtime panic  
**C)** The pointer is automatically converted to a zero-value struct  
**D)** The compiler catches all nil dereferences  

<details><summary>💡 Answer</summary>

**B) Runtime panic on dereference**

```go
func process(p *Person) {
    fmt.Println(p.Name)  // PANIC if p == nil
}

process(nil)  // passes — the panic happens inside
```

The function could check: `if p == nil { return }`. The compiler cannot catch nil dereferences at compile time — they're runtime errors.

</details>

---

### Q15: When should you prefer returning a value vs returning a pointer from a function?

**A)** Always return pointers — avoids copying  
**B)** Return a value when the function creates small data that doesn't need shared mutation; return a pointer when the data is large, needs to express "absence" (nil), or must be mutated by callers  
**C)** Return pointers only for structs  
**D)** The caller decides — both always work  

<details><summary>💡 Answer</summary>

**B) Value for small/immutable; pointer for large/nil-able/mutable**

The book's guidance: returning a pointer is a signal that the data is meant to be shared and possibly mutated. Returning a value means "here's your own copy." For types that can legitimately be absent, pointers allow `nil` as a sentinel — values cannot represent absence without a separate boolean.

</details>

---

### Q16: What does `*` mean in a type declaration like `var p *int`?

**A)** Multiply  
**B)** "Pointer to" — `*int` is the type "pointer to int"  
**C)** Dereference  
**D)** Optional value  

<details><summary>💡 Answer</summary>

**B) "Pointer to" in a type position**

`*` has three distinct uses in Go:
1. In type position: `*int` means "pointer to int"
2. On the left side of assignment: `*p = 5` means "write through pointer"
3. In expression position: `*p` means "read through pointer"

Context determines meaning. This takes getting used to.

</details>

---

### Q17: What is the output?
```go
a := []int{1, 2, 3}
b := a
b[0] = 99
fmt.Println(a[0])
```

**A)** `1` — slices are value types  
**B)** `99` — slice assignment copies the header (pointer), so both share the backing array  
**C)** Compile error  
**D)** `0`  

<details><summary>💡 Answer</summary>

**B) `99` — slices use pointer semantics**

Slices contain a pointer to a backing array. Assigning a slice copies the pointer — both `a` and `b` point to the same array. This is pointer semantics without an explicit `*`. The book covers this to illustrate that value vs pointer semantics applies to composite types too.

</details>

---

### Q18: What is `unsafe.Pointer`?

**A)** A pointer type that bypasses Go's type safety — allows pointer arithmetic and conversion between unrelated pointer types; should almost never be used in application code  
**B)** A pointer that automatically catches nil dereferences  
**C)** A pointer for use in `unsafe` packages only  
**D)** A deprecated type  

<details><summary>💡 Answer</summary>

**A) Bypasses type safety — very rarely used**

`unsafe.Pointer` can be converted to/from any pointer type and to `uintptr` (allowing arithmetic). It's used in the `sync`, `reflect`, and `syscall` packages for low-level operations. The book says: if you're using `unsafe`, you have a very specialized need and must understand the exact guarantees you're giving up.

</details>

---

### Q19: What is the output?
```go
x := 42
p1 := &x
p2 := &x
fmt.Println(p1 == p2)
```

**A)** `false` — `p1` and `p2` are different pointers  
**B)** `true` — both pointers hold the same address (address of `x`)  
**C)** Compile error  
**D)** `false` — pointer equality compares the pointed-to values  

<details><summary>💡 Answer</summary>

**B) `true` — same address**

Pointer equality compares the addresses they hold, not the values they point to. Both `p1` and `p2` hold the address of `x`, so they're equal. To compare pointed-to values: `*p1 == *p2`.

</details>

---

### Q20: The book introduces pointers after functions and before interfaces. Why is understanding pointers important for interfaces?

**A)** Interfaces are implemented with pointers internally  
**B)** Methods on pointer receivers vs value receivers determine whether a pointer type or value type satisfies an interface — confusing this leads to compile errors  
**C)** Interfaces require pointer parameters  
**D)** Pointers must be used when calling interface methods  

<details><summary>💡 Answer</summary>

**B) Pointer vs value receivers determine interface satisfaction**

A value of type `T` automatically has the method set of `T`. A value of type `*T` has the method set of both `T` and `*T`. If an interface requires a method that's defined with a pointer receiver, only `*T` (not `T`) satisfies the interface. This is covered in Chapter 7 and is much clearer once pointer semantics are understood.

</details>

---

## 📊 Score

| Score | Result |
|---|---|
| 19–20 ✅ | **Excellent.** Pointer semantics are clear. |
| 16–18 ✅ | **Ready for interfaces — review escape analysis and nil pointer risks.** |
| 12–15 ⚠️ | **Study pointer semantics carefully — incorrect pointer use in interfaces causes compile errors.** |
| Below 12 ❌ | **Reread Chapter 6 — pointers underpin everything in Chapter 7.** |