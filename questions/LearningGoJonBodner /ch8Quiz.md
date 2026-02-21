# 📘 Learning Go — Chapter 8 Quiz
## Errors

**Questions:** 22 | **Time:** 30 minutes | **Passing Score:** 18/22 (82%)

---

### Q1: What is the `error` interface in Go?

**A)** A built-in struct with a `Message` field  
**B)** `interface { Error() string }` — any type with an `Error() string` method is an error  
**C)** A keyword like in Python  
**D)** A special return type that triggers exceptions  

<details><summary>💡 Answer</summary>

**B) `interface { Error() string }`**

The entire error system is built on this single-method interface. `errors.New`, `fmt.Errorf`, and every custom error type all implement it. Because it's just an interface, you can make any type an error by adding `Error() string`.

</details>

---

### Q2: What is a sentinel error?

**A)** An error that terminates the program  
**B)** A package-level error variable used for comparison — signals a specific known condition  
**C)** An error in the standard library  
**D)** An error that wraps another error  

<details><summary>💡 Answer</summary>

**B) A package-level error variable for identity comparison**

```go
var ErrNotFound = errors.New("not found")

if err == ErrNotFound { ... }  // old-style comparison
if errors.Is(err, ErrNotFound) { ... }  // correct — handles wrapping
```

Examples: `io.EOF`, `sql.ErrNoRows`, `os.ErrNotExist`. The name starts with `Err` by convention. The book says: use sentinel errors when callers need to distinguish specific error conditions.

</details>

---

### Q3: What is the difference between `errors.New("msg")` and `fmt.Errorf("msg")`?

**A)** They are identical  
**B)** `errors.New` creates a simple error; `fmt.Errorf` creates a formatted error and supports wrapping with `%w`  
**C)** `fmt.Errorf` is deprecated  
**D)** `errors.New` can wrap errors; `fmt.Errorf` cannot  

<details><summary>💡 Answer</summary>

**B) `errors.New` = simple; `fmt.Errorf` = formatted + wrapping with `%w`**

```go
e1 := errors.New("file not found")  // simple, fixed string
e2 := fmt.Errorf("reading %s: %w", filename, e1)  // formatted, wraps e1
```

`%w` is the wrapping verb — it stores the wrapped error so `errors.Is` and `errors.As` can unwrap it.

</details>

---

### Q4: What does error wrapping mean and why is it important?

**A)** Storing the original error inside a new error — allows callers to inspect the full error chain  
**B)** Compressing the error message  
**C)** Converting one error type to another  
**D)** Retrying the operation that caused the error  

<details><summary>💡 Answer</summary>

**A) Storing the original error inside a new — preserves context chain**

```go
// Wrapping adds context without losing the original:
if err := os.Open(path); err != nil {
    return fmt.Errorf("openConfig %s: %w", path, err)
}
```

Callers can later use `errors.Is(err, os.ErrNotExist)` and it will unwrap the chain to find the original error. This preserves the full context of what went wrong and where.

</details>

---

### Q5: What is `errors.Is(err, target)` and why is it better than `err == target`?

**A)** No difference for non-wrapped errors  
**B)** `errors.Is` unwraps the error chain — it returns `true` if any error in the chain matches `target`, not just the outermost error  
**C)** `errors.Is` is slower  
**D)** `errors.Is` only works with sentinel errors  

<details><summary>💡 Answer</summary>

**B) `errors.Is` traverses the unwrap chain**

```go
baseErr := ErrNotFound
wrapped := fmt.Errorf("get user: %w", baseErr)
moreWrapped := fmt.Errorf("handler: %w", wrapped)

moreWrapped == ErrNotFound         // false — only compares outermost
errors.Is(moreWrapped, ErrNotFound) // true — traverses chain
```

Always use `errors.Is` to check for sentinel errors. `==` breaks as soon as the error is wrapped even once.

</details>

---

### Q6: What is `errors.As(err, &target)` used for?

