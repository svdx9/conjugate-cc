---
id: TASK-008.01
title: Define API Contract Boundaries
status: Done
assignee: []
created_date: '2026-04-09 15:50'
updated_date: '2026-04-09 16:13'
labels: []
dependencies: []
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update the backend and frontend schemas to define the API contract boundaries for user sign-in and magic link authentication.
<!-- SECTION:DESCRIPTION:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Found the API schema file at docs/schema/v1/api.yaml. Need to add authentication endpoints and models for user sign-in and magic link functionality.

## Detailed Implementation Plan

### Objective
Update the OpenAPI schema with magic link authentication endpoints and regenerate both backend (Go) and frontend (TypeScript) code.

### Specific Implementation Steps

#### 1. Update OpenAPI Schema (docs/schema/v1/api.yaml)

**Add Authentication Endpoints:**

1. **POST /v1/auth/magiclink/request**
   - OperationId: RequestMagicLink
   - Request body: `{ email: string }` (required, format: email)
   - Responses: 202 (Accepted), 422 (Validation Error)

2. **GET /v1/auth/magiclink/verify**
   - OperationId: GetMagicLinkVerify  
   - Query parameter: `token` (string, required)
   - Responses: 200 (HTML), 400 (Missing token), 410 (Token invalid)

3. **POST /v1/auth/magiclink/verify**
   - OperationId: PostMagicLinkVerify
   - Request body: `application/x-www-form-urlencoded` with `token` field
   - Responses: 302 (Redirect with session), 401 (Invalid token)

4. **DELETE /v1/auth/session**
   - OperationId: DeleteSession
   - Security: Session cookie authentication
   - Responses: 204 (Success), 401 (Unauthorized)

#### 2. Add Required Models to components/schemas
- MagicLinkRequest: `{ email: string }`
- ErrorResponse: Standard error format with code, message, details
- ValidationError: Field-specific validation errors

#### 3. Add Security Scheme
- sessionCookie: Cookie-based authentication scheme
- Apply to DELETE /v1/auth/session endpoint

#### 4. Code Regeneration
- Backend: Run `cd backend && make generate`
- Frontend: Run `cd frontend && npm run generate`

### Estimated Time: 1-2 hours

### Dependencies: None

### Risk Factors:
- Schema syntax errors (mitigated by validation)
- Generation tool compatibility (mitigated by testing)
- Type conflicts in generated code (mitigated by review)

### Verification Steps:
1. Schema validates without errors
2. Backend code compiles successfully
3. Frontend TypeScript types compile without errors
4. No breaking changes to existing endpoints
5. Generated code follows existing patterns
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
