---
id: TASK-1.1
title: Create repository scaffolding and local development workflow
status: Done
assignee:
  - Codex
created_date: '2026-03-02 18:07'
updated_date: '2026-03-02 18:31'
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
- [x] #1 The repository contains an agreed baseline structure for backend and frontend code, plus any root-level files needed for local development.
- [x] #2 The project root contains a Makefile for backend workflows with invocable targets for test, lint, build, debug-build, and format.
- [x] #3 The Makefile backend build targets pass the service git SHA and build time into the backend application entrypoint at build time.
- [x] #4 Project documentation explains the intended app structure and how a contributor should install dependencies and start the project locally, including use of the Makefile.
- [x] #5 The task includes a verification step showing the documented local workflow and Makefile commands match the repository layout and commands.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Inspect the empty repository baseline (`README.md`, root backlog files) and confirm this task should establish the first tracked project structure rather than adapt existing code.
2. Add the shared root scaffolding needed by follow-on tasks: create top-level `backend/` and `frontend/` directories plus any minimal root files required for local development and repository hygiene.
3. Add a root `Makefile` focused on backend workflows with `test`, `lint`, `build`, `debug-build`, and `format` targets, keeping target names and command layout aligned with the Go tooling expected by later backend bootstrap work.
4. Wire backend build metadata into the `build` and `debug-build` targets via `-ldflags` so the backend entrypoint can receive git SHA and build time values once `TASK-1.2` adds the Go service variables.
5. Expand `README.md` to document the intended repo layout, required local tools, how backend and frontend responsibilities are split, and how contributors should use the root `Makefile` plus frontend commands during local setup.
6. Verify the documented workflow against the repository contents by checking that the scaffolded paths exist, the `Makefile` exposes the documented targets, and the README instructions match the actual commands introduced in this task.
7. Leave backend service implementation, frontend app generation, and cross-stack startup wiring to `TASK-1.2`, `TASK-1.3`, and `TASK-1.5` respectively so this task stays limited to shared scaffolding and developer workflow setup.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research summary: the working tree currently contains only `README.md`, `AGENTS.md`, and Backlog metadata, so `TASK-1.1` needs to define the initial repository scaffold from scratch rather than refactor existing backend or frontend code.

Plan recorded as a proposed implementation plan pending user approval. No code changes made.

User approved the recorded implementation plan on 2026-03-02. Proceeding with repository scaffolding, root Makefile, and documentation only.

Implemented the shared repository scaffold for follow-on work: added `backend/cmd/`, `backend/internal/`, `backend/tools/`, and `frontend/src/` along with root `.gitignore`, root `Makefile`, and expanded `README.md` guidance.

Verification: `find backend frontend -maxdepth 2 -type d | sort` confirmed the scaffolded directory layout.

Verification: `make -n format lint test build debug-build` showed the documented root workflow targets and confirmed that `build` and `debug-build` inject `main.serviceGitSHA` and `main.serviceBuildTime` via `-ldflags`.

Verification: `make build` fails fast with the expected guard message because `backend/go.mod` is intentionally left for `TASK-1.2`, which keeps this task scoped to shared scaffolding rather than backend service bootstrap.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Created the initial shared repository scaffold and local developer workflow for the MVP. Added root `.gitignore`, a backend-focused root `Makefile` with test/lint/build/debug-build/format targets and build metadata injection, empty backend/frontend structure for follow-on tasks, and README documentation covering layout, prerequisites, backend workflow, frontend workflow, and verification expectations.
<!-- SECTION:FINAL_SUMMARY:END -->
