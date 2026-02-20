# Groupie Tracker Filters Project Guide

> **Before you start:** This project builds on groupie-tracker. Your data fetching, server, and templates must all be working. Think through how filtering works in any website you use before writing any code.

---

## Objectives

By completing this project you will learn:

1. **Filter Logic** — Implementing range and checkbox filters that narrow a dataset
2. **Form Handling** — Reading multiple filter values from a single form submission
3. **Goroutines** — Using concurrency to fetch multiple API endpoints simultaneously
4. **Location Matching** — Handling hierarchical location data (city is part of region is part of country)
5. **Data Aggregation** — Deriving filter bounds (min/max years, member counts) from the dataset itself

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Groupie Tracker (Completed)
- All four API endpoints fetched and decoded
- Artist list and detail pages working

### 2. Go Goroutines and WaitGroups
- Launching goroutines with `go func()`
- `sync.WaitGroup` — waiting for multiple goroutines
- Search: **"golang sync WaitGroup concurrent HTTP requests"**

### 3. HTML Form Controls
- `<input type="range">` — slider for range filters
- `<input type="checkbox">` — multi-select filters
- `name` and `value` attributes — how form data maps to server parameters
- Search: **"HTML input type range example"**
- Search: **"HTML checkbox form POST multiple values"**

### 4. URL Query Parameters
- How to read multiple values for the same key: `r.URL.Query()["members"]`
- Search: **"golang read multiple query parameters same key"**

---

## Project Structure

```
groupie-tracker-filters/
├── main.go
├── handlers.go
├── api.go
├── filter.go         ← new file for all filter logic
├── templates/
│   ├── index.html    ← now includes filter controls
│   └── artist.html
├── static/
│   └── style.css
└── go.mod
```

---

## Milestone 1 — Fetch API Data Concurrently with Goroutines

**Goal:** All four API endpoints are fetched at the same time instead of one after the other.

**Questions to answer before writing anything:**
- What is the current bottleneck when fetching four endpoints sequentially?
- How do you launch multiple goroutines and wait for all of them to finish?
- How do you safely collect results from multiple goroutines without a race condition?
- What happens if one of the four requests fails while the others succeed?

**Code Placeholder:**
```go
// api.go

func fetchAllData() ([]Artist, []Location, []Dates, []Relation, error) {
    // Create variables to hold each result and each error

    // Create a WaitGroup
    // Launch 4 goroutines — one per endpoint
    // Each goroutine fetches its endpoint and stores the result
    // Use a mutex or channels to safely write results from goroutines

    // Wait for all goroutines to finish
    // Check if any errors occurred
    // Return all four results
}
```

**Resources:**
- Search: **"golang fetch multiple URLs concurrently"**
- https://pkg.go.dev/sync#WaitGroup

**Verify:** Add timing logs before and after — concurrent fetching should be noticeably faster than sequential.

---

## Milestone 2 — Define the Four Filters

**Goal:** Understand what each filter covers and how it maps to the data.

**This milestone has no code.** Answer these questions first:

| Filter | Data source | Type | What it filters |
|---|---|---|---|
| Creation date | `artist.CreationDate` | Range (min–max year) | Artists created between two years |
| First album date | `artist.FirstAlbum` | Range (min–max year) | Artists whose first album is between two years |
| Number of members | `len(artist.Members)` | Checkbox (1, 2, 3, 4, 5+) | Artists with selected member counts |
| Concert locations | `relation.DatesLocations` keys | Checkbox | Artists with concerts in selected locations |

**Questions to answer:**
- `FirstAlbum` is a string like `"13-07-1998"`. How do you extract just the year from it?
- For the location filter, if a user selects "Washington, USA", should "Seattle, Washington, USA" match? (Yes — read the spec hint carefully.)
- How do you derive the minimum and maximum years across all artists to set your range slider bounds?
- How do you build the list of all unique locations for the checkbox options?

---

## Milestone 3 — Build the Filter UI

**Goal:** The home page has working filter controls above the artist list.

**Questions to answer:**
- How does an HTML range input send its min and max values in a form?
- How do checkboxes with the same `name` send multiple values to the server?
- Should filters apply on form submit, or live as the user changes them?
- How do you pre-fill the filter controls with the user's previous selections after a page reload?

