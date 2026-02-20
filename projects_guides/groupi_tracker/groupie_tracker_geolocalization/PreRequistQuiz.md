# 🎯 Groupie Tracker Geolocalization Prerequisites Quiz
## Geocoding APIs · Coordinates · Map Embedding · Mutex Caching · JS Map Libraries

**Time Limit:** 50 minutes  
**Total Questions:** 27  
**Passing Score:** 21/27 (78%)

> ✅ Pass → You're ready to start Groupie Tracker Geolocalization  
> ⚠️ This project requires understanding TWO external APIs (geocoding + maps) and passing Go data to JavaScript. If you score 21–23, study Section 3 carefully before starting.

---

## 📋 SECTION 1: GEOCODING CONCEPTS (5 Questions)

### Q1: What is geocoding?

**A)** Encoding text in Go  
**B)** Converting a human-readable address or place name into geographic coordinates (latitude and longitude)  
**C)** Compressing geographic data  
**D)** Drawing maps in HTML  

<details><summary>💡 Answer</summary>

**B) Converting address/place name → latitude and longitude**

```
"Berlin, Germany"  →  lat: 52.5200, lng: 13.4050
"New York, USA"    →  lat: 40.7128, lng: -74.0060
```

The reverse (coordinates → address) is called **reverse geocoding**. This project uses forward geocoding: you have location strings from the API and need their coordinates to place map markers.

</details>

---

### Q2: The Groupie Tracker API returns a location like `"new_york-usa"`. What must you do to it before sending it to a geocoding API?

**A)** Nothing — all geocoding APIs accept any format  
**B)** Normalize it: replace `_` with space, replace `-` with `,` or space → `"New York, USA"` or `"new york usa"`  
**C)** URL-encode it only  
**D)** Convert it to coordinates manually  

<details><summary>💡 Answer</summary>

**B) Normalize: replace underscores with spaces, hyphens with commas**

```go
func normalizeForGeocoding(loc string) string {
    loc = strings.ReplaceAll(loc, "_", " ")
    loc = strings.ReplaceAll(loc, "-", ", ")
    return loc
}
// "new_york-usa" → "new york, usa"
// "saint_etienne-france" → "saint etienne, france"
```

Sending `"new_york-usa"` to Nominatim or similar will likely return no results or wrong results. Always normalize first.

</details>

---

### Q3: A geocoding API call costs you one HTTP request. One artist might have 8 concert locations. If your app has 52 artists, what is the maximum number of geocoding requests you might make at startup without caching?

**A)** 52  
**B)** Up to 52 × 8 = 416 (or however many unique locations there are)  
**C)** 1  
**D)** It's limited by Go automatically  

<details><summary>💡 Answer</summary>

**B) Up to ~52 × 8 = ~416 requests**

Without caching, you'd geocode every location for every artist — even locations shared across multiple artists (e.g. "London, UK" might appear for 20 different artists). A cache ensures each unique location is geocoded only once. This is both a performance and rate-limit concern — free geocoding APIs typically allow 1 request/second or 1000/day.

</details>

---

### Q4: What are the coordinates for latitude and longitude?

**A)** X and Y pixel positions on a screen  
**B)** Latitude: north-south position (-90 to +90, positive = north); Longitude: east-west position (-180 to +180, positive = east)  
**C)** Longitude: north-south; Latitude: east-west  
**D)** Both range from 0 to 360  

<details><summary>💡 Answer</summary>

**B) Lat: north-south (-90 to +90); Lng: east-west (-180 to +180)**

Common trap: some geocoding APIs return `[lng, lat]` (GeoJSON convention), others return `[lat, lng]`. Always check the API docs. Swapping them will place your markers in the wrong hemisphere.

Berlin example: `lat: 52.52, lng: 13.40` → Northern Europe ✓

</details>

---

### Q5: What HTTP status code from a geocoding API means "too many requests"?

**A)** 400  
**B)** 404  
**C)** 429  
**D)** 503  

