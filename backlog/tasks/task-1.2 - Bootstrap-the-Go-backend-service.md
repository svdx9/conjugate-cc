---
id: TASK-1.2
title: Bootstrap the Go backend service
status: Done
assignee:
  - Codex
created_date: '2026-03-02 18:07'
updated_date: '2026-03-02 18:43'
labels:
  - mvp
  - backend
dependencies:
  - TASK-1.1
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-1
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the first runnable backend service for the conjugation drill application under the repository's backend area. The service should be suitable for local development and provide the minimal API surface needed for an MVP frontend to confirm the backend is alive.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The backend can be started locally and exposes a minimal HTTP surface for health or status checking.
- [x] #2 The backend exposes an endpoint that returns the service git SHA and build time.
- [x] #3 The git SHA and build time values are injected into the application entrypoint at build time rather than hard-coded.
- [x] #4 Basic backend tests or equivalent automated verification cover the initial server behavior that the MVP relies on, including the build metadata endpoint.
- [x] #5 Documentation is updated with the command(s) needed to run and verify the backend locally.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Initialize the backend Go module under `backend/` and create the first application entrypoint at `backend/cmd/main.go`, wiring the root `Makefile` expectations (`./cmd`, `main.serviceGitSHA`, and `main.serviceBuildTime`) into a runnable binary.
2. Add the minimal backend composition structure required by the Go backend skill: `backend/internal/config/` for typed startup config, `backend/internal/http/` for router/server wiring, and a small feature package for status/health behavior so the entrypoint stays focused on dependency assembly.
3. Define the initial HTTP contract schema-first under `docs/schema/v1/` with an OpenAPI document and deterministic `config.yaml`, covering a health/status endpoint and a build metadata endpoint that returns git SHA and build time in a stable response shape.
4. Generate and commit the initial `oapi-codegen` output into `backend/internal/api/v1/api.gen.go`, then implement the generated server interface through thin handlers that delegate to the status feature rather than embedding logic in routing code.
5. Keep configuration minimal and local-dev friendly by parsing server settings in `backend/internal/config` and providing sane defaults for the MVP startup path while preserving explicit validation and deterministic startup/shutdown behavior.
6. Update the root developer workflow as needed for this backend bootstrap, likely by extending the existing root `Makefile` with any additional backend support targets required by the OpenAPI/codegen flow while preserving the repo-local tooling model introduced in `TASK-1.1`.
7. Add backend tests focused on the MVP contract: verify the router or handlers serve the health/status response, verify the metadata endpoint returns the build values injected at build time, and verify the backend can be exercised through `go test ./...` from the existing root `make test` workflow.
8. Update `README.md` so contributors can install backend tools, run the backend locally, hit the initial verification endpoints, and understand how the generated API contract fits into the backend structure.
9. Verify the task by running the relevant backend tests and at least one local startup/build path that proves the root `Makefile` now works end to end for the bootstrapped backend service.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Loaded the repo-pinned Go backend skill during planning because `TASK-1.2` will modify files under `backend/` and needs to follow the enforced backend architecture and schema-first API rules.

Research summary: the current repository contains only the shared scaffold from `TASK-1.1` plus a root `Makefile` that expects a Go app at `backend/cmd/api` and injects `main.serviceGitSHA` / `main.serviceBuildTime` via `-ldflags`. No backend module, config, API schema, or handlers exist yet.

Plan recorded as a proposed implementation plan pending user approval. No code changes made for `TASK-1.2`.

User approved the recorded implementation plan on 2026-03-02. Proceeding with backend bootstrap implementation under the repo-pinned Go backend skill.

Implemented the first backend module under `backend/`: added `backend/go.mod`, `backend/cmd/api/main.go`, typed config loading, chi-based HTTP server wiring, a small status feature, generated OpenAPI server code, and backend tests for the MVP health and build-info endpoints.

Added the schema-first backend contract under `docs/schema/v1/` with deterministic `config.yaml` and `conjugate.yaml`, then generated and committed `backend/internal/api/v1/api.gen.go` via repo-local `oapi-codegen`.

Updated the root `Makefile` to support backend code generation and startup with repo-local tools via `generate`, `backend-run`, and `backend-dev`, while preserving build metadata injection through `main.serviceGitSHA` and `main.serviceBuildTime`.

Verification: `cd backend && GOTOOLCHAIN=local GOPROXY=file:///Users/steve/go/pkg/mod/cache/download GOSUMDB=off GOCACHE=$(pwd)/../.cache/go-build go test ./...` passed.

Verification: `GOTOOLCHAIN=local GOPROXY=file:///Users/steve/go/pkg/mod/cache/download GOSUMDB=off make test` passed through the root workflow.

Verification: `GOTOOLCHAIN=local GOPROXY=file:///Users/steve/go/pkg/mod/cache/download GOSUMDB=off make build` succeeded and rebuilt `backend/internal/api/v1/api.gen.go` before producing the backend binary.

Verification: launching `./backend/bin/conjugate-api` reaches the server startup path and attempts to listen on `:8080`, but this sandbox forbids binding local ports (`bind: operation not permitted`), so endpoint serving was verified through router tests rather than an actual live socket in this environment.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Bootstrapped the first runnable Go backend service for the MVP with a schema-first OpenAPI contract, generated chi server boilerplate, typed startup config, structured request logging, graceful HTTP server lifecycle, and status/build-info endpoints. The root Makefile and README now document generation, test, build, and local backend run workflows, and automated verification covers the health and build metadata endpoints plus the root `make test` and `make build` paths.
<!-- SECTION:FINAL_SUMMARY:END -->
