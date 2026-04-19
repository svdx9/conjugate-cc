---
id: TASK-008.09
title: Auth HTTP handlers and server integration
status: To Do
assignee: []
created_date: '2026-04-11 17:57'
labels:
  - backend
  - authentication
dependencies:
  - TASK-008.08
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Wire the auth service into HTTP handlers and integrate with the server.

- Implement auth HTTP handlers for all 4 endpoints (request, GET verify, POST verify, DELETE session)
- Replace CompositeHandler stubs with real auth handlers
- Add session auth middleware for protected routes
- Add custom header check (X-Requested-With) for CSRF on state-changing endpoints
- Set session cookie with correct attributes (HttpOnly, SameSite=Lax, conditional Secure)
- HTTP handler tests with httptest
- Wire everything together in cmd/server/main.go
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All 4 auth endpoints return correct responses per OpenAPI spec
- [ ] #2 Session cookie set with HttpOnly, SameSite=Lax, conditional Secure
- [ ] #3 Session auth middleware validates session cookie on protected routes
- [ ] #4 CSRF custom header check enforced on state-changing endpoints
- [ ] #5 GET /verify does NOT consume the token
- [ ] #6 POST /verify atomically consumes token and creates session
- [ ] #7 CompositeHandler stubs are fully replaced
- [ ] #8 HTTP handler tests cover success and error cases
- [ ] #9 All code compiles and tests pass
<!-- AC:END -->
