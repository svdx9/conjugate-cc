---
id: TASK-1.2
title: Bootstrap the Go backend service
status: To Do
assignee: []
created_date: '2026-03-02 18:07'
updated_date: '2026-03-02 18:13'
labels:
  - mvp
  - backend
dependencies:
  - TASK-1.1
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-1
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the first runnable backend service for the conjugation drill application under the repository's backend area. The service should be suitable for local development and provide the minimal API surface needed for an MVP frontend to confirm the backend is alive.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The backend can be started locally and exposes a minimal HTTP surface for health or status checking.
- [ ] #2 The backend exposes an endpoint that returns the service git SHA and build time.
- [ ] #3 The git SHA and build time values are injected into the application entrypoint at build time rather than hard-coded.
- [ ] #4 Basic backend tests or equivalent automated verification cover the initial server behavior that the MVP relies on, including the build metadata endpoint.
- [ ] #5 Documentation is updated with the command(s) needed to run and verify the backend locally.
<!-- AC:END -->
