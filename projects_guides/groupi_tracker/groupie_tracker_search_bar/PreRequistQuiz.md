# 🎯 Groupie Tracker Search Bar Prerequisites Quiz
## JavaScript Events · Fetch API · Go JSON Endpoints · Case-Insensitive Matching · DOM Manipulation

**Time Limit:** 50 minutes  
**Total Questions:** 27  
**Passing Score:** 21/27 (78%)

> ✅ Pass → You're ready to start Groupie Tracker Search Bar  
> ⚠️ This project introduces JavaScript client-side logic as a first-class requirement. If you score 21–23, spend extra time on Sections 2 and 3 before starting.

---

## 📋 SECTION 1: RETURNING JSON FROM GO (5 Questions)

### Q1: Your Go handler currently executes an HTML template. To support a search endpoint, you need to return JSON instead. What changes?

**A)** Nothing — Go automatically converts template output to JSON  
**B)** Set `Content-Type: application/json`, then encode your data with `json.NewEncoder(w).Encode(data)` instead of executing a template  
**C)** Use `fmt.Fprintf(w, "%v", data)`  
**D)** Return a different status code  

<details><summary>💡 Answer</summary>

**B) Set `Content-Type: application/json` and use `json.NewEncoder(w).Encode(data)`**

```go
func searchHandler(w http.ResponseWriter, r *http.Request) {
    // ... search logic ...
    
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(suggestions); err != nil {
        http.Error(w, err.Error(), 500)
    }
}
```

`json.NewEncoder(w).Encode` writes JSON directly to the `ResponseWriter`. The `Content-Type` header tells the browser and client code to expect JSON, not HTML.

</details>

---

### Q2: What does `json.NewEncoder(w).Encode(data)` output when `data` is an empty `[]Suggestion{}`?

**A)** `null`  
**B)** `""`  
**C)** `[]` followed by a newline  
**D)** Nothing — empty slices are skipped  

<details><summary>💡 Answer</summary>

**C) `[]` followed by a newline**

An empty Go slice encodes as `[]` (empty JSON array). `json.NewEncoder.Encode` appends a `\n` after the JSON. This is important: JavaScript's `fetch` and `JSON.parse` both handle `[]` correctly — they return an empty array, not an error.

</details>

---

### Q3: Your `Suggestion` struct has an unexported field. What happens when you encode it to JSON?

```go
type Suggestion struct {
    value    string  // unexported
    Category string  // exported
}
```

**A)** Both fields are included  
**B)** Only `Category` is included — unexported fields are invisible to `encoding/json`  
**C)** The entire struct encodes as `null`  
**D)** `json.Marshal` panics  

<details><summary>💡 Answer</summary>

**B) Only exported (capitalized) fields are included**

`encoding/json` uses reflection and can only see exported fields. The `value` field is silently ignored. Always export all fields you want serialized. Add a `json:"value"` tag to control the JSON key name:

```go
type Suggestion struct {
    Value    string `json:"value"`
    Category string `json:"category"`
    ArtistID int    `json:"artistId"`
}
```

</details>

---

### Q4: You want the search endpoint to return an empty array when the query is empty, rather than no response or an error. Why is `[]Suggestion{}` (not `nil`) the right choice?

**A)** No difference — `nil` and empty slice encode identically  
**B)** `nil` encodes as JSON `null`; `[]Suggestion{}` encodes as `[]` — JavaScript code expecting an array will handle `[]` correctly but may break on `null`  
**C)** `nil` causes a panic  
**D)** Empty slice is slower  

<details><summary>💡 Answer</summary>

**B) `nil` → `null`; `[]Suggestion{}` → `[]`**

```go
var a []Suggestion       // nil slice → "null"
b := []Suggestion{}      // empty slice → "[]"
```

In JavaScript: `null.forEach(...)` throws. `[].forEach(...)` runs without error. Always initialize slices before encoding to avoid this trap.

