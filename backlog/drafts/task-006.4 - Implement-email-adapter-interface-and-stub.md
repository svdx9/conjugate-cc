---
id: TASK-006.4
title: Implement email adapter interface and stub
status: To Do
assignee: []
created_date: '2026-03-12 01:11'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - authentication
dependencies: []
parent_task_id: TASK-006
priority: medium
ordinal: 15000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define an `EmailSender` interface and a log-only stub implementation so the magic link request handler can be wired and tested without a real SMTP/API dependency.

Interface (owned by the auth feature package):
```go
type EmailSender interface {
    SendMagicLink(ctx context.Context, toEmail, magicLinkURL string) error
}
```

Implementations:
- `LogEmailSender` — logs the magic link URL via slog (used in dev/test)
- Real implementation deferred to a future task; the interface makes it swappable.

Wire `LogEmailSender` into the composition root for now. Config should read an `EMAIL_SENDER` env var (`log` as default) so a real sender can be plugged in later without code changes.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 EmailSender interface defined in the auth feature package
- [ ] #2 LogEmailSender logs the recipient and magic link URL at INFO level
- [ ] #3 Composition root wires LogEmailSender by default
- [ ] #4 Unit test confirms LogEmailSender does not return an error
<!-- AC:END -->
