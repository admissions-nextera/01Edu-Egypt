# Push-Swap Project Guide - Learn to Think Like a Programmer

## 📋 What This Guide IS and IS NOT

### ❌ This Guide Will NOT:
- Give you code to copy
- Provide ready-to-use functions
- Solve the problem for you
- Let you skip the learning process

### ✅ This Guide WILL:
- Teach you HOW to think about the problem
- Guide you through breaking it down
- Point you to concepts you need to learn
- Ask questions that help you discover solutions
- Show you WHERE to research, not WHAT to write

**Remember**: The goal is to make YOU a programmer who can solve problems, not to finish this project quickly.

---

## 🎯 Before You Start: Essential Understanding

### What is This Project Really Teaching You?

This isn't just about sorting numbers. It's teaching you:
1. **How to work with constraints** (only certain operations allowed)
2. **How to optimize** (minimize moves)
3. **How to think algorithmically** (step-by-step problem solving)
4. **How to test systematically** (verify your solution works)

### The Big Question
**"How do I sort numbers when I can only use these 11 operations?"**

Don't rush to code. First, understand the problem deeply.

---

## 📚 Phase 1: Understanding (NO CODE YET!)

### Step 1: Play With Stacks on Paper

**Activity**: Get two pieces of paper. Label them "Stack A" and "Stack B".

Write these numbers on small pieces of paper: 3, 1, 2

Put them on Stack A (3 on top).

**Now try**:
- What happens if you "sa" (swap first two)?
- What happens if you "ra" (rotate)?
- What happens if you "pb" (push to B)?

**Spend 30 minutes** trying different operations. Feel how stacks work.

**Questions to answer for yourself**:
- How do I get the smallest number to the top?
- How do I get the largest number to the bottom?
- When should I use Stack B?

### Step 2: Understanding Each Operation

For EACH of the 11 operations, you need to understand:
- What it does physically to the stack
- When it would be useful
- What problem it solves

**Exercise**: 
Draw "before and after" for each operation. Do this by hand.

Example format (you do the rest):
```
Operation: sa (swap first two of stack a)
Before:     After:
[5]         [3]
[3]         [5]
[1]         [1]
```

**Don't move to code until you can explain what each operation does in your own words.**

---

## 🧠 Phase 2: Problem Decomposition

### The Wrong Approach
❌ "I'll just start coding the checker program"
❌ "Let me look for a sorting algorithm online"
❌ "I'll implement all operations first"

### The Right Approach
✅ "What's the simplest version of this problem?"
✅ "Can I solve it for 2 numbers first?"
✅ "What patterns do I notice?"

### Start Small: The 2-Element Case

**Question**: You have two numbers in Stack A. How do you sort them?

**Don't code yet!** Think through these questions:
1. How many possibilities are there? (List them)
2. For each possibility, what operations do you need?
3. What's the maximum operations needed?

**Write your answer in plain English** before touching Go.

Example thinking process:
```
If Stack A has [2, 1] (2 on top, 1 below):
- Problem: 2 > 1, but 1 should be on top
- Solution idea: I need to swap them
- Operation: "sa"
- Result: [1, 2] ✓

If Stack A has [1, 2]:
- Already sorted!
- Operations needed: 0
```

**Your turn**: Write out the logic for 2 elements completely before coding.

---

### Next: The 3-Element Case

**Challenge**: Can you sort 3 numbers with maximum 3 operations?

**Don't code yet!** Answer these:
1. How many ways can 3 numbers be arranged? (List all possibilities)
2. For each arrangement, what's the path to sorted?
3. What patterns do you see?

**Example thinking** (you complete the rest):
```
Possibility 1: [1, 2, 3] - Already sorted, 0 operations
Possibility 2: [1, 3, 2] - What operations get us to [1, 2, 3]?
  - Idea 1: ?
  - Idea 2: ?
  - Best solution: ?

Continue for all 6 possibilities...
```

**Key Insight**: There are only 6 ways to arrange 3 numbers. You can handle each case!

---

## 🔍 Phase 3: Research & Learning

### What Do You Need to Learn?

