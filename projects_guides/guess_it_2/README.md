# Guess It 2 Project Guide

> **Before you start:** Complete linear-stats first. This project uses the linear regression line as its core prediction tool. If your slope and intercept calculations are not correct, this project will not score well. Also re-read guess-it-1 — the input/output contract is identical.

---

## Objectives

By completing this project you will learn:

1. **Predictive Modeling** — Using a fitted line to predict future values instead of just describing past ones
2. **Residual Analysis** — Measuring how far actual values deviate from the predicted line
3. **Dynamic Regression** — Updating the regression line as new data points arrive
4. **Prediction Intervals** — Building a range around a linear prediction using residual spread
5. **Scoring Optimization** — Tightening your range while maintaining accuracy

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Linear Stats (Completed)
- Linear regression slope and intercept — working correctly
- Pearson correlation coefficient — working correctly

### 2. Guess It 1 (Completed)
- Stdin streaming loop with immediate stdout flushing
- `student/` folder and `script.sh` structure

### 3. Residuals
- What a residual is: the difference between an actual value and the predicted value
- Search: **"linear regression residuals explained"**
- Search: **"prediction interval linear regression"**

### 4. Why Linear Regression Beats Simple Mean Here
- In guess-it-1, you used mean ± stddev — a horizontal band
- In guess-it-2, the data may have a trend (going up or down over time)
- A linear regression captures that trend — your prediction center is now the line, not the mean

---

## Project Structure

```
student/
├── your-program.{go,js,py,rs}
└── script.sh
```

---

## Milestone 1 — Understand the Upgrade from Guess It 1

**This milestone has no code.**

In guess-it-1 your strategy was:
```
center = mean of all seen values
range  = center ± (k × standard_deviation)
```

In guess-it-2 your strategy will be:
```
center = predicted y at the next x, using the regression line
range  = center ± (k × spread_of_residuals)
```

**Questions to answer before writing anything:**
- If your data is `189, 113, 121, 114, 145, 110` and you have seen 5 values (indices 0–4), what x value do you use to predict the next one?
- What is a residual for data point `i`? (It is `actual_y[i] - predicted_y[i]`.)
- What does the spread of residuals tell you that the standard deviation of raw values does not?
- What should you print when you have fewer data points than needed to fit a regression line?

---

## Milestone 2 — Reuse and Adapt Linear Regression

**Goal:** At any point in the stream, compute the linear regression line for all values seen so far, then predict the next value.

**Questions to answer:**
- Your linear-stats program computes regression over a fixed file. How do you adapt this to run on a growing list?
- After seeing N values (indices 0 to N-1), what is the x value for the next prediction? (It is N.)
- How many data points do you need before a regression line is meaningful? (At least 2.)

**Code Placeholder:**
```
function predictNext(xs, ys):
    // Compute linear regression slope and intercept from xs and ys
    // next_x = length of xs (the index of the next point)
    // predicted_y = slope * next_x + intercept
    // Return predicted_y
```

**Verify:** For the first 3 points `(0,189), (1,113), (2,121)`:
- Compute slope and intercept by hand
- Predict y at x=3
- Confirm your function returns the same value

---

## Milestone 3 — Compute Residual Spread

**Goal:** Measure how much the actual values deviate from the regression line. Use this as the half-width of your range.

**Questions to answer:**
- For each point `(xi, yi)`, what is the residual? (`yi - (slope*xi + intercept)`)
- What statistic of the residuals gives a good measure of their spread? (Standard deviation of residuals.)
- How does this differ from the standard deviation of raw y values?

**Code Placeholder:**
```
function residualStdDev(xs, ys):
    // Compute slope and intercept for xs and ys
    // For each point i:
    //   predicted = slope * xs[i] + intercept
    //   residual = ys[i] - predicted
    // Compute the standard deviation of all residuals
    // Return it
```

**Verify:** For a dataset where all points lie exactly on a line, residual standard deviation should be 0.

---

## Milestone 4 — Build the Prediction Range

**Goal:** Combine the predicted next value with the residual spread to produce a range.

