# 🎯 Groupie Tracker Geolocalization Prerequisites Quiz
## External APIs · Geocoding · JSON Parsing · sync.Mutex Caching · Leaflet.js · JavaScript Maps

**Time Limit:** 55 minutes  
**Total Questions:** 30  
**Passing Score:** 24/30 (80%)

> Questions are tagged: 🟢 Easy · 🟡 Medium · 🔴 Hard  
> All topics are general — no specific project knowledge required.

---

## 📋 SECTION 1: CALLING EXTERNAL APIs (8 Questions)

### Q1 🟢 — What is geocoding?

**A)** Encoding geographic data as binary  
**B)** Converting a human-readable address or location name into geographic coordinates (latitude and longitude)  
**C)** Drawing maps in SVG  
**D)** Compressing GPS data  

<details><summary>💡 Answer</summary>

**B) Converting address/name → latitude + longitude**

```
"New York, USA"      → { lat: 40.7128, lon: -74.0060 }
"Paris, France"      → { lat: 48.8566, lon:   2.3522 }
"Tokyo, Japan"       → { lat: 35.6762, lon: 139.6503 }
```

Geocoding is performed by APIs (Google Maps, OpenStreetMap Nominatim, etc.) that maintain massive databases of place names and their coordinates. Reverse geocoding does the opposite: coordinates → human-readable address.

</details>

---

### Q2 🟢 — The Nominatim geocoding API returns JSON. What Go struct would you use to decode this response?

```json
[{"lat": "48.8566", "lon": "2.3522", "display_name": "Paris, France"}]
```

**A)**
```go
type Result struct {
    Lat  float64 `json:"lat"`
    Lon  float64 `json:"lon"`
}
```
**B)**
```go
type Result struct {
    Lat         string `json:"lat"`
    Lon         string `json:"lon"`
    DisplayName string `json:"display_name"`
}
var results []Result
```
**C)**
```go
type Result struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}
```
**D)**
```go
var results map[string]string
```

<details><summary>💡 Answer</summary>

**B) String fields — the JSON values are quoted strings, not numbers**

```go
type GeoResult struct {
    Lat         string `json:"lat"`
    Lon         string `json:"lon"`
    DisplayName string `json:"display_name"`
}

var results []GeoResult
json.NewDecoder(resp.Body).Decode(&results)

// Then convert to float64 for calculations:
lat, err := strconv.ParseFloat(results[0].Lat, 64)
lon, err := strconv.ParseFloat(results[0].Lon, 64)
```

Nominatim returns coordinates as JSON **strings** (`"48.8566"`), not numbers (`48.8566`). If you use `float64` fields, the JSON decoder silently leaves them as `0.0`. Always look at the raw API response before writing your struct.

</details>

---

### Q3 🟢 — How do you add a custom header to an outgoing HTTP request in Go?

**A)** `http.Get(url, headers)`  
**B)** Build a request with `http.NewRequest`, call `req.Header.Set(key, value)`, then use a client to `Do(req)`  
**C)** `http.SetHeader(key, value)` globally  
**D)** Pass headers as query parameters  

<details><summary>💡 Answer</summary>

**B) `http.NewRequest` → `req.Header.Set` → `client.Do(req)`**

```go
req, err := http.NewRequest("GET", geocodeURL, nil)
if err != nil { return err }

// Many geocoding APIs require a User-Agent header:
req.Header.Set("User-Agent", "MyApp/1.0 (myemail@example.com)")
req.Header.Set("Accept", "application/json")

client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Do(req)
if err != nil { return err }
defer resp.Body.Close()
```

Nominatim specifically requires a `User-Agent` header that identifies your app — requests without it may be blocked. Always set a timeout to avoid goroutines hanging on slow external services.

</details>

---

### Q4 🟡 — How do you add query parameters to a URL programmatically in Go?

**A)** String concatenation: `url + "?q=" + query`  
**B)** Use `url.Values` to build and encode parameters safely  
**C)** `http.AddParams(url, params)`  
**D)** Put parameters in the request body  

<details><summary>💡 Answer</summary>

**B) `url.Values` — handles URL encoding automatically**