**For the Checker Program**, you need to understand:
1. How to read from stdin (research: "Go read from standard input")
2. How to store a list of numbers (research: "Go slices")
3. How to check if a list is sorted (think: how would YOU check?)

**For the Push-Swap Program**, you need to understand:
1. Sorting algorithms (research: "types of sorting algorithms")
2. Why normal sorting won't work here (think: what's different?)
3. How to optimize (research: "algorithm complexity")

### Research Tasks (Do These Before Coding!)

**Task 1**: Read about these sorting algorithms:
- Bubble Sort
- Insertion Sort  
- Quick Sort

**Question**: Why can't you just use these directly? What's different about your constraints?

**Task 2**: Research "stack data structure"
- What is LIFO?
- What operations do stacks support?
- How do stacks work in memory?

**Task 3**: Learn about algorithm complexity
- What does O(n) mean?
- What does O(n²) mean?
- Why does it matter?

**Write a summary** (in your own words) of what you learned. Don't copy definitions.

---

## 🛠️ Phase 4: Building Blocks (NOW You Can Code)

### Building Block 1: Stack Implementation

**Questions to answer BEFORE coding**:
1. How will you represent a stack in Go? (What data type?)
2. What functions does a stack need?
3. How do you handle "pop from empty stack"?

**What you need to implement** (figure out HOW):
- A way to store numbers
- A way to add a number (push)
- A way to remove a number (pop)
- A way to see the top number (peek)
- A way to check if empty

**Don't look for code examples yet!** Try implementing based on:
- Go slices documentation
- Your understanding of what a stack does

**Test Strategy**: 
Write tests BEFORE implementation. What should happen if:
- You push 3 numbers, then pop them?
- You try to pop from empty stack?
- You peek at empty stack?

---

### Building Block 2: The 11 Operations

**For EACH operation**, you need to think through:

**Example thinking process for "sa" (swap first two of stack a)**:

Questions:
1. What if stack has 0 elements? (What should happen?)
2. What if stack has 1 element? (Can you swap?)
3. What if stack has 2+ elements? (How do you swap top two?)

Algorithm in plain English:
```
To swap first two elements of a stack:
1. Check if stack has at least 2 elements
2. If not, do nothing (or return error?)
3. If yes:
   - Remove the first element, store it
   - Remove the second element, store it  
   - Put the second element back (it's now first)
   - Put the first element back (it's now second)
```

**Your turn**: Write the algorithm in English for:
- pa (push from b to a)
- ra (rotate a)
- rra (reverse rotate a)

**Only after writing algorithms in English**, then translate to Go.

---

### Building Block 3: Checker Program

**What does the checker need to do?** (Answer in English first)
1. ?
2. ?
3. ?

**Break it into steps**:
```
Step 1: Parse input arguments
  - What could go wrong? (non-numbers? duplicates?)
  - How to check for errors?
  
Step 2: Create stacks from input
  - Which stack do numbers start in?
  - What order?

Step 3: Read instructions from stdin
  - How do you read line by line?
  - What if instruction is invalid?

Step 4: Execute each instruction
  - How do you match string to operation?
  - What if something goes wrong?

Step 5: Check if sorted
  - What does "sorted" mean?
  - What must be true about stack A?
  - What must be true about stack B?
```

**Implement one step at a time. Test each step before moving to next.**

---

## 🎯 Phase 5: Solving Strategy

### The Wrong Way to Approach Sorting

❌ "I'll just try random operations until it's sorted"
❌ "I'll implement bubble sort with these operations"
❌ "I'll guess and hope it works"

### The Right Way to Think

✅ Start with what you CAN solve (2, 3 elements)
✅ Build up to harder problems (4, 5 elements)
✅ Look for patterns
✅ Optimize later

---

### Strategy for 3 Elements: Decision Tree Approach

**Concept**: There are only 6 arrangements. Handle each one.

**Your task**: Create a decision tree on paper.

```
Start with 3 numbers in stack A.
Look at the arrangement.

Is it [1,2,3]? → Done! (0 operations)
Is it [1,3,2]? → What operations? (You figure out)
Is it [2,1,3]? → What operations? (You figure out)
Is it [2,3,1]? → What operations? (You figure out)
Is it [3,1,2]? → What operations? (You figure out)
Is it [3,2,1]? → What operations? (You figure out)
```

