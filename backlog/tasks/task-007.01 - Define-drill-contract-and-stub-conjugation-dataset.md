---
id: TASK-007.01
title: Define drill contract and stub conjugation dataset
status: Done
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-09 09:36'
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
- [ ] #4 The stub provider can drive both the single-input drill (TASK-007.02) and the six-input form (TASK-007.03) without modification.
<!-- AC:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Implemented drill contract and stub provider with the following:

1. ✅ TypeScript types defined in `frontend/src/features/drills/types.ts`:
   - Pronoun type with all French subject pronouns
   - Tense type with common French tenses
   - DrillPrompt interface (infinitive, tense, pronoun)
   - ExpectedAnswer interface supporting single words, aux+participle, and reflexive forms
   - DrillItem and DrillData types for complete drill datasets

2. ✅ **DrillProvider interface** defined in `frontend/src/features/drills/provider.ts`:
   - Clear contract with JSDoc documentation
   - `getDrillData(verb: string, tense: string): DrillData` method
   - `getDrillItem(verb: string, tense: string, pronoun: string): DrillItem` method for specific conjugations
   - Explicitly documented as the contract for future conjugation engine
   - Returns DrillData for synchronous access

3. ✅ Stub implementation with hardcoded data for 3 verbs:
   - être (present tense) - irregular verb
   - avoir (present tense) - irregular verb  
   - se laver (present tense) - reflexive verb with proper flagging
    - Clearly marked as temporary implementation to be replaced by real engine
    - Added getDrillItem method for specific conjugations

4. ✅ Contract supports both single-input and six-input drill formats:
   - Single DrillItem contains all needed data
   - Array of items provides full conjugation set
   - ExpectedAnswer supports all required answer formats
    - getDrillItem method enables targeted drill experiences

5. ✅ Comprehensive test suite in `provider.test.ts` with 7 tests covering:
   - Basic functionality for each verb
   - Case insensitivity
   - Reflexive verb handling
   - Fallback behavior for unknown verbs
   - Pronoun filtering functionality
   - Full conjugation set retrieval

6. ✅ Clean exports via `index.ts` for easy importing

7. ✅ All linting errors resolved and TypeScript compilation passes

8. ✅ Clear documentation showing how to swap stub for real conjugation engine

The implementation is ready to drive both TASK-007.02 (single-input drill) and TASK-007.03 (six-input form) without modification. The DrillProvider interface provides a clear contract that the future conjugation engine must implement, now with enhanced pronoun filtering capabilities.
<!-- SECTION:FINAL_SUMMARY:END -->
