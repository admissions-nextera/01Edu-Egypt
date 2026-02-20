# Push-Swap Project Guide

> **Before you start:** Read about stacks as a data structure and work through the example in the spec by hand on paper. Every instruction must make sense to you before you write any code. You cannot optimize something you do not understand.

---

## Objectives

By completing this project you will learn:

1. **Stack Data Structure** — How a stack works and how to implement push, pop, peek, and rotate operations
2. **Sorting Algorithms** — Understanding sorting complexity and why the number of operations matters
3. **Non-Comparative Sorting** — Thinking about sorting differently from bubble sort or quicksort
4. **Algorithm Design** — Choosing the right strategy based on input size
5. **Two-Program Architecture** — Building a generator (push-swap) and a validator (checker) that work together
6. **Input Validation** — Detecting and reporting errors cleanly

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Stacks
- What a stack is (LIFO — Last In, First Out)
- Push, pop, peek operations
- Search: **"stack data structure explained"**

### 2. The 11 Instructions
Before writing any code, simulate all 11 instructions by hand on a small example:
- `pa`, `pb` — push between stacks
- `sa`, `sb`, `ss` — swap top two elements
- `ra`, `rb`, `rr` — rotate up (first becomes last)
- `rra`, `rrb`, `rrr` — reverse rotate (last becomes first)

### 3. Sorting Complexity
- What O(n), O(n log n), O(n²) means in terms of number of operations
- Search: **"sorting algorithm complexity comparison"**
- https://en.wikipedia.org/wiki/Sorting_algorithm

### 4. Go Fundamentals
- Slices as stacks (append, slice operations)
- `os.Args` for arguments
- `fmt.Fprintln(os.Stderr, ...)` for error output
- `bufio.Scanner` for reading stdin

---

## Project Structure

```
push-swap/
├── main_pushswap.go    ← push-swap program entry point
├── main_checker.go     ← checker program entry point
├── stack.go            ← stack operations
├── instructions.go     ← the 11 instruction implementations
├── validate.go         ← argument parsing and validation
├── sort.go             ← sorting algorithms
├── go.mod
└── push-swap_test.go
```

---

## Milestone 1 — Build the Stack and All 11 Instructions

**Goal:** You have a working stack implementation and all 11 instructions execute correctly.

**Questions to answer before writing anything:**
- How will you represent a stack in Go? A slice where index 0 is the top?
- What happens when you call `pa` on an empty stack b? (Nothing — handle gracefully.)
- What does `ra` do to a stack with only one element? (Nothing changes.)
- What is the difference between `ra` and `rra`?

**Code Placeholder:**
```go
// stack.go

type Stack struct {
    // A slice of integers representing the stack
    // Index 0 = top of the stack
}

func (s *Stack) push(val int)   { /* add to top */ }
func (s *Stack) pop() int       { /* remove and return top */ }
func (s *Stack) peek() int      { /* return top without removing */ }
func (s *Stack) isEmpty() bool  { /* return true if empty */ }
func (s *Stack) size() int      { /* return number of elements */ }
```

```go
// instructions.go

func pa(a, b *Stack) { /* move top of b to top of a */ }
func pb(a, b *Stack) { /* move top of a to top of b */ }
func sa(a *Stack)    { /* swap top two elements of a */ }
func sb(b *Stack)    { /* swap top two elements of b */ }
func ss(a, b *Stack) { /* sa and sb */ }
func ra(a *Stack)    { /* rotate a up: first becomes last */ }
func rb(b *Stack)    { /* rotate b up */ }
func rr(a, b *Stack) { /* ra and rb */ }
func rra(a *Stack)   { /* reverse rotate a: last becomes first */ }
func rrb(b *Stack)   { /* reverse rotate b */ }
func rrr(a, b *Stack){ /* rra and rrb */ }
```

**Verify:** Work through the example in the spec step by step. After each instruction, print both stacks and confirm they match the spec's diagram.

---

## Milestone 2 — Parse and Validate Arguments

**Goal:** Both programs correctly parse the integer list from arguments, detect all errors, and print `Error` to stderr for invalid input.

**Questions to answer:**
- Arguments can arrive as `./push-swap 2 1 3` OR as `./push-swap "2 1 3"` — how do you handle both?
- What are the three error cases? (Non-integer, duplicate, integer overflow)
- What is Go's `int` size on 64-bit systems? Does it matter for overflow detection?
- What does "display nothing" mean for zero arguments — no output at all, not even a newline?

**Code Placeholder:**
```go
// validate.go

func parseArgs(args []string) ([]int, error) {
    // 1. Join all args into one string and split on spaces
    //    This handles both "2 1 3" as one arg and 2 1 3 as separate args

    // 2. For each token:
    //    Convert to integer — return error if not a valid integer
    //    Check if value already seen — return error if duplicate

    // 3. Return the slice of integers
}
```