```go
import "net/url"

base := "https://nominatim.openstreetmap.org/search"
params := url.Values{}
params.Set("q", "New York, USA")   // automatically URL-encodes spaces, special chars
params.Set("format", "json")
params.Set("limit", "1")

fullURL := base + "?" + params.Encode()
// "https://nominatim.openstreetmap.org/search?format=json&limit=1&q=New+York%2C+USA"
```

String concatenation breaks for inputs with spaces, commas, or other special characters. `url.Values.Encode()` handles percent-encoding correctly.

</details>

---

### Q5 🟡 — A geocoding API returns HTTP 429. What does this mean and how should you handle it?

**A)** 429 = Internal Server Error — retry immediately  
**B)** 429 = Too Many Requests — you've hit the rate limit; back off and retry after a delay  
**C)** 429 = Not Found — the location doesn't exist  
**D)** 429 = Authentication required  

<details><summary>💡 Answer</summary>

**B) 429 = Too Many Requests — implement rate limiting and backoff**

```go
if resp.StatusCode == http.StatusTooManyRequests {
    // Check Retry-After header if present:
    retryAfter := resp.Header.Get("Retry-After")
    delay := 1 * time.Second
    if retryAfter != "" {
        if seconds, err := strconv.Atoi(retryAfter); err == nil {
            delay = time.Duration(seconds) * time.Second
        }
    }
    time.Sleep(delay)
    // retry the request
}
```

Free geocoding APIs often have strict rate limits (e.g., 1 request/second). Caching geocoding results locally is the best way to stay under limits.

</details>

---

### Q6 🟡 — What is the correct way to normalize a location string like `"new_york-usa"` before geocoding?

**A)** Send it as-is — APIs handle any format  
**B)** Replace underscores with spaces and hyphens with commas: `"new york, usa"` — then geocode  
**C)** Convert to uppercase  
**D)** Remove all punctuation  

<details><summary>💡 Answer</summary>

**B) Replace `_` with space and `-` with `, ` — then geocode**

```go
func normalizeLocation(loc string) string {
    s := strings.ReplaceAll(loc, "_", " ")  // "new york-usa"
    s = strings.ReplaceAll(s, "-", ", ")    // "new york, usa"
    return s
}

// "london-uk"       → "london, uk"
// "new_york-usa"    → "new york, usa"
// "paris-france"    → "paris, france"
```

API response location keys often use URL-safe encodings. Geocoding APIs work best with natural language: "city, country" format. Normalizing first dramatically improves match quality.

</details>

---

### Q7 🔴 — What does `resp.Body` contain when a geocoding API returns an empty array `[]`?

**A)** `nil`  
**B)** The literal bytes `[]` — parsing succeeds but the resulting Go slice is empty (length 0), meaning the location was not found  
**C)** An error is returned by `http.Get`  
**D)** The response body is skipped  

<details><summary>💡 Answer</summary>

**B) `[]` in body — parsing succeeds, slice is empty — check length**

```go
var results []GeoResult
if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
    return err
}

if len(results) == 0 {
    return fmt.Errorf("location not found: %s", query)
}

// Safe to use results[0] now:
lat, _ := strconv.ParseFloat(results[0].Lat, 64)
```

An empty JSON array `[]` is valid JSON. The decoder succeeds but produces an empty slice. Always check `len(results) > 0` before accessing `results[0]` — panic on index out of range is one of the most common bugs.

</details>

---

### Q8 🔴 — What is the risk of making one geocoding API request per page load, and how do you mitigate it?

**A)** No risk — geocoding APIs are always fast  
**B)** Each request adds network latency, risks rate-limiting, and costs money on paid APIs — mitigate with an in-memory cache keyed by normalized location string  
**C)** Risk: the coordinates change over time  
**D)** Risk: the request blocks the UI  

<details><summary>💡 Answer</summary>

**B) Latency + rate limits + cost — mitigate with in-memory caching**

```go
var (
    geoCache = make(map[string][2]float64)  // location → [lat, lon]
    geoCacheMu sync.RWMutex
)

func geocode(location string) ([2]float64, error) {
    key := normalizeLocation(location)

    geoCacheMu.RLock()
    if coords, ok := geoCache[key]; ok {
        geoCacheMu.RUnlock()
        return coords, nil  // cache hit — no API call
    }
    geoCacheMu.RUnlock()

    coords, err := callGeoAPI(key)  // slow network call
    if err != nil { return [2]float64{}, err }

    geoCacheMu.Lock()
    geoCache[key] = coords
    geoCacheMu.Unlock()

    return coords, nil
}
```

