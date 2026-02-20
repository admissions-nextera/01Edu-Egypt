# 🎯 ASCII-Art-Stylize Prerequisites Quiz
## CSS Fundamentals · Static File Serving · Responsive Design · User Feedback

**Time Limit:** 45 minutes  
**Total Questions:** 25  
**Passing Score:** 20/25 (80%)

> ✅ Pass → You're ready to start ASCII-Art-Stylize  
> ⚠️ Also Required → ASCII-Art-Web must be fully working before you add styling

---

## 📋 SECTION 1: SERVING STATIC FILES IN GO (5 Questions)

### Q1: What does `http.FileServer(http.Dir("static"))` do?

**A)** Creates a directory named "static"  
**B)** Returns an HTTP handler that serves files from the `"static"` directory  
**C)** Serves a single file named "static"  
**D)** Watches the `"static"` directory for changes  

<details><summary>💡 Answer</summary>

**B) Returns an HTTP handler that serves files from the `"static"` directory**

`http.FileServer` creates a handler that reads files from a given directory and sends them as HTTP responses. Combined with `HandleFunc`, it allows the browser to load CSS, images, and other static assets.

</details>

---

### Q2: You register a file server like this:
```go
http.Handle("/static/", http.FileServer(http.Dir("static")))
```
The browser requests `/static/style.css`. What path does the file server look for on disk?

**A)** `static/style.css`  
**B)** `static/static/style.css`  
**C)** `style.css`  
**D)** `/static/style.css` (absolute path)  

<details><summary>💡 Answer</summary>

**B) `static/static/style.css` — which is wrong**

The handler receives the full path `/static/style.css` and looks for it inside `http.Dir("static")`, resulting in `static/static/style.css`. That file doesn't exist. This is why you need `http.StripPrefix`.

</details>

---

### Q3: What does `http.StripPrefix("/static/", handler)` do?

**A)** Removes `/static/` from all URLs in the entire app  
**B)** Wraps a handler so that `/static/` is stripped from the request path before passing it to the handler  
**C)** Serves files only from the `/static/` URL  
**D)** Prevents access to the `/static/` directory  

<details><summary>💡 Answer</summary>

**B) Wraps a handler — `/static/` is stripped from the request path before the inner handler sees it**

With `http.StripPrefix`:
```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
```
Browser requests `/static/style.css` → StripPrefix removes `/static/` → FileServer looks for `style.css` inside `http.Dir("static")` → finds `static/style.css` ✓

</details>

---

### Q4: The browser loads your page but CSS is not applied. You check DevTools Network and see `style.css` returns `404`. What are the two most likely causes?

**A)** The CSS syntax is wrong, or the colors are invalid  
**B)** The static file handler is not registered, or the path in `<link href="...">` doesn't match the registered route  
**C)** You need to restart the browser  
**D)** Go doesn't support CSS files  

<details><summary>💡 Answer</summary>

**B) The static file handler is not registered, or the `<link href>` path doesn't match**

Check both: (1) Is `http.Handle("/static/", ...)` registered in main before `ListenAndServe`? (2) Does `<link href="/static/style.css">` exactly match the registered route prefix? A mismatch in either causes 404.

</details>

---

### Q5: What is the correct HTML tag to link an external CSS file?

**A)** `<style src="/static/style.css">`  
**B)** `<css href="/static/style.css">`  
**C)** `<link rel="stylesheet" href="/static/style.css">`  
**D)** `<script src="/static/style.css">`  

<details><summary>💡 Answer</summary>

**C) `<link rel="stylesheet" href="/static/style.css">`**

`<link>` goes inside `<head>`. The `rel="stylesheet"` attribute tells the browser this is CSS. The `href` must match the URL path your Go file server is registered to handle.

</details>

---

## 📋 SECTION 2: CSS FUNDAMENTALS (8 Questions)

### Q6: What do CSS custom properties (variables) look like? How do you define and use them?

**A)** `$primary: #333; color: $primary;`  
**B)** `--primary: #333;` in `:root {}`, then `color: var(--primary);`  
**C)** `@variable primary #333;`  
**D)** `const primary = "#333";`  

<details><summary>💡 Answer</summary>

**B) `--primary: #333;` in `:root {}`, then `color: var(--primary);`**

```css
:root {
    --color-primary: #3a86ff;
    --spacing-base: 1rem;
}

button {
    background: var(--color-primary);
    padding: var(--spacing-base);
}
```

