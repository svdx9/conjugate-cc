---
id: TASK-001.7
title: Add client-side routing with @solidjs/router
status: To Do
assignee: []
created_date: '2026-03-10 16:24'
updated_date: '2026-03-14 16:27'
labels:
  - frontend
  - routing
dependencies: []
parent_task_id: TASK-001
priority: medium
ordinal: 27000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Install `@solidjs/router` and wire up the SPA router in `App.tsx`. Replace all plain `<a>` tags used for internal navigation with the router's `<A>` component to prevent full-page reloads on link clicks.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 @solidjs/router is installed and App is wrapped in <Router>
- [x] #2 Nav links (/, /drills, /verbs) use <A> and navigate without a full-page reload
- [x] #3 Start Drilling CTA uses <A>
- [x] #4 All existing tests pass
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Implementation Plan

### Research Summary

Current state:
- `App.tsx` renders `<HomePage />` + `<DevFooter />` with no router
- `Nav.tsx` has 3 plain `<a>` tags: `/`, `/drills`, `/verbs`
- `HomePage.tsx` has 1 plain `<a>` tag: `/drills` (the Start CTA)
- `App.test.tsx` renders `<App />` and asserts on `role="link"` elements
- `@solidjs/router` is **not** installed yet

### Steps

1. **Install `@solidjs/router`**
   - `cd frontend && npm install @solidjs/router`

2. **Create `src/pages/DrillsPage.tsx` and `src/pages/VerbsPage.tsx`**
   - Minimal stub components (a heading only) so the router has real components to mount at `/drills` and `/verbs`

3. **Rewrite `src/App.tsx`**
   - Import `Router`, `Route` from `@solidjs/router`
   - Define a root layout component (`RootLayout`) that renders `<DevFooter />` below the matched route via `<Outlet />`
   - Wire up routes:
     ```tsx
     <Router>
       <Route path="/" component={RootLayout}>
         <Route path="/" component={HomePage} />
         <Route path="/drills" component={DrillsPage} />
         <Route path="/verbs" component={VerbsPage} />
       </Route>
     </Router>
     ```
   - `DevFooter` lives in the root layout so it is always visible regardless of active route

4. **Update `src/components/Nav.tsx`**
   - Import `A` from `@solidjs/router`
   - Replace all three `<a href="...">` with `<A href="...">`

5. **Update `src/pages/HomePage.tsx`**
   - Import `A` from `@solidjs/router`
   - Replace the Start CTA `<a href="/drills">` with `<A href="/drills">`

6. **Update `src/App.test.tsx`**
   - `<A>` renders a native `<a>` element, so `getByRole("link", ...)` assertions still work without change
   - `<Router>` requires `window.location` which jsdom provides; no `MemoryRouter` wrapper should be needed
   - If the Router fails to initialise in jsdom (e.g. throws on missing history API), wrap the render with `MemoryRouter` from `@solidjs/router` instead of `Router`

7. **Run tests**
   - `cd frontend && npm test` — all 4 existing assertions must pass

### Notes

- No new public-facing routes or pages beyond stubs are in scope for this task
- `DevFooter` remains always-mounted via the root layout, matching current behaviour
- The `<A>` component from `@solidjs/router` renders a native `<a>`, so existing CSS class names and test selectors need no changes
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Installed @solidjs/router v0.15.4. Rewrote App.tsx with Router + nested Route definitions and a RootLayout component that renders DevFooter persistently via props.children. Created stub DrillsPage and VerbsPage. Replaced all plain `<a>` tags in Nav.tsx and HomePage.tsx with `<A>` from @solidjs/router. All 4 existing tests pass unchanged.
<!-- SECTION:FINAL_SUMMARY:END -->
