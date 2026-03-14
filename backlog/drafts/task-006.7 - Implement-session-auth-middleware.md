---
id: TASK-006.7
title: Implement session auth middleware
status: To Do
assignee: []
created_date: '2026-03-12 01:12'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - authentication
dependencies:
  - TASK-006.2
parent_task_id: TASK-006
priority: high
ordinal: 18000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement a chi middleware that authenticates requests on protected routes via the session cookie.

Logic:
1. Read the session cookie value; return 401 if absent.
2. Call `GetSessionByID`; return 401 if not found.
3. Return 401 if `invalidated_at` is set or `expires_at` is in the past.
4. Store the authenticated session (including `user_id`) in the request context via a typed context key.
5. Call next handler on success.

Expose a helper `SessionFromContext(ctx context.Context) (*Session, bool)` for handlers that need the caller's identity.

Wire the middleware onto all routes that require authentication. The magiclink request and verify endpoints MUST NOT be behind this middleware.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Request with valid session cookie reaches the next handler
- [ ] #2 Request with no cookie returns 401
- [ ] #3 Request with invalidated session returns 401
- [ ] #4 Request with expired session returns 401
- [ ] #5 Authenticated session (including user_id) is retrievable from context via SessionFromContext
- [ ] #6 Unit tests cover all four cases above
<!-- AC:END -->
