# 🎯 Prerequisites Quiz
## Go HTTP Client · JSON · net/http Server · html/template · HTML Basics

**Time Limit:** 60 minutes  
**Total Questions:** 35  
**Passing Score:** 28/35 (80%)

> Questions are tagged: 🟢 Easy · 🟡 Medium · 🔴 Hard  
> All topics are general — no specific project knowledge required.

---

## 📋 SECTION 1: GO HTTP CLIENT (8 Questions)

### Q1 🟢 — What does `http.Get(url)` return?

**A)** The response body as a string  
**B)** A `*http.Response` and an `error`  
**C)** Just an `error`  
**D)** A `[]byte` of the body  

<details><summary>💡 Answer</summary>

**B) `*http.Response` and an `error`**

```go
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    // network failed — no response
}
// resp.StatusCode, resp.Body, resp.Header available here
```

The `error` is non-nil only when the request itself fails (DNS failure, timeout, connection refused). A 404 or 500 response is NOT a Go error — it's a valid response with a bad status code.

</details>

---

### Q2 🟢 — Why must you call `resp.Body.Close()` after reading a response?

**A)** To save the response to disk  
**B)** To release the underlying TCP connection back to the pool — not closing it leaks connections  
**C)** To decrypt the response  
**D)** It's optional — the garbage collector handles it  

<details><summary>💡 Answer</summary>

**B) To release the TCP connection back to the connection pool**

```go
resp, err := http.Get(url)
if err != nil { return err }
defer resp.Body.Close()  // always — even if you don't read the body
```

`defer` placed immediately after the error check guarantees `Close()` runs when the function exits, even on early returns or panics. Without it, your program slowly leaks connections until it can no longer make new ones.

</details>

---

### Q3 🟡 — What is the difference between `err != nil` and checking `resp.StatusCode` after `http.Get`?

**A)** They check the same thing  
**B)** `err != nil` means the network/transport failed; a non-200 `StatusCode` means the server responded but with an error — the transport succeeded  
**C)** `resp.StatusCode` is always 200 when `err == nil`  
**D)** `err` contains the status code  

<details><summary>💡 Answer</summary>

**B) They test completely different failure modes**

```go
resp, err := http.Get(url)
if err != nil {
    return err  // network failure — DNS, timeout, refused
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("server returned: %d", resp.StatusCode)
}
```

A server returning 404 or 500 gives you `err == nil` (the HTTP conversation succeeded) but `resp.StatusCode != 200`. Always check both.

</details>

---

### Q4 🟡 — What is the most efficient way to decode a JSON response body directly from `resp.Body`?

**A)** `body, _ := io.ReadAll(resp.Body); json.Unmarshal(body, &target)`  
**B)** `json.NewDecoder(resp.Body).Decode(&target)`  
**C)** `json.Parse(resp.Body)`  
**D)** `resp.Body.Decode(&target)`  

<details><summary>💡 Answer</summary>

**B) `json.NewDecoder(resp.Body).Decode(&target)`**

```go
var result MyStruct
if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    return err
}
```

`json.NewDecoder` streams directly from the `io.Reader` without buffering the entire body in memory first. For large responses this is more efficient. Option A works correctly but reads all bytes into memory before decoding.

</details>

---

### Q5 🟡 — How do you add a custom header (e.g. `User-Agent`) to an HTTP GET request?

**A)** `http.Get(url, "User-Agent: MyApp")`  
**B)** Use `http.NewRequest` + `req.Header.Set` + `http.DefaultClient.Do(req)`  
**C)** `http.SetHeader("User-Agent", "MyApp"); http.Get(url)`  
**D)** Pass a header map as the second argument to `http.Get`  

<details><summary>💡 Answer</summary>

**B) `http.NewRequest` → set headers → `client.Do(req)`**

