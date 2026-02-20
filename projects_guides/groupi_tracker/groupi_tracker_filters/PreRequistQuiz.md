# 🎯 Groupie Tracker Filters Prerequisites Quiz
## Goroutines · WaitGroups · Race Conditions · Filter Logic · HTML Form Controls

**Time Limit:** 50 minutes  
**Total Questions:** 28  
**Passing Score:** 22/28 (78%)

> ✅ Pass → You're ready to start Groupie Tracker Filters  
> ⚠️ Also Required → Groupie Tracker must be fully working with all four endpoints

---

## 📋 SECTION 1: GOROUTINES & CONCURRENCY (8 Questions)

### Q1: What is a goroutine?

**A)** A Go function that returns an error  
**B)** A lightweight thread managed by the Go runtime — started with the `go` keyword, runs concurrently with other code  
**C)** A loop that retries failed operations  
**D)** A type of channel  

<details><summary>💡 Answer</summary>

**B) A lightweight thread managed by the Go runtime, started with `go`**

```go
go fetchArtists()    // starts concurrently, doesn't wait
go fetchLocations()  // starts concurrently at the same time
```

Goroutines are cheap (a few KB of stack vs MB for OS threads). You can run thousands of them. The Go scheduler multiplexes them onto OS threads.

</details>

---

### Q2: You launch 4 goroutines to fetch 4 API endpoints. How do you wait for ALL of them to finish before proceeding?

**A)** `time.Sleep(5 * time.Second)`  
**B)** `sync.WaitGroup` — call `wg.Add(4)` before launching, `wg.Done()` in each goroutine, `wg.Wait()` after launching  
**C)** `go.Wait()` built-in function  
**D)** Check all results in a for loop  

<details><summary>💡 Answer</summary>

**B) `sync.WaitGroup`**

```go
var wg sync.WaitGroup
wg.Add(4)

go func() { defer wg.Done(); artists, err = fetchArtists() }()
go func() { defer wg.Done(); locations, err = fetchLocations() }()
go func() { defer wg.Done(); dates, err = fetchDates() }()
go func() { defer wg.Done(); relations, err = fetchRelations() }()

wg.Wait()  // blocks until all 4 goroutines call Done()
```

`time.Sleep` is wrong — you don't know how long the requests take. `wg.Wait()` is precise.

</details>

---

### Q3: What is a race condition?

**A)** When one goroutine runs faster than another  
**B)** When two goroutines access the same memory location concurrently, and at least one is writing — the outcome is undefined and non-deterministic  
**C)** When the network is slow  
**D)** A type of deadlock  

<details><summary>💡 Answer</summary>

**B) Concurrent access to shared memory with at least one writer**

```go
// RACE: both goroutines write to `err` without synchronization
var err error
go func() { err = doSomething() }()
go func() { err = doOther() }()
```

The final value of `err` is unpredictable. Race conditions cause bugs that appear randomly and are hard to reproduce. Go has a built-in race detector: `go run -race .`

</details>

---

### Q4: How do you safely store results from multiple goroutines?

**A)** Use regular variables — goroutines are synchronized automatically  
**B)** Use `sync.Mutex` to protect writes, or use separate variables per goroutine and collect after `wg.Wait()`  
**C)** Use global variables  
**D)** Goroutines can't share data  

<details><summary>💡 Answer</summary>

**B) Use a mutex for writes OR use separate variables and collect after `wg.Wait()`**

The cleanest approach for independent fetches: each goroutine writes to its own dedicated variable (no sharing), then after `wg.Wait()` all results are collected:

```go
var (
    artists   []Artist
    locations []Location
    fetchErr  error
    mu        sync.Mutex
)

go func() {
    defer wg.Done()
    result, err := fetchArtists()
    mu.Lock()
    artists = result
    if err != nil { fetchErr = err }
    mu.Unlock()
}()
```

</details>

---

### Q5: What does `go run -race .` do?

**A)** Runs the program faster  
**B)** Enables the race detector — reports any data races at runtime with goroutine stack traces  
**C)** Benchmarks goroutine performance  
**D)** Limits the program to one goroutine  

<details><summary>💡 Answer</summary>

**B) Enables the race detector**

The race detector instruments your code to detect concurrent memory access. Always run it during development when using goroutines. It adds overhead (~5–10x slower) so don't ship it in production — but every race it finds would be a real bug.

</details>

---

### Q6: `defer wg.Done()` vs `wg.Done()` at the end of the goroutine — which is safer and why?