Cache geocoding results in memory for the lifetime of the server process — coordinates don't change.

</details>

---

## 📋 SECTION 2: CACHING WITH sync.Mutex (6 Questions)

### Q9 🟢 — Why must a cache shared between goroutines be protected by a mutex?

**A)** Caches are read-only, so no protection needed  
**B)** Multiple goroutines may read and write the cache map simultaneously — concurrent map access without synchronization causes a panic ("concurrent map read and map write")  
**C)** Mutex makes the cache faster  
**D)** Only needed if goroutines run on different CPU cores  

<details><summary>💡 Answer</summary>

**B) Concurrent map access without sync panics in Go**

```go
// Go 1.6+ detects concurrent map writes and panics at runtime:
// "fatal error: concurrent map read and map write"

// CORRECT — protect the map:
var (
    cache = make(map[string]Result)
    mu    sync.RWMutex
)

// Read:  mu.RLock() / mu.RUnlock()
// Write: mu.Lock()  / mu.Unlock()
```

Even `map` reads concurrent with writes are unsafe — the panic fires on any concurrent access where at least one is a write. This is not a "sometimes happens" bug — it will reliably crash your server under load.

</details>

---

### Q10 🟡 — What is the "check-then-act" pattern for a cache lookup?

**A)** Check the cache; act on the result  
**B)** Check under a read lock; if missing, upgrade to a write lock and check AGAIN before fetching — avoids duplicate fetches from concurrent misses  
**C)** Check without a lock; only lock when writing  
**D)** Always use a write lock for both reads and writes  

<details><summary>💡 Answer</summary>

**B) Read lock → miss → write lock → check again (double-check locking)**

```go
mu.RLock()
if v, ok := cache[key]; ok {
    mu.RUnlock()
    return v, nil  // cache hit
}
mu.RUnlock()

// Another goroutine may have added the key between the two locks:
mu.Lock()
if v, ok := cache[key]; ok {  // CHECK AGAIN under write lock
    mu.Unlock()
    return v, nil  // another goroutine already fetched it
}
v, err := fetch(key)  // only fetch if still missing
if err == nil {
    cache[key] = v
}
mu.Unlock()
return v, err
```

Without the second check, two goroutines that both miss the read lock would both call `fetch(key)` — the "thundering herd". The double-check prevents duplicate fetches.

</details>

---

### Q11 🟡 — What happens if you call `mu.Lock()` while already holding `mu.Lock()`?

**A)** The second lock is ignored  
**B)** The goroutine deadlocks — `sync.Mutex` is not reentrant; it blocks waiting for itself forever  
**C)** A panic occurs immediately  
**D)** The second lock is queued  

<details><summary>💡 Answer</summary>

**B) Deadlock — mutex is not reentrant**

```go
func processData() {
    mu.Lock()
    defer mu.Unlock()
    // ...
    saveToCache(result)  // BUG if saveToCache also calls mu.Lock()
}

func saveToCache(v Result) {
    mu.Lock()         // goroutine already holds mu → deadlock
    defer mu.Unlock()
    cache["key"] = v
}
```

Fix: either don't call locking functions from within a lock, or create an internal "unsafe" version that assumes the lock is held:

```go
// Internal — caller must hold mu:
func saveToCacheUnsafe(v Result) { cache["key"] = v }
```

</details>

---

### Q12 🟡 — Why use `sync.RWMutex` instead of `sync.Mutex` for a read-heavy cache?

**A)** `RWMutex` is always faster  
**B)** `RWMutex.RLock()` allows multiple concurrent readers — in a cache that is read far more often than written, this significantly reduces contention  
**C)** `RWMutex` doesn't require `Unlock`  
**D)** `RWMutex` handles network calls  

<details><summary>💡 Answer</summary>

**B) Multiple concurrent readers — less contention for read-heavy caches**

```go
// sync.Mutex — only 1 goroutine at a time, even for reads:
mu.Lock()   // all other goroutines blocked — even readers
v := cache[key]
mu.Unlock()

// sync.RWMutex — multiple readers simultaneously:
mu.RLock()  // other RLock() holders proceed concurrently
v := cache[key]
mu.RUnlock()
// Only mu.Lock() blocks all access

// Result for a cache hit 99% of the time:
// Mutex:   goroutines queue even for reads → slow
// RWMutex: goroutines read in parallel → fast
```