<details><summary>💡 Answer</summary>

**C) 429 Too Many Requests**

If you hit a geocoding API's rate limit, it returns `429`. Your code must handle this — check `resp.StatusCode` and either wait and retry, return an error, or skip the location. Always check the API's rate limit before designing your caching strategy.

</details>

---

## 📋 SECTION 2: MAKING GEOCODING REQUESTS IN GO (6 Questions)

### Q6: You need to build a geocoding request URL with a query parameter. Which produces `https://nominatim.openstreetmap.org/search?q=berlin+germany&format=json`?

**A)** `"https://nominatim.openstreetmap.org/search?q=" + location + "&format=json"`  
**B)** Using `url.Values` to properly encode the query  
**C)** Both work, but B is safer for locations containing special characters (spaces, commas, ampersands)  
**D)** Neither works  

<details><summary>💡 Answer</summary>

**C) Both work, but B is safer for special characters**

```go
// Safe approach using url.Values:
params := url.Values{}
params.Set("q", normalizedLocation)
params.Set("format", "json")
reqURL := "https://nominatim.openstreetmap.org/search?" + params.Encode()
```

String concatenation breaks if the location contains `&`, `=`, or non-ASCII characters. `url.Values.Encode()` handles this correctly.

</details>

---

### Q7: The Nominatim geocoding API returns a JSON array. Each element has `"lat"` and `"lon"` as string fields:
```json
[{"lat": "52.5200", "lon": "13.4050", "display_name": "Berlin..."}]
```

What Go struct decodes this correctly?

**A)**
```go
type GeoResult struct {
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
}
```
**B)**
```go
type GeoResult struct {
    Lat string `json:"lat"`
    Lon string `json:"lon"`
}
```
**C)**
```go
type GeoResult struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}
```
**D)**
```go
type GeoResult struct {
    lat string
    lon string
}
```

<details><summary>💡 Answer</summary>

**B) `Lat string` and `Lon string`**

The JSON values are quoted — they are JSON strings, not JSON numbers. Using `float64` would fail because Go's JSON decoder is strict: a JSON string `"52.5200"` cannot decode into `float64`. After decoding, convert with `strconv.ParseFloat(result.Lat, 64)`.

Option D has unexported fields — `json.Unmarshal` cannot populate them.

</details>

---

### Q8: After decoding, what does `strconv.ParseFloat("52.5200", 64)` return?

**A)** `52` (integer)  
**B)** `52.52` (float64) and `nil` error  
**C)** `52.5200` (float64) and `nil` error  
**D)** Error — strings with trailing zeros are invalid  

<details><summary>💡 Answer</summary>

**C) `52.5200` (float64) and `nil` error**

`strconv.ParseFloat` converts the string to a 64-bit floating point number. `52.5200` and `52.52` represent the same float — the trailing zeros don't matter numerically. The `64` parameter specifies the precision (use 64 for `float64`).

</details>

---

### Q9: The geocoding API returns an empty array `[]` for a location it couldn't find. How should your `geocode` function handle this?

**A)** Return coordinates `{0, 0}` — that's the "null island" in the ocean  
**B)** Return an error: "location not found" — the caller can skip this location  
**C)** Panic  
**D)** Return coordinates `{-1, -1}`  

<details><summary>💡 Answer</summary>

**B) Return an error — let the caller decide what to do**

`{0, 0}` is a real point (off the coast of Africa in the Gulf of Guinea) — placing a marker there would be misleading. An error lets `geocodeAll` skip the location gracefully while still plotting the ones that did resolve.

```go
if len(results) == 0 {
    return Coordinates{}, fmt.Errorf("no results for location: %s", location)
}
```

</details>

---

### Q10: How should you add a `User-Agent` header to your Nominatim requests? (Nominatim requires it.)

**A)** Nominatim doesn't require headers  
**B)** Use `http.NewRequest` and `req.Header.Set("User-Agent", "MyApp/1.0")`  
**C)** Add it to the URL query string  
**D)** Set it globally for all requests  

