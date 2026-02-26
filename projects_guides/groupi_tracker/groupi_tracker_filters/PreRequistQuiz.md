# 🎯 Groupie Tracker Filters Prerequisites Quiz
## Goroutines · sync.WaitGroup · sync.Mutex · Race Conditions · Filter Logic

**Time Limit:** 55 minutes  
**Total Questions:** 30  
**Passing Score:** 24/30 (80%)

> Questions are tagged: 🟢 Easy · 🟡 Medium · 🔴 Hard  
> All topics are general — no specific project knowledge required.

---

## 📋 SECTION 1: GOROUTINES AND CONCURRENCY BASICS (7 Questions)

### Q1 🟢 — What does the `go` keyword do?

**A)** Marks a function as asynchronous  
**B)** Launches a new goroutine — a lightweight thread managed by the Go runtime that runs the function call concurrently  
**C)** Compiles the function separately  
**D)** Creates a promise  

<details><summary>💡 Answer</summary>

**B) Launches a goroutine — concurrent execution**

```go
func sayHello(name string) {
    fmt.Println("Hello,", name)
}

go sayHello("Alice")   // starts concurrently — does NOT block
go sayHello("Bob")     // another goroutine
// main continues immediately — Alice and Bob print in undefined order
```

Goroutines are not OS threads — they're multiplexed by the Go runtime onto actual threads, making them very cheap (~2KB of stack). You can have thousands of goroutines where hundreds of OS threads would fail.

</details>

---

### Q2 🟢 — What is a race condition?

**A)** Two goroutines competing to run first — the faster one wins  
**B)** When two or more goroutines access shared data concurrently, at least one access is a write, and there's no synchronization — the result depends on timing and is unpredictable  
**C)** A goroutine that runs too fast  
**D)** Using `go` in a for loop  

<details><summary>💡 Answer</summary>

**B) Concurrent unsynchronized access where at least one access is a write**

```go
// RACE CONDITION:
var count int  // shared
go func() { count++ }()  // write
go func() { count++ }()  // write

// Two goroutines increment count simultaneously.
// count++ is actually: read → add 1 → write.
// Both may read 0, both write 1 — final value: 1, not 2.
// This is a silent data corruption bug.
```

Race conditions are among the hardest bugs to debug because they're timing-dependent — they may work fine 999 times and fail on the 1000th. Use `go run -race .` to detect them.

</details>

---

### Q3 🟡 — How does `sync.WaitGroup` work?

**A)** It pauses all goroutines for a fixed duration  
**B)** A counter: `Add(n)` registers n goroutines, each calls `Done()` when finished, and `Wait()` blocks until the counter reaches zero  
**C)** It limits how many goroutines can run concurrently  
**D)** It queues goroutines and runs them one at a time  

<details><summary>💡 Answer</summary>

**B) A counter for waiting on multiple goroutines**

```go
var wg sync.WaitGroup

wg.Add(3)  // expecting 3 goroutines to finish

go func() {
    defer wg.Done()  // decrements counter when this goroutine returns
    doWork1()
}()
go func() {
    defer wg.Done()
    doWork2()
}()
go func() {
    defer wg.Done()
    doWork3()
}()

wg.Wait()  // blocks until all 3 call Done()
fmt.Println("all goroutines finished")
```

`WaitGroup` is the standard tool for "fire off N goroutines, wait for all of them." `Add` must be called before the goroutines start.

</details>

---

### Q4 🟡 — Why should you use `defer wg.Done()` instead of calling `wg.Done()` at the end of the goroutine body?

**A)** `defer` is faster  
**B)** If the goroutine returns early (due to an error or `return` statement) or panics, `defer wg.Done()` still runs — without `defer`, the counter never reaches zero and `wg.Wait()` blocks forever  
**C)** `defer wg.Done()` can be called multiple times safely  
**D)** No difference — the goroutine always runs to the end  

<details><summary>💡 Answer</summary>

**B) `defer` guarantees `Done()` runs even on early return or panic**

