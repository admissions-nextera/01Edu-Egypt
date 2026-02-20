# 🎯 ASCII-Art-Web Prerequisites Quiz
## HTTP Basics · net/http · html/template · HTML Forms · Status Codes

**Time Limit:** 50 minutes  
**Total Questions:** 28  
**Passing Score:** 22/28 (78%)

> ✅ Pass → You're ready to start ASCII-Art-Web  
> ⚠️ Also Required → ASCII-Art must be fully working and `render` must return a string

---

## 📋 SECTION 1: HTTP FUNDAMENTALS (7 Questions)

### Q1: What is the key difference between an HTTP GET and POST request?

**A)** GET is faster than POST  
**B)** GET requests data (parameters in the URL); POST sends data (in the request body, invisible in the URL)  
**C)** POST is encrypted; GET is not  
**D)** They are interchangeable — the difference is just convention  

<details><summary>💡 Answer</summary>

**B) GET requests data (parameters in URL); POST sends data in the request body**

For this project: the main page loads with GET `/`. The form submits user input (text + banner choice) with POST `/ascii-art`. Using GET for a form that sends user content would expose the data in the URL and violate HTTP semantics.

</details>

---

### Q2: What does HTTP status code `200` mean?

**A)** Redirect  
**B)** OK — the request succeeded  
**C)** Not Found  
**D)** Internal Server Error  

<details><summary>💡 Answer</summary>

**B) OK — the request succeeded**

The standard HTTP status codes you must know for this project:
- `200` OK — success
- `400` Bad Request — client sent invalid input
- `404` Not Found — resource doesn't exist
- `500` Internal Server Error — something broke on the server

</details>

---

### Q3: A user visits a URL that your server doesn't have a handler for. What status code should be returned?

**A)** 200  
**B)** 400  
**C)** 404  
**D)** 500  

<details><summary>💡 Answer</summary>

**C) 404**

`404 Not Found` means the server received the request but has no resource at that path. The browser (or curl) should get a `404` response, not a crash.

</details>

---

### Q4: A user submits the form with the text field empty. What status code should your server return?

**A)** 200 — return an empty result  
**B)** 404 — no content found  
**C)** 400 — the client sent invalid (empty) input  
**D)** 500 — the server can't handle it  

<details><summary>💡 Answer</summary>

**C) 400 — the client sent invalid input**

`400 Bad Request` is for client errors — invalid input, missing required fields, wrong format. The server is working fine; the input is the problem. `500` would be wrong because the server itself didn't fail.

</details>

---

### Q5: Your server fails to parse a template file. What status code should be returned?

**A)** 400  
**B)** 404  
**C)** 500  
**D)** 200 with an error message in the body  

<details><summary>💡 Answer</summary>

**C) 500**

`500 Internal Server Error` means the server itself failed — not because of bad client input. A missing or unparseable template is a server-side misconfiguration. The client did nothing wrong.

</details>

---

### Q6: What happens if you try to set a response header AFTER calling `w.Write()`?

**A)** The header is set normally  
**B)** The header is silently ignored — once the body is written, headers are locked  
**C)** The server panics  
**D)** The header is queued and sent with the next response  

<details><summary>💡 Answer</summary>

**B) The header is silently ignored — headers are locked after writing the body**

Go will log a "superfluous response.WriteHeader call" warning. Always set headers BEFORE writing the body:
```go
// CORRECT order:
w.Header().Set("Content-Type", "text/html")
w.WriteHeader(400)
w.Write([]byte("error message"))
```

</details>

---

### Q7: What is a request handler in Go's `net/http` package? What is its function signature?

**A)** `func handler(request string) string`  
**B)** `func handler(w http.ResponseWriter, r *http.Request)`  
**C)** `func handler(w http.Writer, r http.Request)`  
**D)** `func handler(ctx context.Context, w http.ResponseWriter)`  

<details><summary>💡 Answer</summary>

