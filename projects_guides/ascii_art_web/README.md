# ASCII-Art-Web Project Guide

> **Before you start:** This project builds on ascii-art. Your `render` function must return a string before you start here. Open a few websites and inspect their HTML source with DevTools. You cannot build a web interface you do not understand.

---

## Objectives

By completing this project you will learn:

1. **HTTP Server** — Starting a Go server, registering routes, and handling requests
2. **HTTP Methods** — The difference between GET and POST and when to use each
3. **HTML Templates** — Rendering dynamic data into HTML using Go's `html/template` package
4. **HTML Forms** — Sending user input from the browser to the server
5. **HTTP Status Codes** — Returning the correct status for each situation
6. **Static Files** — Serving CSS, images, and other assets from a Go server

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art (Completed)
- `loadBanner`, `getCharLines`, `render` — all working and `render` returns a string

### 2. HTTP Basics
- What is a GET request? What is a POST request?
- What is an HTTP status code? What do 200, 400, 404, and 500 mean?
- Search: **"HTTP methods GET vs POST explained"**
- Search: **"HTTP status codes explained"**

### 3. Go HTTP Package
- `http.ListenAndServe`
- `http.HandleFunc`
- `http.ResponseWriter` and `*http.Request`
- Search: **"golang net/http tutorial"**
- https://pkg.go.dev/net/http

### 4. HTML Basics
- `<form>`, `<input>`, `<textarea>`, `<select>`, `<button>`
- `action` and `method` attributes on a form
- How a form POST sends data to the server
- Search: **"HTML form POST request explained"**

### 5. Go Templates
- `html/template` — `ParseFiles`, `Execute`
- Passing data from Go to an HTML template
- Search: **"golang html/template tutorial"**
- https://pkg.go.dev/html/template

**If any of these are unfamiliar, read about them before writing any code.**

---

## Project Structure

```
ascii-art-web/
├── main.go
├── handlers.go         ← HTTP handler functions
├── banner.go           ← your existing render logic
├── templates/
│   ├── index.html
│   └── result.html     ← optional, depending on your approach
├── static/
│   └── style.css       ← added in ascii-art-stylize
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
├── README.md
└── go.mod
```

---

## Milestone 1 — Start the HTTP Server

**Goal:** The server starts and responds to requests.
```
go run .
```
Opens `http://localhost:8080` in your browser and shows something (even a blank page).

**Questions to answer before writing anything:**
- What function starts an HTTP server in Go?
- What does `http.HandleFunc` do and what arguments does it take?
- What port will you use? Is it hardcoded or configurable?
- What does a handler function's signature look like in Go?

**Code Placeholder:**
```go
// main.go

func main() {
    // Register the GET "/" route to its handler
    // Register the POST "/ascii-art" route to its handler

    // Print a message so you know the server started
    // Start the server on port 8080
    // Handle the error from ListenAndServe
}
```

**Resources:**
- https://gobyexample.com/http-servers
- https://pkg.go.dev/net/http#HandleFunc

**Verify:** `go run .` starts without crashing. `curl http://localhost:8080` returns something.

---

## Milestone 2 — Serve the Main HTML Page (GET `/`)

**Goal:** Visiting `http://localhost:8080` in a browser shows your main page with:
- A text input or textarea for the string to render
- A way to choose a banner (radio buttons, select, or similar)
- A submit button

**Questions to answer:**
- How do you load and parse an HTML file in Go using `html/template`?
- How do you send the rendered template as the HTTP response?
- What HTTP status code should a successful page load return?
- What status code should you return if the template file is not found?

**Code Placeholder:**
```go
// handlers.go

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Check that the request method is GET
    // If not, return 400 Bad Request

    // Check that the URL path is exactly "/"
    // If not, return 404 Not Found

    // Parse the template file from the templates/ directory
    // If parsing fails, return 500 Internal Server Error

    // Execute the template with w as the writer
    // If execution fails, return 500 Internal Server Error
}
```

```html
<!-- templates/index.html -->
<!-- A complete HTML page with: -->
<!-- - A <form> with method="POST" action="/ascii-art" -->
<!-- - A <textarea> or <input> for the text -->
<!-- - Radio buttons or <select> for banner choice (standard, shadow, thinkertoy) -->
<!-- - A <button type="submit"> -->
<!-- - A place to display the ascii art result (empty for now) -->
```

**Resources:**
- https://pkg.go.dev/html/template
- Search: **"HTML form radio buttons example"**
- Search: **"golang template ParseFiles Execute"**

