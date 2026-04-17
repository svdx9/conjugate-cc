---
id: TASK-007.06
title: Wire orchestration layer into drill UI
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-17 13:46'
labels:
  - frontend
  - drills
  - integration
dependencies:
  - TASK-007.03
  - TASK-007.04
  - TASK-007.05
references:
  - backlog/docs/doc-002 - Frontend-Architecture.md
  - frontend/src/features/drills/hooks/useDrill.ts
  - frontend/src/features/drills/DrillsPage.tsx
parent_task_id: TASK-007
ordinal: 76000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Integrate the full orchestration stack: VerbSetProvider → DrillTestBuilder → ConjugationEngine → drill UI components.

This is more than a provider swap. It requires:
- DrillsPage wires up VerbSetProvider + DrillTestBuilder + DrillBehaviour
- useDrill hook refactored to accept a DrillProvider via parameter or SolidJS context (not the singleton import)
- Both SingleInputDrill and VerbDrill receive DrillData from the orchestration layer
- Session scoring displayed (from DrillTestBuilder)
- "Next verb" / "drill complete" flow driven by DrillTestBuilder
- StubDrillProvider remains available for testing

See: `backlog/docs/doc-002 - Frontend-Architecture.md` §2 (layer diagram), §8 (changes)
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 DrillsPage wires VerbSetProvider, DrillTestBuilder, DrillBehaviour, and ConjugationEngine together.
- [ ] #2 useDrill hook accepts a DrillProvider via parameter or SolidJS context — no singleton import.
- [ ] #3 Both SingleInputDrill and VerbDrill work with the orchestration layer.
- [ ] #4 Session scoring from DrillTestBuilder is displayed in the UI.
- [ ] #5 Next verb / drill complete flow is driven by DrillTestBuilder sequencing.
- [ ] #6 StubDrillProvider remains available for testing purposes.
- [ ] #7 No regressions in existing SingleInputDrill functionality.
<!-- AC:END -->
