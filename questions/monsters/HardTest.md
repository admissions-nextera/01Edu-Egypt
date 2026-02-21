# 🔥 Go Chapters 2-6 - HARD FINAL TEST
## Are You Ready for Chapter 7? Let's Find Out! 💪

**Time Limit:** 90 minutes  
**Passing Score:** 16/20 (80%)  
**Difficulty:** Hard (combines multiple concepts)

**Rules:**
- No looking at previous solutions
- Try to solve without running code first
- Explain your reasoning
- If you get stuck, that's OK - shows what to review!

---

## Problem 1: Shadowing + Closures + Defer (HARD)
```go
package main
import "fmt"

func mystery() {
    x := 1
    
    defer func() {
        fmt.Print(x, " ")
    }()
    
    for i := 0; i < 3; i++ {
        x := x + 1
        defer func() {
            fmt.Print(x, " ")
        }()
    }
    
    x = 10
}

func main() {
    mystery()
}
```
**Question:** What gets printed? (Order and values)

**Difficulty:** ⭐⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** `4 4 4 10`

**Explanation:**
1. Outer `x` starts at 1
2. First defer captures outer `x` (by reference in closure)
3. Loop runs 3 times:
   - Each iteration: `x := x + 1` creates **new inner x** (shadowing)
   - Inner `x` = outer `x` + 1 = 2, then 3, then 4
   - Each defer captures the **same inner loop variable** (same memory location)
   - Last value of inner `x` is 4
4. After loop, outer `x = 10`
5. Defers execute LIFO:
   - Loop defer 3: prints `4` (last value of inner x)
   - Loop defer 2: prints `4` (same variable!)
   - Loop defer 1: prints `4` (same variable!)
   - First defer: prints `10` (outer x)

**Concepts Tested:**
- Shadowing with `:=`
- Closure capture by reference
- Defer LIFO order
- Loop variable capture trap

</details>

---

## Problem 2: Slice Capacity + Append + Sharing (HARD)
```go
package main
import "fmt"

func main() {
    s1 := make([]int, 3, 6)
    s1[0], s1[1], s1[2] = 1, 2, 3
    
    s2 := s1[1:3]
    s3 := s1[1:3:4]
    
    s2 = append(s2, 100)
    s3 = append(s3, 200)
    
    s1[2] = 999
    
    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(s3)
}
```
**Question:** What does each slice contain?

**Difficulty:** ⭐⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
[1 2 999]
[2 999 100]
[2 3 200]
```

**Explanation:**
1. `s1 = [1, 2, 3]` with len=3, cap=6
2. `s2 = s1[1:3]` → `[2, 3]` with len=2, cap=5 (shares underlying array)
3. `s3 = s1[1:3:4]` → `[2, 3]` with len=2, cap=3 (limited capacity)
4. `s2 = append(s2, 100)` → fits in capacity, modifies `s1[3]` → `s2 = [2, 3, 100]`
5. `s3 = append(s3, 200)` → **exceeds capacity (3)**, creates new array → `s3 = [2, 3, 200]`
6. `s1[2] = 999` → modifies s1 and affects s2 (still sharing) but NOT s3 (new array)
7. Final state:
   - `s1 = [1, 2, 999]` (index 3 is beyond len, not visible)
   - `s2 = [2, 999, 100]` (shares s1's array, sees the 999)
   - `s3 = [2, 3, 200]` (independent array)

**Concepts Tested:**
- Slice capacity and sharing
- Full slice expression `[low:high:max]`
- Append reallocation
- Underlying array modifications

</details>

---

## Problem 3: Variadic + Slice Unpacking + Pointers (HARD)
```go
package main
import "fmt"

func modify(nums ...*int) {
    for i := range nums {
        *nums[i] = *nums[i] * 2
    }
}

func main() {
    a, b, c := 1, 2, 3
    slice := []*int{&a, &b, &c}
    modify(slice...)
    fmt.Println(a, b, c)
}
```
**Question:** What are the final values of a, b, c?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** `2 4 6`

**Explanation:**
1. `slice` contains **pointers** to a, b, c
2. `modify(slice...)` unpacks slice of pointers into variadic parameter
3. Inside `modify`, `nums` is `[]*int` containing pointers
4. Loop dereferences and modifies:
   - `*nums[0]` = `a` = `1 * 2` = 2
   - `*nums[1]` = `b` = `2 * 2` = 4
   - `*nums[2]` = `c` = `3 * 2` = 6
5. Original variables are modified through pointers

**Concepts Tested:**
- Variadic functions with pointers
- Slice unpacking with `...`
- Pointer dereferencing
- Modifying through pointers

</details>

---

## Problem 4: Named Returns + Defer + Panic Recovery (HARD)
```go
package main
import "fmt"