**Verify:** The page loads in a browser. The form elements are visible. No result appears yet.

---

## Milestone 3 — Handle the POST Request (`POST /ascii-art`)

**Goal:** When the form is submitted, the server receives the text and banner choice, renders the ASCII art, and displays it.

**Questions to answer:**
- How do you read form values from a POST request in Go?
- What is `r.FormValue` and how do you use it?
- What should happen if the text field is empty?
- What should happen if the banner name is invalid or the file is not found?
- How do you pass the rendered ASCII art back to the HTML template to display it?

**Code Placeholder:**
```go
// handlers.go

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
    // Check that the method is POST
    // If not, return 400 Bad Request

    // Read "text" and "banner" from the form values

    // Validate: if text is empty, return 400 Bad Request
    // Validate: if banner is not one of the three valid names, return 400 Bad Request

    // Load the banner file
    // If file not found, return 404 Not Found
    // If any other error, return 500 Internal Server Error

    // Render the ASCII art using your render function

    // Pass the result to the template and render it
    // Set status 200 OK
}
```

**Template data structure:**
```go
// Define a struct to pass data to your template
type PageData struct {
    // The rendered ASCII art result
    // The original input text (so the form stays filled)
    // The selected banner (so the selection stays active)
    // Any error message to display
}
```

**Verify:**
- Type text in the form, choose a banner, click submit
- The ASCII art appears on the page
- Submitting an empty form does not crash the server

---

## Milestone 4 — Correct HTTP Status Codes

**Goal:** Every endpoint returns the right status code for every situation.

**Questions to answer:**
- How do you set a custom status code in Go before writing the response?
- What is the difference between `w.WriteHeader(404)` and `http.NotFound(w, r)`?
- If you call `w.WriteHeader` after writing the body, what happens?

**Status code map for this project:**

| Situation | Code |
|---|---|
| Successful page load or render | 200 |
| Wrong HTTP method | 400 |
| Invalid or missing form input | 400 |
| Template or banner file not found | 404 |
| URL path not registered | 404 |
| Any unhandled server error | 500 |

**Verify with curl:**
```bash
curl -o /dev/null -s -w "%{http_code}" http://localhost:8080/
# Should print 200

curl -o /dev/null -s -w "%{http_code}" http://localhost:8080/notexist
# Should print 404

curl -X GET -o /dev/null -s -w "%{http_code}" http://localhost:8080/ascii-art
# Should print 400 (wrong method for this endpoint)
```

---

## Milestone 5 — README.md

**Goal:** Create `README.md` in the project root with these four sections.

```markdown
# ASCII-Art-Web

## Description
...

## Authors
...

## Usage
...

## Implementation Details
...
```

**Questions to answer:**
- What does each section need to contain? Read the spec carefully.
- The Implementation Details section should describe the algorithm — how does your render pipeline work from input string to ASCII art output?

---

## Debugging Checklist

- Does the server crash on startup? Check that your template files exist at the exact path you are parsing.
- Does form submission return a blank page? Check that your handler is actually executing the template with data, not just writing an empty response.
- Do you get "superfluous response.WriteHeader call" in your logs? You called `w.WriteHeader` twice — once before and once inside a template execution or `http.Error` call.
- Does your 404 handler trigger for every unregistered path? Remember that `http.HandleFunc("/", ...)` matches all paths unless you check `r.URL.Path == "/"` explicitly.
- Does the ASCII art display in the browser with collapsed spaces? HTML collapses whitespace. Wrap the result in a `<pre>` tag.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `net/http` | Start server, handle requests | https://pkg.go.dev/net/http |
| `html/template` | Parse and execute HTML templates | https://pkg.go.dev/html/template |
| `os` | Read banner files | https://pkg.go.dev/os |
| `fmt` | Format error messages | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] `go run .` starts the server without errors
- [ ] `GET /` returns the main page with status 200
- [ ] Main page has text input, banner selector, and submit button
- [ ] `POST /ascii-art` renders the ASCII art and displays it
- [ ] ASCII art displayed inside `<pre>` tags so spacing is preserved
- [ ] Wrong method returns 400
- [ ] Empty text input returns 400
- [ ] Unknown banner returns 400 or 404
- [ ] Missing template file returns 500
- [ ] Unregistered paths return 404
- [ ] All three banners work correctly
- [ ] `README.md` has all four required sections
- [ ] Only standard Go packages used