```go
// DANGEROUS — if fetchArtists returns early, Done() never called:
go func() {
    data, err := fetchArtists()
    if err != nil {
        log.Println(err)
        return            // wg.Done() at the bottom never runs → deadlock!
    }
    process(data)
    wg.Done()            // only runs on success
}()

// SAFE — defer runs on every exit path:
go func() {
    defer wg.Done()      // always runs, no matter what
    data, err := fetchArtists()
    if err != nil {
        log.Println(err)
        return
    }
    process(data)
}()
```

Always use `defer wg.Done()` as the first line of a goroutine that participates in a `WaitGroup`.

</details>

---

### Q5 🟡 — What happens to the program when all goroutines are blocked waiting for each other?

**A)** The program runs forever  
**B)** The Go runtime detects the deadlock and panics: "all goroutines are asleep — deadlock!"  
**C)** The goroutines are automatically cancelled  
**D)** Nothing — Go handles this gracefully  

<details><summary>💡 Answer</summary>

**B) The runtime detects and panics with a deadlock message**

```go
// Classic deadlock:
var wg sync.WaitGroup
wg.Add(1)
// goroutine never started — Done() never called
wg.Wait()  // blocks forever → runtime detects → panic

// Also causes deadlock:
wg.Add(2)  // expects 2 Done calls
go func() { defer wg.Done(); }()  // only 1 goroutine — 1 Done call
wg.Wait()  // blocks forever waiting for 2nd Done
```

The runtime can only detect deadlocks where ALL goroutines are blocked. If just some are stuck, it won't detect it — so test your WaitGroup counts carefully.

</details>

---

### Q6 🔴 — What is the output? Is there a bug?

```go
var wg sync.WaitGroup
results := make([]int, 3)

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        results[i] = i * i
    }()
}
wg.Wait()
fmt.Println(results)
```

**A)** `[0 1 4]` — correct  
**B)** `[9 9 9]` or a panic — `i` is captured by reference; by the time goroutines run, `i == 3`, which is out of bounds for `results[i]`  
**C)** `[0 0 0]`  
**D)** Random values — race condition on `results`  

<details><summary>💡 Answer</summary>

**B) Bug — `i` captured by reference, likely out of bounds panic or wrong values**

```go
// The goroutine closure captures i by reference.
// By the time a goroutine runs, the loop may have advanced i to 3.
// results[3] is out of bounds → panic.

// FIX — pass i as a parameter:
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(n int) {   // n is a copy of i at this moment
        defer wg.Done()
        results[n] = n * n
    }(i)               // i is evaluated and passed NOW
}
```

This is the goroutine-in-loop closure bug — one of Go's most famous gotchas. Always pass the loop variable as a parameter to the goroutine function.

</details>

---

### Q7 🔴 — You launch 4 goroutines to fetch data from 4 API endpoints. One fails. How do you collect all results AND the error?

**A)** Check `err` globally after `wg.Wait()`  
**B)** Use a shared struct (protected by mutex) or a results channel to collect both values and errors from each goroutine  
**C)** Goroutines can't return errors  
**D)** Use `panic` inside the goroutine  

<details><summary>💡 Answer</summary>

**B) Shared struct with mutex, or results channel**

```go
type result struct {
    data interface{}
    err  error
}

results := make([]result, 4)
var wg sync.WaitGroup

for i, url := range urls {
    wg.Add(1)
    go func(idx int, u string) {
        defer wg.Done()
        data, err := fetch(u)
        results[idx] = result{data, err}  // safe: each goroutine writes its own index
    }(i, url)
}

wg.Wait()

for i, r := range results {
    if r.err != nil {
        fmt.Println("error on", i, ":", r.err)
    }
}
```

Each goroutine writes to its own index — no mutex needed when goroutines don't share indices. For dynamic fan-out, use a channel of results instead.

</details>

---

## 📋 SECTION 2: sync.Mutex AND SHARED STATE (8 Questions)

### Q8 🟢 — What does `sync.Mutex` do?

**A)** Prevents goroutines from starting  
**B)** Ensures only one goroutine can execute the code between `Lock()` and `Unlock()` at a time — others block until the lock is released  
**C)** Limits the number of goroutines  
**D)** Detects deadlocks  

<details><summary>💡 Answer</summary>

**B) Mutual exclusion — one goroutine at a time in the critical section**

