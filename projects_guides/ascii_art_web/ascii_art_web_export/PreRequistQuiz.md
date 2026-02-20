# 🎯 ASCII-Art-Web-Export Prerequisites Quiz
## HTTP Headers · Content-Disposition · Content-Type · Content-Length · File Downloads

**Time Limit:** 40 minutes  
**Total Questions:** 25  
**Passing Score:** 20/25 (80%)

> ✅ Pass → You're ready to start ASCII-Art-Web-Export  
> ⚠️ Also Required → ASCII-Art-Web must be fully working before adding export

---

## 📋 SECTION 1: HTTP RESPONSE HEADERS (7 Questions)

### Q1: What is an HTTP response header?

**A)** The first line of HTML on a page  
**B)** A key-value pair sent by the server before the response body that tells the browser how to handle the response  
**C)** A Go struct field  
**D)** The URL of the request  

<details><summary>💡 Answer</summary>

**B) A key-value pair sent before the response body that tells the browser how to handle the response**

Headers control: what type of data is being sent (`Content-Type`), how to handle it (`Content-Disposition`), how large it is (`Content-Length`), caching, authentication, and more. The browser reads headers first to decide what to do with the body.

</details>

---

### Q2: How do you set a response header in Go's `net/http`?

**A)** `r.Header.Set("key", "value")`  
**B)** `w.Header().Set("key", "value")`  
**C)** `http.SetHeader(w, "key", "value")`  
**D)** `w.WriteHeader("key: value")`  

<details><summary>💡 Answer</summary>

**B) `w.Header().Set("key", "value")`**

`w` is the `http.ResponseWriter`. `w.Header()` returns the `http.Header` map that will be sent with the response. `r.Header` is the **incoming** request header — read-only in most cases. `w.Header()` is what you write to.

</details>

---

### Q3: What is the correct order of operations when writing a custom HTTP response?

**A)** `w.Write(body)` → `w.Header().Set(...)` → `w.WriteHeader(status)`  
**B)** `w.Header().Set(...)` → `w.WriteHeader(status)` → `w.Write(body)`  
**C)** `w.WriteHeader(status)` → `w.Header().Set(...)` → `w.Write(body)`  
**D)** The order doesn't matter  

<details><summary>💡 Answer</summary>

**B) Set headers → WriteHeader (status) → Write (body)**

Headers must be set BEFORE `WriteHeader` or `Write`. Once you call either of those, headers are sent to the client and any subsequent `Header().Set()` calls are silently ignored. Go will print a "superfluous response.WriteHeader call" warning if you try.

</details>

---

### Q4: What is the purpose of the `Content-Type` header?

**A)** It tells the server what type of data the client is sending  
**B)** It tells the browser what type of data the server is sending, so the browser knows how to handle it  
**C)** It controls the file download speed  
**D)** It sets the character encoding of the URL  

<details><summary>💡 Answer</summary>

**B) It tells the browser what type of data is being sent**

Without `Content-Type`, the browser guesses (MIME sniffing) — sometimes incorrectly. For a text file: `text/plain; charset=utf-8`. For HTML: `text/html; charset=utf-8`. For JSON: `application/json`. Explicit is always better.

</details>

---

### Q5: What is the `Content-Type` value for a plain text file with UTF-8 encoding?

**A)** `text/txt`  
**B)** `application/text`  
**C)** `text/plain; charset=utf-8`  
**D)** `plain/text; encoding=utf8`  

<details><summary>💡 Answer</summary>

**C) `text/plain; charset=utf-8`**

The format is `type/subtype; parameter=value`. For text files: `text/plain`. The `charset=utf-8` parameter tells the browser how to decode the bytes. For the ASCII art export, this is the correct value.

</details>

---

### Q6: What does `Content-Length` tell the browser?

**A)** The maximum request size allowed  
**B)** The size of the response body in bytes — allows the browser to show a progress bar and verify completeness  
**C)** The number of lines in the response  
**D)** The number of HTTP headers  

<details><summary>💡 Answer</summary>

**B) The size of the response body in bytes**

```go
content := render(banner, input)
w.Header().Set("Content-Length", strconv.Itoa(len([]byte(content))))
```

Note: use `len([]byte(content))` not `len(content)` — for pure ASCII they're equal, but `len(string)` counts bytes anyway in Go, so it's fine. Being explicit is safer.

</details>

---

### Q7: You call `w.Header().Set("Content-Type", "text/plain")` but the browser still displays the file as HTML. What might be wrong?

