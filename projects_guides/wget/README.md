# Wget Project Guide

## 📋 Project Overview
Build a command-line tool that replicates core functionalities of GNU wget - a utility for downloading files from the web. This project combines HTTP networking, concurrent programming, file system operations, HTML parsing, and command-line interface design. You'll learn how professional download managers work and implement features like progress tracking, rate limiting, background downloads, and website mirroring.

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **HTTP Protocol**: Making HTTP/HTTPS requests, handling responses
2. **File I/O**: Downloading and saving files efficiently
3. **Concurrency**: Downloading multiple files simultaneously
4. **Progress Tracking**: Real-time download progress with bars
5. **Rate Limiting**: Controlling download speed
6. **HTML Parsing**: Extracting links and resources from web pages
7. **Recursive Downloading**: Following links to mirror entire websites
8. **Command-Line Interfaces**: Parsing flags and arguments
9. **Error Handling**: Graceful handling of network errors
10. **File System Operations**: Creating directories, managing paths

---

## 📚 Prerequisites - Topics You Must Know

### 1. **HTTP Basics**
- HTTP methods: GET, POST
- HTTP status codes: 200 OK, 404 Not Found, 500 Server Error
- HTTP headers: Content-Length, Content-Type
- Request/Response cycle
- URLs and URL parsing

### 2. **Go HTTP Package**
- `net/http` package:
  - `http.Get()` - Make GET requests
  - `http.Client` - Custom HTTP clients
  - `http.Response` - Handle responses
  - `io.Copy()` - Copy response body to file
  - `io.Reader` - Reading data streams

### 3. **File Operations**
- `os` package:
  - `os.Create()` - Create files
  - `os.OpenFile()` - Open with permissions
  - `os.MkdirAll()` - Create directories
  - `filepath.Join()` - Construct paths
  - `filepath.Base()` - Get filename from path

### 4. **Concurrency**
- Goroutines: `go function()`
- Channels: Communication between goroutines
- WaitGroups: Waiting for goroutines to complete
- Mutexes: Preventing race conditions

### 5. **Command-Line Arguments**
- `os.Args` - Access arguments
- `flag` package:
  - Defining flags
  - Parsing flags
  - Getting flag values
- Custom flag parsing

### 6. **Time Operations**
- `time` package:
  - `time.Now()` - Current time
  - `time.Format()` - Format timestamps
  - `time.Since()` - Calculate duration
  - `time.Sleep()` - Delays

### 7. **HTML Parsing** (for mirroring)
- `golang.org/x/net/html` package:
  - Parsing HTML documents
  - Traversing DOM tree
  - Finding tags and attributes

### 8. **URL Operations**
- `net/url` package:
  - `url.Parse()` - Parse URLs
  - `url.ResolveReference()` - Resolve relative URLs
  - URL components: scheme, host, path

---

## 🌐 Understanding HTTP Downloads

### **How HTTP Download Works**

```
Client (Your Program)          Server
       |                          |
       |  1. HTTP GET Request    |
       |------------------------>|
       |                          |
       |  2. HTTP Response       |
       |    Status: 200 OK       |
       |    Content-Length: 1024 |
       |    Content-Type: image  |
       |<------------------------|
       |                          |
       |  3. Download Data       |
       |    (Stream of bytes)    |
       |<------------------------|
       |<------------------------|
       |<------------------------|
       |  4. Save to file        |
```

**Key Concepts**:
1. **Request**: Ask server for a resource (URL)
2. **Response**: Server replies with status and headers
3. **Body Stream**: Actual file content comes as stream
4. **Save**: Write stream to local file

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Basic HTTP Download** 📥

#### Step 1: Project Setup
```
wget/
├── main.go
├── downloader/
│   └── downloader.go
├── progress/
│   └── progress.go
├── utils/
│   └── utils.go
└── go.mod
```

#### Step 2: Initialize Module
```bash
go mod init wget
# Install HTML parsing library
go get golang.org/x/net/html
```

#### Step 3: Basic Download Function
In `downloader/downloader.go`:

```go
package downloader

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

// DownloadFile downloads a file from URL and saves it
func DownloadFile(url, filepath string) error {
    // 1. Make HTTP GET request
    // 2. Check response status
    // 3. Get Content-Length
    // 4. Create output file
    // 5. Copy response body to file
    // 6. Return any errors
}
```