**Exercise**: Work out ALL 6 cases on paper before coding.

**Optimization question**: What's the maximum operations needed for 3 elements? (Prove it!)

---

### Strategy for 4-5 Elements: Simplification Approach

**Key Insight**: Can you turn a 5-element problem into a 3-element problem?

**Guiding questions**:
1. What if you push some numbers to stack B?
2. Could you sort the remaining numbers in stack A?
3. Then bring numbers back from stack B?

**Think through this scenario**:
```
Stack A: [5, 2, 4, 1, 3]
Goal: Sorted in ascending order

Idea: What if you push the smallest numbers to B?
- Find smallest (1)
- Move it to top of A (how?)
- Push to B
- Repeat for next smallest (2)
Now A has [5, 4, 3] and B has [1, 2]
- Sort A (3 elements - you can do this!)
- Push back from B

Does this work? Try it on paper!
```

**Your task**: 
1. Test this strategy on paper with different 5-element arrangements
2. Count operations
3. Can you do better than 12 operations?

---

### Strategy for 100 Elements: Research Needed

**You cannot figure this out by yourself easily.** This is where research comes in.

**Research Topics**:
1. "Radix sort algorithm"
2. "Bucket sort"
3. "Divide and conquer sorting"
4. "Chunk-based sorting"

**Key concept to understand**: "Normalization" or "Ranking"
- Instead of sorting actual values, sort their ranks
- Example: [42, 7, 105, 3] → ranks: [2, 1, 3, 0]

**Questions to research and answer**:
1. What does it mean to divide numbers into "chunks"?
2. How could chunks help with sorting?
3. What's the benefit of using stack B strategically?

**After research, develop YOUR algorithm on paper**:
```
Your Algorithm Pseudocode (write this yourself):
1. ?
2. ?
3. ?
```

**Test your algorithm on paper with 10 numbers before coding!**

---

## 🧪 Phase 6: Testing Strategy

### Test Small Before Big

**Test Progression**:
```
1. Test with 2 elements
   - Try: [1,2], [2,1]
   - Verify: operations correct, count is minimal

2. Test with 3 elements
   - Try: all 6 arrangements
   - Verify: all work, max 3 operations

3. Test with 5 elements
   - Try: 10 random arrangements
   - Verify: all work, max 12 operations

4. Test with 100 elements
   - Try: several random sets
   - Verify: all work, less than 700 operations
```

### How to Generate Test Cases

**Don't test randomly!** Think about:
- Best case (already sorted)
- Worst case (reverse sorted)
- Random cases
- Edge cases (duplicates should error!)

**Test command pattern**:
```bash
# Create test input
ARG="3 2 1"; ./push-swap "$ARG"

# Verify with checker
ARG="3 2 1"; ./push-swap "$ARG" | ./checker "$ARG"

# Should output: OK
```

---

## 🔍 Phase 7: Debugging Your Logic

### When It Doesn't Work

**Don't immediately look for solutions!** Debug systematically:

**Step 1: Isolate the Problem**
- Does it fail on all inputs or specific ones?
- What's the smallest input that fails?
- Can you reproduce the bug consistently?

**Step 2: Trace Execution**
- Add print statements to see stack states
- After each operation, print both stacks
- Compare expected vs actual

**Step 3: Check Your Assumptions**
- Did you understand the operation correctly?
- Are you modifying the right stack?
- Are you checking conditions properly?

### Common Logic Errors (Don't Copy, Learn From!)

**Pattern 1: Off-by-one errors**
Question: Are you counting from 0 or 1?

**Pattern 2: Empty stack operations**
Question: Did you check if stack is empty before popping?

**Pattern 3: Wrong stack modified**
Question: Are you operating on A when you meant B?

**Pattern 4: Incorrect rotation logic**
Question: Did you understand what "rotate" means?

---

## 💭 Phase 8: Optimization Thinking

### First Make It Work, Then Make It Fast

**Step 1: Get a working solution** (even if slow)
**Step 2: Measure operations** (how many for 100 elements?)
**Step 3: Find bottlenecks** (where are extra operations?)
**Step 4: Optimize** (reduce unnecessary operations)

