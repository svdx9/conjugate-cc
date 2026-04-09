---
id: TASK-007.1
title: Define drill contract and stub conjugation dataset
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-09 12:00'
labels:
  - frontend
  - drills
  - data
dependencies:
  - TASK-001.5
references:
  - README.md
parent_task_id: TASK-007
ordinal: 71000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define the TypeScript types and provider interface for conjugation drill data. The contract must support:

- A drill prompt consisting of an infinitive, tense, and pronoun.
- An expected answer that can be a single word (e.g. "suis"), aux+participle (e.g. "ai mange"), or reflexive pronoun+verb (e.g. "me lave").
- A provider interface that returns drill data for a given verb/tense, so that both the stub implementation and the future conjugation engine conform to the same contract.

Implement a hardcoded stub provider with data for 2-3 verbs (e.g. etre present, avoir present, se laver present).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 TypeScript types are defined for drill prompt (infinitive, tense, pronoun) and expected answer (supporting single word, aux+word, and reflexive forms).
- [ ] #2 A provider interface is defined that returns drill data for a given verb and tense.
- [ ] #3 A stub provider implements the interface with hardcoded data for at least 2 verbs.
- [ ] #4 The stub provider can drive both the single-input drill (TASK-007.2) and the six-input form (TASK-007.3) without modification.
<!-- AC:END -->