<details><summary>💡 Answer</summary>

**B) Use `http.NewRequest` + `req.Header.Set` + `http.DefaultClient.Do(req)`**

```go
req, err := http.NewRequest("GET", reqURL, nil)
if err != nil { return Coordinates{}, err }
req.Header.Set("User-Agent", "groupie-tracker-app/1.0 (your@email.com)")

resp, err := http.DefaultClient.Do(req)
```

Nominatim's usage policy requires a descriptive `User-Agent`. Requests without one may be blocked. `http.Get` uses a generic user-agent by default.

</details>

---

### Q11: Geocoding the same location (e.g., "London, UK") for 15 different artists without a cache means 15 HTTP requests for the same data. Besides being wasteful, what is the other risk?

**A)** London moves  
**B)** Hitting the geocoding API's rate limit (typically 1 req/sec for free tiers) causing 429 errors  
**C)** Go can only make 10 outbound requests  
**D)** The coordinates change between requests  

<details><summary>💡 Answer</summary>

**B) Rate limit violations — 429 errors**

Free geocoding APIs (especially Nominatim) have strict rate limits. Making 15 requests for the same location hits the limit and gets subsequent requests rejected. A simple in-memory cache: `map[string]Coordinates` + `sync.Mutex` solves both problems — one request per unique location, ever.

</details>

---

## 📋 SECTION 3: MUTEX CACHE (5 Questions)

### Q12: You implement a geocoding cache as a global `map[string]Coordinates`. Multiple goroutines might call `geocodeCached` simultaneously. What problem can occur?

**A)** The cache grows too large  
**B)** Concurrent map writes in Go cause a panic — maps are not safe for concurrent access without synchronization  
**C)** The coordinates get overwritten with wrong values  
**D)** No problem — maps handle concurrency automatically  

<details><summary>💡 Answer</summary>

**B) Concurrent map writes cause a panic**

Go maps are not concurrent-safe. Simultaneous reads are fine; a write concurrent with anything causes a `concurrent map read and map write` fatal panic. Always protect a shared map with `sync.Mutex` when multiple goroutines might access it.

</details>

---

### Q13: What is the correct pattern for a thread-safe cache check + insert?

**A)**
```go
mu.Lock()
if v, ok := cache[key]; ok { return v, nil }
result, err := geocode(key)
cache[key] = result
mu.Unlock()
```
**B)**
```go
mu.Lock()
if v, ok := cache[key]; ok { mu.Unlock(); return v, nil }
mu.Unlock()
result, err := geocode(key)
mu.Lock()
cache[key] = result
mu.Unlock()
```
**C)**
```go
result, err := geocode(key)
mu.Lock()
cache[key] = result
mu.Unlock()
```
**D)** Both A and B are correct; B is better  

<details><summary>💡 Answer</summary>

**D) Both work; B is better**

Option A holds the lock for the entire geocoding HTTP call — blocking ALL other goroutines from reading the cache while one is making a network request. Option B: check cache (brief lock), release while doing I/O, re-lock only to write. This is more efficient. Note: option B has a subtle "double-check" issue (two goroutines might both miss the cache and both geocode the same location) — acceptable for this project since the result is the same.

</details>

---

### Q14: What does `sync.RWMutex` offer over `sync.Mutex` for a cache?

**A)** It's faster in all cases  
**B)** `RWMutex` allows multiple concurrent **reads** but exclusive **writes** — a cache with many reads and few writes benefits significantly  
**C)** It automatically persists the cache to disk  
**D)** It works without locking  

<details><summary>💡 Answer</summary>

**B) Multiple concurrent reads, exclusive writes**

```go
var mu sync.RWMutex

// Reading (many goroutines can do this simultaneously):
mu.RLock()
v, ok := cache[key]
mu.RUnlock()

// Writing (exclusive):
mu.Lock()
cache[key] = value
mu.Unlock()
```