**Implementation Steps**:

**A. Make HTTP Request**
```go
resp, err := http.Get(url)
if err != nil {
    return fmt.Errorf("request failed: %w", err)
}
defer resp.Body.Close()
```

**B. Check Status**
```go
if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("bad status: %s", resp.Status)
}
```

**C. Create File**
```go
out, err := os.Create(filepath)
if err != nil {
    return err
}
defer out.Close()
```

**D. Copy Data**
```go
_, err = io.Copy(out, resp.Body)
return err
```

**Test**:
```go
func main() {
    url := "https://example.com/file.txt"
    err := DownloadFile(url, "file.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

---

#### Step 4: Extract Filename from URL
Create utility function:

```go
func GetFilenameFromURL(url string) string {
    // Parse URL
    // Get path
    // Extract filename using filepath.Base()
    // Handle query parameters
}
```

**Implementation**:
```go
import (
    "net/url"
    "path/filepath"
    "strings"
)

func GetFilenameFromURL(urlStr string) string {
    u, err := url.Parse(urlStr)
    if err != nil {
        return "download"
    }
    
    // Get path, remove query
    path := strings.Split(u.Path, "?")[0]
    
    // Extract filename
    filename := filepath.Base(path)
    
    if filename == "" || filename == "/" {
        return "index.html"
    }
    
    return filename
}
```

**Test Cases**:
```
"https://example.com/file.zip" → "file.zip"
"https://example.com/path/to/image.jpg" → "image.jpg"
"https://example.com/" → "index.html"
"https://example.com/file?query=1" → "file"
```

---

### **Phase 2: Display Information** 📊

#### Step 5: Implement Timestamp Formatting
```go
func FormatTimestamp(t time.Time) string {
    // Format as "yyyy-mm-dd hh:mm:ss"
    // Use time.Format() with layout
}
```

**Go Time Format**:
```go
// Go uses reference time: Mon Jan 2 15:04:05 MST 2006
layout := "2006-01-02 15:04:05"
return t.Format(layout)
```

**Test**:
```go
now := time.Now()
fmt.Println("start at", FormatTimestamp(now))
// Output: start at 2024-01-15 14:30:45
```

---

#### Step 6: Display Download Information
```go
type DownloadInfo struct {
    URL           string
    Filename      string
    ContentLength int64
    Status        string
}

func DisplayInfo(info DownloadInfo) {
    // Print start time
    // Print status
    // Print content size (with MB/GB formatting)
    // Print filename
}
```

**Size Formatting**:
```go
func FormatSize(bytes int64) string {
    const (
        KB = 1024
        MB = KB * 1024
        GB = MB * 1024
    )
    
    switch {
    case bytes >= GB:
        return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
    case bytes >= MB:
        return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
    case bytes >= KB:
        return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
    default:
        return fmt.Sprintf("%d bytes", bytes)
    }
}
```

**Output Example**:
```
start at 2024-01-15 14:30:45
sending request, awaiting response... status 200 OK
content size: 56370 [~0.06MB]
saving file to: ./file.jpg
```

---

### **Phase 3: Progress Bar** 📈

#### Step 7: Implement Progress Tracking
Create a progress tracker:

```go
type ProgressReader struct {
    Reader      io.Reader
    Total       int64
    Current     int64
    StartTime   time.Time
    LastUpdate  time.Time
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
    // Read from underlying reader
    // Update current bytes
    // Display progress
    // Return bytes read
}
```

**Key Concepts**:
- Wrap `io.Reader` to track bytes read
- Update progress after each read
- Calculate percentage and speed
- Display progress bar

**Implementation**:
```go
func (pr *ProgressReader) Read(p []byte) (int, error) {
    n, err := pr.Reader.Read(p)
    pr.Current += int64(n)
    
    // Update progress display (throttle updates)
    if time.Since(pr.LastUpdate) > 100*time.Millisecond {
        pr.DisplayProgress()
        pr.LastUpdate = time.Now()
    }
    
    return n, err
}

