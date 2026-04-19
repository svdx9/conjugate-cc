---
id: TASK-008.05
title: Implement rate limiting for authentication endpoints
status: To Do
assignee: []
created_date: '2026-04-11 17:37'
labels:
  - backend
  - security
  - authentication
dependencies: []
parent_task_id: TASK-008
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add rate limiting to the magic link request endpoint (POST /api/v1/auth/magiclink/request) to prevent email flooding abuse.

Rate limiting should be applied:
- Per IP address: limit the number of magic link requests from a single IP
- Per target email: limit the number of magic link requests for a single email address

This is a security requirement identified during the auth flow design review (SEC-007).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 POST /api/v1/auth/magiclink/request is rate-limited per IP address
- [ ] #2 POST /api/v1/auth/magiclink/request is rate-limited per target email address
- [ ] #3 Rate-limited requests return 429 Too Many Requests with appropriate Retry-After header
- [ ] #4 Rate limit configuration is externalized (not hardcoded)
- [ ] #5 Rate limiting does not affect other endpoints
<!-- AC:END -->