</details>

---

### Q5: Verify with curl that your search endpoint works:
```bash
curl "http://localhost:8080/search?q=queen"
```
The response starts with `[`. Then you test:
```bash
curl "http://localhost:8080/search?q="
```
This returns `null` instead of `[]`. What is the bug?

**A)** The URL is wrong  
**B)** Your search function returns `nil` when the query is empty instead of returning `[]Suggestion{}`  
**C)** curl doesn't support empty query strings  
**D)** The status code is wrong  

<details><summary>💡 Answer</summary>

**B) Returning `nil` instead of `[]Suggestion{}`**

```go
func search(query string, ...) []Suggestion {
    if query == "" {
        return []Suggestion{}  // NOT: return nil
    }
    // ...
}
```

This is a simple fix but a real JavaScript breakage if you don't catch it.

</details>

---

## 📋 SECTION 2: JAVASCRIPT FETCH API (7 Questions)

### Q6: What is the JavaScript `fetch` API used for?

**A)** Fetching local files from disk  
**B)** Making HTTP requests from the browser to a server without reloading the page  
**C)** Loading images faster  
**D)** It's a Go package  

<details><summary>💡 Answer</summary>

**B) Making HTTP requests from the browser without page reload**

`fetch` is the modern way to do AJAX (Asynchronous JavaScript And XML). It returns a Promise that resolves with the server's response. This is how your search bar calls `/search?q=...` and updates the DOM without a full page reload.

</details>

---

### Q7: What is the output of this JavaScript?
```javascript
fetch('/search?q=queen')
  .then(r => r.json())
  .then(data => console.log(data));
```

**A)** Synchronously logs the response immediately  
**B)** Logs the parsed JSON response array after the HTTP request completes (asynchronously)  
**C)** Logs the raw response string  
**D)** Throws an error — fetch requires `async/await`  

<details><summary>💡 Answer</summary>

**B) Logs the parsed JSON array asynchronously after completion**

The chain is: `fetch(url)` → resolves with Response → `.then(r => r.json())` → parses body as JSON → `.then(data => ...)` → data is the parsed value. This is asynchronous — other JavaScript runs while the request is in flight.

</details>

---

### Q8: The fetch request completes but `r.ok` is `false` (server returned 400). What does the `.then(r => r.json())` chain do?

**A)** Automatically throws an error  
**B)** Continues — it parses the response body as JSON regardless of status code. You must explicitly check `r.ok` before parsing.  
**C)** Goes to `.catch()`  
**D)** Returns `null`  

<details><summary>💡 Answer</summary>

**B) `fetch` doesn't throw on 4xx/5xx — you must check `r.ok`**

```javascript
fetch('/search?q=queen')
  .then(r => {
    if (!r.ok) throw new Error(`HTTP error: ${r.status}`);
    return r.json();
  })
  .then(data => displaySuggestions(data))
  .catch(err => console.error(err));
```

This is the same concept as Go's `resp.StatusCode` check — the network didn't fail, but the server returned an error status.

</details>

---

### Q9: Which JavaScript event fires EVERY TIME the user types a character in an input field?

**A)** `change` — fires when input loses focus  
**B)** `keyup` — fires on key release  
**C)** `input` — fires on every value change including typing, paste, autocomplete  
**D)** `submit` — fires when the form is submitted  

<details><summary>💡 Answer</summary>

**C) `input` — fires on every value change**

```javascript
input.addEventListener('input', function() {
    const query = this.value;
    // fires on every character typed, deleted, or pasted
});
```

`change` only fires when focus leaves the field — too late for live suggestions. `keyup` fires on every key press but misses paste and autocomplete. `input` is the correct event for "fire on every text change."

</details>

---

### Q10: You want to wait 300ms after the user stops typing before making the fetch request (debouncing). How do you implement this?