func (pr *ProgressReader) DisplayProgress() {
    // Calculate percentage
    percent := float64(pr.Current) / float64(pr.Total) * 100
    
    // Calculate speed
    elapsed := time.Since(pr.StartTime).Seconds()
    speed := float64(pr.Current) / elapsed
    
    // Calculate remaining time
    remaining := float64(pr.Total-pr.Current) / speed
    
    // Display progress bar
    // Format: downloaded / total [======>    ] percent% speed remaining
}
```

---

#### Step 8: Create Progress Bar Display
```go
func CreateProgressBar(current, total int64, width int) string {
    // Calculate filled portion
    // Create bar with = for filled, spaces for empty
    // Add > at the end of filled portion
}
```

**Implementation**:
```go
func CreateProgressBar(current, total int64, width int) string {
    if total <= 0 {
        return ""
    }
    
    percent := float64(current) / float64(total)
    filled := int(percent * float64(width))
    
    bar := ""
    for i := 0; i < width; i++ {
        if i < filled-1 {
            bar += "="
        } else if i == filled-1 {
            bar += ">"
        } else {
            bar += " "
        }
    }
    
    return "[" + bar + "]"
}
```

**Example Output**:
```
 55.05 KiB / 55.05 KiB [==================================>] 100.00% 1.24 MiB/s 0s
```

---

### **Phase 4: Command-Line Flags** 🚩

#### Step 9: Parse Command-Line Arguments
```go
type Config struct {
    URL          string
    OutputFile   string
    OutputDir    string
    Background   bool
    RateLimit    int64  // bytes per second
    InputFile    string
    Mirror       bool
    Reject       []string
    Exclude      []string
    ConvertLinks bool
}

func ParseFlags() (*Config, error) {
    // Use flag package or custom parsing
    // Handle all flags: -B, -O, -P, --rate-limit, -i, --mirror, etc.
    // Validate arguments
    // Return config
}
```

**Flag Package Example**:
```go
import "flag"

func ParseFlags() *Config {
    cfg := &Config{}
    
    flag.StringVar(&cfg.OutputFile, "O", "", "output filename")
    flag.StringVar(&cfg.OutputDir, "P", ".", "output directory")
    flag.BoolVar(&cfg.Background, "B", false, "background download")
    // ... more flags
    
    flag.Parse()
    
    // Get URL from remaining args
    args := flag.Args()
    if len(args) > 0 {
        cfg.URL = args[0]
    }
    
    return cfg
}
```

**Custom Flag Parsing** (for `--rate-limit=400k`):
```go
func ParseRateLimit(s string) (int64, error) {
    // Parse "400k" or "2M" format
    // k = 1024 bytes/sec
    // M = 1024*1024 bytes/sec
}
```

**Test**:
```bash
go run . -O=output.jpg -P=~/Downloads/ https://example.com/file.jpg
```

---

### **Phase 5: Background Download** 🌙

#### Step 10: Implement Background Mode
```go
func DownloadInBackground(url string, config Config) error {
    // 1. Open log file "wget-log"
    // 2. Redirect stdout to log file
    // 3. Start download
    // 4. Write all output to log
}
```

**Implementation**:
```go
func RunBackgroundDownload(url string) error {
    fmt.Println("Output will be written to \"wget-log\".")
    
    // Create log file
    logFile, err := os.Create("wget-log")
    if err != nil {
        return err
    }
    defer logFile.Close()
    
    // Redirect output
    oldStdout := os.Stdout
    os.Stdout = logFile
    
    // Perform download
    err = Download(url)
    
    // Restore stdout
    os.Stdout = oldStdout
    
    return err
}
```

**Alternative** (using goroutine):
```go
func RunBackgroundDownload(url string) {
    go func() {
        // Download in goroutine
        // All output goes to log file
    }()
    
    fmt.Println("Output will be written to \"wget-log\".")
    // Main program exits, but goroutine continues
}
```

---

### **Phase 6: Rate Limiting** 🐌

#### Step 11: Implement Rate Limiter
Create a rate-limited reader:

```go
type RateLimitedReader struct {
    Reader    io.Reader
    RateLimit int64 // bytes per second
    LastRead  time.Time
}

