---
id: TASK-006.5
title: Implement magic link request handler (POST /api/v1/auth/magiclink/request)
status: To Do
assignee: []
created_date: '2026-03-12 01:11'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - authentication
dependencies:
  - TASK-006.1
  - TASK-006.2
  - TASK-006.3
  - TASK-006.4
parent_task_id: TASK-006
priority: high
ordinal: 16000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the handler for POST /api/v1/auth/magiclink/request.

Logic:
1. Validate request body — email must be a well-formed address.
2. UpsertUser by email (find-or-create); obtain the user_id.
3. Generate a cryptographically random token (32 bytes, base64url encoded).
4. Hash the token with SHA-256; store only the hash in magiclinks.
5. Set expires_at = now() + MAGIC_LINK_TTL (default 5 minutes, configurable via env).
6. Construct the magic link URL: `<BASE_URL>/api/v1/auth/magiclink/verify?token=<raw_token>`.
7. Call EmailSender.SendMagicLink.
8. Always respond 202 Accepted regardless of whether the email was previously known (avoids enumeration).

Config dependencies (add to `internal/config`):
- `BASE_URL` (required) — used to construct the magic link URL.
- `MAGIC_LINK_TTL` (optional, default `5m`) — parsed as a Go duration; controls token lifetime.

Steps 2–5 SHOULD run inside a single DB transaction to keep user upsert and token insert atomic.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Handler returns 202 for a valid email address
- [ ] #2 Handler returns 422 for a malformed email
- [ ] #3 Raw token is never stored; only the SHA-256 hash is persisted
- [ ] #4 Token expires_at matches now() + MAGIC_LINK_TTL (default 5 minutes)
- [ ] #5 MAGIC_LINK_TTL is read from config and defaults to 5m if unset
- [ ] #6 User upsert and token insert are atomic (single transaction)
- [ ] #7 HTTP handler test covers happy path and invalid email cases
<!-- AC:END -->