**A)** `setTimeout` inside the `input` listener — cancel the previous timeout on each new keystroke  
**B)** `setInterval` that polls every 300ms  
**C)** Check the timestamp of each keypress  
**D)** Debouncing is not possible in JavaScript  

<details><summary>💡 Answer</summary>

**A) `setTimeout` with cancel-previous (debouncing)**

```javascript
let debounceTimer;
input.addEventListener('input', function() {
    clearTimeout(debounceTimer);  // cancel previous timer
    debounceTimer = setTimeout(() => {
        fetch(`/search?q=${encodeURIComponent(this.value)}`)
          .then(r => r.json())
          .then(displaySuggestions);
    }, 300);  // wait 300ms after last keystroke
});
```

Without debouncing, typing "queen" makes 5 requests (q, qu, que, quee, queen). With debouncing, only 1 request is made after the user pauses. Optional for this project but a major UX improvement.

</details>

---

### Q11: How do you navigate to `/artist?id=5` when the user clicks a suggestion?

**A)** `fetch('/artist?id=5')`  
**B)** `window.location.href = '/artist?id=5'`  
**C)** `document.navigate('/artist?id=5')`  
**D)** `router.push('/artist?id=5')`  

<details><summary>💡 Answer</summary>

**B) `window.location.href = '/artist?id=5'`**

Setting `window.location.href` navigates the browser to the given URL — a full page load, which is exactly what you want when clicking a search result. `fetch` makes a background request without navigation. `router.push` is a React/Vue concept, not vanilla JavaScript.

</details>

---

### Q12: You create suggestion elements dynamically in JavaScript. Why must you call `suggestionBox.innerHTML = ''` before adding new results?

**A)** To reset the CSS  
**B)** To clear old suggestions so new results replace them rather than accumulating below the old ones  
**C)** For performance  
**D)** You don't need to — browser handles it  

<details><summary>💡 Answer</summary>

**B) To clear old suggestions before adding new ones**

Without clearing, each keystroke appends new suggestions below the previous ones. After typing 5 characters, you'd have 5 sets of results stacked. Always clear the container before populating it:

```javascript
suggestionBox.innerHTML = '';  // or: while (box.firstChild) box.removeChild(box.firstChild)
suggestions.forEach(s => {
    const div = document.createElement('div');
    div.textContent = `${s.value} - ${s.category}`;
    div.addEventListener('click', () => window.location.href = `/artist?id=${s.artistId}`);
    suggestionBox.appendChild(div);
});
```

</details>

---

## 📋 SECTION 3: SEARCH LOGIC IN GO (7 Questions)

### Q13: The search must cover 5 categories. For artist name `"Queen"` with member `"Freddie Mercury"`, creation date `1970`, first album `"1973-07-13"`, and location `"london-uk"` — the query `"1"` should match what?

**A)** Nothing — "1" is too short  
**B)** Creation date (1970 contains "1"), first album (1973 contains "1"), and potentially location (london contains no "1" — so only the date matches)  
**C)** All categories — "1" appears in every field  
**D)** Only exact matches  

<details><summary>💡 Answer</summary>

**B) Creation date (1970 contains "1") and first album (1973 contains "1")**

The spec says substring matching. `strings.Contains("1970", "1")` = `true`. `strings.Contains("Freddie Mercury", "1")` = `false`. `strings.Contains("london-uk", "1")` = `false`. The creation date and first album fields both contain "1" as a digit.

</details>

---

### Q14: How do you implement case-insensitive substring matching in Go?

**A)** `a == b` — Go strings are case-insensitive by default  
**B)** `strings.Contains(strings.ToLower(haystack), strings.ToLower(needle))`  
**C)** `strings.EqualFold(haystack, needle)`  
**D)** `regexp.MatchString("(?i)"+needle, haystack)`  

<details><summary>💡 Answer</summary>

**B) `strings.Contains(strings.ToLower(a), strings.ToLower(b))`**

