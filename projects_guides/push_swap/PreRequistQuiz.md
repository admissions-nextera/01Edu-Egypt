# 🎯 Push-Swap Project Prerequisites Quiz
## Stacks · 11 Instructions · Sorting Complexity · Algorithm Design · Go Fundamentals · Input Validation

**Time Limit:** 50 minutes  
**Total Questions:** 30  
**Passing Score:** 24/30 (80%)

> Questions are tagged: 🟢 Easy · 🟡 Medium · 🔴 Hard  
> All topics are general — no specific project knowledge required.

---

## 📋 SECTION 1: STACKS (6 Questions)

### Q1 🟢 — What does LIFO mean and which data structure uses it?

**A)** Last In, First Out — Queue  
**B)** Last In, First Out — Stack  
**C)** Last In, First Out — Linked List  
**D)** Least Important, First Out — Priority Queue  

<details><summary>💡 Answer</summary>

**B) Last In, First Out — Stack**

```
Push 1 → [1]
Push 2 → [2, 1]
Push 3 → [3, 2, 1]
Pop    → returns 3 → [2, 1]   ← last pushed = first popped
Pop    → returns 2 → [1]
```

A stack works like a pile of plates — you can only add or remove from the top. The last element you pushed is always the first one you get back when you pop.

</details>

---

### Q2 🟢 — What are the three core operations of a stack?

**A)** Insert, Delete, Search  
**B)** Enqueue, Dequeue, Peek  
**C)** Push, Pop, Peek  
**D)** Add, Remove, Find  

<details><summary>💡 Answer</summary>

**C) Push, Pop, Peek**

```go
type Stack struct {
    data []int  // index 0 = top of the stack
}

func (s *Stack) push(val int) { s.data = append([]int{val}, s.data...) }
func (s *Stack) pop() int     { top := s.data[0]; s.data = s.data[1:]; return top }
func (s *Stack) peek() int    { return s.data[0] }  // read without removing
```

- **Push** — add an element to the top
- **Pop** — remove and return the top element
- **Peek** — read the top element without removing it

</details>

---

### Q3 🟢 — What should `pop()` do if called on an empty stack?

**A)** Return 0  
**B)** Return -1  
**C)** Handle it gracefully — either return an error or do nothing; never crash  
**D)** It will never happen, so no check is needed  

<details><summary>💡 Answer</summary>

**C) Handle it gracefully — never crash**

```go
func (s *Stack) pop() (int, bool) {
    if s.isEmpty() {
        return 0, false  // signal that pop failed
    }
    top := s.data[0]
    s.data = s.data[1:]
    return top, true
}

// In instructions.go — pa() must guard against empty b:
func pa(a, b *Stack) {
    if b.isEmpty() { return }  // nothing to push — do nothing
    a.push(b.pop())
}
```

In push-swap, instructions like `pa` and `pb` are called on a potentially empty stack. Your implementation must silently do nothing rather than panicking on an empty pop.

</details>

---

### Q4 🟡 — In Go, which slice layout makes it easiest to implement a stack where index 0 is the top?

**A)** Append to the end and read the last index  
**B)** Append to the front and read index 0  
**C)** Use index 0 as the top — push prepends, pop reads and removes index 0  
**D)** Use a linked list instead of a slice  

<details><summary>💡 Answer</summary>

**C) Index 0 = top — prepend on push, remove from front on pop**

```go
// Push: prepend value to make it the new top
func (s *Stack) push(val int) {
    s.data = append([]int{val}, s.data...)
}

// Pop: remove from front (the top)
func (s *Stack) pop() int {
    top := s.data[0]
    s.data = s.data[1:]
    return top
}

// Stack after pushing 1, 2, 3:
// s.data = [3, 2, 1]
//           ↑ top
```

This layout means `s.data[0]` is always the top — it maps naturally to stack diagrams and makes instruction logic easier to read and reason about.

</details>

---

### Q5 🟡 — What is the time complexity of a push operation on a stack implemented with a Go slice?

**A)** O(n) — you must shift all elements  
**B)** O(1) amortized — appending to the end of a slice (if top = last index)  
**C)** O(log n)  
**D)** O(n²)  

<details><summary>💡 Answer</summary>

**B) O(1) amortized — append to the end is efficient**