**A)** Type assertion for errors — extracts a specific error type from the chain  
**B)** Converting one error type to another  
**C)** Checking if two errors have the same message  
**D)** Unwrapping the outermost error  

<details><summary>💡 Answer</summary>

**A) Type assertion that traverses the chain**

```go
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string { ... }

var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Println("validation failed on field:", ve.Field)
}
```

`errors.As` unwraps the chain looking for an error that can be assigned to `target`. It sets `target` to the found error and returns `true`. Always use `errors.As` instead of direct type assertions on wrapped errors.

</details>

---

### Q7: What is the difference between `errors.Is` and `errors.As`?

**A)** No difference  
**B)** `errors.Is` checks for a specific error VALUE (sentinel); `errors.As` checks for a specific error TYPE and extracts it  
**C)** `errors.As` is faster  
**D)** `errors.Is` works with types; `errors.As` works with values  

<details><summary>💡 Answer</summary>

**B) `Is` = value/sentinel; `As` = type extraction**

```go
errors.Is(err, ErrNotFound)    // is this specific error in the chain?
errors.As(err, &myCustomErr)   // is there a *MyCustomError in the chain? (extract it)
```

Use `Is` when you care about a specific error instance. Use `As` when you care about the type and need to access its fields.

</details>

---

### Q8: What is `panic` and when should you use it?

**A)** A way to return errors — use it instead of returning error values  
**B)** A function that immediately stops the current goroutine and unwinds the stack — should only be used for unrecoverable programming errors, not expected failures  
**C)** A performance optimization  
**D)** The Go equivalent of try-catch  

<details><summary>💡 Answer</summary>

**B) For unrecoverable errors only — not for expected failures**

The book is very clear: panics are for programming errors (nil dereference, out-of-bounds, impossible states). For expected failures (file not found, network timeout, invalid input), always return errors. If external callers should handle it — return an error. Only `panic` for things that indicate broken invariants.

</details>

---

### Q9: What is `recover()` and where must it be used?

**A)** A built-in that prevents any panics  
**B)** A built-in that captures a panic's value — must be called inside a deferred function to work  
**C)** A method on the `error` interface  
**D)** A package in the standard library  

<details><summary>💡 Answer</summary>

**B) Captures panic value — only works inside a deferred function**

```go
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    return a / b, nil  // panics if b == 0
}
```

If called outside a deferred function during normal execution, `recover()` returns `nil` and does nothing. After `recover()`, execution continues after the deferred function — the panicking goroutine does NOT resume from the panic point.

</details>

---

### Q10: The book says to treat panics from external libraries how?

**A)** Let them propagate — not your problem  
**B)** Wrap them in a deferred recover at your API boundary and convert them to errors — your callers should receive errors, not panics  
**C)** Log and exit  
**D)** Re-panic with a more descriptive message  

<details><summary>💡 Answer</summary>

**B) Recover at API boundaries, convert to errors**

```go
func CallExternalLib() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("external library panicked: %v", r)
        }
    }()
    return externalLib.DoThing()
}
```

Your library's callers shouldn't have to deal with panics from dependencies you chose. Convert to errors at your boundary.

</details>

---

### Q11: What is a custom error type and when should you create one?

**A)** Any struct implementing `Error() string` — create one when callers need to inspect specific fields (not just the error message)  
**B)** A type that extends the built-in error  
**C)** An error defined in a separate file  
**D)** Required for all non-trivial packages  

<details><summary>💡 Answer</summary>

**A) A struct implementing `Error() string` — use when fields matter**

```go
type HTTPError struct {
    Code    int
    Message string
}
func (e *HTTPError) Error() string {
    return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Caller can access Code:
var httpErr *HTTPError
if errors.As(err, &httpErr) {
    if httpErr.Code == 404 { handleNotFound() }
}
```

If callers only need the message, `errors.New` or `fmt.Errorf` is sufficient. Create a custom type when callers need to branch on error fields.

</details>

---

### Q12: What is the `Unwrap() error` method?

