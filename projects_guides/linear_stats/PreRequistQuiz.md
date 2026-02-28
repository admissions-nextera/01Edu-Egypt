# 🔥 Go + Math Quiz — Linear Regression & Pearson Correlation

---

## BLOCK 1 — The Line y = mx + b

### Problem 1: Reading the Formula ⭐
Given the line:
```
y = 2x + 3
```
**Question:** What is the slope? What is the intercept? What is y when x = 5?

**Key Concept:** `m` controls steepness (rise per step). `b` is where the line crosses the y-axis (when x = 0)!

---

### Problem 2: Slope Sign Meaning ⭐
```
Line A: y =  3x + 1
Line B: y = -3x + 1
Line C: y =  0x + 5
```
**Question:** What does each slope tell you about the relationship between x and y?

**Key Concept:** Sign of slope reveals direction of the relationship — positive, negative, or none!

---

### Problem 3: Intercept Meaning ⭐
```
y = 4x + 7
```
**Question:** Without any calculation, what is y when x = 0? Why?

**Key Concept:** The intercept `b` is always the value of y when x = 0 — no math needed!

---

### Problem 4: Predict with a Line ⭐⭐
You fit a line to some data and get:
```
m = 1.5
b = 2.0
```
**Question:** Predict y for x = 0, x = 4, and x = 10.

**Key Concept:** Once you have m and b, prediction is just plug-and-calculate!

---

### Problem 5: Two Points → Slope ⭐⭐
```
Point 1: (2, 5)
Point 2: (6, 13)
```
**Question:** Calculate the slope between these two points using:
```
m = (y2 - y1) / (x2 - x1)
```

**Key Concept:** Slope = rise over run — change in y divided by change in x!

---

## BLOCK 2 — Index as X-Axis

### Problem 6: What is the X? ⭐
You have this dataset read from a file:
```
10.5
14.2
18.7
22.1
```
**Question:** In linear regression on this data, what are the x values? What are the y values?

**Key Concept:** When data has no explicit x, the index (0, 1, 2, …) becomes the x-axis!

---

### Problem 7: Index vs Value Confusion ⭐⭐
```go
values := []float64{5.0, 10.0, 15.0, 20.0}

// Version A
for i, v := range values {
    x := v   // ← using the value as x
    y := float64(i)
    fmt.Println(x, y)
}

// Version B
for i, v := range values {
    x := float64(i)  // ← using the index as x
    y := v
    fmt.Println(x, y)
}
```
**Question:** Which version is correct for linear regression on file data? What does each print?

**Key Concept:** In file-based regression, `x = float64(i)` and `y = values[i]` — never swap them!

---

### Problem 8: Mean of Indices ⭐⭐
```go
values := []float64{3.1, 6.2, 9.3, 12.4, 15.5}
n := len(values)
```
**Question:** What is the mean of the x values (indices 0 to n-1)? Write the formula and calculate it.

**Key Concept:** Mean of indices 0…n-1 is always `(n-1) / 2.0` — a useful shortcut!

---

### Problem 9: Why Index Not Value? ⭐⭐⭐
You have temperatures recorded each day:
```
Day 0: 20°C
Day 1: 22°C
Day 2: 19°C
Day 3: 25°C
```
**Question:** Why is it wrong to use the temperature as x? What would break in the regression?

**Key Concept:** x = what you control or observe in order (time/position). y = what changes as a result!

---

## BLOCK 3 — Linear Regression Formula

### Problem 10: The Slope Formula ⭐
The slope formula for linear regression is:
```
m = Σ((xi - meanX)(yi - meanY)) / Σ((xi - meanX)²)
```
**Question:** What does the numerator measure? What does the denominator measure?

**Key Concept:** Slope = covariance of x,y divided by variance of x!

---

### Problem 11: Calculate Slope by Hand ⭐⭐
```
Data: (0, 2), (1, 4), (2, 6)
meanX = 1.0
meanY = 4.0
```
**Question:** Calculate the slope `m` step by step.

**Key Concept:** Work through the formula term by term — numerator sums products, denominator sums squares!

---

### Problem 12: The Intercept Formula ⭐
Once you have slope m, the intercept is:
```
b = meanY - m * meanX
```
**Question:** Using `m = 2.0`, `meanX = 1.0`, `meanY = 4.0`, calculate b.

**Key Concept:** `b = meanY - m * meanX` — the line always passes through the point (meanX, meanY)!

---

### Problem 13: Slope Formula in Go ⭐⭐
```go
func slope(xs, ys []float64, meanX, meanY float64) float64 {
    num, den := 0.0, 0.0
    for i := range xs {
        num += (xs[i] - meanX) * (ys[i] - meanY)
        den += (xs[i] - meanX) * (xs[i] - meanX)
    }
    return num / den
}
```
**Question:** What does this function return if all x values are identical? Is this a bug or expected?

