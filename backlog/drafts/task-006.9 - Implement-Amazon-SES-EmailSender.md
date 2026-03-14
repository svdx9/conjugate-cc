---
id: TASK-006.9
title: Implement Amazon SES EmailSender
status: To Do
assignee: []
created_date: '2026-03-12 01:27'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - authentication
  - email
dependencies:
  - TASK-006.4
parent_task_id: TASK-006
priority: medium
ordinal: 20000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement a concrete `EmailSender` backed by Amazon SES that can replace the `LogEmailSender` stub in production.

Use the AWS SDK for Go v2 (`github.com/aws/aws-sdk-go-v2`) — specifically the `ses` or `sesv2` client. Prefer SES v2 (`sesv2`) as it is the current recommended API.

Implementation:
- `SESEmailSender` struct implementing the `EmailSender` interface defined in TASK-006.4.
- Sends a plain-text email (and optionally HTML) containing the magic link URL.
- AWS credentials MUST be sourced from the default credential chain (env vars, instance profile, etc.) — no hardcoded keys.

Config (add to `internal/config`):
- `SES_FROM_ADDRESS` (required when `EMAIL_SENDER=ses`) — verified SES sender address.
- `SES_REGION` (required when `EMAIL_SENDER=ses`) — AWS region (e.g. `us-east-1`).
- `EMAIL_SENDER` env var (introduced in TASK-006.4): wire `SESEmailSender` when value is `ses`.

Wiring:
- Composition root reads `EMAIL_SENDER`; selects `SESEmailSender` or `LogEmailSender` accordingly.
- `SESEmailSender` MUST be constructed with an injected `sesv2.Client` so it is testable.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 SESEmailSender implements the EmailSender interface
- [ ] #2 EMAIL_SENDER=ses wires SESEmailSender in the composition root
- [ ] #3 SES_FROM_ADDRESS and SES_REGION are required config when EMAIL_SENDER=ses; startup fails fast if absent
- [ ] #4 AWS credentials are never hardcoded; default credential chain is used
- [ ] #5 Unit test stubs the sesv2 client and asserts SendEmail is called with correct To/From/body
- [ ] #6 LogEmailSender remains the default (EMAIL_SENDER unset or 'log')
<!-- AC:END -->