`strings.EqualFold` checks full equality (not substring). For substring: normalize both strings to lowercase first, then use `Contains`. Option D (regex) works but is overkill and slower for this use case.

```go
func matchesQuery(field, query string) bool {
    return strings.Contains(strings.ToLower(field), strings.ToLower(query))
}
```

</details>

---

### Q15: The search must check member names. `artist.Members` is `[]string`. How do you check if the query matches ANY member's name?

**A)** `strings.Contains(fmt.Sprint(artist.Members), query)` — stringify the whole slice  
**B)** Loop through each member and check `matchesQuery(member, query)`  
**C)** `artist.Members == query`  
**D)** `slices.Contains(artist.Members, query)`  

<details><summary>💡 Answer</summary>

**B) Loop through each member and check individually**

```go
for _, member := range artist.Members {
    if matchesQuery(member, query) {
        suggestions = append(suggestions, Suggestion{
            Value:    member,
            Category: "member",
            ArtistID: artist.ID,
        })
    }
}
```

Option A stringifies the slice as `[Freddie Mercury Brian May...]` — a substring search on this could match partial words at the slice boundary. Option D is exact match only.

</details>

---

### Q16: The query is `"phil"`. An artist named `"Phil Collins"` exists as both an artist name and a member of another band. How many suggestions should your `search` function return for this query?

**A)** 1 — deduplicate  
**B)** 2 — one for `artist/band` category and one for `member` category  
**C)** Depends on how many artists match  
**D)** 0 — "phil" is not an exact match  

<details><summary>💡 Answer</summary>

**B) 2 (or more) — one per category match, not one per artist**

The spec says return one suggestion per category-match pair. "Phil Collins" as an artist name → `{value: "Phil Collins", category: "artist/band"}`. "Phil Collins" as a member of another artist → `{value: "Phil Collins", category: "member", artistId: <that other artist's ID>}`. Both should appear.

</details>

---

### Q17: How do you check if `artist.CreationDate` (an `int`, e.g. `1970`) matches the query string `"197"`?

**A)** `artist.CreationDate == 197`  
**B)** `strings.Contains(strconv.Itoa(artist.CreationDate), query)`  
**C)** `artist.CreationDate > 197`  
**D)** `fmt.Sprintf("%d", artist.CreationDate) == query`  

<details><summary>💡 Answer</summary>

**B) Convert the int to string, then use `strings.Contains`**

```go
creationStr := strconv.Itoa(artist.CreationDate)
if matchesQuery(creationStr, query) {
    suggestions = append(suggestions, Suggestion{
        Value:    creationStr,
        Category: "creation date",
        ArtistID: artist.ID,
    })
}
```

This allows `"197"` to match `1970` and `1971`, etc. Option D requires an exact match.

</details>

---

### Q18: For the location category, `relation.DatesLocations` is `map[string][]string`. How do you search through all concert locations for a match?

**A)** `strings.Contains(fmt.Sprint(relation.DatesLocations), query)`  
**B)** Loop through the map keys: `for location := range relation.DatesLocations { if matchesQuery(location, query) ... }`  
**C)** `relation.DatesLocations[query]`  
**D)** Compare to each value (dates), not keys  

<details><summary>💡 Answer</summary>

**B) Loop through map keys**

The keys of `DatesLocations` are the location strings (e.g. `"london-uk"`, `"berlin-germany"`). Loop through keys and check each one. If a location matches, create a suggestion with `category: "location"` and the artist's ID.

```go
for loc := range relation.DatesLocations {
    if matchesQuery(loc, query) {
        suggestions = append(suggestions, Suggestion{
            Value:    loc,
            Category: "location",
            ArtistID: artist.ID,
        })
    }
}
```

</details>

---

### Q19: Should you limit the number of suggestions returned? What happens if the query is `"a"` and every artist matches?

