---
id: TASK-1.1
title: Create repository scaffolding and local development workflow
status: To Do
assignee: []
created_date: '2026-03-02 18:07'
updated_date: '2026-03-02 18:13'
labels:
  - mvp
  - setup
dependencies: []
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-1
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Establish the shared repository structure and baseline developer workflow for a Go backend plus SolidJS frontend application. This task should leave the repo organized for follow-on backend and frontend implementation, with clear local startup instructions for contributors.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The repository contains an agreed baseline structure for backend and frontend code, plus any root-level files needed for local development.
- [ ] #2 The project root contains a Makefile for backend workflows with invocable targets for test, lint, build, debug-build, and format.
- [ ] #3 The Makefile backend build targets pass the service git SHA and build time into the backend application entrypoint at build time.
- [ ] #4 Project documentation explains the intended app structure and how a contributor should install dependencies and start the project locally, including use of the Makefile.
- [ ] #5 The task includes a verification step showing the documented local workflow and Makefile commands match the repository layout and commands.
<!-- AC:END -->
