# 🎯 Groupie Tracker Visualizations Prerequisites Quiz
## Shneiderman's Rules · CSS Design Systems · Animations · Accessibility · Empty/Error States

**Time Limit:** 50 minutes  
**Total Questions:** 27  
**Passing Score:** 21/27 (78%)

> ✅ Pass → You're ready to start Groupie Tracker Visualizations  
> ⚠️ This project is about design quality, not new Go features. If you score 21–23, do the design audit milestone on paper before writing any CSS.

---

## 📋 SECTION 1: SHNEIDERMAN'S 8 GOLDEN RULES (8 Questions)

### Q1: Shneiderman's Rule 1 is "Strive for Consistency." What does this mean concretely for your site?

**A)** Use the same background color everywhere  
**B)** Every button, spacing value, font, border-radius, and color must come from the same design system — no element should look like it belongs to a different site  
**C)** Every page must have the same layout  
**D)** All text must be the same size  

<details><summary>💡 Answer</summary>

**B) Every element shares the same design tokens — nothing looks out of place**

Consistency means: if buttons have `border-radius: 8px` on the home page, they have `8px` on the detail page. If the primary color is `#3a86ff`, that exact value (via a CSS variable) is used everywhere — not `#3a85ff` on one page and `#4090ff` on another. Use CSS variables as the single source of truth.

</details>

---

### Q2: Rule 3 is "Offer Informative Feedback." Which of your site's states currently has NO feedback?

**A)** A successfully loaded page  
**B)** The page while artist data is loading for the first time, an empty filter/search result, and a server error page  
**C)** When the user hovers a card  
**D)** When the URL is correct  

<details><summary>💡 Answer</summary>

**B) Loading state, empty state, and error state**

Most first implementations have: loading? (blank page until data arrives), empty? (nothing renders, no explanation), error? (Go's default plain text error). All three need designed states: a spinner or skeleton, a friendly "No results" message, and a styled error page.

</details>

---

### Q3: Rule 5 is "Offer Simple Error Handling." A user filters to "creation date 2050–2060" and gets zero results. Which response follows this rule?

**A)** Display a blank page with no explanation  
**B)** Show a 404 error page  
**C)** Show a friendly empty-state message: "No artists match your filters. Try adjusting the date range."  
**D)** Crash with a 500 error  

<details><summary>💡 Answer</summary>

**C) A friendly, actionable empty-state message**

Error handling is not just about server errors — it's about any state where the user doesn't get what they expected. The message should: tell them what happened, suggest how to fix it, and not make them feel stupid. Never show a blank page when zero results is a valid state.

</details>

---

### Q4: Rule 7 is "Support Internal Locus of Control." What does this mean for your site's navigation?

**A)** The user must follow a strict order of pages  
**B)** The user should always feel in control: a "Back to artists" link on every detail page, an active indicator on nav links, and no dead ends  
**C)** Use only back-navigation  
**D)** Navigation should be hidden until needed  

<details><summary>💡 Answer</summary>

**B) User always knows where they are and how to get anywhere**

Feeling "trapped" on a page violates rule 7. Concretely: every detail page has a clear "← Back" or home link, the active page is highlighted in the nav, and there are no pages where the user must use the browser back button to continue.

</details>

---

### Q5: Rule 8 is "Reduce Short-Term Memory Load." Which design pattern directly addresses this rule?

**A)** Using small font sizes  
**B)** Showing all relevant information on the current screen instead of requiring the user to remember information from a previous screen  
**C)** Adding more pages  
**D)** Using animations  

<details><summary>💡 Answer</summary>

**B) Show all relevant information without requiring memory of previous screens**

On the artist detail page: show the artist name prominently at the top (not just a back-navigation breadcrumb) so the user knows whose page they're on. On filter results: show the active filters above the results so the user doesn't need to remember what they filtered by.

</details>

---

### Q6: Rule 2 is "Enable Frequent Users to Use Shortcuts." Which is a realistic shortcut for your site?

**A)** A command-line interface  
**B)** Clicking an artist name anywhere it appears navigates directly to their page; pressing Escape closes the search dropdown  
**C)** Memorizing keyboard shortcuts from a documentation page  
**D)** Only shows shortcuts to admin users  

<details><summary>💡 Answer</summary>

**B) Clickable links and keyboard shortcuts for common actions**

