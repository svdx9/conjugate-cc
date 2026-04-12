---
id: TASK-008.02
title: Implement Backend Handlers and Database Storage
status: Done
assignee: []
created_date: '2026-04-09 15:50'
updated_date: '2026-04-12 05:28'
labels: []
dependencies: []
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Build the backend handlers and database storage for user authentication and state management.
<!-- SECTION:DESCRIPTION:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Backend Implementation Plan

See doc-001 (Magic Link Authentication Flow) for the full auth flow design and security requirements.

### Objective
Implement the backend handlers and database storage for magic link authentication, building on the API contract defined in task 008.01.

### Implementation Steps

#### 1. Database Infrastructure Setup
- Extend `internal/config/config.go` with `DATABASE_URL` and `SESSION_COOKIE_SECURE`
- Create migration files in `sql/migrations/`:
  - `000001_create_users_table` (id UUID, email, created_at, updated_at)
  - `000002_create_magic_links_table` (id, user_id, token_hash SHA-256, expires_at, consumed_at)
  - `000003_create_sessions_table` (id, user_id, token_hash, expires_at, created_at)
- Set up `sqlc.yaml` and write queries in `sql/queries/`
- Generate sqlc code into `internal/db/queries/`

#### 2. Auth Feature Implementation
- `internal/auth/` package:
  - Repository interfaces (user, magic link, session)
  - Service layer: token generation (crypto/rand, 32 bytes), SHA-256 hashing, constant-time comparison
  - HTTP handlers implementing generated ServerInterface methods
- `internal/db/` implements auth repository interfaces

#### 3. Email (stub only)
- `internal/email/` package with sender interface
- Stub sender that logs to stdout (production SES delivery is TASK-008.06)

#### 4. HTTP Server Integration
- Replace CompositeHandler stubs with real auth handlers
- Add session auth middleware
- Add custom header check (X-Requested-With) for CSRF on state-changing endpoints

#### 5. Testing
- Unit tests for auth service (token generation, hashing, expiry)
- HTTP handler tests with httptest
- Repository tests against test database

### Security Requirements (from doc-001)
- Tokens: crypto/rand, SHA-256 storage, constant-time compare, 15min TTL
- Sessions: HttpOnly, SameSite=Lax, conditional Secure
- CSRF: SameSite=Lax + custom header
- GET /verify does NOT consume token (link-scanner safe)
- Atomic token consumption via DB transaction

### Dependencies
- Task 008.01 (API Contract) — completed
- PostgreSQL database for development
- Email adapter (can be stubbed initially)
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
SUPERSEDED — not implemented. Scope was too large and has been broken into three focused tasks with explicit dependency ordering:

- TASK-008.07: Database infrastructure setup (config, migrations, sqlc, pool) — depends on 008.01
- TASK-008.08: Auth service layer and email stub — depends on 008.07
- TASK-008.09: Auth HTTP handlers and server integration — depends on 008.08

See doc-001 (Magic Link Authentication Flow) for the full auth flow design and security requirements.
<!-- SECTION:FINAL_SUMMARY:END -->