```go
req, err := http.NewRequest("GET", url, nil)
if err != nil { return err }

req.Header.Set("User-Agent", "MyApp/1.0")
req.Header.Set("Accept", "application/json")

resp, err := http.DefaultClient.Do(req)
```

`http.Get` is a convenience wrapper that creates a request internally. You can't customize headers through it. Build your own request with `http.NewRequest` when you need control.

</details>

---

### Q6 🟡 — What happens if you call `json.NewDecoder(resp.Body).Decode(&target)` but the server actually returned an error body (e.g. `{"error":"not found"}`)?

**A)** The decode succeeds and target gets the error message  
**B)** Decode succeeds (the JSON is valid), but `target` will have zero values for fields that don't match the error JSON's structure — you should check `StatusCode` BEFORE decoding  
**C)** Decode returns an error  
**D)** The program panics  

<details><summary>💡 Answer</summary>

**B) Decode succeeds but target will be mostly zero — check status code first**

```go
if resp.StatusCode != http.StatusOK {
    // read error body separately before returning
    var errResp ErrorResponse
    json.NewDecoder(resp.Body).Decode(&errResp)
    return fmt.Errorf("API error %d: %s", resp.StatusCode, errResp.Message)
}
// safe to decode success response
json.NewDecoder(resp.Body).Decode(&target)
```

Always check the status code before deciding which struct to decode into.

</details>

---

### Q7 🔴 — What is the default timeout for `http.DefaultClient`? Why is this a problem?

**A)** 30 seconds — long enough for any server  
**B)** No timeout — `http.DefaultClient` has no timeout by default; a hung server can block your goroutine forever  
**C)** 10 seconds  
**D)** 60 seconds  

<details><summary>💡 Answer</summary>

**B) No timeout — this is a serious production concern**

```go
// http.DefaultClient has Timeout: 0 (no timeout)

// Create a client with a timeout:
client := &http.Client{
    Timeout: 10 * time.Second,
}
resp, err := client.Get(url)
```

In production, always use a custom client with a timeout. A server that never responds will hold your goroutine (and its memory) indefinitely if you use `http.DefaultClient`.

</details>

---

### Q8 🔴 — You read `resp.Body` once. Can you read it again?

**A)** Yes — the body resets after reading  
**B)** No — `resp.Body` is a stream (`io.ReadCloser`); once read, the data is consumed and cannot be read again  
**C)** Yes — but only if you call `resp.Body.Reset()` first  
**D)** Yes — use `resp.Body.Seek(0, 0)`  

<details><summary>💡 Answer</summary>

**B) No — streams are one-way; read once**

```go
// If you need to read the body twice:
body, err := io.ReadAll(resp.Body)
resp.Body.Close()

// Now use body ([]byte) as many times as needed:
json.Unmarshal(body, &result)
log.Printf("raw body: %s", body)
```

Read into `[]byte` with `io.ReadAll` if you need multiple reads. This is also useful when debugging — print the raw body before decoding to see exactly what the server sent.

</details>

---

## 📋 SECTION 2: JSON ENCODING & DECODING (9 Questions)

### Q9 🟢 — What does a struct tag like `` `json:"name"` `` do?

**A)** Documents the field for godoc  
**B)** Maps the Go field name to a specific JSON key during marshal and unmarshal  
**C)** Makes the field required  
**D)** Sets a default value  

<details><summary>💡 Answer</summary>

**B) Maps the Go field name to its JSON key**

```go
type User struct {
    FirstName string `json:"first_name"` // JSON key: "first_name"
    Age       int    `json:"age"`
    Password  string `json:"-"`          // excluded from JSON entirely
}
```

Without the tag, `encoding/json` uses the field name exactly (case-sensitive match). The tag controls the JSON key name. `json:"-"` means "never marshal or unmarshal this field."

</details>

---

### Q10 🟢 — What is the zero value for a field that is present in the Go struct but missing from the incoming JSON?

