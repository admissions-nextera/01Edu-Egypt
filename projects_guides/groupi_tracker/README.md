# Groupie Tracker Project Guide

> **Before you start:** Open https://groupietrackers.herokuapp.com/api in your browser. Read every endpoint. Then open https://rickandmortyapi.com/ and explore how a well-structured API works. You are building a website that consumes real data — understand the data first.

---

## Objectives

By completing this project you will learn:

1. **Consuming a REST API** — Making HTTP requests to an external API and reading the response
2. **JSON Parsing** — Decoding JSON into Go structs
3. **Data Relationships** — Linking data across multiple API endpoints into a unified model
4. **HTTP Server** — Serving an HTML website from a Go backend
5. **HTML Templates** — Rendering dynamic data from Go into HTML pages
6. **Client-Server Events** — Implementing a feature where a user action triggers a server request and updates the page
7. **Error Handling** — Keeping the server alive and informative under all conditions

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Go HTTP Client
- `http.Get` — making a GET request to an external URL
- Reading and closing `resp.Body`
- Search: **"golang http GET request external API"**

### 2. JSON in Go
- `json.Unmarshal` or `json.NewDecoder` — decoding JSON into a struct
- How to define Go structs that match a JSON structure
- What struct tags like `json:"name"` do
- Search: **"golang JSON decode struct tags"**
- https://pkg.go.dev/encoding/json

### 3. HTTP Server and Templates
- `net/http` — `HandleFunc`, `ListenAndServe`
- `html/template` — `ParseFiles`, `Execute`
- HTTP status codes: 200, 404, 400, 500

### 4. HTML Basics
- Cards, lists, tables — ways to display structured data
- `<a href>` for navigation between pages

**Read before starting:**
- https://groupietrackers.herokuapp.com/api — explore every endpoint
- Search: **"golang JSON decode tutorial"**
- https://pkg.go.dev/encoding/json

---

## Project Structure

```
groupie-tracker/
├── main.go
├── handlers.go       ← HTTP handler functions
├── api.go            ← API fetching and data structs
├── templates/
│   ├── index.html    ← artist listing page
│   └── artist.html   ← single artist detail page
├── static/
│   └── style.css
└── go.mod
```

---

## Milestone 1 — Understand the API

**This milestone has no code.**

Visit each of these URLs and read the JSON response:
- `https://groupietrackers.herokuapp.com/api`
- `https://groupietrackers.herokuapp.com/api/artists`
- `https://groupietrackers.herokuapp.com/api/locations`
- `https://groupietrackers.herokuapp.com/api/dates`
- `https://groupietrackers.herokuapp.com/api/relation`

**Questions to answer before writing anything:**
- What fields does each artist object contain?
- What does the `relation` endpoint add that the others do not?
- How are locations and dates linked to a specific artist?
- What Go struct would you write to represent one artist? What types do its fields need?
- What is the JSON key for the artist's image URL? What about their members?

Sketch your Go structs on paper before writing any code.

---

## Milestone 2 — Fetch and Decode the API Data

**Goal:** Your program fetches the artists from the API and prints their names to the terminal. No server yet.

**Questions to answer:**
- How do you make a GET request to a URL in Go?
- How do you read the entire response body?
- How do you decode JSON into a slice of structs?
- What should happen if the API request fails?

**Code Placeholder:**
```go
// api.go

type Artist struct {
    // Define fields that match the API JSON response
    // Use struct tags: `json:"fieldName"`
    // What types are: ID, Name, Members, CreationDate, FirstAlbum, Image?
}

type Location struct {
    // Define fields matching the locations API response
}

type Dates struct {
    // Define fields matching the dates API response
}

type Relation struct {
    // Define fields matching the relation API response
    // The locations-dates link is a map — what type?
}

func fetchArtists() ([]Artist, error) {
    // Make a GET request to the artists endpoint
    // Decode the JSON response into []Artist
    // Return the slice and any error
}

func fetchRelations() ([]Relation, error) {
    // Same pattern for the relations endpoint
}
```

**Resources:**
- Search: **"golang json.NewDecoder http response"**
- Search: **"golang struct json tags example"**

**Verify:** `go run .` (no server yet) fetches and prints all artist names without crashing.

---

## Milestone 3 — Start the Server and Serve the Artist List

**Goal:** `http://localhost:8080` shows a page listing all artists with their name and image.