**A)** Go overrides your `Content-Type`  
**B)** You called `w.Write()` before setting the header — headers were already sent with the default `Content-Type: text/html`  
**C)** The browser doesn't respect `Content-Type`  
**D)** `text/plain` is not a valid content type  

<details><summary>💡 Answer</summary>

**B) You called `w.Write()` before setting the header**

If you write any body bytes before setting headers, Go auto-sends the headers with the default `Content-Type: text/html; charset=utf-8`. Your subsequent `Header().Set()` call is too late. Always set headers first.

</details>

---

## 📋 SECTION 2: CONTENT-DISPOSITION & FILE DOWNLOADS (7 Questions)

### Q8: What is the `Content-Disposition` header used for?

**A)** Specifying the character encoding  
**B)** Controlling whether the browser displays the content inline or triggers a file download  
**C)** Setting the response status code  
**D)** Specifying the server's timezone  

<details><summary>💡 Answer</summary>

**B) Controlling inline display vs file download**

Without this header, the browser decides based on `Content-Type`. With it, you explicitly control the behavior:
- `Content-Disposition: inline` — display in browser
- `Content-Disposition: attachment; filename="file.txt"` — download prompt

</details>

---

### Q9: What is the exact `Content-Disposition` header value that tells the browser to download the file and name it `"ascii-art.txt"`?

**A)** `Content-Disposition: download; name="ascii-art.txt"`  
**B)** `Content-Disposition: attachment; filename="ascii-art.txt"`  
**C)** `Content-Disposition: save; file="ascii-art.txt"`  
**D)** `Content-Disposition: file; filename=ascii-art.txt`  

<details><summary>💡 Answer</summary>

**B) `Content-Disposition: attachment; filename="ascii-art.txt"`**

`attachment` triggers the download prompt. `filename=` (with quoted value) sets the suggested save name. Some sources write it without quotes — quotes are recommended for filenames with spaces or special characters.

</details>

---

### Q10: What is the difference between `Content-Disposition: inline` and `Content-Disposition: attachment`?

**A)** `inline` is for images; `attachment` is for text  
**B)** `inline` displays the content directly in the browser; `attachment` triggers a Save dialog  
**C)** They are identical — the `filename=` part is what matters  
**D)** `inline` requires JavaScript; `attachment` does not  

<details><summary>💡 Answer</summary>

**B) `inline` displays in browser; `attachment` triggers a Save dialog**

If you set `Content-Disposition: inline; filename="ascii-art.txt"` with `Content-Type: text/plain`, the browser renders the text directly. Change `inline` to `attachment` and the browser prompts the user to save the file instead of displaying it.

</details>

---

### Q11: You open DevTools → Network → click a download link on a real website. What three headers should you find in the response that make the download work?

**A)** `Authorization`, `Cookie`, `Cache-Control`  
**B)** `Content-Type`, `Content-Disposition`, `Content-Length`  
**C)** `X-Download`, `File-Name`, `File-Size`  
**D)** `Transfer-Encoding`, `Keep-Alive`, `ETag`  

<details><summary>💡 Answer</summary>

**B) `Content-Type`, `Content-Disposition`, `Content-Length`**

These three headers work together:
- `Content-Type: text/plain; charset=utf-8` — what the data is
- `Content-Disposition: attachment; filename="ascii-art.txt"` — how to handle it (download)
- `Content-Length: 1024` — how big it is

</details>

---

### Q12: How do you calculate `Content-Length` for a rendered ASCII art string in Go?

**A)** `len(result)` always  
**B)** `len([]byte(result))` — bytes, not characters  
**C)** `strings.Count(result, "")` — count characters  
**D)** `result.Size()`  

<details><summary>💡 Answer</summary>

**B) `len([]byte(result))`**

HTTP `Content-Length` must be the number of **bytes** in the body. For pure ASCII content (which this project produces), `len(result)` and `len([]byte(result))` are identical. Using `[]byte(result)` explicitly is self-documenting and correct for any content.

```go
body := []byte(result)
w.Header().Set("Content-Length", strconv.Itoa(len(body)))
w.Write(body)
```

</details>

---

### Q13: The browser downloads the file but it's empty. What is the most likely cause?

**A)** The `Content-Length` was wrong  
**B)** `w.Write(body)` was called before setting the headers  
**C)** The render function returned an empty string because of an unhandled error that should have returned a 400/500 instead  
**D)** The browser cached the old response  

<details><summary>💡 Answer</summary>

