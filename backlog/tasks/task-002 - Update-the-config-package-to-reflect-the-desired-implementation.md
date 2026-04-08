---
id: TASK-002
title: Update the config package to reflect the desired implementation
status: Done
assignee: []
created_date: '2026-03-29 17:56'
updated_date: '2026-03-29 00:00'
labels: []
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
implement as per for go-config-management skill

**MANDATORY REQUIREMENTS:**
	- The internal.config package only responsibilty is parsing environment variable into the config.Config struct. No other functionality is present in this package
	- The only exported function is config.FromEnv
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 `config.FromEnv()` is the only exported function in the config package
- [x] #2 `Config` is returned as a value type (not pointer)
- [x] #3 Port and host are validated at parse time with exported sentinel errors
- [x] #4 `getEnvOrDefault` trims whitespace; empty string falls back to default
- [x] #5 Backend starts cleanly with `make dev` via air hot-reload
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### Changes

| File | Action |
|------|--------|
| `backend/internal/config/config.go` | Rewrite — replace `Load()` with `FromEnv()`, value return, exported errors, whitespace trimming |
| `backend/cmd/server/main.go` | Update call site: `config.Load()` → `config.FromEnv()` |
| `backend/Makefile` | Add `dev` target wrapping `air -c .air.toml` |
| `backend/.air.toml` | New — air hot-reload config, proxies port 8080→8081 |

### Config package changes

- Removed `type envKey string` — plain string constants suffice
- Removed `Load() *Config` and `Validate()` — replaced by `FromEnv() Config`
- Added `getEnvOrDefault` (with `strings.TrimSpace`) and `getEnvOrDefaultInt`
- Exported `ErrPortOutOfRange` and `ErrInvalidHost` (were unexported)
- `validateHost` accepts both valid IPs and resolvable hostnames (via `net.LookupHost`)
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Rewrote `internal/config` to follow the go-config-management skill contract. `FromEnv()` is now the sole exported function, returning `Config` by value. Helper functions trim whitespace and validate at parse time. Exported sentinel errors allow callers to type-assert. Added `make dev` target and `.air.toml` for hot-reload during local development.
<!-- SECTION:FINAL_SUMMARY:END -->