Use `RWMutex` when you have many reads and few writes — the classic cache pattern.

</details>

---

### Q13 🔴 — What is a normalized cache key and why does it matter?

**A)** A key that's been encrypted  
**B)** A standardized form of the key that treats equivalent inputs as identical — prevents duplicate cache entries for the same location written differently  
**C)** A key sorted alphabetically  
**D)** A key validated against a schema  

<details><summary>💡 Answer</summary>

**B) Standardized key — equivalent inputs map to the same entry**

```go
func cacheKey(location string) string {
    s := strings.ToLower(location)        // "Paris" == "paris"
    s = strings.TrimSpace(s)              // " paris " == "paris"
    s = strings.ReplaceAll(s, "_", " ")   // "new_york" == "new york"
    s = strings.ReplaceAll(s, "-", ", ")  // "paris-france" == "paris, france"
    return s
}

// Without normalization:
cache["Paris, France"] = coords
cache["paris, france"] = coords  // duplicate! 2 geocode calls for same location

// With normalization:
key := cacheKey("Paris, France")   // "paris, france"
key2 := cacheKey("paris,france")   // "paris, france" — same key, one entry
```

Poor cache keys waste memory and network calls. Normalize before looking up and before storing.

</details>

---

### Q14 🔴 — You cache geocoding results in a `map` at server startup. Is this safe to read from request handlers without a mutex?

**A)** Yes — the map is only written once before any goroutine reads it  
**B)** Depends on whether writes are fully complete before serving requests — if the map is populated before `http.ListenAndServe` is called and never written again, it is safe (read-only shared state requires no mutex)  
**C)** No — maps always require a mutex, even read-only  
**D)** Only safe if the map is declared as `const`  

<details><summary>💡 Answer</summary>

**B) Safe if all writes complete before any goroutine can read**

```go
// SAFE: populate cache in main before starting server:
cache := buildCache()    // synchronous — all writes done here
// No goroutines yet — no races possible
http.HandleFunc("/", handler)  // handler will only read cache
http.ListenAndServe(":8080", nil)  // goroutines start here

// UNSAFE: write to cache from a goroutine after server starts:
go func() {
    cache["key"] = geocode("location")  // concurrent write — race!
}()
http.ListenAndServe(":8080", nil)
```

Go's memory model guarantees: writes before goroutine creation are visible to the goroutine. A read-only map shared across goroutines after initialization requires no mutex.

</details>

---

## 📋 SECTION 3: LEAFLET.JS AND MAPS IN THE BROWSER (8 Questions)

### Q15 🟢 — What is Leaflet.js?

**A)** A Go package for maps  
**B)** A lightweight JavaScript library for interactive maps — renders tile-based maps in the browser, supports markers, popups, and layers  
**C)** A CSS framework  
**D)** A geocoding API  

<details><summary>💡 Answer</summary>

**B) A JavaScript library for interactive browser maps**

```html
<!-- Include via CDN: -->
<link rel="stylesheet" href="https://unpkg.com/leaflet/dist/leaflet.css">
<script src="https://unpkg.com/leaflet/dist/leaflet.js"></script>

<!-- Container div: -->
<div id="map" style="height: 400px;"></div>

<script>
const map = L.map('map').setView([48.8566, 2.3522], 13);  // Paris, zoom 13
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap'
}).addTo(map);
</script>
```

Leaflet uses OpenStreetMap tiles by default — free and open source. The `setView([lat, lon], zoom)` call sets the initial center and zoom level.

</details>

---

### Q16 🟢 — How do you add a marker to a Leaflet map?

**A)** `map.addPin([lat, lon])`  
**B)** `L.marker([lat, lon]).addTo(map)`  
**C)** `map.createMarker(lat, lon)`  
**D)** `L.point(lat, lon).addTo(map)`  

<details><summary>💡 Answer</summary>

**B) `L.marker([lat, lon]).addTo(map)`**

```javascript
// Basic marker:
L.marker([48.8566, 2.3522]).addTo(map);

// Marker with popup:
L.marker([48.8566, 2.3522])
  .addTo(map)
  .bindPopup("<b>Paris</b><br>City of Light")
  .openPopup();

// Marker with click handler:
L.marker([lat, lon])
  .addTo(map)
  .on('click', function() {
      console.log('marker clicked');
  });
```

