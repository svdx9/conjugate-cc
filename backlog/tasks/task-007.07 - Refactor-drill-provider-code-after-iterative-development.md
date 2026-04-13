---
id: TASK-007.07
title: Refactor drill provider code after iterative development
status: To Do
assignee: []
created_date: '2026-04-13 15:30'
updated_date: '2026-04-13 15:30'
labels:
  - frontend
  - refactor
  - drills
dependencies:
  - TASK-007.5
references:
  - frontend/src/features/drills/provider.ts
  - frontend/src/features/drills/types.ts
  - frontend/src/shared/types.ts
  - frontend/src/features/drills/provider.test.ts
parent_task_id: TASK-007
priority: medium
ordinal: 77000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Clean up accumulated cruft in the drill provider and shared types after multiple rounds of iterative development on TASK-007.1 through TASK-007.6. The code works and tests pass, but readability and maintainability have degraded.

This should be done after TASK-007.5 (conjugation engine port) since that task will replace the stub provider — no point polishing code that's about to be rewritten. The shared types cleanup and any patterns established here should still apply post-engine-port.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All existing tests continue to pass after refactoring
- [ ] #2 Windsurf watermark comments removed from provider.ts
- [ ] #3 Redundant `as Tense` casts eliminated from verbData (use buildDrillData helper or properly typed structure)
- [ ] #4 validateString removed — fold empty-check directly into validateTense/validatePronoun/normalizeVerb
- [ ] #5 Overly verbose comments pruned (remove comments that restate the code)
- [ ] #6 Unused CommonErrorCode type removed from shared/types.ts
- [ ] #7 verbData repetition reduced via a compact data format and builder function
- [ ] #8 Pronoun synthesis logic extracted into its own named function
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Issues Identified

### 1. Windsurf watermark comments (provider.ts)
Three `/**** Windsurf Command ****/` banners at lines 159, 176, 222 — noise from tooling that should never be committed.

### 2. Excessive `as Tense` casts in verbData
Every single DrillItem repeats `'présent' as Tense` because the outer record is typed as `Record<string, Record<string, DrillData>>` instead of using `Tense` keys. This results in ~30 unnecessary casts. Fix by either:
- Typing verbData as `Record<string, Partial<Record<Tense, DrillData>>>` and casting at the outer level, or
- Using a `buildDrillData(verb, tense, entries)` helper that constructs DrillData from compact tuples like `['je', 'suis']`.

### 3. Over-engineered validateString
The `typeof str !== 'string'` check is pointless — TypeScript already enforces the parameter type. The entire `validateString` function exists just to check emptiness and is called by three other validators. Inline the empty check directly.

### 4. Redundant validation chaining
`validateTense` calls `validateString` then checks `validTenses.includes()`. The string validation is unnecessary since `includes()` on an empty/falsy string just returns `false`. Same for `validatePronoun`. Collapse each into a single check.

### 5. Verbose/obvious comments
Many comments just restate what the code does:
- `// Return the validated string` above `return success(str)`
- `// Check if the string is empty or not defined` above `if (!str ...)`
- `// Normalize the verb string` above `const normalizedVerb = ...`
Remove these. Keep comments that explain *why*, not *what*.

### 6. Unused CommonErrorCode type (shared/types.ts)
`CommonErrorCode` is defined at shared/types.ts:43-49 but never referenced anywhere. Remove it.

### 7. verbData repetition
Each DrillItem in verbData manually repeats the verb name and tense that are already encoded in the parent keys. A `buildDrillData` helper taking compact `[pronoun, text, extras?]` tuples would cut the data block roughly in half.

### 8. Inline pronoun synthesis
The pronoun synthesis logic (deriving elle/on from il, elles from ils) at lines 278-304 is embedded in getDrillData. Extract it into a named `synthesizePronouns(items)` function for clarity.

## Proposed provider.ts structure

```ts
type ConjugationEntry = [Pronoun, string, Partial<ExpectedAnswer>?];

function buildDrillData(verb: string, tense: Tense, entries: ConjugationEntry[]): DrillData {
  return {
    verb, tense,
    items: entries.map(([pronoun, text, extra]) => ({
      prompt: { infinitive: verb, tense, pronoun },
      expectedAnswer: { text, ...extra },
    })),
  };
}

// Data becomes compact:
const verbData = {
  être: {
    présent: buildDrillData('être', 'présent', [
      ['je', 'suis'],
      ['tu', 'es'],
      // ...
    ]),
  },
  // ...
};

// Validators become single-check functions:
function validateTense(tense: string): Result<Tense> {
  if (!tense || !validTenses.includes(tense as Tense)) {
    return error(...);
  }
  return success(tense as Tense);
}

function synthesizePronouns(items: DrillItem[]): DrillItem[] {
  // extracted from getDrillData
}
```
<!-- SECTION:PLAN:END -->