```go
var (
    mu    sync.Mutex
    count int
)

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++  // only one goroutine runs this at a time
}

// 100 concurrent goroutines, each calls increment():
// Final count is guaranteed to be 100 (no race)
```

`Lock()` acquires the mutex — if already locked, the caller blocks. `Unlock()` releases it. The section between Lock and Unlock is the "critical section" — protected from concurrent access.

</details>

---

### Q9 🟢 — Why should you always use `defer mu.Unlock()` instead of calling `mu.Unlock()` at the end?

**A)** `defer` is required by the compiler  
**B)** If the function returns early or panics, `defer mu.Unlock()` ensures the mutex is released — without `defer`, other goroutines waiting on `Lock()` are blocked forever  
**C)** `defer mu.Unlock()` is faster  
**D)** No difference  

<details><summary>💡 Answer</summary>

**B) `defer` guarantees unlock on every exit path**

```go
func updateMap(key string, value int) {
    mu.Lock()
    defer mu.Unlock()  // runs on every return path, including errors

    if !isValid(key) {
        return  // mutex is still released — defer runs
    }
    myMap[key] = value
}
```

Without `defer`, every early `return` must manually call `mu.Unlock()`. Miss one and you have a deadlock. With `defer`, it's impossible to forget.

</details>

---

### Q10 🟡 — What is the difference between `sync.Mutex` and `sync.RWMutex`?

**A)** `RWMutex` is faster for all use cases  
**B)** `sync.Mutex` allows one goroutine at a time (any operation); `sync.RWMutex` allows multiple concurrent readers OR one exclusive writer — ideal for read-heavy shared data  
**C)** `RWMutex` is for reading only  
**D)** They are identical  

<details><summary>💡 Answer</summary>

**B) `Mutex` = one at a time; `RWMutex` = many readers OR one writer**

```go
var (
    mu    sync.RWMutex
    cache map[string]string
)

// Reading — multiple goroutines can hold RLock simultaneously:
func get(key string) string {
    mu.RLock()
    defer mu.RUnlock()
    return cache[key]
}

// Writing — exclusive, blocks all readers and writers:
func set(key, value string) {
    mu.Lock()
    defer mu.Unlock()
    cache[key] = value
}
```

Use `RWMutex` when reads are frequent and writes are rare. For write-heavy code, a plain `Mutex` is simpler and often just as fast.

</details>

---

### Q11 🟡 — What is a deadlock caused by mutex misuse?

**A)** Forgetting to create the mutex  
**B)** A goroutine tries to Lock a mutex it already holds — since `sync.Mutex` is not reentrant, it blocks forever waiting for itself  
**C)** Using the mutex in a goroutine  
**D)** Calling Unlock without Lock  

<details><summary>💡 Answer</summary>

**B) Re-locking a non-reentrant mutex — goroutine blocks forever on itself**

```go
// DEADLOCK:
func broadcast(msg string) {
    mu.Lock()
    defer mu.Unlock()
    // ...
    sendHistory(conn)  // sendHistory also calls mu.Lock() — DEADLOCK!
}

func sendHistory(conn net.Conn) {
    mu.Lock()          // goroutine already holds mu → blocks on itself
    defer mu.Unlock()
    // never reached
}
```

`sync.Mutex` is NOT reentrant (unlike Java's `synchronized`). If you hold the lock and try to lock again, you wait forever. Fix: don't call locking functions from within a lock, or use a separate internal function that assumes the lock is already held.

</details>

---

### Q12 🟡 — What does `go run -race .` do?

**A)** Runs the program faster using race conditions  
**B)** Enables the Go race detector — instruments the binary to detect concurrent unsynchronized accesses at runtime; reports exact file/line of the race  
**C)** Compiles for a race car  
**D)** Disables goroutines  

<details><summary>💡 Answer</summary>

**B) Enables the race detector — finds concurrent access bugs at runtime**

```bash
go run -race .
go test -race ./...
go build -race -o myapp .

# When a race is detected, you get:
# ==================
# WARNING: DATA RACE
# Write at 0x00c000018080 by goroutine 7:
#   main.main.func1()
#       /home/user/main.go:15 +0x44
# Previous read at 0x00c000018080 by goroutine 6:
#   main.main.func2()
#       /home/user/main.go:22 +0x38
```

