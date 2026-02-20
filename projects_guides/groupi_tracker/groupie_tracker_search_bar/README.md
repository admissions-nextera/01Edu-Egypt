# Groupie Tracker Search Bar Project Guide

> **Before you start:** This project builds on groupie-tracker. Use the search bar on any website you know — notice how suggestions appear as you type, what they show, and how selecting one works. You are building exactly that.

---

## Objectives

By completing this project you will learn:

1. **Real-Time Search** — Filtering data as the user types, not only on form submit
2. **Typing Suggestions** — Displaying contextual autocomplete options below an input
3. **JavaScript Events** — Listening for keyboard input and responding dynamically
4. **Case-Insensitive Matching** — Normalizing strings before comparison
5. **Search Categories** — Labeling each result with what type of data matched
6. **Client-Server Interaction** — Deciding what logic belongs on the server vs the client

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Groupie Tracker (Completed)
- All artist data loaded and accessible
- Artist detail page working

### 2. JavaScript Basics
- `document.getElementById`, `addEventListener`
- `input` event on a text field — fires on every keystroke
- Creating and appending DOM elements
- Search: **"JavaScript input event listener typing"**
- Search: **"JavaScript create element appendChild"**

### 3. Fetch API (JavaScript)
- How to call a Go endpoint from JavaScript without reloading the page
- `fetch(url).then(r => r.json()).then(data => ...)`
- Search: **"JavaScript fetch API tutorial"**

### 4. Go JSON Responses
- Returning JSON from a Go handler instead of HTML
- `json.NewEncoder(w).Encode(data)`
- Setting `Content-Type: application/json`
- Search: **"golang handler return JSON response"**

---

## Project Structure

```
groupie-tracker-search-bar/
├── main.go
├── handlers.go
├── api.go
├── search.go         ← search and suggestion logic
├── templates/
│   ├── index.html    ← now includes the search bar
│   └── artist.html
├── static/
│   ├── style.css
│   └── search.js     ← typing suggestion JavaScript
└── go.mod
```

---

## Milestone 1 — Define What Is Searchable

**This milestone has no code.**

The spec requires these five search categories:

| Category | Example match |
|---|---|
| Artist/band name | "Queen" → `Queen - artist/band` |
| Member name | "Freddie Mercury" → `Freddie Mercury - member` |
| Concert location | "London" → `London, UK - location` |
| First album date | "1973" → `1973 - first album` |
| Creation date | "1970" → `1970 - creation date` |

**Questions to answer before writing anything:**
- One search term can match multiple categories. If the user types `"phil"`, what results should appear? (See spec example.)
- What does each suggestion need to contain? (The matched value and its category label.)
- Is search case-sensitive? (No — the spec says case-insensitive.)
- When a suggestion is clicked, where should the user go?

---

## Milestone 2 — Search Endpoint (`GET /search`)

**Goal:** `GET /search?q=queen` returns a JSON array of suggestions matching the query across all five categories.

**Questions to answer:**
- What Go struct represents one suggestion?
- How do you search through artist names, member names, locations, and dates for a substring match?
- How do you handle the case where the same value appears in multiple categories? (Return one result per category match.)
- What do you return if the query is empty or too short?

**Code Placeholder:**
```go
// search.go

type Suggestion struct {
    // The matched value (e.g. "Freddie Mercury")
    // The category label (e.g. "member")
    // The artist ID to link to (so clicking navigates to the right page)
}

func search(query string, artists []Artist, relations []Relation) []Suggestion {
    // If query is empty, return empty slice

    // Normalize query to lowercase

    // For each artist:
    //   Check artist.Name — if it contains query, add suggestion with category "artist/band"
    //   Check each member in artist.Members — if name contains query, add suggestion with category "member"
    //   Check artist.FirstAlbum — if it contains query, add suggestion with category "first album"
    //   Check artist.CreationDate — if it matches query, add suggestion with category "creation date"
    //   Check each location in the artist's relation — if it contains query, add suggestion with category "location"

    // Return all collected suggestions
}
```

```go
// handlers.go

func searchHandler(w http.ResponseWriter, r *http.Request) {
    // Read "q" query parameter
    // If empty, return empty JSON array with status 200

    // Call search() with the query and loaded data

    // Set Content-Type to application/json
    // Encode and write the suggestions slice as JSON
}
```

**Verify:**
```bash
curl "http://localhost:8080/search?q=queen"
# Should return JSON array of suggestions

curl "http://localhost:8080/search?q=QUEEN"
# Should return the same results (case-insensitive)

curl "http://localhost:8080/search?q="
# Should return empty array
```

---

## Milestone 3 — Search Bar HTML and Basic Display

