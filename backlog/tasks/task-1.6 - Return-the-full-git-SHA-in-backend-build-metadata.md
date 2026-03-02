---
id: TASK-1.6
title: Return the full git SHA in backend build metadata
status: To Do
assignee: []
created_date: '2026-03-02 19:03'
labels:
  - mvp
  - backend
  - follow-up
dependencies:
  - TASK-1.2
references:
  - README.md
  - Makefile
  - backlog/tasks/task-1.2 - Bootstrap-the-Go-backend-service.md
parent_task_id: TASK-1
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update the backend build metadata workflow so the build-info surface returns the full git commit SHA instead of the shortened SHA currently passed in by the root Makefile. This should preserve the existing build metadata endpoint contract apart from the SHA length change and keep the documentation and verification steps aligned with the new behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The backend build workflow passes the full git commit SHA into the application entrypoint instead of the shortened SHA.
- [ ] #2 The build metadata endpoint returns the full git SHA value consistently in local verification and automated tests.
- [ ] #3 Documentation and any verification steps that mention the build SHA are updated to reflect the full SHA behavior.
<!-- AC:END -->