**A)** They are identical  
**B)** `defer wg.Done()` is safer — it runs even if the goroutine panics or returns early, preventing `wg.Wait()` from blocking forever  
**C)** `wg.Done()` at the end is safer — defer has overhead  
**D)** `defer` doesn't work inside goroutines  

<details><summary>💡 Answer</summary>

**B) `defer wg.Done()` is safer**

If the goroutine returns early (due to an error check) or panics, `wg.Done()` at the end would never be called and `wg.Wait()` would block forever. `defer` guarantees it runs regardless. Always use `defer wg.Done()` as the first line inside a goroutine that participates in a WaitGroup.

</details>

---

### Q7: How does fetching 4 endpoints concurrently improve performance compared to sequentially?

**A)** It doesn't — concurrency adds overhead  
**B)** Sequentially: total time ≈ sum of all request times. Concurrently: total time ≈ slowest individual request  
**C)** Concurrency makes each request faster  
**D)** Only if the endpoints are on the same server  

<details><summary>💡 Answer</summary>

**B) Concurrent total time ≈ slowest request; sequential total time ≈ sum of all**

If each endpoint takes ~200ms: Sequential = 4 × 200ms = 800ms. Concurrent = max(200ms, 200ms, 200ms, 200ms) = 200ms. For I/O-bound work like HTTP requests, concurrency gives roughly 4x speedup with 4 goroutines.

</details>

---

### Q8: What is a deadlock and what causes it with WaitGroups?

**A)** When the program runs out of memory  
**B)** When `wg.Wait()` blocks forever because some goroutines never called `wg.Done()` — usually because `Add` count doesn't match `Done` count  
**C)** When two goroutines call `wg.Wait()` simultaneously  
**D)** When goroutines are faster than the main goroutine  

<details><summary>💡 Answer</summary>

**B) `wg.Wait()` blocks forever when `Done` count doesn't reach `Add` count**

Common causes:
- `wg.Add(4)` but one goroutine panics and `defer wg.Done()` isn't used
- `wg.Add(3)` but you actually launch 4 goroutines
- A goroutine returns early without calling `Done`

Go detects simple deadlocks and panics with "all goroutines are asleep - deadlock!"

</details>

---

## 📋 SECTION 2: FILTER LOGIC (7 Questions)

### Q9: You need to filter artists by creation date range. The user provides `minYear=1970` and `maxYear=1985`. How should you handle the case where `minYear > maxYear`?

**A)** Panic  
**B)** Return an empty result  
**C)** Swap them silently, or return a 400 error with a clear message  
**D)** Ignore it and apply only the minYear filter  

<details><summary>💡 Answer</summary>

**C) Swap them silently, or return 400 with a clear message**

Both approaches are defensible. Swapping is more user-friendly (assume the user made a mistake). Returning 400 is more correct (invalid input from the client). Pick one and be consistent. Never silently produce wrong results.

</details>

---

### Q10: `artist.FirstAlbum` is the string `"13-07-1998"`. How do you extract the year `1998` from it?

**A)** `artist.FirstAlbum[0:4]`  
**B)** `strings.Split(artist.FirstAlbum, "-")[2]` → then `strconv.Atoi`  
**C)** `strconv.Atoi(artist.FirstAlbum)`  
**D)** `time.Parse("02-01-2006", artist.FirstAlbum).Year()`  

<details><summary>💡 Answer</summary>

**B or D — both work; D is more robust**

The format is `"DD-MM-YYYY"`. `strings.Split(..., "-")[2]` gives `"1998"`, then `strconv.Atoi` converts it. This works if the format is guaranteed. `time.Parse` is more robust if the format might vary — it also validates the date.

```go
// Simple approach:
parts := strings.Split(artist.FirstAlbum, "-")
year, _ := strconv.Atoi(parts[2])

// Robust approach:
t, err := time.Parse("02-01-2006", artist.FirstAlbum)
year := t.Year()
```

</details>

---

### Q11: What should `applyFilters` return when NO filters are set (all filter fields at defaults)?

**A)** An empty slice  
**B)** The full unfiltered artists slice  
**C)** An error  
**D)** Nil  

<details><summary>💡 Answer</summary>

**B) The full unfiltered artists slice**

When no filters are active, all artists match by default. Your filter logic should pass every artist through if the filter parameters represent "no constraint." For a range filter: if `minYear = 0` and `maxYear = 0`, treat as "no year constraint." Make sure defaults are sensible (e.g., `minYear = 0`, `maxYear = 9999`).

</details>

---

### Q12: How do you filter by number of members when the user selects multiple checkboxes (e.g., 2 and 4 members)?