**A)** The decode returns an error  
**B)** The field silently gets its Go zero value (`0`, `""`, `false`, `nil`)  
**C)** The field is set to `nil`  
**D)** The decode panics  

<details><summary>💡 Answer</summary>

**B) The field gets its zero value — no error, silent**

```go
type Item struct {
    Name  string `json:"name"`
    Count int    `json:"count"`
}

data := []byte(`{"name": "apple"}`)
var item Item
json.Unmarshal(data, &item)
// item.Name == "apple", item.Count == 0 (key missing from JSON)
```

This is the most common JSON debugging trap. If a field is zero after decoding, the JSON key probably doesn't match the struct tag. Tags are **case-sensitive**.

</details>

---

### Q11 🟢 — What type represents a JSON object with unknown keys in Go?

**A)** `interface{}`  
**B)** `map[string]interface{}`  
**C)** `struct{}`  
**D)** `json.Object`  

<details><summary>💡 Answer</summary>

**B) `map[string]interface{}`**

```go
var data map[string]interface{}
json.Unmarshal(body, &data)

name := data["name"].(string)    // type assertion
count := data["count"].(float64) // JSON numbers decode as float64
```

When you don't know the JSON structure ahead of time, use `map[string]interface{}`. Note: all JSON numbers become `float64`, not `int`, when decoded into `interface{}`.

</details>

---

### Q12 🟡 — What Go type should you use for this JSON field?
```json
{ "tags": ["go", "web", "api"] }
```

**A)** `string`  
**B)** `map[string]string`  
**C)** `[]string`  
**D)** `[3]string`  

<details><summary>💡 Answer</summary>

**C) `[]string`**

```go
type Article struct {
    Tags []string `json:"tags"`
}
```

JSON arrays map to Go slices. The size is dynamic so a slice is correct — not an array (which has a fixed, compile-time size). A nil slice field will encode as `null`; use `[]string{}` or the `omitempty` option to avoid this.

</details>

---

### Q13 🟡 — What does the `omitempty` option in a struct tag do?

**A)** Skips the field during decode if it's empty  
**B)** During marshal (encoding to JSON), omits the field if it has its zero value (`0`, `""`, `false`, `nil`, empty slice/map)  
**C)** Makes the field optional during unmarshal  
**D)** Converts empty strings to `null`  

<details><summary>💡 Answer</summary>

**B) Omits the field from the JSON output when it is zero/empty**

```go
type Response struct {
    Name  string `json:"name"`
    Error string `json:"error,omitempty"` // omitted if ""
    Count int    `json:"count,omitempty"` // omitted if 0
}
```

Useful for APIs where you don't want to include `"error": ""` in every successful response. Has no effect on unmarshal — missing keys are still handled the same way.

</details>

---

### Q14 🟡 — What is the correct Go type for this JSON field?
```json
{
  "concerts": {
    "london-uk": ["2023-07-01", "2023-07-02"],
    "berlin-de": ["2023-08-15"]
  }
}
```

**A)** `map[string]string`  
**B)** `[][]string`  
**C)** `map[string][]string`  
**D)** `map[string]interface{}`  

<details><summary>💡 Answer</summary>

**C) `map[string][]string`**

```go
type Event struct {
    Concerts map[string][]string `json:"concerts"`
}
```

The outer object has string keys (location names), and each value is an array of strings (dates). This maps exactly to `map[string][]string`. Getting this type right is one of the most important JSON skills — always look at the structure carefully before writing the Go type.

</details>

---

### Q15 🟡 — How do you marshal a Go struct to a JSON `[]byte`?

**A)** `json.Encode(myStruct)`  
**B)** `json.Marshal(myStruct)` → returns `([]byte, error)`  
**C)** `json.Stringify(myStruct)`  
**D)** `fmt.Sprintf("%j", myStruct)`  

<details><summary>💡 Answer</summary>

**B) `json.Marshal(myStruct)` returns `([]byte, error)`**

