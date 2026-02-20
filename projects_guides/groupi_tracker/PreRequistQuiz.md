# 🎯 Groupie Tracker Prerequisites Quiz
## HTTP Client · JSON Decoding · Struct Tags · Data Relationships · API Consumption

**Time Limit:** 50 minutes  
**Total Questions:** 28  
**Passing Score:** 22/28 (78%)

> ✅ Pass → You're ready to start Groupie Tracker  
> ⚠️ Also Required → ASCII-Art-Web must be complete — you already know `net/http` server and `html/template`

---

## 📋 SECTION 1: MAKING HTTP REQUESTS AS A CLIENT (6 Questions)

### Q1: What is the difference between Go's `net/http` package used as a **server** vs as a **client**?

**A)** They use completely different packages  
**B)** As a server you call `http.ListenAndServe` and write `http.ResponseWriter`; as a client you call `http.Get` to fetch data from an external URL  
**C)** The client mode is slower  
**D)** Client mode requires a separate import  

<details><summary>💡 Answer</summary>

**B) Server writes responses; client makes requests to external URLs**

```go
// As a server (you've done this):
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)

// As a client (new for this project):
resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
```

The same `net/http` package does both. This project adds the client side.

</details>

---

### Q2: What is the correct way to make a GET request to an external URL and read the response body?

**A)**
```go
body := http.Get("https://api.example.com/data")
```
**B)**
```go
resp, err := http.Get("https://api.example.com/data")
if err != nil { return err }
defer resp.Body.Close()
data, err := io.ReadAll(resp.Body)
```
**C)**
```go
resp := http.Fetch("https://api.example.com/data")
return resp.JSON()
```
**D)**
```go
resp, _ := http.Get("https://api.example.com/data")
return resp.Body
```

<details><summary>💡 Answer</summary>

**B)**

Three critical points:
1. `http.Get` returns `(*Response, error)` — always check the error
2. `defer resp.Body.Close()` — always close the body to avoid resource leaks
3. `io.ReadAll(resp.Body)` — read the body bytes before doing anything with them

Option D ignores the error AND doesn't close the body. Option A won't compile.

</details>

---

### Q3: You make a request to the API but get HTTP status 500. Your `err` variable is `nil`. What does this mean?

**A)** The request succeeded  
**B)** The HTTP transport succeeded (no network error) but the server returned a 500 error — you must check `resp.StatusCode`, not just `err`  
**C)** Go automatically retries on 500  
**D)** `err` would never be nil on a 500  

<details><summary>💡 Answer</summary>

**B) You must check `resp.StatusCode` separately from `err`**

`err != nil` means the network/transport failed (DNS, connection refused, timeout). A 4xx or 5xx response is NOT a Go error — the HTTP transport worked fine. Always check both:

```go
resp, err := http.Get(url)
if err != nil { return err }
defer resp.Body.Close()
if resp.StatusCode != 200 {
    return fmt.Errorf("API returned status %d", resp.StatusCode)
}
```

</details>

---

### Q4: Why must you call `resp.Body.Close()` after reading from it?

**A)** To save the data  
**B)** To release the underlying network connection back to the connection pool — not closing leaks connections and eventually exhausts available connections  
**C)** To prevent the response from being cached  
**D)** It's optional — Go GC handles it  

<details><summary>💡 Answer</summary>

**B) To release the network connection back to the pool**

HTTP keep-alive reuses connections. If you don't close the body, the connection is never returned and you'll eventually run out. `defer resp.Body.Close()` placed immediately after checking `err` is the idiomatic pattern — it runs even if subsequent code panics.

</details>

---

### Q5: The API returns a JSON array. You call `http.Get` and get the response. What is the most efficient way to decode the JSON directly from the response without reading all bytes into memory first?

**A)** `json.Unmarshal(resp.Body, &data)`  
**B)** `json.NewDecoder(resp.Body).Decode(&data)`  
**C)** `ioutil.ReadAll(resp.Body)` then `json.Unmarshal`  
**D)** `resp.Body.JSON(&data)`  

<details><summary>💡 Answer</summary>

**B) `json.NewDecoder(resp.Body).Decode(&data)`**

`json.NewDecoder` reads directly from the `io.Reader` stream without buffering everything in memory first. For large API responses this is more efficient. `json.Unmarshal` requires the full bytes upfront. Both work correctly — `NewDecoder` is the idiomatic choice for HTTP responses.