**A)** No limit — return all matches  
**B)** Yes — limit to 10–20 suggestions for usability and performance  
**C)** Return only 1 suggestion  
**D)** The browser automatically limits the dropdown  

<details><summary>💡 Answer</summary>

**B) Yes — limit for usability**

A query of `"a"` could match dozens or hundreds of artists across all 5 categories — hundreds of DOM elements. A practical limit (10–20 total, or 3–5 per category) keeps the dropdown manageable and the response fast.

```go
if len(suggestions) >= 20 { break }
```

Add this check inside your search loops.

</details>

---

## 📋 SECTION 4: DOM MANIPULATION (5 Questions)

### Q20: How do you create a new `<div>` element and set its text in JavaScript?

**A)** `document.write('<div>text</div>')`  
**B)** `const el = document.createElement('div'); el.textContent = 'text';`  
**C)** `innerHTML += '<div>text</div>'`  
**D)** `new HTMLElement('div', 'text')`  

<details><summary>💡 Answer</summary>

**B) `document.createElement` + `.textContent`**

```javascript
const el = document.createElement('div');
el.textContent = `${suggestion.value} - ${suggestion.category}`;
el.className = 'suggestion-item';
el.addEventListener('click', () => window.location.href = `/artist?id=${suggestion.artistId}`);
suggestionBox.appendChild(el);
```

Avoid `innerHTML` with user data — it can cause XSS if any suggestion value contains HTML characters. `.textContent` is always safe.

</details>

---

### Q21: Why is `el.textContent = data` safer than `el.innerHTML = data` when displaying search results?

**A)** `textContent` is faster  
**B)** `innerHTML` parses the string as HTML — if `data` contains `<script>alert('xss')</script>`, it executes. `textContent` treats the string as literal text, never as HTML.  
**C)** `innerHTML` requires a DOM node  
**D)** No difference for text content  

<details><summary>💡 Answer</summary>

**B) `innerHTML` can execute scripts; `textContent` is always safe**

Artist names and location names come from an external API — always treat them as untrusted. If an artist name contained `<img onerror="...">`, using `innerHTML` would execute it. `textContent` always escapes everything.

</details>

---

### Q22: The suggestion dropdown should close when the user clicks anywhere outside it. How do you implement this?

**A)** Add `blur` event to the input  
**B)** Add a `click` event listener on `document` — check if the click target is outside the input AND the suggestion box; if so, clear the box  
**C)** Add `mouseleave` to the suggestion box  
**D)** It closes automatically  

<details><summary>💡 Answer</summary>

**B) Document-level click listener with target check**

```javascript
document.addEventListener('click', function(e) {
    if (!input.contains(e.target) && !suggestionBox.contains(e.target)) {
        suggestionBox.innerHTML = '';
        suggestionBox.style.display = 'none';
    }
});
```

`blur` fires when the input loses focus — but clicking a suggestion briefly takes focus away, causing the dropdown to close before the click registers. The document-level click is the correct approach.

</details>

---

### Q23: How do you show/hide the suggestion dropdown container?

**A)** `display: none` in CSS, toggle with `element.style.display = 'block'`  
**B)** Remove and re-add the element from the DOM  
**C)** Use `visibility: hidden`  
**D)** Move it off-screen with `position: absolute; top: -9999px`  

<details><summary>💡 Answer</summary>

**A) Toggle `display` with CSS + JavaScript**

```css
.suggestions { display: none; }
.suggestions.visible { display: block; }
```

```javascript
// Show:
suggestionBox.style.display = 'block';
// or: suggestionBox.classList.add('visible');

// Hide:
suggestionBox.innerHTML = '';
suggestionBox.style.display = 'none';
```

Using a CSS class toggle is cleaner than inline styles. `visibility: hidden` hides visually but still takes up space in the layout.

</details>

---

### Q24: Your suggestion dropdown appears but is hidden behind other page elements. Which CSS property fixes this?