**Verify:**
```bash
./push-swap "0 one 2 3"      # prints Error to stderr
./push-swap "1 2 2 3"        # prints Error (duplicate)
./push-swap                   # prints nothing
./push-swap "3 2 1"          # no error, proceeds to sort
```

---

## Milestone 3 — Build the Checker Program

**Goal:** checker validates that a sequence of instructions sorts the stack correctly. Build and test this before push-swap — it becomes your verification tool.

**Questions to answer:**
- How do you read instructions from stdin line by line until EOF?
- How do you map the string `"pa"` to the `pa()` function?
- What constitutes an invalid instruction? (Wrong name, wrong format.)
- When is the result `OK`? (Stack a is sorted ascending AND stack b is empty.)

**Code Placeholder:**
```go
// main_checker.go

func main() {
    // 1. Parse and validate arguments → build stack a
    //    If no arguments: exit silently
    //    If invalid: print Error to stderr and exit

    // 2. Read instructions from stdin line by line
    //    For each instruction:
    //      If unrecognized: print Error to stderr and exit
    //      Execute the instruction on the stacks

    // 3. Check if stack a is sorted ascending and stack b is empty
    //    If yes: print "OK"
    //    If no:  print "KO"
}

func isSorted(a *Stack) bool {
    // Return true if every element is <= the next element
}
```

**Verify:**
```bash
echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"   # OK
./checker "3 2 1 0" <<< $'sa\nrra\npb'                  # KO
echo -e "rra\npb\nsa\n" | ./checker "3 2 one 0"         # Error
```

---

## Milestone 4 — Trivial Cases in push-swap

**Goal:** Handle the cases where no sorting is needed or only one move is needed.

**Questions to answer:**
- If the stack is already sorted, what should push-swap print? (Nothing.)
- For 2 elements, what is the maximum number of instructions needed?
- For 3 elements, how many possible orderings are there and what is the optimal instruction sequence for each?

**Code Placeholder:**
```go
// sort.go

func alreadySorted(a *Stack) bool {
    // Check if stack a is in ascending order
}

func sort2(a *Stack) []string {
    // If already sorted: return empty
    // If not: return ["sa"]
}

func sort3(a *Stack) []string {
    // There are 6 possible orderings of 3 elements
    // For each ordering, determine and return the minimal instruction sequence
    // Maximum 3 instructions for any 3-element case
}
```

**Verify:** Test all 6 possible orderings of 3 elements manually. Confirm the minimum instruction count for each.

---

## Milestone 5 — Sorting Algorithm for Small Inputs (≤ 5)

**Goal:** For 4 and 5 elements, use a strategy that stays within a small instruction budget.

**Questions to answer:**
- For 5 elements, what is a reasonable maximum instruction count? (≤ 12 is achievable.)
- What is a common approach: push the smallest elements to b, sort what remains in a using sort3, then push back in order?
- How do you find the position of the minimum element in a stack?
- How do you rotate a stack to bring a specific element to the top using the fewest rotations (ra vs rra)?

**Code Placeholder:**
```go
// sort.go

func sort5(a, b *Stack, instructions *[]string) {
    // Strategy:
    // 1. Push the two smallest elements to b (pb)
    // 2. Sort the remaining 3 elements in a using sort3
    // 3. Push back from b in the correct order

    // Helper needed:
    // - findMinIndex(s *Stack) int  →  position from top of smallest element
    // - rotateTo(s *Stack, index int, isA bool) []string
    //     returns the cheapest rotation instructions to bring index to top
    //     (ra if index <= size/2, rra otherwise)
}
```

**Verify:**
```bash
ARG="4 67 3 87 23"
./push-swap "$ARG" | wc -l          # should be ≤ 12
./push-swap "$ARG" | ./checker "$ARG"   # OK
```

---

## Milestone 6 — Sorting Algorithm for Large Inputs

**Goal:** For inputs larger than 5, implement an efficient algorithm that minimizes total instruction count.

**Questions to answer before choosing an algorithm:**
- What are the well-known approaches for push-swap? Research: **"push swap algorithm radix"**, **"push swap algorithm turkish sort"**, **"push swap algorithm chunks"**
- For ~100 numbers, what is the acceptable instruction budget? (≤ 700)
- For ~500 numbers, what is the acceptable instruction budget? (≤ 5500)
- Which algorithm gives the best operation count for each size range?

**Two common strategies:**

**Chunk-based sort:**
```
1. Divide the sorted range into chunks (e.g. 5 chunks of 20 for 100 elements)
2. Push elements from a to b chunk by chunk, smallest chunks first
3. For each element being pushed to b, rotate b so larger elements go to bottom
4. Pull back from b in sorted order
```

**Radix sort (bitwise):**
```
1. Normalize values to 0..N-1 (assign rank)
2. For each bit position (LSB to MSB):
   - Elements with 0 in this bit → push to b
   - Elements with 1 in this bit → rotate in a
   - Push all from b back to a
3. Repeat for log2(N) passes
```

