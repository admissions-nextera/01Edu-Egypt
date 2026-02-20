# Math Skills Project Guide

> **Before you start:** Read the Wikipedia pages for each of the four statistics before writing any code. You cannot implement a formula you do not understand. Work through each one manually on a small dataset (5 numbers) with pen and paper first.

---

## Objectives

By completing this project you will learn:

1. **Average (Mean)** — The sum of all values divided by the count
2. **Median** — The middle value of a sorted dataset
3. **Variance** — How spread out the values are from the mean
4. **Standard Deviation** — The square root of the variance, in the same units as the data
5. **File Reading** — Reading numerical data from a plain text file
6. **Rounding** — Correctly rounding float results to integers

---

## Prerequisites — Topics You Must Know Before Starting

### 1. The Four Statistics
Read each one before writing any code:
- https://en.wikipedia.org/wiki/Average
- https://en.wikipedia.org/wiki/Median
- https://en.wikipedia.org/wiki/Variance
- https://en.wikipedia.org/wiki/Standard_deviation

### 2. Your Language's Standard Library
- How to read a file and split it into lines
- How to convert a string to a number
- How to sort a list/slice/array of numbers
- How to compute a square root
- How to round a float to the nearest integer

### 3. Math Foundations
- What the difference is between population variance and sample variance — the spec uses a **statistical population**, which affects the formula denominator
- Search: **"population variance vs sample variance"**

---

## Project Structure

```
math-skills/
├── your-program.{go,js,py,rs}
└── data.txt
```

---

## Milestone 1 — Read the File

**Goal:** Your program reads a file passed as a command-line argument and prints the numbers to confirm they loaded correctly.

**Questions to answer before writing anything:**
- How do you access command-line arguments in your chosen language?
- How do you read a file line by line?
- How do you convert each line from a string to a number?
- What should happen if the file does not exist or cannot be read?

**Code Placeholder:**
```
function readData(filename):
    // Open the file at the given path
    // Read all lines
    // For each line:
    //   Trim whitespace
    //   Skip empty lines
    //   Convert to number and add to a list
    // Return the list of numbers
```

**Verify:** Run your program with a small test file. Print the loaded numbers and confirm they match the file contents exactly.

---

## Milestone 2 — Average

**Goal:** Print the correct average, rounded to the nearest integer.

**Formula:**
```
average = sum of all values / count of values
```

**Questions to answer:**
- How do you sum a list of numbers in your language?
- How do you round a float to the nearest integer? Is `int(3.7)` the same as rounding? (No — check the difference.)
- What should happen if the dataset is empty?

**Code Placeholder:**
```
function average(data):
    // Sum all values
    // Divide by the count
    // Round to nearest integer
    // Return the result
```

**Verify manually:** For `[189, 113, 121, 114, 145, 110]`:
- Sum = 792
- Count = 6
- Average = 792 / 6 = 132
- Does your program print `132`?

---

## Milestone 3 — Median

**Goal:** Print the correct median, rounded to the nearest integer.

**Formula:**
```
Sort the data.
If count is odd:  median = middle value
If count is even: median = average of the two middle values
```

**Questions to answer:**
- How do you sort a list of numbers in your language?
- For an even-count dataset, which two indices are the "middle two"?
- Does sorting modify the original list? Does that matter for later calculations?

**Code Placeholder:**
```
function median(data):
    // Make a sorted copy of data (do not modify the original)
    // If count is odd:
    //   Return the element at index count/2 (integer division)
    // If count is even:
    //   Return the average of elements at index (count/2 - 1) and (count/2)
    //   Round the result
```

**Verify manually:** For `[189, 113, 121, 114, 145, 110]`:
- Sorted: `[110, 113, 114, 121, 145, 189]`
- Count = 6 (even) → middle two: index 2 and 3 → `114` and `121`
- Median = (114 + 121) / 2 = 117.5 → rounds to `118`
- Does your program print `118`?

---

