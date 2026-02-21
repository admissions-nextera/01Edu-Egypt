# 🎓 Go Web Development - Mastery Quiz
## HTTP Basics & Routing Final Assessment

**Time Limit:** 45 minutes  
**Passing Score:** 28/35 (80%)  
**Instructions:** No code execution allowed - predict output mentally!

---

## 📋 SECTION 1: OUTPUT PREDICTION (10 Questions)

### Q1: What's the output when you visit http://localhost:8080/?
```go
package main
import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `Hello, !`

**Explanation:** r.URL.Path is "/", so r.URL.Path[1:] is an empty string.
</details>

---

### Q2: What HTTP status code is sent?
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Success")
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `200`

**Explanation:** If WriteHeader() is not called, Go defaults to 200 OK when Write() is called.
</details>

---

### Q3: What's printed to browser when you visit /user?
```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Home")
    })
    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "User")
    })
    http.ListenAndServe(":8080", nil)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `User`

**Explanation:** /user is more specific than /, so the /user handler is matched.
</details>

---

### Q4: What does r.Method contain for this request?
```go
// Request: GET /api/users HTTP/1.1

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, r.Method)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `GET`

**Explanation:** r.Method contains the HTTP method as an uppercase string.
</details>

---

### Q5: What's the output?
```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(404)
    w.WriteHeader(200)
    fmt.Fprintf(w, "OK")
}
```
**Status Code Sent:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `404`

**Explanation:** WriteHeader() can only be called once. Subsequent calls are ignored. First call (404) wins.
</details>

---

### Q6: What query parameter value is retrieved?
```go
// URL: /search?q=golang&limit=10

func handler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    fmt.Fprintf(w, query)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `golang`

**Explanation:** r.URL.Query().Get("q") extracts the 'q' parameter value.
</details>

---

### Q7: Which handler is called for /api/users?
```go
func main() {
    http.HandleFunc("/api/", apiHandler)
    http.HandleFunc("/api/users", usersHandler)
    http.ListenAndServe(":8080", nil)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `usersHandler`

**Explanation:** Exact match /api/users is more specific than subtree pattern /api/, so usersHandler wins.
</details>

---

### Q8: What's in the response?
```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"status":"ok"}`)
}
```
**Content-Type header:** `_______________`  
**Body:** `_______________`

<details><summary>💡 Solution</summary>
**Content-Type:** `application/json`  
**Body:** `{"status":"ok"}`

**Explanation:** Headers must be set before writing body. Both are correctly sent.
</details>

---

### Q9: What's the output?
```go
// URL: /hello?name=Alice

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", 405)
        return
    }
    name := r.URL.Query().Get("name")
    fmt.Fprintf(w, "Hello, %s", name)
}
```
**For GET request:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `Hello, Alice`

**Explanation:** Method is GET, so the error check passes and name parameter is used.
</details>

---

### Q10: What status code is sent?
```go
func handler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Not Found", 404)
}
```
**Answer:** `_______________`

<details><summary>💡 Solution</summary>
**Answer:** `404`

**Explanation:** http.Error() is a helper that sets status code, Content-Type, and writes message.
</details>

---

## 📋 SECTION 2: CODE COMPLETION (10 Questions)

### Q11: Complete the basic handler
```go
func homeHandler(___ http.ResponseWriter, ___ *http.Request) {
    fmt.Fprintf(___, "Welcome Home")
}
```

<details><summary>💡 Solution</summary>
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome Home")
}
```
</details>

---

### Q12: Complete route registration
```go
func main() {
    http.___("/", homeHandler)
    http.___("/about", aboutHandler)
    http.ListenAndServe(":8080", ___)
}
```

<details><summary>💡 Solution</summary>
```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/about", aboutHandler)
http.ListenAndServe(":8080", nil)
```
</details>

---

### Q13: Complete method check
```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.___ != "___" {
        http.Error(w, "Method not allowed", 405)
        return
    }
    // Handle POST
}
```

<details><summary>💡 Solution</summary>
```go
if r.Method != "POST" {
    // or use: http.MethodPost
}
```
</details>

---

### Q14: Complete query parameter extraction
```go
// URL: /search?term=golang&page=2

