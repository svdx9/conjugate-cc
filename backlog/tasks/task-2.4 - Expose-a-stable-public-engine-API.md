---
id: TASK-2.4
title: Expose a stable public engine API
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-02 19:05'
labels:
  - domain
  - engine
  - api
dependencies:
  - TASK-2.3
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-2
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define the outward-facing API for the conjugation engine so other parts of the application can consume it without reaching into rule tables or implementation internals. The public surface should support both focused form lookup and richer engine-driven outputs needed by drill or table features, while staying framework-agnostic.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The engine exposes a small, documented public API with typed input and output contracts suitable for downstream consumers.
- [ ] #2 The public API hides internal rule tables, helper utilities, and exception-resolution details behind stable entrypoints.
- [ ] #3 Invalid or unsupported requests return explicit, typed failures or result states instead of consumer-hostile runtime ambiguity.
- [ ] #4 Public exports are organized so future consumers can generate a single conjugated form or a broader conjugation view without duplicating rule logic outside the engine.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Identify the minimal set of entrypoints needed from the existing engine, such as single-form lookup and grouped-form generation.
2. Define typed request and response contracts for those entrypoints and move them into the package boundary established earlier.
3. Hide internal implementation modules behind index-level exports so consumers cannot accidentally depend on unstable internals.
4. Document error semantics and intended usage patterns for application code that will call the engine.
<!-- SECTION:PLAN:END -->

