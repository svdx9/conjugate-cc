---
id: TASK-1.3
title: Bootstrap the SolidJS frontend application
status: Done
assignee:
  - Codex
created_date: '2026-03-02 18:07'
updated_date: '2026-03-02 19:27'
labels:
  - mvp
  - frontend
dependencies:
  - TASK-1.1
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-1
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the first runnable SolidJS frontend for the conjugation drill application under the repository's frontend area. The result should provide a clean development entry point that future feature work can build on.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The frontend can be started locally and renders an application shell in the browser.
- [x] #2 The initial frontend setup includes an appropriate verification step for the first render path used by the MVP.
- [x] #3 Documentation is updated with the command(s) needed to run and verify the frontend locally.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Bootstrap the first repo-local SolidJS frontend toolchain under `frontend/` with the minimal runnable files and npm scripts required for local development: `package.json`, `package-lock.json`, `vite.config.ts`, `tsconfig.json`, `index.html`, and the TypeScript entry files for a Solid app.
2. Create the deterministic frontend structure required by the repo-pinned SolidJS skill so later work can build without reshuffling files: `src/app/`, `src/features/`, `src/shared/`, `src/assets/`, and `src/styles/`.
3. Implement only a minimal application shell for this task, with a root app component and a simple shell screen that proves the frontend boots and renders in the browser, while leaving the real MVP front-page content and drill entry experience to `TASK-1.4`.
4. Keep the dependency set small and aligned with existing repo expectations by using SolidJS, TypeScript, Vite, and the standard Solid Vite plugin, plus only the test dependencies needed to verify the first render path.
5. Add the initial automated verification for the frontend bootstrap using Vitest and `@solidjs/testing-library`, with a component test that renders the shell and asserts the user-visible application content expected from the first render.
6. Add a production build script as part of the frontend package workflow so the bootstrap produces a buildable app, even though the primary acceptance target is the local dev startup path.
7. Update `README.md` so contributors can work from `frontend/`, install dependencies with `npm install`, start the app with `npm run dev`, and run the frontend verification command(s) introduced in this task.
8. Verify the task by confirming the documented frontend commands match the created files and scripts, that the shell render test passes, and that the frontend can be built and started locally through the documented npm workflow.
9. Keep explicit task boundaries: do not add backend integration, shared cross-stack env wiring, or the final landing-page experience in this task, because those belong to `TASK-1.5` and `TASK-1.4`.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Loaded the repo-pinned `javascript-typescript-solid-vite-web-frontend` skill during planning because `TASK-1.3` will introduce the initial SolidJS frontend under `frontend/`.

Research summary: the repository currently contains only `frontend/src/.gitkeep` for the frontend area, while the root `README.md` already reserves `frontend/` for a SolidJS app worked from with `npm install` and `npm run dev`.

Plan recorded as a proposed implementation plan pending user approval. No frontend code changes have been made for `TASK-1.3`.

User approved the recorded implementation plan on 2026-03-02. Proceeding with the frontend bootstrap implementation under the repo-pinned SolidJS frontend skill.

Implemented the frontend bootstrap under `frontend/` with a SolidJS + TypeScript + Vite toolchain, a minimal application shell component, project-local styling, and the initial Vitest render test for the first user-visible shell state.

Updated `README.md` so contributors can install frontend dependencies, run the Vite dev server, execute the render-path test, and build the frontend bundle from `frontend/`.

Verification: `cd frontend && npm install` succeeded and generated the tracked `package-lock.json` for the repo-local frontend workflow.

Verification: `cd frontend && npm run test` passed, confirming the shell heading, placeholder body copy, and placeholder action render correctly through the first app render path.

Verification: `cd frontend && npm run build` passed and produced the Vite production bundle for the bootstrapped frontend.

Verification: `cd frontend && npm run dev -- --host 127.0.0.1 --strictPort` reaches Vite startup but cannot bind to `127.0.0.1:5173` in this sandbox (`listen EPERM`), so live browser serving could not be exercised here even though the documented dev command is correct.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Bootstrapped the first runnable SolidJS frontend under `frontend/` with a minimal Vite + TypeScript toolchain, a reusable application shell, and a component test that verifies the initial render path. Updated the repository documentation so contributors can install dependencies, run the dev server, execute the frontend test, and build the app locally.
<!-- SECTION:FINAL_SUMMARY:END -->
