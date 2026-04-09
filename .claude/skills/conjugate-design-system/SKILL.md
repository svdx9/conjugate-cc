---
name: conjugate-design-system
description: Conjugate.cc Design System skill for creating consistent, minimal UI components following the project's design tokens and patterns. Use when building or modifying frontend components, pages, or layouts in the conjugate-cc application.
license: MIT
metadata:
  author: Mistral Vibe
  version: "1.0.0"
  compatibility: "SolidJS 1.9.3, Tailwind CSS 4.0.0, Vite 6.0.7"
---

# Conjugate.cc Design System Skill

## Overview

This skill provides guidance for implementing the Conjugate.cc design system when building or modifying frontend components. It ensures consistency with the project's minimal aesthetic, dark mode support, and Tailwind CSS implementation patterns.

## When to Use This Skill

**Activate this skill when:**
- Creating new UI components
- Modifying existing components
- Implementing new pages
- Ensuring design consistency
- Working with Tailwind CSS classes
- Implementing dark mode support

**Do NOT use for:**
- Backend API changes
- Database modifications
- Non-UI related tasks
- General code reviews unrelated to design

## Design Tokens

All color tokens are defined as CSS custom properties in `src/styles/app.css` via Tailwind v4's `@theme` directive.

### Colors (Light / Dark)

| Token               | Light        | Dark         | Tailwind class                                       |
|---------------------|-------------|-------------|------------------------------------------------------|
| Surface             | `#ffffff`   | `#111111`   | `bg-surface` / `dark:bg-surface-dark`                |
| Text primary        | `#000000`   | `#ffffff`   | `text-text-primary` / `dark:text-text-primary-dark`  |
| Border              | `#00000026` | `#ffffff26` | `border-border` / `dark:border-border-dark`          |
| Highlight (gold)    | `#ebc61d`   | `#ebc61d`   | `text-highlight`, `hover:text-highlight`             |
| Button background   | `#f8f8f8`   | —           | `bg-btn-bg`                                          |
| Button text         | `#374151`   | —           | `text-btn-text`                                      |

### Dark Mode Implementation

- Configured via `@custom-variant dark` in `app.css` (class-based)
- `darkMode.ts` toggles the `dark` class on `document.documentElement`
- Persists to localStorage
- **Rule**: Components MUST use `dark:` Tailwind variants for styling
- **Exception**: `isDarkMode` may be imported only for non-styling logic (e.g., toggle button icon)

## Layout System

### Container Structure

```jsx
// Max width: max-w-6xl (1152px)
// Responsive padding: px-4 (mobile), sm:px-6, lg:px-8
// Centered: mx-auto
```

### Standard Page Structure

```jsx
// App.tsx (layout route)
<div class="flex min-h-screen flex-col bg-surface text-text-primary transition-colors dark:bg-surface-dark dark:text-text-primary-dark">
  <Navigation />
  <main class="flex-1">{props.children}</main>
  <Footer />
</div>
```

### PageShell Component

All pages MUST use the `PageShell` component:

```jsx
import PageShell from '../../shared/PageShell';

<PageShell>
  {/* page content */}
</PageShell>
```

## Typography System

### Font Family
- **Primary**: Inter via Tailwind (`--font-sans`)

### Text Sizes
- **Page Title**: `text-4xl font-bold` (mobile) → `text-6xl` (large screens)
- **Body Text**: `text-lg` for descriptions
- **Navigation/Footer Links**: `text-sm font-medium`

## Component Patterns

### Header Component

```jsx
<header class="sticky top-0 z-50 bg-surface transition-colors dark:bg-surface-dark">
  <nav class="px-4 sm:px-6 lg:px-8">
    <div class="flex h-16 items-center justify-between border-b border-border px-2 dark:border-border-dark">
      {/* content */}
    </div>
  </nav>
</header>
```

### Button Component

```jsx
<A
  href="/path"
  class="inline-flex h-10 items-center bg-btn-bg px-8 text-sm font-medium text-btn-text transition-colors hover:bg-highlight"
>
  Text
</A>
```

### Footer Component

```jsx
<footer class="border-t border-border bg-surface py-8 transition-colors dark:border-border-dark dark:bg-surface-dark">
  <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
    <div class="flex justify-center gap-8">
      {/* links with hover:text-highlight */}
    </div>
  </div>
</footer>
```

## Spacing System

### Vertical Padding
- **Mobile**: `py-12`
- **Small screens**: `sm:py-16`
- **Large screens**: `lg:py-20`

### Element Gaps
- **Flex containers**: `gap-4`
- **Footer links**: `gap-8`

## Styling Rules (MUST FOLLOW)

1. **No inline styles** - Use Tailwind utility classes only
2. **No JS hover handlers** - Use Tailwind `hover:` and `focus-visible:` variants
3. **No hardcoded hex values** - Use token classes from `@theme`
4. **Dark mode via `dark:` variant** - Never use conditional JS logic for colors
5. **Consistent button styling** - Height: `h-10`, Radius: 0, Border: none

## Design Principles

1. **Simplicity First**: No gradients, shadows, or decorative effects
2. **Content Focus**: Design supports content without distraction
3. **Responsive**: Mobile-first approach that scales to desktop
4. **Accessibility**: High contrast, semantic HTML, keyboard-navigable states
5. **Consistency**: Repeat patterns across all pages via shared components

## Routing Structure

```
/
  ├── (home/landing)
  ├── /drills
  ├── /verbs
  ├── /help
  ├── /contact
  └── /cookie-policy
```

All pages use the same layout wrapper (header, main content area, footer) and `PageShell` for content containers.

## Workflow Instructions

### When Creating a New Component:

1. **Check existing patterns** in `src/components/`
2. **Use PageShell** as the outer container
3. **Apply design tokens** via Tailwind classes
4. **Implement dark mode** using `dark:` variants
5. **Follow spacing system** for consistent layout
6. **Test responsiveness** across breakpoints

### When Modifying an Existing Component:

1. **Preserve existing structure** unless there's a compelling reason to change
2. **Update design tokens** if colors or spacing changes
3. **Maintain dark mode support**
4. **Keep accessibility features** (focus states, semantic HTML)

### Common Pitfalls to Avoid:

- ❌ Hardcoding colors (`bg-[#ffffff]`)
- ❌ Using inline styles (`style={{ color: 'red' }}`)
- ❌ JavaScript hover handlers
- ❌ Inconsistent spacing
- ❌ Breaking dark mode support

## Tools and Resources

- **Tailwind CSS Docs**: https://tailwindcss.com/docs
- **Project Styles**: `src/styles/app.css`
- **Dark Mode Logic**: `src/app/darkMode.ts`
- **Shared Components**: `src/components/`
- **PageShell**: `src/shared/PageShell.tsx`

## Example Implementation

### Creating a New Page Component

```jsx
import PageShell from '../../shared/PageShell';
import { A } from '@solidjs/router';

export default function NewPage() {
  return (
    <PageShell>
      <div class="space-y-8">
        <h1 class="text-4xl font-bold sm:text-6xl">Page Title</h1>
        <p class="text-lg">Description text goes here.</p>
        
        <div class="flex gap-4">
          <A
            href="/action"
            class="inline-flex h-10 items-center bg-btn-bg px-8 text-sm font-medium text-btn-text transition-colors hover:bg-highlight"
          >
            Action Button
          </A>
        </div>
      </div>
    </PageShell>
  );
}
```

This implementation follows all design system guidelines: proper PageShell usage, correct typography, consistent button styling, and responsive spacing.
