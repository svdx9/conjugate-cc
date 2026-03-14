---
id: TASK-006.8
title: Implement logout handler (DELETE /api/v1/auth/session)
status: To Do
assignee: []
created_date: '2026-03-12 01:12'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - authentication
dependencies:
  - TASK-006.2
  - TASK-006.3
  - TASK-006.7
parent_task_id: TASK-006
priority: medium
ordinal: 19000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the handler for DELETE /api/v1/auth/session.

This route MUST be behind the session auth middleware (TASK-006.7).

Logic:
1. Read the authenticated session from context (set by middleware).
2. Call InvalidateSession for the session ID.
3. Clear the session cookie (Set-Cookie with Max-Age=0).
4. Respond 204 No Content.

No request body required.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Authenticated request invalidates the session in the DB and responds 204
- [ ] #2 Unauthenticated request returns 401 (rejected by middleware before reaching handler)
- [ ] #3 Session cookie is cleared in the response (Max-Age=0)
- [ ] #4 Subsequent requests with the same cookie return 401 (session is invalidated)
- [ ] #5 Handler test covers happy path
<!-- AC:END -->
