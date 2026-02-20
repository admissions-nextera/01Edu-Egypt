# Groupie Tracker Geolocalization Project Guide

> **Before you start:** This project builds on groupie-tracker. Open Google Maps and search for "Germany Mainz". Look at the coordinates in the URL. That number pair — latitude and longitude — is what this project is about. Understand what geocoding is before writing any code.

---

## Objectives

By completing this project you will learn:

1. **Geocoding** — Converting a human-readable address into geographic coordinates (latitude, longitude)
2. **Map APIs** — Embedding an interactive map and placing markers on it
3. **External API Integration** — Making requests to a third-party geocoding service from Go
4. **Coordinate Systems** — Understanding latitude, longitude, and how map markers work
5. **Per-Artist View** — Showing all concert locations of one specific artist on a map

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Groupie Tracker (Completed)
- Artist detail page working
- Relation data (locations per artist) accessible

### 2. What Geocoding Is
- The difference between an address and coordinates
- What a geocoding API does
- Search: **"what is geocoding explained"**
- Search: **"geocoding API free options"**

### 3. Making HTTP Requests from Go to External APIs
- `http.Get` with a constructed URL
- Reading and decoding a JSON response
- Handling API keys securely

### 4. Map Embedding
- How to embed a map in an HTML page
- What a map marker is and what it needs (coordinates, label)
- Search: **"embed interactive map HTML JavaScript"**

**Read before starting:**
- https://rapidapi.com/blog/top-map-apis/ — choose one geocoding and one map API
- Your chosen APIs' documentation — read the request format and response format completely before writing any code

---

## Project Structure

```
groupie-tracker-geolocalization/
├── main.go
├── handlers.go
├── api.go
├── geo.go            ← geocoding logic
├── templates/
│   ├── index.html
│   └── artist.html   ← now includes an embedded map
├── static/
│   └── style.css
└── go.mod
```

---

## Milestone 1 — Choose Your APIs

**This milestone has no code.**

You need two things: a geocoding API and a map display solution. Research your options and answer these questions before writing anything:

**For geocoding (address → coordinates):**
- What free geocoding APIs are available? (Options: OpenStreetMap Nominatim, Geocode.xyz, Positionstack, others)
- Does your chosen API require an API key?
- What does a request look like? What does the response JSON look like?
- What rate limits apply? Will it handle all the artist locations without hitting the limit?

**For map display:**
- Will you use a JavaScript map library (Leaflet.js, Google Maps, OpenLayers)?
- How do you add a marker at a specific latitude/longitude?
- How do you add a label or popup to a marker?
- Does your choice require an API key?

Write down the exact request URL format and the JSON path to the coordinates in the response before moving on.

---

## Milestone 2 — Geocode One Address

**Goal:** Given a string like `"germany-mainz"`, your program returns the coordinates `{49.0, 8.27}`.

**Questions to answer:**
- The API location strings use hyphens and underscores: `"new_york-usa"`. How do you convert this to a readable address for the geocoding API?
- What Go struct represents a coordinate pair?
- What should your function return if the geocoding API cannot find the address?
- What should happen if the geocoding API is unreachable?

**Code Placeholder:**
```go
// geo.go

type Coordinates struct {
    // Latitude  float64
    // Longitude float64
}

func geocode(location string) (Coordinates, error) {
    // 1. Normalize the location string
    //    Replace _ with space, replace - with , etc.
    //    Make it readable for the geocoding API

    // 2. Build the request URL
    //    Include the normalized location as a query parameter
    //    Include API key if required

    // 3. Make the GET request

    // 4. Decode the JSON response
    //    Extract latitude and longitude from the correct field path

    // 5. Return the Coordinates struct
}
```

**Verify:** Call `geocode("germany-mainz")` and print the result. Check it against Google Maps.

---

## Milestone 3 — Geocode All Locations for One Artist

**Goal:** Given an artist's relation data, return coordinates for all their concert locations.

**Questions to answer:**
- How many locations can a single artist have?
- Should you geocode all locations every time the page loads, or cache the results?
- What is the performance impact of making one HTTP request per location sequentially? What would be faster?

**Code Placeholder:**
```go
// geo.go

type LocationCoordinate struct {
    // Location name (original string)
    // Coordinates
}

func geocodeAll(locations []string) []LocationCoordinate {
    // For each location, call geocode()
    // Collect results — skip locations that fail
    // Return all successfully geocoded locations

    // Optional: do this concurrently with goroutines for speed
}
```