</details>

---

### Q6: Should you fetch the API data on every request to your server, or once at startup?

**A)** Every request — to get fresh data  
**B)** Once at startup — the artist data is static; re-fetching on every page load is wasteful and slow  
**C)** Every 5 minutes using a ticker  
**D)** Never — hardcode the data  

<details><summary>💡 Answer</summary>

**B) Once at startup**

The Groupie Tracker API contains static band/artist data that doesn't change. Fetching it once at startup and storing it in memory means: fast page loads (no outbound HTTP per request), no rate-limiting risk, and the server works even if the API is temporarily unreachable after startup.

</details>

---

## 📋 SECTION 2: JSON DECODING & STRUCT TAGS (8 Questions)

### Q7: What is a JSON struct tag in Go and why is it needed?

**A)** A comment explaining the field  
**B)** A backtick annotation that maps a Go struct field name to its JSON key name — needed because Go uses PascalCase but JSON typically uses camelCase or snake_case  
**C)** A validation rule  
**D)** A default value  

<details><summary>💡 Answer</summary>

**B) A backtick annotation mapping Go field name to JSON key**

```go
type Artist struct {
    ID           int      `json:"id"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Image        string   `json:"image"`
}
```

Without the tag, Go would look for a JSON key `"ID"` (exact match) and silently leave the field as zero value if not found. Tags ensure correct mapping.

</details>

---

### Q8: What is the output?
```go
type Artist struct {
    Name string `json:"name"`
}

data := []byte(`{"name": "Queen"}`)
var a Artist
json.Unmarshal(data, &a)
fmt.Println(a.Name)
```

**A)** `"Queen"` (with quotes)  
**B)** `Queen`  
**C)** empty string  
**D)** Error  

<details><summary>💡 Answer</summary>

**B) `Queen`**

`json.Unmarshal` correctly decodes the JSON string `"Queen"` into the Go `string` field `Name`. The quotes are part of JSON syntax — the resulting Go string doesn't have them.

</details>

---

### Q9: What happens if you define a struct field but the JSON key doesn't exist in the response?

**A)** `json.Unmarshal` returns an error  
**B)** The field gets its zero value — `0` for int, `""` for string, `nil` for slice  
**C)** The program panics  
**D)** The entire struct is left empty  

<details><summary>💡 Answer</summary>

**B) The field gets its zero value — silent, no error**

This is a common debugging trap. If your struct tag doesn't exactly match the JSON key (case-sensitive), the field is silently left at zero. Always print the decoded struct immediately after decoding to verify all fields are populated.

</details>

---

### Q10: The API returns locations as an array of objects, each with an `"index"` field containing a nested array of location strings. What Go struct represents this?

```json
{
  "index": [
    {
      "id": 1,
      "locations": ["saint_etienne-france", "seattle-usa"]
    }
  ]
}
```

**A)**
```go
type LocationsResponse struct {
    Index []struct {
        ID        int      `json:"id"`
        Locations []string `json:"locations"`
    } `json:"index"`
}
```
**B)**
```go
type LocationsResponse struct {
    Locations string `json:"locations"`
}
```
**C)**
```go
type LocationsResponse struct {
    Index []string `json:"index"`
}
```
**D)**
```go
type LocationsResponse map[string][]string
```

<details><summary>💡 Answer</summary>

**A)**

Nested JSON objects require nested Go structs (or anonymous structs inline). The outer `"index"` key maps to a slice of objects. Each object has `"id"` (int) and `"locations"` ([]string). The struct tags must match the JSON keys exactly.

</details>

---

### Q11: The `relation` API endpoint returns:
```json
{
  "index": [
    {
      "id": 1,
      "datesLocations": {
        "berlin-germany": ["25-06-2019"],
        "london-uk": ["01-07-2019", "02-07-2019"]
      }
    }
  ]
}
```

What Go type should `DatesLocations` be?

**A)** `[]string`  
**B)** `map[string]string`  
**C)** `map[string][]string`  
**D)** `[][]string`  

<details><summary>💡 Answer</summary>

**C) `map[string][]string`**

The keys are location strings (e.g. `"berlin-germany"`). Each value is an array of date strings. In Go: `map[string][]string`. The struct tag would be `` `json:"datesLocations"` ``.

This is the critical data structure in the Groupie Tracker project — getting this type right unlocks the entire relation endpoint.

</details>

---

### Q12: You decode the artists successfully but `CreationDate` is always `0`. The JSON has `"creationDate": 1970`. What is the most likely cause?

**A)** Integers can't be decoded from JSON  
**B)** The struct tag is wrong — it might say `json:"creation_date"` or `json:"CreationDate"` instead of `json:"creationDate"`  
**C)** `int` is the wrong type — use `int64`  
**D)** JSON decoding is case-insensitive so it always works  

<details><summary>💡 Answer</summary>

**B) The struct tag doesn't exactly match the JSON key**

JSON decoding IS case-insensitive for the default case (no tag), but struct tags are matched exactly. `json:"creation_date"` won't match `"creationDate"`. Print the raw JSON before decoding and compare every key character by character with your struct tags.

</details>

---

### Q13: `json.Unmarshal` vs `json.NewDecoder().Decode()` — which requires `[]byte` and which requires an `io.Reader`?

**A)** Both take `[]byte`  
**B)** `json.Unmarshal` takes `[]byte`; `json.NewDecoder` takes an `io.Reader`  
**C)** Both take `io.Reader`  
**D)** `json.Unmarshal` takes `string`  

<details><summary>💡 Answer</summary>

**B) `Unmarshal` = `[]byte`; `NewDecoder` = `io.Reader`**

```go
// Unmarshal — needs bytes in memory
data, _ := io.ReadAll(resp.Body)
json.Unmarshal(data, &result)

