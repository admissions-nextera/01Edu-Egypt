# Groupie Tracker Visualizations Project Guide

> **Before you start:** This project builds on groupie-tracker. Read Shneiderman's 8 Golden Rules completely before writing a single line: https://www.interaction-design.org/literature/article/shneiderman-s-eight-golden-rules-will-help-you-design-better-interfaces — then audit your existing site against each rule.

---

## Objectives

By completing this project you will learn:

1. **Interface Design Principles** — Applying Shneiderman's 8 Golden Rules to a real product
2. **CSS Mastery** — Using advanced CSS to build compelling data visualizations
3. **Data Visualization** — Choosing the right visual representation for each data type
4. **Consistency** — Creating a unified design system across every page
5. **Feedback and Interactivity** — Making the UI respond to every user action
6. **Accessibility** — Ensuring the interface works for all users regardless of color or ability

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Groupie Tracker (Completed)
- All data pages working
- Basic HTML structure in place

### 2. CSS Intermediate
- CSS Grid and Flexbox for complex layouts
- CSS animations and transitions
- CSS pseudo-classes: `:hover`, `:focus`, `:active`
- CSS variables for a design system
- Search: **"CSS Grid complete guide"**
- Search: **"CSS transitions animations tutorial"**

### 3. Shneiderman's 8 Golden Rules
Read each rule and understand what it means for your site before proceeding:
- https://www.interaction-design.org/literature/article/shneiderman-s-eight-golden-rules-will-help-you-design-better-interfaces

---

## Project Structure

```
groupie-tracker-visualizations/
├── main.go
├── handlers.go
├── api.go
├── templates/
│   ├── index.html
│   └── artist.html
├── static/
│   ├── style.css       ← comprehensive design system
│   └── (any JS files)
└── go.mod
```

---

## Milestone 1 — Audit Against Shneiderman's Rules

**This milestone has no code.**

Go through your current site and score yourself on each rule. Be honest.

| Rule | What it means for your site | Current status |
|---|---|---|
| Strive for consistency | Same fonts, spacing, colors, button styles everywhere | |
| Enable shortcuts | Can power users do things faster? Keyboard nav? | |
| Offer informative feedback | Does every action produce a visible response? | |
| Design dialogue to yield closure | Does every process have a clear start, middle, and end? | |
| Offer simple error handling | Are errors clear, non-blaming, and recoverable? | |
| Permit easy reversal | Can the user undo or go back from any state? | |
| Support internal locus of control | Does the user feel in control, not trapped? | |
| Reduce short-term memory load | Is everything visible that needs to be visible? | |

**For each rule you are failing, write down specifically what you will change.**

---

## Milestone 2 — Build a Design System

**Goal:** Define all visual tokens in CSS variables before touching any component styles.

**Questions to answer:**
- What is your typographic scale? (h1, h2, h3, body, small — what sizes?)
- What is your color palette? (Primary, secondary, background, surface, error, success — what values?)
- What is your spacing scale? (4px, 8px, 16px, 24px, 32px, 48px — or your own scale?)
- What is your border radius convention? (0, 4px, 8px, full — pick one and stick to it.)
- What are your shadow levels?

**Code Placeholder:**
```css
/* static/style.css */

:root {
    /* Typography */
    /* --font-family: ... */
    /* --font-size-sm: ... */
    /* --font-size-base: ... */
    /* --font-size-lg: ... */
    /* --font-size-xl: ... */
    /* --font-size-2xl: ... */

    /* Colors */
    /* --color-primary: ... */
    /* --color-primary-dark: ... */
    /* --color-background: ... */
    /* --color-surface: ... */
    /* --color-text: ... */
    /* --color-text-muted: ... */
    /* --color-error: ... */
    /* --color-success: ... */

    /* Spacing */
    /* --space-xs: ... */
    /* --space-sm: ... */
    /* --space-md: ... */
    /* --space-lg: ... */
    /* --space-xl: ... */

    /* Other */
    /* --radius: ... */
    /* --shadow: ... */
    /* --transition: ... */
}
```

**Verify:** Every color, size, and spacing value in your CSS comes from a variable — no hardcoded values anywhere.

---

## Milestone 3 — Artist Card Visualization

**Goal:** The artist list uses visually rich cards that communicate key data at a glance.

**Questions to answer:**
- What information should be visible on the card without clicking? (Image, name, creation year, member count?)
- How does the card change on hover to indicate it is clickable?
- How do you display the member count visually — a number, icons, dots?
- How do you handle artists with very long names without breaking the layout?

**Code Placeholder:**
```css
/* Artist card */
/* Grid or flex layout for the card grid */
/* Card container: background, border-radius, shadow, overflow hidden */
/* Image: aspect ratio maintained, fills card width */
/* Card body: padding, spacing between elements */
/* Hover state: lift effect with transform and shadow transition */
/* Member count: visual treatment */
/* Creation year: muted color, smaller size */
```

