---
id: TASK-001.1
title: Create repository scaffolding and local development workflow
status: To Do
assignee:
  - Codex
created_date: '2026-03-02 18:07'
updated_date: '2026-03-14 16:27'
labels:
  - mvp
  - setup
dependencies: []
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-001
ordinal: 31000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Establish the shared repository structure and baseline developer workflow for a Go backend plus SolidJS frontend application.

This task should leave the repo organized for follow-on backend and frontend implementation, with clear local startup instructions for contributors.

/backend for the go backend
/frontend for the frontend using: solidjs, solidjs router, tailwindcss, typescript, eslint, prettier and vite for building/testing
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The repository contains an agreed baseline structure for backend and frontend code, plus any root-level files needed for local development.
- [ ] #2 The project root contains a Makefile for backend workflows with invocable targets for test, lint, build, debug-build, and format.
- [ ] #3 The Makefile backend build targets pass the service git SHA and build time into the backend application entrypoint at build time.
- [ ] #4 Project documentation explains the intended app structure and how a contributor should install dependencies and start the project locally, including use of the Makefile.
- [ ] #5 The task includes a verification step showing the documented local workflow and Makefile commands match the repository layout and commands.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Inspect the empty repository baseline (`README.md`, root backlog files) and confirm this task should establish the first tracked project structure rather than adapt existing code.
2. Add the shared root scaffolding needed by follow-on tasks: create top-level `backend/` and `frontend/` directories plus any minimal root files required for local development and repository hygiene.
3. Add a `Makefile` in /backend focused on backend workflows with `test`, `lint`, `build`, `debug-build`, and `format` targets, keeping target names and command layout aligned with the Go tooling expected by later backend bootstrap work.
4. Expand `README.md` to document the intended repo layout, required local tools, how backend and frontend responsibilities are split, and how contributors should use the root `Makefile` plus frontend commands during local setup.
5. Verify the documented workflow against the repository contents by checking that the scaffolded paths exist, the `Makefile` exposes the documented targets, and the README instructions match the actual commands introduced in this task.
7. Leave backend service implementation, frontend app generation, and cross-stack startup wiring to `TASK-001` subtasks so this task stays limited to shared scaffolding and developer workflow setup.
<!-- SECTION:PLAN:END -->

## Implementation Notes


## Final Summary
