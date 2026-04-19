---
id: TASK-008.04
title: Build Frontend UI Components
status: To Do
assignee: []
created_date: '2026-04-09 15:50'
updated_date: '2026-04-12 05:29'
labels: []
dependencies:
  - TASK-008.09
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the frontend UI components for signup, signin, logout, and indicating the current user.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Sign-in page with email input form that calls POST /api/v1/auth/magiclink/request
- [ ] #2 "Check your email" confirmation screen shown after requesting magic link
- [ ] #3 Magic link verification page that calls GET /api/v1/auth/magiclink/verify and shows confirmation UI
- [ ] #4 "Sign in as user@example.com?" confirmation with button that calls POST /api/v1/auth/magiclink/verify
- [ ] #5 Error pages for invalid/expired tokens (400, 410, 401 responses)
- [ ] #6 Logout button that calls DELETE /api/v1/auth/session and clears local auth state
- [ ] #7 Current user indicator showing signed-in user email
- [ ] #8 Auth state persists across page navigation via session cookie
- [ ] #9 Frontend sends X-Requested-With header on state-changing requests
<!-- AC:END -->
