---
id: TASK-007
title: Deliver the conjugation drill interface
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-09 12:00'
labels:
  - frontend
  - ui
  - drills
  - epic
dependencies:
  - TASK-001.5
references:
  - README.md
ordinal: 7000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Epic tracker for the conjugation drill interface, delivered in three phases:

1. **Single-input drill with stub data** — show one pronoun + infinitive + tense, user types the conjugated verb phrase (single word, aux+word, or reflexive pronoun+verb). Uses hardcoded stub data.
2. **Full conjugation form** — expand to all 6 standard pronouns (je, tu, il/elle, nous, vous, ils/elles) with submit and results summary.
3. **Conjugation engine** — port and adapt an existing conjugation engine implementation to replace stub data, enabling drills for all French verbs.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Child tasks exist for the drill contract/stub, single-input component, six-input form, results summary, engine port, and engine integration.
- [ ] #2 Phase 1 delivers a working single-input drill with stub data.
- [ ] #3 Phase 2 expands to a full six-pronoun conjugation form with scoring.
- [ ] #4 Phase 3 replaces stubs with a ported conjugation engine covering regular and irregular verbs.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Complete `TASK-007.01` to define the drill contract (types + provider interface) and stub dataset.
2. Complete `TASK-007.02` to build the single-input drill component.
3. Complete `TASK-007.03` to expand to the six-input full conjugation form.
4. Complete `TASK-007.04` to add the results summary after submission.
5. Complete `TASK-007.05` to port and adapt the existing conjugation engine.
6. Complete `TASK-007.06` to integrate the engine with the drill interface, replacing stubs.
<!-- SECTION:PLAN:END -->
