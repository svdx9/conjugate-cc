---
id: TASK-006.3
title: Update OpenAPI spec with auth endpoints
status: To Do
assignee: []
created_date: '2026-03-12 01:11'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - api
  - authentication
dependencies: []
parent_task_id: TASK-006
priority: high
ordinal: 14000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add the four auth endpoints to docs/schema/v1/*.yaml and regenerate the Go server boilerplate via oapi-codegen.

Endpoints to define:

- POST /api/v1/auth/magiclink/request
  - Request body: `{ email: string }`
  - Response 202: accepted (avoids email enumeration)
  - Response 422: validation error

- GET /api/v1/auth/magiclink/verify
  - Query param: `token` (string, required)
  - Response 200: HTML confirmation page (token valid, not yet consumed)
  - Response 400: missing token param
  - Response 410: token not found, already used, or expired

- POST /api/v1/auth/magiclink/verify
  - Request body: `application/x-www-form-urlencoded`, field `token`
  - Response 302: redirect to / with Set-Cookie session header
  - Response 401: invalid/expired/used token

- DELETE /api/v1/auth/session
  - Auth: session cookie (defined as security scheme)
  - Response 204: session invalidated
  - Response 401: not authenticated

Also define the session cookie security scheme in the spec components.

Run `make generate` after spec changes and commit the regenerated file.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All four endpoints present in the OpenAPI spec with correct methods, paths, and response shapes
- [ ] #2 GET and POST /magiclink/verify are distinct operations in the spec
- [ ] #3 make generate succeeds and produces updated api.gen.go
- [ ] #4 Generated code is committed
- [ ] #5 No manual edits to api.gen.go
<!-- AC:END -->
