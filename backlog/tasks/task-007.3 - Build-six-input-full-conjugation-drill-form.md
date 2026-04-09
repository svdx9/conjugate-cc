---
id: TASK-007.3
title: Build six-input full conjugation drill form
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-09 12:00'
labels:
  - frontend
  - ui
  - drills
dependencies:
  - TASK-007.2
references:
  - README.md
parent_task_id: TASK-007
ordinal: 73000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Expand the single-input drill into a full conjugation form showing all 6 standard French pronouns:

- je
- tu
- il/elle
- nous
- vous
- ils/elles

Each pronoun has its own text input. The form displays the infinitive and tense at the top. The user fills in all 6 conjugated forms and submits them together.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The form displays the infinitive and tense prominently at the top.
- [ ] #2 All 6 pronouns are shown with a corresponding text input per row.
- [ ] #3 The layout is readable on desktop and mobile with clear alignment between pronoun and input.
- [ ] #4 A submit button sends all 6 answers for validation.
- [ ] #5 The form consumes data from the same provider interface defined in TASK-007.1.
<!-- AC:END -->
