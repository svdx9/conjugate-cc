---
id: TASK-2.1
title: Establish the conjugation engine package boundary
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-02 19:05'
labels:
  - domain
  - frontend
  - engine
dependencies:
  - TASK-1.3
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-2
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define the package location, directory shape, and public surface for the conjugation engine so it is explicitly treated as pure domain logic within this repository. The boundary should mirror the intent of the existing `conjuflo` implementation while making it impossible for rendering or browser concerns to leak into the engine.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 A dedicated engine area exists under the application codebase, for example `frontend/src/domain/conjugator/`, with a structure that clearly separates public entrypoints from internal rule/data modules.
- [ ] #2 The engine package imports no UI framework modules, DOM APIs, router code, styling assets, or application state containers.
- [ ] #3 A short architecture note documents what belongs inside the engine boundary, what must stay outside it, and how consumers are expected to call into it.
- [ ] #4 A lightweight verification step guards the boundary, such as import-boundary tests, lint rules, or a package-level review checklist committed alongside the module.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Choose the concrete package location for the engine and create the top-level directories for API, rules, data, and tests.
2. Define the initial public entrypoint and any internal-only modules so future work has a stable import path from day one.
3. Add documentation describing the purity constraint for the engine and explicitly ban UI or browser dependencies from this package.
4. Add the smallest enforceable boundary check the current toolchain supports so regressions are caught automatically.
<!-- SECTION:PLAN:END -->

