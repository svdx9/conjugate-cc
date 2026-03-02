---
id: TASK-2.5
title: Build deterministic engine tests and parity fixtures
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-02 19:05'
labels:
  - domain
  - engine
  - test
dependencies:
  - TASK-2.4
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-2
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Capture the expected behavior of the conjugation engine in automated tests so the ported logic remains trustworthy as the application grows. The tests should emphasize parity with the existing implementation, deterministic outputs, and coverage of both representative and troublesome verb families.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Automated tests cover representative regular, irregular, and edge-case conjugation scenarios that reflect the behavior of the reference engine.
- [ ] #2 The test suite exercises both the public API and selected internal rule helpers where direct coverage improves confidence in tricky transformations.
- [ ] #3 Any parity fixtures or snapshot-style expectations are stored in version control in a readable form that supports review.
- [ ] #4 The verification workflow is documented so future engine changes can prove they did not regress existing conjugation behavior.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Assemble a representative fixture set from the existing implementation, prioritizing verbs and forms that cover the main rule families plus known exceptions.
2. Add tests for the public API contracts first, then supplement them with focused helper-level tests where failures would otherwise be hard to diagnose.
3. Keep expected outputs deterministic and reviewable, avoiding opaque generated blobs where a table or fixture file would be clearer.
4. Document the exact test command and what evidence constitutes parity with the imported engine behavior.
<!-- SECTION:PLAN:END -->

