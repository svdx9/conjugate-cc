---
id: TASK-002.6
title: Integrate the engine through a clean consumer boundary
status: To Do
assignee: []
created_date: '2026-03-02 19:05'
updated_date: '2026-03-14 16:26'
labels:
  - domain
  - engine
  - integration
dependencies:
  - TASK-002.5
references:
  - README.md
  - AGENTS.md
parent_task_id: TASK-002
ordinal: 5000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Wire the completed engine into the current codebase through a narrow consumer-facing seam so application code can use conjugation results without absorbing engine internals. This task should finish the separation story by documenting where orchestration ends and where UI or app-specific presentation begins.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 At least one real application consumer calls the engine only through its public API and does not import internal rule or data modules.
- [ ] #2 Any mapping between engine outputs and UI-facing view models happens outside the engine package.
- [ ] #3 Repository documentation explains where the engine lives, how consumers should import it, and what architectural boundaries must be preserved.
- [ ] #4 Verification demonstrates that the engine remains independently testable even after application integration is added.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Identify the first application-level consumer for the engine and add a thin adapter if the consumer needs a view-specific shape.
2. Keep all output shaping, labeling, and presentation concerns outside the engine package, using the engine result as input rather than extending the engine with UI state.
3. Update repository documentation to explain the separation between engine internals, public API, and consumer code.
4. Verify both the standalone engine tests and the first integrated usage path so the boundary is enforced in practice rather than only described.
<!-- SECTION:PLAN:END -->