**Verify:** For an artist with 5+ concert locations, all are geocoded and printed correctly.

---

## Milestone 4 — Cache Geocoding Results

**Goal:** Geocoding results are stored so the same address is never looked up twice.

**Questions to answer:**
- Where do you store the cache — in memory (a map) or in a file?
- If two artist pages are loaded at the same time and both need the same location, what prevents a double request?
- What is the cache key — the raw location string or the normalized one?

**Code Placeholder:**
```go
// geo.go

var (
    // geocache map[string]Coordinates
    // cacheMu  sync.Mutex
)

func geocodeCached(location string) (Coordinates, error) {
    // Lock the mutex
    // Check if the location is already in the cache
    // If yes: return cached coordinates, unlock
    // If no: unlock, call geocode(), lock again, store in cache, return
}
```

**Resources:**
- Search: **"golang in-memory cache map mutex"**

---

## Milestone 5 — Embed the Map on the Artist Detail Page

**Goal:** The artist detail page shows an interactive map with a marker for each concert location.

**Questions to answer:**
- How do you pass a slice of coordinates from Go to an HTML template?
- How do you use those coordinates in JavaScript inside the template to place map markers?
- What should the map show if no coordinates could be geocoded for an artist?

**Code Placeholder:**
```go
// handlers.go

type ArtistPageData struct {
    // Artist
    // Relation (dates and locations)
    // LocationCoordinates []LocationCoordinate  ← new field
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
    // ... existing logic ...

    // Geocode all locations for this artist
    // Add coordinates to the page data struct

    // Execute template with the enriched data
}
```

```html
<!-- In artist.html -->

<!-- Include your map library (e.g. Leaflet.js via CDN) -->
<!-- <link rel="stylesheet" href="https://...leaflet.css"> -->
<!-- <script src="https://...leaflet.js"></script> -->

<!-- Map container -->
<!-- <div id="map" style="height: 400px;"></div> -->

<!-- Initialize map with markers from template data -->
<!-- <script>
    var map = L.map('map').setView([0, 0], 2);
    // Add tile layer
    // Loop over coordinates from Go template and add markers
</script> -->
```

**Resources:**
- https://leafletjs.com/examples/quickstart/ — if using Leaflet
- Search: **"golang template pass data to JavaScript"**

**Verify:** The artist page shows a map. Each concert location has a marker. Clicking a marker shows the location name.

---

## Milestone 6 — Error Handling

**Goal:** The page still loads correctly even when geocoding fails.

**Questions to answer:**
- What should the map look like if zero locations were successfully geocoded?
- Should a failed geocoding call be visible to the user or silently skipped?
- What HTTP status should the server return if the geocoding API is completely unreachable?

**Verify:**
- Disconnect from the internet and load an artist page — the server does not crash
- An artist with no geocoded locations still shows their other info correctly

---

## Debugging Checklist

- Are coordinates coming back as `0, 0`? Print the raw geocoding API response — you are likely reading the wrong JSON field path.
- Is the map not appearing? Check the browser console for JavaScript errors. The map library script may not be loading.
- Are markers in the wrong place? Latitude and longitude may be swapped — geocoding APIs are inconsistent about which comes first in the response.
- Is the geocoding API returning 429 (rate limit)? You are making too many requests too fast. Add caching and consider a small delay between requests.
- Are location strings not being recognized by the geocoding API? Print the normalized string before the request to check how it looks.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `net/http` | Call geocoding API, serve pages | https://pkg.go.dev/net/http |
| `encoding/json` | Decode geocoding API response | https://pkg.go.dev/encoding/json |
| `sync` | Mutex for geocoding cache | https://pkg.go.dev/sync |
| `strings` | Normalize location strings | https://pkg.go.dev/strings |
| `fmt` | Build geocoding request URLs | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] Artist detail page includes an interactive map
- [ ] Each concert location has a marker on the map
- [ ] Markers are labeled with the location name
- [ ] Geocoding results are cached — same address never looked up twice
- [ ] Map loads correctly when some locations cannot be geocoded
- [ ] Server does not crash if the geocoding API is unreachable
- [ ] Location strings are normalized before geocoding
- [ ] All previous groupie-tracker pages still work correctly
- [ ] Unit tests for geocoding and location normalization