Shortcuts don't have to be complex: every artist card is fully clickable (not just the "View" button), pressing `Escape` closes the search bar, the search bar focuses when pressing `/`. These small interactions add up to a faster experience for returning users.

</details>

---

### Q7: Rule 4 is "Design Dialogue to Yield Closure." Your filter form submits and the page reloads. What must be visible to indicate the filter was applied?

**A)** A loading spinner that never goes away  
**B)** The active filter values must still be visible in the form controls AND ideally an indicator showing "X results found" or which filters are active  
**C)** A success toast notification  
**D)** Nothing — the results are self-explanatory  

<details><summary>💡 Answer</summary>

**B) Active filters retained + result count**

"Closure" means the user knows the action completed. If the filter form resets after submission, the user doesn't know what's applied. Retained filter values plus a result count ("Showing 5 of 52 artists") gives clear closure — the action happened, this is the outcome.

</details>

---

### Q8: Rule 6 is "Permit Easy Reversal of Actions." On your artist detail page, what is the minimum implementation of this rule?

**A)** A redo button  
**B)** A clearly visible "Back to all artists" link that returns the user to the unfiltered home page  
**C)** The browser back button is sufficient — no extra code needed  
**D)** An undo history panel  

<details><summary>💡 Answer</summary>

**B) A visible "Back" or home link on every page**

While the browser back button technically works, relying on it violates rule 6 — the user shouldn't need external browser controls to navigate your site. An explicit link is always better. Bonus: if filters were active, link back to the filtered state too.

</details>

---

## 📋 SECTION 2: CSS DESIGN SYSTEMS (7 Questions)

### Q9: What is a CSS design token?

**A)** A security token for CSS  
**B)** An atomic, named design decision stored as a CSS variable — color, spacing, typography, radius values that every component references  
**C)** A class name  
**D)** A CSS media query  

<details><summary>💡 Answer</summary>

**B) Named, atomic design decisions stored as CSS variables**

```css
:root {
    --color-primary:    #3a86ff;
    --color-bg:         #0f172a;
    --color-surface:    #1e293b;
    --space-md:         1rem;
    --space-lg:         1.5rem;
    --radius:           8px;
    --shadow-card:      0 4px 6px -1px rgba(0,0,0,0.1);
    --transition:       0.2s ease;
}
```

Every component that needs a color says `color: var(--color-primary)` — never `color: #3a86ff`. This is the design system's single source of truth.

</details>

---

### Q10: A developer hardcoded `color: #3a86ff` in 12 different places. The designer changes the primary color to `#4f46e5`. How many changes does this require vs using a CSS variable?

**A)** Same work — find and replace is easy  
**B)** 12 changes with hardcoded values vs 1 change (`--color-primary: #4f46e5` in `:root`) with variables  
**C)** Variables are harder to change  
**D)** The browser caches the old color  

<details><summary>💡 Answer</summary>

**B) 12 changes hardcoded vs 1 change with variables**

This is the core value proposition of design tokens. In a larger site, one color might appear in 50 places. A CSS variable change in `:root` propagates instantly everywhere. Hardcoded values require a search-and-replace that is error-prone and easy to miss.

</details>

---

### Q11: A CSS transition is defined on the **base** state of a button. Why does this matter?

```css
/* Option A — base state: */
button { transition: background 0.2s ease; }
button:hover { background: blue; }

/* Option B — hover state only: */
button { background: white; }
button:hover { background: blue; transition: background 0.2s ease; }
```

**A)** No difference  
**B)** Option A animates both entering AND leaving hover. Option B only animates when hovering over (transition entering), not when mousing out — the exit snaps instantly  
**C)** Option B is correct  
**D)** Transitions on base states are ignored  

<details><summary>💡 Answer</summary>

**B) Transition on base state animates both directions**

The transition property on the base element applies when transitioning FROM that state in both directions. On the hover state, it only applies when transitioning INTO hover. Always put `transition` on the base state for smooth bi-directional animation.

</details>

---

### Q12: What is a CSS skeleton loading screen?

**A)** A wireframe of the page layout  
**B)** Placeholder elements with a pulsing animation that mimic the shape of incoming content — shown while data is loading  
**C)** An animation of bones  
**D)** A page with no CSS  

<details><summary>💡 Answer</summary>

