---
id: TASK-007.5
title: Port and adapt existing conjugation engine
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-09 12:00'
labels:
  - frontend
  - domain
  - engine
dependencies:
  - TASK-007.1
references:
  - README.md
parent_task_id: TASK-007
ordinal: 75000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Port an existing conjugation engine implementation (to be provided by the user) into the project. The engine must:

- Be reviewed and adapted to fit the project's patterns and the provider interface defined in TASK-007.1.
- Remain a pure domain module with no UI dependencies.
- Cover regular verbs (-er, -ir, -re) and common irregular verbs (etre, avoir, aller, faire, etc.).
- Handle compound tenses (auxiliary selection: etre vs avoir for passe compose).
- Handle reflexive verbs (reflexive pronoun agreement).

The existing code will need review and potential refactoring to conform to the drill contract.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The engine is ported into the project as a pure domain module with no UI imports.
- [ ] #2 The engine implements the provider interface from TASK-007.1.
- [ ] #3 Regular verbs (-er, -ir, -re) conjugate correctly across supported tenses.
- [ ] #4 Common irregular verbs (at minimum: etre, avoir, aller, faire) conjugate correctly.
- [ ] #5 Tests verify correctness for representative regular and irregular verbs.
<!-- AC:END -->