**B) `func handler(w http.ResponseWriter, r *http.Request)`**

- `w http.ResponseWriter` — you write your response to this (headers, status, body)
- `r *http.Request` — contains everything about the incoming request (method, URL, headers, body, form data)

This exact signature is required by `http.HandleFunc`.

</details>

---

## 📋 SECTION 2: net/http PACKAGE (6 Questions)

### Q8: What is the correct way to start an HTTP server on port 8080?

**A)** `http.Start(":8080", nil)`  
**B)** `http.ListenAndServe(":8080", nil)`  
**C)** `http.Serve(8080)`  
**D)** `net.Listen("tcp", ":8080")`  

<details><summary>💡 Answer</summary>

**B) `http.ListenAndServe(":8080", nil)`**

`nil` as the second argument uses the default `ServeMux`. `ListenAndServe` blocks forever (or until the server errors). Always check its return value:
```go
if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
}
```

</details>

---

### Q9: What does `http.HandleFunc("/", homeHandler)` do?

**A)** Runs `homeHandler` immediately  
**B)** Registers `homeHandler` to be called whenever a request comes in for the path `"/"`  
**C)** Registers `homeHandler` for all paths  
**D)** Serves the file at `"/"` from disk  

<details><summary>💡 Answer</summary>

**C) Registers `homeHandler` for ALL paths — not just `"/"`**