**C) The render function returned empty due to an unhandled error**

An empty downloaded file means the body is empty. Check: was the banner file loaded? Did `render` produce output? Did you write the result to the response? Add logging before the write to see what `result` contains before sending.

</details>

---

### Q14: The file downloads with the name `download` instead of `ascii-art.txt`. What is wrong?

**A)** The browser always names files `download`  
**B)** The `Content-Disposition` header is missing or missing the `filename=` part  
**C)** The `Content-Type` header is wrong  
**D)** You need to use a different HTTP method  

<details><summary>💡 Answer</summary>

**B) `Content-Disposition` is missing or missing `filename=`**

When `Content-Disposition: attachment` is set without `filename=`, the browser uses the URL path as the filename (often resulting in just `download` for a `/download` endpoint). Always include `filename="ascii-art.txt"` with the attachment disposition.

</details>

---

## 📋 SECTION 3: IMPLEMENTING THE DOWNLOAD ENDPOINT (6 Questions)

### Q15: Should your `/download` endpoint use GET or POST? Why?

**A)** GET — because it's retrieving a file  
**B)** POST — because the text and banner values come from the form, which is a POST  
**C)** Either works identically  
**D)** PUT — because you're creating a new file  

<details><summary>💡 Answer</summary>

**B) POST — because form data is sent in the request body**

The download button/form sends `text` and `banner` values. Form submissions use POST (for data, not a simple URL navigation). A GET download would need the data in the URL query string — possible but exposes the content and has URL length limits.

</details>

---

### Q16: Your download form needs to send the same `text` and `banner` values that the render form sends. What is the cleanest implementation?

**A)** Use JavaScript to copy the values  
**B)** A second `<form>` with hidden inputs for `text` and `banner`, pointing to `/download` with `method="POST"`  
**C)** Add a second action to the existing form  
**D)** Use cookies to pass values between forms  

<details><summary>💡 Answer</summary>

**B) A second `<form>` with hidden inputs pointing to `/download`**

```html
<form method="POST" action="/download">
    <input type="hidden" name="text" value="{{ .InputText }}">
    <input type="hidden" name="banner" value="{{ .SelectedBanner }}">
    <button type="submit">Download as .txt</button>
</form>
```

The hidden inputs carry the current values. This requires your `PageData` struct to store `InputText` and `SelectedBanner` after the render step.

</details>

---

### Q17: What is the complete, correct sequence inside `downloadHandler`?

**A)** Validate → Render → Set headers → Write body  
**B)** Set headers → Validate → Render → Write body  
**C)** Render → Set headers → Validate → Write body  
**D)** Set headers → Write body → Validate → Render  

<details><summary>💡 Answer</summary>

**A) Validate → Render → Set headers → Write body**

Validate first (fail fast before doing I/O). Render after validation (only render valid input). Set headers after rendering (now you know `Content-Length`). Write body last. Never set headers before you have the content — you can't know `Content-Length` until you have the string.

</details>

---

### Q18: The download endpoint returns a 500 error. How should your handler return the error — as an HTTP error or as a downloaded file with the error message inside?

**A)** As a downloaded file — the browser expects a file  
**B)** As an HTTP error response (`http.Error(w, "message", 500)`) — never send a broken download  
**C)** Both simultaneously  
**D)** Silently fail and return an empty file  

<details><summary>💡 Answer</summary>

**B) As an HTTP error response**

Never set `Content-Disposition: attachment` and then write an error message as the body — the user would download a file containing "Internal Server Error". Return a proper HTTP error status code and message instead. The download only starts if the render succeeds.

</details>

---

### Q19: What does `strconv.Itoa(len(body))` produce for a 512-byte body?

**A)** `"512 bytes"`  
**B)** `"512"`  
**C)** `512`  
**D)** `0x200`  

<details><summary>💡 Answer</summary>

**B) `"512"` — a string**

`strconv.Itoa` converts an `int` to its decimal string representation. HTTP headers are strings, so you need the string `"512"`, not the integer `512`.

```go
w.Header().Set("Content-Length", strconv.Itoa(len(body)))
```

</details>

---

### Q20: A curl test for your download endpoint:
```bash
curl -X POST http://localhost:8080/download \
  -d "text=Hello&banner=standard" \
  -o output.txt -v
```
What should you look for in the verbose output (`-v`) to confirm the download is working?

**A)** `HTTP/1.1 200 OK` only  
**B)** `Content-Disposition: attachment; filename="ascii-art.txt"` and `Content-Type: text/plain` in the response headers  
**C)** The response body length  
**D)** The request method  