`L.marker` takes `[latitude, longitude]` — latitude first. The popup content can include HTML.

</details>

---

### Q17 🟡 — What coordinate order does Leaflet use for `L.marker`?

**A)** `[longitude, latitude]` — same as GeoJSON  
**B)** `[latitude, longitude]` — Leaflet always uses lat first  
**C)** `[x, y]` — pixel coordinates  
**D)** Any order — Leaflet detects automatically  

<details><summary>💡 Answer</summary>

**B) `[latitude, longitude]` — Leaflet uses lat first, opposite of GeoJSON**

```javascript
// Leaflet: [lat, lon]
L.marker([48.8566, 2.3522])  // Paris: lat=48.8566, lon=2.3522

// GeoJSON: [lon, lat] (longitude FIRST — the opposite!)
{
    "type": "Point",
    "coordinates": [2.3522, 48.8566]  // GeoJSON: lon first
}
```

This is one of the most common map bugs. Swapping lat and lon places your marker in the middle of an ocean. Always double-check the coordinate order for each library you use.

</details>

---

### Q18 🟡 — How do you pass Go data (a slice of coordinates) to JavaScript in an HTML template safely?

**A)** `{{ .Coords }}`  
**B)** `json.Marshal` the data in Go, then use `template.JS()` to inject it into a `<script>` tag without HTML escaping  
**C)** Use `fmt.Sprintf` to build JavaScript  
**D)** Store in a cookie  

<details><summary>💡 Answer</summary>

**B) `json.Marshal` + `template.JS()` to prevent double-escaping**

```go
type Marker struct {
    Lat  float64
    Lon  float64
    Name string
}

markers := []Marker{{48.85, 2.35, "Paris"}, {51.51, -0.12, "London"}}
jsonData, _ := json.Marshal(markers)

data := PageData{
    MarkersJSON: template.JS(jsonData),  // marks as safe JavaScript
}
```

```html
<script>
const markers = {{ .MarkersJSON }};  // outputs raw JSON, not HTML-escaped
markers.forEach(m => {
    L.marker([m.Lat, m.Lon]).addTo(map).bindPopup(m.Name);
});
</script>
```

Without `template.JS()`, `html/template` would escape `<`, `>`, `&` and `"` in the JSON, breaking the JavaScript. `template.JS` tells the template engine this value is safe JavaScript that shouldn't be escaped.

</details>

---

### Q19 🟡 — What does `L.map('map').fitBounds(bounds)` do?

**A)** Creates a map  
**B)** Adjusts the map's center and zoom level to fit all the given coordinates in the visible area  
**C)** Draws a bounding box  
**D)** Limits the map to a specific region  

<details><summary>💡 Answer</summary>

**B) Auto-adjusts view to show all markers**

```javascript
const markers = [
    L.marker([48.85, 2.35]).addTo(map),   // Paris
    L.marker([51.51, -0.12]).addTo(map),  // London
    L.marker([52.52, 13.40]).addTo(map),  // Berlin
];

// Create a bounds object from all marker positions:
const group = L.featureGroup(markers);
map.fitBounds(group.getBounds().pad(0.1));  // 0.1 = 10% padding
```

`fitBounds` is essential when you have a variable number of markers spread across a map — hard-coding a center and zoom level would never work for all sets of locations.

</details>

---

### Q20 🔴 — What is `template.JS` type and when should you NOT use it?

**A)** A type alias for `string` — always use it for JavaScript  
**B)** A type that marks a value as safe JavaScript — it prevents HTML escaping. Only use it for data you control (Go-generated JSON); NEVER use it for user-supplied strings — it bypasses XSS protection  
**C)** A function that validates JavaScript syntax  
**D)** Required for all `<script>` tags  

<details><summary>💡 Answer</summary>

**B) `template.JS` bypasses escaping — only for trusted, Go-generated data**

```go
// SAFE — data comes from your own Go struct:
jsonData, _ := json.Marshal(myStruct)
safe := template.JS(jsonData)

// DANGEROUS — user input bypasses XSS protection:
userInput := r.FormValue("name")
unsafe := template.JS(userInput)  // if user enters: "}; alert('xss');//
// This injects JavaScript into your page!
```

