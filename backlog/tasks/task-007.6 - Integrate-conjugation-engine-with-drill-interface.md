---
id: TASK-007.6
title: Integrate conjugation engine with drill interface
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-09 12:00'
labels:
  - frontend
  - drills
  - integration
dependencies:
  - TASK-007.4
  - TASK-007.5
references:
  - README.md
parent_task_id: TASK-007
ordinal: 76000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Replace the stub conjugation data provider with the ported conjugation engine. The drill interface should now be able to present any French verb in any supported tense, rather than being limited to the hardcoded stub verbs.

This task wires the engine provider into the drill components without changing the UI — the provider interface swap should be seamless.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The drill interface uses the engine provider instead of the stub provider.
- [ ] #2 Drills can be generated for any verb the engine supports, not just stub verbs.
- [ ] #3 The single-input drill (TASK-007.2) and six-input form (TASK-007.3) both work with the engine provider.
- [ ] #4 No UI changes are required — the swap is at the provider level only.
- [ ] #5 The stub provider remains available for testing purposes.
<!-- AC:END -->
