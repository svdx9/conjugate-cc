---
id: TASK-006.10
title: Compose magic link emails using go templ
status: To Do
assignee: []
created_date: '2026-03-12 01:28'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - authentication
  - email
dependencies:
  - TASK-006.4
  - TASK-006.9
parent_task_id: TASK-006
priority: medium
ordinal: 21000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Use `templ` (already in the approved stack) to define the magic link email template and generate the rendering code.

Template outputs:
- HTML body — branded email with a clear call-to-action button/link.
- Plain-text body — fallback for clients that do not render HTML.

Template location: `internal/auth/email/magiclink.templ`

The template receives a single data struct:
```go
type MagicLinkEmailData struct {
    MagicLinkURL string
    ExpiresIn    string // human-readable, e.g. "15 minutes"
}
```

Generated code (`magiclink_templ.go`) MUST be committed. Generation runs via `make generate` (add a `templ generate` step if not already present).

Integration:
- `SESEmailSender` (TASK-006.9) calls the generated render functions to produce HTML and plain-text bodies before calling the SES API.
- `LogEmailSender` (TASK-006.4) MAY render the plain-text template and log the output, or continue logging the raw URL — either is acceptable for a dev stub.
- The `EmailSender` interface signature (`SendMagicLink(ctx, toEmail, magicLinkURL string) error`) does NOT change; rendering is an internal implementation detail of each sender.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 magiclink.templ defines both an HTML component and a plain-text component
- [ ] #2 templ generate produces magiclink_templ.go with no errors
- [ ] #3 Generated file is committed
- [ ] #4 make generate includes the templ generation step
- [ ] #5 SESEmailSender uses the generated HTML and plain-text render functions when building the SES SendEmail request
- [ ] #6 Template renders the magic link URL and expiry duration correctly (unit test against rendered output)
<!-- AC:END -->