```go
type Point struct {
    X int `json:"x"`
    Y int `json:"y"`
}

p := Point{X: 3, Y: 7}
data, err := json.Marshal(p)
if err != nil { return err }
fmt.Println(string(data))  // {"x":3,"y":7}
```

`json.MarshalIndent(v, "", "  ")` produces pretty-printed output with 2-space indentation.

</details>

---

### Q16 🔴 — What happens if a JSON struct field is unexported (lowercase)?

**A)** `json.Marshal` panics  
**B)** The field is silently ignored during both marshal and unmarshal — only exported (uppercase) fields are visible to `encoding/json`  
**C)** An error is returned  
**D)** The field is included with its Go name  

<details><summary>💡 Answer</summary>

**B) Silently ignored — unexported fields are invisible to `encoding/json`**

```go
type User struct {
    Name     string `json:"name"`   // exported — included in JSON
    password string                  // unexported — NEVER in JSON
}

u := User{Name: "Alice", password: "secret"}
data, _ := json.Marshal(u)
fmt.Println(string(data)) // {"name":"Alice"} — password excluded
```

This is a security feature but also a frequent bug: fields that should appear in JSON are accidentally lowercase. Always check capitalization if a field isn't appearing in your JSON output.

</details>

---

### Q17 🔴 — What does `json.Unmarshal` do if the JSON contains a key that doesn't exist in the target struct?

**A)** Returns an error  
**B)** Panics  
**C)** Silently ignores the unknown key — only known fields are populated  
**D)** Stores it in a catch-all field  

<details><summary>💡 Answer</summary>

**C) Unknown keys are silently ignored by default**

```go
type Minimal struct {
    Name string `json:"name"`
}

data := []byte(`{"name":"Alice","age":30,"email":"a@b.com"}`)
var m Minimal
json.Unmarshal(data, &m)
// m.Name == "Alice" — age and email are silently dropped
```

This is by design — it allows JSON APIs to add new fields without breaking existing Go clients. To detect unknown fields (e.g. for strict parsing), use `json.Decoder` with `DisallowUnknownFields()`.

</details>

---

## 📋 SECTION 3: net/http SERVER (9 Questions)

### Q18 🟢 — What does `http.HandleFunc("/", myHandler)` register?

**A)** A handler for exactly the path `/`  
**B)** A handler for `/` AND all paths that don't match any other registered handler (catch-all)  
**C)** A handler only for the root domain  
**D)** A handler for all paths starting with `/`  

<details><summary>💡 Answer</summary>

**B) A catch-all — matches `/` AND all unregistered paths**

```go
http.HandleFunc("/", homeHandler)
// This handles: /, /anything, /foo/bar, /favicon.ico, etc.

// Fix — check the exact path:
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    // handle home page
}
```

The trailing-slash rule in Go's default mux: a pattern ending in `/` (like `/`) is a subtree — it matches everything under it. A pattern without trailing slash (like `/about`) matches exactly.

</details>

---

### Q19 🟢 — What is the correct HTTP status code for each situation?

| Situation | Code |
|---|---|
| Request processed successfully | ? |
| Client sent an invalid/missing parameter | ? |
| The requested resource doesn't exist | ? |
| Server encountered an unexpected error | ? |

**A)** 200, 400, 404, 500  
**B)** 200, 404, 400, 503  
**C)** 201, 400, 404, 500  
**D)** 200, 403, 404, 500  

<details><summary>💡 Answer</summary>

**A) 200, 400, 404, 500**

```go
// 200 OK — success
w.WriteHeader(http.StatusOK)  // or just write a body (200 is default)

// 400 Bad Request — client error (bad input)
http.Error(w, "invalid id", http.StatusBadRequest)

// 404 Not Found — resource doesn't exist
http.NotFound(w, r)  // or http.Error(w, "not found", http.StatusNotFound)

// 500 Internal Server Error — server-side failure
http.Error(w, "internal error", http.StatusInternalServerError)
```