Defining all your design tokens in `:root` makes the whole design consistent and easy to change.

</details>

---

### Q7: What is the CSS box model? What four properties make it up?

**A)** Color, font, border, background  
**B)** Content, padding, border, margin  
**C)** Width, height, top, left  
**D)** Display, position, float, clear  

<details><summary>💡 Answer</summary>

**B) Content, padding, border, margin**

From inside out:
- **Content** — the actual text or element
- **Padding** — space inside the border (pushes content away from border)
- **Border** — the visible edge
- **Margin** — space outside the border (pushes other elements away)

`box-sizing: border-box` (in your reset) makes `width` include padding and border — use it always.

</details>

---

### Q8: How do you center a div horizontally on the page using CSS?

**A)** `text-align: center`  
**B)** `margin: 0 auto` with a defined `width` or `max-width`  
**C)** `position: center`  
**D)** `float: center`  

<details><summary>💡 Answer</summary>

**B) `margin: 0 auto` with a defined `max-width`**

```css
.container {
    max-width: 800px;
    margin: 0 auto;
    padding: 0 1rem;
}
```

`margin: auto` on left and right distributes available space equally on both sides. Without a `max-width`, the element fills the full width and centering has no effect.

</details>

---

### Q9: What does this CSS do?
```css
.form-row {
    display: flex;
    gap: 1rem;
    align-items: center;
}
```

**A)** Stacks elements vertically  
**B)** Places child elements side by side horizontally, with 1rem gap between them, vertically centered  
**C)** Centers the `.form-row` element on the page  
**D)** Creates a grid layout  

<details><summary>💡 Answer</summary>

**B) Places child elements side by side horizontally with gap and vertical centering**

Flexbox makes horizontal layouts easy:
- `display: flex` activates flex layout
- `gap: 1rem` adds space between children
- `align-items: center` vertically centers children within the row

</details>

---

### Q10: What is a CSS `@media` query used for?

**A)** Loading images  
**B)** Playing audio  
**C)** Applying different styles when the screen size or other conditions match  
**D)** Importing other CSS files  

<details><summary>💡 Answer</summary>

**C) Applying different styles when screen size or other conditions match**

```css
/* Desktop: side-by-side layout */
.form-row { display: flex; }

/* Mobile: stacked layout */
@media (max-width: 768px) {
    .form-row { flex-direction: column; }
}
```

Media queries are how you make a responsive design — different layouts for different screen sizes.

</details>

---

### Q11: What does `<meta name="viewport" content="width=device-width, initial-scale=1">` do?

**A)** Sets the page title  
**B)** Tells mobile browsers to use the device's real width instead of zooming out to show a "desktop" view  
**C)** Enables dark mode  
**D)** Loads a CSS framework  

<details><summary>💡 Answer</summary>

**B) Tells mobile browsers to use the device's real width, making responsive CSS work correctly**

Without this tag, mobile browsers render your page at ~980px and scale it down — your media queries fire at wrong sizes. With it, `max-width: 768px` actually triggers at 768 CSS pixels on a phone. Always include it in `<head>`.

</details>

---

### Q12: How do you add a hover effect to a button?

**A)** `button.hover { background: blue; }`  
**B)** `button:hover { background: blue; }`  
**C)** `button[hover] { background: blue; }`  
**D)** `@hover button { background: blue; }`  

<details><summary>💡 Answer</summary>

**B) `button:hover { background: blue; }`**

`:hover` is a CSS pseudo-class that applies when the user's mouse is over the element. For the active (click) state: `button:active { ... }`. These give users visual feedback that the button responds to interaction.

</details>

---

### Q13: What is the correct font for displaying ASCII art?

**A)** `font-family: serif`  
**B)** `font-family: sans-serif`  
**C)** `font-family: monospace`  
**D)** `font-family: cursive`  

<details><summary>💡 Answer</summary>

**C) `font-family: monospace`**

ASCII art is drawn assuming every character is exactly the same width. A monospace font (like `Courier New`, `Consolas`, or `monospace`) guarantees this. With a proportional font (serif, sans-serif), characters have different widths and the art looks broken.

</details>

---

## 📋 SECTION 3: DESIGN CONSISTENCY (5 Questions)

### Q14: You want a consistent design. What is the FIRST thing to define before writing any CSS rules?

**A)** The hover states of buttons  
**B)** Your design tokens — color palette, font(s), spacing scale — as CSS variables in `:root`  
**C)** The layout grid  
**D)** The mobile breakpoints  