func calculator(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
            result = -1
        }
    }()
    
    result = a / b
    return result, nil
}

func main() {
    r1, e1 := calculator(10, 2)
    r2, e2 := calculator(10, 0)
    fmt.Println(r1, e1)
    fmt.Println(r2, e2)
}
```
**Question:** What gets printed?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
5 <nil>
-1 panic: runtime error: integer divide by zero
```

**Explanation:**
1. First call `calculator(10, 2)`:
   - `10 / 2 = 5`, no panic
   - Returns `5, nil`
2. Second call `calculator(10, 0)`:
   - `10 / 0` causes **panic**
   - `recover()` catches panic
   - Defer modifies named returns: `result = -1`, `err = error message`
   - Returns `-1, error`

**Concepts Tested:**
- Named return values
- Defer modifying returns
- Panic and recover
- Error handling pattern

</details>

---

## Problem 5: Map + Pointer Values + Modification (HARD)
```go
package main
import "fmt"

type Counter struct {
    count int
}

func main() {
    m1 := map[string]Counter{
        "a": {count: 1},
    }
    
    m2 := map[string]*Counter{
        "a": {count: 1},
    }
    
    // m1["a"].count = 10  // Line A
    
    temp := m1["a"]
    temp.count = 10
    m1["a"] = temp
    
    m2["a"].count = 10
    
    fmt.Println(m1["a"].count)
    fmt.Println(m2["a"].count)
}
```
**Question:** Why does Line A not compile? What prints?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** `10` and `10`

**Why Line A doesn't compile:**
- Map values are **not addressable**
- Cannot modify struct fields directly in map
- Error: "cannot assign to struct field m1["a"].count in map"

**Explanation:**
1. `m1` - map of **values**: must extract, modify, reassign
2. `m2` - map of **pointers**: can modify directly
3. For `m1`: must do the extract-modify-reassign dance
4. For `m2`: `m2["a"]` is a pointer, can modify through it

**Concepts Tested:**
- Map value addressability
- Map of values vs pointers
- Struct modification patterns

</details>

---

## Problem 6: For-Range + Closure + Goroutines Preview (HARD)
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    values := []int{1, 2, 3, 4, 5}
    
    for _, v := range values {
        go func() {
            fmt.Print(v, " ")
        }()
    }
    
    time.Sleep(100 * time.Millisecond)
}
```
**Question:** What typically gets printed? Why?

**Difficulty:** ⭐⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** Most likely `5 5 5 5 5` (order may vary)

**Explanation:**
1. Loop creates 5 goroutines (concurrent functions)
2. Each closure captures **same variable** `v` by reference
3. By the time goroutines execute, loop has finished
4. `v` has its last value: `5`
5. All goroutines print `5`

**Fix:**
```go
for _, v := range values {
    v := v  // Shadow to capture each value
    go func() {
        fmt.Print(v, " ")
    }()
}
```

**Or pass as parameter:**
```go
for _, v := range values {
    go func(n int) {
        fmt.Print(n, " ")
    }(v)
}
```

**Concepts Tested:**
- Closure loop variable trap
- Concurrent execution (preview of Chapter 9)
- Variable capture by reference

</details>

---

## Problem 7: Switch + Fallthrough + Init Statement (HARD)
```go
package main
import "fmt"

func main() {
    switch x := 5; {
    case x < 3:
        fmt.Print("A")
    case x < 7:
        fmt.Print("B")
        x = 2
        fallthrough
    case x < 5:
        fmt.Print("C")
    default:
        fmt.Print("D")
    }
}
```
**Question:** What gets printed?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** `BC`

**Explanation:**
1. `x := 5` initializes x
2. Blank switch evaluates boolean conditions
3. `x < 3` → false
4. `x < 7` → true, prints "B"
5. `x = 2` (modifies x, but doesn't affect next case check!)
6. `fallthrough` executes next case **without checking condition**
7. Prints "C" (even though `x < 5` is now true, condition isn't checked)

**Important:** `fallthrough` does NOT re-evaluate the next condition!

**Concepts Tested:**
- Blank switch
- Switch init statement
- Fallthrough behavior
- Variable scope in switch

</details>

---

## Problem 8: Slice Copy + Underlying Array (HARD)
```go
package main
import "fmt"