func handler(w http.ResponseWriter, r *http.Request) {
    term := r.URL.___().Get("___")
    page := r.URL.___().Get("___")
}
```

<details><summary>💡 Solution</summary>
```go
term := r.URL.Query().Get("term")
page := r.URL.Query().Get("page")
```
</details>

---

### Q15: Complete header setting
```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.___().Set("___", "application/json")
    w.___().Set("___", "no-cache")
    fmt.Fprintf(w, `{"status":"ok"}`)
}
```

<details><summary>💡 Solution</summary>
```go
w.Header().Set("Content-Type", "application/json")
w.Header().Set("Cache-Control", "no-cache")
```
</details>

---

### Q16: Complete custom ServeMux
```go
func main() {
    mux := http.___()
    mux.___(", homeHandler)
    mux.___("/api", apiHandler)
    http.ListenAndServe(":8080", ___)
}
```

<details><summary>💡 Solution</summary>
```go
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/api", apiHandler)
http.ListenAndServe(":8080", mux)
```
</details>

---

### Q17: Complete redirect
```go
func oldHandler(w http.ResponseWriter, r *http.Request) {
    http.___(w, r, "/new-url", http.___)
}
```

<details><summary>💡 Solution</summary>
```go
http.Redirect(w, r, "/new-url", http.StatusFound)
// or http.StatusMovedPermanently (301)
```
</details>

---

### Q18: Complete file server
```go
func main() {
    fs := http.FileServer(http.___("./static"))
    http.Handle("/static/", http.StripPrefix("___", fs))
    http.ListenAndServe(":8080", nil)
}
```

<details><summary>💡 Solution</summary>
```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```
</details>

---

### Q19: Complete form value retrieval
```go
// POST form data: username=alice&email=alice@example.com

func handler(w http.ResponseWriter, r *http.Request) {
    username := r.___("username")
    email := r.___("email")
    fmt.Fprintf(w, "%s: %s", username, email)
}
```

<details><summary>💡 Solution</summary>
```go
username := r.FormValue("username")
email := r.FormValue("email")
// or r.PostFormValue() for POST-only
```
</details>

---

### Q20: Complete status code setting
```go
func handler(w http.ResponseWriter, r *http.Request) {
    if !isAuthenticated(r) {
        w.___(http.___)
        fmt.Fprintf(w, "Unauthorized")
        return
    }
    // ... rest of handler
}
```

<details><summary>💡 Solution</summary>
```go
w.WriteHeader(http.StatusUnauthorized)
// or w.WriteHeader(401)
```
</details>

---

## 📋 SECTION 3: BUG FIXING (5 Questions)

### Q21: Fix this code
```go
// BUG: Headers not being set!
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `{"status":"ok"}`)
    w.Header().Set("Content-Type", "application/json")
}
```

<details><summary>💡 Solution</summary>
**Bug:** Headers must be set BEFORE writing response body!

**Fix:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"status":"ok"}`)
}
```
</details>

---

### Q22: Fix this code
```go
// BUG: Can't get query parameters!
func handler(w http.ResponseWriter, r *http.Request) {
    search := r.Query["q"]  // Wrong!
    fmt.Fprintf(w, search)
}
```

<details><summary>💡 Solution</summary>
**Bug:** Incorrect method to get query parameters!

**Fix:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    search := r.URL.Query().Get("q")
    fmt.Fprintf(w, search)
}
```
</details>

---

### Q23: Fix this code
```go
// BUG: Server doesn't start!
func main() {
    http.HandleFunc("/", homeHandler)
    // Missing something!
}
```

<details><summary>💡 Solution</summary>
**Bug:** No ListenAndServe call!

**Fix:**
```go
func main() {
    http.HandleFunc("/", homeHandler)
    http.ListenAndServe(":8080", nil)
}
```
</details>

---

### Q24: Fix this code
```go
// BUG: Wrong handler matched!
func main() {
    http.HandleFunc("/api", apiHandler)
    http.HandleFunc("/api/users", usersHandler)
}

// Visiting /api/users calls apiHandler instead of usersHandler!
```

<details><summary>💡 Solution</summary>
**Bug:** Need trailing slash for subtree pattern!

**Fix:**
```go
func main() {
    http.HandleFunc("/api/", apiHandler)  // Trailing slash!
    http.HandleFunc("/api/users", usersHandler)
}
```
**Explanation:** Without trailing slash, "/api" only matches exactly "/api", not "/api/users". With "/api/", it becomes a subtree pattern, but more specific "/api/users" still takes precedence.
</details>

