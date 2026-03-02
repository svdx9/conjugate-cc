---
id: TASK-1
title: Deliver MVP skeleton for the conjugation drill application
status: To Do
assignee: []
created_date: '2026-03-02 18:07'
updated_date: '2026-03-02 18:29'
labels:
  - mvp
  - planning
dependencies: []
references:
  - README.md
  - AGENTS.md
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the initial tracked work needed to stand up a minimum viable project for a conjugation drill application with a Go backend and a SolidJS frontend. The end state is a newcomer-clonable repository that can be started locally and shows an MVP front page in the browser.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Child tasks exist for shared setup, backend bootstrap, frontend bootstrap, landing page implementation, and final local MVP integration.
- [ ] #2 When all child tasks are complete, the repository can be started locally and the MVP front page is visible in a browser.
- [ ] #3 The tracked work includes test or verification expectations and documentation updates needed for an independent agent to complete each child task.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Complete `TASK-1.1` first to establish repository scaffold, root developer workflow, and contributor documentation.
2. Complete `TASK-1.2` next to add the first runnable Go backend service and wire the build metadata variables expected by the root `Makefile`.
3. Complete `TASK-1.3` after the shared scaffold is in place to add the first runnable SolidJS frontend under `frontend/`.
4. Complete `TASK-1.4` once the frontend shell exists to implement the MVP landing page within the established frontend structure.
5. Complete `TASK-1.5` last to connect frontend and backend into a single documented local startup workflow and verify the full MVP path end to end.
<!-- SECTION:PLAN:END -->
