# ASCII-Art-Web-Export Project Guide

> **Before you start:** This project builds on ascii-art-web. The server must be working. Open your browser's DevTools Network tab and watch what happens when you download a file from any website — inspect the response headers carefully.

---

## Objectives

By completing this project you will learn:

1. **HTTP Headers** — How headers control how a browser handles a response
2. **File Download via HTTP** — Sending a file as a response instead of an HTML page
3. **Content-Disposition** — The header that tells the browser to download rather than display
4. **Content-Type** — How to correctly identify the type of data being sent
5. **Content-Length** — Why the browser needs to know the file size before receiving it
6. **Export Formats** — Understanding how different file formats carry the same data

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art-Web (Completed)
- Working GET `/` and POST `/ascii-art` endpoints
- `render` function returns a string

### 2. HTTP Headers
- What is an HTTP response header?
- Read about these three headers before writing any code:
  - `Content-Type` — https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
  - `Content-Disposition` — https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
  - `Content-Length` — https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length
- Search: **"HTTP Content-Disposition attachment filename"**

### 3. Setting Headers in Go
- `w.Header().Set(key, value)` — must be called before `w.WriteHeader` or `w.Write`
- Search: **"golang http response headers"**

### 4. File Permissions
- What `0644` means for file permissions (read/write for user, read for others)
- Search: **"linux file permissions 644 explained"**

**If any of these are unfamiliar, read about them before writing any code.**

---

## Project Structure

```
ascii-art-web-export/
├── main.go
├── handlers.go
├── banner.go
├── templates/
│   └── index.html
├── static/
│   └── style.css
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Understand File Download via HTTP

**This milestone has no code.**

Open any website that lets you download a file. Open DevTools → Network → click the download link → click the response and look at its headers.

**Questions to answer:**
- What value does `Content-Disposition` have when a browser is told to download a file?
- What is the difference between `inline` and `attachment` in `Content-Disposition`?
- What does the `filename=` part of `Content-Disposition` do?
- What `Content-Type` value would you use for a plain text file?
- What happens in the browser if you set `Content-Disposition: attachment` but forget `Content-Type`?

Write the exact header values you will use before moving on.

---

## Milestone 2 — Add a Download Button to the Page

**Goal:** The page now has a button or link that triggers the export. Clicking it downloads the ASCII art as a file.

**Questions to answer:**
- Should the download be triggered by a new form, a link with `href`, or a POST to a new endpoint?
- What URL should the download button point to?
- How does the current rendered ASCII art get passed to the export endpoint? Via a form field? Via a query parameter? Via session state?
- What should happen if the user clicks download before rendering anything?

**Code Placeholder:**
```html
<!-- In your template -->

<!-- The existing form for rendering ASCII art stays as-is -->

<!-- Add a download button or form that: -->
<!--   - Sends the text and banner to a new endpoint /download -->
<!--   - OR submits the already-rendered result to /download -->
<!--   - Should only be enabled/visible when there is a result to download -->
```

---

## Milestone 3 — Create the Export Endpoint (`GET` or `POST /download`)

**Goal:** Requesting `/download` with valid input renders the ASCII art and sends it as a downloadable file with the correct headers.

**Questions to answer:**
- What HTTP method makes sense for this endpoint?
- What is the exact format of the `Content-Disposition` header for a file named `ascii-art.txt`?
- What `Content-Type` value is correct for a `.txt` file?
- How do you calculate `Content-Length` from a string in Go?
- In Go, what is the correct order: set headers, then write status, then write body?

**Code Placeholder:**
```go
// handlers.go

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Validate the request method

    // 2. Read text and banner from the form/query

    // 3. Validate input — return 400 if invalid

    // 4. Load the banner file — return 404 if not found

    // 5. Render the ASCII art to a string

    // 6. Set the response headers:
    //    Content-Type: text/plain; charset=utf-8
    //    Content-Disposition: attachment; filename="ascii-art.txt"
    //    Content-Length: (length of the rendered string in bytes)

    // 7. Write the rendered string as the response body
}
```

**Resources:**
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
- Search: **"golang http response headers Content-Disposition"**
- https://pkg.go.dev/net/http#ResponseWriter

**Verify:**
- Click the download button — the browser opens a Save dialog
- Open the downloaded file — it contains the correct ASCII art
- The file has the correct name (`ascii-art.txt` or your chosen name)

---

## Milestone 4 — File Permissions

**Goal:** The downloaded file is created with read and write permissions for the user.

**Questions to answer:**
- When the browser saves the file, who controls its permissions — the server or the browser?
- If you write the art to a temporary file on the server before sending it, what permissions must you use?
- What is the octal value for read+write for owner, read-only for group and others?

**Note:** If you are streaming the response directly (writing to `w` without creating a file on disk), research whether the spec requires a server-side file to be created. Read the spec again.

**Verify:**
```bash
ls -l ~/Downloads/ascii-art.txt
# Should show -rw-r--r-- or similar (readable and writable by owner)
```

---

## Milestone 5 — Error Handling

**Goal:** Every error state is handled gracefully and the user sees a clear message.

**Questions to answer:**
- What should the download endpoint return if the banner file is missing?
- What should happen if the text field is empty when the user tries to download?
- If rendering fails, should the server return a file with an error message or an HTTP error code?

**Code Placeholder:**
```go
// Extend downloadHandler with error cases:

    // If banner file not found:
    //   Return 404 Not Found with a readable error message (not a broken download)

    // If text is empty:
    //   Return 400 Bad Request

    // If rendering fails for any reason:
    //   Return 500 Internal Server Error
