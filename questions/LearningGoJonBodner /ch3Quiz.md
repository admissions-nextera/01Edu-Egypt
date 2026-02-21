# 📘 Learning Go — Chapter 3 Quiz
## Composite Types: Arrays, Slices, Maps, and Structs

**Time Limit:** 45 minutes  
**Total Questions:** 28  
**Passing Score:** 22/28 (78%)

> This quiz covers: arrays (fixed-size, value semantics), slices (len, cap, append, make, copy), maps (declaration, nil maps, comma-ok), and structs (fields, anonymous structs, embedding).

---

## 📋 SECTION 1: ARRAYS (4 Questions)

### Q1: What is the key difference between an array and a slice in Go?

**A)** Arrays are faster; slices are safer  
**B)** An array has a fixed, compile-time size that is part of its type; a slice is dynamically sized and is a view into an underlying array  
**C)** Arrays are reference types; slices are value types  
**D)** Arrays can only hold primitive types  

<details><summary>💡 Answer</summary>

**B) Arrays are fixed-size value types; slices are dynamic views**

```go
var a [5]int       // [5]int — size 5 is PART OF THE TYPE
var b [3]int       // [3]int — different type from [5]int!
// a = b           // compile error: mismatched types

s := []int{1, 2, 3}  // slice — size not part of type
```

Because size is part of array's type, you almost never use arrays directly in Go — slices are almost always the right choice. Arrays are useful for: SHA-256 hashes `[32]byte`, hardware registers, or when you explicitly want value semantics.

</details>

---

### Q2: What happens when you assign one array to another in Go?

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 99
fmt.Println(a[0])
```

**A)** `99` — `b` is a reference to `a`  
**B)** `1` — arrays are value types; `b` is a complete copy of `a`  
**C)** Compile error  
**D)** Panic  

<details><summary>💡 Answer</summary>

**B) `1` — arrays have value semantics (copying)**

Assigning an array copies ALL its elements. `b` is an entirely independent copy. Modifying `b` never affects `a`. This is unlike slices, which share the underlying array. Array value semantics can be expensive for large arrays — prefer slices or pass arrays by pointer when size matters.

</details>

---

### Q3: How do you declare an array of 5 integers where all elements are zero?

**A)** `var a = array[5]int`  
**B)** `var a [5]int`  
**C)** `a := make([5]int)`  
**D)** `a := [5]int{}`  

<details><summary>💡 Answer</summary>

**B and D are both valid — B uses `var`, D uses `:=` with composite literal**

```go
var a [5]int      // all zeros — zero value for int
b := [5]int{}     // also all zeros — explicit composite literal
c := [3]int{1, 2, 3}  // initialized values
d := [...]int{4, 5, 6}  // size inferred from values
```

`[...]` lets the compiler count the elements. `make` does NOT work with arrays — only slices, maps, and channels.

</details>

---

### Q4: Are `[3]int` and `[4]int` the same type?

**A)** Yes — size is irrelevant for type equality  
**B)** No — the size is part of the array type; they are completely different types and cannot be assigned to each other  
**C)** They are the same underlying type  
**D)** Yes if the elements are the same type  

<details><summary>💡 Answer</summary>

**B) No — `[3]int` and `[4]int` are different types**

This is the fundamental property of arrays that makes them rarely useful in Go. You can't write a function that accepts "any size array" — the size must be fixed in the function signature. Slices solve this: `[]int` accepts any length.

</details>

---

## 📋 SECTION 2: SLICES (12 Questions)

### Q5: What are the three components of a slice header?

**A)** Start, end, type  
**B)** Pointer to underlying array, length, capacity  
**C)** Address, size, element type  
**D)** Data, length, growth factor  

<details><summary>💡 Answer</summary>

**B) Pointer to underlying array, length, capacity**

```
slice header: [ptr | len | cap]
                ↓
underlying array: [0][1][2][3][4][5][6][7]
                   ↑___↑___↑       ↑_______↑
                   slice view      capacity