**Code Placeholder:**
```html
<!-- In index.html -->

<!-- Filter form — wraps all filter controls and the artist list -->
<!-- <form method="GET" action="/"> -->

<!-- Creation date range filter -->
<!-- Two inputs: min year and max year, or a dual-handle range slider -->

<!-- First album range filter -->
<!-- Same pattern -->

<!-- Number of members checkboxes -->
<!-- One checkbox per possible member count (1, 2, 3, 4, 5+) -->

<!-- Locations checkboxes -->
<!-- One checkbox per unique location from the data -->

<!-- Submit button (or auto-submit on change) -->

<!-- Artist cards rendered here, filtered by the server -->
```

---

## Milestone 4 — Implement the Filter Logic

**Goal:** When the form is submitted, the server filters the artist list and returns only matching artists.

**Questions to answer:**
- Where does the filter logic live — in the handler or in a separate function?
- How do you parse the year range values from query parameters?
- For the location filter, how do you check if an artist has any concert in a location that contains the selected string?
- If no filters are set, should all artists be returned? (Yes.)
- What happens if a filter value is invalid (e.g. min year > max year)?

**Code Placeholder:**
```go
// filter.go

type FilterParams struct {
    // Creation date min and max
    // First album year min and max
    // Selected member counts (slice)
    // Selected locations (slice)
}

func parseFilters(r *http.Request) FilterParams {
    // Read creation date min/max from query params
    // Read first album min/max from query params
    // Read member counts (multiple values, same key)
    // Read locations (multiple values, same key)
    // Return the parsed FilterParams
}

func applyFilters(artists []Artist, relations []Relation, params FilterParams) []Artist {
    result := []Artist{}

    for _, artist := range artists {
        // Check creation date range — skip if outside bounds
        // Check first album year range — skip if outside bounds
        // Check member count — skip if not in selected counts
        // Check location — skip if none of the artist's locations match

        // If all checks pass, append to result
    }

    return result
}
```

**Verify:**
- Set creation date range 1970–1980 — only artists from that era appear
- Select "2 members" checkbox — only duos appear
- Select a location — only artists with concerts there appear
- Combine multiple filters — all conditions apply simultaneously
- No filters selected — all artists appear

---

## Milestone 5 — Location Partial Matching

**Goal:** Selecting "Washington, USA" also matches "Seattle, Washington, USA".

**Questions to answer:**
- How do you check if a location string contains another location string as a substring?
- The API uses hyphen-separated location strings like `"new_york-usa"`. How do you normalize them for comparison?
- Should the match be case-insensitive?

**Code Placeholder:**
```go
// filter.go

func locationMatches(artistLocation string, selectedLocation string) bool {
    // Normalize both strings (lowercase, replace _ with space, replace - with , )
    // Check if artistLocation contains selectedLocation as a substring
}
```

**Verify:**
- Adding "washington-usa" as a filter matches both "washington-usa" and "seattle-washington-usa"
- Matching is not case sensitive

---

## Debugging Checklist

- Do goroutines cause a race condition? Run with `go run -race .` and fix any warnings.
- Are checkbox values not arriving at the server? Check that your checkbox `name` attributes match what you are reading with `r.URL.Query()["name"]`.
- Is the location filter not matching? Print both strings before comparing — there may be underscores, hyphens, or capitalization differences.
- Does filtering break when no filter values are submitted? Make sure your `parseFilters` function handles empty/missing parameters by setting sensible defaults (e.g. full year range, all member counts).
- Are range sliders only sending one value? Make sure you have two separate inputs with different names for min and max.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `sync` | WaitGroup for concurrent fetching | https://pkg.go.dev/sync |
| `strings` | Normalize and match location strings | https://pkg.go.dev/strings |
| `strconv` | Parse year and member count values | https://pkg.go.dev/strconv |
| `net/http` | Handle filtered GET requests | https://pkg.go.dev/net/http |

---

## Submission Checklist

- [ ] All API endpoints fetched concurrently with goroutines
- [ ] `go run -race .` reports no race conditions
- [ ] Filter by creation date range works correctly
- [ ] Filter by first album year range works correctly
- [ ] Filter by number of members (checkbox) works correctly
- [ ] Filter by concert location (checkbox) works correctly
- [ ] Location filter handles partial matches (city matches region)
- [ ] Multiple filters applied simultaneously narrow results correctly
- [ ] No filters selected returns all artists
- [ ] Filter controls retain their values after page reload
- [ ] Invalid filter values handled gracefully
- [ ] Unit tests for `applyFilters` and `locationMatches`