**Code Placeholder:**
```go
// sort.go

func sortLarge(a, b *Stack, instructions *[]string) {
    // Choose and implement your strategy here
    // Make every instruction call also append to instructions slice
}

func normalize(a *Stack) {
    // Replace values in a with their rank (0 = smallest, N-1 = largest)
    // This makes bitwise operations meaningful
}
```

**Verify:**
```bash
# Test with 100 random numbers
ARG=$(python3 -c "import random; l=list(range(100)); random.shuffle(l); print(' '.join(map(str,l)))")
./push-swap "$ARG" | wc -l               # ≤ 700
./push-swap "$ARG" | ./checker "$ARG"    # OK

# Test with 500 random numbers  
ARG=$(python3 -c "import random; l=list(range(500)); random.shuffle(l); print(' '.join(map(str,l)))")
./push-swap "$ARG" | wc -l               # ≤ 5500
./push-swap "$ARG" | ./checker "$ARG"    # OK
```

---

## Milestone 7 — Collect and Print Instructions

**Goal:** push-swap prints the instruction sequence to stdout, one instruction per line.

**Questions to answer:**
- Should your sort functions print instructions as they go, or collect them and print at the end?
- Is there any optimization pass you can do on the instruction list before printing? (e.g. `ra` followed by `rra` cancel out — can you detect and remove such pairs?)
- What does the output look like for an already-sorted input? (Empty — no output at all.)

**Code Placeholder:**
```go
// main_pushswap.go

func main() {
    // 1. Parse and validate arguments
    //    No arguments → exit silently
    //    Invalid → print Error to stderr

    // 2. Build stack a from the arguments

    // 3. If already sorted: exit (print nothing)

    // 4. Run the appropriate sort based on size:
    //    size == 2: sort2
    //    size == 3: sort3
    //    size <= 5: sort5
    //    size >  5: sortLarge

    // 5. Print each instruction on its own line
}
```

---

## Milestone 8 — Unit Tests

**Goal:** Every instruction and the core sorting logic is tested.

**Code Placeholder:**
```go
// push-swap_test.go

func TestPa(t *testing.T) {
    // Set up a with elements, b with elements
    // Call pa
    // Assert top of a is what was top of b
    // Assert b lost its top element
}

func TestSort3(t *testing.T) {
    // Test all 6 orderings of [1,2,3]
    // For each: build stack, call sort3, verify sorted and instruction count ≤ 3
}

func TestIsSorted(t *testing.T) {
    // Test sorted input → true
    // Test unsorted input → false
    // Test single element → true
    // Test empty stack → true
}

func TestParseArgs(t *testing.T) {
    // Test valid input
    // Test non-integer → error
    // Test duplicate → error
    // Test empty → no error, empty slice
}
```

---

## Debugging Checklist

- Does checker return KO for a sequence that looks correct? Simulate the instructions step by step by hand and find where the stacks diverge from your expectation.
- Is your instruction count too high for 100 elements? Profile which part of your algorithm generates the most instructions — usually it is the "bring element to top" rotation that can be improved by choosing ra vs rra based on position.
- Does push-swap print instructions but checker still says KO? Run `./push-swap "$ARG" | ./checker "$ARG"` — if KO, one of your instructions does the wrong thing. Add debug prints inside the checker to show stack state after each instruction.
- Are duplicates not being caught? Make sure your validation checks for duplicates after parsing all values, not just adjacent ones.
- Does the program crash on an empty stack operation? Every instruction that reads from a stack must check `isEmpty()` first.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Args, stderr output | https://pkg.go.dev/os |
| `fmt` | Print instructions, errors | https://pkg.go.dev/fmt |
| `bufio` | Read instructions from stdin in checker | https://pkg.go.dev/bufio |
| `strconv` | Parse integer arguments | https://pkg.go.dev/strconv |
| `strings` | Split argument string on spaces | https://pkg.go.dev/strings |

---

## Submission Checklist

- [ ] All 11 instructions implemented and correct
- [ ] Stack handles empty-stack operations gracefully
- [ ] Arguments parsed correctly for both `"2 1 3"` and `2 1 3` formats
- [ ] Non-integer arguments print `Error` to stderr
- [ ] Duplicate arguments print `Error` to stderr
- [ ] Zero arguments: both programs exit silently
- [ ] Already-sorted input: push-swap prints nothing
- [ ] sort2 uses at most 1 instruction
- [ ] sort3 uses at most 3 instructions
- [ ] sort5 uses at most 12 instructions
- [ ] 100 elements sorted in ≤ 700 instructions
- [ ] 500 elements sorted in ≤ 5500 instructions
- [ ] `./push-swap "$ARG" | ./checker "$ARG"` always returns OK
- [ ] checker prints OK or KO on valid sorted/unsorted results
- [ ] checker prints Error for invalid instructions
- [ ] Unit tests written for instructions, validation, and sort cases