```go
// Efficient: append to end, treat last index as top
func (s *Stack) push(val int) {
    s.data = append(s.data, val)  // O(1) amortized
}
func (s *Stack) pop() int {
    top := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]  // O(1)
    return top
}
```

Go's `append` doubles the backing array when full, making the average cost O(1). Prepending (shifting all elements forward) would be O(n). For push-swap, either layout works — pick one and stay consistent.

</details>

---

### Q6 🔴 — You have a stack with elements `[5, 3, 1]` (top to bottom). After one `rotate` operation, what is the new order?

**A)** `[1, 5, 3]`  
**B)** `[3, 1, 5]`  
**C)** `[1, 3, 5]`  
**D)** `[5, 3, 1]` — no change  

<details><summary>💡 Answer</summary>

**B) `[3, 1, 5]`**

```
Before rotate:  [5, 3, 1]  (5 is top)
                 ↑ top

After rotate (ra):  top element moves to the bottom
                [3, 1, 5]  (3 is now top, 5 went to bottom)
                 ↑ top
```

`ra` (rotate a) shifts every element up by one — the first element becomes the last. It's like taking a card off the top of a deck and placing it at the bottom.

</details>

---

## 📋 SECTION 2: THE 11 INSTRUCTIONS (8 Questions)

### Q7 🟢 — What does `pb` do?

**A)** Push the top of stack b to the top of stack a  
**B)** Push the top of stack a to the top of stack b  
**C)** Swap the top two elements of stack b  
**D)** Rotate stack b  

<details><summary>💡 Answer</summary>

**B) Push the top of stack a to the top of stack b**

```
Before pb:
a = [2, 1, 3]   b = [4, 5]
     ↑ top            ↑ top

After pb:
a = [1, 3]       b = [2, 4, 5]
     ↑ top             ↑ top  ← 2 moved from a to b
```

`pa` and `pb` are your only tools to move elements between stacks. All other instructions operate within a single stack. When b is empty, `pb` still works — it simply moves the top of a to the empty b.

</details>

---

### Q8 🟢 — What is the difference between `ra` and `rra`?

**A)** `ra` rotates left; `rra` rotates right  
**B)** `ra` moves the top element to the bottom; `rra` moves the bottom element to the top  
**C)** `ra` swaps the top two; `rra` reverses the whole stack  
**D)** No difference — they produce the same result  

<details><summary>💡 Answer</summary>

**B) `ra` = top → bottom; `rra` = bottom → top**

```
Stack: [A, B, C, D]  (A is top)

After ra:   [B, C, D, A]   ← A went to bottom
After rra:  [D, A, B, C]   ← D came to top
```

Think of the stack as a circular ring — `ra` spins it clockwise, `rra` spins it counter-clockwise. If you apply `ra` and then `rra` on the same stack, you get back to the original order.

</details>

---

### Q9 🟢 — What does `ss` do?

**A)** Swap the tops of both stacks simultaneously  
**B)** Sort both stacks  
**C)** Rotate both stacks simultaneously  
**D)** Push to both stacks simultaneously  

<details><summary>💡 Answer</summary>

**A) Swap the tops of both stacks simultaneously — equivalent to `sa` + `sb` in one instruction**

```go
func ss(a, b *Stack) {
    sa(a)  // swap top two of a
    sb(b)  // swap top two of b
}

// Both swaps happen — this counts as ONE instruction, not two
```

The combined instructions (`ss`, `rr`, `rrr`) let you apply the same operation on both stacks at once, counting as a single instruction. Using them when applicable reduces your total instruction count.

</details>

---

### Q10 🟡 — Stack a is `[3, 1, 2]` (top to bottom). What is the result after `sa`?

**A)** `[1, 3, 2]`  
**B)** `[3, 2, 1]`  
**C)** `[2, 1, 3]`  
**D)** `[1, 2, 3]`  

<details><summary>💡 Answer</summary>

**A) `[1, 3, 2]`**

```
Before sa:  [3, 1, 2]
             ↑ ↑ top two elements

After sa:   [1, 3, 2]
             ↑ top  ← only the top two are swapped; 2 stays at bottom
```

`sa` only ever touches the top two elements. If the stack has fewer than two elements, `sa` does nothing.

</details>

---

### Q11 🟡 — How many instructions does `rr` count as?

**A)** 2  
**B)** 0  
**C)** 1  
**D)** Depends on stack sizes  

<details><summary>💡 Answer</summary>