**A)** A method on all error values  
**B)** An optional method on error types that returns the wrapped error — used by `errors.Is` and `errors.As` to traverse the chain  
**C)** A way to get the error message  
**D)** Required for all custom errors  

<details><summary>💡 Answer</summary>

**B) Optional method enabling chain traversal**

```go
type WrappedError struct {
    msg  string
    err  error
}
func (e *WrappedError) Error() string { return e.msg + ": " + e.err.Error() }
func (e *WrappedError) Unwrap() error { return e.err }  // enables errors.Is/As
```

`fmt.Errorf("...: %w", err)` automatically implements `Unwrap()`. Your custom wrapping types should too.

</details>

---

### Q13: Should you always wrap errors? What does the book say about when NOT to wrap?

**A)** Always wrap — more context is always better  
**B)** Don't wrap when you're generating a new error from scratch (not propagating an existing one), or when wrapping would expose implementation details  
**C)** Never wrap — it adds overhead  
**D)** Only wrap in the `main` package  

<details><summary>💡 Answer</summary>

**B) Don't wrap when creating new errors or when it leaks implementation details**

If your function calls an internal helper and the helper returns `sql.ErrNoRows`, you might wrap with context: `fmt.Errorf("getUser(%d): %w", id, err)`. But exposing `sql.ErrNoRows` to your callers means they now depend on your SQL implementation. Sometimes `errors.New("user not found")` is better encapsulation.

</details>

---

### Q14: What is the idiomatic way to handle multiple potential errors in sequence?

**A)** Collect all errors in a slice and return them  
**B)** Return immediately on the first error — don't continue if one step fails  
**C)** Use `panic` for subsequent errors  
**D)** Use `select` to handle them concurrently  

<details><summary>💡 Answer</summary>

**B) Return immediately on first error**

```go
func setup() error {
    if err := initDB(); err != nil {
        return fmt.Errorf("setup: initDB: %w", err)
    }
    if err := initCache(); err != nil {
        return fmt.Errorf("setup: initCache: %w", err)
    }
    return nil
}
```

This is the Go error handling pattern. Each operation checks its result before proceeding. Callers know: if there's no error, all steps succeeded.

</details>

---

### Q15: What does the book say about the "errors as values" philosophy?

**A)** Errors should be constants  
**B)** Errors are just values — they can be stored, passed around, ignored (intentionally), and composed like any other value. No special syntax or exception machinery.  
**C)** Errors should always be logged  
**D)** Error values should never escape their package  

<details><summary>💡 Answer</summary>

**B) Errors are regular values — no special mechanics**

This is Go's fundamental philosophy difference from exception-based languages. There's no `try/catch`, no stack unwinding (except panic). Errors flow through your program as return values. You handle them explicitly at each call site. This makes error handling visible and traceable.

</details>

---

### Q16: What is the output?
```go
err := fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", io.EOF))
fmt.Println(errors.Is(err, io.EOF))
```

**A)** `false` — `io.EOF` is buried too deep  
**B)** `true` — `errors.Is` traverses the entire chain  
**C)** Panic  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) `true`**

`errors.Is` unwraps repeatedly until it finds a match or exhausts the chain. Two levels of wrapping doesn't matter — `io.EOF` is at the deepest level and `errors.Is` finds it.

</details>

---

### Q17: What is the `%v` vs `%+v` vs `%#v` format verb for errors?

**A)** All identical for errors  
**B)** `%v` = calls `Error()` string; `%+v` and `%#v` may include extra detail if the type implements the `fmt.Formatter` interface  
**C)** `%#v` always panics with errors  
**D)** Only `%v` works with errors  

<details><summary>💡 Answer</summary>

**B) `%v` calls `Error()`; richer formats depend on the type**

For standard errors: `%v` and `%s` both call `Error()`. Custom error types can implement `fmt.Formatter` for richer output with `%+v`. This is used by libraries like `pkg/errors` (which also provides stack traces with `%+v`).

</details>

---

### Q18: The book says to add context to errors when propagating them up the call stack. What does good error context look like?