The race detector adds ~5–10× overhead — use it during development and testing, not in production. Always run `go test -race ./...` in CI.

</details>

---

### Q13 🔴 — Is this code safe from race conditions?

```go
var mu sync.Mutex
var data []string

func addItem(item string) {
    mu.Lock()
    data = append(data, item)
    mu.Unlock()
}

func getItems() []string {
    mu.Lock()
    result := data
    mu.Unlock()
    return result
}
```

**A)** Yes — all accesses are protected  
**B)** Partially — `addItem` is safe, but `getItems` returns the slice header (pointer+len+cap). The caller can read the slice while another goroutine calls `addItem` and triggers a reallocation, making the returned slice point to freed memory  
**C)** No — both functions have race conditions  
**D)** Yes — the mutex covers all operations  

<details><summary>💡 Answer</summary>

**B) Partially safe — returned slice is a snapshot but underlying array may be reallocated**

```go
// SAFER — return a copy:
func getItems() []string {
    mu.Lock()
    defer mu.Unlock()
    result := make([]string, len(data))
    copy(result, data)
    return result  // independent copy — safe to use after lock is released
}
```

Returning a slice from under a mutex is subtle — the caller holds the slice header, but the underlying array is still shared. If `append` in `addItem` causes reallocation, the old array (pointed to by the returned slice) is still valid (Go GC keeps it alive), but it won't see new items. Returning a copy is always safer.

</details>

---

### Q14 🔴 — What is the correct way to protect a map accessed by multiple goroutines?

**A)** Use `sync.Map` always  
**B)** Wrap all map accesses in a `sync.Mutex` or `sync.RWMutex`; or use `sync.Map` for append-only/load-heavy workloads  
**C)** Maps are automatically thread-safe in Go  
**D)** Use channels to serialize access  

<details><summary>💡 Answer</summary>

**B) Wrap in mutex, or use `sync.Map` for specific patterns**

```go
// Option 1: RWMutex (most common):
var (
    mu   sync.RWMutex
    cache = make(map[string]int)
)
func get(k string) (int, bool) {
    mu.RLock(); defer mu.RUnlock()
    v, ok := cache[k]; return v, ok
}
func set(k string, v int) {
    mu.Lock(); defer mu.Unlock()
    cache[k] = v
}

// Option 2: sync.Map (best for many goroutines, mostly reading, rarely writing):
var m sync.Map
m.Store("key", 42)
v, ok := m.Load("key")
```

Go's built-in `map` panics on concurrent write+write or read+write. `sync.Map` has different trade-offs (harder API, no generics). For most cases, `RWMutex` + regular map is simpler and more efficient.

</details>

---

### Q15 🔴 — What is the "thundering herd" problem in concurrent caching?

**A)** Many goroutines crashing at once  
**B)** When a cache is empty, many goroutines detect the miss simultaneously and all rush to fetch the same data — producing N redundant network calls instead of 1  
**C)** A cache that grows without limit  
**D)** Cache eviction causing slowdowns  

<details><summary>💡 Answer</summary>

**B) Many concurrent cache misses triggering redundant work**

```go
// Without protection — thundering herd:
func getWithCache(key string) Data {
    mu.RLock()
    if v, ok := cache[key]; ok {
        mu.RUnlock()
        return v  // cache hit
    }
    mu.RUnlock()
    // 50 goroutines all reach here simultaneously — all call fetch(key)!
    v := fetch(key)  // 50 redundant network calls
    mu.Lock()
    cache[key] = v
    mu.Unlock()
    return v
}

// Fix: use sync.Map's LoadOrStore, or double-check pattern:
mu.Lock()
if v, ok := cache[key]; ok {  // check again under write lock
    mu.Unlock(); return v
}
v := fetch(key)
cache[key] = v
mu.Unlock()
return v
```

The solution is to serialize the first fetch for a given key. `golang.org/x/sync/singleflight` is designed exactly for this.

</details>

---

## 📋 SECTION 3: FILTER LOGIC AND DATA PROCESSING (8 Questions)

