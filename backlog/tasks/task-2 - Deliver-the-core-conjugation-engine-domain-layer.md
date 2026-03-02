---
id: TASK-2
title: Deliver the core conjugation engine domain layer
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-02 19:05'
labels:
  - domain
  - planning
  - engine
dependencies:
  - TASK-1.3
references:
  - README.md
  - AGENTS.md
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the task set needed to bring the conjugation engine into this repository as pure domain logic, based on the existing implementation currently living outside this codebase at `https://github.com/svdx9/conjuflo/tree/main/web/ui/src/domain/conjugator`. The delivered engine must remain cleanly separated from UI concerns so it can be exercised, tested, and reused independently of any frontend screens.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Child tasks exist for package boundary definition, rule/data porting, conjugation pipeline implementation, public API design, automated verification, and consumer integration.
- [ ] #2 The resulting plan preserves a strict domain boundary so the engine can be executed without importing UI, routing, rendering, or browser-specific code.
- [ ] #3 The tracked work defines verification and documentation expectations sufficient for an independent agent to implement the engine with predictable behavior.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Complete `TASK-2.1` first to establish the engine package boundary, public surface, and architectural constraints that keep the module free of UI dependencies.
2. Complete `TASK-2.2` next to port the language model primitives, static rule data, and typed lookup structures that the engine depends on.
3. Complete `TASK-2.3` after the core data structures exist to implement the conjugation pipeline itself, including classification, rule selection, and exception handling.
4. Complete `TASK-2.4` once the engine internals are working to expose a stable public API for single-form and table-generation use cases without leaking implementation details.
5. Complete `TASK-2.5` to capture the expected behavior with deterministic automated tests, including representative regular, irregular, and edge-case verbs.
6. Complete `TASK-2.6` last to integrate the engine into the current repository structure, document the clean separation rules, and verify that consumers use the module through the intended boundary.
<!-- SECTION:PLAN:END -->

