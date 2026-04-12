---
id: TASK-008.07
title: 'Database infrastructure setup (config, migrations, sqlc, pool)'
status: To Do
assignee: []
created_date: '2026-04-11 17:56'
labels:
  - backend
  - database
  - authentication
dependencies:
  - TASK-008.01
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Set up the database infrastructure required for magic link authentication.

- Extend `internal/config/config.go` with `DATABASE_URL` and `SESSION_COOKIE_SECURE`
- Create migration files in `sql/migrations/` for users, magic_links, and sessions tables
- Set up `sqlc.yaml` and write queries in `sql/queries/`
- Generate sqlc code into `internal/db/queries/`
- Create `internal/db/` package with pool wiring and transaction helpers
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Config includes DATABASE_URL and SESSION_COOKIE_SECURE with validation
- [ ] #2 Migration files exist for users, magic_links, and sessions tables
- [ ] #3 Migrations apply and roll back cleanly via migrate CLI
- [ ] #4 sqlc generates query code without errors
- [ ] #5 internal/db/ package provides pool setup and transaction helpers
- [ ] #6 All code compiles and existing tests pass
<!-- AC:END -->
