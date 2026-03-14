---
id: TASK-003.3
title: Add show-answer and hide-answer state to the drill form
status: To Do
assignee: []
created_date: '2026-03-02 23:31'
updated_date: '2026-03-14 16:26'
labels:
  - frontend
  - ui
  - interaction
dependencies:
  - TASK-003.2
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-003
ordinal: 7000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add the answer reveal interaction to the drill form. When the learner presses `Show answer`, the form should replace the visible input content with the correct `etre` answers in place, and the control label should change to `Hide answer`; when the learner presses `Hide answer`, their original typed responses should reappear in the same inputs.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Pressing `Show answer` reveals the six correct answers inline where the learner inputs normally appear and updates the control label to `Hide answer`.
- [ ] #2 Pressing `Hide answer` restores the learner's previously typed responses without clearing, reordering, or mutating them.
- [ ] #3 The task includes verification that repeated toggling between shown and hidden answers preserves the learner's draft state across all six inputs.
<!-- AC:END -->