```

- `len`: number of elements accessible in the slice
- `cap`: total elements from the slice's start to the end of the underlying array
- `ptr`: where in the array the slice starts

`len(s)` and `cap(s)` access these values.

</details>

---

### Q6: What is the zero value of a slice, and is it usable?

**A)** `[]` — an empty but usable slice  
**B)** `nil` — a nil slice; `len` and `cap` return 0, and `append` works on it, but direct indexing panics  
**C)** `nil` — a nil slice that panics on any operation  
**D)** An empty slice — identical in behavior to `[]int{}`  

<details><summary>💡 Answer</summary>

**B) `nil` — usable with `len`, `cap`, `append`, and `range`; panics on index**

```go
var s []int        // nil slice
fmt.Println(s == nil)   // true
fmt.Println(len(s))     // 0
fmt.Println(cap(s))     // 0
s = append(s, 1)        // works — append handles nil slices
s[0]                    // panic if len is still 0
```

A nil slice and an empty slice (`[]int{}`) behave identically for `len`, `cap`, `append`, and `range`. They differ only in `s == nil`.

</details>

---

### Q7: What does `append` return when it runs out of capacity?

**A)** An error  
**B)** A new slice backed by a new, larger array — the original slice's backing array is unchanged  
**C)** The same slice, extended in place  
**D)** A panic  

<details><summary>💡 Answer</summary>

**B) A new slice backed by a new, larger array**

```go
s := make([]int, 3, 3)  // len=3, cap=3
t := append(s, 4)       // cap exceeded — new array allocated
// s still points to old array; t points to new array
```

**Always assign `append` back to the same variable:**
```go
s = append(s, 4)  // correct
append(s, 4)      // wrong — result discarded, silent bug
```

Go doubles the capacity when growing (roughly), making amortized appends O(1).

</details>

---

### Q8: What is the output?
```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]
b[0] = 99
fmt.Println(a)
```

**A)** `[1 2 3 4 5]` — `b` is a copy  
**B)** `[1 99 3 4 5]` — `b` shares the underlying array with `a`  
**C)** Compile error  
**D)** `[99 99 3 4 5]`  

<details><summary>💡 Answer</summary>

**B) `[1 99 3 4 5]` — slices share the underlying array**

`a[1:3]` creates a slice that views elements 1 and 2 of `a`'s underlying array. `b[0]` is `a[1]` — they are the same memory location. This sharing is a feature (efficient) but a common source of bugs. Use `copy` when you need independence:

```go
b := make([]int, 2)
copy(b, a[1:3])
b[0] = 99        // now safe — doesn't affect a
```

</details>

---

### Q9: What is the difference between `make([]int, 5)` and `make([]int, 0, 5)`?

**A)** No difference  
**B)** `make([]int, 5)` creates a slice with len=5 and cap=5 (all zeros); `make([]int, 0, 5)` creates a slice with len=0 and cap=5 (nothing accessible yet, but room for 5 without reallocation)  
**C)** `make([]int, 0, 5)` is invalid  
**D)** `make([]int, 5)` allocates on the stack; `make([]int, 0, 5)` on the heap  

<details><summary>💡 Answer</summary>

**B) len=5 zeros vs len=0 with pre-allocated capacity**

```go
a := make([]int, 5)     // [0 0 0 0 0] — 5 accessible zeros
b := make([]int, 0, 5)  // [] — empty but capacity for 5 before reallocation

a[0] = 1                // valid — len is 5
b[0] = 1                // panic — len is 0

b = append(b, 1)        // correct way to add to b
```

Use `make([]T, 0, n)` when you know the approximate final size and will use `append` to fill it.

</details>

---

### Q10: What does `copy(dst, src)` return, and what happens if they have different lengths?

**A)** Nothing — it panics if lengths differ  
**B)** The number of elements copied — `min(len(dst), len(src))` elements are copied  
**C)** A new slice  
**D)** `true` if the copy succeeded  

<details><summary>💡 Answer</summary>

**B) Returns the number of elements copied — copies `min(len(dst), len(src))` elements**

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)
n := copy(dst, src)   // copies 3 elements (min of 3 and 5)
fmt.Println(n)        // 3
fmt.Println(dst)      // [1 2 3]
```

`copy` never panics due to length mismatch — it just copies as many as fit. `dst` must already be allocated (it copies INTO existing elements, it doesn't grow `dst`).

</details>

---

### Q11: What is the output?
```go
s := []int{1, 2, 3}
s = append(s[:1], s[2:]...)
fmt.Println(s)
```

**A)** `[1 2 3]`  
**B)** `[1 3]` — removes element at index 1  
**C)** `[2 3]`  
**D)** Panic  

<details><summary>💡 Answer</summary>

**B) `[1 3]` — removes element at index 1**

This is the idiomatic pattern to remove an element from a slice without preserving order:
- `s[:1]` = `[1]`
- `s[2:]` = `[3]`
- `append([1], [3]...)` = `[1 3]`

Note: this modifies the original underlying array. If you had other slices sharing the same array, they'd see unexpected data. The order-preserving remove uses the same trick but can be O(n).

</details>

---

### Q12: When you pass a slice to a function and the function appends to it, does the caller's slice change?

