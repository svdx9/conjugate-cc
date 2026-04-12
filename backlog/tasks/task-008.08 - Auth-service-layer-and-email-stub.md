---
id: TASK-008.08
title: Auth service layer and email stub
status: To Do
assignee: []
created_date: '2026-04-11 17:57'
labels:
  - backend
  - authentication
dependencies:
  - TASK-008.07
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the auth service layer and a stub email sender for magic link authentication.

- Create `internal/auth/` package with repository interfaces (user, magic link, session)
- Implement auth service: token generation (crypto/rand, 32 bytes), SHA-256 hashing, constant-time comparison, 15min TTL
- `internal/db/` implements auth repository interfaces using sqlc-generated queries
- Create `internal/email/` package with sender interface and stub sender (logs to stdout)
- Unit tests for auth service (token generation, hashing, expiry, atomic consumption)
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Auth repository interfaces defined in internal/auth/
- [ ] #2 Auth service implements token generation with crypto/rand (32 bytes)
- [ ] #3 Tokens stored as SHA-256 hashes and compared with constant-time comparison
- [ ] #4 Token TTL of 15 minutes is enforced
- [ ] #5 internal/db/ implements auth repository interfaces
- [ ] #6 Email sender interface exists with a stub implementation that logs to stdout
- [ ] #7 Unit tests cover token generation, hashing, expiry, and consumption logic
- [ ] #8 All code compiles and tests pass
<!-- AC:END -->