func (r *RateLimitedReader) Read(p []byte) (int, error) {
    // Calculate how much we can read based on time elapsed
    // Sleep if we're reading too fast
    // Read from underlying reader
}
```

**Implementation**:
```go
func (r *RateLimitedReader) Read(p []byte) (int, error) {
    if r.RateLimit <= 0 {
        return r.Reader.Read(p)
    }
    
    // Calculate time since last read
    if !r.LastRead.IsZero() {
        elapsed := time.Since(r.LastRead)
        expectedDuration := time.Duration(int64(len(p)) * int64(time.Second) / r.RateLimit)
        
        if expectedDuration > elapsed {
            time.Sleep(expectedDuration - elapsed)
        }
    }
    
    n, err := r.Reader.Read(p)
    r.LastRead = time.Now()
    
    return n, err
}
```

**Algorithm**:
```
Rate: 400 KB/s = 400,000 bytes/sec
Read buffer: 8192 bytes

Time per buffer = 8192 / 400000 = 0.02048 seconds
If we read faster, sleep the difference
```

**Test**:
```bash
# Should take ~2.5 seconds for 1MB at 400k/s
go run . --rate-limit=400k https://example.com/1mb-file.zip
```

---

### **Phase 7: Multiple Downloads** 📚

#### Step 12: Read URLs from File
```go
func ReadURLsFromFile(filename string) ([]string, error) {
    // Open file
    // Read line by line
    // Trim whitespace
    // Skip empty lines
    // Return slice of URLs
}
```

**Implementation**:
```go
import "bufio"

func ReadURLsFromFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var urls []string
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && !strings.HasPrefix(line, "#") {
            urls = append(urls, line)
        }
    }
    
    return urls, scanner.Err()
}
```

---

#### Step 13: Download Multiple Files Concurrently
```go
func DownloadMultiple(urls []string) error {
    // Use WaitGroup to wait for all downloads
    // Start goroutine for each URL
    // Track completion
}
```

**Implementation**:
```go
import "sync"

