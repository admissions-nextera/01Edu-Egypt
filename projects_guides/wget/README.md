# Wget Project Guide

> **Before you start:** Run `man wget` and use the real `wget` command for every feature before you build it. You cannot recreate something you have never seen.

---

## Objectives

By completing this project you will learn:

1. **HTTP Protocol** — How HTTP requests and responses work, status codes, headers, and content types
2. **File System Operations** — Creating directories, writing files, walking directory trees
3. **I/O Streaming** — Reading data in chunks instead of all at once, which is essential for large files and progress tracking
4. **Concurrency** — Running multiple downloads at the same time with goroutines and WaitGroups
5. **Recursion** — Writing a recursive crawler that follows links without getting stuck in loops
6. **CLI Argument Parsing** — Building a clean flag parser for a real-world tool
7. **Rate Limiting** — Controlling how fast data flows using time and sleep
8. **URL Manipulation** — Parsing, resolving, and transforming URLs

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Go Basics
- Functions, structs, error handling
- `os.Args` — reading command-line arguments
- `os.Create`, `os.MkdirAll` — creating files and directories

### 2. HTTP in Go
- How to make a GET request with `net/http`
- What `http.Response` contains: status code, headers, body
- How to read a response body as a stream

### 3. Concurrency
- Goroutines (`go func()`)
- `sync.WaitGroup` — waiting for multiple goroutines to finish
- `sync.Mutex` — protecting shared data from simultaneous access

### 4. Time and Formatting
- `time.Now()` and `time.Since()`
- Go's reference time format — search: **"golang time format 2006"**

### 5. Strings and Paths
- `strings.HasPrefix`, `strings.HasSuffix`, `strings.TrimPrefix`
- `path/filepath.Base`, `filepath.Join`, `filepath.Rel`
- `net/url.Parse`, `url.ResolveReference`

**If any of these are unfamiliar, read about them before writing any code.**

---

## Project Structure

```
wget/
├── main.go        — entry point, argument parsing
├── download.go    — single file download logic
├── progress.go    — progress bar
├── mirror.go      — website mirroring
└── go.mod
```

---

## Milestone 1 — Download a Single File

**Goal:**
```
go run . https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
```
Must produce this exact output:
```
start at 2017-10-14 03:46:06
sending request, awaiting response... status 200 OK
content size: 56370 [~0.06MB]
saving file to: ./EMtmPFLWkAA8CIS.jpg
 55.05 KiB / 55.05 KiB [====================================] 100.00% 1.24 MiB/s 0s

Downloaded [https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg]
finished at 2017-10-14 03:46:07
```

**Questions to answer before writing anything:**
- How do you make a GET request and read the response in Go?
- What does `resp.ContentLength` return if the server does not report it?
- How do you read a response body in chunks to track how much has been downloaded?
- How do you print a line that updates itself in place without creating a new line?
- What format string produces `2017-10-14 03:46:06` in Go?

**Code Placeholder:**
```go
// download.go

func download(url string, outputName string, outputPath string, rateLimit int64) error {
    // 1. Record and print start time

    // 2. Make the HTTP GET request
    //    Print "sending request, awaiting response..."

    // 3. Check if status is 200 OK
    //    If not, print the status and return an error

    // 4. Print content size in bytes and MB

    // 5. Determine the save path (path + filename)
    //    Print "saving file to: ..."

    // 6. Open/create the destination file

    // 7. Read the response body in chunks
    //    After each chunk: write to file, update progress, apply rate limit if set

    // 8. Print blank line, "Downloaded [url]", and finish time
}
```

**Resources:**
- Search: **"golang net/http GET request"**
- Search: **"golang read http response body chunks"**
- Search: **"golang time format layout reference time"**
- https://pkg.go.dev/net/http

**Verify:** Run with a real URL and compare output character by character against the spec.

---

## Milestone 2 — Progress Bar

**Goal:** While downloading, print a single line that updates in place:
```
 55.05 KiB / 55.05 KiB [====================================] 100.00% 1.24 MiB/s 0s
```

**Questions to answer:**
- How do you overwrite the current terminal line without moving to a new one?
- How do you calculate download speed and estimated time remaining?
- How do you format bytes into KiB or MiB depending on size?