**C) 1 — `rr` is a single instruction that rotates both stacks**

```go
func rr(a, b *Stack) {
    ra(a)  // rotate a
    rb(b)  // rotate b
}
// Output: one line "rr" — NOT two lines "ra" then "rb"
```

This is important for optimization. If you need to rotate both stacks in the same direction at the same time, using `rr` instead of `ra` + `rb` saves one instruction in your output. The same applies to `ss` and `rrr`.

</details>

---

### Q12 🟡 — Stack a is `[9, 5, 2, 7]`. After `rra`, what does it look like?

**A)** `[5, 2, 7, 9]`  
**B)** `[7, 9, 5, 2]`  
**C)** `[2, 9, 5, 7]`  
**D)** `[9, 7, 5, 2]`  

<details><summary>💡 Answer</summary>

**B) `[7, 9, 5, 2]`**

```
Before rra: [9, 5, 2, 7]   (9 is top, 7 is bottom)

rra = reverse rotate: last element comes to the top
After rra:  [7, 9, 5, 2]   (7 is now top)
```

`rra` grabs the bottom element (7) and places it at the top. All other elements shift down by one position. Use `rra` when your target element is in the lower half of the stack — it's cheaper than rotating all the way around with `ra`.

</details>

---

### Q13 🔴 — Which instruction sequence correctly sorts a 2-element stack a = `[2, 1]` into ascending order (smallest on top)?

**A)** No instructions needed  
**B)** `ra`  
**C)** `sa`  
**D)** `pa` then `pb`  

<details><summary>💡 Answer</summary>

**C) `sa`**

```
Before: a = [2, 1]   ← 2 is on top, but 1 should be on top for ascending order
                         (smallest at top = ascending when read top to bottom)

After sa: a = [1, 2]  ← sorted: 1 (top) < 2 (bottom) ✓
```

Wait — ascending order for push-swap means the smallest element is at the top of stack a when done. One `sa` is both necessary and sufficient for any 2-element unsorted stack.

</details>

---

### Q14 🔴 — What is the maximum number of instructions needed to sort any 3-element stack?

**A)** 6  
**B)** 5  
**C)** 3  
**D)** 2  

<details><summary>💡 Answer</summary>

**C) 3 — no ordering of 3 elements requires more than 3 instructions**

```
All 6 orderings of [1, 2, 3] and their solutions:

[1,2,3] → already sorted → 0 instructions
[1,3,2] → sa             → 1 instruction
[2,1,3] → sa, ra (wait—check) actually: rra, sa or ra → varies
[2,3,1] → ra             → 1 instruction  (becomes [3,1,2]... recalculate all by hand!)
[3,1,2] → rra            → 1 instruction
[3,2,1] → sa, rra        → 2 instructions
```

Work through all 6 cases by hand on paper before writing a single line of code. You need to find the minimum instruction sequence for each — this is the foundation of sort3.

</details>

---

## 📋 SECTION 3: SORTING COMPLEXITY (5 Questions)

### Q15 🟢 — What does O(n²) mean in terms of operations for a list of 100 elements?

**A)** About 100 operations  
**B)** About 200 operations  
**C)** About 10,000 operations  
**D)** About 1,000 operations  

<details><summary>💡 Answer</summary>

**C) About 10,000 operations**

```
O(n²) means operations grow proportionally to n squared:

n = 10    →   100 operations
n = 100   →  10,000 operations
n = 500   → 250,000 operations

Push-swap budget for 500 elements: ≤ 5,500 instructions
O(n²) for 500 elements:            250,000 operations  ← way too slow

This is why bubble sort won't pass push-swap grading.
```

A naïve O(n²) sorting approach will generate far too many instructions. You need a more efficient algorithm — O(n log n) or even O(n) strategies — to stay within the instruction budget.

</details>

---

### Q16 🟢 — Which of these sorting algorithms has the best average-case complexity?

**A)** Bubble Sort — O(n²)  
**B)** Selection Sort — O(n²)  
**C)** Merge Sort — O(n log n)  
**D)** Insertion Sort — O(n²)  

<details><summary>💡 Answer</summary>

**C) Merge Sort — O(n log n)**

```
For n = 500:

Bubble Sort:   500² = 250,000 comparisons
Merge Sort:    500 × log₂(500) ≈ 500 × 9 = 4,500 comparisons

This is why algorithms that approximate merge sort or radix sort
are used in push-swap for large inputs.
```

