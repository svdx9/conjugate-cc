---
id: TASK-008
title: User Sign-In Functionality & Magic Link Authentication
status: In Progress
assignee: []
created_date: '2026-04-09 15:50'
updated_date: '2026-04-12 05:28'
labels: []
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement user sign-in functionality and authentication via a magic link system. This includes defining API contracts, building frontend UI components, and implementing backend handlers and database storage.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All auth endpoints defined in OpenAPI spec and code generated (008.01)
- [ ] #2 Database infrastructure in place with migrations for users, magic_links, sessions (008.07)
- [ ] #3 Auth service implements secure token generation, hashing, and session management (008.08)
- [ ] #4 All 4 auth HTTP handlers return correct responses per spec (008.09)
- [ ] #5 Frontend UI components for sign-in, sign-up, logout, and current user indicator (008.04)
- [ ] #6 Magic link flow works end-to-end (request → verify → confirm → session)
- [ ] #7 Sessions persist in PostgreSQL and survive server restart
- [ ] #8 Expired or used tokens are rejected
- [ ] #9 Logout invalidates the session immediately
<!-- AC:END -->
