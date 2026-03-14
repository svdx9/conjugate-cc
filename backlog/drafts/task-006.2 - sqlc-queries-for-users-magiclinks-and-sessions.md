---
id: TASK-006.2
title: 'sqlc queries for users, magiclinks, and sessions'
status: To Do
assignee: []
created_date: '2026-03-12 01:11'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - database
  - authentication
dependencies:
  - TASK-006.1
parent_task_id: TASK-006
priority: high
ordinal: 13000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Write sqlc-managed SQL query files and regenerate Go code for all three tables.

Required queries:

users:
- UpsertUser — insert on conflict (email) do nothing, return the row; used during magic link request to find-or-create the user

magiclinks:
- InsertMagicLink — insert a new magic link record for a given user_id
- GetMagicLinkByTokenHash — look up by token_hash (used during verify)
- MarkMagicLinkUsed — set used_at = now() by id; used atomically in the verify transaction
- DeleteExpiredMagicLinks — cleanup helper

sessions:
- InsertSession — create a new session for a given user_id
- GetSessionByID — look up a session by id (used in middleware)
- InvalidateSession — set invalidated_at = now() by id (logout)
- DeleteExpiredSessions — cleanup helper
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 sqlc generate produces no errors
- [ ] #2 Generated Go types match the migration schema (table and column names)
- [ ] #3 All listed queries have corresponding .sql files under sql/queries/
- [ ] #4 Generated code is committed
<!-- AC:END -->