**Verify:**
- Cards are uniform in size regardless of content length
- Hover produces a smooth, visible transition
- The grid is responsive — fewer columns on narrow screens

---

## Milestone 4 — Artist Detail Visualization

**Goal:** The artist detail page presents data in a way that is immediately scannable and visually organized.

**Questions to answer:**
- How do you display the concert locations and dates? (Timeline? Table? Grouped list?)
- How do you visually separate different sections (bio, members, concerts)?
- How do you show the list of members — plain text, avatars, chips?
- What visual hierarchy makes the most important information (name, image) prominent?

**Code Placeholder:**
```css
/* Detail page layout */
/* Two-column layout: artist info left, concerts right — collapses to single on mobile */

/* Members section */
/* Pill/chip style for each member name */

/* Concert list */
/* Clear grouping of location and its dates */
/* Visual separator between locations */

/* Section headers */
/* Consistent with your typography scale */
```

---

## Milestone 5 — Informative Feedback (Rule 3)

**Goal:** Every user action that changes the page produces a visible, immediate response.

**Questions to answer:**
- When data is loading (API fetch), what does the user see?
- When an error occurs (404, 500), what does the page show?
- When a search returns no results, is there a clear empty state?
- When a form is submitted, is there any indication it was received?

**Code Placeholder:**
```css
/* Loading state */
/* Skeleton cards or a spinner while data loads */

/* Empty state */
/* Centered message with icon when no artists match filters or search */

/* Error state */
/* Clear, styled error message — not just a blank page */

/* Transition on list update */
/* Smooth fade or slide when artist cards appear/disappear */
```

```html
<!-- In your templates, add these states: -->
<!-- Loading: shown while Go fetches data (if applicable) -->
<!-- Empty: shown when filter/search returns no results -->
<!-- {{ if .Error }}<div class="error-state">...</div>{{ end }} -->
<!-- {{ if not .Artists }}<div class="empty-state">...</div>{{ end }} -->
```

---

## Milestone 6 — Shortcuts and Locus of Control (Rules 2 and 7)

**Goal:** Power users can navigate the site faster, and every user always knows where they are.

**Questions to answer:**
- How does the user know which page they are on? (Active nav link, breadcrumb, page title?)
- Can the user get back to the home page from anywhere?
- Is there a keyboard shortcut or fast path to the search bar?
- Can the user jump directly to an artist they have visited before?

**Code Placeholder:**
```css
/* Navigation */
/* Active page indicator on nav links */
/* Breadcrumb on the detail page */

/* Focus styles */
/* Visible :focus ring on all interactive elements */
/* Never remove outline without replacing it */
```

---

## Milestone 7 — Accessibility Check

**Goal:** The site is usable without relying on color alone, and keyboard navigation works throughout.

**Questions to answer:**
- Does every image have a meaningful `alt` attribute?
- Can you navigate the entire site using only Tab and Enter?
- Does your text have sufficient contrast against its background?
- Do your error and success states use more than just color to communicate their meaning?

**Verify tools:**
- Search: **"WCAG contrast checker online"** — paste your text and background colors
- Search: **"axe browser extension accessibility"** — install and run on your pages
- Tab through the entire site and confirm every link and button is reachable

---

## Debugging Checklist

- Does the layout break at certain screen widths? Open DevTools and drag the viewport — find the exact breakpoint where it breaks and add a media query.
- Do CSS transitions not trigger? Check that you are using `transition` on the element, not on `:hover`. The transition property belongs on the base state.
- Are cards not the same height despite different content amounts? Use `align-items: stretch` on the grid or flex container.
- Does the design look inconsistent? Open DevTools and check for any hardcoded color or spacing values not using your CSS variables.
- Do error states not appear? Check your template conditionals — make sure the error field is being passed to the template when it should be.

---

## Key Concepts and Resources

| Concept | What to Search |
|---|---|
| Shneiderman's rules | "Shneiderman 8 golden rules examples" |
| CSS design system | "CSS design tokens variables tutorial" |
| CSS card hover effects | "CSS card hover lift shadow transition" |
| CSS skeleton loading | "CSS skeleton screen loading animation" |
| CSS empty state design | "empty state UI design best practices" |
| Accessibility audit | "WCAG 2.1 checklist web accessibility" |
| Color contrast | "WCAG AA contrast ratio checker" |

---

## Submission Checklist

- [ ] All 8 Shneiderman rules visibly addressed in the design
- [ ] Design system defined as CSS variables — no hardcoded values
- [ ] Artist cards are consistent, responsive, and hover-animated
- [ ] Artist detail page has clear visual hierarchy and data grouping
- [ ] Loading, empty, and error states are all designed and visible
- [ ] Navigation shows active page clearly
- [ ] Keyboard navigation works for all interactive elements
- [ ] All images have meaningful `alt` attributes
- [ ] Text passes WCAG AA color contrast on all backgrounds
- [ ] Error states use more than color alone
- [ ] Layout is fully responsive from mobile to desktop
- [ ] The design is consistent — same fonts, spacing, and colors everywhere