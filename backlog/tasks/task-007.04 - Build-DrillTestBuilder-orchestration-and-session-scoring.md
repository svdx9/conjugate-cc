---
id: TASK-007.04
title: Build DrillTestBuilder orchestration and session scoring
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-17 13:46'
labels:
  - frontend
  - drills
  - orchestration
dependencies:
  - TASK-007.03
references:
  - backlog/docs/doc-002 - Frontend-Architecture.md
  - frontend/src/features/drills/types.ts
parent_task_id: TASK-007
ordinal: 74000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Build the DrillTestBuilder (`drill-test-builder.ts`), DrillBehaviour (`drill-behaviour.ts`), and VerbSetProvider (`verb-set-provider.ts`) — the orchestration layer that sequences verbs, selects the next drill, and tracks session scoring.

**DrillTestBuilder** is not a UI component. It:
- Receives a VerbSetProvider (what to drill) and a DrillBehaviour (how to select)
- Sequences through verb+tense pairs across the session
- Requests conjugation data from the ConjugationEngine for the current pair
- Tracks session scoring (correct/total, per-verb results)
- Signals "drill complete" when the set is exhausted

**DrillBehaviour** is a strategy object:
- `random` — picks randomly from the set
- `sequential` — iterates in order; returns null when exhausted

**VerbSetProvider**:
- Returns an array of `{ verb, tense }` pairs
- Stub implementation with hardcoded array for now

See: `backlog/docs/doc-002 - Frontend-Architecture.md` §4-6
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 DrillTestBuilder sequences through verb+tense pairs from a VerbSetProvider.
- [ ] #2 DrillBehaviour supports random and sequential selection strategies.
- [ ] #3 StubVerbSetProvider returns a hardcoded array of verb+tense pairs.
- [ ] #4 Session scoring tracks correct/total and per-verb results (SessionScore, VerbDrillResult types).
- [ ] #5 DrillTestBuilder signals drill complete when the set is exhausted (sequential mode).
- [ ] #6 Unit tests verify sequencing, scoring, and both behaviour strategies.
- [ ] #7 No UI or framework imports — pure TypeScript modules.
<!-- AC:END -->