---

### Q25: Fix this code
```go
// BUG: Can't read request body!
func handler(w http.ResponseWriter, r *http.Request) {
    body := r.Body.ReadAll()  // Wrong!
    fmt.Fprintf(w, string(body))
}
```

<details><summary>💡 Solution</summary>
**Bug:** Body is io.ReadCloser, use io.ReadAll!

**Fix:**
```go
import "io"

func handler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading body", 500)
        return
    }
    defer r.Body.Close()
    fmt.Fprintf(w, string(body))
}
```
</details>

---

## 📋 SECTION 4: CONCEPT QUESTIONS (10 Questions)

### Q26: Multiple Choice
What package provides HTTP server functionality?
- A) http/server
- B) net/http
- C) web/http
- D) server/http

<details><summary>💡 Solution</summary>
**B) net/http**

Standard library package for HTTP.
</details>

---

### Q27: Multiple Choice
What's the default status code if you don't call WriteHeader()?
- A) 0
- B) 200
- C) 404
- D) 500

<details><summary>💡 Solution</summary>
**B) 200**

Go defaults to 200 OK.
</details>

---

### Q28: Multiple Choice
What is http.ResponseWriter?
- A) A struct
- B) An interface
- C) A function
- D) A constant

<details><summary>💡 Solution</summary>
**B) An interface**

It defines methods: Header(), Write(), WriteHeader().
</details>

---

### Q29: Multiple Choice
What does ServeMux do?
- A) Serves static files
- B) Routes requests to handlers
- C) Handles HTTP methods
- D) Manages sessions

<details><summary>💡 Solution</summary>
**B) Routes requests to handlers**

ServeMux is a request multiplexer/router.
</details>

---

### Q30: Multiple Choice
What's the correct handler signature?
- A) func(http.ResponseWriter, http.Request)
- B) func(http.ResponseWriter, *http.Request)
- C) func(*http.ResponseWriter, *http.Request)
- D) func(ResponseWriter, Request)

<details><summary>💡 Solution</summary>
**B) func(http.ResponseWriter, *http.Request)**

ResponseWriter is interface (no pointer), Request is pointer.
</details>

---

### Q31: True/False
You can call WriteHeader() multiple times to change the status code.

<details><summary>💡 Solution</summary>
**FALSE**

WriteHeader() can only be called once. Subsequent calls are ignored.
</details>

---

### Q32: True/False
Headers must be set before writing the response body.

<details><summary>💡 Solution</summary>
**TRUE**

Headers are sent before the body. Setting them after Write() has no effect.
</details>

---

### Q33: Fill in the blank
To get form value 'email', use: `r.___("email")`

<details><summary>💡 Solution</summary>
```go
r.FormValue("email")
// or r.PostFormValue("email") for POST only
```
</details>

---

### Q34: Multiple Choice
What does http.Handle() expect as the second parameter?
- A) A function
- B) An http.Handler interface
- C) A string
- D) A ServeMux

<details><summary>💡 Solution</summary>
**B) An http.Handler interface**

Must implement ServeHTTP(ResponseWriter, *Request).
</details>

---

### Q35: Multiple Choice
What's the difference between http.Error and manually setting status + writing?
- A) No difference
- B) http.Error sets Content-Type to text/plain
- C) http.Error adds a newline
- D) Both B and C

<details><summary>💡 Solution</summary>
**D) Both B and C**

http.Error sets Content-Type and adds trailing newline.
</details>

---

## 🎯 SCORING GUIDE

**Count your correct answers:**

### 35/35 - 🏆 PERFECT MASTER
- Ready for advanced routing & middleware!
- Consider building a REST API

### 32-34 - 🔥 EXCELLENT
- Very strong HTTP fundamentals
- Ready for templates & databases
- Review missed concepts

### 28-31 - ✅ PASS
- Good grasp of HTTP basics
- Ready for next topics
- Practice routing patterns

### 24-27 - ⚠️ BORDERLINE
- Need more practice
- Build more handlers
- Review Let's Go Ch 2

### Below 24 - 🔄 NEED REVIEW
- Re-study HTTP fundamentals
- Practice basic servers
- Take quiz again in 2-3 days

---

## 📊 TOPIC BREAKDOWN

**Check which areas need work:**