For a geocoding cache that is written once per location and read many times thereafter, `RWMutex` eliminates unnecessary contention on reads.

</details>

---

### Q15: What is the cache key — the raw API location string or the normalized one?

**A)** Raw — to preserve the original data  
**B)** Normalized — so that "new_york-usa" and "new york, usa" (if both appeared) map to the same cache entry  
**C)** The coordinates themselves  
**D)** The artist ID  

<details><summary>💡 Answer</summary>

**B) Normalized**

Normalize the location string BEFORE looking up or storing in the cache. This ensures that any variation of the same location (different formatting from different parts of the API) hits the same cache entry.

```go
func geocodeCached(rawLoc string) (Coordinates, error) {
    key := normalizeForGeocoding(rawLoc)
    // check cache with `key`, geocode with `key`
}
```

</details>

---

### Q16: Your geocoding cache is an in-memory `map[string]Coordinates`. What happens to cached data when the server restarts?

**A)** It persists automatically  
**B)** It is lost — the cache starts empty on every startup  
**C)** It is saved to a file automatically  
**D)** Docker preserves it  

<details><summary>💡 Answer</summary>

**B) Lost on restart — in-memory only**

An in-memory cache is sufficient for this project — the server stays up and accumulates results as artists are viewed. For a production system you'd persist to Redis or a file. For this project, the cache is warm after the first round of artist page loads.

</details>

---

## 📋 SECTION 4: EMBEDDING MAPS & PASSING DATA TO JAVASCRIPT (7 Questions)

### Q17: You want to embed an interactive map on the artist detail page. Which approach works without a server-side library?

**A)** `<img src="map.png">` — static image  
**B)** A JavaScript map library (e.g. Leaflet.js) loaded via CDN, initialized with coordinates passed from Go template  
**C)** Go renders the map as SVG  
**D)** A `<map>` HTML element  

<details><summary>💡 Answer</summary>

**B) JavaScript map library via CDN + coordinates from Go template**

```html
<link rel="stylesheet" href="https://unpkg.com/leaflet@1.9/dist/leaflet.css">
<script src="https://unpkg.com/leaflet@1.9/dist/leaflet.js"></script>
<div id="map" style="height: 400px;"></div>
<script>
    var map = L.map('map').setView([0, 0], 2);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png').addTo(map);
    // Add markers from Go template data
</script>
```

Leaflet.js is free, open-source, and works with OpenStreetMap tiles. No API key required.

</details>

---

### Q18: You have a Go slice of `LocationCoordinate` structs. How do you make this data available to JavaScript in the HTML template?

**A)** You can't — Go and JavaScript can't share data directly  
**B)** Use `{{ range }}` to emit JavaScript code inside a `<script>` tag that builds an array  
**C)** Use a hidden `<div>` with the data as attributes  
**D)** Both B and C work  

<details><summary>💡 Answer</summary>

**D) Both approaches work; B (emitting JS array) is cleaner**

```html
<script>
var locations = [
    {{ range .Locations }}
    { lat: {{ .Lat }}, lng: {{ .Lng }}, name: "{{ .Name }}" },
    {{ end }}
];
</script>
```

Or use `json` encoding in the template to emit a JSON array directly:
```html
<script>
var locations = {{ .LocationsJSON }};  // pre-encoded JSON string from Go
</script>
```

The second approach is safer — the first can break if location names contain quotes.

</details>

---

### Q19: In Go, how do you convert a slice of structs to a JSON string to embed in a template?

**A)** `fmt.Sprintf("%v", slice)`  
**B)** `json.Marshal(slice)` → `[]byte` → `string(bytes)` → pass to template  
**C)** `string(slice)`  
**D)** Manually loop and build the JSON string  

<details><summary>💡 Answer</summary>

**B) `json.Marshal` → bytes → string → template**

```go
locJSON, err := json.Marshal(locationCoords)
if err != nil { /* handle */ }

data := ArtistPageData{
    // ...
    LocationsJSON: string(locJSON),
}
```

