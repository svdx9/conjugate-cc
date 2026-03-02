---
id: TASK-3
title: Deliver the conjugation drill interface with stubbed etre data
status: To Do
assignee: []
created_date: '2026-03-02 23:31'
updated_date: '2026-03-02 23:32'
labels:
  - frontend
  - ui
  - drills
dependencies:
  - TASK-1.5
references:
  - README.md
  - AGENTS.md
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Plan and track the first playable conjugation drill interface in the SolidJS frontend. This slice should use a stubbed `etre` verb dataset instead of a real conjugation engine, while still delivering the intended learner flow from prompt, to answer entry, to answer reveal, to a submitted results page with score summary.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Child tasks exist for the stub drill data contract, the six-input drill form UI, the show-answer and hide-answer behavior, and the submitted results and score page.
- [ ] #2 The tracked work makes it explicit that `etre` is temporary stub data and that the conjugation engine will replace it later without changing the learner-facing drill flow.
- [ ] #3 Each child task includes verification expectations that are specific enough for an independent agent to implement and validate the feature.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Complete `TASK-3.1` first to define the temporary `etre`-only drill contract, including pronoun ordering, expected answers, and score inputs.
2. Complete `TASK-3.2` next to build the main drill page layout with pronouns on the left and six answer inputs on the right.
3. Complete `TASK-3.3` after the base form exists to add the show-answer and hide-answer interaction while preserving typed user responses.
4. Complete `TASK-3.4` last to submit the learner's answers into a dedicated results page that shows correct answers and a summary score.
<!-- SECTION:PLAN:END -->
