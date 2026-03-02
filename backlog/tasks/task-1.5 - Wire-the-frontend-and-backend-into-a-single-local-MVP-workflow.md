---
id: TASK-1.5
title: Wire the frontend and backend into a single local MVP workflow
status: To Do
assignee: []
created_date: '2026-03-02 18:07'
updated_date: '2026-03-02 18:07'
labels:
  - mvp
  - integration
dependencies:
  - TASK-1.2
  - TASK-1.3
  - TASK-1.4
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-1
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Connect the initial frontend and backend so the repository behaves like one MVP application during local development. The outcome should make it straightforward for a contributor to run both parts together and confirm that the front page is backed by a live application stack.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 A contributor can follow repository documentation to start the backend and frontend together and view the MVP front page locally.
- [ ] #2 The MVP front page demonstrates live integration with the backend through a minimal user-visible signal such as status, readiness, or configuration data.
- [ ] #3 The integration path has a repeatable verification step that confirms the local stack is working end to end.
<!-- AC:END -->