**A)** Yes — slices are reference types  
**B)** It depends: if append doesn't exceed capacity, the underlying array is shared and the caller might see changes; if it does exceed capacity, a new array is allocated and the caller's slice is unaffected  
**C)** No — slices are always copied  
**D)** Yes always — slices are passed by reference  

<details><summary>💡 Answer</summary>

**B) It depends on whether capacity is exceeded**

```go
func addToSlice(s []int) {
    s = append(s, 99)  // modifies s locally
    // caller's slice header is unchanged (len, cap, ptr are copied)
}
```

The slice header (ptr, len, cap) is passed by value. If append allocates a new array, only the local copy of the header is updated. The caller's `len` never changes from the function's `append`. To modify the caller's slice, return the new slice and assign it.

</details>

---

### Q13: What does `a[2:5]` require about index `5` and the slice `a`?

**A)** `a` must have at least 6 elements (index 5 must be valid)  
**B)** `5` must be ≤ `cap(a)` — it can equal `len(a)` (exclusive end is ok to be at `len`); panics if `5 > cap(a)`  
**C)** `5` must be < `len(a)`  
**D)** Any indices are fine — out-of-bounds returns zero values  

<details><summary>💡 Answer</summary>

**B) The high bound can go up to `cap(a)` — but not beyond**

```go
a := make([]int, 5, 10)  // len=5, cap=10
b := a[2:8]   // valid — 8 <= cap(10)
c := a[2:11]  // panic — 11 > cap(10)
d := a[2:5]   // valid — exactly at len
```

Slice expressions use `cap` for the upper bound, not `len`. This allows "re-extending" a slice to reclaim elements that were sliced away (as long as the underlying array still exists).

</details>

---

### Q14: What is the three-index slice expression `a[low:high:max]` and why would you use it?

**A)** A syntax error  
**B)** Creates a slice with `len = high-low` and `cap = max-low` — limits the capacity of the resulting slice to prevent accidental sharing of the backing array  
**C)** Allocates a new array of size `max`  
**D)** Same as `a[low:high]` — the third index is ignored  

<details><summary>💡 Answer</summary>

**B) Limits the capacity of the slice — protects the backing array**

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3:3]  // len=2, cap=2 — can't see past index 3
// b = append(b, 99) would allocate a new array (cap exceeded)
// without :3, b would share array with a through index 4
```

Useful when returning a sub-slice from a function — prevents callers from accidentally writing into the original array via `append`.

</details>

---

### Q15: What does `for i, v := range s` give you for each iteration?

**A)** `i` = element value, `v` = index  
**B)** `i` = index (0-based), `v` = a **copy** of the element at that index  
**C)** `i` = index, `v` = a pointer to the element  
**D)** `i` = index, `v` = the element by reference  

<details><summary>💡 Answer</summary>

**B) `i` = index, `v` = a copy of the element**

```go
s := []int{10, 20, 30}
for i, v := range s {
    v = 99    // modifies the copy — does NOT change s
    s[i] = 99 // this DOES change s
}
```

`v` is a copy. To modify elements in place, use `s[i] = newValue` instead of `v = newValue`. Use `for _, v := range s` when you don't need the index; `for i := range s` when you don't need the value.

</details>

---

## 📋 SECTION 3: MAPS (7 Questions)

### Q16: What is the zero value of a map, and what happens if you try to write to it?

**A)** `{}` — an empty map; writing works fine  
**B)** `nil` — reading returns zero values safely; writing causes a panic  
**C)** `nil` — all operations panic  
**D)** An empty map — identical to `map[string]int{}`  

<details><summary>💡 Answer</summary>

**B) `nil` — reads return zero values; writes panic**

```go
var m map[string]int    // nil map
fmt.Println(m["key"])  // 0 — safe, returns zero value
m["key"] = 1          // PANIC: assignment to entry in nil map
```

Always initialize a map before writing:
```go
m := make(map[string]int)
m := map[string]int{}
```

</details>

---

### Q17: How do you check whether a key exists in a map?

**A)** `if m["key"] != nil`  
**B)** `val, ok := m["key"]` — the comma-ok idiom: `ok` is `true` if the key exists  
**C)** `if m.Contains("key")`  
**D)** `if m["key"] != 0` (for `map[string]int`)  

<details><summary>💡 Answer</summary>

**B) The comma-ok idiom: `val, ok := m["key"]`**

```go
m := map[string]int{"a": 1}
val, ok := m["a"]   // val=1, ok=true
val, ok = m["z"]    // val=0, ok=false (key doesn't exist)
```

Option D is wrong: `m["z"]` returns `0` (zero value) even if `"z"` was never set, AND if you had explicitly stored `0` for `"z"`. The only reliable way to distinguish "key absent" from "key has zero value" is the comma-ok idiom.

</details>

---

### Q18: How do you delete a key from a map?

**A)** `m["key"] = nil`  
**B)** `m.Remove("key")`  
**C)** `delete(m, "key")`  
**D)** `m["key"] = map.zero`  

<details><summary>💡 Answer</summary>

**C) `delete(m, "key")`**

```go
m := map[string]int{"a": 1, "b": 2}
delete(m, "a")
fmt.Println(m)  // map[b:2]
```

`delete` on a non-existent key is a no-op — it does not panic. `delete` on a nil map panics (same as writing to a nil map).

</details>

---

### Q19: Is it safe to iterate over a map with `for k, v := range m` while other goroutines are reading from it?

**A)** Yes — map reads are concurrent-safe  
**B)** No — concurrent map reads AND iteration without synchronization are a race condition; concurrent read+write is even worse and causes a panic  
**C)** Yes — Go maps are internally synchronized  
**D)** Only if using `sync.Map`  

<details><summary>💡 Answer</summary>

**B) No — maps are not concurrent-safe**

Go's built-in map is not safe for concurrent access. Multiple concurrent readers ARE safe (Go 1.6+). But a concurrent read + write (any write), or iteration + write, causes a race condition. In Go 1.6+, the runtime detects concurrent map writes and panics with "concurrent map read and map write." Use `sync.Mutex` or `sync.RWMutex` to protect maps accessed from multiple goroutines.

</details>

---

### Q20: What is the output order of this code?
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
for k, v := range m {
    fmt.Printf("%s:%d ", k, v)
}
```