// NewDecoder — streams directly from the body
json.NewDecoder(resp.Body).Decode(&result)
```

`resp.Body` is an `io.Reader`, so `NewDecoder` is the direct fit. If you need to inspect the raw bytes (for debugging), read them first with `io.ReadAll`, then decode.

</details>

---

### Q14: How do you check if `json.Unmarshal` or `json.NewDecoder().Decode()` failed?

**A)** Check if the struct is empty  
**B)** Both return an `error` — check `if err != nil`  
**C)** They never fail  
**D)** Use `recover()` to catch panics  

<details><summary>💡 Answer</summary>

**B) Both return `error` — check `if err != nil`**

```go
if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
    return nil, fmt.Errorf("failed to decode artists: %w", err)
}
```

A decode error means: the JSON is malformed, or the target type doesn't match. Always check — silently ignoring decode errors leads to empty structs and mysterious bugs.

</details>

---

## 📋 SECTION 3: DATA RELATIONSHIPS (5 Questions)

### Q15: The API has separate endpoints for artists, locations, dates, and relations. How do they connect to each other?

**A)** Each response contains all the data for that artist  
**B)** Each artist and relation entry has an `id` field — you match them by comparing `artist.ID` with `relation.ID`  
**C)** They connect via the order they appear (first artist matches first location)  
**D)** The artist name is the key  

<details><summary>💡 Answer</summary>

**B) Match via `id` field**

```go
// Find the relation for a specific artist:
for _, rel := range relations {
    if rel.ID == artist.ID {
        // this is the matching relation data
    }
}
```

Always match by ID — never by position. The order of items in different endpoints is not guaranteed to be the same.

</details>

---

### Q16: For the artist detail page, you need to show concert locations and their dates. Which endpoint provides this combined data?

**A)** `/api/artists` — it includes all concert data  
**B)** `/api/locations` and `/api/dates` — you combine them manually  
**C)** `/api/relation` — it directly links locations to dates for each artist  
**D)** `/api/artists/1` — detail endpoints include all data  

<details><summary>💡 Answer</summary>

**C) `/api/relation`**

The `relation` endpoint returns `datesLocations: map[string][]string` — exactly what you need: each location mapped to its list of dates for that artist. This is the most useful endpoint for the detail page.

</details>

---

### Q17: You want to display the detail page for artist with ID 5. How do you find the artist in your pre-loaded slice?

**A)** `artists[5]` — index directly  
**B)** Loop through the slice and compare `artist.ID == 5`  
**C)** `artists.Find(5)`  
**D)** The API always returns them sorted so index 4 is always ID 5  

<details><summary>💡 Answer</summary>

**B) Loop and compare `artist.ID == 5`**

```go
var found *Artist
for i, a := range artists {
    if a.ID == id {
        found = &artists[i]
        break
    }
}
if found == nil {
    http.NotFound(w, r)
    return
}
```

Never assume the slice is sorted or that index maps to ID. IDs can have gaps. Always search by value.

</details>

---

### Q18: A user visits `/artist?id=abc`. How should your handler respond?

**A)** Return 404 — artist not found  
**B)** Return 400 — `"abc"` is not a valid integer ID  
**C)** Return 500 — internal error  
**D)** Return 200 with an empty page  

<details><summary>💡 Answer</summary>

**B) Return 400 — invalid input**

`"abc"` is a client error — they sent a non-integer where an integer is expected. Use `strconv.Atoi` to parse and check:

```go
id, err := strconv.Atoi(r.URL.Query().Get("id"))
if err != nil {
    http.Error(w, "Invalid artist ID", http.StatusBadRequest)
    return
}
```

`400` = bad request from client. `404` = valid ID that doesn't exist. Get the distinction right.

</details>

---

### Q19: How do you read a query parameter from the URL `"/artist?id=5"` in Go?

**A)** `r.FormValue("id")`  
**B)** `r.URL.Query().Get("id")`  
**C)** `r.QueryParams["id"]`  
**D)** Both A and B work  

<details><summary>💡 Answer</summary>

**D) Both A and B work — but understand the difference**

`r.FormValue("id")` reads from form POST data AND URL query parameters. `r.URL.Query().Get("id")` reads only from the URL query string. For a GET request with query parameters, both return the same thing. `r.URL.Query()` is more explicit and preferred for query-param-only reads.

</details>

---

## 📋 SECTION 4: HTTP SERVER & TEMPLATES (5 Questions)

### Q20: Your home handler is registered as `http.HandleFunc("/", homeHandler)`. A user visits `/favicon.ico`. What happens without extra code?

**A)** Returns 404 automatically  
**B)** `homeHandler` is called — you must explicitly check `r.URL.Path` and return 404 for unknown paths  
**C)** Go serves it from the current directory automatically  
**D)** The browser doesn't request `/favicon.ico`  

<details><summary>💡 Answer</summary>

**B) `homeHandler` is called — you must check `r.URL.Path`**

The `"/"` pattern in Go's default mux matches ALL unregistered paths. Add this to every handler registered on `"/"`:

```go
if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
}
```

</details>

---

### Q21: You want to pass both the artists slice AND the selected artist's relation data to a template. How do you do this?

**A)** Call `tmpl.Execute` twice with different data  
**B)** Define a struct that holds all the data, pass that struct to `Execute`  
**C)** Use global variables in the template  
**D)** You can only pass one value to a template  

<details><summary>💡 Answer</summary>

**B) Define a struct holding all the data, pass the struct**

```go
type ArtistPageData struct {
    Artist   Artist
    Relation Relation
    Error    string
}