**Code Placeholder:**
```go
// progress.go

func printProgress(downloaded int64, total int64, elapsed time.Duration) {
    // 1. Calculate percentage

    // 2. Calculate download speed (bytes per second)

    // 3. Calculate estimated time remaining

    // 4. Build the progress bar string using = and spaces

    // 5. Format downloaded and total as KiB or MiB

    // 6. Print everything on one line using \r (not \n)
}
```

**Verify:** Watch the bar fill from 0% to 100% on a single line without scrolling.

---

## Milestone 3 — Non-200 Responses

**Goal:** If the server returns anything other than 200, print the status and exit. No file is created.

**Code Placeholder:**
```go
// In download.go, after making the request

    // Check status code
    // If not 200: print the status, return an error, do not proceed
```

**Verify:** Try a URL that returns 404. Your program should exit cleanly with a message.

---

## Milestone 4 — Flag `-O` (Save Under a Different Name)

**Goal:**
```
go run . -O=meme.jpg <url>
```
Saves the file as `meme.jpg` instead of the name from the URL.

**Questions to answer:**
- Can Go's built-in `flag` package handle `-O=value`? Test it before deciding.
- How do you extract the filename from a URL for when `-O` is not given?

**Code Placeholder:**
```go
// main.go or flags.go

type Config struct {
    URL        string
    OutputName string  // -O flag
    OutputPath string  // -P flag (next milestone)
    RateLimit  int64   // --rate-limit flag
    Background bool    // -B flag
    InputFile  string  // -i flag
    Mirror     bool    // --mirror flag
    Reject     []string
    Exclude    []string
    ConvertLinks bool
}

func parseArgs(args []string) Config {
    // Loop through args
    // For each arg:
    //   - Check if it matches a known flag pattern (e.g. strings.HasPrefix(arg, "-O="))
    //   - Extract the value after "="
    //   - Store in the correct field of Config
    //   - If it does not match any flag, treat it as the URL
}
```

**Verify:**
```bash
go run . -O=meme.jpg <url>
ls -l meme.jpg
```

---

## Milestone 5 — Flag `-P` (Save to a Directory)

**Goal:**
```
go run . -P=~/Downloads/ -O=meme.jpg <url>
```
Saves the file into the specified directory.

**Questions to answer:**
- How do you expand `~` to the real home directory in Go?
- What happens if the directory does not exist?
- How do you combine `-P` and `-O` into a final save path?

**Code Placeholder:**
```go
func resolveSavePath(cfg Config) string {
    // 1. Determine the filename:
    //    Use cfg.OutputName if set, otherwise extract from URL

    // 2. Determine the directory:
    //    Use cfg.OutputPath if set, otherwise use "./"

    // 3. Expand "~/" to the real home directory if present

    // 4. Join directory and filename into the final path
}
```

**Verify:**
```bash
go run . -P=/tmp/ -O=test.jpg <url>
ls -l /tmp/test.jpg
```

---

## Milestone 6 — Flag `--rate-limit`

**Goal:**
```
go run . --rate-limit=400k <url>
go run . --rate-limit=2M <url>
```
Limits download speed to the given rate.

**Questions to answer:**
- How do you convert `400k` and `2M` to bytes per second?
- You are reading in chunks. After each chunk, how do you calculate how long to sleep to stay at the target speed?

**Code Placeholder:**
```go
func parseRateLimit(s string) int64 {
    // Check if s ends with "k" → multiply by 1024
    // Check if s ends with "M" → multiply by 1024 * 1024
    // Otherwise treat as raw bytes per second
}

func throttle(bytesJustWritten int64, rateLimit int64) {
    // Calculate how long this chunk should have taken at the given rate
    // Sleep for that duration
}
```

**Verify:** Download a large file with `--rate-limit=50k`. The progress bar speed should stay near 50 KiB/s.

---

## Milestone 7 — Flag `-B` (Background Download)

**Goal:**
```
go run . -B <url>
Output will be written to "wget-log".
$             ← shell returns immediately
```
The download continues in the background. All output goes to `wget-log`.

**Questions to answer:**
- How do you re-launch your own program as a background process?
- How do you redirect a child process's stdout to a file?
- What is the difference between `cmd.Run()` and `cmd.Start()`?