**A)** Always `a:1 b:2 c:3` — maps are stored in insertion order  
**B)** Random — Go intentionally randomizes map iteration order  
**C)** Alphabetical — Go sorts map keys during iteration  
**D)** Undefined but consistent across runs  

<details><summary>💡 Answer</summary>

**B) Random — intentionally randomized**

Go deliberately randomizes map iteration order (since Go 1.0). If you need sorted output, collect the keys into a slice, sort it, then iterate:
```go
keys := make([]string, 0, len(m))
for k := range m { keys = append(keys, k) }
sort.Strings(keys)
for _, k := range keys { fmt.Println(k, m[k]) }
```

</details>

---

### Q21: What is the idiomatic way to use a map as a set of strings?

**A)** `map[string]bool` — set `m[key] = true` to add; check with `if m[key]`  
**B)** `map[string]int` — use 0 and 1  
**C)** `[]string` with `slices.Contains`  
**D)** `map[string]struct{}` — use `m[key] = struct{}{}` to add; check with comma-ok  

<details><summary>💡 Answer</summary>

**D) `map[string]struct{}` — or `map[string]bool` for readability**

```go
// Memory-efficient: struct{} has zero size
set := map[string]struct{}{}
set["alice"] = struct{}{}
_, exists := set["alice"]   // true

// More readable:
seen := map[string]bool{}
seen["alice"] = true
if seen["alice"] { ... }
```

`map[string]struct{}` uses no memory for values (struct{} is zero-sized). `map[string]bool` is more readable. The book shows both — prefer `struct{}` in performance-sensitive code.

</details>

---

### Q22: Can you use a slice as a map key?

**A)** Yes — any type can be a map key  
**B)** No — map keys must be comparable (support `==`); slices, maps, and functions are not comparable  
**C)** Yes — but only `[]byte`  
**D)** Only if the slice has length ≤ 8  

<details><summary>💡 Answer</summary>

**B) No — slices are not comparable and cannot be map keys**

Map keys must be comparable (`==` operator must work). Slices, maps, and functions are not comparable in Go. Valid key types: all numeric types, `string`, `bool`, pointers, arrays (of comparable element type), structs (all comparable fields). Use `fmt.Sprint(slice)` as a string key if you must, or design around the limitation.

</details>

---

## 📋 SECTION 4: STRUCTS (5 Questions)

### Q23: What is the zero value of a struct?

**A)** `nil`  
**B)** Each field gets its own zero value — no fields are uninitialized  
**C)** Compile error — structs must be initialized explicitly  
**D)** All fields are `nil`  

<details><summary>💡 Answer</summary>

**B) Each field gets its own zero value**

```go
type Person struct {
    Name string
    Age  int
    Active bool
}
var p Person
// p.Name == "", p.Age == 0, p.Active == false
```

Struct zero values are fully initialized and usable. This applies recursively — nested struct fields also get zero values.

</details>

---

### Q24: What is an anonymous struct and when would you use one?

