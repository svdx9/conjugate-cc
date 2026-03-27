# Conjugate.cc Design System

Minimal, clean design with black text on white background. Focus on clarity and simplicity.
Supports dark mode via a `.dark` class on `<html>`, toggled by `src/app/darkMode.ts`.

## Design Tokens

All color tokens are defined as CSS custom properties in `src/styles/app.css` via Tailwind v4's `@theme` directive. Use the corresponding Tailwind utility classes — never hardcode hex values in components.

### Colors (Light / Dark)
| Token               | Light        | Dark         | Tailwind class                                       |
|---------------------|-------------|-------------|------------------------------------------------------|
| Surface             | `#ffffff`   | `#111111`   | `bg-surface` / `dark:bg-surface-dark`                |
| Text primary        | `#000000`   | `#ffffff`   | `text-text-primary` / `dark:text-text-primary-dark`  |
| Border              | `#00000026` | `#ffffff26` | `border-border` / `dark:border-border-dark`          |
| Highlight (gold)    | `#ebc61d`   | `#ebc61d`   | `text-highlight`, `hover:text-highlight`             |
| Button background   | `#f8f8f8`   | —           | `bg-btn-bg`                                          |
| Button text         | `#374151`   | —           | `text-btn-text`                                      |

### Dark mode
- Configured via `@custom-variant dark` in `app.css` (class-based, not media-query).
- `darkMode.ts` toggles the `dark` class on `document.documentElement` and persists to localStorage.
- Components MUST NOT import `isDarkMode` for styling. Use `dark:` Tailwind variants instead.
- `isDarkMode` may be imported only for non-styling logic (e.g. the toggle button icon).

### Layout
- **Max Width**: `max-w-6xl` (1152px)
- **Responsive Padding**:
  - Mobile: `px-4` (16px)
  - Small: `sm:px-6` (24px)
  - Large: `lg:px-8` (32px)
- **Container**: Centered with `mx-auto`

### Typography
- **Font Family**: Inter via Tailwind (`--font-sans`)
- **Page Title**: `text-4xl font-bold` (mobile) → `text-6xl` (large)
- **Body Text**: `text-lg` for descriptions
- **Nav/Footer Links**: `text-sm font-medium`

### Buttons
- **Height**: 40px (`h-10`)
- **Border Radius**: 0 (no rounding)
- **Border**: None
- **Padding**: `px-4` (nav links), `px-8` (CTA buttons)
- **Hover**: Use Tailwind `hover:` variant — never use JS `onMouseEnter`/`onMouseLeave` for hover styles

### Spacing
- **Section Vertical Padding**:
  - Mobile: `py-12`
  - Small: `sm:py-16`
  - Large: `lg:py-20`
- **Gap Between Elements**: `gap-4` (flex), `gap-8` (footer links)

## Component Patterns

### Header
```jsx
<header class="sticky top-0 z-50 bg-surface transition-colors dark:bg-surface-dark">
  <nav class="px-4 sm:px-6 lg:px-8">
    <div class="flex h-16 items-center justify-between border-b border-border px-2 dark:border-border-dark">
      {/* content */}
    </div>
  </nav>
</header>
```

### Layout Wrapper (App.tsx — layout route)
```jsx
<div class="flex min-h-screen flex-col bg-surface text-text-primary transition-colors dark:bg-surface-dark dark:text-text-primary-dark">
  <Navigation />
  <main class="flex-1">{props.children}</main>
  <Footer />
</div>
```

### Page Container (PageShell — `src/shared/PageShell.tsx`)
```jsx
import PageShell from '../../shared/PageShell';

<PageShell>
  {/* page content */}
</PageShell>
```

All pages MUST use PageShell for their outer container.

### Button Link
```jsx
<A
  href="/path"
  class="inline-flex h-10 items-center bg-btn-bg px-8 text-sm font-medium text-btn-text transition-colors hover:bg-highlight"
>
  Text
</A>
```

### Footer
```jsx
<footer class="border-t border-border bg-surface py-8 transition-colors dark:border-border-dark dark:bg-surface-dark">
  <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
    <div class="flex justify-center gap-8">
      {/* links with hover:text-highlight */}
    </div>
  </div>
</footer>
```

## Styling Rules

1. **No inline styles** — all colors, borders, and backgrounds via Tailwind utility classes
2. **No JS hover handlers** — use Tailwind `hover:` and `focus-visible:` variants
3. **No hardcoded hex values** in components — use token classes from `@theme`
4. **Dark mode via `dark:` variant only** — never conditional JS logic for colors

## Principles

1. **Simplicity First**: No gradients, shadows, or decorative effects
2. **Content Focus**: Design supports content, doesn't distract
3. **Responsive**: Mobile-first, scales to desktop
4. **Accessibility**: High contrast, semantic HTML, keyboard-navigable hover/focus states
5. **Consistency**: Repeat patterns across all pages via shared components

## Routing Structure

```
/ (home/landing)
/drills
/verbs
/help
/contact
/cookie-policy
```

All pages use the same Layout route wrapper (header, main content area, footer).
Pages use `<PageShell>` for their content container.