This is a critical Go gotcha. `http.HandleFunc("/", ...)` acts as a catch-all — it matches any path that doesn't have a more specific handler. You must explicitly check `r.URL.Path == "/"` inside the handler and return `404` for everything else:
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    // ... handle "/"
}
```

</details>

---

### Q10: How do you read a form value from a POST request in Go?

**A)** `r.Body["text"]`  
**B)** `r.Query("text")`  
**C)** `r.FormValue("text")`  
**D)** `r.Form["text"][0]`  

<details><summary>💡 Answer</summary>

**C) `r.FormValue("text")`**

`r.FormValue(key)` returns the first value for the named form field. It works for both POST form data and URL query parameters. Returns an empty string if the key doesn't exist — which is why you must check for empty values manually.

</details>

---

### Q11: What is the correct way to return a 400 status with an error message?

**A)**
```go
w.Write([]byte("Bad Request"))
w.WriteHeader(400)
```
**B)**
```go
http.Error(w, "Bad Request: text cannot be empty", http.StatusBadRequest)
```
**C)**
```go
w.WriteHeader(400)
// (nothing else)
```
**D)**
```go
return http.StatusBadRequest
```

<details><summary>💡 Answer</summary>

**B) `http.Error(w, "message", http.StatusBadRequest)`**

`http.Error` sets the correct `Content-Type`, writes the status code, and writes the message body — all in one call. Option A has the wrong order (body before status). Option C writes no body. Option D doesn't compile — handler functions don't return HTTP statuses.

</details>

---

### Q12: How do you check which HTTP method was used in a request?

**A)** `r.HttpMethod`  
**B)** `r.Method`  
**C)** `http.GetMethod(r)`  
**D)** `r.Header.Get("Method")`  

<details><summary>💡 Answer</summary>

**B) `r.Method`**

```go
if r.Method != http.MethodPost {
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
}
```

Use the constants: `http.MethodGet`, `http.MethodPost`, `http.MethodPut`, etc. — don't hardcode the strings `"GET"` or `"POST"`.

</details>

---

### Q13: What is the output when this handler is hit?
```go
func handler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "oops", 500)
    w.WriteHeader(200)
    fmt.Fprintln(w, "all good")
}
```

**A)** Status 200, body "all good"  
**B)** Status 500, body "oops\n", and a warning about superfluous WriteHeader  
**C)** Status 500, body "oops\nall good"  
**D)** Compile error  

<details><summary>💡 Answer</summary>

**B) Status 500, body "oops\n", and a warning about superfluous WriteHeader**

`http.Error` writes the status AND body. The subsequent `w.WriteHeader(200)` is ignored (headers already sent) and Go logs a warning. The `fmt.Fprintln` may or may not append to the body depending on buffering, but the status will be 500. Always `return` after sending an error.

</details>

---

## 📋 SECTION 3: html/template (6 Questions)

### Q14: What is the difference between `text/template` and `html/template`?

**A)** No functional difference  
**B)** `html/template` automatically escapes HTML entities to prevent XSS attacks; `text/template` does not  
**C)** `html/template` is faster  
**D)** `text/template` supports more syntax features  

<details><summary>💡 Answer</summary>

**B) `html/template` automatically escapes HTML entities to prevent XSS attacks**

If a user types `<script>alert('xss')</script>` as input and you use `text/template`, it renders as raw HTML — dangerous. `html/template` converts `<` to `&lt;`, making it safe. Always use `html/template` for web pages.

</details>

---

### Q15: Which is the correct way to parse and execute a template file?

**A)**
```go
tmpl := template.New("index.html")
tmpl.Execute(w, data)
```
**B)**
```go
tmpl, err := template.ParseFiles("templates/index.html")
if err != nil { http.Error(w, err.Error(), 500); return }
err = tmpl.Execute(w, data)
if err != nil { http.Error(w, err.Error(), 500); return }
```
**C)**
```go
tmpl := template.Must(template.ParseFiles("templates/index.html"))
tmpl.Execute(w, nil)
```
**D)**
```go
http.ServeFile(w, r, "templates/index.html")
```

<details><summary>💡 Answer</summary>

**B)**

Option B is correct for a web handler: parse, check error, execute, check error. Option C uses `template.Must` which panics on error — acceptable at startup, not inside a per-request handler. Option D serves the raw file without template processing. Option A doesn't load anything.

</details>

---

### Q16: How do you display a variable called `Result` from Go inside an HTML template?

**A)** `$Result`  
**B)** `{Result}`  
**C)** `{{ .Result }}`  
**D)** `<%= Result %>`  

<details><summary>💡 Answer</summary>

**C) `{{ .Result }}`**

Go templates use `{{ }}` delimiters. The `.` refers to the data value passed to `Execute`. If you pass a struct, `.Result` accesses the `Result` field. The field must be exported (start with a capital letter).

</details>

---

### Q17: You pass this struct to a template:
```go
type PageData struct {
    Result string
    Error  string
}
```
How do you show the error message only if it's not empty?

**A)** `<p>{{ if .Error != "" }}{{ .Error }}{{ end }}</p>`  
**B)** `<p>{{ if .Error }}{{ .Error }}{{ end }}</p>`  
**C)** `<p>{{ .Error || "" }}</p>`  
**D)** `<p>{{ when .Error }}{{ .Error }}</p>`  

<details><summary>💡 Answer</summary>

**B) `{{ if .Error }}{{ .Error }}{{ end }}`**

In Go templates, `{{ if .Error }}` evaluates to true if `.Error` is a non-empty string (or non-zero value). The syntax `{{ if .Error != "" }}` is not valid Go template syntax — comparison operators work differently in templates.

</details>

---

### Q18: The ASCII art result contains spaces and newlines. You display it with `{{ .Result }}` in a `<p>` tag. What problem occurs?

**A)** No problem — HTML renders it correctly  
**B)** The browser collapses all whitespace — the art looks broken  
**C)** The template escapes newlines as `\n` in the output  
**D)** The `<p>` tag prevents rendering  

<details><summary>💡 Answer</summary>

**B) The browser collapses all whitespace — the art looks broken**

HTML treats multiple spaces and newlines as a single space. The fix: wrap the result in `<pre>` tags, which preserve whitespace exactly:
```html
<pre>{{ .Result }}</pre>
```
`<pre>` uses a monospace font by default and preserves all whitespace — essential for ASCII art.

</details>

---

### Q19: `template.ParseFiles("templates/index.html")` fails with "no such file or directory". What are the two most likely causes?

**A)** The file doesn't exist, or the working directory when running `go run .` is not the project root  
**B)** The template syntax is wrong  
**C)** The file has a `.html` extension instead of `.tmpl`  
**D)** You need to use an absolute path  

<details><summary>💡 Answer</summary>

**A) The file doesn't exist, or the working directory is not the project root**

`ParseFiles` uses a path relative to the working directory. When you run `go run .` from the project root, the working directory is the project root. If you run from a different directory, the relative path breaks. Print `os.Getwd()` to debug.

</details>

---

## 📋 SECTION 4: HTML FORMS (5 Questions)

### Q20: What does `<form method="POST" action="/ascii-art">` do?

**A)** Sends a GET request to `/ascii-art` when submitted  
**B)** Sends a POST request to `/ascii-art` when submitted, with all form field values in the request body  
**C)** Navigates the browser to `/ascii-art` without sending data  
**D)** Calls a JavaScript function named `ascii-art`  

<details><summary>💡 Answer</summary>

**B) Sends a POST request to `/ascii-art` when submitted, with all form field values in the request body**

The `method` attribute controls GET vs POST. The `action` attribute is the URL to send to. Each `<input>`, `<select>`, or `<textarea>` inside the form is included in the submission using its `name` attribute as the key.

</details>

---

### Q21: You have `<textarea name="text"></textarea>` and `<select name="banner">...</select>` in your form. What does `r.FormValue("text")` return after submission?

**A)** The HTML of the textarea element  
**B)** The text the user typed into the textarea  
**C)** `"text"` (the name attribute)  
**D)** `nil`  

<details><summary>💡 Answer</summary>

**B) The text the user typed into the textarea**

`r.FormValue("text")` looks up the submitted value for the form field with `name="text"`. For a textarea, this is whatever the user typed. For a select, it's the `value` of the selected `<option>`.

</details>

---

### Q22: How do you create radio buttons for banner selection (standard, shadow, thinkertoy) so that only one can be selected at a time?

**A)**
```html
<input type="radio" name="banner1" value="standard">
<input type="radio" name="banner2" value="shadow">
<input type="radio" name="banner3" value="thinkertoy">
```
**B)**
```html
<input type="radio" name="banner" value="standard">
<input type="radio" name="banner" value="shadow">
<input type="radio" name="banner" value="thinkertoy">
```
**C)**
```html
<input type="checkbox" name="banner" value="standard">
<input type="checkbox" name="banner" value="shadow">
```
**D)**
```html
<select name="banner">
    <option>standard</option>