In the template: `var locations = {{ .LocationsJSON }};` — emits valid JSON directly. If using `html/template`, the JSON may be HTML-escaped — use `template.JS(locJSON)` to mark it as safe JavaScript.

</details>

---

### Q20: What does `template.JS(value)` do?

**A)** Executes JavaScript code  
**B)** Marks a string as safe JavaScript — prevents `html/template` from HTML-escaping characters like `<`, `>`, `&` that would break JSON embedded in a script tag  
**C)** Lints the JavaScript for errors  
**D)** Compresses the JavaScript  

<details><summary>💡 Answer</summary>

**B) Marks a string as safe JavaScript, preventing HTML escaping**

`html/template` escapes `<`, `>`, `&`, `"` to prevent XSS. This breaks JSON containing these characters. `template.JS(jsonString)` tells the template engine "this is already safe JavaScript, don't escape it."

```go
// In handler:
data.LocationsJS = template.JS(locJSON)
// In template:
var locs = {{ .LocationsJS }};  // emitted raw, not escaped
```

</details>

---

### Q21: Using Leaflet.js, how do you add a marker at Berlin (lat: 52.52, lng: 13.40) with a popup showing "Berlin, Germany"?

**A)** `L.point(52.52, 13.40).addTo(map)`  
**B)** `L.marker([52.52, 13.40]).addTo(map).bindPopup("Berlin, Germany")`  
**C)** `map.addMarker(52.52, 13.40, "Berlin, Germany")`  
**D)** `<marker lat="52.52" lng="13.40">`  

<details><summary>💡 Answer</summary>

**B) `L.marker([lat, lng]).addTo(map).bindPopup("text")`**

```javascript
// Leaflet expects [lat, lng] (note: latitude FIRST)
L.marker([52.52, 13.40])
 .addTo(map)
 .bindPopup("Berlin, Germany");
```

Common mistake: using `[lng, lat]` — this is GeoJSON convention (which Leaflet does NOT use for `L.marker`). Leaflet always expects `[latitude, longitude]`.

</details>

---

### Q22: What should the artist detail page show when ZERO locations were successfully geocoded?

**A)** A blank map centered on the ocean  
**B)** Hide the map entirely and show a message like "Map data unavailable"  
**C)** Show a map centered on the world with no markers — it's fine  
**D)** Return a 500 error  

<details><summary>💡 Answer</summary>

**B or C — both are acceptable; B is more informative**

A blank world map is confusing. If zero locations were geocoded, it's better UX to either hide the map and explain why, or show the map with a message. The server should NOT return 500 — the artist's other information is still valid and should display. Only the map portion is degraded.

</details>

---

### Q23: The geocoding API is completely unreachable (network error). What should your `artistHandler` do?

**A)** Return 500 Internal Server Error  
**B)** Return the artist page without the map (skip geocoding, show a message, still show all other artist info)  
**C)** Retry indefinitely until it succeeds  
**D)** Show a blank page  

<details><summary>💡 Answer</summary>

**B) Degrade gracefully — show the page without map data**

The artist's name, members, creation date, first album, and concert locations are all available from the main API. The map is an enhancement. If geocoding fails, show the page without it — never let an optional feature break the entire page.

```go
coords, err := geocodeAll(rel.DatesLocations)
if err != nil {
    // log the error, coords will be empty
    log.Printf("geocoding failed: %v", err)
}
// pass coords (possibly empty) to template — template handles empty gracefully
```

</details>

---

## 📋 SECTION 5: INTEGRATION (4 Questions)

### Q24: Where in your application should geocoding happen — at startup for all artists, or lazily when each artist's page is viewed?

**A)** At startup — geocode everything before the server starts  
**B)** Lazily — geocode only the locations for the artist currently being viewed, cache results  
**C)** In a background goroutine that runs forever  
**D)** In the template  

<details><summary>💡 Answer</summary>

**B) Lazily with caching**

