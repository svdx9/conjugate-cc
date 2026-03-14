---
id: TASK-005
title: Add frontend CI gates (test + lint) and fix make test to cover frontend
status: To Do
assignee: []
created_date: '2026-03-11 18:03'
updated_date: '2026-03-14 16:26'
labels:
  - ci
  - frontend
  - dx
dependencies: []
references:
  - frontend/package.json
  - frontend/src/shared/api/client.ts
  - .github/workflows/lint.yml
  - .github/workflows/test.yml
  - Makefile
priority: high
ordinal: 10000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
PR #15 introduced `openapi-fetch` as a frontend dependency and new components that import it. No CI workflow covers frontend files, and `make test` only runs Go tests. As a result, a broken import shipped to main undetected — `npm test` fails immediately with "Failed to resolve import 'openapi-fetch'".

Two gaps must be closed:

1. **No frontend CI** — `lint.yml` and `test.yml` watch only `backend/**/*.go` path patterns. Any change under `frontend/**` merges with zero automated checks.
2. **`make test` misses frontend** — the CLAUDE.md pre-commit rule says "run `make test` before committing", but that only executes `go test ./...` in `backend/`. A developer following the rule would not catch a broken frontend build or failing vitest suite.

Fix both gaps so that a PR touching `frontend/**` must pass frontend checks before it can merge, and so the local `make test` catches frontend regressions too.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 A `frontend-tests` job is added to the existing `.github/workflows/test.yml` workflow (not a new file) with path filters so it only triggers when `frontend/**` files change
- [ ] #2 The job runs `npm ci` to install dependencies from the lockfile (not `npm install`)
- [ ] #3 The job runs `npm test` (vitest) and fails the PR if any test fails
- [ ] #4 A `typecheck` script (`tsc --noEmit`) is added to `frontend/package.json` and the job runs `npm run typecheck` to catch TypeScript errors, failing the PR if any type error exists
- [ ] #5 The job runs `npm run build` to catch Vite build-time errors (e.g. unresolvable imports) and fails the PR if the build fails
- [ ] #6 The root `make test` target runs both Go tests and frontend tests so the local pre-commit gate is complete
- [ ] #7 All three checks (test, typecheck, build) would have caught the openapi-fetch bug before merge — verified by confirming a checkout without `npm install` would fail the build step
<!-- AC:END -->
