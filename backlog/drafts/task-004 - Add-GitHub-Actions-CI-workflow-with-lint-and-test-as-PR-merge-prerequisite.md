---
id: TASK-004
title: Add GitHub Actions CI workflow with lint and test as PR merge prerequisite
status: To Do
assignee: []
created_date: '2026-03-04 14:15'
updated_date: '2026-03-14 16:26'
labels:
  - ci
  - dx
  - backend
dependencies: []
priority: high
ordinal: 24000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Context

A regression was merged into the `claude/migrate-from-codex` PR because there was no automated CI gate enforcing `make test` (or `make build-dev`) before merge. The broken test (`handler_test.go` calling `NewHandler` with the wrong number of arguments) existed in the codebase and was never caught because nothing prevented the PR from being merged without tests passing.

## Goal

Add a GitHub Actions workflow that runs lint and tests on every PR, and configure the repository to require this check to pass before a PR can be merged (branch protection rule).

## Implementation Notes

- Create `.github/workflows/ci.yml` that runs on `pull_request` events targeting `main`
- The workflow should run `make build-dev` (which runs `generate → format → tidy → lint → test → build`) in the `backend` directory
- Alternatively, run `make lint` and `make test` as separate steps for clearer failure output
- Set up the Go version to match what the project uses (check `backend/go.mod` for the Go version)
- Cache Go modules and the build cache (`.cache/go-build`) for faster runs
- After the workflow exists, enable a branch protection rule on `main` requiring the CI check to pass before merging

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 `.github/workflows/ci.yml` exists and triggers on pull requests to `main`
- [ ] #2 The workflow runs `make lint` and `make test` (or equivalent) against the backend
- [ ] #3 A failing test or lint error causes the workflow to fail and blocks merge
- [ ] #4 The Go build cache is cached between runs to avoid unnecessary reinstallation of tools
- [ ] #5 Branch protection on `main` requires the CI check to pass before a PR can be merged
<!-- SECTION:DESCRIPTION:END -->

<!-- AC:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
CI workflow added via PR #11. GitHub Actions workflow for lint and test on PRs was implemented and merged before the task finalization convention was established in PR #24.
<!-- SECTION:FINAL_SUMMARY:END -->