**Code Placeholder:**
```go
func runInBackground(args []string) {
    // 1. Print "Output will be written to \"wget-log\"."

    // 2. Build a new args slice — same as current args but without "-B"

    // 3. Create (or truncate) the file "wget-log"

    // 4. Create an exec.Command for os.Args[0] with the new args
    //    Set its Stdout and Stderr to the log file

    // 5. Start the command (do not wait for it)

    // 6. Return — the parent process exits, child continues
}
```

**Resource:** Search: **"golang os/exec Start background process"**

**Verify:**
```bash
go run . -B <url>   # returns immediately
cat wget-log        # shows full output after a moment
```

---

## Milestone 8 — Flag `-i` (Multiple Concurrent Downloads)

**Goal:**
```
go run . -i=download.txt
```
Reads URLs from the file and downloads all of them at the same time.

**Questions to answer:**
- How do you read a text file line by line in Go?
- How do you wait for multiple goroutines to all finish?
- What is the goroutine loop variable capture bug and why will it affect you here?

**Code Placeholder:**
```go
func downloadMultiple(inputFile string, cfg Config) {
    // 1. Open and read the file line by line
    //    Collect all non-empty URLs into a slice

    // 2. Create a sync.WaitGroup

    // 3. For each URL:
    //    - Add 1 to the WaitGroup
    //    - Launch a goroutine that:
    //        a. Calls download() with that URL
    //        b. Calls wg.Done() when finished
    //    - Be careful: pass the URL as a function argument, do not capture it directly

    // 4. Wait for all goroutines to finish
}
```

**Resources:**
- Search: **"golang sync WaitGroup example"**
- Search: **"golang goroutine loop variable capture"** — read this before writing the loop

**Verify:** Put 3 URLs in a file. All 3 download simultaneously and all finish before the program exits.

---

## Milestone 9 — Mirror a Website (`--mirror`)

**Goal:**
```
go run . --mirror https://example.com
```
Downloads the entire site into a folder named after the domain, preserving directory structure.

This milestone is the most complex. Solve each step completely before moving to the next.

---

### Step 9.1 — Save One Page with Correct Path

**Questions to answer:**
- How do you turn `https://example.com/about/team` into a local path like `example.com/about/team.html`?
- What do you save when the URL path ends in `/` or is empty?

**Code Placeholder:**
```go
func buildLocalPath(host string, urlPath string) string {
    // 1. Start with the host as the root folder

    // 2. Append the URL path

    // 3. If the path ends in "/" or is empty, append "index.html"

    // 4. If the path has no file extension, append "/index.html"

    // 5. Return the final local path
}
```

---

### Step 9.2 — Extract Links from HTML

**Questions to answer:**
- How do you parse HTML token by token in Go?
- How do you resolve a relative URL like `/about` into `https://example.com/about`?

**Code Placeholder:**
```go
func extractLinks(body io.Reader, baseURL *url.URL) []string {
    // 1. Create an HTML tokenizer from the body

    // 2. Loop over tokens until ErrorToken (end of document)

    // 3. For each StartTag or SelfClosingTag token:
    //    - If tag is "a" or "link": look for "href" attribute
    //    - If tag is "img" or "script": look for "src" attribute
    //    - Resolve the found value against baseURL
    //    - Add to results if non-empty and not mailto: or javascript:

    // 4. Return all collected links
}
```

**Resources:**
- Search: **"golang x/net/html tokenizer example"**
- Search: **"golang url ResolveReference"**
- `go get golang.org/x/net/html`

---

### Step 9.3 — Recursive Crawl Without Infinite Loops

**Questions to answer:**
- How do you prevent visiting the same URL twice across concurrent goroutines?
- How do you make sure you only follow links on the same domain?

**Code Placeholder:**
```go
type Crawler struct {
    // base URL of the site being mirrored
    // map of visited URLs (needs a mutex to be safe across goroutines)
    // config (for -R, -X, --convert-links)
}

func (c *Crawler) crawl(rawURL string, wg *sync.WaitGroup) {
    defer wg.Done()

    // 1. Lock the mutex, check if URL was already visited
    //    If yes: unlock and return
    //    If no: mark as visited, unlock

    // 2. Parse the URL — if it is not on the same domain, return

    // 3. Check -X: if the path starts with an excluded directory, return

    // 4. Check -R: if the path ends with a rejected extension, return

    // 5. Make the HTTP GET request

    // 6. Build the local save path and create necessary directories

    // 7. If the response is HTML:
    //    - Read the full body
    //    - Save it to disk
    //    - Extract all links
    //    - For each link: wg.Add(1) and go c.crawl(link, wg)
    //    If not HTML:
    //    - Stream directly to disk

}
```