func main() {
    src := []int{1, 2, 3, 4, 5}
    dst := make([]int, 3)
    
    n := copy(dst, src)
    
    src[0] = 100
    dst[0] = 200
    
    fmt.Println(n)
    fmt.Println(src)
    fmt.Println(dst)
}
```
**Question:** What does copy return? Are slices independent?

**Difficulty:** ⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
3
[100 2 3 4 5]
[200 2 3]
```

**Explanation:**
1. `copy(dst, src)` copies **minimum** of len(dst) and len(src) = 3 elements
2. Returns number of elements copied: `3`
3. `dst` is now `[1, 2, 3]` (only 3 elements fit)
4. `copy` creates **independent** copy (new underlying array)
5. Modifying `src[0]` doesn't affect `dst`
6. Modifying `dst[0]` doesn't affect `src`

**Concepts Tested:**
- copy() behavior
- Slice independence after copy
- Understanding len vs cap

</details>

---

## Problem 9: Type Conversion + Overflow + Const (HARD)
```go
package main
import "fmt"

func main() {
    const big = 1000000
    const small int8 = 127
    
    var a int8 = 127
    a = a + 1
    
    var b int8 = small + 1
    
    fmt.Println(a)
    fmt.Println(b)
}
```
**Question:** What gets printed? Does line with b compile?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
- Line `var b int8 = small + 1` **DOES NOT COMPILE**
- Error: "constant 128 overflows int8"

If we remove that line:
```
-128
```

**Explanation:**
1. `a = 127 + 1` → runtime overflow → wraps to `-128`
2. `small + 1` → compile-time constant calculation → `128`
3. `128` doesn't fit in `int8` → **compile-time error**
4. **Typed constants** are checked at compile time!
5. Variable overflow is **silent** (no error, just wraps)

**Concepts Tested:**
- Typed constants vs variables
- Compile-time vs runtime overflow
- Integer overflow behavior

</details>

---

## Problem 10: Pointer + Slice + Function Modification (HARD)
```go
package main
import "fmt"

func modifySlice(s []int) {
    s[0] = 100
    s = append(s, 200)
}

func modifyPointer(s *[]int) {
    (*s)[0] = 100
    *s = append(*s, 200)
}

func main() {
    a := []int{1, 2, 3}
    modifySlice(a)
    fmt.Println(a)
    
    b := []int{1, 2, 3}
    modifyPointer(&b)
    fmt.Println(b)
}
```
**Question:** What does each slice contain after?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
[100 2 3]
[100 2 3 200]
```

**Explanation:**
1. `modifySlice(a)`:
   - `s[0] = 100` modifies shared underlying array → affects `a`
   - `s = append(s, 200)` modifies **local copy** of slice header
   - Original `a` doesn't get the append
2. `modifyPointer(&b)`:
   - `(*s)[0] = 100` modifies underlying array → affects `b`
   - `*s = append(*s, 200)` modifies **original slice header** → `b` gets append

**Lesson:** Use pointer to slice ONLY if you need to modify the slice itself (len/cap)

**Concepts Tested:**
- Slice as reference type
- When to use pointer to slice
- Append with/without pointer

</details>

---

## Problem 11: Multi-Level Shadowing (HARD)
```go
package main
import "fmt"

func main() {
    x := 1
    fmt.Print(x, " ")
    
    {
        x := 2
        fmt.Print(x, " ")
        
        {
            x := 3
            fmt.Print(x, " ")
        }
        
        fmt.Print(x, " ")
    }
    
    fmt.Print(x, " ")
}
```
**Question:** What sequence gets printed?

**Difficulty:** ⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** `1 2 3 2 1`

**Explanation:**
1. Outer scope: `x = 1`, print `1`
2. First block: `x := 2` shadows outer, print `2`
3. Inner block: `x := 3` shadows both, print `3`
4. Exit inner block, back to first block: `x = 2`, print `2`
5. Exit first block, back to outer: `x = 1`, print `1`

**Visualization:**
```
Level 0: x = 1 ─────────────────┐
         print 1                │
                                │
Level 1: x = 2 ───────────┐    │
         print 2          │    │
                          │    │
Level 2: x = 3            │    │
         print 3          │    │
         (exit level 2)   │    │
                          │    │
         print 2 (L1) ────┘    │
         (exit level 1)        │
                               │
         print 1 (L0) ─────────┘