</select>
```

<details><summary>💡 Answer</summary>

**B)**

Radio buttons with the **same `name`** attribute form a group — only one can be selected at a time. The selected button's `value` is submitted. Option A gives each button a different name (they act independently). A `<select>` (option D) also works, but the question specifically asks about radio buttons.

</details>

---

### Q23: What attribute on a radio button or option makes it selected by default when the page loads?

**A)** `default`  
**B)** `selected` (for radio: `checked`)  
**C)** `active`  
**D)** `value="default"`  

<details><summary>💡 Answer</summary>

**B) `checked` for radio buttons, `selected` for `<option>` elements**

```html
<input type="radio" name="banner" value="standard" checked>
<select name="banner">
    <option value="standard" selected>Standard</option>
</select>
```

For the form to stay filled after submission, your template must conditionally add `checked`/`selected` based on the submitted value (stored in `PageData`).

</details>

---

### Q24: After the form submits and the page reloads with results, the user sees the banner selector reset to default. How do you keep the user's selection?

**A)** Use JavaScript to remember it  
**B)** Store the selected banner in `PageData` and conditionally render `checked` or `selected` in the template  
**C)** Use browser cookies automatically  
**D)** It's not possible without JavaScript  

<details><summary>💡 Answer</summary>

**B) Store the selected banner in `PageData` and conditionally render the `checked`/`selected` attribute**

```go
type PageData struct {
    Result         string
    SelectedBanner string
}
```
```html
<input type="radio" name="banner" value="standard"
    {{ if eq .SelectedBanner "standard" }}checked{{ end }}>