**A)** `overflow: visible`  
**B)** `z-index: 1000` combined with `position: absolute`  
**C)** `float: above`  
**D)** `display: overlay`  

<details><summary>💡 Answer</summary>

**B) `z-index: 1000` with `position: absolute`**

```css
.suggestions {
    position: absolute;
    z-index: 1000;
    background: white;
    border: 1px solid #ccc;
    width: 100%;
    top: 100%;   /* position directly below the input */
    left: 0;
}
```

`z-index` only works on positioned elements (`position: relative`, `absolute`, or `fixed`). The parent container should have `position: relative` so the dropdown positions correctly below the input.

</details>

---

## 📋 SECTION 5: INTEGRATION (3 Questions)

### Q25: Your search endpoint is `GET /search`. The home page handler is registered on `GET /` (catch-all). What happens if someone visits `GET /search` without you explicitly registering it?

**A)** Go returns 404 automatically  
**B)** The `/` handler is called — you must register `/search` explicitly  
**C)** Go routes to the closest matching handler  
**D)** An error is returned  

<details><summary>💡 Answer</summary>

**B) The `/` catch-all handler is called**

This would cause your home handler to try to parse an HTML template for a search request — likely returning HTML instead of JSON, breaking the JavaScript `fetch`. Always register `/search` explicitly:

```go
http.HandleFunc("/search", searchHandler)
http.HandleFunc("/", homeHandler)
```

Register specific routes BEFORE the catch-all `/`.

</details>

---

### Q26: Your search endpoint returns correct JSON in the browser DevTools Network tab, but the suggestions don't appear on the page. What should you check first?

**A)** The Go handler  
**B)** Open the browser console — there is likely a JavaScript error in the `fetch` callback or the DOM manipulation code  
**C)** The CSS  
**D)** The server logs  

<details><summary>💡 Answer</summary>

**B) The browser console for JavaScript errors**

If the JSON is correct but the UI doesn't update, the problem is in JavaScript. Open DevTools → Console. A `TypeError`, `ReferenceError`, or `SyntaxError` in your fetch callback will silently swallow the response without updating the DOM. Fix the JS error first.

</details>

---

### Q27: A user types `"queen"` very fast. Three fetch requests go out before the first one returns. The second response arrives last (out of order). What could the user see?

**A)** The correct final results for "queen"  
**B)** Results from an earlier shorter query — "q" or "qu" results overwriting the final results  
**C)** No results — all responses are discarded  
**D)** An error  

<details><summary>💡 Answer</summary>

**B) Out-of-order responses can overwrite more recent results**

This is the "stale closure" / "race condition in UI" problem. Solutions:
1. **Debouncing** (simplest): only send a request after 300ms of inactivity — avoids making multiple requests at all
2. **Abort controller**: cancel previous pending requests when a new one starts
3. **Sequence numbering**: track request order, ignore responses that arrive out of order

For this project, debouncing is the simplest and most effective solution.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 25–27 ✅ | **Excellent.** Strong Go + JavaScript fundamentals — start immediately. |
| 21–24 ✅ | **Ready.** Review missed questions, especially the fetch chain and DOM manipulation. |
| 16–20 ⚠️ | **Study first.** Work through the JavaScript fetch API tutorial and DOM manipulation before starting. |
| Below 16 ❌ | **Not ready.** This project requires JavaScript knowledge — work through basic JS DOM manipulation tutorials first. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q5 | `json.NewEncoder`, `Content-Type: application/json`, nil vs empty slice encoding |
| Q6–Q12 | `fetch` chain, `r.ok` check, `input` event, debouncing, `window.location.href`, clearing DOM |
| Q13–Q19 | Search logic per category, case-insensitive match, `strconv.Itoa`, map key iteration, result limits |
| Q20–Q24 | `createElement`, `textContent` vs `innerHTML` (XSS), document click to close, `z-index` |
| Q25–Q27 | Route registration order, browser console debugging, out-of-order fetch responses |