```

**Concepts Tested:**
- Nested block scopes
- Multiple level shadowing
- Variable lifetime

</details>

---

## Problem 12: Struct + Pointer + Nil (HARD)
```go
package main
import "fmt"

type Node struct {
    value int
    next  *Node
}

func (n *Node) Add(val int) {
    current := n
    for current.next != nil {
        current = current.next
    }
    current.next = &Node{value: val}
}

func main() {
    var head *Node
    head.Add(1)  // What happens?
}
```
**Question:** Does this work? What happens?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** **PANIC!** "invalid memory address or nil pointer dereference"

**Explanation:**
1. `var head *Node` → `head` is `nil`
2. `head.Add(1)` → can call method on nil receiver (Go allows this!)
3. Inside `Add`, `n` is `nil`
4. `current := n` → `current` is `nil`
5. `current.next` tries to access field of `nil` pointer → **PANIC**

**Fix - handle nil receiver:**
```go
func (n *Node) Add(val int) {
    if n == nil {
        // Can't modify nil receiver!
        return
    }
    // ... rest of code
}
```

**Or initialize properly:**
```go
head := &Node{}  // Not nil
head.Add(1)      // Works!
```

**Concepts Tested:**
- Nil pointer receivers (legal but dangerous)
- Pointer method receivers
- Linked list patterns
- Nil checking

</details>

---

## Problem 13: Defer + Named Returns + Multiple Returns (HARD)
```go
package main
import "fmt"

func divide(a, b int) (result int, err string) {
    defer func() {
        if result < 0 {
            err = "negative result"
        }
    }()
    
    if b == 0 {
        err = "division by zero"
        return
    }
    
    result = a / b
    return
}

func main() {
    r1, e1 := divide(10, 2)
    r2, e2 := divide(10, 0)
    r3, e3 := divide(-10, 2)
    
    fmt.Println(r1, e1)
    fmt.Println(r2, e2)
    fmt.Println(r3, e3)
}
```
**Question:** What gets printed for each call?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
5 
0 division by zero
-5 negative result
```

**Explanation:**
1. `divide(10, 2)`:
   - `result = 5`, `err = ""`
   - Defer runs: `5 < 0` is false, no change
   - Returns `5, ""`
2. `divide(10, 0)`:
   - `err = "division by zero"`
   - Bare `return` (returns `result=0`, `err="division by zero"`)
   - Defer runs: `0 < 0` is false, no change
   - Returns `0, "division by zero"`
3. `divide(-10, 2)`:
   - `result = -5`, `err = ""`
   - Defer runs: `-5 < 0` is true, sets `err = "negative result"`
   - Returns `-5, "negative result"`

**Concepts Tested:**
- Named returns with defer
- Bare return behavior
- Defer modifying returns
- Multiple return values

</details>

---

## Problem 14: Variadic + Empty + Nil (HARD)
```go
package main
import "fmt"

func sum(nums ...int) int {
    if nums == nil {
        return -1
    }
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum())
    fmt.Println(sum(1, 2, 3))
    
    var s []int
    fmt.Println(sum(s...))
}
```
**Question:** What gets printed for each call?

**Difficulty:** ⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
0
6
0
```

**Explanation:**
1. `sum()` → `nums` is **empty slice**, NOT nil! Returns `0`
2. `sum(1, 2, 3)` → returns `1 + 2 + 3 = 6`
3. `sum(s...)` → `s` is nil, unpacked into variadic
   - Inside function, `nums` is nil slice
   - But the check `nums == nil` is actually **never true** with variadic!
   - Wait... actually this is a tricky edge case!
   - Unpacking nil slice `s...` creates nil variadic parameter
   - `nums == nil` is true? Let me reconsider...

**Actually, correct answer:**
- All print `0` because variadic creates empty slice even from nil
- The nil check is misleading - variadic parameters are slices, can be nil

**Let me correct:**
```
0  (empty variadic)
6  (sum of values)
0  (nil unpacked is still nil, but range over nil works)
```

**Concepts Tested:**
- Variadic functions
- Nil vs empty slice
- Range over nil slice (works fine!)

</details>

---

## Problem 15: Closure + Pointer + Slice (HARD)
```go
package main
import "fmt"

