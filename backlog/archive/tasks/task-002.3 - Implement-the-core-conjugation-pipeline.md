---
id: TASK-002.3
title: Implement the core conjugation pipeline
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-02 19:05'
labels:
  - domain
  - engine
dependencies:
  - TASK-002.2
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-002
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the pure conjugation execution path that turns a normalized request into one or more conjugated forms. This task should capture the real rule flow from the reference engine, including verb classification, stem derivation, rule selection, irregular overrides, and any orthographic or morphology adjustments needed for correct output.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The engine can resolve a conjugation request into the expected output by applying the rule pipeline rather than relying on UI-layer branching.
- [ ] #2 The implementation handles both regular rule-driven cases and explicit exception or override paths represented in the ported reference logic.
- [ ] #3 Intermediate engine steps are organized into small internal helpers or modules so classification, override resolution, and output assembly remain testable in isolation.
- [ ] #4 The pipeline is deterministic and side-effect free so repeated invocations with the same input return the same result.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Implement request normalization and validation so downstream rule logic receives a canonical input shape.
2. Add verb classification and rule-resolution helpers that locate the correct stem, ending set, and exception path.
3. Port any irregular, spelling-change, or stem-change behavior from the reference engine into dedicated internal modules instead of scattering special cases.
4. Compose the helpers into a pure orchestration function that returns final conjugated forms without mutating shared state.
<!-- SECTION:PLAN:END -->

