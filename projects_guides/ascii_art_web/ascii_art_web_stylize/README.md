# ASCII-Art-Stylize Project Guide

> **Before you start:** This project builds on ascii-art-web. The server must be working before you add styling. Open 5 websites you like and inspect their CSS in DevTools. Notice what makes them feel clean, consistent, and easy to use.

---

## Objectives

By completing this project you will learn:

1. **CSS Fundamentals** — Selectors, properties, layout, typography, and color
2. **Linking CSS to HTML** — How a browser loads and applies an external stylesheet
3. **Responsive Design** — Making a page look correct on different screen sizes
4. **Serving Static Files** — Handling requests for CSS, images, and other assets in Go
5. **User Feedback** — Giving the user clear signals about what is happening
6. **Consistency** — Applying uniform design decisions across all states of the page

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art-Web (Completed)
- Working GET `/` and POST `/ascii-art` endpoints
- HTML templates loading and rendering correctly

### 2. CSS Basics
- How to link a CSS file to an HTML page with `<link>`
- Selectors: element, class, id
- The box model: margin, padding, border, width, height
- Flexbox or Grid for layout
- Search: **"CSS flexbox guide"**
- Search: **"CSS box model explained"**

### 3. Serving Static Files in Go
- `http.FileServer` — serving a directory of static assets
- `http.StripPrefix` — why you need it when serving from a subdirectory
- Search: **"golang serve static files http.FileServer"**

### 4. Responsive Design
- What `@media` queries are and how they work
- The `<meta name="viewport">` tag — why it matters on mobile
- Search: **"CSS media queries tutorial"**

**If any of these are unfamiliar, read about them before writing any code.**

---

## Project Structure

```
ascii-art-stylize/
├── main.go
├── handlers.go
├── banner.go
├── templates/
│   └── index.html
├── static/
│   ├── style.css       ← main stylesheet
│   └── (images, fonts if needed)
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Serve Static Files

**Goal:** The browser can load `static/style.css` from the server.

**Questions to answer before writing anything:**
- How do you register a route in Go that serves all files from a directory?
- What does `http.StripPrefix` do and why is it needed for `/static/`?
- What URL does the browser request when your HTML has `<link href="/static/style.css">`?

**Code Placeholder:**
```go
// main.go

func main() {
    // Register your existing routes (GET / and POST /ascii-art)

    // Register a file server for the /static/ path
    // Strip the /static/ prefix so it maps to the static/ directory
    // Use http.FileServer with http.Dir

    // Start the server
}
```

**Verify:**
- Add any text to `static/style.css` (e.g. `body { background: red; }`)
- Reload the page — the background should change
- Check DevTools Network tab — `style.css` should return 200

---

## Milestone 2 — Design Principles (No Code)

**This milestone has no code.** Before writing any CSS, answer these design questions for your site:

**Consistency:**
- What font will you use? (One or two at most)
- What is your color palette? (Pick 2–3 colors that work together)
- What border radius, spacing scale, and font size scale will you use?
- Write these down as CSS variables before you start styling.

**Responsiveness:**
- What should the layout look like on a desktop?
- What should it look like on a phone?
- Where are your breakpoints?

**User Feedback:**
- How does the user know the form was submitted?
- How does the user know an error occurred?
- What happens to the button during and after submit?
- Is the ASCII art result easy to read?

Sketching on paper before writing CSS saves hours. Do it.

---

## Milestone 3 — CSS Variables and Base Styles

**Goal:** Define your design tokens and apply base styles.

**Questions to answer:**
- What are CSS custom properties (variables) and how do you define and use them?
- Why is a CSS reset or normalizer useful?

**Code Placeholder:**
```css
/* static/style.css */

/* 1. Define CSS variables for your color palette, font, spacing, etc. */
:root {
    /* --color-primary: ... */
    /* --color-background: ... */
    /* --font-main: ... */
    /* --spacing-unit: ... */
}

/* 2. Base reset */
/* Remove default margin and padding, set box-sizing */

/* 3. Body styles */
/* Font, background color, min-height */