func main() {
    values := []int{1, 2, 3}
    funcs := []func(){}
    
    for i := range values {
        funcs = append(funcs, func() {
            fmt.Print(values[i], " ")
        })
    }
    
    values[0] = 100
    
    for _, f := range funcs {
        f()
    }
}
```
**Question:** What gets printed?

**Difficulty:** ⭐⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** `3 3 3` (or panic if accessing out of bounds!)

**Wait, let me recalculate:**
- Loop variable `i` is shared by all closures
- After loop ends, `i = 2` (last index)
- All closures print `values[2]`
- But we modified `values[0]`, not `values[2]`
- So prints `3 3 3`

**Actually:** Could also print `100 100 100` if... no wait.

**Correct answer:** `3 3 3`

**Explanation:**
1. All closures capture same variable `i`
2. After loop, `i = 2` (last value)
3. All closures execute `values[2]` → prints `3` three times
4. The modification `values[0] = 100` doesn't matter

**Two traps combined:**
- Loop variable capture (all closures see same `i`)
- Slice modification (red herring!)

**Concepts Tested:**
- Closure loop variable trap
- Slice reference semantics
- Multiple trap combination

</details>

---

## Problem 16: Rune Iteration + String Indexing (HARD)
```go
package main
import "fmt"

func main() {
    s := "Hello, 世界"
    
    fmt.Println(len(s))
    
    count := 0
    for range s {
        count++
    }
    fmt.Println(count)
    
    fmt.Printf("%T\n", s[0])
}
```
**Question:** What gets printed for each?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
13
9
uint8
```

**Explanation:**
1. `len(s)` returns **byte count**, not character count
   - "Hello, " = 7 bytes
   - "世" = 3 bytes (UTF-8)
   - "界" = 3 bytes (UTF-8)
   - Total = 13 bytes
2. `for range s` iterates over **runes** (characters)
   - "H" "e" "l" "l" "o" "," " " "世" "界" = 9 characters
3. `s[0]` returns **byte** (uint8), not rune

**Concepts Tested:**
- String length vs character count
- UTF-8 encoding
- Rune vs byte
- Range over string

</details>

---

## Problem 17: Switch Type Assertion Preview (HARD)
```go
package main
import "fmt"

func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v*2)
    case string:
        fmt.Printf("string: %s!\n", v)
    case []int:
        fmt.Printf("slice: len=%d\n", len(v))
    default:
        fmt.Printf("unknown: %T\n", v)
    }
}

func main() {
    describe(42)
    describe("hello")
    describe([]int{1, 2, 3})
    describe(true)
}
```
**Question:** What gets printed? (Preview of interfaces)

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
int: 84
string: hello!
slice: len=3
unknown: bool
```

**Explanation:**
1. `interface{}` can hold any type (empty interface)
2. Type switch checks actual type
3. `v` has correct type in each case
4. First: `42` is int → `42 * 2 = 84`
5. Second: `"hello"` is string → prints with "!"
6. Third: `[]int{1,2,3}` → len is 3
7. Fourth: `true` is bool → no case matches, default

**Concepts Tested:**
- Empty interface (preview)
- Type switch (preview of Chapter 7)
- Type assertions

</details>

---

## Problem 18: Slice Re-slicing Edge Cases (HARD)
```go
package main
import "fmt"

func main() {
    s := []int{0, 1, 2, 3, 4, 5}
    
    a := s[2:5]
    b := a[:cap(a)]
    c := s[:0]
    
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(len(c), cap(c))
}
```
**Question:** What does each slice see?

**Difficulty:** ⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:**
```
[2 3 4]
[2 3 4 5]
[]
0 6
```

**Explanation:**
1. `s[2:5]` → `a = [2, 3, 4]`, len=3, cap=4 (to end of s)
2. `a[:cap(a)]` → extends `a` to full capacity → `b = [2, 3, 4, 5]`
3. `s[:0]` → `c` has len=0 but still shares underlying array
4. `c` has cap=6 (full capacity of s)
5. `c` is empty but can be appended to reuse the array

**Key insight:** `s[:0]` creates zero-length slice that shares array

**Concepts Tested:**
- Slice capacity
- Re-slicing to capacity
- Zero-length slices
- Underlying array sharing

</details>

---

## Problem 19: Const + Untyped + Type Inference (HARD)
```go
package main
import "fmt"

func main() {
    const a = 10
    const b int = 10
    
    var x float64 = a
    var y float64 = b  // Does this compile?
    
    fmt.Println(x, y)
}
```
**Question:** Does line with y compile?

**Difficulty:** ⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** **NO!** Line with `y` doesn't compile.

**Error:** "cannot use b (type int) as type float64 in assignment"

**Explanation:**
1. `const a = 10` → **untyped** constant
   - Can be assigned to any numeric type
   - `var x float64 = a` works!
2. `const b int = 10` → **typed** constant (type is int)
   - Cannot assign int to float64 without conversion
   - `var y float64 = b` fails!
   - Must use: `var y float64 = float64(b)`

**Concepts Tested:**
- Typed vs untyped constants
- Type conversion requirements
- Const flexibility

</details>

---

## Problem 20: Ultimate Combined Challenge (HARD)
```go
package main
import "fmt"