Only use `template.JS` for data you construct yourself in Go. User input must always go through `{{ .UserInput }}` (normal template action) which HTML-escapes it properly.

</details>

---

### Q21 🔴 — What is "lazy geocoding" and why is it preferred over geocoding everything at startup?

**A)** A geocoding library that is slow  
**B)** Geocoding only when a location is first requested — avoids startup delay, avoids rate-limiting for locations never actually viewed, and spreads the load over time  
**C)** Caching results to disk  
**D)** Using GPS instead of an API  

<details><summary>💡 Answer</summary>

**B) Geocode on first use — not all at startup**

```go
// EAGER (at startup) — geocodes ALL locations before serving any request:
func init() {
    for _, loc := range allLocations {
        coords, _ := geocode(loc)  // if you have 500 locations, 500 API calls at startup!
        cache[loc] = coords
    }
}

// LAZY (on first request) — only geocode when needed:
func getCoords(loc string) ([2]float64, error) {
    if coords, ok := cache[loc]; ok { return coords, nil }
    coords, err := geocode(loc)  // called only if this location was never requested
    if err == nil { cache[loc] = coords }
    return coords, err
}
```

Lazy geocoding: faster startup, no wasted API calls for unused locations, lower rate-limit risk. The tradeoff: first request for a location is slower.

</details>

---

### Q22 🔴 — How do you handle a failed geocoding call gracefully?

**A)** Panic — a map without all markers is useless  
**B)** Log the error, skip the failed location, and render the page with the successfully geocoded locations — don't block the whole page for one failure  
**C)** Return an HTTP 500 error  
**D)** Retry indefinitely until it succeeds  

<details><summary>💡 Answer</summary>

**B) Graceful degradation — show partial data, don't block the page**

```go
type MarkerData struct {
    Lat  float64
    Lon  float64
    Name string
}

var markers []MarkerData
for _, location := range locations {
    coords, err := geocodeWithCache(location)
    if err != nil {
        log.Printf("geocoding %s failed: %v (skipping)", location, err)
        continue  // skip this location, continue with others
    }
    markers = append(markers, MarkerData{coords[0], coords[1], location})
}

// Render page with whatever markers we have:
tmpl.Execute(w, PageData{Markers: markers})
```

Graceful degradation: if 1 out of 20 locations fails to geocode, show 19 markers rather than breaking the whole page. Log the failure for debugging.

</details>

---

## 📋 SECTION 4: PASSING DATA FROM GO TO JAVASCRIPT (8 Questions)

### Q23 🟢 — How does JavaScript receive data from a Go HTML template?

**A)** Through HTTP cookies  
**B)** Via AJAX calls to a JSON endpoint, or embedded directly into the HTML using template actions (`{{ . }}`) in `<script>` tags  
**C)** Via WebSockets only  
**D)** JavaScript can't access Go data  

<details><summary>💡 Answer</summary>

**B) Embedded in HTML via template actions, or via AJAX to a JSON endpoint**

```html
<!-- Method 1: embedded JSON in script tag (simpler, synchronous) -->
<script>
const allMarkers = {{ .MarkersJSON }};
</script>

<!-- Method 2: fetch from a JSON endpoint (better for large data) -->
<script>
fetch('/api/markers')
  .then(r => r.json())
  .then(data => {
      data.forEach(m => L.marker([m.Lat, m.Lon]).addTo(map));
  });
</script>
```

For map data that's ready at page load time, embedding JSON directly is simpler. For data that changes dynamically (user interactions), use fetch/AJAX.

</details>

---

### Q24 🟡 — Why does `json.Marshal` with struct fields produce camelCase JSON keys in JavaScript?

**A)** Go automatically converts PascalCase to camelCase  
**B)** It doesn't — Go uses the exact field name (PascalCase) unless you add `json:"camelCase"` struct tags  
**C)** JavaScript requires camelCase  
**D)** `json.Marshal` always lowercases keys  

<details><summary>💡 Answer</summary>

**B) Without tags, Go uses exact field names — add tags for JavaScript conventions**