### Questions to Ask Yourself

**For operation count**:
1. Am I doing unnecessary rotations?
2. Am I moving numbers back and forth?
3. Could I combine operations?
4. Am I using stack B effectively?

**For algorithm choice**:
1. Why did I choose this approach?
2. Are there better algorithms for constrained sorting?
3. What's the theoretical minimum operations?
4. How close am I to the minimum?

### Optimization Ideas to Research

**Don't implement these immediately!** First understand the concept:

1. **Operation compression**: 
   - Can "sa" followed by "sb" become "ss"?
   - Think: how to detect these patterns?

2. **Reverse operations canceling**:
   - Does "ra" followed by "rra" do nothing?
   - Think: how to detect and remove?

3. **Better algorithms**:
   - Research: "Turk algorithm for push-swap"
   - Research: "Greedy algorithms for stack sorting"
   - Understand concept, then implement YOUR version

---

## 📝 Learning Checklist

### Before You Start Coding

- [ ] I can explain what each operation does in my own words
- [ ] I've worked through 2-3 element cases on paper
- [ ] I understand what LIFO means
- [ ] I know what "sorted" means for this project
- [ ] I've researched sorting algorithms

### After Basic Implementation

- [ ] My checker correctly validates operations
- [ ] I can sort 3 elements with ≤3 operations
- [ ] I can sort 5 elements with ≤12 operations
- [ ] My code handles errors (duplicates, invalid input)
- [ ] I've tested edge cases

### For Large Numbers

- [ ] I've researched chunk-based approaches
- [ ] I understand the concept of normalization
- [ ] I've tested my algorithm on paper first
- [ ] I can sort 100 elements with <700 operations
- [ ] I've verified with checker multiple times

### Code Quality

- [ ] My code is readable (clear variable names)
- [ ] I have comments explaining WHY, not what
- [ ] I've organized code into logical functions
- [ ] Each function does ONE thing
- [ ] I've tested each function independently

---

## 🎓 What You Should Learn (Not Copy!)

### Go Concepts to Research and Understand

**Don't just find code examples!** Read documentation and understand:

1. **Slices**
   - Read: Go blog on slices
   - Understand: how they grow, how to manipulate
   - Practice: create, append, remove elements

2. **Reading stdin**
   - Read: bufio.Scanner documentation
   - Understand: how to read line by line
   - Practice: simple program that echoes stdin

3. **String to Int conversion**
   - Read: strconv package
   - Understand: what errors can occur
   - Practice: convert with error handling

4. **Command-line arguments**
   - Read: os.Args documentation
   - Understand: what os.Args[0] is
   - Practice: program that prints all arguments

### Algorithm Concepts to Research

**Read about these, then think how they apply**:

1. **Sorting Algorithm Comparison**
   - Read about bubble, insertion, quick, merge sort
   - Understand time complexity of each
   - Question: Why can't you use these directly?

2. **Stack-Based Algorithms**
   - Research: stack-based sorting
   - Understand: when stacks are useful
   - Question: How is your problem different?

3. **Optimization Techniques**
   - Research: greedy algorithms
   - Research: divide and conquer
   - Question: Which applies to your problem?

---

## 🤔 Thought Exercises (No Coding!)

### Exercise 1: Manual Sorting
Take these numbers: [5, 2, 8, 1, 9]

Using only the 11 operations, write out step-by-step how to sort them.
Track the state after each operation.

**Goal**: Feel the problem before coding it.

### Exercise 2: Operation Analysis
For each operation, answer:
- When is this operation useful?
- When should you NOT use it?
- Can it be combined with others?

### Exercise 3: Strategy Comparison
Compare two strategies on paper:
- Strategy A: Push all to B, then back sorted
- Strategy B: Use B for chunks

Which is better? Why? Prove it with an example.

### Exercise 4: Edge Case Discovery
List every way the program could fail:
- Invalid inputs
- Empty inputs
- Already sorted
- Reverse sorted
- Duplicates
- Non-numbers
- Very large numbers

How should each be handled?

---

## 🎯 Your Implementation Plan