**A)** `artist.Members == selectedCount`  
**B)** Check if `len(artist.Members)` is in the selected counts slice  
**C)** Filter only by the first selected count  
**D)** Use the AND operator — artist must have BOTH 2 AND 4 members  

<details><summary>💡 Answer</summary>

**B) Check if `len(artist.Members)` is in the selected counts slice**

```go
func hasSelectedMemberCount(artist Artist, counts []int) bool {
    if len(counts) == 0 { return true }  // no filter = accept all
    for _, c := range counts {
        if len(artist.Members) == c { return true }
    }
    return false
}
```

Multiple checkboxes mean OR logic — artist matches if their member count equals ANY of the selected values.

</details>

---

### Q13: For the "5+ members" checkbox, how do you check if an artist qualifies?

**A)** `len(artist.Members) == 5`  
**B)** `len(artist.Members) >= 5`  
**C)** `len(artist.Members) > 5`  
**D)** `artist.Members[5] != ""`  

<details><summary>💡 Answer</summary>

**B) `len(artist.Members) >= 5`**

"5+" means 5 or more. The checkbox value "5" represents "5 or more." Handle this as a special case:

```go
if selectedCount == 5 {
    return len(artist.Members) >= 5
}
return len(artist.Members) == selectedCount
```

</details>

---

### Q14: You apply 4 filters simultaneously. The logic between filters should be AND or OR?

**A)** OR — artist must match any one filter  
**B)** AND — artist must match all active filters  
**C)** Depends on which filter  
**D)** The user chooses  

<details><summary>💡 Answer</summary>

**B) AND — the artist must pass ALL active filters**

This is standard filtering behavior. If you filter by "creation date 1970–1985" AND "2 members", only artists that match BOTH criteria appear. Within a single filter that accepts multiple values (like member count checkboxes), the logic is OR. Between different filters, the logic is AND.

</details>

---

### Q15: How do you compute the minimum and maximum creation years across all artists dynamically (to set range slider bounds)?

**A)** Hardcode 1900 and 2024  
**B)** Loop through all artists and track the running min and max of `artist.CreationDate`  
**C)** The API provides min/max in a separate endpoint  
**D)** Sort the artists by creation date and take first and last  

<details><summary>💡 Answer</summary>

**B) Loop and track running min/max**

```go
func yearBounds(artists []Artist) (min, max int) {
    min, max = artists[0].CreationDate, artists[0].CreationDate
    for _, a := range artists[1:] {
        if a.CreationDate < min { min = a.CreationDate }
        if a.CreationDate > max { max = a.CreationDate }
    }
    return
}
```

Dynamic bounds are better than hardcoded values — they adapt if the dataset changes.

</details>

---

## 📋 SECTION 3: HTML FORM CONTROLS (6 Questions)

### Q16: How does an `<input type="range">` send its value in a form?

**A)** It doesn't — range inputs require JavaScript  
**B)** When the form is submitted, the slider's current value is sent as the value of the input's `name` attribute  
**C)** It sends the percentage position of the slider  
**D)** It sends two values: current and max  

<details><summary>💡 Answer</summary>

**B) The current value is sent under the input's `name`**

```html
<input type="range" name="minYear" min="1960" max="2024" value="1970">
```

When submitted, the form sends `minYear=1970` (or whatever the user dragged it to). For a dual-range filter (min AND max), you need two separate inputs with different `name` attributes.

</details>

---

### Q17: How do multiple checkboxes with the same `name` attribute send data?

```html
<input type="checkbox" name="members" value="1">
<input type="checkbox" name="members" value="2">
<input type="checkbox" name="members" value="3">
```

**A)** Only the last checked value is sent  
**B)** Each checked checkbox sends a separate `members=value` entry — the server receives them as a list  
**C)** They can't share the same name  
**D)** They are concatenated: `members=1,2,3`  

<details><summary>💡 Answer</summary>

**B) Each checked checkbox sends a separate entry — read as a list on the server**

If the user checks values 2 and 3, the form sends `members=2&members=3`. In Go: `r.URL.Query()["members"]` returns `[]string{"2", "3"}`. Note: `r.URL.Query().Get("members")` only returns the first value — use the multi-value form when checkboxes are involved.

</details>

---

### Q18: How do you read multiple values for the same query parameter key in Go?

**A)** `r.URL.Query().Get("members")` — returns all values  
**B)** `r.URL.Query()["members"]` — returns `[]string` of all values for that key  
**C)** `r.FormValue("members")` — returns all values as a comma-separated string  
**D)** Multiple values for the same key are not supported  

