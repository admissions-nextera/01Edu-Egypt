# Linear Stats Project Guide

> **Before you start:** Read the Wikipedia articles for both concepts completely before writing any code. Work through the formulas manually on a small dataset — 5 points is enough. If you cannot compute the answer by hand, you cannot verify your program is correct.

---

## Objectives

By completing this project you will learn:

1. **Linear Regression** — Finding the line that best fits a set of data points
2. **Pearson Correlation Coefficient** — Measuring how strongly two variables are linearly related
3. **Index as X-Axis** — Understanding that the line number is the x value, not the data value
4. **Decimal Precision** — Formatting floating-point output to a specific number of decimal places
5. **Mathematical Derivation** — Translating a statistical formula into working code

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Math Skills (Completed)
- Mean calculation — you will need it here
- Reading a data file

### 2. Linear Regression Concepts
- What a line `y = mx + b` represents
- What "best fit" means (minimizing squared error)
- The formula for slope (m) and intercept (b)
- https://en.wikipedia.org/wiki/Linear_regression
- Search: **"linear regression slope intercept formula explained"**

### 3. Pearson Correlation Coefficient
- What a value of 1, -1, and 0 means
- The formula
- https://en.wikipedia.org/wiki/Pearson_correlation_coefficient
- Search: **"Pearson correlation coefficient formula step by step"**

### 4. Floating Point Formatting
- How to format a float to exactly 6 decimal places
- How to format a float to exactly 10 decimal places
- Search: **"format float decimal places [your language]"**

---

## Project Structure

```
linear-stats/
├── your-program.{go,js,py,rs}
└── data.txt
```

---

## Milestone 1 — Understand the X and Y Values

**This milestone has no code.**

The data file contains only y values. The x values are the line numbers:
```
Line 0: 189   →  point (0, 189)
Line 1: 113   →  point (1, 113)
Line 2: 121   →  point (2, 121)
...
```

**Questions to answer before writing anything:**
- What is the mean of the x values for a dataset of N points? (It is always `(N-1) / 2` — do you see why?)
- What is the mean of the y values?
- In the Pearson formula, what are you measuring the correlation between — two data columns, or x (index) and y (value)?

---

## Milestone 2 — Read the File and Build (x, y) Pairs

**Goal:** Read the data file and produce a list of `(x, y)` pairs where x is the line index and y is the value.

**Code Placeholder:**
```
function loadData(filename):
    // Read all lines from the file
    // For each non-empty line at index i:
    //   Convert line to float
    //   Record the pair (i, value)
    // Return a list of (x, y) pairs or two separate lists: xs and ys
```

**Verify:** Print the first 5 pairs from your data file and confirm they look like `(0, 189), (1, 113), ...`

---

## Milestone 3 — Linear Regression Line

**Goal:** Compute the slope and intercept of the best-fit line.

**Formulas:**
```
mean_x = mean of all x values
mean_y = mean of all y values

slope (a) = Σ[(xi - mean_x)(yi - mean_y)] / Σ[(xi - mean_x)²]

intercept (b) = mean_y - slope × mean_x
```

**Questions to answer:**
- How do you compute the two summations in the slope formula?
- What should you do if the denominator `Σ[(xi - mean_x)²]` is zero? (This means all x values are identical — impossible in this project since x is always the index.)
- The output requires 6 decimal places. How do you format this in your language?

**Code Placeholder:**
```
function linearRegression(xs, ys):
    // Compute mean of xs and mean of ys

    // Compute numerator: sum of (xi - mean_x) * (yi - mean_y) for all i
    // Compute denominator: sum of (xi - mean_x)² for all i

    // slope = numerator / denominator
    // intercept = mean_y - slope * mean_x

    // Return (slope, intercept)
```

**Expected output format:**
```
Linear Regression Line: y = 2.123456x + 45.678901
```

**Verify manually:** For points `(0,2), (1,4), (2,5), (3,4), (4,5)`:
- mean_x = 2, mean_y = 4
- numerator = (0-2)(2-4)+(1-2)(4-4)+(2-2)(5-4)+(3-2)(4-4)+(4-2)(5-4) = 4+0+0+0+2 = 6
- denominator = 4+1+0+1+4 = 10
- slope = 0.6, intercept = 4 - 0.6×2 = 2.8
- Output: `y = 0.600000x + 2.800000`