**B) Placeholder shapes with animation mimicking future content**

```css
.skeleton {
    background: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
    background-size: 200% 100%;
    animation: shimmer 1.5s infinite;
    border-radius: var(--radius);
    height: 1rem;
}
@keyframes shimmer {
    0% { background-position: 200% 0; }
    100% { background-position: -200% 0; }
}
```

Skeletons reduce perceived loading time and prevent layout shift — better UX than a spinner.

</details>

---

### Q13: Your artist cards have different heights because some artist names are longer. How do you make all cards the same height in a CSS grid?

**A)** Set a fixed `height` on each card  
**B)** Use `align-items: stretch` on the grid container (default) — flex/grid stretches children to fill the row height  
**C)** Use JavaScript to measure and equalize heights  
**D)** Truncate long names  

<details><summary>💡 Answer</summary>

**B) `align-items: stretch` (grid/flex default)**

In CSS Grid and Flexbox, `align-items: stretch` (the default) makes all children in a row the same height — equal to the tallest item. Make the card's inner layout flexbox with `flex-direction: column` and push the footer to the bottom with `margin-top: auto` on the last element.

</details>

---

### Q14: What does `overflow: hidden` on an artist card do to a large image inside it?

**A)** Hides the entire card  
**B)** Clips the image to the card's boundaries — required when using `border-radius` to prevent image corners from appearing outside the rounded card  
**C)** Makes the image scroll inside the card  
**D)** Converts the image to a placeholder  

<details><summary>💡 Answer</summary>

**B) Clips content to the card boundaries — essential for rounded cards with images**

```css
.artist-card {
    border-radius: var(--radius);
    overflow: hidden;  /* prevents image from bleeding outside the border-radius */
}
.artist-card img {
    width: 100%;
    aspect-ratio: 1;  /* square image */
    object-fit: cover; /* fill without distortion */
}
```

Without `overflow: hidden`, the image's corners appear outside the card's rounded corners.

</details>

---

### Q15: What does `object-fit: cover` do to an image?

**A)** Makes the image transparent  
**B)** Scales the image to fill its container while maintaining aspect ratio — crops the edges if needed rather than distorting  
**C)** Stretches the image to fill the container regardless of ratio  
**D)** Adds a border around the image  

<details><summary>💡 Answer</summary>

**B) Fills container while maintaining ratio — crops instead of distorting**

```css
img {
    width: 100%;
    height: 200px;
    object-fit: cover;  /* fills the 200px tall box, crops top/bottom if needed */
}
```

Without `object-fit`, images stretch and distort. `contain` shows the full image but leaves empty space. `cover` is almost always the right choice for cards and thumbnails.

</details>

---

## 📋 SECTION 3: ACCESSIBILITY (5 Questions)

### Q16: What is WCAG and why does it matter?

**A)** A CSS framework  
**B)** Web Content Accessibility Guidelines — an international standard defining how to make web content accessible to people with disabilities, including minimum color contrast ratios  
**C)** A JavaScript testing tool  
**D)** A Go code quality standard  

<details><summary>💡 Answer</summary>

**B) International accessibility standard — includes color contrast, keyboard nav, and more**

WCAG AA requires at minimum:
- 4.5:1 contrast ratio for normal text
- 3:1 for large text (18px+ or 14px+ bold)
- Keyboard navigation for all interactive elements
- Alt text for all images

Tools: browser extensions like axe, or online checkers like `contrast-ratio.com`.

</details>

---

### Q17: What is the minimum contrast ratio required for normal body text under WCAG AA?

**A)** 2:1  
**B)** 3:1  
**C)** 4.5:1  
**D)** 7:1 (AAA)  

<details><summary>💡 Answer</summary>

**C) 4.5:1 for normal text (WCAG AA)**

Light gray text (`#aaa`) on white background has about a 2:1 ratio — fails WCAG. Dark gray (`#767676`) on white is exactly 4.5:1 — barely passes. Dark text on dark background or light text on light background are common failures. Always verify with a tool before finalizing colors.

</details>

---

### Q18: Why should you NEVER use `outline: none` on a focused element without providing a replacement?

**A)** `outline` doesn't work in Chrome  
**B)** The focus ring is the only visual indicator for keyboard users navigating by Tab — removing it makes the site unusable for keyboard navigation  
**C)** It's a CSS syntax error  
**D)** It removes all styling  