The distinction between 400 and 404 matters: 400 = "your request is malformed," 404 = "your request is valid but the thing doesn't exist."

</details>

---

### Q20 🟢 — How do you write a response body to a `http.ResponseWriter`?

**A)** `w.Send("hello")`  
**B)** `w.Write([]byte("hello"))` or `fmt.Fprintf(w, "hello %s", name)`  
**C)** `w.Body = "hello"`  
**D)** `http.Write(w, "hello")`  

<details><summary>💡 Answer</summary>

**B) `w.Write([]byte(...))` or `fmt.Fprintf(w, ...)`**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
    // or:
    w.Write([]byte("Hello!"))
}
```

`http.ResponseWriter` implements `io.Writer`, so `fmt.Fprintf` works directly. The first call to `Write` sends headers (including a 200 status) if `WriteHeader` hasn't been called yet.

</details>

---

### Q21 🟡 — What is the rule about calling `w.Header().Set()` relative to `w.WriteHeader()` or `w.Write()`?

**A)** Headers can be set at any time — they are sent at the end  
**B)** Headers MUST be set before `WriteHeader` or `Write` — once the response body starts, headers are already sent and any header changes are silently ignored  
**C)** Headers must be set after `WriteHeader`  
**D)** Headers are automatically set — you never need to set them manually  

<details><summary>💡 Answer</summary>

**B) Set headers BEFORE writing the body**

```go
// CORRECT:
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(result)

// WRONG — header is ignored, a warning is logged:
json.NewEncoder(w).Encode(result)     // Write called — headers sent
w.Header().Set("Content-Type", "application/json")  // too late!
```

HTTP responses send headers first, then the body. Once the first byte of the body is written, headers are locked. This is one of the most common bugs in Go HTTP handlers.

</details>

---

### Q22 🟡 — How do you read a URL query parameter from `r.URL` for the URL `/search?q=golang&limit=10`?

**A)** `r.Params["q"]`  
**B)** `r.URL.Query().Get("q")`  
**C)** `r.QueryParam("q")`  
**D)** `r.Form["q"]`  

<details><summary>💡 Answer</summary>

**B) `r.URL.Query().Get("q")`**

```go
func searchHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")          // "golang"
    limit := r.URL.Query().Get("limit")  // "10" (always a string)

    // For multiple values with the same key:
    tags := r.URL.Query()["tag"]  // []string — if URL had ?tag=a&tag=b
}
```

`r.URL.Query()` returns a `url.Values` (which is `map[string][]string`). `.Get()` returns the first value for a key, or `""` if missing.

</details>

---

### Q23 🟡 — How do you read a form field from a POST request with `Content-Type: application/x-www-form-urlencoded`?

**A)** `r.Body.Get("fieldname")`  
**B)** `r.FormValue("fieldname")`  
**C)** `r.PostParam("fieldname")`  
**D)** `r.URL.Query().Get("fieldname")`  

<details><summary>💡 Answer</summary>

**B) `r.FormValue("fieldname")`**

```go
func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    name := r.FormValue("name")  // reads from POST body
    email := r.FormValue("email")
}
```

`r.FormValue` parses the request body (for POST) and query string (for GET) and returns the first value for the named key. It calls `r.ParseForm()` internally. `r.URL.Query().Get()` reads ONLY the query string, not the POST body.

</details>

---

### Q24 🟡 — What is the purpose of `http.ListenAndServe(":8080", nil)`?

**A)** Creates a new HTTP client  
**B)** Starts a TCP server on port 8080 that routes requests using the default `ServeMux` (the `nil` argument) — this call blocks until the server stops  
**C)** Listens for one connection then exits  
**D)** The `nil` means the server doesn't handle any requests  

<details><summary>💡 Answer</summary>

**B) Starts the TCP server, blocks, routes via default mux when `nil` is passed**

```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/api", apiHandler)

