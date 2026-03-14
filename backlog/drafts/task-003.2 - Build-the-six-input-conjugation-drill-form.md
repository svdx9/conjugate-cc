---
id: TASK-003.2
title: Build the six-input conjugation drill form
status: To Do
assignee: []
created_date: '2026-03-02 23:31'
updated_date: '2026-03-14 16:27'
labels:
  - frontend
  - ui
  - drills
dependencies:
  - TASK-001.5
  - TASK-002.2
  - TASK-003.1
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-003
ordinal: 32000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the main conjugation drill page as a form-driven interface. The page should present the target verb infinitive, show the six pronouns in a left-hand column, and render six aligned learner input fields on the right so the user can enter their responses for the `etre` drill round.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The drill page renders the `etre` infinitive prompt together with six pronouns on the left and six corresponding answer inputs on the right.
- [ ] #2 The layout is readable and usable on common desktop and mobile widths, with clear row alignment between each pronoun and its input.
- [ ] #3 The form includes both a `Show answer` control and a `Submit answer` control in positions that support the intended drill flow.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Context

- Design tokens (dark theme): `bg-bg` (#0d0f14), `bg-card` (#151922), `text-ink` (#f5f1ea), `text-muted` (#a99f94), `border-line` (#2a2f3a), `text-accent` (#4d5d6a), `text-sage` (#6f8a78), `text-danger` (#D70A53), `text-success` (#6BB342)
- Fonts: `font-sans` (IBM Plex Sans), `font-serif` (Source Serif 4)
- `DrillsPage.tsx` is currently a stub (`<main><h1>Drills</h1></main>`)
- `etreDrillSet` is ready in `features/drills/data/etre-stub.ts`
- TASK-003.3 adds show/hide state per row — TASK-003.2 only needs both buttons present and positioned; they don't need to be functional yet

## Layout

```
┌─────────────────────────────────┐
│ Présent                         │  ← tense label (muted, eyebrow style)
│ être                            │  ← infinitive (serif, large)
│                                 │
│  je    [________________]       │
│  tu    [________________]       │
│  il    [________________]       │
│  nous  [________________]       │
│  vous  [________________]       │
│  ils   [________________]       │
│                                 │
│         [Show answers] [Submit] │  ← controls right-aligned
└─────────────────────────────────┘
```

## Files to change

### 1. `frontend/src/pages/DrillsPage.tsx` (replace stub)

```tsx
import { createSignal, For } from "solid-js";
import { etreDrillSet } from "../features/drills/data/etre-stub";

export function DrillsPage() {
  const [inputs, setInputs] = createSignal<string[]>(
    etreDrillSet.rows.map(() => "")
  );

  return (
    <main class="flex-1 flex items-center justify-center px-4 py-12">
      <div class="w-full max-w-sm">
        <p class="font-sans text-[0.65rem] font-bold uppercase tracking-[0.34em] text-muted mb-1">
          {etreDrillSet.tense}
        </p>
        <h2 class="font-serif text-3xl sm:text-4xl text-ink mb-8">
          {etreDrillSet.infinitive}
        </h2>
        <form onSubmit={(e) => e.preventDefault()} class="space-y-3">
          <For each={etreDrillSet.rows}>
            {(row, i) => (
              <div class="flex items-center gap-4">
                <span class="w-10 text-right font-sans text-sm text-muted shrink-0">
                  {row.pronoun}
                </span>
                <input
                  id={`row-${i()}`}
                  type="text"
                  value={inputs()[i()]}
                  onInput={(e) =>
                    setInputs((prev) => {
                      const next = [...prev];
                      next[i()] = e.currentTarget.value;
                      return next;
                    })
                  }
                  autocomplete="off"
                  autocorrect="off"
                  autocapitalize="off"
                  spellcheck={false}
                  class="flex-1 bg-card border border-line rounded-lg px-3 py-2 text-sm text-ink placeholder-muted focus:outline-none focus:border-accent transition-colors"
                />
              </div>
            )}
          </For>
          <div class="flex justify-end gap-3 pt-2">
            <button
              type="button"
              class="px-4 py-2 border border-line text-muted text-sm font-medium rounded-lg hover:text-ink hover:border-accent transition-colors"
            >
              Show answers
            </button>
            <button
              type="submit"
              class="px-4 py-2 bg-card border border-line text-sage text-sm font-medium rounded-lg hover:border-sage transition-colors"
            >
              Submit
            </button>
          </div>
        </form>
      </div>
    </main>
  );
}
```

### 2. `frontend/src/pages/DrillsPage.test.tsx` (new)

Five tests:
1. Renders the infinitive "être"
2. Renders the tense label "Présent"
3. Renders exactly 6 text inputs
4. Renders all 6 pronouns (je, tu, il, nous, vous, ils)
5. Renders both "Show answers" and "Submit" buttons

## Sequence

1. Replace `DrillsPage.tsx`
2. Create `DrillsPage.test.tsx`
3. Run `npm test` — all tests pass

## Out of scope

- Show/hide per-row answer reveal (TASK-003.3)
- Submit handler that navigates to results (TASK-003.4)
- Any routing changes
<!-- SECTION:PLAN:END -->