**Questions to answer:**
- When should you fetch the API data — once at startup or on every request?
- What are the trade-offs of each approach?
- How do you pass a slice of artists to an HTML template?
- What HTTP status code do you return if the API fetch fails at startup?

**Code Placeholder:**
```go
// main.go

func main() {
    // Fetch all data from the API at startup
    // If fetching fails, log the error and exit

    // Register routes
    // Start the server on port 8080
}
```

```go
// handlers.go

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Check method is GET and path is exactly "/"
    // Return 404 for unknown paths
    // Return 400 for wrong methods

    // Parse and execute the index template
    // Pass the artists slice to the template
    // Return 500 if template fails
}
```

**Verify:** The page loads and shows all artist names and images.

---

## Milestone 4 — Artist Detail Page

**Goal:** Clicking an artist shows a detail page with all their information: name, image, members, creation date, first album, and concert locations with dates.

**Questions to answer:**
- How do you pass an artist ID through a URL? (e.g. `/artist?id=1`)
- How do you read a query parameter from `r.URL.Query()`?
- How do you find the relation data for a specific artist?
- What should happen if someone requests `/artist?id=999` and that ID does not exist?

**Code Placeholder:**
```go
// handlers.go

func artistHandler(w http.ResponseWriter, r *http.Request) {
    // Check method is GET

    // Read "id" from query parameters
    // Validate: return 400 if id is missing or not a valid integer

    // Find the artist with that ID in your loaded data
    // Return 404 if not found

    // Find the matching relation data for this artist

    // Build a combined struct with artist + relation data
    // Parse and execute the artist template
    // Return 500 if template fails
}
```

**Verify:**
- Click an artist — the detail page loads correctly
- Visit `/artist?id=999` — returns 404
- Visit `/artist?id=abc` — returns 400

---

## Milestone 5 — Client-Server Event

**Goal:** Implement one feature where a user action triggers a request to the server and the page updates with new information. This is the required "event" in the spec.

**Questions to answer:**
- What feature will you build? Ideas: clicking a location shows a list of artists who played there, clicking a member shows all artists they belong to, a "show more" button that fetches more details.
- What new endpoint will this feature call?
- What data does the server need to receive from the client to respond correctly?
- What does the server return — a full HTML page or just a piece of data?

**Code Placeholder:**
```go
// handlers.go

func yourEventHandler(w http.ResponseWriter, r *http.Request) {
    // Read the input from the request (query param, form value, or path)
    // Validate it
    // Look up the relevant data
    // Return the result — either a full template or a partial response
}
```

**Verify:** The event triggers when expected, the server responds, and the page updates. It does not crash under any input.

---

## Milestone 6 — Error Handling and Stability

**Goal:** The server never crashes. Every error state shows a meaningful page.

**Questions to answer:**
- What happens if a template file is missing?
- What happens if the API is unreachable at startup?
- What happens if someone navigates to a URL your server does not handle?
- How do you serve a custom 404 page instead of Go's default?

**Verify:**
- Delete a template file temporarily — server returns 500, does not crash
- Visit any random URL — returns 404 with a readable message
- Submit any form with unexpected values — server handles it gracefully

---

## Debugging Checklist

- Does JSON decoding silently fail? Print the raw response body before decoding to see what you are actually receiving.
- Are struct fields empty after decoding? Check your JSON struct tags — they must exactly match the JSON key names (case-sensitive).
- Does artist data appear but location data not? The `relation` endpoint links them — make sure you are fetching and matching it correctly.
- Does the server crash on a bad request? Every handler must validate its inputs before touching any data.
- Does your home handler serve every path? Add `if r.URL.Path != "/" { http.NotFound(w, r); return }` at the top.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `net/http` | Server, client requests | https://pkg.go.dev/net/http |
| `encoding/json` | Decode API responses | https://pkg.go.dev/encoding/json |
| `html/template` | Render HTML with data | https://pkg.go.dev/html/template |
| `strconv` | Parse artist ID from query string | https://pkg.go.dev/strconv |
| `fmt` | Format error messages | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] All four API endpoints fetched and decoded correctly
- [ ] Home page displays all artists
- [ ] Artist detail page shows all fields including concert locations and dates
- [ ] At least one client-server event implemented
- [ ] GET `/` returns 200
- [ ] Unknown paths return 404
- [ ] Wrong methods return 400
- [ ] Server errors return 500
- [ ] Server never crashes on any input
- [ ] All pages render correctly with no template errors
- [ ] Unit tests written for at least the data fetching and decoding logic