<details><summary>💡 Answer</summary>

**B) Your design tokens as CSS variables in `:root`**

Design tokens are the atomic values of your system: primary color, background color, border radius, font size scale, spacing scale. Defining them first in `:root` means every component references variables rather than hardcoded values — change a variable and the whole page updates.

</details>

---

### Q15: Your form uses `border-radius: 8px` on the input, `border-radius: 4px` on the button, and `border-radius: 12px` on the result box. What design principle does this violate?

**A)** Accessibility  
**B)** Consistency — all interactive elements should use the same border-radius from your design tokens  
**C)** Performance  
**D)** Responsiveness  

<details><summary>💡 Answer</summary>

**B) Consistency — mismatched border-radius creates visual noise**

Pick one `--border-radius` variable and apply it everywhere. Inconsistent rounding makes the page feel unintentional. Every element that looks similar (inputs, buttons, cards) should share the same radius.

</details>

---

### Q16: How do you visually indicate which banner is currently selected (e.g., highlight the "shadow" radio button)?

**A)** Change the background of the entire form  
**B)** Use CSS `:checked` pseudo-class combined with a sibling selector to style the label  
**C)** Use JavaScript to add a class  
**D)** It's not possible with pure CSS  

<details><summary>💡 Answer</summary>

**B) Use CSS `:checked` pseudo-class with a sibling selector**

```css
input[type="radio"]:checked + label {
    background: var(--color-primary);
    color: white;
}
```

When a radio button is `:checked`, the `+` selector targets the immediately following `<label>`. No JavaScript needed. This provides immediate visual feedback about which banner is active.

</details>

---

### Q17: Your error message uses the same font color and background as the rest of the page. What is the problem?

**A)** It will display in the wrong position  
**B)** The user won't notice the error — it lacks visual distinction  
**C)** It will cause a CSS parse error  
**D)** No problem — errors don't need special styling  

<details><summary>💡 Answer</summary>

**B) The user won't notice the error — error messages must stand out visually**

Error messages should use a distinct color (typically red or orange), possibly a contrasting background, and perhaps a border or icon. The user must immediately know something went wrong — not have to hunt for the message.

```css
.error {
    color: #d93025;
    background: #fce8e6;
    border: 1px solid #d93025;
    padding: 0.75rem;
    border-radius: var(--border-radius);
}
```

</details>

---

### Q18: What is color contrast and why does it matter?

**A)** The difference between two colors' brightness — it determines readability for all users including those with visual impairments  
**B)** The number of colors used on the page  
**C)** How saturated the colors are  
**D)** Whether dark mode is supported  

<details><summary>💡 Answer</summary>

**A) The difference between two colors' brightness — determines readability**

WCAG (Web Content Accessibility Guidelines) requires at least 4.5:1 contrast ratio for normal text. Light gray text on white background might look elegant but fails accessibility. Use a contrast checker before finalizing your color choices.

</details>

---

## 📋 SECTION 4: CSS TRICKY CASES (4 Questions)

### Q19: The ASCII art result is very long and overflows its container. Which CSS property prevents it from breaking the page layout?

**A)** `overflow: hidden`  
**B)** `overflow-x: auto` or `overflow: auto` on the container  
**C)** `word-break: break-all`  
**D)** `white-space: nowrap`  

<details><summary>💡 Answer</summary>

**B) `overflow-x: auto`**

```css
.ascii-result {
    overflow-x: auto;
    white-space: pre;
    font-family: monospace;
}
```

`overflow-x: auto` adds a horizontal scrollbar only when needed. `white-space: pre` (already implied by `<pre>`) prevents line wrapping. This keeps the layout intact while letting users scroll to see wide art.

</details>

---

### Q20: What does `box-sizing: border-box` do and why should you set it globally?

**A)** Adds a visible border box around every element  
**B)** Changes the width model so `width` includes padding and border, making layout math intuitive  
**C)** Enables the CSS box model  
**D)** Prevents elements from having margins  

<details><summary>💡 Answer</summary>

**B) Makes `width` include padding and border**

Without it (default): `width: 300px` + `padding: 20px` = 340px total. With `border-box`: `width: 300px` means 300px total, regardless of padding. Set it globally in your reset:
```css
*, *::before, *::after {
    box-sizing: border-box;
}
```

</details>

---

### Q21: You add `transition: background 0.2s ease` to your button. What effect does this create?

