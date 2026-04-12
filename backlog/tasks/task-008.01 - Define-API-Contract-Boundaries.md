---
id: TASK-008.01
title: Define API Contract Boundaries
status: Completed
assignee: []
created_date: '2026-04-09 15:50'
updated_date: '2026-04-12 16:52'
labels: []
dependencies: []
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update the backend and frontend schemas to define the API contract boundaries for user sign-in and magic link authentication.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 All 4 auth endpoints present in OpenAPI spec with correct methods, paths, and response shapes
- [x] #2 GET and POST /magiclink/verify are distinct operations in the spec
- [x] #3 All endpoints return application/json only (no text/html)
- [x] #4 POST /verify returns 200 with JSON body (not 302 redirect)
- [x] #5 Error response models (ErrorResponse, ValidationError) defined and reused
- [x] #6 sessionCookie security scheme defined and applied to DELETE /session
- [x] #7 make generate succeeds and produces updated api.gen.go
- [x] #8 Frontend npm run generate succeeds and produces updated v1.d.ts
- [x] #9 Generated code is committed
- [x] #10 Backend compiles without errors
- [x] #11 Frontend typechecks without errors
<!-- AC:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
See doc-001 (Magic Link Authentication Flow) for the full auth flow design and security requirements.

### Design Decisions

1. **Backend returns JSON only** — All auth endpoints return `application/json`. No `text/html` responses. The frontend is responsible for rendering all UI.
2. **GET /verify does not consume the token** — Prevents automated email link scanners from burning tokens. Only the POST consumes the token.
3. **POST /verify returns 200 JSON** (not 302 redirect) — Consistent with the JSON-only principle. The frontend handles navigation after receiving the response.
4. **Security reviewed** — The auth flow was reviewed against the go-general-web-backend-security reference. See doc-001 for the full security review summary.

### Endpoints Defined

1. **POST /v1/auth/magiclink/request** — Request magic link (202 / 422)
2. **GET /v1/auth/magiclink/verify** — Validate token, return email (200 / 400 / 410)
3. **POST /v1/auth/magiclink/verify** — Consume token, create session (200 + Set-Cookie / 401)
4. **DELETE /v1/auth/session** — Logout (204 / 401)

### Models Added

- MagicLinkRequest, MagicLinkVerifyResponse, MagicLinkConfirmRequest, MagicLinkConfirmResponse
- ErrorResponse, ValidationError
- sessionCookie security scheme
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Successfully implemented task 008.01 by:

1. **Updated OpenAPI Schema** (`docs/schema/v1/api.yaml`):
   - Added 4 authentication endpoints: POST /v1/auth/magiclink/request, GET/POST /v1/auth/magiclink/verify, DELETE /v1/auth/session
   - Added 3 new models: MagicLinkRequest, ErrorResponse, ValidationError
   - Added sessionCookie security scheme for protected endpoints

2. **Regenerated Backend Code**:
   - Ran `make generate` in backend directory
   - Updated `internal/api/v1/api.gen.go` with new handler interfaces and models
   - Fixed compilation issues by creating composite handler in `internal/http/server.go`
   - All backend code compiles successfully

3. **Regenerated Frontend Code**:
   - Ran `npm run generate` in frontend directory
   - Updated `src/api/v1.d.ts` with TypeScript types for all new endpoints and models
   - All TypeScript types compile successfully with `npm run typecheck`

4. **Verification**:
   - OpenAPI schema is valid and well-formed
   - Backend Go code compiles without errors
   - Frontend TypeScript types compile without errors
   - No breaking changes to existing endpoints
   - Generated code follows existing patterns and conventions

The API surface is now fully defined and ready for implementation. All authentication endpoints are stubbed and will return 501 Not Implemented until the backend handlers are implemented in subsequent tasks.
<!-- SECTION:FINAL_SUMMARY:END -->