**A)** Only the original error message — don't add anything  
**B)** Each wrapping adds the operation name: `"reading config: opening file: permission denied"` — the chain reads like a stack trace  
**C)** The full stack trace in the error message  
**D)** The error code only  

<details><summary>💡 Answer</summary>

**B) Each level adds its operation name — chain tells the story**

```go
// Deep in the stack:
return fmt.Errorf("os.Open: %w", err)

// Middle level:
return fmt.Errorf("readConfig: %w", err)

// Top level:
return fmt.Errorf("startup: %w", err)

// Result: "startup: readConfig: os.Open: permission denied"
```

Reading this error message tells you exactly what failed and the path through the code. The book calls this "error context chains."

</details>

---

### Q19: Can you define multiple `init` functions and can they cause errors?

**A)** No — only one `init` per package  
**B)** Yes — multiple `init` functions are allowed per package and per file; they cannot return errors (signature is `func init()`) — use `panic` for initialization failures  
**C)** `init` functions can return `error`  
**D)** `init` functions run in reverse file order  

<details><summary>💡 Answer</summary>

**B) Multiple allowed; no error return; use panic for failures**

```go
func init() {
    if err := setup(); err != nil {
        panic(fmt.Sprintf("init failed: %v", err))  // only option — can't return error
    }
}
```

The book advises keeping `init` functions minimal and avoiding complex logic there precisely because they can't return errors cleanly.

</details>

---

### Q20: What is the book's advice on using `log.Fatal` vs returning errors?

**A)** Always use `log.Fatal` for simplicity  
**B)** Return errors from libraries and most functions; use `log.Fatal` only in `main` or `init` when the program genuinely cannot proceed  
**C)** `log.Fatal` is deprecated  
**D)** Use `log.Fatal` for all errors above a certain severity  

<details><summary>💡 Answer</summary>

**B) Libraries return errors; `main` can use `log.Fatal`**

`log.Fatal` calls `os.Exit(1)` — no deferred functions run. Libraries should never call `log.Fatal` — it takes control away from the caller. `main` can use it when a startup failure means there's nothing to do but exit. The book emphasizes: libraries return errors; callers decide what to do with them.

</details>

---

### Q21: What is the `errors.Join` function (added in Go 1.20)?

**A)** Concatenates error messages  
**B)** Combines multiple errors into one — useful when you want to report several errors at once; `errors.Is` and `errors.As` work on joined errors  
**C)** Merges two error chains  
**D)** It doesn't exist  

<details><summary>💡 Answer</summary>

**B) Combines multiple errors into one value**

```go
err1 := errors.New("database connection failed")
err2 := errors.New("cache unavailable")
combined := errors.Join(err1, err2)
fmt.Println(combined)
// database connection failed
// cache unavailable
```

`errors.Is(combined, err1)` returns `true`. Useful for validation where you want to report all failures at once.

</details>

---

### Q22: Why does the book say NOT to use `panic` for error handling in libraries?

**A)** Performance reasons  
**B)** It removes the caller's ability to handle the error gracefully — the caller can't use `if err != nil`; the only option is `recover` in a defer, which is awkward  
**C)** Panics are too slow  
**D)** Library panics are caught automatically  

<details><summary>💡 Answer</summary>

**B) Panics bypass the caller's normal error handling**

A library function that panics forces every caller to use `defer recover()` defensively. This is a terrible API design. The contract for a library function is: succeed or return an error. Panics are for the library's OWN programming errors (asserts, impossible states), not for expected runtime failures that callers should handle.

</details>

---

## 📊 Score

| Score | Result |
|---|---|
| 21–22 ✅ | **Excellent.** Solid error handling foundation. |
| 18–20 ✅ | **Ready.** Review `errors.Is` vs `errors.As` and wrapping. |
| 14–17 ⚠️ | **Study error wrapping and the `errors` package — used in every real program.** |
| Below 14 ❌ | **Reread Chapter 8 — proper error handling is non-negotiable in Go.** |