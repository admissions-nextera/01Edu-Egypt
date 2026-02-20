# Guess It 1 Project Guide

> **Before you start:** Complete math-skills first. This project uses average, variance, and standard deviation as building blocks. If those calculations are not working correctly, this project will not score well.

---

## Objectives

By completing this project you will learn:

1. **Standard Input Streaming** — Reading numbers one at a time as they arrive, not all at once
2. **Rolling Statistics** — Updating statistical calculations incrementally as new data comes in
3. **Confidence Intervals** — Using standard deviation to define a prediction range
4. **Scoring Optimization** — Balancing range width against accuracy
5. **Scripting** — Packaging your program to run from a shell script

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Math Skills (Completed)
- Average, variance, and standard deviation — working correctly

### 2. Standard Input in Your Language
- How to read from stdin line by line, blocking until input arrives
- How to print to stdout and flush immediately (important — the tester reads your output line by line)
- Search: **"read stdin line by line [your language]"**
- Search: **"flush stdout [your language]"**

### 3. Statistics Concepts
- What a standard deviation tells you about a dataset
- What a confidence interval is and how standard deviation is used to build one
- Search: **"standard deviation confidence interval explained"**
- Search: **"68-95-99.7 rule normal distribution"**

### 4. Shell Scripting
- How to write an executable shell script
- What `chmod +x` does
- Search: **"write executable shell script tutorial"**

---

## Project Structure

```
student/
├── your-program.{go,js,py,rs}
└── script.sh
```

The `student/` folder sits inside the tester's root directory when testing.

---

## Milestone 1 — Understand the Input/Output Contract

**This milestone has no code.**

Read the spec interaction carefully:

```
189       ← you receive this from stdin
120 200   ← you print this to stdout (your range prediction for the NEXT number)
113       ← you receive the next number
160 230   ← you print your prediction for the one after that
...
```

**Questions to answer before writing anything:**
- When you receive the first number, how many data points do you have? Can you make a meaningful prediction yet?
- After receiving 2 numbers, can you compute a standard deviation?
- What should you print for the very first number when you have no history to base a prediction on?
- The output is two numbers separated by a space: `lower upper`. What happens if lower > upper?

---

## Milestone 2 — Read Numbers from Standard Input

**Goal:** Your program reads numbers from stdin one at a time and prints `"received: N"` for each — no prediction yet.

**Questions to answer:**
- How does reading from stdin differ from reading from a file?
- How do you know when stdin is finished (EOF)?
- Do you need to print anything before reading the next line?

**Code Placeholder:**
```
main:
    // Initialize an empty list to store numbers seen so far

    // Loop:
    //   Read one line from stdin
    //   If EOF or read error: break
    //   Convert line to number
    //   Add to the list
    //   Print "received: " + number  (temporary, replace later)
```

**Verify:** Run your program. Type numbers manually. Confirm each is received and echoed. Press Ctrl+D (EOF) to end.

---

## Milestone 3 — Build the Prediction Strategy

**This milestone has no code.**

Your goal is to predict a range `[lower, upper]` such that the next number falls inside it as often as possible, while keeping the range as narrow as possible (wider ranges score less).

**Questions to answer:**
- What does the standard deviation of your seen data tell you about where the next value might land?
- The 68-95-99.7 rule says: in a normal distribution, ~68% of values fall within 1 standard deviation of the mean, ~95% within 2. Which multiplier gives you the best balance for this dataset?
- Your prediction center: should it be the mean? The last value? Something else?
- What should you output when you have fewer than 2 data points and cannot yet compute a standard deviation?

Draw your strategy on paper: `center ± (multiplier × stddev)`.

---

## Milestone 4 — Implement the Prediction

**Goal:** After receiving each number, print the predicted range for the next one.

**Questions to answer:**
- Do you recalculate mean and standard deviation from scratch each time, or update them incrementally?
- What is the minimum number of data points needed before your formula produces a meaningful range?
- How do you handle the case where standard deviation is 0 (all values so far are identical)?

**Code Placeholder:**
```
main:
    // Initialize an empty list

    // Loop:
    //   Read one line from stdin
    //   If EOF: break
    //   Convert to number, add to list

    //   If not enough data for a prediction:
    //     Print a wide default range (e.g. "0 1000")
    //     Continue

    //   Calculate mean of the list
    //   Calculate standard deviation of the list

    //   Calculate: lower = mean - (multiplier * stddev)
    //   Calculate: upper = mean + (multiplier * stddev)

    //   Round lower and upper to integers
    //   Print lower + " " + upper
    //   Flush stdout immediately
```

**Critical:** Flush stdout after every print. The tester reads your output line by line — if you buffer, the tester will hang waiting for output that is stuck in your buffer.

**Resources:**
- Search: **"flush stdout immediately [your language]"**

**Verify:** Run your program manually. Feed numbers one at a time. Confirm a range prints after each input.

---

## Milestone 5 — Write the Shell Script

**Goal:** Create `student/script.sh` that runs your program.

**Questions to answer:**
- What command runs your program from the tester's root directory?
- Your program is inside `student/` — what path do you use?
- Is the script executable? (`chmod +x script.sh`)

**Code Placeholder:**
```sh
#!/bin/sh
# Run from the tester's root directory
# Your program is in ./student/

# Example for Python:
# python3 ./student/solution.py

# Example for Go (pre-compiled binary):
# ./student/solution

# Example for JavaScript:
# node ./student/solution.js
```

**Verify:**
```bash
chmod +x student/script.sh
./student/script.sh
# Type numbers manually — confirm the program runs and produces ranges
```

---

## Milestone 6 — Tune Your Multiplier

**Goal:** Find the multiplier that maximizes your score.

The scoring rule: if your range correctly contains the next number, you score points inversely proportional to the range size. A range of 50 scores more than a range of 200.

**Questions to answer:**
- What happens to your accuracy as you decrease the multiplier?
- What happens to your score per correct guess as you increase the multiplier?
- Try multipliers of 1.0, 1.5, 2.0, 2.5. Which gives the best balance?
- Does using the mean as the center give the best accuracy, or is there a better center?

**Verify:** Download the tester zip from the spec. Run it with `Data 1`, `Data 2`, and `Data 3`. Compare scores at different multipliers.

---

## Debugging Checklist

- Does the tester hang without receiving output? You are buffering stdout. Flush after every print.
- Is your range sometimes `lower > upper`? Standard deviation cannot be negative, so this should not happen — check your subtraction order.
- Is your score always 0? Your range may never contain the next number. Try a much wider range first to confirm the pipeline works, then narrow it.
- Does the program crash after a few inputs? Check that you handle the case where only 1 data point exists (standard deviation of a single value is 0 or undefined).
- Is the script not running? Check `chmod +x` and confirm the path in the script is correct relative to the tester's root.

---

## Key Concepts

| Concept | How It Is Used Here |
|---|---|
| Mean | Center of your prediction range |
| Standard Deviation | Half-width of your prediction range |
| Multiplier | Controls how many standard deviations wide the range is |
| Confidence Interval | `[mean - k*σ, mean + k*σ]` where k is your multiplier |

---

## Submission Checklist

- [ ] Program reads numbers from stdin one at a time
- [ ] Prints a range after each number received
- [ ] Range format is exactly `lower upper` (two integers, space-separated)
- [ ] Stdout is flushed after every print
- [ ] Handles fewer than 2 data points without crashing
- [ ] Handles standard deviation of 0 without crashing
- [ ] `student/` folder contains all necessary files
- [ ] `script.sh` is executable and runs the program correctly from tester root
- [ ] Tested with the provided tester on Data 1, Data 2, and Data 3