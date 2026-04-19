# 🏟️ Go Gladiators — Session 1 Exercise Sheet

> **Rules:**
> - First to paste a **working solution** in the group chat wins the round points
> - Code must compile and run — no crashes
> - **Bonus +1 point:** write a unit test for any of your solutions
> - Final score = sum of round points + bonus points

---

## Leaderboard

| Name | R1 | R2 | R3 | R4 | Bonus | Total |
|------|----|----|----|----|----|-------|
|      |    |    |    |    |    |       |
|      |    |    |    |    |    |       |
|      |    |    |    |    |    |       |

---

## Round 1 — Username Formatter
**⏱ 10 minutes · 🏆 1 point**

Write a function with this signature:

```go
func FormatUsername(name, role string) string
```

**Rules:**

| Role      | Format                                      | Example input → output              |
|-----------|---------------------------------------------|--------------------------------------|
| `"admin"` | ALL CAPS + `_ADMIN` suffix                  | `"john doe"` → `"JOHN DOE_ADMIN"`   |
| `"guest"` | all lowercase + `guest_` prefix             | `"John"` → `"guest_john"`           |
| `"mod"`   | Capitalize each word + `.mod` suffix        | `"john doe"` → `"John Doe.mod"`     |
| anything else | return the string `"unknown_role"`     | `"john", "vip"` → `"unknown_role"`  |

**Test it with:**
```
"alice smith", "admin"  →  "ALICE SMITH_ADMIN"
"BOB",         "guest"  →  "guest_bob"
"carol white", "mod"    →  "Carol White.mod"
"dave",        "vip"    →  "unknown_role"
```

---

## Round 2 — Vault Decoder
**⏱ 15 minutes · 🏆 2 points**

You receive a slice of strings. Each item is either:
- **Binary** → contains only `0` and `1` characters
- **Hex** → contains digits `0-9` and letters `A-F` (uppercase) — and has at least one letter
- **Invalid** → anything else

Write a program that:
1. Converts each valid entry to its decimal value
2. Prints the **sum** of all valid values
3. Prints how many entries were **skipped** (invalid)
4. Never crashes on bad input

**Input:**
```go
entries := []string{"10011", "1A3F", "00101010", "ZZZ", "FF", "hello", "101"}
```

**Expected output:**
```
Sum: 6997
Skipped: 2
```

> 💡 Hint: `strconv.ParseInt(s, base, 64)` — figure out the right base for each entry.

---

## Round 3 — Log Auditor
**⏱ 20 minutes · 🏆 3 points**

Write a program that takes a `.log` filename as a command-line argument.

Each line in the file looks like:
```
[INFO] Server started
[ERROR] Connection refused
[WARNING] Memory usage high
[INFO] User logged in
```

Your program must:
1. Count how many lines have each level: `INFO`, `ERROR`, `WARNING`
2. Identify the **most common** level
3. Write a **new file** called `report.txt` with exactly this format:

```
INFO: 10
WARNING: 3
ERROR: 5
Most common: INFO
```

4. If the input file does **not exist**, print this message and exit cleanly — no panic:
```
file not found: <filename>
```

**Test it with this file (`server.log`):**
```
[INFO] App started
[ERROR] Null pointer
[WARNING] Disk space low
[INFO] Request received
[INFO] Response sent
[ERROR] Timeout
```

**Expected `report.txt`:**
```
INFO: 3
WARNING: 1
ERROR: 2
Most common: INFO
```

---

## Round 4 — Broken Dialogue Fixer ☠️
**⏱ 15 minutes · 🏆 4 points**

A chat export was badly formatted. Fix it with these rules:

| Rule | Description |
|------|-------------|
| **1. Speaker line** | Any line starting with a name followed by `>` → reformat as `Name:` + one space + dialogue. Name should be capitalized. |
| **2. Ellipsis** | `...` should always attach directly to the previous word, no space before it |
| **3. Emphasis** | Any `*word*` → remove the `*` and make the word fully UPPERCASE |
| **4. Extra spaces** | Multiple spaces anywhere → single space |

**Input:**
```
john  >   are you *sure* about this ...
MARY>  i think so   ...   maybe
old man  >  *never* trust a *calm* sea ...
```

**Expected output:**
```
John: are you SURE about this...
Mary: i think so... maybe
Old Man: NEVER trust a CALM sea...
```

> ⚠️ Watch out for names with multiple words like `"old man"`.

---

## ⭐ Bonus Point

Write a **unit test** using Go's `testing` package for **any one** of your solutions.

```go
func TestFormatUsername(t *testing.T) {
    result := FormatUsername("alice", "admin")
    expected := "ALICE_ADMIN"
    if result != expected {
        t.Errorf("got %s, want %s", result, expected)
    }
}
```

Run with: `go test ./...`

---

*Good luck — first working solution wins! 🚀*