**A)** The button pulsates continuously  
**B)** The background color change (on hover/active) animates smoothly over 0.2 seconds instead of snapping instantly  
**C)** The button disappears after 0.2 seconds  
**D)** A compile error in the CSS  

<details><summary>💡 Answer</summary>

**B) The hover color change animates smoothly over 0.2 seconds**

```css
button {
    background: var(--color-primary);
    transition: background 0.2s ease, transform 0.1s ease;
}
button:hover { background: var(--color-primary-dark); }
button:active { transform: scale(0.98); }
```

Transitions make interactions feel polished. 0.1–0.3s is the sweet spot — faster feels broken, slower feels sluggish.

</details>

---

### Q22: Your page looks fine on desktop but the text and buttons are tiny on mobile. You already have `@media (max-width: 768px)` rules. What are you probably missing?

**A)** A JavaScript resize listener  
**B)** The viewport meta tag: `<meta name="viewport" content="width=device-width, initial-scale=1">`  
**C)** More CSS  
**D)** A mobile-specific stylesheet  

<details><summary>💡 Answer</summary>

**B) The viewport meta tag**

Without `<meta name="viewport">`, mobile browsers render your page as if it were 980px wide and scale it down to fit the screen. Your `@media (max-width: 768px)` rule never triggers because the browser thinks the viewport is 980px. Adding the viewport tag fixes this immediately.

</details>

---

## 📋 SECTION 5: INTEGRATION (3 Questions)

### Q23: In what order should you add these to your HTML `<head>`?

```
A) <title>
B) <meta charset="utf-8">
C) <meta name="viewport" ...>
D) <link rel="stylesheet" href="/static/style.css">
```

**A)** D, B, C, A  
**B)** B, C, A, D  
**C)** A, B, C, D  
**D)** Any order — HTML `<head>` is order-independent  

<details><summary>💡 Answer</summary>

**B) B, C, A, D — charset first, then viewport, then title, then stylesheet**

The charset declaration should come first so the browser knows how to parse everything after it. Viewport comes early so the browser applies it before layout. Title before CSS is conventional. CSS last ensures it doesn't block parsing unnecessarily (though in `<head>` it blocks render regardless).

</details>

---

### Q24: You add a new `<link>` to a Google Font in your HTML but the font doesn't appear on the page. What is the most likely issue?

**A)** Go doesn't support Google Fonts  
**B)** You forgot to apply the font in CSS: `font-family: 'YourFont', sans-serif`  
**C)** The font file is too large  
**D)** Google Fonts require JavaScript  

<details><summary>💡 Answer</summary>

**B) You forgot to apply the font in CSS**

The `<link>` tag loads the font file, but you still need to tell CSS to use it:
```css
body {
    font-family: 'Inter', sans-serif;  /* Use the font */
}
```
The second value (`sans-serif`) is a fallback — used if the Google Font fails to load.

</details>

---

### Q25: After styling, a tester says "the ASCII art looks great in Chrome but broken in Firefox." What is the most likely cause?

**A)** Firefox doesn't support CSS  
**B)** A non-standard CSS property or value works in Chrome but not Firefox — check for vendor-specific prefixes or unsupported features  
**C)** Firefox caches CSS differently  
**D)** The server sends different CSS to Firefox  

<details><summary>💡 Answer</summary>

**B) A non-standard CSS property or Chrome-only feature**

Different browsers implement CSS features at different times. Check MDN for the property you're using — it shows browser compatibility. Also check for missing vendor prefixes (e.g., some gradient or transition syntax). Open DevTools in Firefox and look for any CSS warnings.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 23–25 ✅ | **Excellent.** Strong CSS foundation — start immediately. |
| 20–22 ✅ | **Ready.** Review the questions you missed before starting. |
| 15–19 ⚠️ | **Study first.** Focus on CSS fundamentals and static file serving. |
| Below 15 ❌ | **Not ready.** Work through a CSS fundamentals course before attempting this project. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q5 | `http.FileServer`, `http.StripPrefix`, linking CSS in HTML |
| Q6–Q13 | CSS variables, box model, flexbox, media queries, viewport, hover states, monospace |
| Q14–Q18 | Design tokens, consistency, `:checked` selector, error styling, color contrast |
| Q19–Q22 | Overflow handling, `box-sizing`, transitions, viewport meta tag |
| Q23–Q25 | HTML `<head>` order, Google Fonts, cross-browser compatibility |