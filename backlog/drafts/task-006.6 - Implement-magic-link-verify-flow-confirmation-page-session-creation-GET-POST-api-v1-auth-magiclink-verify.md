---
id: TASK-006.6
title: >-
  Implement magic link verify flow: confirmation page + session creation (GET +
  POST /api/v1/auth/magiclink/verify)
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
parent_task_id: TASK-006
priority: high
ordinal: 17000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the two-step verify flow that prevents token burn-by-prefetch.

**Why two steps**: GET requests on magic links are routinely prefetched by email clients, security scanners, and corporate proxies. Consuming the token in a GET causes systematic sign-in failures in environments with mailbox security crawling. Token consumption MUST only happen as a result of an explicit user-initiated non-GET action.

**Step 1 — GET /api/v1/auth/magiclink/verify?token=\<raw_token\>**

1. Read `token` query param; return 400 if missing.
2. Hash the token (SHA-256); look up in `magiclinks` by token_hash.
3. If not found, `used_at` is set, or `expires_at` is in the past → return 410 Gone.
4. If valid → respond 200 with a minimal HTML page containing a single form:
   - `<form method="POST" action="/api/v1/auth/magiclink/verify">`
   - Hidden field: `token=<raw_token>`
   - Submit button: "Sign in"
   - Do NOT consume or mutate the token at this step.

**Step 2 — POST /api/v1/auth/magiclink/verify**

1. Read `token` from the form body (application/x-www-form-urlencoded).
2. Hash the token (SHA-256); look up in `magiclinks` by token_hash.
3. Reject 401 if: not found, `used_at` is already set (replay attempt), or `expires_at` is in the past.
4. Within a single DB transaction:
   a. `MarkMagicLinkUsed` — set `used_at = now()`; this is the single-use enforcement gate.
   b. `InsertSession` — create session row linked to the token's `user_id` (expires_at = now() + 30 days).
5. Set session cookie: HttpOnly, SameSite=Lax, Secure (controlled by ENV), Path=/, Max-Age=30 days.
6. Redirect 302 to /.

Steps 4a and 4b MUST be inside one transaction. A second POST with the same token hits the `used_at` check in step 3 and is rejected before any write occurs.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 GET with valid unused token returns 200 with a confirmation form; token is NOT consumed
- [ ] #2 GET with missing token returns 400
- [ ] #3 GET with expired token returns 410
- [ ] #4 GET with already-used token returns 410
- [ ] #5 POST with valid unused token creates a session, sets cookie, and redirects 302 to /
- [ ] #6 POST with expired token returns 401
- [ ] #7 POST with already-used token returns 401 (replay prevention)
- [ ] #8 A second POST with the same token is rejected even if submitted concurrently (used_at gate inside transaction)
- [ ] #9 Token mark-used and session insert are in a single transaction
- [ ] #10 Cookie is HttpOnly and SameSite=Lax
- [ ] #11 Handler tests cover all cases above for both GET and POST
<!-- AC:END -->
