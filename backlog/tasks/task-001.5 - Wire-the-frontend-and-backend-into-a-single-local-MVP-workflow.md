---
id: TASK-001.5
title: Wire the frontend and backend into a single local MVP workflow
status: To Do
assignee: []
created_date: '2026-03-02 18:07'
updated_date: '2026-03-14 16:26'
labels:
  - mvp
  - integration
dependencies:
  - TASK-001.2
  - TASK-001.3
  - TASK-001.4
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-001
ordinal: 23000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Connect the initial frontend and backend so the repository behaves like one MVP application during local development. The outcome should make it straightforward for a contributor to run both parts together and confirm that the front page is backed by a live application stack.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 A contributor can follow repository documentation to start the backend and frontend together and view the MVP front page locally.
- [ ] #2 The MVP front page demonstrates live integration with the backend through a minimal user-visible signal such as status, readiness, or configuration data.
- [ ] #3 The integration path has a repeatable verification step that confirms the local stack is working end to end.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### Context

- Backend: Go + chi, runs on port 8080, exposes `GET /v1/health → {"status":"ok"}` and `GET /v1/build-info → {"gitSha":"...","buildTime":"..."}`
- Frontend: SolidJS + Vite, runs on port 5173, currently has no API calls or proxy config
- Root Makefile has `make dev` (backend via air) but no frontend target
- README documents both processes separately; no single "start everything" command

### Approach

Use the **Vite dev-server proxy** to forward `/v1/*` requests from the browser (port 5173) to the backend (port 8080). This eliminates CORS entirely in development and means the frontend uses simple relative URLs. No backend changes required.

Use **`openapi-typescript`** (dev dependency) to generate TypeScript types from `docs/schema/v1/conjugate.yaml`, and **`openapi-fetch`** (runtime dependency) to make type-safe requests against those types. Generated output is committed and kept in sync via a `make generate` step, mirroring the backend's existing codegen discipline.

Add a `<DevFooter>` component rendered only when `import.meta.env.DEV` is true. It uses the generated client to fetch `/v1/health` and `/v1/build-info` and displays git SHA, build time, and API status as a single horizontal bar. This keeps the production UI completely unaffected.

Mount `<DevFooter>` in `App.tsx` so it appears on every page during development.

Add `make dev-all` and `make frontend-generate` root targets.

### Files to Change

#### 1. `frontend/package.json`
Add dependencies:
- `openapi-fetch` (runtime) — type-safe fetch client
- `openapi-typescript` (devDependency) — generates types from schema
Add script: `"generate": "openapi-typescript ../docs/schema/v1/conjugate.yaml -o src/api/v1.d.ts"`

#### 2. `frontend/src/api/v1.d.ts` (generated, committed)
TypeScript types generated from `docs/schema/v1/conjugate.yaml` via `openapi-typescript`. Never edited by hand.

#### 3. `frontend/src/api/client.ts` (new file)
Initialises and exports a single `openapi-fetch` client instance typed against `v1.d.ts`:
```ts
import createClient from "openapi-fetch";
import type { paths } from "./v1.d.ts";
export const api = createClient<paths>({ baseUrl: "" });
```

#### 4. `frontend/vite.config.ts`
Add a `server.proxy` block forwarding `/v1` to `http://localhost:8080`.

#### 5. `frontend/src/components/DevFooter.tsx` (new file)
- Guarded by `import.meta.env.DEV` — returns `null` in production
- Uses `api` client to fetch `/v1/build-info` and `/v1/health` via `createResource`
- Renders a single horizontal bar: `status: ok  backend ver: abc1234  backend build time: 2026-03-10T12:00:00Z`
- Styled to be visually distinct from app content (muted, monospace, small text)

#### 6. `frontend/src/App.tsx`
Mount `<Show when={import.meta.env.DEV}><DevFooter /></Show>` at the bottom of the app shell.

#### 7. `Makefile` (root)
- Add `frontend-generate` target: runs `npm run generate` in `frontend/`
- Update `generate` target to also run `frontend-generate`
- Add `frontend-dev` and `dev-all` targets (`make -j2 dev frontend-dev`)

#### 8. `README.md`
- Add a **Quick start** section: `make dev-all`
- Keep individual commands for reference
- Add a **Verification** section: `curl http://localhost:8080/v1/health` + browser dev footer check

### Acceptance Criteria Mapping

| AC | How satisfied |
|----|--------------|
| #1 Contributor can start both parts and view front page | `make dev-all` + README quick-start section |
| #2 Front page shows live backend signal | Dev footer shows API status, git SHA, build time |
| #3 Repeatable verification step | `curl` command + browser dev footer check in README |

### Out of Scope
- CORS middleware (not needed with Vite proxy)
- Production reverse-proxy config
- Any new backend endpoints
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Frontend and backend wired into a single local MVP workflow via PR #15. Completed before the task finalization convention was established in PR #24.
<!-- SECTION:FINAL_SUMMARY:END -->