### Week 1: Foundation
- [ ] Understand all operations deeply (paper exercises)
- [ ] Implement stack data structure
- [ ] Implement all 11 operations
- [ ] Test each operation independently

### Week 2: Checker & Small Cases
- [ ] Build checker program
- [ ] Test checker with manual inputs
- [ ] Solve 2-element case
- [ ] Solve 3-element case
- [ ] Solve 5-element case

### Week 3: Large Cases
- [ ] Research chunk-based sorting
- [ ] Design your algorithm on paper
- [ ] Implement algorithm
- [ ] Test with 100 elements
- [ ] Optimize operation count

### Week 4: Polish & Perfect
- [ ] Test edge cases
- [ ] Optimize further
- [ ] Clean up code
- [ ] Add comprehensive tests
- [ ] Final verification

---

## 💡 Hints (Not Solutions!)

### Hint 1: Stack Representation
Think: What Go data type behaves like a stack?
Answer: A slice can work! Top is the END of the slice (why?)

### Hint 2: Operation Efficiency
Think: If you need smallest number on top, which is faster:
- Rotating until it reaches top
- Reverse rotating

Answer: Depends on where the number is! (Think about why)

### Hint 3: Using Stack B
Think: Stack B is temporary storage. When is it useful?
Answer: When you want to "set aside" some numbers while working on others

### Hint 4: Optimization
Think: Can you predict where each number should go?
Answer: Yes! If you know final positions, you can be strategic about moves

---

## 🚫 What NOT To Do

### DON'T:
❌ Search for "push-swap solution in Go"
❌ Copy code from GitHub
❌ Use AI to write the code for you
❌ Skip the paper exercises
❌ Start with 100 elements
❌ Give up when stuck

### DO:
✅ Work through problems on paper first
✅ Understand each operation deeply
✅ Start with small cases (2, 3 elements)
✅ Test incrementally
✅ Ask yourself "why" at each step
✅ Learn from mistakes

---

## 📚 Recommended Research Path

### Phase 1: Basics (Before Coding)
1. Read Go slice documentation
2. Read about stack data structure
3. Understand LIFO concept
4. Learn about sorting algorithms

### Phase 2: Implementation (While Coding Basics)
1. Research Go stdin reading (bufio)
2. Research string parsing in Go
3. Learn about Go error handling
4. Read about testing in Go

### Phase 3: Advanced (For Large Numbers)
1. Research "radix sort"
2. Research "bucket sort"
3. Look for "stack sorting algorithms"
4. Read about algorithm optimization

**Important**: Don't just copy algorithms! Understand the CONCEPTS, then apply to your constraints.

---

## 🎓 Final Wisdom

### This Project is About Learning to Think

**You're not here to finish quickly.**
**You're here to become a programmer who can solve problems.**

### The Real Skills You're Building

1. **Problem Decomposition**: Breaking big problems into small ones
2. **Algorithmic Thinking**: Step-by-step logical reasoning
3. **Constraint Handling**: Solving within limitations
4. **Optimization**: Making solutions better
5. **Testing**: Verifying solutions work
6. **Debugging**: Finding and fixing errors

**These skills matter more than this specific project.**

### When You're Stuck

1. Go back to paper - trace through manually
2. Simplify - test with 2 numbers instead of 100
3. Research the concept - not the code
4. Take a break - fresh mind sees solutions
5. Explain problem to someone - talking helps thinking

### Success Metrics

**Bad metric**: "I finished quickly"
**Good metric**: "I understand how and why it works"

**Bad metric**: "My code works"
**Good metric**: "I can explain my algorithm to others"

**Bad metric**: "I have a solution"
**Good metric**: "I learned to think algorithmically"

---

## 🤝 How to Use This Guide

1. **Read section by section** - Don't skip ahead
2. **Do the exercises** - They're not optional
3. **Answer questions yourself** - Before looking up anything
4. **Write things down** - On paper, not just thinking
5. **Test your understanding** - Can you explain to someone else?

**Remember**: The guide teaches you HOW to think about the problem. YOU must do the actual thinking and coding.

---

**You're not looking for answers. You're building the skill to create answers.**

**Good luck, programmer!** 🚀