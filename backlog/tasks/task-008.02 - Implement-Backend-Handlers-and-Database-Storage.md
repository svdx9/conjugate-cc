---
id: TASK-008.02
title: Implement Backend Handlers and Database Storage
status: To Do
assignee: []
created_date: '2026-04-09 15:50'
updated_date: '2026-04-09 16:05'
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
## API Surface Definition Plan

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
