---
id: TASK-007.05
title: Build conjugation engine as pure domain module
status: To Do
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-17 13:46'
labels:
  - frontend
  - domain
  - engine
dependencies:
  - TASK-007.01
references:
  - backlog/docs/doc-002 - Frontend-Architecture.md
  - frontend/src/features/drills/provider.ts
  - frontend/src/features/drills/types.ts
parent_task_id: TASK-007
ordinal: 75000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Build the conjugation engine in `features/drills/engine/` — a pure domain module that replaces the hardcoded `verbData` in `StubDrillProvider`.

Given a verb infinitive and a tense, the engine produces all pronoun conjugations. It must handle:
- All tenses (présent, imparfait, passé composé, futur, and future additions)
- Regular verbs (-er, -ir, -re)
- Irregular verbs (être, avoir, aller, faire, dire, pouvoir, vouloir, savoir, voir, venir, prendre, etc.)
- Reflexive verbs (pronoun agreement, elision before vowels)
- Compound tenses (auxiliary selection: avoir vs être, past participle agreement)
- Stem-changing verbs (e.g. appeler, acheter, manger)
- Negation (ne...pas, correct placement with compounds and reflexives)
- Multiple valid conjugations via `ExpectedAnswer.alternates?: string[]` (e.g. s'asseoir)

Pronoun synthesis (elle/on from il, elles from ils) carries forward from the existing `synthesizePronouns` function.

See: `backlog/docs/doc-002 - Frontend-Architecture.md` §7 (Conjugation Engine)
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Engine is a pure domain module in engine/ directory — no UI or framework imports.
- [ ] #2 Regular verbs (-er, -ir, -re) conjugate correctly across all supported tenses.
- [ ] #3 Common irregular verbs (at minimum: être, avoir, aller, faire, dire, pouvoir, vouloir, savoir) conjugate correctly.
- [ ] #4 Reflexive verbs produce correct reflexive pronoun + verb forms.
- [ ] #5 Compound tenses select the correct auxiliary (avoir vs être) and apply past participle agreement.
- [ ] #6 Multiple valid conjugations are returned via ExpectedAnswer.alternates for verbs like s'asseoir.
- [ ] #7 Pronoun synthesis derives elle/on from il and elles from ils.
- [ ] #8 Comprehensive unit tests verify regular, irregular, reflexive, and compound tense conjugations.
<!-- AC:END -->