O(n log n) is the theoretical lower bound for comparison-based sorting. Radix sort achieves O(n × k) where k is the number of bit positions — often better in practice.

</details>

---

### Q17 🟡 — For push-swap with 100 elements, what is the acceptable instruction budget?

**A)** ≤ 100 instructions  
**B)** ≤ 300 instructions  
**C)** ≤ 700 instructions  
**D)** ≤ 2000 instructions  

<details><summary>💡 Answer</summary>

**C) ≤ 700 instructions for 100 elements**

```
Grading thresholds (approximate, check your spec):

  5 elements  → ≤ 12 instructions   (full score)
100 elements  → ≤ 700 instructions  (full score)
500 elements  → ≤ 5,500 instructions (full score)

Test yourself:
ARG=$(python3 -c "import random; l=list(range(100)); random.shuffle(l); print(' '.join(map(str,l)))")
./push-swap "$ARG" | wc -l
```

Always verify with `./push-swap "$ARG" | ./checker "$ARG"` to confirm `OK` before counting instructions. A low count that produces `KO` scores zero.

</details>

---

### Q18 🟡 — What is the key insight behind radix sort that makes it different from comparison-based sorts?

**A)** It compares pairs of elements like bubble sort  
**B)** It sorts by processing digits/bits position by position — no direct comparisons between elements  
**C)** It uses recursion like merge sort  
**D)** It sorts by insertion like insertion sort  

<details><summary>💡 Answer</summary>

**B) Sorts digit by digit — no element comparisons**

```
Radix sort on [3, 1, 2, 0] (binary: 11, 01, 10, 00)

Pass 1 — bit 0 (LSB):
  bit=0 → stay in a:  [2(10), 0(00)]
  bit=1 → push to b:  [1(01), 3(11)]
  Push b back → [0, 2, 1, 3]... (normalize first!)

Pass 2 — bit 1:
  bit=0 → stay in a: [1(01), 0(00)]
  bit=1 → push to b: [3(11), 2(10)]
  Push b back → [0, 1, 2, 3] ✓ sorted!
```

For push-swap, you first normalize values to 0..N-1 so bitwise operations make sense, then use each bit position to decide whether to push to b or rotate in a.

</details>

---

### Q19 🔴 — Why must you normalize values before applying radix sort in push-swap?

**A)** To make values fit in an integer  
**B)** Because radix sort requires values to be in the range 0..N-1 — arbitrary integers have inconsistent bit-length and distribution, making bit passes unreliable  
**C)** Normalization is optional — radix works on any integers  
**D)** To detect duplicates  

<details><summary>💡 Answer</summary>

**B) Radix sort requires 0..N-1 — normalize arbitrary integers first**

```go
func normalize(a *Stack) {
    // Extract all values, sort them, assign rank by position
    values := make([]int, a.size())
    for i := range values { values[i] = a.data[i] }

    sorted := make([]int, len(values))
    copy(sorted, values)
    sort.Ints(sorted)  // sort a copy

    // Build rank map: original value → rank (0 = smallest)
    rank := make(map[int]int)
    for i, v := range sorted { rank[v] = i }

    // Replace values in stack with their ranks
    for i := range a.data { a.data[i] = rank[a.data[i]] }
}

// Before: [42, -3, 1000, 7]
// After:  [2,   0,    3, 1]   ← ranks 0–3
```

Without normalization, the number of bit passes you'd need is determined by the largest value (could be millions of bits), not by N.

</details>

---

## 📋 SECTION 4: ALGORITHM STRATEGY (5 Questions)

### Q20 🟢 — For 3 elements in stack a, what is the correct approach in push-swap?

**A)** Use radix sort  
**B)** Use a hardcoded lookup — identify the ordering of 3 elements and apply the known minimum instruction sequence  
**C)** Use bubble sort  
**D)** Push all to b and pull back in order  

<details><summary>💡 Answer</summary>

**B) Hardcoded lookup — there are only 6 orderings**

```go
func sort3(a *Stack) []string {
    x, y, z := a.data[0], a.data[1], a.data[2]

    // Identify ordering and return minimal instruction sequence:
    if x > y && y < z && x < z { return []string{"sa"} }
    if x > y && y > z           { return []string{"sa", "rra"} }
    if x > y && y < z && x > z { return []string{"ra"} }
    if x < y && y > z && x < z { return []string{"rra", "sa"} } // ... etc
    // (work all 6 cases out by hand on paper!)
    return nil  // already sorted
}
```