<details><summary>💡 Answer</summary>

**B) `r.URL.Query()["members"]` returns `[]string`**

```go
memberValues := r.URL.Query()["members"]  // []string{"2", "3"}
for _, v := range memberValues {
    count, err := strconv.Atoi(v)
    // ...
}
```

`r.URL.Query().Get("members")` only returns the FIRST value — always use the map index `["key"]` for checkboxes.

</details>

---

### Q19: How do you make an HTML range input display its current value to the user as they drag it?

**A)** It shows automatically  
**B)** Use JavaScript: listen to the `input` event and update a `<span>` showing the current value  
**C)** Use `<input type="range" showvalue="true">`  
**D)** It's not possible without a third-party library  

<details><summary>💡 Answer</summary>

**B) JavaScript: `input` event → update a display element**

```html
<input type="range" id="minYear" name="minYear" min="1960" max="2024" 
       oninput="document.getElementById('minYearDisplay').textContent = this.value">
<span id="minYearDisplay">1960</span>
```

The `oninput` handler fires on every drag. Without this, the user can't see the value they're selecting until they submit.

</details>

---

### Q20: Should filters apply on form submit or as the user changes them (live)?

**A)** Always live — no submit button needed  
**B)** Either works — live requires JavaScript to submit on change; submit button is simpler with pure HTML  
**C)** Always on submit — live filtering is too slow  
**D)** The spec requires live filtering  

<details><summary>💡 Answer</summary>

**B) Either approach works — submit button is simpler**

For a pure HTML/Go solution: use a submit button (or `<input type="submit">`). For live filtering: add `onchange="this.form.submit()"` to each filter input, or use JavaScript fetch. The submit approach is easier to get right and less prone to race conditions. Choose based on what you can implement correctly.

</details>

---

### Q21: After a filter form is submitted and the page reloads, how do you keep the filter controls showing the user's previous selections?

**A)** Use browser cookies automatically  
**B)** Read the submitted values in the handler and pass them to the template; use template logic to set the `value`, `checked`, or `selected` attributes  
**C)** It's not possible without JavaScript  
**D)** The browser remembers form state automatically  

<details><summary>💡 Answer</summary>

**B) Read submitted values → pass to template → conditionally render `checked`/`value`**

```go
type PageData struct {
    Artists         []Artist
    FilterParams    FilterParams  // includes selected member counts, year bounds, etc.
}
```

```html
<input type="range" name="minYear" value="{{ .FilterParams.MinCreationYear }}">
<input type="checkbox" name="members" value="2"
    {{ if memberSelected .FilterParams.Members 2 }}checked{{ end }}>
```

This is the same principle as keeping the banner selector state in ASCII-Art-Web.

</details>

---

## 📋 SECTION 4: LOCATION MATCHING (4 Questions)

### Q22: The API returns location strings like `"new_york-usa"`. What normalizations should you apply before comparing with user-selected filter values?

**A)** No normalization needed  
**B)** Replace `_` with space, replace `-` with `,`, convert to lowercase  
**C)** Replace `-` with space only  
**D)** Convert to uppercase  

<details><summary>💡 Answer</summary>

**B) Replace `_` with space, replace `-` with `,`, lowercase**

```go
func normalizeLocation(s string) string {
    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, "_", " ")
    s = strings.ReplaceAll(s, "-", ", ")
    return s
}
// "new_york-usa" → "new york, usa"
```

This makes the display readable AND allows substring matching (e.g., "usa" matching "new york, usa").

</details>

---

### Q23: The spec says selecting "washington, usa" should also match "seattle, washington, usa". How do you implement this?

**A)** Exact string equality  
**B)** `strings.Contains(normalizedArtistLocation, normalizedSelectedLocation)`  
**C)** `strings.HasPrefix`  
**D)** Regular expression matching  

<details><summary>💡 Answer</summary>

**B) `strings.Contains(artistLocation, selectedLocation)`**

```go
func locationMatches(artistLoc, selected string) bool {
    a := normalizeLocation(artistLoc)
    s := normalizeLocation(selected)
    return strings.Contains(a, s)
}
// normalizeLocation("seattle-washington-usa") = "seattle, washington, usa"
// normalizeLocation("washington-usa") = "washington, usa"
// strings.Contains("seattle, washington, usa", "washington, usa") = true ✓
```

</details>

---

### Q24: How do you build the list of all unique locations to show as filter checkboxes?