<details><summary>💡 Answer</summary>

**B) The focus ring is critical for keyboard users**

Removing `outline` is a common "aesthetics" decision that breaks accessibility. Instead of removing it, style it:
```css
button:focus-visible {
    outline: 2px solid var(--color-primary);
    outline-offset: 2px;
}
```

`:focus-visible` only shows the ring for keyboard navigation, not mouse clicks — a good compromise.

</details>

---

### Q19: What `alt` text should you use for an artist's photo on the card?

**A)** `alt=""` (empty — decorative image)  
**B)** `alt="photo"` — generic description  
**C)** `alt="Artist photo of Queen"` or `alt="{{ .Name }}"` — descriptive, specific  
**D)** No `alt` attribute at all  

<details><summary>💡 Answer</summary>

**C) Specific, descriptive alt text using the artist's name**

Screen readers read the `alt` attribute aloud. `alt=""` tells a screen reader "ignore this image" (correct for purely decorative images). `alt="photo"` is useless. `alt="Queen"` or `alt="Artist photo of Queen"` gives a blind user meaningful context. Use your template data: `alt="{{ .Name }}"`.

</details>

---

### Q20: An error message on your page uses only red color to signal an error — no icon, no text prefix like "Error:". Why does this violate WCAG?

**A)** Red is not allowed in web design  
**B)** Color-blind users cannot distinguish red from other colors — information conveyed by color alone is inaccessible to roughly 8% of male users  
**C)** Red is too bright  
**D)** It violates rule 5 only  

<details><summary>💡 Answer</summary>

**B) Color alone is inaccessible — combine with text or icon**

About 8% of men have red-green color blindness. To communicate "this is an error" accessibly: use both color AND a text label ("⚠ Error:") or icon. Never rely on color as the sole differentiator between states.

```html
<div class="error" role="alert">
    <span aria-hidden="true">⚠</span>
    Error: Could not load artist data
</div>
```

</details>

---

## 📋 SECTION 4: RESPONSIVE DESIGN & INTERACTIONS (4 Questions)

### Q21: You design the artist card grid with `grid-template-columns: repeat(4, 1fr)`. On a 375px mobile screen, what problem occurs?

**A)** The grid disappears  
**B)** Cards are squished to about 93px wide — images become unreadable and text overflows  
**C)** The grid becomes vertical automatically  
**D)** CSS Grid doesn't work on mobile  

<details><summary>💡 Answer</summary>

**B) Cards are too narrow on mobile**

A 4-column grid on a 375px screen = ~93px per card. Fix with responsive columns:

```css
.artist-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    /* OR use media queries: */
    grid-template-columns: repeat(4, 1fr);
}
@media (max-width: 768px) {
    .artist-grid { grid-template-columns: repeat(2, 1fr); }
}
@media (max-width: 480px) {
    .artist-grid { grid-template-columns: 1fr; }
}
```

`auto-fill` with `minmax` is the most elegant solution — it automatically adjusts columns to fit the available width.

</details>

---

### Q22: What does `@keyframes` do in CSS?

**A)** Defines the frames per second for animations  
**B)** Defines the intermediate states of a CSS animation — the "from" and "to" states or percentage waypoints  
**C)** Imports external animation libraries  
**D)** Triggers a JavaScript function  

<details><summary>💡 Answer</summary>

**B) Defines the intermediate states of a CSS animation**

```css
@keyframes fadeIn {
    from { opacity: 0; transform: translateY(10px); }
    to   { opacity: 1; transform: translateY(0); }
}

.artist-card {
    animation: fadeIn 0.3s ease forwards;
}
```

Combine with `animation` property: `animation: name duration timing-function fill-mode`. `forwards` keeps the final state instead of reverting.

</details>

---

### Q23: A "lift" hover effect on artist cards uses `transform: translateY(-4px)`. How do you animate it smoothly?

**A)** `transition: all 0.2s ease` on the base state  
**B)** `animation: lift 0.2s` in the hover state  
**C)** JavaScript `element.animate()`  
**D)** `transform-transition: 0.2s`  

<details><summary>💡 Answer</summary>

**A) `transition: transform 0.2s ease` on the base state**