```go
// Without tags — JavaScript gets PascalCase:
type Marker struct { Lat float64; Lon float64 }
json.Marshal(Marker{48.85, 2.35})
// {"Lat":48.85,"Lon":2.35}  — JavaScript: m.Lat, m.Lon

// With tags — JavaScript gets camelCase (common convention):
type Marker struct {
    Lat  float64 `json:"lat"`
    Lon  float64 `json:"lon"`
    Name string  `json:"name"`
}
json.Marshal(Marker{48.85, 2.35, "Paris"})
// {"lat":48.85,"lon":2.35,"name":"Paris"}  — JavaScript: m.lat, m.lon
```

Convention: Go structs use PascalCase fields; JSON APIs use camelCase or snake_case keys. Use struct tags to control the JSON output format.

</details>

---

### Q25 🟡 — What does `html/template` do to `{{ .Value }}` when the value contains `<script>` or `"`?

**A)** Nothing — outputs the raw value  
**B)** HTML-escapes it: `<` → `&lt;`, `>` → `&gt;`, `"` → `&#34;` — prevents XSS attacks  
**C)** Removes the value and logs a warning  
**D)** Panics  

<details><summary>💡 Answer</summary>

**B) HTML-escapes all special characters — XSS prevention**

```go
// If .Name == `<script>alert('xss')</script>`:
// Template: <p>{{ .Name }}</p>
// Output:   <p>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</p>
// Browser shows the literal text, doesn't execute the script

// This is why template.JS() must only be used for your own data:
// Template: <script>const name = "{{ .Name }}"</script>  // STILL wrong — breaks JS context
// Use:      <script>const name = {{ .NameJSON }};</script> // where NameJSON is template.JS
```

`html/template` is context-aware — it knows when it's inside an HTML attribute, a script tag, or a URL, and escapes appropriately for each context.

</details>

---

### Q26 🟡 — How do you iterate over Go map data in JavaScript after receiving it as JSON?

**A)** `for (key in data) { ... }`  
**B)** `Object.entries(data).forEach(([key, value]) => { ... })`  
**C)** `data.forEach(entry => { ... })`  
**D)** Both A and B work; B is the modern approach  

<details><summary>💡 Answer</summary>

**D) Both A and B work — B is more modern and avoids prototype issues**

```javascript
// JSON: {"london-uk": ["2023-07-01", "2023-07-02"], "paris-fr": ["2023-08-10"]}
const concerts = {{ .ConcertsJSON }};

// A) for...in (older style — works but can include inherited properties):
for (const location in concerts) {
    if (concerts.hasOwnProperty(location)) {
        const dates = concerts[location];
    }
}

// B) Object.entries (modern — cleaner, no prototype issues):
Object.entries(concerts).forEach(([location, dates]) => {
    L.marker(coordsFor(location)).addTo(map).bindPopup(
        `<b>${location}</b><br>${dates.join('<br>')}`
    );
});
```

`Object.entries` returns `[key, value]` pairs as an array — works the same for JSON objects decoded from Go maps.

</details>

---

### Q27 🔴 — What is the risk of embedding raw Go struct data in a `<script>` tag without `template.JS`?

**A)** No risk — Go data is always safe  
**B)** `html/template` HTML-escapes `<`, `>`, `&`, and `"` — so `{"key":"value"}` becomes `{&#34;key&#34;:&#34;value&#34;}`, which is invalid JavaScript  
**C)** The page crashes  
**D)** The browser ignores the script  

<details><summary>💡 Answer</summary>

**B) HTML escaping breaks JavaScript syntax**

```go
// Without template.JS:
type Data struct { Name string }
d := Data{Name: "it's <great>"}
jsonBytes, _ := json.Marshal(d)
// jsonBytes: {"Name":"it's \u003cgreat\u003e"}
// (json.Marshal escapes < and > as \u003c and \u003e — actually fine)

// The real problem is with raw struct output:
// Template: <script>const x = {{ .Data }}</script>
// Go outputs: &amp;{it&#39;s &lt;great&gt;}  ← not JSON, not JavaScript

// Correct:
jsonBytes, _ := json.Marshal(d)
pageData.DataJSON = template.JS(jsonBytes)
// Template: <script>const x = {{ .DataJSON }};</script>
// Outputs valid JSON: {"Name":"it's \u003cgreat\u003e"}
```

Always marshal to JSON and use `template.JS` for data embedded in `<script>` tags.

</details>

---

### Q28 🔴 — What is the difference between `L.tileLayer` and `L.marker` in Leaflet?