**Goal:** The page has a search input. As the user types, suggestions appear below it.

**Questions to answer:**
- Should the search bar be its own `<form>` or just a standalone `<input>`?
- What HTML element will you use to display the suggestion list below the input?
- How do you position the suggestion dropdown directly below the input?

**Code Placeholder:**
```html
<!-- In index.html -->

<!-- Search bar input -->
<!-- <input type="text" id="search-input" placeholder="Search artists, members, locations..."> -->

<!-- Suggestion dropdown container -->
<!-- <div id="suggestions" class="suggestions-list"></div> -->

<!-- Link to search.js -->
<!-- <script src="/static/search.js"></script> -->
```

---

## Milestone 4 — Typing Suggestions in JavaScript

**Goal:** Every time the user types a character, the client calls `/search?q=...` and displays the results as a dropdown.

**Questions to answer:**
- Which JavaScript event fires on every keystroke in an input field?
- How do you call the Go search endpoint from JavaScript without refreshing the page?
- How do you create and display suggestion items dynamically in the DOM?
- Each suggestion must show the matched value AND the category. How do you format this?
- What should happen when the user clicks a suggestion?

**Code Placeholder:**
```javascript
// static/search.js

const input = document.getElementById('search-input');
const suggestionBox = document.getElementById('suggestions');

input.addEventListener('input', function() {
    const query = this.value.trim();

    // If query is empty: clear and hide suggestion box, return

    // Fetch /search?q=<query>
    // On success:
    //   Clear the suggestion box
    //   For each suggestion:
    //     Create a <div> element
    //     Set its text to: "MatchedValue - category"
    //     Add a click handler that navigates to the artist's page
    //     Append to suggestion box
    //   Show the suggestion box
    // On error: clear suggestion box
});

// Hide suggestions when clicking outside the input or suggestion box
document.addEventListener('click', function(e) {
    // If click target is not the input or the suggestion box:
    //   Clear and hide the suggestion box
});
```

**Verify:**
- Type `"phil"` — suggestions for Phil Collins appear as both `member` and `artist/band`
- Type `"QUEEN"` — returns the same results as `"queen"`
- Type something with no matches — suggestion box is empty or hidden
- Click a suggestion — navigates to the correct artist page
- Click outside the suggestion box — it closes

---

## Milestone 5 — Keyboard Navigation (Optional but Good Practice)

**Goal:** The user can move through suggestions with arrow keys and select with Enter.

**Questions to answer:**
- Which JavaScript event handles arrow key and Enter key presses?
- How do you track which suggestion is currently highlighted?
- How do you visually indicate the highlighted suggestion?

**Code Placeholder:**
```javascript
// Add to search.js

let activeIndex = -1;

input.addEventListener('keydown', function(e) {
    const items = suggestionBox.querySelectorAll('.suggestion-item');

    // If ArrowDown: increment activeIndex, highlight that item
    // If ArrowUp: decrement activeIndex, highlight that item
    // If Enter and activeIndex >= 0: navigate to that suggestion's artist page
    // If Escape: clear and hide suggestion box
});
```

---

## Debugging Checklist

- Are suggestions not appearing? Open the browser console and check for JavaScript errors. Then check the Network tab — is the `/search` request returning 200 with valid JSON?
- Is the search case-sensitive? Make sure you normalize both the query and the data to lowercase before comparing.
- Does `"phil"` only return one result instead of two? Your search is stopping after the first match per artist — make sure you check all five categories for every artist.
- Do suggestions stay visible after clicking one? Your click handler needs to also clear the suggestion box.
- Are locations formatted strangely? The API uses `"new_york-usa"` format — normalize underscores and hyphens for display.

---

## Key Packages and APIs

| Item | What You Use It For | Docs |
|---|---|---|
| `encoding/json` | Return JSON from search endpoint | https://pkg.go.dev/encoding/json |
| `strings` | Case-insensitive matching | https://pkg.go.dev/strings |
| `net/http` | Register `/search` route | https://pkg.go.dev/net/http |
| JavaScript `fetch` | Call search endpoint from browser | https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API |
| JavaScript `input` event | React to every keystroke | https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/input_event |

---

## Submission Checklist

- [ ] Search endpoint `GET /search?q=` returns JSON suggestions
- [ ] Suggestions include all five categories: artist/band, member, location, first album, creation date
- [ ] Each suggestion shows the matched value and its category label
- [ ] Search is case-insensitive
- [ ] Typing suggestions appear below the input as the user types
- [ ] Clicking a suggestion navigates to the correct artist page
- [ ] Clicking outside the suggestion box closes it
- [ ] Empty query returns no suggestions
- [ ] No crashes on any search input
- [ ] Unit tests for the `search` function covering all five categories