### Q16 🟢 — You have a slice of items and want to keep only those matching a condition. What is the idiomatic Go pattern?

**A)** Use `slice.Filter(func)`  
**B)** Loop and append matching items to a new slice  
**C)** Modify the slice in place with a cursor  
**D)** Use `sort.Filter`  

<details><summary>💡 Answer</summary>

**B) Loop and append to a new slice**

```go
type Person struct { Name string; Age int }
people := []Person{{"Alice", 30}, {"Bob", 17}, {"Carol", 25}}

// Keep only adults:
var adults []Person
for _, p := range people {
    if p.Age >= 18 {
        adults = append(adults, p)
    }
}
// adults = [{Alice 30} {Carol 25}]
```

Go has no built-in `filter`. The pattern is explicit and clear. Go 1.23+ adds `slices.Collect(slices.Values(s))` patterns, but the loop approach remains most readable and is used in all existing code.

</details>

---

### Q17 🟢 — When you have multiple filter criteria (e.g. filter by age AND by city), which logical operator connects them?

**A)** OR — an item must match at least one filter  
**B)** AND — an item must match ALL active filters to be included  
**C)** XOR — an item must match exactly one filter  
**D)** NOT — exclude items matching any filter  

<details><summary>💡 Answer</summary>

**B) AND — all active filters must pass**

```go
type Filter struct {
    MinAge  int
    MaxAge  int
    City    string
}

func applyFilters(people []Person, f Filter) []Person {
    var result []Person
    for _, p := range people {
        // ALL conditions must be true:
        if p.Age < f.MinAge { continue }
        if p.Age > f.MaxAge { continue }
        if f.City != "" && p.City != f.City { continue }
        result = append(result, p)
    }
    return result
}
```

Different filter types (age range, city, membership) are combined with AND — a user expects filtering to narrow results, not expand them. Within a single multi-select filter (e.g. multiple cities), items matching ANY selected city are included (OR).

</details>

---

### Q18 🟡 — An HTML form has checkboxes for selecting multiple values. How do you read ALL checked values in Go?

**A)** `r.FormValue("field")` — returns a comma-separated string  
**B)** `r.URL.Query()["field"]` or `r.Form["field"]` — returns a `[]string` of all checked values  
**C)** `r.MultiFormValue("field")`  
**D)** Loop over `r.FormValues`  

<details><summary>💡 Answer</summary>

**B) `r.Form["field"]` or `r.URL.Query()["field"]` returns `[]string`**

```go
// HTML form with multiple checkboxes of the same name:
// <input type="checkbox" name="genre" value="rock">
// <input type="checkbox" name="genre" value="jazz">
// <input type="checkbox" name="genre" value="pop">

// If user checks rock and jazz:
r.ParseForm()
genres := r.Form["genre"]    // ["rock", "jazz"]
// r.FormValue("genre")      // "rock" — only returns the FIRST value!

// Check if a value was selected:
func contains(slice []string, s string) bool {
    for _, v := range slice { if v == s { return true } }
    return false
}
```

`r.FormValue` only returns the first value. For multi-select (checkboxes or multi-select dropdowns), always use `r.Form["field"]` which returns all values.

</details>

---

### Q19 🟡 — You have a range slider for "year from 1950 to 2023". How do you parse and validate the values from a form?

**A)** `r.FormValue("year_min")` returns an `int` directly  
**B)** `strconv.Atoi(r.FormValue("year_min"))` — form values are always strings, must convert and validate  
**C)** `r.IntValue("year_min")`  
**D)** `json.Unmarshal(r.Body, &yearMin)`  

<details><summary>💡 Answer</summary>

**B) `strconv.Atoi` — form values are always strings**

```go
yearMinStr := r.FormValue("year_min")
yearMin, err := strconv.Atoi(yearMinStr)
if err != nil || yearMin < 1950 || yearMin > 2023 {
    http.Error(w, "invalid year", http.StatusBadRequest)
    return
}

yearMaxStr := r.FormValue("year_max")
yearMax, err := strconv.Atoi(yearMaxStr)
if err != nil || yearMax < yearMin || yearMax > 2023 {
    http.Error(w, "invalid year range", http.StatusBadRequest)
    return
}
```