**Key Concept:** Zero denominator in slope = undefined slope. Always guard against it in real code!

---

### Problem 14: Intercept in Go ⭐⭐
```go
func intercept(m, meanX, meanY float64) float64 {
    return meanY - m*meanX
}

func main() {
    fmt.Println(intercept(2.0, 1.0, 4.0))
    fmt.Println(intercept(0.0, 5.0, 3.0))
    fmt.Println(intercept(-1.5, 2.0, 1.0))
}
```
**Question:** What does each line print?

**Key Concept:** When slope is 0, `b = meanY` — the flat line sits exactly at the mean of y!

---

### Problem 15: Full Regression from Scratch ⭐⭐⭐
```
Data values (y): 3, 7, 11, 15
Indices (x):     0, 1,  2,  3
```
**Question:** Calculate meanX, meanY, slope m, and intercept b. Then predict y at x = 5.

**Key Concept:** Perfect linear data produces an exact fit — every point lies on the line!

---

## BLOCK 4 — Pearson Correlation Coefficient

### Problem 16: What Does r Mean? ⭐
**Question:** Match each r value to its meaning:

| r value | Meaning |
|---------|---------|
| 1.0     | ? |
| -1.0    | ? |
| 0.0     | ? |
| 0.85    | ? |
| -0.3    | ? |


**Key Concept:** r is always between -1 and 1. Closer to ±1 = stronger linear relationship!

---

### Problem 17: The Pearson Formula ⭐⭐
```
r = Σ((xi - meanX)(yi - meanY))
    ─────────────────────────────────────────
    sqrt(Σ(xi - meanX)²) * sqrt(Σ(yi - meanY)²)
```
**Question:** What is the relationship between the numerator of r and the numerator of the slope formula?

**Key Concept:** Pearson r normalizes the slope by the spread of y — making it always land in [-1, 1]!

---

### Problem 18: Calculate r by Hand ⭐⭐
```
Data: (0, 2), (1, 4), (2, 6)
meanX = 1.0,  meanY = 4.0
```
**Question:** Calculate the Pearson r step by step.

**Key Concept:** Perfectly linear data always gives r = exactly 1.0 or -1.0!

---

### Problem 19: r in Go ⭐⭐⭐
```go
func pearson(xs, ys []float64, meanX, meanY float64) float64 {
    num, denX, denY := 0.0, 0.0, 0.0
    for i := range xs {
        dx := xs[i] - meanX
        dy := ys[i] - meanY
        num  += dx * dy
        denX += dx * dx
        denY += dy * dy
    }
    return num / math.Sqrt(denX*denY)
}
```
**Question:** What happens if all y values are identical? What does the function return?

**Key Concept:** r requires both variables to vary — zero variance in x or y makes r undefined!

---

### Problem 20: Negative Correlation ⭐⭐
```
Data: (0, 10), (1, 7), (2, 4), (3, 1)
```
**Question:** Without calculating, predict the sign of r and the sign of slope m. Why?

**Key Concept:** Slope and r always share the same sign — both negative for downward trends!

---

### Problem 21: r vs Slope — Key Difference ⭐⭐⭐
```
Dataset A: (0,0), (1,10), (2,20)   → m = 10,  r = 1.0
Dataset B: (0,0), (1,1),  (2,2)    → m = 1,   r = 1.0
```
**Question:** Both datasets have r = 1.0 but very different slopes. What does this tell you?

**Key Concept:** r = 1.0 means perfect linearity, NOT steep slope. r and slope measure different things!

---

## BLOCK 5 — Float Formatting & Precision

### Problem 22: fmt.Sprintf Formatting ⭐
```go
x := 3.14159265358979
fmt.Printf("%.2f\n", x)
fmt.Printf("%.6f\n", x)
fmt.Printf("%.10f\n", x)
```
**Question:** What does each line print?

**Key Concept:** `%.Nf` formats a float to exactly N decimal places, rounding the last digit!

---

### Problem 23: Sprintf vs Printf ⭐
```go
r := 0.987654321
s := fmt.Sprintf("%.6f", r)
fmt.Println(s)
fmt.Printf("%T\n", s)
```
**Question:** What does each line print?

**Key Concept:** `Sprintf` = returns a formatted string. `Printf` = prints directly. Same format verbs!

---

### Problem 24: Precision Trap ⭐⭐
```go
x := 1.0 / 3.0
fmt.Printf("%.6f\n", x)
fmt.Printf("%.10f\n", x)
fmt.Println(x == 0.333333)
```
**Question:** What does each line print?

**Key Concept:** Formatting rounds for display — the underlying float is still the full precision value!

---

