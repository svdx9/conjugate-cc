---
id: TASK-1.4
title: Implement the MVP front page
status: In Progress
assignee:
  - claude
created_date: '2026-03-02 18:07'
updated_date: '2026-03-04 14:12'
labels:
  - mvp
  - frontend
  - ui
dependencies:
  - TASK-1.3
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-1
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Build the first user-visible page for the conjugation drill application. The page should explain the product purpose, present a clear entry point for the future drill experience, and work well on common desktop and mobile widths.

The navigation element will be at the top of the page and will be horizontal with the landing page as first link element, then drills, then 'verbs'
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The front page clearly communicates that the product is a conjugation drill application, includes the app name `conjugate.cc`, and includes a primary call to action or placeholder entry point for drills.
- [x] #2 The page is usable on both desktop and mobile viewport sizes without broken layout or inaccessible content.
- [ ] #3 The task includes a verification step covering the rendered front page state and any new copy or assets are documented in the repository as needed.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Implementation Plan

### Approach
Transform the existing AppShell bootstrap into a proper MVP front page with horizontal navigation.
All CSS follows a **mobile-first** strategy: base styles target mobile viewports, `min-width` media queries add desktop enhancements.

### Steps

1. **Create `Nav` component** (`src/features/shell/components/Nav.tsx`)
   - Semantic `<nav>` with `<ul>/<li>` structure for accessibility
   - Links (in order): Home (`/`), Drills (`/drills`), Verbs (`/verbs`) — plain `<a>` anchors, no router yet
   - Mobile base: links stack vertically or wrap; desktop (`min-width: 640px`): horizontal flex row

2. **Rewrite `AppShell`** (`src/features/shell/screens/AppShell.tsx`)
   - Compose `<Nav>` at top of layout
   - Hero section below nav:
     - App name `conjugate.cc` as primary heading
     - Short tagline communicating conjugation drill purpose
     - Primary CTA button "Start Drilling" (placeholder)
   - Layout: full-page flex column; nav at top, hero fills remaining space centered

3. **Update `src/styles/index.css`**
   - Mobile-first: base styles for small screens, `min-width` breakpoints for wider viewports
   - Nav: base = column-friendly compact layout; `min-width: 640px` = horizontal flex row
   - Hero: base = single-column centered; fluid typography via `clamp()`
   - Preserve existing color palette and design tokens

4. **Update tests** (`src/app/App.test.tsx`)
   - Remove old placeholder text assertions
   - Add: nav link assertions (Home, Drills, Verbs), heading `conjugate.cc`, CTA button "Start Drilling"

5. **Verification**
   - Run `npm run dev`, visually confirm mobile and desktop layouts
   - Log final copy and layout decisions in task notes
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Bootstrapped the full SolidJS + Vite 6 + TypeScript frontend (task-1.3 prerequisite work included). Files created: frontend/package.json, frontend/vite.config.ts, frontend/tsconfig.json, frontend/index.html, frontend/src/index.tsx, frontend/src/app/App.tsx, frontend/src/app/App.test.tsx, frontend/src/test-setup.ts, frontend/src/styles/index.css, frontend/src/features/shell/components/Nav.tsx, frontend/src/features/shell/screens/AppShell.tsx. Vite 6 + vitest 2 required explicit resolve.conditions=['browser'] and server.deps.inline for solid-js packages to load the web (not server) bundle in tests. CSS is mobile-first with min-width: 640px breakpoints. Nav: conjugate.cc (brand, home), Drills, Verbs. Hero: eyebrow, h1 heading conjugate.cc, description, Start Drilling CTA link. All 4 tests pass.
<!-- SECTION:NOTES:END -->