// This blocks — run in goroutine if you need to do other things:
log.Fatal(http.ListenAndServe(":8080", nil))
```

`nil` means "use `http.DefaultServeMux`" — the same mux that `http.HandleFunc` registers to. Passing a custom `*http.ServeMux` gives you an isolated router.

</details>

---

### Q25 🔴 — What is the risk of writing to `w` after calling `http.Error(w, ..., code)`?

**A)** No risk — you can write additional data  
**B)** `http.Error` writes the header and body in one call; any subsequent `w.Write` or `w.WriteHeader` is ignored (headers already sent) — always `return` immediately after error responses  
**C)** The second write overwrites the first  
**D)** It causes a panic  

<details><summary>💡 Answer</summary>

**B) Headers are already sent — always `return` after `http.Error`**

```go
// BUG — continues executing after error response:
func handler(w http.ResponseWriter, r *http.Request) {
    if id == 0 {
        http.Error(w, "bad id", http.StatusBadRequest)
        // NO return — falls through to execute success path!
    }
    // ... renders template anyway, but headers are already sent
}

// CORRECT:
if id == 0 {
    http.Error(w, "bad id", http.StatusBadRequest)
    return  // stop processing
}
```

This is one of the most common bugs in Go HTTP handlers. `http.Error` does not stop function execution.

</details>

---

### Q26 🔴 — What does `http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))` do, and why is `StripPrefix` needed?

**A)** Serves files and strips the directory listing  
**B)** `FileServer` serves files from `./static`; without `StripPrefix`, a request for `/static/style.css` would look for `./static/static/style.css` (doubled path); `StripPrefix` removes `/static/` before FileServer sees the path  
**C)** Strips file extensions from URLs  
**D)** `StripPrefix` is not needed — `FileServer` handles this automatically  

<details><summary>💡 Answer</summary>

**B) `StripPrefix` removes the URL prefix before passing to FileServer**

```go
// Without StripPrefix:
// GET /static/style.css → FileServer looks for ./static/static/style.css (wrong)

// With StripPrefix:
// GET /static/style.css → StripPrefix removes "/static/" → FileServer sees "style.css" → serves ./static/style.css (correct)

http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
```

This is the canonical pattern for serving static files in Go.

</details>

---

## 📋 SECTION 4: html/template (9 Questions)

### Q27 🟢 — How do you output a variable named `.Name` in a Go HTML template?

**A)** `<%= .Name %>`  
**B)** `{{ .Name }}`  
**C)** `{{{ .Name }}}`  
**D)** `${.Name}`  

<details><summary>💡 Answer</summary>

**B) `{{ .Name }}`**

```html
<!-- Template -->
<h1>Hello, {{ .Name }}!</h1>
<p>Age: {{ .Age }}</p>

<!-- In Go: -->
data := struct{ Name string; Age int }{"Alice", 30}
tmpl.Execute(w, data)
```

The `.` (dot) refers to the data passed to `Execute`. `{{ .Name }}` accesses the `Name` field of that data. HTML template automatically escapes the output to prevent XSS.

</details>

---

### Q28 🟢 — How do you iterate over a slice in a template?

**A)** `{{ for item in .Items }}`  
**B)** `{{ range .Items }}{{ . }}{{ end }}`  
**C)** `{{ each .Items as item }}{{ item }}{{ end }}`  
**D)** `{{ loop .Items }}`  

<details><summary>💡 Answer</summary>

**B) `{{ range .Items }}...{{ end }}`**

```html
<ul>
{{ range .Items }}
    <li>{{ .Name }} — {{ .Price }}</li>
{{ end }}
</ul>