```

This keeps the form state consistent after submission — good UX and often tested in the spec.

</details>

---

## 📋 SECTION 5: INTEGRATION & STRUCTURE (4 Questions)

### Q25: What is the correct order of operations in your POST `/ascii-art` handler?

**A)** Render → Validate → Load banner → Return  
**B)** Validate input → Load banner → Render → Return result  
**C)** Load banner → Return result → Validate input  
**D)** Return result → Validate input → Load banner  

<details><summary>💡 Answer</summary>

**B) Validate input → Load banner → Render → Return result**

Always validate first — fail fast before doing expensive work. Load the banner file (I/O operation) only after confirming the input is valid. Render only after confirming the file loaded. Return the result last.

</details>

---

### Q26: Where in `main.go` should you call `http.HandleFunc` — before or after `http.ListenAndServe`?

**A)** After — routes are registered while the server is running  
**B)** Before — all routes must be registered before the server starts listening  
**C)** It doesn't matter  
**D)** In a separate goroutine  

<details><summary>💡 Answer</summary>

**B) Before — all routes must be registered before the server starts**

`http.ListenAndServe` blocks. Any code after it won't run until the server stops. Register all handlers before calling it.

```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/ascii-art", asciiArtHandler)
log.Fatal(http.ListenAndServe(":8080", nil))
```

</details>

---

### Q27: Your template file is at `templates/index.html`. After `template.ParseFiles("templates/index.html")`, which name do you use to execute the template?

**A)** `"templates/index.html"`  
**B)** `"index.html"`  
**C)** `"index"`  
**D)** The first one automatically — you can call `tmpl.Execute(w, data)` directly  

<details><summary>💡 Answer</summary>

**D) `tmpl.Execute(w, data)` works directly when only one file is parsed**

When you call `template.ParseFiles("templates/index.html")`, the returned `*Template` is named `"index.html"` (just the filename, not the path). With a single file, `Execute` uses that template automatically. If you parse multiple files, you'd need `tmpl.ExecuteTemplate(w, "index.html", data)`.

</details>

---

### Q28: The spec says `GET /ascii-art` should return `400 Bad Request`. Why isn't it `405 Method Not Allowed`?

**A)** `405` doesn't exist  
**B)** The spec defines the exact status codes to use — follow the spec even if another code might seem more semantically correct  
**C)** `400` and `405` mean the same thing  
**D)** `405` is only for DELETE requests  

<details><summary>💡 Answer</summary>

**B) Follow the spec exactly**

`405 Method Not Allowed` would be the most semantically precise HTTP code here. However, the project spec explicitly says to use `400`. In real projects, follow the spec. The tests will check for the exact codes the spec defines — not what "should" be correct.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 26–28 ✅ | **Excellent.** Strong HTTP and Go web foundations — start immediately. |
| 22–25 ✅ | **Ready.** Review missed questions, especially template and form handling. |
| 17–21 ⚠️ | **Study first.** HTTP fundamentals and `html/template` need more work. |
| Below 17 ❌ | **Not ready.** Review HTTP methods, status codes, `net/http`, and HTML forms before starting. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q7 | HTTP methods, status codes 200/400/404/500, header ordering |
| Q8–Q13 | `http.ListenAndServe`, `HandleFunc`, `FormValue`, `http.Error`, `r.Method` |
| Q14–Q19 | `html/template` ParseFiles/Execute, `{{ if }}`, `<pre>` for whitespace |
| Q20–Q24 | HTML `<form>`, `method`/`action`, radio buttons, keeping form state |
| Q25–Q28 | Handler order of operations, route registration, spec compliance |