Geocoding at startup means potentially hundreds of API requests before the server is ready — and most artists may never be viewed. Lazy geocoding means: first visit to an artist page triggers geocoding for that artist's locations, then the cache serves all future requests instantly. This is the standard pattern for expensive external API calls.

</details>

---

### Q25: Two users visit the same artist page simultaneously. Both trigger `geocodeCached("berlin-germany")` at the same time. Without proper locking, what could happen?

**A)** Nothing bad — they'd both get the correct result  
**B)** Both see a cache miss, both make HTTP requests to the geocoding API for the same location, both try to write to the cache concurrently — risking a map write panic  
**C)** One request is automatically queued  
**D)** The second request uses the result of the first  

<details><summary>💡 Answer</summary>

**B) Both cache miss, both make requests, both try to write — potential panic**

This is the "thundering herd" problem. For this project the risk is low (it only races on the first access) but the map write panic is real. The mutex pattern prevents it. For a production system, a "single-flight" pattern prevents duplicate requests for the same key.

</details>

---

### Q26: Your `ArtistPageData` struct now includes `LocationsJSON template.JS`. In your `artist.html` template, you write:
```html
<script>var locs = {{ .LocationsJSON }};</script>
```
But the JSON contains `"saint_etienne-france"` which becomes `"saint_etienne\u002dfrance"` when rendered. What causes this?

**A)** Normal — JSON always escapes hyphens  
**B)** You didn't use `template.JS` — `html/template` is escaping the content  
**C)** Leaflet doesn't support non-ASCII characters  
**D)** The hyphen is invalid JSON  

<details><summary>💡 Answer</summary>

**B) You're not using `template.JS`**

`html/template` Unicode-escapes content in script tags for XSS protection. `\u002d` is the Unicode escape for `-`. Wrapping with `template.JS(jsonBytes)` marks it as trusted, bypassing the escaping:

```go
data.LocationsJSON = template.JS(locJSON)  // in handler
```

Without this, the JSON is technically still valid (Unicode escapes are legal JSON), but it's confusing and breaks character-level matching in JavaScript.

</details>

---

### Q27: After adding the map feature, you run your test suite. The test for the artist handler now requires a working geocoding API. How do you avoid making real HTTP calls in tests?

**A)** Skip the test  
**B)** Extract the geocoding function as an interface or function parameter — in tests, pass a mock function that returns fake coordinates  
**C)** Use `time.Sleep` to wait for the API  
**D)** Tests don't need to cover geocoding  

<details><summary>💡 Answer</summary>

**B) Mock the geocoding function (dependency injection)**

```go
// Make geocode a function variable:
var geocodeFunc = geocode  // default to real implementation

func geocodeCached(loc string) (Coordinates, error) {
    return geocodeFunc(loc)
}

// In tests:
geocodeFunc = func(loc string) (Coordinates, error) {
    return Coordinates{Lat: 52.52, Lng: 13.40}, nil
}
```

This keeps tests fast, deterministic, and independent of external services.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 25–27 ✅ | **Excellent.** Start Groupie Tracker Geolocalization. |
| 21–24 ✅ | **Ready.** Study the section you scored lowest on before starting. |
| 16–20 ⚠️ | **Study first.** Practice making HTTP requests to external APIs and passing data from Go to JavaScript. |
| Below 16 ❌ | **Not ready.** This project combines many moving parts — make sure HTTP client code and Go templates are solid first. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q5 | Geocoding concepts, normalization, rate limits, coordinate systems |
| Q6–Q11 | `url.Values`, JSON string fields, `strconv.ParseFloat`, error handling, User-Agent |
| Q12–Q16 | `sync.Mutex`, `sync.RWMutex`, cache patterns, normalized cache keys |
| Q17–Q23 | Leaflet.js basics, `json.Marshal`, `template.JS`, graceful degradation |
| Q24–Q27 | Lazy loading pattern, thundering herd, `template.JS` escaping, mocking for tests |