**Verify:** Run `go run -race . --mirror https://example.com` — fix any races before continuing.

---

### Step 9.4 — Flag `-R` (Reject Extensions)

The check belongs inside `crawl`. See Step 9.3 placeholder comment.

**Verify:**
```bash
go run . --mirror -R=jpg,png https://example.com
# No .jpg or .png files should appear in the output folder
```

---

### Step 9.5 — Flag `-X` (Exclude Paths)

The check belongs inside `crawl`. See Step 9.3 placeholder comment.

**Verify:**
```bash
go run . --mirror -X=/js https://example.com
# Nothing from /js should appear in the output folder
```

---

### Step 9.6 — Flag `--convert-links`

**Goal:** After mirroring, rewrite all absolute URLs in downloaded HTML files to point to local files instead.

**Questions to answer:**
- How do you walk all files in a directory tree in Go?
- How do you calculate a relative path from one local HTML file to another?

**Code Placeholder:**
```go
func convertLinks(rootDir string, baseHost string) error {
    // Walk every file in rootDir
    // For each .html file:
    //   1. Read its content
    //   2. Find all absolute URLs that belong to baseHost
    //   3. Calculate the relative local path from this file to the target file
    //   4. Replace the absolute URL with the relative path
    //   5. Write the updated content back to disk
}
```

**Resources:**
- Search: **"golang filepath Walk"**
- Search: **"golang filepath Rel"**

---

## Debugging Checklist

Go through this before asking for help:

- Have you used the real `wget` to see what the expected behavior is?
- Are you checking the error return of every function call?
- If goroutines are involved, have you run `go run -race .` to detect race conditions?
- If the mirror loops forever, is every access to your visited map inside a mutex lock?
- If the progress bar creates new lines instead of updating, are you using `\r` and not `\n`?
- If concurrent downloads produce garbled output, what is protecting your print calls?

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `net/http` | Make requests, read response body | https://pkg.go.dev/net/http |
| `net/url` | Parse and resolve URLs | https://pkg.go.dev/net/url |
| `os` | Create files, read args, get home dir | https://pkg.go.dev/os |
| `io` | Stream body to file | https://pkg.go.dev/io |
| `bufio` | Read input file line by line | https://pkg.go.dev/bufio |
| `sync` | WaitGroup and Mutex | https://pkg.go.dev/sync |
| `time` | Timestamps, sleep for rate limit | https://pkg.go.dev/time |
| `path/filepath` | Build paths, walk directories | https://pkg.go.dev/path/filepath |
| `strings` | Parse flags, check extensions | https://pkg.go.dev/strings |
| `os/exec` | Launch background process | https://pkg.go.dev/os/exec |
| `golang.org/x/net/html` | Parse HTML for mirroring | https://pkg.go.dev/golang.org/x/net/html |

---

## Submission Checklist

- [ ] Output format matches the spec exactly
- [ ] Non-200 responses print status and exit cleanly
- [ ] `-O` saves under the given name
- [ ] `-P` saves to the given directory
- [ ] `-P` and `-O` combined work correctly
- [ ] `--rate-limit=400k` and `--rate-limit=2M` both slow the download visibly
- [ ] `-B` returns the shell immediately, output written to `wget-log`
- [ ] `-i=file.txt` downloads all URLs concurrently
- [ ] `--mirror` downloads the site into a domain-named folder
- [ ] `--mirror -R=jpg,gif` skips rejected extensions
- [ ] `--mirror -X=/js` skips excluded paths
- [ ] `--mirror --convert-links` rewrites links for offline viewing
- [ ] No URL downloaded twice during mirror
- [ ] Mirror stays on the original domain only
- [ ] `go run -race .` reports no race conditions
- [ ] No crashes on bad URLs, missing files, or network errors
