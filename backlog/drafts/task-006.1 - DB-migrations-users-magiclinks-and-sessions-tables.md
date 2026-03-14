---
id: TASK-006.1
title: 'DB migrations: users, magiclinks, and sessions tables'
status: To Do
assignee: []
created_date: '2026-03-12 01:11'
updated_date: '2026-03-14 16:26'
labels:
  - backend
  - database
  - authentication
dependencies: []
parent_task_id: TASK-006
priority: high
ordinal: 12000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Write and apply PostgreSQL migrations for the three tables required by the magic link auth system.

`users`:
- id (UUID PK)
- email (text, not null, unique)
- created_at (timestamptz, not null, default now())

`magiclinks`:
- id (UUID PK)
- user_id (UUID, not null, FK → users.id ON DELETE CASCADE)
- token_hash (text, not null, unique) — store SHA-256 hash; never the raw token
- expires_at (timestamptz, not null)
- used_at (timestamptz, nullable) — set on first consumption; presence prevents replay
- created_at (timestamptz, not null, default now())

`sessions`:
- id (UUID PK) — this is the session cookie value
- user_id (UUID, not null, FK → users.id ON DELETE CASCADE)
- created_at (timestamptz, not null, default now())
- expires_at (timestamptz, not null)
- invalidated_at (timestamptz, nullable) — set on logout

All migrations must have corresponding down migrations.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Up migration creates all three tables with correct columns, constraints, and foreign keys
- [ ] #2 Down migration drops all three tables cleanly
- [ ] #3 migrate up && migrate down leaves schema unchanged (idempotent round-trip)
- [ ] #4 Unique index on users(email)
- [ ] #5 Unique index on magiclinks(token_hash)
- [ ] #6 Index on sessions(user_id)
<!-- AC:END -->