data := ArtistPageData{Artist: found, Relation: rel}
tmpl.Execute(w, data)
```

In the template: `{{ .Artist.Name }}`, `{{ .Relation.DatesLocations }}`. This is the standard pattern for any page with multiple data sources.

</details>

---

### Q22: In your `index.html` template, how do you iterate over the artists slice?

**A)** `{{ for artist in .Artists }}`  
**B)** `{{ range .Artists }}{{ .Name }}{{ end }}`  
**C)** `{{ each .Artists as artist }}{{ artist.Name }}{{ end }}`  
**D)** `{{ loop .Artists }}`  

<details><summary>💡 Answer</summary>

**B) `{{ range .Artists }}{{ .Name }}{{ end }}`**

Inside `{{ range }}`, the `.` changes to refer to the current element. So `.Name` refers to the artist's name. To access the parent data inside range, define a variable: `{{ range $i, $a := .Artists }}{{ $a.Name }}{{ end }}`.

</details>

---

### Q23: How do you iterate over a `map[string][]string` (the `datesLocations` data) in a Go template?

**A)** `{{ range .DatesLocations }}{{ . }}{{ end }}`  
**B)** `{{ range $location, $dates := .DatesLocations }}{{ $location }}: {{ range $dates }}{{ . }}{{ end }}{{ end }}`  
**C)** `{{ map .DatesLocations }}`  
**D)** Maps can't be used in templates  

<details><summary>💡 Answer</summary>

**B) `{{ range $key, $value := .Map }}`**

```html
{{ range $location, $dates := .DatesLocations }}
    <h3>{{ $location }}</h3>
    <ul>
        {{ range $dates }}
        <li>{{ . }}</li>
        {{ end }}
    </ul>
{{ end }}
```

Go templates support two-variable range for maps: `$key, $value`. The order is not guaranteed — maps in Go templates are iterated in random order.

</details>

---

### Q24: The API is unreachable when your server starts. What should happen?

**A)** Start the server anyway with empty data  
**B)** Log the error and exit — the app cannot function without the data  
**C)** Retry automatically in the background  
**D)** Return 503 on all requests forever  

<details><summary>💡 Answer</summary>

**B) Log the error and exit**

```go
artists, err := fetchArtists()
if err != nil {
    log.Fatalf("Failed to fetch artist data: %v", err)
}
```

Starting with empty data would serve a broken page to every user. Exiting with a clear error message is better — the operator knows to investigate. `log.Fatalf` logs and calls `os.Exit(1)`.

</details>

---

## 📋 SECTION 5: CLIENT-SERVER EVENT (4 Questions)

### Q25: The spec requires at least one "client-server event." What does this mean?

**A)** A user can click a button to reload the page  
**B)** A user action triggers a new request to the server, and the page updates based on the server's response  
**C)** The server sends push notifications to the client  
**D)** A form with a submit button  

<details><summary>💡 Answer</summary>

**B) User action → server request → page update**

Examples that qualify: clicking a location on the detail page to see all artists who played there, a "related artists" button, a "show tour history" button. The key is that the event triggers a go-to-server-and-get-data cycle, not just front-end DOM manipulation.

</details>

---

### Q26: You implement a `/location?name=berlin` endpoint that returns all artists who played in Berlin. What should it return — a full HTML page or JSON?

**A)** Must be JSON  
**B)** Must be a full HTML page  
**C)** Either can work — a full page redirect is simpler; JSON with JavaScript is more dynamic. The spec doesn't require AJAX.  
**D)** It must use WebSockets  

<details><summary>💡 Answer</summary>

**C) Either approach works**

The simplest implementation: a link navigates to `/location?name=berlin` which renders a full HTML page. This requires no JavaScript and is easier to implement correctly. JSON + JavaScript is more impressive but adds complexity. For this project, a full-page response is perfectly acceptable.

</details>

---

### Q27: Your event handler at `/location?name=berlin-germany` must find all artists who have a concert in Berlin. How do you check if an artist has a location that matches?

**A)** `if artist.Location == "berlin-germany"`  
**B)** Loop through the artist's `DatesLocations` map keys and check if any key equals or contains the search string  
**C)** Use `strings.Contains` on a concatenated string of all locations  
**D)** The API provides a pre-built endpoint for this  

<details><summary>💡 Answer</summary>

**B) Loop through `DatesLocations` map keys**

```go
for location := range relation.DatesLocations {
    if strings.Contains(location, searchTerm) {
        // This artist has a concert in the searched location
        matched = append(matched, artist)
        break
    }
}
```

The relation data's `DatesLocations` map has location strings as keys. Loop through them and check for a match.

</details>

---

### Q28: What is the minimum test coverage the spec requires?

**A)** No tests required  
**B)** Unit tests for at least the data fetching and decoding logic  
**C)** Full integration tests  
**D)** Only end-to-end tests  

<details><summary>💡 Answer</summary>

**B) Unit tests for data fetching and decoding logic**

At minimum: test that your fetch functions return the correct types, test that your JSON decoding works with a sample JSON fixture, and test that your lookup/matching logic returns correct results. Write these alongside the code, not as an afterthought.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 26–28 ✅ | **Excellent.** Strong API and JSON foundations — start immediately. |
| 22–25 ✅ | **Ready.** Review missed questions, especially struct tags and response body handling. |
| 17–21 ⚠️ | **Study first.** JSON decoding and HTTP client patterns need more work. |
| Below 17 ❌ | **Not ready.** Review `encoding/json`, `http.Get`, and struct tags before starting. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q6 | `http.Get`, `resp.Body.Close()`, checking `StatusCode`, `json.NewDecoder` |
| Q7–Q14 | Struct tags, `json.Unmarshal`, zero values on decode fail, `map[string][]string` |
| Q15–Q19 | Matching by ID, the relation endpoint, `strconv.Atoi`, query params |
| Q20–Q24 | `/` catch-all handler, template structs, `range` over maps in templates |
| Q25–Q28 | Client-server events, location matching, test requirements |