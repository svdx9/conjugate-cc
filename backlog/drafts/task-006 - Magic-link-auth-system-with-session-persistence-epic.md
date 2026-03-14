---
id: TASK-006
title: Magic link auth system with session persistence (epic)
status: To Do
assignee: []
created_date: '2026-03-12 00:51'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - authentication
dependencies: []
ordinal: 11000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement a magic link authentication system with PostgreSQL-backed session persistence.

Users request a magic link via email; clicking the link verifies the token and creates a durable session stored in the database. Subsequent requests authenticate via a session cookie resolved against the sessions table.

Endpoints:
- POST /api/v1/auth/magiclink/request — accept an email address, generate a single-use token, send a magic link email
- GET  /api/v1/auth/magiclink/verify  — consume the token, create a session, set a session cookie
- DELETE /api/v1/auth/session          — logout; invalidate the session

Parent epic; see sub-tasks TASK-006.1 through TASK-006.8 for implementation breakdown.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All sub-tasks are Done
- [ ] #2 Magic link flow works end-to-end in integration test or manual smoke test
- [ ] #3 Sessions persist in PostgreSQL and survive a server restart
- [ ] #4 Expired or used tokens are rejected with 401
- [ ] #5 Logout invalidates the session immediately
<!-- AC:END -->
