---
id: doc-001
title: Magic Link Authentication Flow
type: other
created_date: '2026-04-11 17:42'
---
# Magic Link Authentication Flow

## Design Principles

1. **Backend returns JSON only** — All auth API endpoints return `application/json`. No HTML rendering on the server.
2. **Frontend owns the UI** — The SolidJS frontend consumes JSON responses and renders all user-facing pages (forms, confirmation screens, error pages).
3. **Stateless API, stateful sessions** — Authentication state is managed via server-side sessions stored in PostgreSQL, referenced by a session cookie.

---

## Authentication Flow

### Step 1: Request Magic Link

```
User enters email
  → Frontend: POST /api/v1/auth/magiclink/request { email }
  → Backend: validates email, creates/finds user, generates token,
             stores SHA-256 hash in DB, sends email with plaintext token in link
  → Backend: returns 202 Accepted (always, to prevent email enumeration)
  → Frontend: shows "Check your email" message
```

**Responsibilities:**

- **Frontend:** Collects email, shows "check your email" confirmation regardless of response
- **Backend:** Validates email format, creates/finds user, generates cryptographic token, stores hashed token in DB, sends email via email adapter

### Step 2: Verify Token (GET — read-only validation)

```
User clicks link in email
  → Browser navigates to frontend route with ?token=...
  → Frontend: GET /api/v1/auth/magiclink/verify?token=...
  → Backend: validates token hash exists, not expired, not consumed
             Does NOT consume the token.
  → Backend: returns 200 { email } or error (400/410)
  → Frontend: renders "Sign in as user@example.com?" confirmation page,
              or renders error page on failure
```

**Error cases:**

- 400: Missing token parameter
- 410: Token not found, already used, or expired

**Responsibilities:**

- **Frontend:** Extracts token from URL, calls GET verify, renders confirmation page with user email, or renders error page on failure
- **Backend:** Validates token without consuming it, returns associated email address

**Security note:** GET does not consume the token. This prevents automated email link scanners and security crawlers from burning tokens before the real user clicks. Token consumption only happens via an explicit POST action (Step 3).

### Step 3: Confirm Sign-In (POST — consumes token, creates session)

```
User clicks "Confirm"
  → Frontend: POST /api/v1/auth/magiclink/verify { token }
  → Backend: atomically consumes token + creates session in single DB transaction
  → Backend: returns 200 { user_id, email } with Set-Cookie session header
  → Frontend: updates auth state, navigates to home
```

**Error cases:**

- 401: Token invalid, expired, or already consumed

**Responsibilities:**

- **Frontend:** Sends POST with token as JSON, handles response, updates auth state in UI, navigates to home
- **Backend:** Atomically consumes token and creates user session in DB, sets session cookie on response

### Step 4: Logout

```
User clicks "Logout"
  → Frontend: DELETE /api/v1/auth/session (with X-Requested-With header)
  → Backend: invalidates session in DB, clears cookie
  → Backend: returns 204 No Content
  → Frontend: clears auth state, redirects to home
```

**Error cases:**

- 401: Not authenticated (no valid session cookie)

**Responsibilities:**

- **Frontend:** Calls DELETE with custom header, clears local auth state, updates UI
- **Backend:** Invalidates session record in DB, clears session cookie

---

## Token Security Requirements

- **Generation:** `crypto/rand` with minimum 32 bytes, URL-safe base64 encoded
- **Storage:** Only the SHA-256 hash of the token is stored in the database; the plaintext token appears only in the email link
- **Verification:** Token hashes are compared using `crypto/subtle.ConstantTimeCompare` to prevent timing attacks
- **Single-use:** Token consumption and session creation happen atomically in a single DB transaction to prevent race conditions
- **TTL:** 15 minutes from creation, enforced server-side
- **GET safety:** GET /verify does NOT consume the token — prevents link-scanner token burning

## Session Security Requirements

- **Session ID:** Generated with `crypto/rand` (minimum 32 bytes)
- **Cookie attributes:**
  - `HttpOnly=true` — prevents JavaScript access (XSS mitigation)
  - `SameSite=Lax` — blocks cross-origin form submissions (CSRF primary defense)
  - `Path=/`
  - `Secure` — conditional on environment via `SESSION_COOKIE_SECURE` config flag (true in production, false in local dev over HTTP)
- **Server-side expiry:** Enforced on every request by checking session record in DB
- **Persistence:** Sessions stored in PostgreSQL and survive server restarts

## CSRF Protection

- **Primary defense:** `SameSite=Lax` session cookies block cross-origin form POSTs and fetch requests
- **Secondary defense:** State-changing API requests from the frontend require a custom header (`X-Requested-With: XMLHttpRequest`). Cross-origin requests with custom headers trigger a CORS preflight, which is blocked since we don't allow arbitrary origins.
- **No CSRF tokens needed** for this architecture — the combination of `SameSite=Lax` and custom header requirement provides sufficient protection

## Rate Limiting

Rate limiting for the magic link request endpoint is tracked separately in TASK-008.05.

- POST /api/v1/auth/magiclink/request: rate-limit by IP address and by target email address
- Rate-limited requests return 429 Too Many Requests with Retry-After header

---

## Security Review Summary

This design was reviewed against the go-general-web-backend-security reference. Key findings addressed:

| ID | Severity | Finding | Mitigation |
|---|---|---|---|
| SEC-001 | Critical | Token randomness | `crypto/rand`, 32+ bytes |
| SEC-002 | High | Token storage | SHA-256 hash in DB, not plaintext |
| SEC-003 | High | Session cookie security | HttpOnly, SameSite=Lax, conditional Secure |
| SEC-004 | High | CSRF protection | SameSite=Lax + custom header |
| SEC-005 | High | Token comparison | `crypto/subtle.ConstantTimeCompare` |
| SEC-006 | Medium | Token expiry | 15 minute TTL, server-side enforcement |
| SEC-007 | Medium | Rate limiting | TASK-008.05 (per-IP and per-email) |
| SEC-008 | Medium | Request body limits | Enforce via middleware |
| SEC-009 | Low | Atomic token consumption | Single DB transaction |