**Basic Handlers (Q1-3, Q11-12, Q23, Q26, Q30):**
- ☐ Master level (7/7)
- ☐ Good (5-6/7)
- ☐ Needs practice (4 or less)

**Request/Response (Q4-6, Q8-10, Q13-15, Q19-20, Q27-28, Q31-33):**
- ☐ Master level (14/15)
- ☐ Good (11-13/15)
- ☐ Needs practice (10 or less)

**Routing & ServeMux (Q7, Q16-18, Q24, Q29, Q34):**
- ☐ Master level (6/6)
- ☐ Good (5/6)
- ☐ Needs practice (4 or less)

**Common Mistakes (Q21-22, Q25, Q35):**
- ☐ Master level (4/4)
- ☐ Good (3/4)
- ☐ Needs practice (2 or less)

**HTTP Fundamentals (All sections):**
- ☐ Master level (33/35)
- ☐ Good (28-32/35)
- ☐ Needs practice (below 28)

---

## 💡 STUDY RECOMMENDATIONS

**Based on your score:**

### If you scored 28-35:
✅ **Move forward!**
- Ready for Quiz 3: Functions & Handlers
- Start building more complex servers
- Learn about templates next

### If you scored 24-27:
📚 **Light review needed**
- Re-read Let's Go Chapter 2
- Practice creating different handler types
- Review ServeMux patterns
- Take quiz again tomorrow

### If you scored below 24:
🔄 **Solid review required**
- Go through Let's Go Ch 2 again
- Read Go Web Programming Ch 3
- Build 5-10 simple servers
- Practice routing patterns
- Retake in 3-5 days

---

## 🎯 PRACTICE CHALLENGES

**After passing (28+), try these:**

### Challenge 1: Multi-Route Server
```go
// Build a server with:
// - Home page: "/"
// - About page: "/about"
// - API endpoint: "/api/status" (returns JSON)
// - 404 for everything else
```

### Challenge 2: Query Parameter Handler
```go
// Build /greet?name=Alice&title=Dr
// Should output: "Hello, Dr Alice!"
// Default to "Friend" if name missing
```

### Challenge 3: Method Router
```go
// Build /data endpoint that:
// - GET: returns data
// - POST: accepts data
// - Others: 405 Method Not Allowed
```

### Challenge 4: Static File Server
```go
// Serve files from ./public directory
// at /static/ URL path
```

---

## 📚 NEXT TOPICS TO STUDY

**After mastering HTTP basics:**

1. ✅ **Templates** - Dynamic HTML generation
2. ✅ **Forms** - Processing user input
3. ✅ **Database** - Storing data with PostgreSQL
4. ✅ **Middleware** - Request/response processing
5. ✅ **Authentication** - User sessions & security

**Keep coding! 💪 Every handler you write makes you better!**

---

## 📝 ANSWER SHEET

**Section 1 (Output Prediction):**
1. _____ 2. _____ 3. _____ 4. _____ 5. _____
6. _____ 7. _____ 8. _____ 9. _____ 10. _____

**Section 2 (Code Completion):**
11-20: Check your code against solutions

**Section 3 (Bug Fixing):**
21-25: Check your fixes

**Section 4 (Concepts):**
26. _____ 27. _____ 28. _____ 29. _____ 30. _____
31. _____ 32. _____ 33. _____ 34. _____ 35. _____

**Total Score: _____/35**

**Performance Breakdown:**
- Basic Handlers: _____/7
- Request/Response: _____/15
- Routing & ServeMux: _____/6
- Common Mistakes: _____/4
- HTTP Fundamentals: _____/35

**Ready to move forward? _____ (Yes/No)**

**Topics to review:** _____________________

---

## 🚀 BONUS: SELF-ASSESSMENT

**Rate yourself (1-5) on these skills:**

- [ ] Creating basic HTTP handlers: _____/5
- [ ] Understanding ResponseWriter: _____/5
- [ ] Working with Request object: _____/5
- [ ] Setting headers correctly: _____/5
- [ ] Routing with ServeMux: _____/5
- [ ] Handling query parameters: _____/5
- [ ] Error handling: _____/5
- [ ] HTTP status codes: _____/5

**Average:** _____/5

**If average < 3:** Review fundamentals  
**If average 3-4:** Practice more  
**If average > 4:** Ready to advance!