There are only 3! = 6 possible orderings. For each one, work out the minimum instruction sequence by hand on paper before writing code.

</details>

---

### Q21 🟡 — What is the chunk-based strategy for sorting large inputs?

**A)** Divide elements into sorted groups in place using ra/rra  
**B)** Divide the value range into chunks, push elements chunk by chunk from a to b (smallest chunks first), then pull back in descending order  
**C)** Push all elements to b, then insertion-sort back into a  
**D)** Use quicksort pivots with pa/pb  

<details><summary>💡 Answer</summary>

**B) Divide value range into chunks, push to b smallest-first, pull back largest-first**

```
Example: 100 elements → 5 chunks of 20

Chunk 0: values 0–19   → push to b first (go to bottom of b)
Chunk 1: values 20–39  → push to b next
...
Chunk 4: values 80–99  → push to b last (sit on top of b)

Pull back: always grab the current maximum from b using rotations
→ a fills from largest to smallest? No — we pull in order so a ends smallest-on-top
```

More chunks = fewer rotations per element, but more pa/pb operations. Tuning chunk count for your N is the key optimization. Around 5 chunks for N=100 and 11 for N=500 are common starting points.

</details>

---

### Q22 🟡 — When rotating a stack to bring element at position `i` to the top, how do you decide between `ra` and `rra`?

**A)** Always use `ra`  
**B)** Use `ra` if `i <= size/2`, use `rra` if `i > size/2` — choose whichever requires fewer rotations  
**C)** Always use `rra`  
**D)** Use `ra` for even positions, `rra` for odd  

<details><summary>💡 Answer</summary>

**B) Use `ra` if element is in the top half, `rra` if in the bottom half**

```go
func rotateTo(s *Stack, targetIndex int, isA bool) []string {
    size := s.size()
    var instructions []string

    if targetIndex <= size/2 {
        // Closer to top — rotate up with ra
        for i := 0; i < targetIndex; i++ {
            if isA { instructions = append(instructions, "ra") } else { ... }
        }
    } else {
        // Closer to bottom — rotate down with rra (fewer moves)
        for i := 0; i < size-targetIndex; i++ {
            if isA { instructions = append(instructions, "rra") } else { ... }
        }
    }
    return instructions
}
```

This is the single most impactful optimization in push-swap. Failing to use `rra` when it's cheaper can double your instruction count.

</details>

---

### Q23 🔴 — Why does the sort strategy change depending on input size (2 vs 3 vs 5 vs large)?

**A)** It doesn't — the same algorithm works for all sizes  
**B)** Larger inputs need more passes and general strategies; small inputs are better handled by hardcoded optimal sequences that can't be beat algorithmically  
**C)** Smaller inputs require radix sort; larger inputs need simple sorts  
**D)** It's a matter of style, not correctness  

<details><summary>💡 Answer</summary>

**B) Small inputs: hardcoded optimal; large inputs: general efficient algorithm**

```go
func sort(a, b *Stack) []string {
    switch a.size() {
    case 0, 1:
        return nil           // already sorted by definition
    case 2:
        return sort2(a)      // at most 1 instruction
    case 3:
        return sort3(a)      // at most 3 instructions — hardcoded
    case 4, 5:
        return sort5(a, b)   // push 1–2 to b, sort3, push back
    default:
        return sortLarge(a, b) // radix or chunk — general strategy
    }
}
```

For N=3, the optimal solution is known and hardcoded. For N=500, no one can enumerate all cases — you need a systematic algorithm. Using `sortLarge` on 3 elements would be correct but wasteful.

</details>

---

### Q24 🔴 — What is the purpose of the `checker` program and when should you build it?

**A)** It generates the push-swap instruction sequence  
**B)** It validates that a given instruction sequence correctly sorts the stack — build it BEFORE push-swap so it becomes your testing tool  
**C)** It benchmarks performance  
**D)** It generates random inputs for testing  

<details><summary>💡 Answer</summary>

**B) Validates instruction sequences — build it first to test push-swap**

