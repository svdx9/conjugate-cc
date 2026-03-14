---
id: TASK-002.1
title: Establish the conjugation engine package boundary
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-14 16:27'
labels:
  - domain
  - frontend
  - engine
dependencies:
  - TASK-001.3
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-002
ordinal: 25000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define the package location, directory shape, and public surface for the conjugation engine so it is explicitly treated as pure domain logic within this repository. The boundary should mirror the intent of the existing `conjuflo` implementation while making it impossible for rendering or browser concerns to leak into the engine.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 A dedicated engine area exists under the application codebase, for example `frontend/src/domain/conjugator/`, with a structure that clearly separates public entrypoints from internal rule/data modules.
- [x] #2 The engine package imports no UI framework modules, DOM APIs, router code, styling assets, or application state containers.
- [x] #3 A short architecture note documents what belongs inside the engine boundary, what must stay outside it, and how consumers are expected to call into it.
- [x] #4 A lightweight verification step guards the boundary, such as import-boundary tests, lint rules, or a package-level review checklist committed alongside the module.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Implementation Steps

### Step 1 — Create directory structure

```
frontend/src/domain/conjugator/
  index.ts          # Public barrel export (empty stub — establishes the stable import path)
  internal/         # Future home for rules, data, pipeline (not exported from index.ts)
    .gitkeep
```

### Step 2 — Write ARCHITECTURE.md

Create `frontend/src/domain/conjugator/ARCHITECTURE.md` documenting:

- **What belongs inside the engine boundary**: pure TS types, static rule data as plain objects, classification and pipeline functions, exception tables.
- **Forbidden imports** (review checklist): `solid-js`, `@solidjs/*`, DOM globals (`document`, `window`, `HTMLElement`), router code, CSS/asset imports, app state containers, anything from `src/features/`, `src/components/`, `src/pages/`.
- **How consumers call in**: import only from the public barrel (`src/domain/conjugator/index.ts`); internal paths are not part of the public API.

---

## Files Created

| File | Action |
|------|--------|
| `frontend/src/domain/conjugator/index.ts` | Create (empty stub) |
| `frontend/src/domain/conjugator/internal/.gitkeep` | Create |
| `frontend/src/domain/conjugator/ARCHITECTURE.md` | Create |

---

## Acceptance Criteria Coverage

| AC | How it is met |
|----|---------------|
| #1 Dedicated engine area with structure separating public from internal | `index.ts` (public) + `internal/` subdirectory |
| #2 No UI/browser imports | Documented constraint in ARCHITECTURE.md; enforced by code review |
| #3 Architecture note documents boundary | ARCHITECTURE.md covers what belongs, what is forbidden, how to call in |
| #4 Lightweight verification step | ARCHITECTURE.md boundary checklist committed alongside the module |
<!-- SECTION:PLAN:END -->