```css
.artist-card {
    transition: transform 0.2s ease, box-shadow 0.2s ease;  /* base state */
}
.artist-card:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-card-hover);
}
```

As established in Q11: transition on the base state animates both entering AND leaving hover.

</details>

---

### Q24: What does `prefers-reduced-motion` media query do and when should you use it?

**A)** Reduces CSS file size  
**B)** Detects when the user has requested reduced animation (accessibility setting) — you should disable or minimize animations when it matches  
**C)** Only works on mobile  
**D)** Reduces the frame rate  

<details><summary>💡 Answer</summary>

**B) Respects user's OS-level "reduce motion" accessibility preference**

```css
@media (prefers-reduced-motion: reduce) {
    *, *::before, *::after {
        animation-duration: 0.01ms !important;
        transition-duration: 0.01ms !important;
    }
}
```

Some users (those with vestibular disorders, epilepsy, or other conditions) configure their OS to minimize motion. Ignoring this setting is an accessibility violation. This single block of CSS respects millions of users who need it.

</details>

---

## 📋 SECTION 5: PUTTING IT TOGETHER (3 Questions)

### Q25: Before writing any CSS, what should you produce first?

**A)** A fully working prototype  
**B)** A sketch/wireframe and a list of CSS variables covering colors, typography, spacing, and radius  
**C)** A detailed CSS file with placeholder values  
**D)** A color palette only  

<details><summary>💡 Answer</summary>

**B) A sketch and CSS variable definitions**

Sketching on paper takes 15 minutes and prevents hours of CSS rewrites. Defining variables before components means every component has a consistent source of truth from the first line. The spec's milestone 2 explicitly says: define your design system before styling any component.

</details>

---

### Q26: You finish styling. A tester says the site looks "inconsistent." What is the most systematic way to audit for consistency?

**A)** Look at it again carefully  
**B)** Open DevTools on each page and check every color, spacing, and font value — verify each one is a CSS variable from `:root`, not a hardcoded value  
**C)** Ask someone else to design it  
**D)** Add more animations  

<details><summary>💡 Answer</summary>

**B) Audit every value in DevTools — verify all come from CSS variables**

Open DevTools → Elements → Styles. For each element, look for hardcoded values (any hex color, px value, or font name not using `var(--...)`). Each one is a potential consistency violation. The goal: zero hardcoded values, 100% through variables.

</details>

---

### Q27: You use `font-weight: bold` in some places and `font-weight: 700` in others, and `font-weight: 600` in a third place. Why is this a problem?

**A)** `bold` and `700` are different weights  
**B)** Inconsistent weight values create visual hierarchy confusion — viewers can't tell which text is "more important" than other text. Define `--font-weight-bold: 700` and use it everywhere  
**C)** CSS doesn't support `bold` as a keyword  
**D)** Screen readers interpret them differently  

<details><summary>💡 Answer</summary>

**B) Inconsistent weights destroy visual hierarchy**

`bold` = `700`. `600` is semi-bold. `800` is extra-bold. If some headings are `700` and some are `600` with no semantic reason, the page's visual hierarchy is arbitrary — the viewer can't learn the pattern. Define two or three weight tokens and use them consistently:

```css
:root {
    --font-weight-normal: 400;
    --font-weight-medium: 500;
    --font-weight-bold:   700;
}
```

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 25–27 ✅ | **Excellent.** Strong design systems and accessibility knowledge — start immediately. |
| 21–24 ✅ | **Ready.** Do the paper design audit (Milestone 1) thoroughly before writing CSS. |
| 16–20 ⚠️ | **Study first.** Read Shneiderman's 8 rules with examples, and work through CSS variables + Flexbox/Grid guides. |
| Below 16 ❌ | **Not ready.** This project is primarily about design quality — solid CSS fundamentals and design principles are prerequisites. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q8 | Shneiderman's 8 Golden Rules — read each one with real examples |
| Q9–Q15 | CSS design tokens, variables, transitions (base vs hover state), skeleton screens, `overflow: hidden`, `object-fit` |
| Q16–Q20 | WCAG AA, 4.5:1 contrast, `outline` focus rings, alt text, color-only errors |
| Q21–Q24 | Responsive grid, `@keyframes`, `prefers-reduced-motion` |
| Q25–Q27 | Design-first workflow, consistency audit with DevTools, font weight tokens |