**Code Placeholder:**
```
main:
    // Initialize empty xs and ys lists

    // Loop:
    //   Read one line from stdin
    //   If EOF: break
    //   Convert to number
    //   current_x = length of xs (before appending)
    //   Append current_x to xs
    //   Append value to ys

    //   If fewer than 2 data points:
    //     Print a wide default range
    //     Flush stdout
    //     Continue

    //   predicted = predictNext(xs, ys)
    //   spread = residualStdDev(xs, ys)

    //   lower = predicted - (multiplier * spread)
    //   upper = predicted + (multiplier * spread)

    //   Round lower and upper to integers
    //   Print lower + " " + upper
    //   Flush stdout immediately
```

**Critical:** Flush stdout after every print — same as guess-it-1.

**Verify:** Run manually. Feed 10+ numbers and watch the range center shift as the trend changes.

---

## Milestone 5 — Handle Edge Cases

**Goal:** The program never crashes regardless of input.

**Questions to answer:**
- What happens when residual standard deviation is 0 (all points exactly on the line)? Your range would be `[predicted, predicted]` — is that valid?
- What if the regression line has a very steep slope and your predicted next value is wildly outside the range of seen values?
- What if only one data point has been seen — is a regression meaningful?

**Code Placeholder:**
```
    // After computing spread:
    //   If spread is 0: use a small minimum range (e.g. ±5)
    //   If spread is very large: cap it at a reasonable maximum

    // After computing lower and upper:
    //   If lower >= upper: set upper = lower + 1 (ensure valid range)
```

---

## Milestone 6 — Tune and Compare with Guess It 1

**Goal:** Find the multiplier that gives the best score on Data 4 and Data 5.

**Questions to answer:**
- Does the linear prediction outperform the mean-based prediction from guess-it-1? On which datasets?
- Is the Pearson correlation coefficient useful here — if r is close to 0 (no linear trend), should you fall back to the guess-it-1 strategy?
- What multiplier (k) of residual standard deviation maximizes your score?

**Strategy hint:**
```
if pearson_correlation is strong (|r| > 0.7):
    use linear prediction ± k × residual_stddev
else:
    fall back to mean ± k × standard_deviation  (guess-it-1 strategy)
```

Try this hybrid approach and compare its score to each strategy alone.

---

## Milestone 7 — Update the Shell Script

**Goal:** `student/script.sh` runs your updated program.

**Verify:**
```bash
chmod +x student/script.sh
./student/script.sh
# Type numbers manually — confirm ranges are produced
```

Run the tester with Data 4 and Data 5. Compare your score to guess-it-1's score on those datasets.

---

## Debugging Checklist

- Does the tester hang? You are buffering stdout. Flush after every print.
- Are predictions drifting far off the actual values? Check that your regression uses `x = index`, not the y values as x.
- Is residual spread always 0? You may be recomputing the regression on only the latest point, not all seen points.
- Does the program crash after receiving 1 point? Guard against fewer than 2 data points before calling regression.
- Is your range always negative or impossibly large? Print your predicted value and spread separately to isolate which is wrong.

---

## Key Concepts

| Concept | How It Is Used Here |
|---|---|
| Regression line `y = ax + b` | Center of the prediction range |
| Residual | `actual - predicted` for each past point |
| Residual standard deviation | Half-width of the prediction range |
| Pearson r | Whether linear or mean-based strategy is better |
| Multiplier k | Trade-off between range width and accuracy |

---

## Submission Checklist

- [ ] Program reads numbers from stdin one at a time
- [ ] Prints a range after each number received
- [ ] Range center is based on the linear regression prediction, not the mean
- [ ] Range width is based on residual standard deviation
- [ ] Stdout is flushed after every print
- [ ] Handles fewer than 2 data points without crashing
- [ ] Handles zero residual spread without producing invalid range
- [ ] `student/` folder contains all necessary files
- [ ] `script.sh` is executable and runs correctly from tester root
- [ ] Tested with the provided tester on Data 4 and Data 5
- [ ] Score is higher than a naive wide-range approach