**A)** A struct with no exported fields  
**B)** A struct type defined inline without a named type — useful for grouping related temporary data or JSON test fixtures  
**C)** A struct that implements no interfaces  
**D)** A struct with only one field  

<details><summary>💡 Answer</summary>

**B) Inline struct type definition — no type name**

```go
// Table-driven tests:
tests := []struct {
    input    string
    expected int
}{
    {"hello", 5},
    {"hi", 2},
}

// Quick grouping:
config := struct {
    host string
    port int
}{"localhost", 8080}
```

Anonymous structs are common in table-driven tests and for grouping data that won't be reused elsewhere.

</details>

---

### Q25: What is the difference between these two struct initializations?
```go
type Point struct { X, Y int }

a := Point{1, 2}
b := Point{X: 1, Y: 2}
```

**A)** `a` is invalid — you must always use field names  
**B)** Both are valid; `a` uses positional initialization (field order matters, fragile); `b` uses named fields (order doesn't matter, more robust)  
**C)** `b` is invalid — you can't name fields in a literal  
**D)** They produce different values  

<details><summary>💡 Answer</summary>

**B) Both valid — positional vs named; prefer named for robustness**

Positional initialization breaks silently if someone adds or reorders struct fields. Named initialization is explicit and resistant to refactoring. The Go style guide recommends named fields for all struct literals outside of tests with very simple structs.

</details>

---

### Q26: What is struct embedding and what does it give you?

**A)** Storing one struct inside another as a named field  
**B)** Including one type inside another without a field name — the outer struct "inherits" the embedded type's methods and fields, accessible directly  
**C)** A syntax error in Go  
**D)** Making one struct a subclass of another  

<details><summary>💡 Answer</summary>

**B) Anonymous field inclusion — promotes methods and fields to the outer type**

```go
type Animal struct {
    Name string
}
func (a Animal) Speak() string { return a.Name + " speaks" }

type Dog struct {
    Animal         // embedded — no field name
    Breed string
}

d := Dog{Animal: Animal{"Rex"}, Breed: "Lab"}
fmt.Println(d.Name)    // promoted — same as d.Animal.Name
fmt.Println(d.Speak()) // promoted — same as d.Animal.Speak()
```

This is composition, not inheritance. The embedded type's methods are "promoted" but the outer type does NOT become the embedded type.

</details>

---

### Q27: Are two struct values with the same fields and values equal with `==`?

**A)** Yes — if all fields are comparable and equal, the structs are equal  
**B)** No — structs are always compared by reference  
**C)** Only if they are the same named type  
**D)** Only if they were created in the same function  

<details><summary>💡 Answer</summary>

**A) Yes — if all fields are comparable and equal**

```go
type Point struct { X, Y int }
a := Point{1, 2}
b := Point{1, 2}
fmt.Println(a == b)  // true

type Bad struct { data []int }  // contains slice — not comparable!
// Bad{} == Bad{} would be a compile error
```

A struct is comparable if and only if ALL its fields are comparable. A struct containing a slice, map, or function is NOT comparable and `==` is a compile error.

</details>

---

### Q28: What does this print?
```go
type Counter struct{ count int }

func increment(c Counter) {
    c.count++
}

func main() {
    var c Counter
    increment(c)
    fmt.Println(c.count)
}
```

**A)** `1`  
**B)** `0` — structs are passed by value; `increment` receives a copy  
**C)** Compile error  
**D)** `1` if the struct is small enough for the optimizer  

<details><summary>💡 Answer</summary>

**B) `0` — structs are value types, passed by copy**

`increment` receives a copy of `c`. Modifying the copy doesn't affect the original. To modify the caller's struct, use a pointer:
```go
func increment(c *Counter) {
    c.count++
}
increment(&c)
fmt.Println(c.count)  // 1
```

This is one of the most important distinctions in Go — value vs pointer semantics.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 26–28 ✅ | **Excellent.** Composite types mastered — proceed to Chapter 4. |
| 22–25 ✅ | **Ready.** Review slice sharing and nil map behavior before moving on. |
| 17–21 ⚠️ | **Review first.** Slice capacity, map nil behavior, and value vs pointer semantics need more work. |
| Below 17 ❌ | **Re-read Chapter 3.** Slices and maps are foundational — every subsequent chapter builds on them. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q4 | Arrays: value type, size in type, no `make` for arrays |
| Q5–Q15 | Slices: header structure, nil vs empty, `append` allocation, sharing, `copy`, three-index, `range` copy |
| Q16–Q22 | Maps: nil map panic, comma-ok, `delete`, iteration order, set pattern, comparable keys |
| Q23–Q28 | Structs: zero values, anonymous structs, positional vs named init, embedding, value semantics |