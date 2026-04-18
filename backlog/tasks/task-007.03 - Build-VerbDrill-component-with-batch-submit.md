---
id: TASK-007.03
title: Build VerbDrill component with batch submit
status: Done
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-17 16:32'
labels:
  - frontend
  - ui
  - drills
dependencies:
  - TASK-007.02
references:
  - backlog/docs/doc-002 - Frontend-Architecture.md
  - frontend/src/features/drills/components/AnswerInput.tsx
  - frontend/src/features/drills/components/SingleInputDrill.tsx
parent_task_id: TASK-007
ordinal: 73000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Build the VerbDrill component (`components/VerbDrill.tsx`) — the full conjugation drill showing all 6 pronoun rows with stacked inputs and batch submission.

This component receives a `DrillData` prop (the type defined in `types.ts`) and renders 6 pronoun rows (je, tu, il/elle, nous, vous, ils/elles), each reusing the existing `AnswerInput` component. The user fills all 6 fields and submits once via a single Submit button. After submission, per-row correct/incorrect feedback is shown inline.

Note: this component only depends on the `DrillData` type, not on the DrillTestBuilder implementation (built in TASK-007.04). DrillTestBuilder will later produce `DrillData` and pass it to VerbDrill.

Requires a new `useVerbDrill` hook managing 6 answer signals, batch submission, and per-row correctness state.

See: `backlog/docs/doc-002 - Frontend-Architecture.md` §3 (UI Layer)
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The VerbDrill component displays verb+tense header and 6 pronoun rows with one AnswerInput per row.
- [x] #2 Each row reuses the existing AnswerInput component — no duplicate input implementation.
- [x] #3 A single Submit button validates all 6 answers as a batch.
- [x] #4 Per-row correct/incorrect feedback is shown inline after submission.
- [x] #5 useVerbDrill hook manages 6 answer signals, batch submission, and per-row correctness state.
- [x] #6 Layout is readable on desktop and mobile with clear pronoun/input alignment.
- [x] #7 Component receives DrillData as props — no direct provider imports.
<!-- AC:END -->