```bash
# Workflow once both are built:
ARG="5 3 1 4 2"
./push-swap "$ARG" | ./checker "$ARG"
# OK  ← push-swap output is a valid sort
# KO  ← push-swap output does NOT sort the stack — bug in your sort

# Checker also catches invalid instructions:
echo "invalid_instruction" | ./checker "1 2 3"
# Error
```

Building `checker` first gives you a tool to verify every sorting function you write. Without it, you're debugging blindly. Think of checker as your automated unit test harness.

</details>

---

## 📋 SECTION 5: INPUT VALIDATION & GO FUNDAMENTALS (6 Questions)

### Q25 🟢 — How do you read command-line arguments in Go?

**A)** `os.Stdin`  
**B)** `os.Args` — a slice of strings where `os.Args[0]` is the program name  
**C)** `flag.Parse()`  
**D)** `fmt.Scan()`  

<details><summary>💡 Answer</summary>

**B) `os.Args` — index 0 is the program name, index 1+ are the arguments**

```go
// ./push-swap 3 2 1
// os.Args = ["./push-swap", "3", "2", "1"]

args := os.Args[1:]  // skip program name → ["3", "2", "1"]

// Handle the case where all numbers are in one string:
// ./push-swap "3 2 1"
// os.Args[1] = "3 2 1"  (one string with spaces)
// Must split: strings.Fields(os.Args[1]) → ["3", "2", "1"]
```

Both `./push-swap 3 2 1` and `./push-swap "3 2 1"` must produce the same result. Join all args into one string and split on whitespace to handle both cases uniformly.

</details>

---

### Q26 🟢 — How do you print an error message to stderr in Go?

**A)** `fmt.Println("Error")`  
**B)** `fmt.Fprintln(os.Stderr, "Error")`  
**C)** `log.Fatal("Error")`  
**D)** `os.Exit(1)`  

<details><summary>💡 Answer</summary>

**B) `fmt.Fprintln(os.Stderr, "Error")`**

```go
func handleError() {
    fmt.Fprintln(os.Stderr, "Error")
    os.Exit(1)
}

// This matters because the checker reads stdout (the instruction list)
// Error messages must go to stderr so they don't pollute the instruction output

// Wrong — this would pipe "Error" into checker:
fmt.Println("Error")    // goes to stdout — breaks checker

// Right — stderr is separate from stdout:
fmt.Fprintln(os.Stderr, "Error")  // goes to stderr — checker unaffected
```

In a pipe like `./push-swap "$ARG" | ./checker "$ARG"`, only stdout is piped. Stderr goes directly to the terminal, keeping your error messages visible without corrupting checker's input.

</details>

---

### Q27 🟡 — What are the three error cases push-swap must detect and how should it report them?

**A)** Empty input, too many arguments, wrong data type — print `error` to stdout  
**B)** Non-integer argument, duplicate value, integer overflow — print `Error` to stderr and exit  
**C)** Negative numbers, floating point numbers, strings — print `Error` to stdout  
**D)** Only non-integers need to be caught — duplicates and overflow are permitted  

<details><summary>💡 Answer</summary>

**B) Non-integer, duplicate, overflow — print `Error` to stderr**

```go
func parseArgs(args []string) ([]int, error) {
    var joined []string
    for _, a := range args { joined = append(joined, strings.Fields(a)...) }

    seen := make(map[int]bool)
    var result []int

    for _, token := range joined {
        // Case 1: non-integer
        n, err := strconv.Atoi(token)
        if err != nil { return nil, fmt.Errorf("Error") }

        // Case 2: integer overflow (Atoi handles int64 range on 64-bit systems)
        // Check against int32 range if spec requires it:
        if n > 2147483647 || n < -2147483648 { return nil, fmt.Errorf("Error") }

        // Case 3: duplicate
        if seen[n] { return nil, fmt.Errorf("Error") }
        seen[n] = true
        result = append(result, n)
    }
    return result, nil
}
```

The spec says "Error" (capital E) — check your spec for the exact string. Zero arguments should print nothing and exit silently.

</details>

---

### Q28 🟡 — What should `push-swap` print if the input is already sorted?

**A)** `Already sorted`  
**B)** Nothing — empty output  
**C)** `0`  
**D)** `OK`  

<details><summary>💡 Answer</summary>

**B) Nothing — absolutely no output**

