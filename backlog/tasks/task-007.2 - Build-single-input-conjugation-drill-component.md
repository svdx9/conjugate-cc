---
id: TASK-007.2
title: Build single-input conjugation drill component
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-09 12:00'
labels:
  - frontend
  - ui
  - drills
dependencies:
  - TASK-007.1
references:
  - README.md
parent_task_id: TASK-007
ordinal: 72000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Build the basic conjugation drill component as a single-input interface. The component displays:

- The **pronoun** (e.g. "je")
- The **infinitive** (e.g. "etre")
- The **tense** (e.g. "Present")

The user sees a single text input where they type the correct conjugated verb phrase. On submit, the component shows whether the answer is correct or incorrect (with the expected answer if wrong).

This is the simplest drill case — one prompt, one input, one answer.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The component displays the pronoun, infinitive, and tense for the current drill prompt.
- [ ] #2 A single text input accepts the user's answer (verb phrase).
- [ ] #3 On submit, the answer is compared to the expected value and correct/incorrect feedback is shown.
- [ ] #4 If incorrect, the correct answer is displayed.
- [ ] #5 The component consumes data from the provider interface defined in TASK-007.1.
<!-- AC:END -->
