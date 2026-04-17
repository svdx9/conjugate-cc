---
id: TASK-007.07
title: Clean up drill code after orchestration integration
status: To Do
assignee: []
created_date: '2026-04-13 15:30'
updated_date: '2026-04-17 13:46'
labels:
  - frontend
  - refactor
  - drills
dependencies:
  - TASK-007.06
references:
  - backlog/docs/doc-002 - Frontend-Architecture.md
  - frontend/src/features/drills/components/AnswerInput.tsx
  - frontend/src/shared/types.ts
parent_task_id: TASK-007
priority: medium
ordinal: 77000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Post-integration cleanup of the drills feature. The original refactoring scope (buildDrillData helper, synthesizePronouns extraction) has already been done. Remaining work focuses on bugs and quality issues surfaced during development:

**Bugs:**
- Answer checking must support `ExpectedAnswer.alternates` (multiple valid conjugations)
- Elision logic missing accented French vowels (àâäéèêëïîôùûü)

**Quality:**
- Remove any remaining Windsurf watermark comments
- Prune verbose comments that restate code
- Remove unused `CommonErrorCode` type from `shared/types.ts` if still present
- Ensure AnswerInput feedback uses design system semantic tokens (not hardcoded green/red)

**Deferred from original plan** (already done or overtaken):
- buildDrillData helper — done
- synthesizePronouns extraction — done
- validateString removal — done
- verbData compact format — done
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Answer checking supports ExpectedAnswer.alternates for multiple valid conjugations.
- [ ] #2 Elision logic handles accented French vowels.
- [ ] #3 No Windsurf watermark comments remain.
- [ ] #4 Feedback UI uses design system semantic tokens, not hardcoded colors.
- [ ] #5 CommonErrorCode type removed from shared/types.ts if unused.
- [ ] #6 All existing tests continue to pass.
<!-- AC:END -->



## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
<!-- SECTION:PLAN:BEGIN -->
## Issues Identified

1. Windsurf watermark comments (provider.ts)
Three `/**** Windsurf Command ****/` banners at lines 159, 176, 222 — noise from tooling that should never be committed.

2. Excessive `as Tense` casts in verbData
Every single DrillItem repeats `'présent' as Tense` because the outer record is typed as `Record<string, Record<string, DrillData>>` instead of using `Tense` keys. This results in ~30 unnecessary casts. Fix by either:
- Typing verbData as `Record<string, Partial<Record<Tense, DrillData>>>` and casting at the outer level, or
- Using a `buildDrillData(verb, tense, entries)` helper that constructs DrillData from compact tuples like `['je', 'suis']`.

3. Over-engineered validateString
The `typeof str !== 'string'` check is pointless — TypeScript already enforces the parameter type. The entire `validateString` function exists just to check emptiness and is called by three other validators. Inline the empty check directly.

4. Redundant validation chaining
`validateTense` calls `validateString` then checks `validTenses.includes()`. The string validation is unnecessary since `includes()` on an empty/falsy string just returns `false`. Same for `validatePronoun`. Collapse each into a single check.

5. Verbose/obvious comments
Many comments just restate what the code does:
- `// Return the validated string` above `return success(str)`
- `// Check if the string is empty or not defined` above `if (!str ...)`
- `// Normalize the verb string` above `const normalizedVerb = ...`
Remove these. Keep comments that explain *why*, not *what*.

6. Unused CommonErrorCode type (shared/types.ts)
`CommonErrorCode` is defined at shared/types.ts:43-49 but never referenced anywhere. Remove it.

7. verbData repetition
Each DrillItem in verbData manually repeats the verb name and tense that are already encoded in the parent keys. A `buildDrillData` helper taking compact `[pronoun, text, extras?]` tuples would cut the data block roughly in half.

8. Inline pronoun synthesis
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

9. **BUG: Score increments on every submit, not just first**
- File: `frontend/src/features/drills/hooks/useDrill.ts:112-115`
- Issue: If user clicks Submit repeatedly, score increments each time
- Fix: Add check to only increment score on first submission — track whether current question has been answered using `answerState() !== 'unanswered'` before incrementing

```typescript
const submitAnswer = () => {
  const item = currentItem();
  if (!item) return;
  const isCorrect = checkAnswer(userAnswer(), item.expectedAnswer.text);
  setAnswerState(isCorrect ? 'correct' : 'incorrect');
  setScore((prev) => ({
    correct: prev.correct + (isCorrect ? 1 : 0),
    total: prev.total + 1,
  }));
};
```

10. **BUG: goToNext doesn't advance index, it loads new random drill**
- File: `frontend/src/features/drills/hooks/useDrill.ts:118-129`
- Issue: Named "goToNext" but loads completely random new drill instead of advancing to next pronoun in current drill. Doesn't match progress indicator
- Fix: Rename to `nextQuestion`, implement proper drill progression — advance `currentIndex` until all items exhausted, then show completion

```typescript
const nextQuestion = () => {
  const verbs = drillProvider.getAvailableVerbs();
  const randomVerb = verbs[Math.floor(Math.random() * verbs.length)];
  // ... loads completely new drill — WRONG
};
```

11. **Duplicate feedback UI**
- Location: `AnswerInput.tsx:86-101` vs `ResultFeedback.tsx`
- Issue: `AnswerInput` embeds its own feedback UI (correct/incorrect alerts) while separate `ResultFeedback` component exists but is unused in `SingleInputDrill`
- Fix: Remove feedback from `AnswerInput`, use `ResultFeedback` as sibling component