HTTP form data is always text. Every number must be parsed and validated. Never trust user input — validate ranges, not just types.

</details>

---

### Q20 🟡 — What does AND-between-filters, OR-within-filter mean for a search with genres ["rock", "jazz"] and years [2000-2010]?

**A)** Include items that are (rock OR jazz) AND (year 2000–2010)  
**B)** Include items that are rock AND jazz AND between 2000 and 2010  
**C)** Include items that are rock OR jazz OR between 2000 and 2010  
**D)** Exclude items that are rock or jazz  

<details><summary>💡 Answer</summary>

**A) (rock OR jazz) AND (year between 2000 and 2010)**

```go
func matches(item Item, genres []string, yearMin, yearMax int) bool {
    // Within genres filter: OR — match any selected genre
    genreMatch := len(genres) == 0  // no filter = all pass
    for _, g := range genres {
        if item.Genre == g {
            genreMatch = true
            break
        }
    }
    if !genreMatch { return false }

    // Between filters: AND — must also pass year filter
    if item.Year < yearMin || item.Year > yearMax {
        return false
    }

    return true
}
```

This is the standard UX behavior: multiple checkboxes in one group = OR (any is acceptable), multiple filter groups = AND (all must pass).

</details>

---

### Q21 🟡 — How do you find unique values from a slice (e.g. all unique locations)?

**A)** `slice.Unique()`  
**B)** Use a `map[string]bool` or `map[string]struct{}` as a "seen" set while iterating  
**C)** Sort the slice and check adjacent elements  
**D)** Use `sync.Map`  

<details><summary>💡 Answer</summary>

**B) Use a map as a seen-set**

```go
func unique(items []string) []string {
    seen := make(map[string]bool)
    var result []string
    for _, item := range items {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }
    return result
}

locations := []string{"Paris", "London", "Paris", "Berlin", "London"}
fmt.Println(unique(locations))  // ["Paris", "London", "Berlin"]
```

The map-as-set pattern is O(n) — the most efficient approach for deduplication. Option C (sort + check adjacent) works but requires O(n log n) and changes order.

</details>

---

### Q22 🔴 — You run 4 goroutines that each append to the same slice. What's wrong and how do you fix it?

```go
var results []string
var wg sync.WaitGroup

for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        data := fetch(u)
        results = append(results, data)  // BUG
    }(url)
}
wg.Wait()
```

**A)** Nothing — slices are thread-safe  
**B)** Race condition — concurrent `append` to the same slice is unsafe (and `go run -race` will flag it). Fix with a mutex or by pre-allocating and writing to specific indices.  
**C)** The goroutines will block each other automatically  
**D)** Only works if all goroutines append the same value  

<details><summary>💡 Answer</summary>

**B) Race condition on `append` — fix with mutex or indexed pre-allocation**

```go
// Fix 1: mutex around append:
var mu sync.Mutex
go func(u string) {
    defer wg.Done()
    data := fetch(u)
    mu.Lock()
    results = append(results, data)
    mu.Unlock()
}(url)

// Fix 2: pre-allocate, each goroutine writes its own index (no mutex needed):
results := make([]string, len(urls))
for i, url := range urls {
    wg.Add(1)
    go func(idx int, u string) {
        defer wg.Done()
        results[idx] = fetch(u)  // safe — each goroutine owns its index
    }(i, url)
}
```

Fix 2 is preferred when you know the count ahead of time — no mutex overhead, results stay in order.

</details>

---

### Q23 🔴 — What is the correct way to make a filtered copy of a map under a read lock?

```go
var mu sync.RWMutex
var allData map[string]Item
```

**A)**
```go
mu.RLock()
filtered := allData
mu.RUnlock()
```
**B)**
```go
mu.RLock()
filtered := make(map[string]Item)
for k, v := range allData {
    if matchesFilter(v) { filtered[k] = v }
}
mu.RUnlock()
```
**C)**
```go
filtered := make(map[string]Item)
for k, v := range allData {
    mu.RLock()
    filtered[k] = v
    mu.RUnlock()
}
```
**D)**
```go
mu.Lock()
filtered := allData
mu.Unlock()
```

<details><summary>💡 Answer</summary>