/* 4. Typography */
/* Headings, paragraphs, the <pre> block for ASCII art */
```

**Resources:**
- Search: **"CSS custom properties variables tutorial"**
- Search: **"CSS reset modern"**

**Verify:** The page uses your fonts and colors. Open DevTools and confirm the variables are applied.

---

## Milestone 4 — Layout

**Goal:** The form and result are laid out clearly. The page is usable on both desktop and mobile.

**Questions to answer:**
- Will you center the form on the page? How?
- Should the banner selector and text input be side by side or stacked?
- What happens to the layout when the screen is narrow?

**Code Placeholder:**
```css
/* Layout for the main container */
/* Center it, set max-width, add padding */

/* Form layout */
/* Stack or side-by-side depending on your design */

/* Banner selector styling */
/* Make it look better than browser defaults */

/* Submit button */
/* Size, color, hover state */

/* Result area (<pre> tag) */
/* Font size, background, padding, overflow handling */

/* Responsive breakpoint */
@media (max-width: 768px) {
    /* Adjust layout for small screens */
}
```

**Verify:**
- Resize your browser window from desktop width to mobile width — the layout adapts cleanly
- The form is usable on a narrow screen without horizontal scrolling

---

## Milestone 5 — Interactivity and Feedback

**Goal:** The user always knows what state the page is in.

**Questions to answer:**
- How do you visually highlight which banner is currently selected?
- How does the button change on hover and on click?
- When an error occurs (400, 500), does the page show a readable message styled differently from the result?
- Is the ASCII art result visually distinct from the rest of the page?

**Code Placeholder:**
```css
/* Button hover state */
/* Change background or add a shadow */

/* Button active state (while being clicked) */

/* Selected/active banner indicator */

/* Error message styling */
/* Different color, maybe a border or background */

/* ASCII art result box */
/* Monospace font, distinct background, scroll if content overflows */
```

```html
<!-- In your template, make sure error messages have a dedicated class -->
<!-- <div class="error-message">{{ .Error }}</div> -->
<!-- <pre class="ascii-result">{{ .Result }}</pre> -->
```

**Verify:**
- Hover over the button — it visually responds
- Submit with empty input — the error message is clearly visible and styled
- The ASCII art result stands out from the surrounding page

---

## Milestone 6 — Consistency Check

**Goal:** Every part of the page follows the same design language.

Go through this checklist yourself before submission:

- Does every element use your CSS variables instead of hardcoded colors?
- Do all interactive elements (inputs, selects, buttons) have the same border style?
- Is spacing consistent — same margin between similar elements?
- Does the page look intentional rather than accidental?
- Is every visible text readable against its background? (Check contrast)

**Resource:** Search: **"WCAG contrast checker"** — paste your colors and verify they pass.

---

## Debugging Checklist

- Does your CSS file return 404? Check that your `http.FileServer` route is registered before `ListenAndServe` and that the path in your HTML `<link>` matches exactly.
- Does styling not apply? Open DevTools → Network and confirm `style.css` is loaded with status 200. Open DevTools → Elements and check which styles are being applied.
- Does the ASCII art render without monospace spacing? Make sure the result is inside a `<pre>` tag with `font-family: monospace`.
- Does the layout break on small screens? Add `<meta name="viewport" content="width=device-width, initial-scale=1">` to your HTML `<head>`.

---

## Key Concepts

| Concept | What to Search |
|---|---|
| Serve static files in Go | "golang http.FileServer static files" |
| CSS variables | "CSS custom properties tutorial" |
| Flexbox layout | "CSS flexbox complete guide" |
| Media queries | "CSS media queries breakpoints" |
| CSS hover and active states | "CSS pseudo-classes hover active" |
| Color contrast | "WCAG color contrast checker" |

---

## Submission Checklist

- [ ] CSS file is served from `/static/` and loads correctly
- [ ] Page uses a consistent color palette defined as CSS variables
- [ ] Layout is clean and centered on desktop
- [ ] Layout adapts to mobile screen sizes
- [ ] Form elements (input, select, button) are styled consistently
- [ ] Submit button has a visible hover and active state
- [ ] Selected banner is visually indicated
- [ ] Error messages are styled distinctly from the result
- [ ] ASCII art displayed in monospace font with correct spacing
- [ ] ASCII art area handles long output without breaking the layout
- [ ] All text has sufficient color contrast against its background
- [ ] Page looks consistent — no mismatched fonts, colors, or spacing