---

## Milestone 4 — Pearson Correlation Coefficient

**Goal:** Compute how strongly the line index and the value correlate.

**Formula:**
```
r = Σ[(xi - mean_x)(yi - mean_y)] / √[Σ(xi - mean_x)² × Σ(yi - mean_y)²]
```

**Questions to answer:**
- Notice the numerator is the same as in the slope formula — can you reuse it?
- What does the denominator require that the slope denominator does not?
- What should you return if the denominator is zero? (Perfect correlation or no variance — handle gracefully.)
- The output requires 10 decimal places.

**Code Placeholder:**
```
function pearsonCorrelation(xs, ys):
    // Compute mean of xs and mean of ys

    // Compute numerator: sum of (xi - mean_x) * (yi - mean_y)
    // Compute sum_sq_x: sum of (xi - mean_x)²
    // Compute sum_sq_y: sum of (yi - mean_y)²

    // denominator = sqrt(sum_sq_x * sum_sq_y)
    // If denominator is 0: handle this edge case

    // r = numerator / denominator
    // Return r
```

**Expected output format:**
```
Pearson Correlation Coefficient: 0.1234567890
```

**Verify manually:** For points `(0,2), (1,4), (2,5), (3,4), (4,5)`:
- numerator = 6 (same as above)
- sum_sq_x = 10
- sum_sq_y = (2-4)²+(4-4)²+(5-4)²+(4-4)²+(5-4)² = 4+0+1+0+1 = 6
- denominator = √(10 × 6) = √60 ≈ 7.746
- r ≈ 6 / 7.746 ≈ 0.7745966692

---

## Milestone 5 — Connect Everything and Format Output

**Goal:** The program reads the file and prints both results in the exact required format.

**Expected format:**
```
Linear Regression Line: y = <6 decimal places>x + <6 decimal places>
Pearson Correlation Coefficient: <10 decimal places>
```

**Questions to answer:**
- What if the intercept is negative? The format becomes `y = 0.500000x + -1.200000` — is that acceptable, or should it be `y = 0.500000x - 1.200000`?
- How do you format a float to exactly N decimal places in your language?

**Code Placeholder:**
```
main:
    // Read filename from argument
    // Load (xs, ys) from file

    // Compute slope and intercept
    // Compute Pearson coefficient

    // Print: "Linear Regression Line: y = {slope:.6f}x + {intercept:.6f}"
    // Print: "Pearson Correlation Coefficient: {r:.10f}"
```

**Verify:** Run with a known dataset and compute both values by hand (or use a calculator). Compare to your program's output digit by digit.

---

## Debugging Checklist

- Is your slope output wrong but your means are correct? Print the numerator and denominator separately and check each against your manual calculation.
- Is your Pearson coefficient always exactly 1.0? You may be using x values that are perfectly correlated by construction. Use a dataset with noisy y values.
- Is the number of decimal places wrong? Check how your language's format string works for exactly 6 vs 10 places — they are different for the two outputs.
- Is the intercept formatted as `+ -3.5` instead of `- 3.5`? Handle the sign of the intercept explicitly if the spec requires a clean format.
- Are floating point rounding errors affecting your last decimal place? This is expected — the auditor's program will have the same precision limits.

---

## Key Formulas Reference

| Output | Formula |
|---|---|
| Slope (a) | `Σ[(xi-x̄)(yi-ȳ)] / Σ(xi-x̄)²` |
| Intercept (b) | `ȳ - a × x̄` |
| Pearson (r) | `Σ[(xi-x̄)(yi-ȳ)] / √[Σ(xi-x̄)² × Σ(yi-ȳ)²]` |

**Resources:**
- https://en.wikipedia.org/wiki/Linear_regression
- https://en.wikipedia.org/wiki/Pearson_correlation_coefficient
- Search: **"linear regression formula derivation simple"**

---

## Submission Checklist

- [ ] Program reads filename from command-line argument
- [ ] Correctly assigns x = line index, y = line value
- [ ] Linear regression slope computed correctly to 6 decimal places
- [ ] Linear regression intercept computed correctly to 6 decimal places
- [ ] Pearson coefficient computed correctly to 10 decimal places
- [ ] Output format matches spec exactly
- [ ] No crash on any valid data file
- [ ] Results match auditor's program on the same dataset