**A)** Hardcode a list of countries  
**B)** Loop through all artists' relation data, collect all location keys into a set (map), extract the keys  
**C)** The API provides a `/locations/unique` endpoint  
**D)** Use only the first artist's locations  

<details><summary>💡 Answer</summary>

**B) Collect all location keys into a map (set) then extract unique values**

```go
func allLocations(relations []Relation) []string {
    seen := map[string]bool{}
    for _, rel := range relations {
        for loc := range rel.DatesLocations {
            seen[loc] = true
        }
    }
    var result []string
    for loc := range seen {
        result = append(result, loc)
    }
    sort.Strings(result)  // sort for consistent display
    return result
}
```

</details>

---

### Q25: When should you compute filter bounds (min/max years, all locations) — on every request or once at startup?

**A)** On every request — the data might change  
**B)** Once at startup alongside the data fetch — the dataset is static, and recomputing on every request is wasteful  
**C)** Only when a filter request comes in  
**D)** Every 60 seconds  

<details><summary>💡 Answer</summary>

**B) Once at startup**

Derive min/max years and the unique location list immediately after fetching the API data. Store them alongside the artist data. Since the data doesn't change, these values are stable for the lifetime of the server.

</details>

---

## 📋 SECTION 5: TRICKY INTEGRATION CASES (3 Questions)

### Q26: You run `go run -race .` and see this output:
```
WARNING: DATA RACE
Write at 0x00c000... by goroutine 7
Read at 0x00c000... by goroutine 8
```
What does this mean and how do you fix it?

**A)** Your code has a bug that could cause random, hard-to-reproduce failures — add a `sync.Mutex` around the shared variable  
**B)** Normal output — ignore it  
**C)** Your goroutines are too slow  
**D)** A network error  

<details><summary>💡 Answer</summary>

**A) A real race condition — add mutex protection**

```go
var (
    result []Artist
    mu     sync.Mutex
)

go func() {
    defer wg.Done()
    data, _ := fetchArtists()
    mu.Lock()
    result = data
    mu.Unlock()
}()
```

Never ignore race warnings — they represent real bugs that will manifest as mysterious crashes or wrong data in production.

</details>

---

### Q27: Your filter form uses `method="GET"`. After submitting, you see the filter values in the URL: `/?minYear=1970&maxYear=1985&members=2&members=4`. Is this correct behavior?

**A)** No — filter data should use POST  
**B)** Yes — using GET for filters means the filtered URL can be bookmarked and shared, which is better UX  
**C)** It exposes private data  
**D)** GET forms don't support multiple values for the same key  

<details><summary>💡 Answer</summary>

**B) Yes — GET is correct for filters**

Filters are query parameters, not data being submitted to the server. Using GET means the filtered URL is bookmarkable and shareable. This is the correct semantic. POST is for actions that change server state — filtering is a read operation.

</details>

---

### Q28: A user applies the creation date filter with `minYear=1980` and `maxYear=1975`. Your `applyFilters` function receives these values. Which behavior is most user-friendly?

**A)** Return an error 400  
**B)** Return all artists (ignore the invalid filter)  
**C)** Swap minYear and maxYear silently and apply the filter correctly  
**D)** Return an empty result  

<details><summary>💡 Answer</summary>

**C) Swap and apply correctly**

```go
if params.MinCreationYear > params.MaxCreationYear {
    params.MinCreationYear, params.MaxCreationYear = 
        params.MaxCreationYear, params.MinCreationYear
}
```

Returning an error for an inverted range is unnecessarily strict — it's likely a user mistake. Silently correcting it and showing results is more forgiving. Never return empty results for what could be a valid (just inverted) query.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 26–28 ✅ | **Excellent.** Start Groupie Tracker Filters. |
| 22–25 ✅ | **Ready.** Review missed questions — especially WaitGroup and location matching. |
| 17–21 ⚠️ | **Study first.** Practice goroutine patterns with `sync.WaitGroup` before starting. |
| Below 17 ❌ | **Not ready.** Goroutines and race conditions will block you. Work through the sync package docs and examples first. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q8 | Goroutines, `sync.WaitGroup`, race conditions, `-race` flag, deadlocks |
| Q9–Q15 | Filter logic, AND vs OR, year extraction, member count `>=`, dynamic bounds |
| Q16–Q21 | Range inputs, checkbox multi-values, `r.URL.Query()["key"]`, form state retention |
| Q22–Q25 | Location normalization, `strings.Contains` partial matching, unique location set |
| Q26–Q28 | Race detector, GET for filters, graceful invalid-input handling |