# Review Issues — commits `5032b49` and `e47913f`

## `5032b49` — feat: task-007.03

### 1. Commented-out code in VerbDrill
- **File:** `frontend/src/features/drills/components/VerbDrill.tsx:63-67`
- Dead JSX block for pronoun labels. Should be removed or used.

### 2. Hardcoded verb/tense in FullDrillPage
- **File:** `frontend/src/features/drills/FullDrillPage.tsx:14`
- `drillProvider.getDrillData('être', 'présent')` is hardcoded. Likely a placeholder.

### 3. Redundant guard in VerbDrill
- **File:** `frontend/src/features/drills/components/VerbDrill.tsx:23-25`
- `<Show when={!props.drillData}>` will never render because `FullDrillPage` already guards `drillData()` before rendering `<VerbDrill>`.

### 4. Pronoun mapping fragility
- **File:** `frontend/src/features/drills/components/VerbDrill.tsx`
- `pronounOrder` uses display labels (`il/elle`, `ils/elles`) manually mapped to core pronouns (`il`, `ils`). The `correctAnswer()` derivation searches for both `il` and `elle` items separately. This dual-mapping logic is duplicated between the component and the hook's `corePronouns`.

### 5. AnswerInput lost keyboard interaction
- **File:** `frontend/src/features/drills/components/AnswerInput.tsx`
- `onSubmit`, `onReset`, and Enter key handling were removed. Quick Drill compensates with its own button, but Full Drill has no keyboard submit path (Enter does nothing).

### 6. AnswerInput lost auto-focus
- **File:** `frontend/src/features/drills/components/AnswerInput.tsx`
- The `createEffect` that focused the input on state changes was removed entirely. Neither consumer re-implements it.

### 7. lint-staged weakened
- **File:** `frontend/package.json`
- Typecheck (`npm run typecheck`) and tests (`npm test`) were removed from the pre-commit lint-staged config. Type errors and test failures can now be committed.

## `e47913f` — fix: enable Next button after submitting answer

### 8. canGoNext is redundant
- **File:** `frontend/src/features/drills/components/SingleInputDrill.tsx`
- `canGoNext()` returns `state.answerState() !== 'unanswered'`, identical to `showNext()`. When `showNext()` is true, `!canGoNext()` is always `false`, so the Next button is never disabled. The helper adds no value.