## Milestone 4 — Variance

**Goal:** Print the correct variance, rounded to the nearest integer.

**Formula (population variance):**
```
variance = sum of (each value - mean)² / count
```

**Questions to answer:**
- Are you computing **population** variance (divide by N) or **sample** variance (divide by N-1)? The spec says "statistical population" — which formula is correct?
- How do you square a number in your language?
- Should you reuse the average from Milestone 2, or recalculate it?

**Code Placeholder:**
```
function variance(data):
    // Calculate the mean of the data
    // For each value:
    //   Compute (value - mean)²
    //   Add to a running sum
    // Divide the sum by the count (N, not N-1 — this is population variance)
    // Round and return
```

**Verify manually:** For `[189, 113, 121, 114, 145, 110]`, mean = 132:
- Deviations squared: `(189-132)²=3249`, `(113-132)²=361`, `(121-132)²=121`, `(114-132)²=324`, `(145-132)²=169`, `(110-132)²=484`
- Sum = 4708
- Variance = 4708 / 6 = 784.67 → rounds to `785`
- Does your program print `785`?

---

## Milestone 5 — Standard Deviation

**Goal:** Print the correct standard deviation, rounded to the nearest integer.

**Formula:**
```
standard deviation = square root of variance
```

**Questions to answer:**
- Which function in your language computes the square root?
- Should you compute variance first and then take the square root, or rewrite the whole formula from scratch? (Reuse your variance function.)

**Code Placeholder:**
```
function standardDeviation(data):
    // Compute variance using your variance function
    // Take the square root of the variance
    // Round and return
```

**Verify manually:** Variance = 785 (from previous step):
- √785 ≈ 28.02 → rounds to `28`
- Does your program print `28`?

---

## Milestone 6 — Connect Everything and Format Output

**Goal:** The program reads the file, computes all four values, and prints them in the exact format:
```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

**Questions to answer:**
- Is the output format exact — correct capitalization, colon, single space?
- Do your rounded results match what the auditor's program will produce?

**Code Placeholder:**
```
main:
    // Read filename from command-line argument
    // If no argument, print usage message and exit

    // Load data from file
    // If file cannot be read, print error and exit

    // Compute and print:
    //   "Average: " + average(data)
    //   "Median: " + median(data)
    //   "Variance: " + variance(data)
    //   "Standard Deviation: " + standardDeviation(data)
```

**Verify:** Create a `data.txt` with a small known dataset. Compute the expected results by hand. Run your program and compare.

---

## Debugging Checklist

- Is your average slightly off? Check whether you are doing integer division instead of float division before rounding.
- Is your median wrong for even-count datasets? Check the indices — for 6 elements, the middle two are at index 2 and 3 (0-based).
- Is your variance using N-1 instead of N? The spec specifies a statistical population — use N.
- Is your standard deviation wrong but variance is right? Check that you are taking the square root of the unrounded variance, not the rounded one.
- Are empty lines in the file causing a parse error? Trim and skip empty lines before converting to numbers.

---

## Key Concepts

| Concept | Formula | Resource |
|---|---|---|
| Average | Σx / N | https://en.wikipedia.org/wiki/Average |
| Median | Middle of sorted data | https://en.wikipedia.org/wiki/Median |
| Population Variance | Σ(x - mean)² / N | https://en.wikipedia.org/wiki/Variance |
| Standard Deviation | √variance | https://en.wikipedia.org/wiki/Standard_deviation |

---

## Submission Checklist

- [ ] Program reads filename from command-line argument
- [ ] Correctly parses all numbers from the file
- [ ] Average is computed and rounded correctly
- [ ] Median is computed correctly for both odd and even count datasets
- [ ] Variance uses population formula (divide by N, not N-1)
- [ ] Standard deviation is the square root of variance
- [ ] All values printed as rounded integers
- [ ] Output format matches spec exactly (capitalization, spacing)
- [ ] Program does not crash on an empty file or missing argument