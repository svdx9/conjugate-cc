---
id: TASK-3.1
title: Define the stubbed etre drill contract
status: To Do
assignee: []
created_date: '2026-03-02 23:31'
updated_date: '2026-03-02 23:32'
labels:
  - frontend
  - drills
  - data
dependencies:
  - TASK-1.5
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-3
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define the temporary frontend-side data contract for a single conjugation drill round using only the verb `etre`. The contract should provide the infinitive prompt, the six pronouns in display order, the corresponding correct forms, and the score inputs needed by the later results page, while staying simple enough to replace with a real conjugation engine later.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The task defines a single temporary drill dataset for `etre` that includes six pronouns and six matching correct answers in the exact order expected by the drill form.
- [ ] #2 The contract makes clear that the learner sees the infinitive prompt and enters the missing verb endings or full forms into six separate inputs.
- [ ] #3 The task includes verification that the stub contract can drive both the in-form answer reveal state and the submitted results summary without requiring the future conjugation engine.
<!-- AC:END -->
