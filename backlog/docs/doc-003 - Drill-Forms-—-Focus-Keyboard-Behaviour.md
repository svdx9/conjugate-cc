---
id: doc-003
title: Drill Forms — Focus & Keyboard Behaviour
type: other
created_date: '2026-04-18 18:37'
---
# Drill Forms — Focus & Keyboard Behaviour

Documents the focus management and keyboard interaction patterns for the Quick Drill and Full Drill forms.

## Quick Drill (SingleInputDrill)

- **On load**: The answer input auto-focuses so the user can start typing immediately.
- **On submit**: Focus moves from the input to the "Next" button. This lets the user press Enter/Space to advance without reaching for the mouse.
- **On next**: Focus returns to the input for the new question.

### Keyboard flow
`type answer` → `Enter` (submit) → focus moves to **Next** button → `Enter` (next question) → focus returns to input → repeat

## Full Drill (VerbDrill)

- **On load**: Only the first row (`je`) auto-focuses. The remaining 5 rows (`tu`, `il`, `nous`, `vous`, `ils`) do **not** auto-focus.
- **On submit**: Focus moves to the "Next" button. The user can review scores and press Enter/Space to restart.
- **On next (reset)**: Focus returns to the first row (`je`), not the last row.
- **Enter key in inputs**: Does **not** submit the form. Users must use the Submit button. This prevents accidental early submission when filling out 6 fields.

### Keyboard flow
`Tab` through inputs filling answers → click/tab to **Submit** → focus moves to **Next** button → `Enter` (reset) → focus returns to `je` input → repeat

## Implementation Details

- `AnswerInput` accepts an `autoFocus` prop (default: `true`). When `false`, the component's `createEffect` skips focusing.
- `VerbDrill` passes `autoFocus={pronounKey === 'je'}` so only the first row captures focus.
- `SingleInputDrill` omits `autoFocus`, inheriting the default `true` behaviour.
- Both drills use a `buttonRef` + `createEffect` to move focus to the action button after submission.