**A)** No difference  
**B)** `L.tileLayer` adds the background map image (from OpenStreetMap tiles); `L.marker` adds an interactive pin at specific coordinates — a map needs a tile layer to be visible  
**C)** `L.tileLayer` renders coordinates; `L.marker` renders images  
**D)** `L.tileLayer` is for mobile; `L.marker` is for desktop  

<details><summary>💡 Answer</summary>

**B) TileLayer = background map; Marker = interactive pin**

```javascript
// Without tileLayer — just a grey box:
const map = L.map('map').setView([51.5, -0.1], 13);
L.marker([51.5, -0.1]).addTo(map);  // marker visible but no map background

// With tileLayer — complete map:
const map = L.map('map').setView([51.5, -0.1], 13);
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '© OpenStreetMap'
}).addTo(map);  // adds map background
L.marker([51.5, -0.1]).addTo(map);
```

A tile layer is essential for context. The `{z}/{x}/{y}` placeholders are filled by Leaflet with the current zoom level and tile coordinates when fetching map tiles.

</details>

---

### Q29 🔴 — How do you handle a map that needs to display markers loaded asynchronously (after page load)?

**A)** Reload the page when data is ready  
**B)** Fetch the data via AJAX, then programmatically add markers using the Leaflet API after the data arrives  
**C)** Pre-load all possible markers at startup  
**D)** Use server-sent events  

<details><summary>💡 Answer</summary>

**B) AJAX fetch then programmatically add markers**

```javascript
const map = L.map('map').setView([0, 0], 2);
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png').addTo(map);

// Load markers after page renders:
fetch('/api/locations')
  .then(r => r.json())
  .then(locations => {
      const markers = locations.map(loc =>
          L.marker([loc.lat, loc.lon]).bindPopup(loc.name).addTo(map)
      );
      // Fit map to show all markers:
      if (markers.length > 0) {
          map.fitBounds(L.featureGroup(markers).getBounds().pad(0.1));
      }
  })
  .catch(err => console.error('failed to load locations:', err));
```

This pattern allows the page to load immediately while map data fetches in the background. Show a loading indicator while fetching.

</details>

---

### Q30 🔴 — What is the minimum Leaflet setup needed to show a world map with one marker at London?

**A)**
```html
<script>L.map('map').setView([51.5, -0.09], 13);</script>
```
**B)**
```html
<link rel="stylesheet" href="leaflet.css">
<div id="map" style="height:400px"></div>
<script src="leaflet.js"></script>
<script>
const map = L.map('map').setView([51.505, -0.09], 13);
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png').addTo(map);
L.marker([51.505, -0.09]).addTo(map).bindPopup('London').openPopup();
</script>
```
**C)**
```html
<script>
const map = new Map('map', 51.505, -0.09, 13);
map.addMarker('London');
</script>
```
**D)** A is correct but missing the marker  

<details><summary>💡 Answer</summary>

**B) — CSS, container div with height, JS library, tileLayer, then marker**

The four requirements for any Leaflet map:
1. **Leaflet CSS** — without it, the map controls look broken
2. **Container div with a height** — Leaflet can't render into a zero-height element
3. **Leaflet JS** — the library
4. **A tileLayer** — the background map; without it you get a grey box

Missing any one of these is the most common "why isn't my map showing?" problem.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 28–30 ✅ | **Exceptional** — geocoding, caching, and maps mastered. |
| 24–27 ✅ | **Ready** — review any missed sections before starting. |
| 18–23 ⚠️ | **Study first** — external APIs + JavaScript integration need more attention. |
| Below 18 ❌ | **Not ready** — review the `net/http` client, `sync` package, and complete the Leaflet.js quickstart. |

---

## 🔍 Review Map

| Missed | Topic to Study |
|---|---|
| Q1–Q8 | Geocoding concepts, Nominatim JSON (string coords!), `url.Values`, 429 handling, location normalization |
| Q9–Q14 | `sync.RWMutex` for caches, double-check locking, non-reentrant mutex, `RLock` vs `Lock`, normalized cache keys |
| Q15–Q22 | Leaflet setup, `L.marker([lat,lon])`, lat-first order, `template.JS`, `fitBounds`, graceful degradation |
| Q23–Q30 | Embedding Go data in JavaScript, struct tags for JSON keys, HTML escaping in scripts, Leaflet tileLayer + marker setup |