---
id: TASK-008.06
title: Implement production email delivery (Amazon SES)
status: To Do
assignee: []
created_date: '2026-04-11 17:54'
labels:
  - backend
  - email
  - authentication
dependencies:
  - TASK-008.02
parent_task_id: TASK-008
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the production email sender using Amazon SES for delivering magic link emails.

Task 008.02 will use a stub/log email sender for development. This task covers building the real email delivery infrastructure:

- Implement `internal/email/` sender interface backed by Amazon SES
- Magic link email template using `templ`
- SES configuration (region, credentials) via environment variables
- Email delivery error handling and retry logic
- Integration testing with SES sandbox
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Amazon SES email sender implements the email sender interface
- [ ] #2 Magic link emails are delivered via SES in production
- [ ] #3 Email templates are rendered using templ
- [ ] #4 SES configuration is externalized via environment variables
- [ ] #5 Stub sender remains available for local development
- [ ] #6 Email delivery failures are logged and handled gracefully
<!-- AC:END -->