```go
func main() {
    args := os.Args[1:]
    if len(args) == 0 { os.Exit(0) }  // zero args: silent exit

    nums, err := parseArgs(args)
    if err != nil { fmt.Fprintln(os.Stderr, "Error"); os.Exit(1) }

    a := buildStack(nums)

    if isSorted(&a) { os.Exit(0) }  // already sorted: print nothing and exit

    instructions := sort(&a, &b)
    for _, inst := range instructions { fmt.Println(inst) }
}
```

Any output for an already-sorted input would be picked up by checker, which would try to execute it as an instruction and may fail. An empty instruction sequence on a sorted stack is correct — checker will confirm `OK`.

</details>

---

### Q29 🟡 — How do you read instructions from stdin line by line in the checker program?

**A)** `fmt.Scan(&line)`  
**B)** `bufio.Scanner` with a loop calling `.Scan()` until it returns false  
**C)** `os.Stdin.Read()`  
**D)** `ioutil.ReadAll(os.Stdin)` then split on newlines  

<details><summary>💡 Answer</summary>

**B) `bufio.Scanner` — reads one line at a time until EOF**

```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    line := scanner.Text()  // one instruction per line, no newline char

    switch line {
    case "pa":  pa(&a, &b)
    case "pb":  pb(&a, &b)
    case "sa":  sa(&a)
    case "sb":  sb(&b)
    case "ss":  ss(&a, &b)
    case "ra":  ra(&a)
    case "rb":  rb(&b)
    case "rr":  rr(&a, &b)
    case "rra": rra(&a)
    case "rrb": rrb(&b)
    case "rrr": rrr(&a, &b)
    default:
        fmt.Fprintln(os.Stderr, "Error")
        os.Exit(1)
    }
}
```

`bufio.Scanner` is the idiomatic Go way to read line-by-line. It handles EOF cleanly — the loop exits naturally when stdin closes.

</details>

---

### Q30 🔴 — You run `./push-swap "1 2 3" | ./checker "1 2 3"` and get `KO`. What is wrong and how do you debug it?

**A)** The checker is broken — ignore it  
**B)** push-swap outputs instructions that do not correctly sort the stack — simulate them step by step by hand or add debug output to checker to print stack state after each instruction  
**C)** The input is wrong — try different numbers  
**D)** `KO` is the correct result for an already-sorted input  

<details><summary>💡 Answer</summary>

**B) push-swap instructions are incorrect — simulate step by step to find the bug**

```bash
# Step 1: See what instructions push-swap generates
./push-swap "1 2 3"
# e.g., outputs: ra, sa  ← what are these doing?

# Step 2: Simulate manually
# Stack a = [1, 2, 3]  (1 is top)
# After ra:  [2, 3, 1]
# After sa:  [3, 2, 1]  ← NOT sorted! Bug found.

# Step 3: Add debug output to checker temporarily
fmt.Fprintf(os.Stderr, "After %s: a=%v b=%v\n", line, a.data, b.data)

# Step 4: Check isSorted logic
func isSorted(a *Stack) bool {
    for i := 0; i < len(a.data)-1; i++ {
        if a.data[i] > a.data[i+1] { return false }
    }
    return true
}
```

`KO` means either the instructions are wrong, or your `isSorted` function has a bug. Simulate step by step — there's no shortcut. Build checker with debug output before you test push-swap.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 28–30 ✅ | **Exceptional** — stacks, instructions, algorithms, and Go all mastered. You're ready to code. |
| 24–27 ✅ | **Ready** — review any missed sections carefully before starting. |
| 18–23 ⚠️ | **Study first** — sorting algorithms and the 11 instructions need more attention. |
| Below 18 ❌ | **Not ready** — review stacks, Big-O complexity, and Go slice fundamentals before beginning. |

---

## 🔍 Review Map

| Missed | Topic to Study |
|---|---|
| Q1–Q6 | Stack LIFO concept, push/pop/peek, rotate vs reverse-rotate, Go slice layout |
| Q7–Q14 | All 11 instructions, pa/pb/sa/sb/ss/ra/rb/rr/rra/rrb/rrr — simulate each by hand |
| Q15–Q19 | Big-O notation (n, n log n, n²), instruction budgets, radix sort, normalization |
| Q20–Q24 | sort2/sort3/sort5 hardcoded strategies, chunk sort, rotation direction optimization, checker purpose |
| Q25–Q30 | `os.Args`, `fmt.Fprintln(os.Stderr)`, error cases, `bufio.Scanner`, sorted detection, debugging KO |