func DownloadMultiple(urls []string) error {
    var wg sync.WaitGroup
    errors := make(chan error, len(urls))
    
    for _, url := range urls {
        wg.Add(1)
        
        go func(u string) {
            defer wg.Done()
            
            filename := GetFilenameFromURL(u)
            err := DownloadFile(u, filename)
            
            if err != nil {
                errors <- err
            } else {
                fmt.Printf("finished %s\n", filename)
            }
        }(url)
    }
    
    wg.Wait()
    close(errors)
    
    // Check for errors
    for err := range errors {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

**Key Concepts**:
- **WaitGroup**: Wait for all goroutines to finish
- **Goroutines**: Each download runs independently
- **Channel**: Collect errors from goroutines
- **Defer**: Ensure WaitGroup.Done() is called

**Test**:
```bash
echo "https://example.com/file1.zip" > download.txt
echo "https://example.com/file2.zip" >> download.txt
go run . -i=download.txt
```

---

### **Phase 8: Website Mirroring** 🪞

#### Step 14: Download and Parse HTML
```go
func DownloadHTML(url string) (string, error) {
    // Download HTML content
    // Return as string
}

func ParseHTML(html string) ([]string, error) {
    // Parse HTML using golang.org/x/net/html
    // Find all <a>, <link>, <img> tags
    // Extract href and src attributes
    // Return list of URLs
}
```

**HTML Parsing Example**:
```go
import "golang.org/x/net/html"

func ParseHTML(htmlContent string) ([]string, error) {
    doc, err := html.Parse(strings.NewReader(htmlContent))
    if err != nil {
        return nil, err
    }
    
    var urls []string
    var traverse func(*html.Node)
    
    traverse = func(n *html.Node) {
        if n.Type == html.ElementNode {
            // Check for <a href="...">
            if n.Data == "a" {
                for _, attr := range n.Attr {
                    if attr.Key == "href" {
                        urls = append(urls, attr.Val)
                    }
                }
            }
            
            // Check for <img src="...">
            if n.Data == "img" {
                for _, attr := range n.Attr {
                    if attr.Key == "src" {
                        urls = append(urls, attr.Val)
                    }
                }
            }
            
            // Check for <link href="...">
            if n.Data == "link" {
                for _, attr := range n.Attr {
                    if attr.Key == "href" {
                        urls = append(urls, attr.Val)
                    }
                }
            }
        }
        
        // Traverse children
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            traverse(c)
        }
    }
    
    traverse(doc)
    return urls, nil
}
```

---

#### Step 15: Resolve Relative URLs
```go
func ResolveURL(baseURL, relativeURL string) (string, error) {
    // Parse base URL
    // Parse relative URL
    // Resolve using url.ResolveReference
    // Return absolute URL
}
```

**Implementation**:
```go
func ResolveURL(baseURL, relativeURL string) (string, error) {
    base, err := url.Parse(baseURL)
    if err != nil {
        return "", err
    }
    
    rel, err := url.Parse(relativeURL)
    if err != nil {
        return "", err
    }
    
    // Resolve relative to base
    absolute := base.ResolveReference(rel)
    return absolute.String(), nil
}
```

**Examples**:
```
Base: https://example.com/page/index.html
Relative: image.jpg
Result: https://example.com/page/image.jpg

Base: https://example.com/page/index.html
Relative: /css/style.css
Result: https://example.com/css/style.css

Base: https://example.com/
Relative: ../other/file.js
Result: https://example.com/other/file.js
```

---

#### Step 16: Recursive Download (Mirroring)
```go
func MirrorWebsite(startURL string, options MirrorOptions) error {
    // 1. Create directory for website
    // 2. Keep track of visited URLs (avoid loops)
    // 3. Download page
    // 4. Parse HTML for links
    // 5. Filter links based on options (reject, exclude)
    // 6. Recursively download linked resources
    // 7. Save files in proper directory structure
}
```

**Data Structures**:
```go
type MirrorOptions struct {
    BaseURL      string
    OutputDir    string
    Reject       []string   // File extensions to skip
    Exclude      []string   // Paths to skip
    ConvertLinks bool
    MaxDepth     int
}

type Crawler struct {
    Visited   map[string]bool
    Queue     []string
    Options   MirrorOptions
    mu        sync.Mutex
}
```

**Algorithm**:
```
1. Start with initial URL
2. Add to queue
3. While queue not empty:
   a. Take URL from queue
   b. If already visited, skip
   c. Mark as visited
   d. Download resource
   e. If HTML, parse for links
   f. Filter links (reject, exclude)
   g. Add valid links to queue
   h. Save file to proper location
```

**Implementation Structure**:
```go
func (c *Crawler) Mirror() error {
    for len(c.Queue) > 0 {
        currentURL := c.Queue[0]
        c.Queue = c.Queue[1:]
        
        if c.IsVisited(currentURL) {
            continue
        }
        
        c.MarkVisited(currentURL)
        
        // Download resource
        content, err := c.Download(currentURL)
        if err != nil {
            continue
        }
        
        // Save to file
        filepath := c.URLToFilepath(currentURL)
        c.SaveFile(filepath, content)
        
        // If HTML, extract and queue links
        if c.IsHTML(currentURL) {
            links := c.ExtractLinks(content, currentURL)
            filtered := c.FilterLinks(links)
            c.AddToQueue(filtered)
        }
    }
    
    return nil
}
```

---

#### Step 17: Filter Links
```go
func (c *Crawler) FilterLinks(links []string) []string {
    var filtered []string
    
    for _, link := range links {
        // Check if should reject (file extension)
        if c.ShouldReject(link) {
            continue
        }
        
        // Check if should exclude (path)
        if c.ShouldExclude(link) {
            continue
        }
        
        // Check if same domain
        if !c.IsSameDomain(link) {
            continue
        }
        
        filtered = append(filtered, link)
    }
    
    return filtered
}
```

**Reject Implementation**:
```go
func (c *Crawler) ShouldReject(url string) bool {
    for _, ext := range c.Options.Reject {
        if strings.HasSuffix(url, ext) {
            return true
        }
    }
    return false
}
```

**Exclude Implementation**:
```go
func (c *Crawler) ShouldExclude(url string) bool {
    u, err := url.Parse(url)
    if err != nil {
        return false
    }
    
    for _, path := range c.Options.Exclude {
        if strings.HasPrefix(u.Path, path) {
            return true
        }
    }
    return false
}
```

---

#### Step 18: Create Directory Structure
```go
func (c *Crawler) URLToFilepath(url string) string {
    // Parse URL
    // Create directory structure based on URL path
    // Use domain as base directory
    // Example: https://example.com/css/style.css
    //       -> example.com/css/style.css
}
```

**Implementation**:
```go
func URLToFilepath(baseDir, urlStr string) string {
    u, err := url.Parse(urlStr)
    if err != nil {
        return filepath.Join(baseDir, "index.html")
    }
    
    // Use domain as directory
    domain := u.Host
    
    // Get path
    path := u.Path
    if path == "" || path == "/" {
        path = "/index.html"
    }
    
    // Combine
    fullPath := filepath.Join(baseDir, domain, path)
    
    // Create directories
    dir := filepath.Dir(fullPath)
    os.MkdirAll(dir, 0755)
    
    return fullPath
}
```

**Example**:
```
URL: https://www.example.com/css/style.css
Result: ./www.example.com/css/style.css

URL: https://www.example.com/
Result: ./www.example.com/index.html
```

---

#### Step 19: Convert Links for Offline Viewing
```go
func ConvertLinks(htmlContent, baseURL string) (string, error) {
    // Parse HTML
    // Find all links (href, src)
    // Convert absolute URLs to relative paths
    // Save modified HTML
}
```

**Algorithm**:
```
1. Parse HTML document
2. For each link found:
   a. If external, keep as-is
   b. If internal, convert to relative path
   c. Example: https://example.com/css/style.css → ../css/style.css
3. Reconstruct HTML with modified links
```

**Implementation Hint**:
```go
func ConvertToRelativePath(from, to string) string {
    // Calculate relative path from 'from' file to 'to' file
    // Use filepath.Rel() or custom logic
}
```

---

### **Phase 9: Testing & Edge Cases** 🧪

#### Step 20: Handle Edge Cases

**Network Errors**:
```go
func DownloadWithRetry(url string, maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        err := Download(url)
        if err == nil {
            return nil
        }
        
        // Wait before retry
        time.Sleep(time.Second * time.Duration(i+1))
    }
    return fmt.Errorf("failed after %d retries", maxRetries)
}
```

**Large Files**:
```go
// Use buffered reading/writing
// Don't load entire file into memory
buf := make([]byte, 32*1024) // 32KB buffer
io.CopyBuffer(dst, src, buf)
```

**Circular Links** (for mirroring):
```go
// Keep track of visited URLs
visited := make(map[string]bool)

if visited[url] {
    return // Already processed
}
visited[url] = true
```

**File Name Conflicts**:
```go
func GetUniqueFilename(path string) string {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return path
    }
    
    // Add number suffix: file.txt → file(1).txt
    ext := filepath.Ext(path)
    base := strings.TrimSuffix(path, ext)
    
    for i := 1; ; i++ {
        newPath := fmt.Sprintf("%s(%d)%s", base, i, ext)
        if _, err := os.Stat(newPath); os.IsNotExist(err) {
            return newPath
        }
    }
}
```

---

## 🐛 Common Issues and Solutions

### Issue 1: Progress Bar Flickering
**Problem**: Progress bar updates too frequently
**Solution**: Throttle updates to every 100ms
```go
if time.Since(lastUpdate) < 100*time.Millisecond {
    return
}
```

### Issue 2: Deadlock in Concurrent Downloads
**Problem**: WaitGroup never completes
**Solution**: Always call `wg.Done()` using defer
```go
go func() {
    defer wg.Done()
    // download code
}()
```

### Issue 3: Rate Limit Not Working
**Problem**: Download faster than limit
**Solution**: Calculate sleep time correctly based on bytes read

### Issue 4: Mirror Downloading Forever
**Problem**: Circular links or external links
**Solution**: Track visited URLs and check domain

### Issue 5: Wrong File Paths
**Problem**: Files saved in wrong directories
**Solution**: Use `filepath.Join()` and create directories with `os.MkdirAll()`

---

## 📋 Testing Checklist

**Basic Download**:
- [ ] Downloads file from URL
- [ ] Saves with correct filename
- [ ] Displays start/end time
- [ ] Shows status code
- [ ] Shows content size
- [ ] Shows progress bar

**Flags**:
- [ ] `-O` saves with custom name
- [ ] `-P` saves to custom directory
- [ ] `-B` runs in background, outputs to log
- [ ] `--rate-limit` limits download speed
- [ ] `-i` downloads multiple files

**Mirroring**:
- [ ] `--mirror` downloads website
- [ ] Creates proper directory structure
- [ ] `-R` rejects specified file types
- [ ] `-X` excludes specified paths
- [ ] `--convert-links` converts for offline viewing
- [ ] Handles relative URLs
- [ ] Avoids circular links

**Edge Cases**:
- [ ] Handles 404 errors gracefully
- [ ] Works with redirects
- [ ] Handles large files (>100MB)
- [ ] Handles network interruptions
- [ ] Prevents filename conflicts

---

## ✅ Submission Checklist

**Code Quality**:
- [ ] Well-organized package structure
- [ ] Clear function names
- [ ] Comments explain complex logic
- [ ] No hardcoded values
- [ ] Proper error handling
- [ ] Concurrent code is safe (no race conditions)

**Functionality**:
- [ ] All basic features work
- [ ] All flags implemented
- [ ] Progress bar accurate
- [ ] Mirroring works correctly
- [ ] Handles errors gracefully

**Testing**:
- [ ] Tested with small files
- [ ] Tested with large files
- [ ] Tested with slow connections
- [ ] Tested mirroring on real websites
- [ ] Tested all flag combinations

---

## 📖 Key Concepts Reference

### **HTTP Status Codes**
- 200: OK (success)
- 301/302: Redirect
- 404: Not Found
- 500: Server Error

### **Content-Length Header**
- Tells you file size
- Used for progress calculation
- May be missing (chunked transfer)

### **Goroutines vs Threads**
- Goroutines are lightweight
- Managed by Go runtime
- Thousands can run concurrently

### **Channels**
- Communication between goroutines
- Buffered vs unbuffered
- Close to signal completion

### **Mutexes**
- Protect shared data
- Use when multiple goroutines access same variable
- Lock/Unlock pattern

---

## 🚀 Pro Tips

1. **Start Simple**: Get basic download working first
2. **Test Early**: Test each feature before moving on
3. **Use Buffers**: Don't load entire files in memory
4. **Handle Errors**: Network can fail at any time
5. **Log Everything**: Especially for background downloads
6. **Rate Limiting**: Use time-based throttling
7. **Progress Updates**: Throttle to avoid performance issues
8. **URL Parsing**: Always validate and sanitize URLs
9. **Concurrent Safety**: Use mutexes for shared state
10. **Recursion Limits**: Set max depth for mirroring

---

## 💡 Extension Ideas

After completing basic requirements:

1. **Resume Downloads**: Support partial downloads
2. **Parallel Chunks**: Download file in multiple parts
3. **Better UI**: Use terminal UI library for better display
4. **Config File**: Read settings from file
5. **Compression**: Support gzip, deflate
6. **Authentication**: Handle HTTP basic auth
7. **Cookies**: Maintain session cookies
8. **Robots.txt**: Respect website crawling rules
9. **Sitemap**: Use sitemap.xml for mirroring
10. **Database**: Store download history in SQLite

---

## 📚 Learning Resources

**HTTP & Networking**:
- [MDN HTTP Guide](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [HTTP Status Codes](https://httpstatuses.com/)
- [Go net/http Package](https://pkg.go.dev/net/http)

**Concurrency**:
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)

**HTML Parsing**:
- [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html)
- [HTML Parsing Tutorial](https://www.alexedwards.net/blog/parsing-html-with-go)

**File System**:
- [Go os Package](https://pkg.go.dev/os)
- [filepath Package](https://pkg.go.dev/path/filepath)

---

## 🎓 Implementation Phases Summary

**Week 1**: Basic download + progress bar + flags (-O, -P)
**Week 2**: Background download (-B) + rate limiting + multiple files (-i)
**Week 3**: HTML parsing + basic mirroring
**Week 4**: Advanced mirroring (reject, exclude, convert-links) + testing

---

## 🔍 Debugging Strategies

**Network Issues**:
```go
// Enable verbose HTTP logging
http.DefaultTransport.(*http.Transport).DisableKeepAlives = false
// Check raw HTTP traffic
```

**Progress Bar Problems**:
```go
// Log progress calculations
fmt.Fprintf(os.Stderr, "Debug: current=%d total=%d\n", current, total)
```

**Concurrent Issues**:
```go
// Use -race flag
go run -race . <args>
```

**Mirroring Issues**:
```go
// Log each URL processed
fmt.Println("Processing:", url)
fmt.Println("Links found:", len(links))
```

---

**Remember**: wget is a complex tool built over many years. Your goal is to learn the concepts, not replicate every feature. Focus on understanding HTTP, concurrency, and file operations. Start simple and build up! 🌐💻