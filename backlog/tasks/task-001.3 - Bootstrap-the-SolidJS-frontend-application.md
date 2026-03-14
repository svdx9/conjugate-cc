---
id: TASK-001.3
title: Bootstrap the SolidJS frontend application
status: To Do
assignee:
  - Codex
created_date: '2026-03-02 18:07'
updated_date: '2026-03-14 17:55'
labels:
  - mvp
  - frontend
dependencies:
  - TASK-001.1
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-001
ordinal: 30000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the first runnable SolidJS frontend for the conjugation drill application under the repository's frontend area. The result should provide a clean development entry point that future feature work can build on.

Include iint (eslint) and formatting (prettier) dev tools. Pay careful attention to the solidjs lint/formatting configuration.

Ensure typescript typecheck exists in dev tooling.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The frontend can be started locally and renders an application shell in the browser.
- [x] #2 The initial frontend setup includes an appropriate verification step for the first render path used by the MVP.
- [x] #3 Documentation is updated with the command(s) needed to run and verify the frontend locally.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Bootstrap the first repo-local SolidJS frontend toolchain under `frontend/` with the minimal runnable files and npm scripts required for local development: `package.json`, `package-lock.json`, `vite.config.ts`, `tsconfig.json`, `index.html`, `.prettierrc`, `eslint.config.js`, `and the TypeScript entry files for a Solid app.
2. Create the deterministic frontend structure required by the repo-pinned SolidJS skill so later work can build without reshuffling files: `src/app/`, `src/features/`, `src/shared/`, `src/assets/`, and `src/styles/`.
3. Implement only a minimal application shell for this task, with a root app component and a simple shell screen that proves the frontend boots and renders in the browser, while leaving the real MVP front-page content and drill entry experience to `TASK-001.4`.
4. Keep the dependency set small and aligned with existing repo expectations by using SolidJS, TypeScript, Vite, and the standard Solid Vite plugin, plus only the test dependencies needed to verify the first render path.
5. Add the initial automated verification for the frontend bootstrap using Vitest and `@solidjs/testing-library`, with a component test that renders the shell and asserts the user-visible application content expected from the first render.
6. Add a production build script as part of the frontend package workflow so the bootstrap produces a buildable app, even though the primary acceptance target is the local dev startup path.
7. Update `README.md` so contributors can work from `frontend/`, install dependencies with `npm install`, start the app with `npm run dev`, and run the frontend verification command(s) introduced in this task.
8. Verify the task by confirming the documented frontend commands match the created files and scripts, that the shell render test passes, and that the frontend can be built and started locally through the documented npm workflow.
9. Keep explicit task boundaries: do not add backend integration, shared cross-stack env wiring, or the final landing-page experience in this task, because those belong to `TASK-001.5` and `TASK-001.4`.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Final Summary
<!-- SECTION:NOTES:END -->