**B) Hold the lock for the entire iteration**

Option A: just copies the map header (pointer) — both maps share the same backing data, still subject to races. Option C: acquires/releases the lock on every iteration — another goroutine could write between iterations, corrupting the iteration. Option D: uses a write lock for a read — works but blocks all readers unnecessarily. Option B holds the read lock for the entire range, which is safe and allows other readers to proceed concurrently.

</details>

---

## 📋 SECTION 4: CHANNELS AND GOROUTINE PATTERNS (7 Questions)

### Q24 🟢 — What is a channel in Go?

**A)** A network connection  
**B)** A typed conduit for sending and receiving values between goroutines — provides synchronized communication  
**C)** A file descriptor  
**D)** A mutex variant  

<details><summary>💡 Answer</summary>

**B) A typed pipe for goroutine communication**

```go
ch := make(chan int)       // unbuffered channel of ints

go func() {
    ch <- 42               // send — blocks until someone receives
}()

value := <-ch              // receive — blocks until someone sends
fmt.Println(value)         // 42

// Buffered channel — doesn't block until buffer is full:
bch := make(chan string, 3)
bch <- "a"                 // doesn't block (buffer has room)
bch <- "b"
bch <- "c"
// bch <- "d"              // would block — buffer full
```

Channels implement Go's philosophy: "Don't communicate by sharing memory; share memory by communicating."

</details>

---

### Q25 🟡 — What is the difference between a buffered and unbuffered channel?

**A)** Buffered channels are faster  
**B)** Unbuffered: both sender and receiver must be ready simultaneously (synchronous); buffered: sender can proceed until the buffer is full, then blocks  
**C)** Buffered channels can hold any type  
**D)** Unbuffered channels can't be closed  

<details><summary>💡 Answer</summary>

**B) Unbuffered = synchronous rendezvous; buffered = async up to capacity**

```go
// Unbuffered — goroutines synchronize:
ch := make(chan int)
go func() { ch <- 1 }()  // blocks until main receives
fmt.Println(<-ch)         // receive unblocks sender

// Buffered — sender proceeds without waiting:
bch := make(chan int, 2)
bch <- 1  // immediate
bch <- 2  // immediate
// bch <- 3  // would block — buffer full
fmt.Println(<-bch)  // 1
fmt.Println(<-bch)  // 2
```

Use unbuffered channels for strict synchronization. Use buffered channels to decouple producers and consumers or to limit concurrency.

</details>

---

### Q26 🟡 — How do you close a channel and why is it important?

**A)** `ch.Close()` — marks the channel as done  
**B)** `close(ch)` — signals to receivers that no more values will be sent; receivers get zero values after close  
**C)** Channels cannot be closed  
**D)** Channels close automatically when the goroutine exits  

<details><summary>💡 Answer</summary>

**B) `close(ch)` — signal that sending is complete**

```go
ch := make(chan int)

go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)  // signal: no more values
}()

// Range over channel — exits when channel is closed:
for v := range ch {
    fmt.Println(v)  // 0, 1, 2, 3, 4
}

// Check if channel is closed:
v, ok := <-ch
// ok == false means channel is closed and empty
```

Only the sender should close a channel. Closing a nil channel or a channel twice panics. Never close a channel from the receiver side.

</details>

---

### Q27 🟡 — What does `select` do with channels?

**A)** Selects the fastest channel  
**B)** Waits on multiple channel operations and executes the first one that's ready — if multiple are ready, it picks one at random  
**C)** Sorts channels by priority  
**D)** Closes all channels at once  

<details><summary>💡 Answer</summary>

**B) Waits on multiple channel operations — runs the first ready one**

```go
ch1 := make(chan string)
ch2 := make(chan string)

select {
case msg1 := <-ch1:
    fmt.Println("received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("received from ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("timeout!")  // default if nothing else is ready in 1s
default:
    fmt.Println("nothing ready")  // non-blocking — immediate if nothing ready
}
```

`select` is how Go programs do non-blocking channel operations, timeouts, and multiplexing multiple data sources. Without `default`, `select` blocks until at least one case is ready.

</details>

---

### Q28 🔴 — What is a goroutine leak and how do you detect one?