<!-- With index: -->
{{ range $i, $item := .Items }}
    <li>{{ $i }}: {{ $item.Name }}</li>
{{ end }}
```

Inside `{{ range }}`, `.` changes to refer to the current element. Use `$i, $item :=` syntax when you need both index and value.

</details>

---

### Q29 🟢 — How do you load and execute a template from a file?

**A)** `template.Open("page.html").Execute(w, data)`  
**B)** `template.ParseFiles("page.html")` → `tmpl.Execute(w, data)`  
**C)** `template.Load("page.html", data)`  
**D)** `html.RenderFile("page.html", w, data)`  

<details><summary>💡 Answer</summary>

**B) `ParseFiles` then `Execute`**

```go
tmpl, err := template.ParseFiles("templates/page.html")
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
if err := tmpl.Execute(w, data); err != nil {
    log.Println("template execute error:", err)
}
```

`ParseFiles` can accept multiple files: `template.ParseFiles("base.html", "page.html")`. Templates are usually parsed once at startup and cached, not re-parsed on every request.

</details>

---

### Q30 🟡 — What is the `{{ if }}` template action used for?

**A)** Importing other templates  
**B)** Conditionally rendering a block based on a value's truthiness  
**C)** Iterating over a collection  
**D)** Defining a template function  

<details><summary>💡 Answer</summary>

**B) Conditional rendering**

```html
{{ if .Error }}
    <div class="error">{{ .Error }}</div>
{{ else if .Warning }}
    <div class="warning">{{ .Warning }}</div>
{{ else }}
    <div class="success">All good!</div>
{{ end }}

<!-- Check a boolean field: -->
{{ if .IsAdmin }}
    <a href="/admin">Admin Panel</a>
{{ end }}
```

In templates, "falsy" values are: `false`, `0`, `nil`, empty string, empty slice/map/channel. Everything else is truthy.

</details>

---

### Q31 🟡 — What is the difference between `{{ .Field }}` and `{{ .Method }}`?

**A)** There is no difference  
**B)** Both work — `html/template` can access struct fields and call methods that return one or two values (second must be `error`)  
**C)** Only fields can be accessed in templates  
**D)** Methods require parentheses: `{{ .Method() }}`  

<details><summary>💡 Answer</summary>

**B) Both fields and methods are accessible — no parentheses needed**

```go
type User struct {
    FirstName string
    LastName  string
}
func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}
```

```html
{{ .FirstName }}   <!-- field access -->
{{ .FullName }}    <!-- method call, no parentheses -->
```

Methods called from templates must return either one value or `(value, error)`. If a method returns an error, the template execution is halted.

</details>

---

### Q32 🟡 — How does `html/template` differ from `text/template`?

**A)** They are identical  
**B)** `html/template` automatically HTML-escapes all values injected into the template, preventing XSS attacks; `text/template` does no escaping  
**C)** `html/template` is faster  
**D)** `text/template` only works with plain text files  

<details><summary>💡 Answer</summary>

**B) `html/template` auto-escapes HTML; `text/template` does not**

```go
// With html/template:
data := "<script>alert('xss')</script>"
// Template: <p>{{ . }}</p>
// Output: <p>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</p>

// With text/template:
// Output: <p><script>alert('xss')</script></p>  ← XSS!
```

Always use `html/template` for web pages. Only use `text/template` for non-HTML output (emails, config files, etc.).

</details>

---

### Q33 🟡 — How do you iterate over a `map[string][]string` in a template?

**A)** `{{ range .Map }}{{ . }}{{ end }}`  
**B)** `{{ range $key, $values := .Map }}{{ $key }}: {{ range $values }}{{ . }}{{ end }}{{ end }}`  
**C)** Maps can't be iterated in templates  
**D)** `{{ for k, v := range .Map }}`  

<details><summary>💡 Answer</summary>

**B) Nested `range` with `$key, $values`**

```html
{{ range $location, $dates := .Concerts }}
    <h3>{{ $location }}</h3>
    <ul>
        {{ range $dates }}
        <li>{{ . }}</li>
        {{ end }}
    </ul>
{{ end }}
```

Two-variable `range` gives you both key and value. The inner `range $dates` iterates the string slice; `.` inside it refers to each date string. Note: map iteration order is random.

</details>

---

### Q34 🔴 — What happens if `tmpl.Execute(w, data)` fails after the response body has already started writing?

**A)** The partial response is discarded and a 500 is sent  
**B)** Nothing visible to the client — headers are already sent with 200; the client receives a partial or corrupted page  
**C)** The template retries  
**D)** A panic is sent to the client  

<details><summary>💡 Answer</summary>

**B) The client gets a partial page — you can't send a 500 after starting the body**

```go
// WRONG pattern:
tmpl.Execute(w, data)   // body starts writing immediately
// if Execute fails partway through, client gets half a page