### Problem 25: Formatting Negative Floats ⭐⭐
```go
values := []float64{-1.5678901234, 0.0, -0.0000001}
for _, v := range values {
    fmt.Printf("%.6f\n", v)
}
```
**Question:** What does each line print?

**Key Concept:** `%.6f` always prints exactly 6 decimal places — even for zero and very small values!

---

### Problem 26: Format in a Function ⭐⭐
```go
func formatResult(label string, value float64) string {
    return fmt.Sprintf("%s: %.6f", label, value)
}

func main() {
    fmt.Println(formatResult("slope", 2.0))
    fmt.Println(formatResult("r", 1.0))
}
```
**Question:** What does each line print?

**Key Concept:** `%.6f` always pads with zeros to reach exactly 6 places — `2.0` becomes `2.000000`!

---

## BLOCK 6 — Translating Formulas to Code

### Problem 27: Mean in Go ⭐
```go
func mean(data []float64) float64 {
    sum := 0.0
    for _, v := range data {
        sum += v
    }
    return sum / float64(len(data))
}
```
**Question:** What happens if `data` is empty? How would you fix it?

**Key Concept:** Always guard against empty slices before dividing by `len` — it produces NaN silently!

---

### Problem 28: Build X Slice from Indices ⭐⭐
```go
ys := []float64{10.5, 14.2, 18.7}
xs := make([]float64, len(ys))
for i := range ys {
    xs[i] = float64(i)
}
fmt.Println(xs)
```
**Question:** What does this print? Why is `float64(i)` needed?

**Key Concept:** Always convert index `i` to `float64(i)` before using it in float calculations!

---

### Problem 29: Numerator and Denominator Loop ⭐⭐
```go
func slope(xs, ys []float64) float64 {
    mx := mean(xs)
    my := mean(ys)
    num, den := 0.0, 0.0
    for i := range xs {
        num += (xs[i] - mx) * (ys[i] - my)
        den += (xs[i] - mx) * (xs[i] - mx)
    }
    return num / den
}
```
**Question:** Can you rewrite `(xs[i] - mx) * (xs[i] - mx)` more concisely? Does it change the result?

**Key Concept:** Extract repeated subexpressions into variables — cleaner and avoids double computation!

---

### Problem 30: Pearson vs Slope Code Difference ⭐⭐⭐
```go
// Slope
func slope(xs, ys []float64) float64 {
    // ...
    return num / den
}

// Pearson
func pearson(xs, ys []float64) float64 {
    // ...
    return num / math.Sqrt(denX * denY)
}
```
**Question:** Both functions have the same numerator loop. What is the ONLY difference in their denominators?

**Key Concept:** Pearson needs TWO denominator accumulators — one for x spread, one for y spread!

---

### Problem 31: Putting It Together ⭐⭐⭐
```go
data := []float64{2.0, 4.0, 6.0, 8.0}
xs   := []float64{0.0, 1.0, 2.0, 3.0}
```
**Question:** Without running code — what will slope m, intercept b, and Pearson r be? Explain why without calculation.

**Key Concept:** Arithmetic sequences (constant step in y) always produce r = 1.0 and slope = that step size!

---

### Problem 32: The Full Pipeline ⭐⭐⭐
```go
data := []float64{5.0, 3.0, 7.0, 1.0}

// Step 1: build xs
// Step 2: compute means
// Step 3: compute slope
// Step 4: compute intercept
// Step 5: compute pearson
// Step 6: print with 6 decimal places
```
**Question:** Fill in what each step produces for this data. Then write what the final output lines look like.

**Key Concept:** Negative r and negative slope always go together — this data is a weak downward trend!

---

## 🏆 Quick Reference Card

| Concept | Formula / Rule |
|---------|----------------|
| Line equation | `y = mx + b` |
| Slope meaning | Rise per unit of x |
| Intercept meaning | y when x = 0 |
| X values in file data | Index: 0, 1, 2, … |
| Y values in file data | The actual data values |
| Mean of indices 0…n-1 | `(n-1) / 2.0` |
| Slope formula | `Σ(dx·dy) / Σ(dx²)` where `dx = xi - meanX` |
| Intercept formula | `b = meanY - m * meanX` |
| Pearson formula | `Σ(dx·dy) / sqrt(Σ(dx²) · Σ(dy²))` |
| r range | Always between -1 and +1 |
| r = 1.0 | Perfect positive linear fit |
| r = 0.0 | No linear relationship |
| Slope vs r | Same sign, different meaning: steepness vs strength |
| Zero denominator | Undefined slope/r — always guard! |
| `%.6f` | Exactly 6 decimal places, rounds last digit |
| `%.10f` | Exactly 10 decimal places |
| `Sprintf` vs `Printf` | Returns string vs prints directly |
| `float64(i)` | Always convert index to float64 before math |

**Master these and statistics becomes just arithmetic! 💪🔥**