12. **Accented vowels not handled in elision**
- Location: `AnswerInput.tsx:3-4`
- Issue: Only checks ASCII vowels, missing French accented vowels

```typescript
const VOWELS = 'aeiouAEIOU';  // WRONG
const VOWEL_LIKE_CONSONANTS = 'hH';
```
- Fix: Expand to `'aeiouàâäéèêëïîôùûüAEIOUÀÂÄÉÈÊËÏÎÔÙÛÜ'`

13. **onReset prop optional but required**
- Location: `AnswerInput.tsx:18-27`
- Issue: `onReset` is optional but component logic assumes it exists when `answerState !== 'unanswered'`
- Fix: Make `onReset` required, add runtime checks

14. **Duplicate elision logic**
- Location: Multiple files with same VOWELS constant defined
- Fix: Extract to shared utility

15. **Hardcoded green colors not in design system**
- Location: `AnswerInput.tsx:87`
- Issue: Uses arbitrary Tailwind colors instead of semantic tokens

```typescript
<div class="... border-green-500 bg-green-50 ... dark:bg-green-900/20">
  <p class="... text-green-700 ... dark:text-green-400">Correct!</p>
```

16. **Type assertion bypassing safety**
- Location: `useDrill.ts:57`
- Issue: Casts `t` to `Tense` without validation
```typescript
const result = drillProvider.getDrillData(v, t as Tense);
```
- Fix: Accept `Tense` type directly or validate before calling

17. **Missing error code type narrowing**
- Location: `useDrill.ts:66`
- Issue: `code` typed as `string`, not literal union
```typescript
if (result.code === 'NOT_FOUND' && result.details?.availableTenses) {
```

18. **Unused type export**
- Location: `shared/types.ts:43-49`
- Issue: `CommonErrorCode` type defined but never used in Result type

19. **Missing cleanup in effect with dynamic dependency**
- Location: `useDrill.ts:54-77`
- Issue: `loadDrillData` called during render, pattern could break with async

20. **Unnecessary createMemo for simple derived values**
- Location: `SingleInputDrill.tsx:14-18`
- Issue: Simple derived values don't need explicit memoization
```typescript
const totalQuestions = createMemo(() => state.drillData()?.items.length ?? 0);
const questionNumber = createMemo(() => state.currentItem() ? state.currentIndex() + 1 : 0);
```

21. **Effect with missing cleanup consideration**
- Location: `AnswerInput.tsx:32-36`
- Issue: Effect could cause focus stealing
```typescript
createEffect(() => {
  if (props.answerState === 'unanswered' && inputRef) {
    inputRef.focus();
  }
});
```

22. **Array shuffling bias**
- Location: `useDrill.ts:59`
- Issue: `sort(() => Math.random() - 0.5)` has bias
```typescript
const shuffled = [...result.data.items].sort(() => Math.random() - 0.5);
```
- Fix: Use Fisher-Yates shuffle

23. **Progress indicator creates array on every render**
- Location: `SingleInputDrill.tsx:59-71`
- Issue: `Array.from({ length: totalQuestions() })` creates garbage every render
- Fix: Memoize the array

24. **VerbDropdown missing focus management**
- Location: `VerbDropdown.tsx91-105`
- Issue: Options don't have tabIndex sync with highlightedIndex. Keyboard nav updates highlightedIndex but doesn't move focus — screen readers won't announce changes

```typescript
<button
  role="option"
  aria-selected={option === props.value}
  onClick={(e) => {...}}
  onMouseEnter={() => setHighlightedIndex(index())}
  class="w-full px-3 py-2 text-left text-foreground hover:bg-accent focus:bg-accent"
>
```

- Fix: Add `tabIndex={0}`, sync focus with highlightedIndex, implement focus trap, add `aria-activedescendant`

25. **Missing aria-label on progress indicators**
- Location: `SingleInputDrill.tsx:59-71`
- Issue: Progress dots purely visual, no ARIA equivalent
- Fix: Wrap in `role="progressbar"` with `aria-valuenow`, `aria-valuemax`

26. **Color contrast on success state**
- Location: `AnswerInput.tsx:87`
- Issue: `text-green-400` on `bg-green-900/20` may not meet WCAG AA (4.5:1) in dark mode
- Fix: Use `text-green-300` or add semantic success colors to design system

27. **No tests for useDrill hook**
- Contains critical business logic with state management

28. **No tests for UI components**
- `AnswerInput`, `SingleInputDrill`, `VerbDropdown` untested

29. **No tests for elision logic**
- `canElide` and `getElidedPronoun` functions contain language-specific logic

30. **Test assertion anti-pattern**
- Location: `provider.test.ts:12-14`
```typescript
if (isError(result)) {
  expect.fail(`Expected data but got error: ${result.error}`);
}
```
- Fix: Use type guards with early returns

31. **Unused VerbDropdown**
- Location: `DrillsPage.tsx:20`
- Issue: Component implemented but not used. Drill hardcoded to `verb="être"` with no way to change

32. **Missing trailing newlines**
- Files: `AnswerInput.tsx`, `ResultFeedback.tsx`, `VerbDropdown.tsx`, `components/index.ts`
-

33. **VerbDropdown not integrated**
- Despite being built, the dropdown is not wired to DrillsPage — appears intentional for MVP but should be integrated or explicitly deferred

---
<!-- SECTION:PLAN:END -->