// BETTER pattern: render to buffer first
var buf bytes.Buffer
if err := tmpl.Execute(&buf, data); err != nil {
    http.Error(w, "template error", 500)
    return
}
buf.WriteTo(w)  // only write to w if rendering succeeded
```

For critical pages, render to `bytes.Buffer` first. For simpler cases, log the error and accept that the page may be partially rendered.

</details>

---

### Q35 🔴 — What does `{{ template "header" . }}` do inside a template?

**A)** Imports the file named `header`  
**B)** Calls a named template called `"header"` and passes the current dot value (`.`) to it as its data  
**C)** Outputs the string `"header"`  
**D)** Renders the `<header>` HTML element  

<details><summary>💡 Answer</summary>

**B) Calls a named sub-template, passing the current data**

```html
<!-- In base.html: -->
{{ define "layout" }}
<html>
<body>
    {{ template "content" . }}
</body>
</html>
{{ end }}

<!-- In page.html: -->
{{ define "content" }}
<h1>{{ .Title }}</h1>
{{ end }}
```

```go
// Load both files:
tmpl := template.Must(template.ParseFiles("base.html", "page.html"))
tmpl.ExecuteTemplate(w, "layout", data)
```

Named templates are how Go implements template inheritance and composition. The `.` (current data) is explicitly passed to the sub-template — it doesn't flow through automatically.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 33–35 ✅ | **Exceptional** — very strong foundation across all four topics. |
| 28–32 ✅ | **Ready** — review individual missed sections before starting. |
| 21–27 ⚠️ | **Study first** — identify which section you scored lowest on and work through it before proceeding. |
| Below 21 ❌ | **Not ready** — multiple topics need work. Budget at least a week on HTTP, JSON, and templates before starting. |

---

## 🔍 Missed Questions Guide

| Missed | What to study | Resources |
|---|---|---|
| Q1–Q8 | `http.Get`, `resp.Body.Close()`, status codes vs errors, custom requests, timeouts | `pkg.go.dev/net/http` |
| Q9–Q17 | Struct tags, zero on missing key, `map[string]interface{}`, `omitempty`, marshal vs unmarshal | `pkg.go.dev/encoding/json` |
| Q18–Q26 | Handler registration, status codes, `WriteHeader` ordering, query params, form values, `FileServer` | `pkg.go.dev/net/http` |
| Q27–Q35 | `{{ range }}`, `{{ if }}`, `ParseFiles`, `Execute`, escaping, named templates | `pkg.go.dev/html/template` |

---

## 🧪 Difficulty Breakdown

| Difficulty | Questions | Topics tested |
|---|---|---|
| 🟢 Easy (11) | Q1, Q2, Q9, Q10, Q11, Q18, Q19, Q20, Q27, Q28, Q29 | Core syntax, zero values, basic calls |
| 🟡 Medium (15) | Q3, Q4, Q5, Q6, Q12, Q13, Q14, Q15, Q21, Q22, Q23, Q24, Q30, Q31, Q32, Q33 | Combining concepts, real gotchas |
| 🔴 Hard (9) | Q7, Q8, Q16, Q17, Q25, Q26, Q34, Q35 | Production concerns, subtle behavior, composition |