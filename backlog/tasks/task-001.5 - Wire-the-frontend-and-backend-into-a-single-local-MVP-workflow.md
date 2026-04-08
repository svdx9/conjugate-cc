---
id: TASK-001.5
title: Wire the frontend and backend into a single local MVP workflow
status: Done
assignee: []
created_date: '2026-03-02 18:07'
updated_date: '2026-03-29 00:00'
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
- [x] #1 A contributor can follow repository documentation to start the backend and frontend together and view the MVP front page locally.
- [x] #2 The MVP front page demonstrates live integration with the backend through a minimal user-visible signal such as status, readiness, or configuration data.
- [x] #3 The integration path has a repeatable verification step that confirms the local stack is working end to end.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### Context

The frontend (SolidJS + Vite, port 3000) and backend (Go + chi, port 8080) run independently with no integration. A contributor must start each separately and has no way to confirm the stack works end-to-end. This task connects them so the front page shows a live backend signal and a single command starts everything.

### Corrected Facts (vs original backlog plan)

| Original plan says | Actual |
|--------------------|--------|
| Schema at docs/schema/v1/conjugate.yaml | docs/schema/v1/api.yaml |
| Endpoints /v1/health, /v1/build-info | /v1/status, /v1/metadata |
| Root Makefile exists | Only backend/Makefile exists |
| Footer at components/Footer.tsx | src/app/Footer.tsx |
| Dev server port 5173 | Port 3000 |
| Create DevFooter component | Modify existing Footer component |

### Approach

Per the user's context: add backend status to the existing Footer (not a separate DevFooter). Use a sync poller with createSignal + setInterval for health checks, and openapi-fetch + openapi-typescript for type-safe API calls.

### Steps

#### 1. Add Vite dev-server proxy

File: `frontend/vite.config.ts`

Add server.proxy block to forward /api requests to http://localhost:8080:
```ts
server: {
  port: 3000,
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true,
    },
  },
},
```

#### 2. Install frontend dependencies

```bash
cd frontend
npm install openapi-fetch
npm install -D openapi-typescript
```

Add generate script to frontend/package.json:
```json
"generate": "openapi-typescript ../docs/schema/v1/api.yaml -o src/api/v1.d.ts"
```

#### 3. Generate TypeScript types

File (generated, committed): `frontend/src/api/v1.d.ts`

Run `npm run generate` to produce types from the OpenAPI schema.

#### 4. Create API client

File (new): `frontend/src/api/client.ts`
```ts
import createClient from "openapi-fetch";
import type { paths } from "./v1";
export const api = createClient<paths>({ baseUrl: "" });
```

#### 5. Update Footer with backend status indicator

File: `frontend/src/app/Footer.tsx`

Add a health poller using the pattern from user context:
- createSignal<boolean>(false) for backendAvailable
- setInterval every 30s calling /v1/status via the typed client
- Green/red dot indicator showing backend status
- When backend is available, also fetch /v1/metadata to show git SHA and build time
- Guard the status indicator behind import.meta.env.DEV so production footer stays clean

Styling follows the design system: small monospace text, muted, using existing tokens (text-highlight for connected, standard text for disconnected).

#### 6. Create root Makefile

File (new): `Makefile` (project root)

Targets:
- `frontend-generate`: runs `npm run generate --prefix frontend`
- `generate`: runs backend generate + frontend-generate
- `frontend-dev`: runs `npm run dev --prefix frontend`
- `dev`: runs `make -C backend build && make -C backend air/dev` (backend)
- `dev-all`: runs backend dev + frontend dev in parallel (`make -j2 dev frontend-dev`)

#### 7. Update README.md

- Add Quick Start section with `make dev-all`
- Add Verification section: `curl http://localhost:8080/api/v1/status` + check footer indicator in browser
- Keep existing individual commands for reference

#### 8. Update backlog task file

Set status to "In Progress" in this task file.

### Files to Modify/Create

| File | Action |
|------|--------|
| frontend/vite.config.ts | Edit — add proxy |
| frontend/package.json | Edit — add deps + generate script |
| frontend/src/api/v1.d.ts | Generate (committed) |
| frontend/src/api/client.ts | Create |
| frontend/src/app/Footer.tsx | Edit — add status indicator |
| Makefile | Create (root) |
| README.md | Edit |
| backlog/tasks/task-001.5...md | Edit — status to In Progress |

### Acceptance Criteria Mapping

| AC | How satisfied |
|----|--------------|
| #1 Contributor can start both and view front page | `make dev-all` + README quick-start |
| #2 Front page shows live backend signal | Footer status indicator (green/red + SHA + build time) |
| #3 Repeatable verification step | curl + browser footer check documented in README |

### Verification

1. `cd frontend && npm install && npm run generate` — types file created without errors
2. `make dev-all` from root — both servers start
3. Open http://localhost:3000 — landing page loads, footer shows green indicator with git SHA
4. Stop backend — footer indicator turns red within 30s
5. `curl http://localhost:8080/api/v1/status` returns `{"status":"ok"}`
6. `cd backend && make test && make lint` — no errors (no backend changes expected)
7. `cd frontend && npm run lint && npm test` — no errors
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Frontend and backend are wired together for local development. The Vite dev server proxies `/v1` to the Go backend on port 8080. A typed API client (`openapi-fetch` + generated types from the OpenAPI schema) hits `/v1/status` and `/v1/metadata` every 30 seconds. The result is surfaced in the footer as a green/red dot with git SHA and build time, visible only in DEV mode. A root `Makefile` adds `make dev-all` to start both servers in parallel. The README quick-start and verification steps document the full workflow. The footer layout was also updated to match the navbar width and structure (indicator on left, links on right).
<!-- SECTION:FINAL_SUMMARY:END -->