```

**Verify with curl:**
```bash
# Missing banner
curl -X POST http://localhost:8080/download -d "text=hello&banner=invalid" -v
# Should return 400 or 404

# Empty text
curl -X POST http://localhost:8080/download -d "text=&banner=standard" -v
# Should return 400
```

---

## Milestone 6 — Multiple Export Formats (Optional but Recommended)

**Goal:** The user can choose to download as `.txt`, `.html`, or another format.

**Questions to answer:**
- What `Content-Type` value is correct for an HTML file?
- What `Content-Disposition` filename should you use for each format?
- If you export as HTML, how do you wrap the ASCII art so it displays correctly in a browser?

**Code Placeholder:**
```go
// Extend downloadHandler to read a "format" parameter

    // If format is "txt":
    //   Content-Type: text/plain
    //   Body: raw rendered string

    // If format is "html":
    //   Content-Type: text/html
    //   Body: HTML page with <pre> containing the rendered art

    // Default: txt
```

---

## Debugging Checklist

- Does the browser display the file content instead of downloading it? You are missing `Content-Disposition: attachment` or it is set to `inline`.
- Does the downloaded file have garbled content? Check that your `Content-Type` includes `charset=utf-8`.
- Does the download fail silently? Check DevTools Network — look at the response headers and status code.
- Does the file save with the wrong name? Verify the `filename=` part of `Content-Disposition` is quoted correctly.
- Are headers not being set? Remember — `w.Header().Set(...)` must be called BEFORE `w.WriteHeader(...)` or `w.Write(...)`. Once you write the body, headers are locked.

---

## Key HTTP Headers Reference

| Header | Purpose | Example Value |
|---|---|---|
| `Content-Type` | Declares the data format | `text/plain; charset=utf-8` |
| `Content-Disposition` | Controls display vs download | `attachment; filename="ascii-art.txt"` |
| `Content-Length` | Size of the body in bytes | `1024` |

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `net/http` | Handle requests, set headers, write response | https://pkg.go.dev/net/http |
| `fmt` | Format error messages | https://pkg.go.dev/fmt |
| `strconv` | Convert int to string for Content-Length | https://pkg.go.dev/strconv |
| `os` | Write temp file if needed | https://pkg.go.dev/os |

---

## Submission Checklist

- [ ] Download button or link is visible on the page
- [ ] Clicking it downloads the ASCII art as a file
- [ ] `Content-Type` header is set correctly
- [ ] `Content-Disposition: attachment` header is set with a filename
- [ ] `Content-Length` header is set correctly
- [ ] Downloaded file contains the correct ASCII art
- [ ] Downloaded file has correct read/write permissions
- [ ] Empty input returns 400 before attempting download
- [ ] Missing banner returns 404
- [ ] All previous web endpoint tests still pass
- [ ] Only standard Go packages used