**A)** A goroutine that runs faster than expected  
**B)** A goroutine that never terminates because it's blocked waiting on a channel or mutex that will never be satisfied — detected by watching goroutine count with `runtime.NumGoroutine()` or `pprof`  
**C)** A goroutine that uses too much memory  
**D)** A goroutine created in a loop  

<details><summary>💡 Answer</summary>

**B) A goroutine blocked permanently — watch goroutine count to detect**

```go
// Leak: goroutine waiting on a channel that nobody will send to:
func leaky() {
    ch := make(chan int)
    go func() {
        val := <-ch  // blocks forever — nobody sends to ch
        process(val)
    }()
    // ch goes out of scope, goroutine is stuck forever
}

// Fix: use context for cancellation:
func withCancel(ctx context.Context) {
    ch := make(chan int)
    go func() {
        select {
        case val := <-ch:
            process(val)
        case <-ctx.Done():
            return  // goroutine exits when context is cancelled
        }
    }()
}
```

Goroutine leaks slowly exhaust memory. Monitor with `runtime.NumGoroutine()` — if it grows unbounded, you have leaks. Use `context.Context` for cancellation.

</details>

---

### Q29 🔴 — What is the "fan-out, fan-in" concurrency pattern?

**A)** Starting many goroutines and stopping them all  
**B)** Fan-out: one goroutine distributes work to N workers; fan-in: N goroutines produce results that are merged into one result channel  
**C)** Balancing load between servers  
**D)** Using WaitGroup with channels  

<details><summary>💡 Answer</summary>

**B) Distribute work to N workers (fan-out), collect results into one stream (fan-in)**

```go
// Fan-out: send work to N goroutines
func fanOut(input <-chan string, workers int) []<-chan Result {
    channels := make([]<-chan Result, workers)
    for i := 0; i < workers; i++ {
        channels[i] = worker(input)  // each worker reads from shared input
    }
    return channels
}

// Fan-in: merge N result channels into one
func fanIn(channels ...<-chan Result) <-chan Result {
    merged := make(chan Result)
    var wg sync.WaitGroup
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan Result) {
            defer wg.Done()
            for r := range c { merged <- r }
        }(ch)
    }
    go func() { wg.Wait(); close(merged) }()
    return merged
}
```

This pattern maximizes CPU utilization for embarrassingly parallel work like fetching multiple URLs concurrently.

</details>

---

### Q30 🔴 — What is the output?

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

for v := range ch {
    fmt.Print(v, " ")
}
fmt.Println()

v, ok := <-ch
fmt.Println(v, ok)
```

**A)** `1 2 3` then panic  
**B)** `1 2 3` then `0 false` — ranging over a closed channel drains it; receiving from a closed empty channel returns zero value and `false`  
**C)** `1 2 3` then blocks forever  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) `1 2 3 ` then `0 false`**

Buffered channels retain their values after closing. `range` drains the buffer and exits when the channel is closed and empty. Receiving from a closed, empty channel immediately returns the zero value (`0` for `int`) and `false`. This is the correct way to check if a channel is closed: `v, ok := <-ch`.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 28–30 ✅ | **Exceptional** — concurrency mastered. |
| 24–27 ✅ | **Ready** — review any missed sections before starting. |
| 18–23 ⚠️ | **Study first** — goroutines and mutex patterns are central to this project. |
| Below 18 ❌ | **Not ready** — work through "Concurrency" in Go by Example and the `sync` package docs. |

---

## 🔍 Review Map

| Missed | Topic to Study |
|---|---|
| Q1–Q7 | `go` keyword, race conditions, `WaitGroup`, `defer wg.Done()`, deadlock detection, closure-in-loop bug |
| Q8–Q15 | `sync.Mutex`, `defer mu.Unlock()`, `RWMutex`, reentrant deadlock, race detector, thundering herd |
| Q16–Q23 | Filter logic, `r.Form["field"]` for multi-select, `strconv.Atoi` for form numbers, AND/OR filter semantics, unique values, concurrent slice append |
| Q24–Q30 | Channels, buffered vs unbuffered, `close`, `select`, goroutine leaks, fan-out/fan-in, closed channel behavior |