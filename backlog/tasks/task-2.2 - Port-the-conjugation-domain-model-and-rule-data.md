---
id: TASK-2.2
title: Port the conjugation domain model and rule data
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-02 19:05'
labels:
  - domain
  - engine
  - data
dependencies:
  - TASK-2.1
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-2
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Bring over the typed grammatical model and static rule data that the conjugation engine depends on from the existing reference implementation. This includes the canonical domain concepts, lookup structures, and exception data required before any conjugation pipeline can behave deterministically.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The engine defines typed domain primitives for the conjugation request and result space, including verb identity, grammatical categories, and engine output.
- [ ] #2 Static rule data from the reference implementation is represented in a deterministic, version-controlled form inside the engine package rather than being embedded ad hoc in UI code.
- [ ] #3 Internal data structures are normalized so the engine can resolve rules and overrides without stringly typed branching spread across the codebase.
- [ ] #4 The task documents any assumptions or intentional deviations from the reference implementation before downstream engine logic depends on them.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Identify the reference implementation modules that define the domain vocabulary, enumerations, and static conjugation datasets.
2. Port those concepts into local typed modules with names that match the current repository conventions.
3. Normalize any nested literals or lookup tables into engine-owned data modules that are easy to inspect and test.
4. Record the mapping between the reference structures and the new package layout so later tasks can verify parity rather than re-deriving assumptions.
<!-- SECTION:PLAN:END -->