func process(nums []int) (result []int) {
    result = make([]int, 0, len(nums))
    
    defer func() {
        for i := range result {
            result[i] = result[i] * 2
        }
    }()
    
    for _, n := range nums {
        if n%2 == 0 {
            result = append(result, n)
        }
    }
    
    return
}

func main() {
    input := []int{1, 2, 3, 4, 5, 6}
    output := process(input)
    fmt.Println(output)
}
```
**Question:** What gets returned?

**Difficulty:** ⭐⭐⭐⭐⭐

<details>
<summary>Click for Answer</summary>

**Answer:** `[4 8 12]`

**Explanation:**
Step by step:
1. `result` is named return, starts as empty slice
2. Loop filters even numbers: `[2, 4, 6]`
3. About to return `[2, 4, 6]`
4. **Defer executes** before return:
   - Modifies named return `result`
   - Doubles each value: `[4, 8, 12]`
5. Returns modified `result`

**Concepts Combined:**
- Named returns
- Defer modifying returns
- Slice filtering
- For-range iteration
- Append operations

</details>

---

## 🏆 SCORING

**Count your correct answers:**

- **18-20 correct:** 🔥 **MASTER LEVEL** - Ready for Chapter 7!
- **16-17 correct:** 💪 **STRONG** - Review weak areas, then proceed
- **14-15 correct:** ⚠️ **GOOD** - Review Chapters 4-6, practice more
- **12-13 correct:** 📚 **NEEDS WORK** - Go back and review all chapters
- **Below 12:** 🔄 **RESTART** - Fundamentals need reinforcement

---

## 📊 CONCEPT BREAKDOWN

Check which concepts you missed:

**Shadowing & Scope:**
- Problems 1, 7, 11

**Slices (Capacity, Sharing, Operations):**
- Problems 2, 8, 10, 18, 23

**Closures:**
- Problems 1, 6, 15

**Defer:**
- Problems 1, 4, 13, 16, 20

**Pointers:**
- Problems 3, 5, 10, 12

**Maps:**
- Problem 5

**Type System:**
- Problems 9, 16, 17, 19

**Functions:**
- Problems 3, 14

**Control Structures:**
- Problem 7

**Combined Concepts:**
- Problems 4, 13, 15, 20

---

## 💡 WHAT TO DO NEXT

### If you scored 16+:
✅ **You're ready for Chapter 7!**
- You have solid fundamentals
- Minor gaps are normal
- Review concepts you missed
- **Proceed with confidence!**

### If you scored 14-15:
⚠️ **Almost there!**
- Review the chapters with most mistakes
- Re-do practice problems for weak areas
- Focus on combined concepts
- Re-take this test in 2-3 days

### If you scored 12-13:
📚 **Need more practice**
- Go back to Chapters 4-6
- Work through all practice problems again
- Focus on understanding WHY, not just WHAT
- Build small programs to practice
- Re-test in 1 week

### If you scored below 12:
🔄 **Start over with stronger foundation**
- Review Chapters 2-3 (basics)
- Work through easy problems first
- Build confidence before medium problems
- Don't rush - understanding > speed
- Consider working with someone for pair review

---

## 🎯 KEY LEARNING INSIGHTS

**Most Challenging Concepts (in order):**
1. **Shadowing + Closures** (Problem 1, 15) - Combines 3 concepts
2. **Slice capacity + sharing** (Problem 2) - Memory model understanding
3. **Defer + Named Returns** (Problems 4, 13, 20) - Execution order
4. **Type system edge cases** (Problems 9, 19) - Const vs var differences
5. **Closure loop traps** (Problems 6, 15) - Famous gotcha

**If you missed these, you MUST understand them before Chapter 7!**

---

## 📝 REFLECTION QUESTIONS

After taking this test, answer honestly:

1. **Which concept surprised you most?**
2. **Which problem took longest to solve?**
3. **Did you have to run code, or could you predict output?**
4. **Which previous chapter needs most review?**
5. **Are you confident explaining answers to a friend?**

**If you can explain all 20 answers clearly, you're ready! 🚀**

**Good luck! You've got this! 💪🔥**