<details><summary>💡 Answer</summary>

**B) `Content-Disposition: attachment` and `Content-Type: text/plain` in response headers**

The `-v` flag shows all request and response headers. You want to verify:
1. Status is `200 OK`
2. `Content-Disposition: attachment; filename="ascii-art.txt"` is present
3. `Content-Type: text/plain; charset=utf-8` is present
4. `Content-Length` matches the file size

</details>

---

## 📋 SECTION 4: FILE PERMISSIONS & EDGE CASES (5 Questions)

### Q21: If you write the ASCII art to a temporary file on disk before sending it, what permission should the file have?

**A)** `0777` — everyone can read and write  
**B)** `0644` — owner read/write, group and others read-only  
**C)** `0600` — owner read/write only  
**D)** `0400` — read-only  

<details><summary>💡 Answer</summary>

**B) `0644`**

`0644` is the standard permission for text files:
- Owner: read (4) + write (2) = 6
- Group: read (4) = 4
- Others: read (4) = 4

`0777` gives everyone execute permission — unnecessary and a security risk for text files.

</details>

---

### Q22: In octal permission `0644`, what does the leading `0` mean in Go code?

**A)** It means the file will be hidden  
**B)** It is the Go syntax for an octal literal — `0644` means the number 420 in decimal  
**C)** It means "no permission for the owner"  
**D)** It is ignored  

<details><summary>💡 Answer</summary>

**B) It is Go's octal literal syntax — `0644` = 420 decimal**

In Go, integer literals starting with `0` are octal (base 8). `0644` = `6×64 + 4×8 + 4×1 = 384 + 32 + 4 = 420` in decimal. This is passed to `os.WriteFile` or `os.OpenFile` as a `fs.FileMode`.

</details>

---

### Q23: Should your download endpoint create a temporary file on the server, or stream the content directly from memory?

**A)** Always create a file — it's more reliable  
**B)** Streaming directly from memory is simpler and doesn't leave temp files on the server — prefer it unless the spec requires a server-side file  
**C)** Always use a file — HTTP requires it  
**D)** It depends on the file size  

<details><summary>💡 Answer</summary>

**B) Stream directly from memory — simpler and no cleanup needed**

```go
body := []byte(render(banner, input))
w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
w.Header().Set("Content-Length", strconv.Itoa(len(body)))
w.Write(body)
```

No file on disk. No cleanup. No permission to worry about on the server side. Read the spec carefully — if it requires creating a file on disk, do that; otherwise, stream.

</details>

---

### Q24: What should happen when the user clicks Download but hasn't rendered any art yet (the form is blank)?

**A)** Download a blank file  
**B)** Download a file with "no content"  
**C)** Return `400 Bad Request` — empty text is invalid input  
**D)** Return `404 Not Found`  

<details><summary>💡 Answer</summary>

**C) Return `400 Bad Request`**

Same validation as the render endpoint: empty text is client error → `400`. Never trigger a download for invalid input. The user should see an error response, not a broken or empty file.

</details>

---

### Q25: Your download works in Chrome but not in Firefox — Firefox opens the text in the browser instead of downloading it. What header is likely missing?

**A)** `Content-Length`  
**B)** `Content-Disposition: attachment`  
**C)** `Content-Type`  
**D)** `Transfer-Encoding`  

<details><summary>💡 Answer</summary>

**B) `Content-Disposition: attachment`**

Without `attachment`, different browsers make different decisions based on `Content-Type` alone. Chrome might download `text/plain` while Firefox displays it inline. `Content-Disposition: attachment` forces the download behavior consistently across all browsers. This is the most common download header bug.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 23–25 ✅ | **Excellent.** You understand HTTP headers deeply — start immediately. |
| 20–22 ✅ | **Ready.** Review missed questions before starting. |
| 15–19 ⚠️ | **Study first.** Focus on `Content-Disposition`, header ordering, and the download flow. |
| Below 15 ❌ | **Not ready.** Review HTTP headers and `net/http` response writing before starting. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q7 | HTTP response headers, `w.Header().Set()`, header ordering, `Content-Type` |
| Q8–Q14 | `Content-Disposition`, `inline` vs `attachment`, `Content-Length`, debugging downloads |
| Q15–Q20 | Download handler structure, hidden form inputs, error responses, curl testing |
| Q21–Q25 | File permissions, octal literals, streaming vs file, input validation, cross-browser |