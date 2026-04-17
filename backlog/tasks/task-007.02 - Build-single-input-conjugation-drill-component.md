---
id: TASK-007.02
title: Build single-input conjugation drill component
status: Done
assignee: []
created_date: '2026-04-09 12:00'
updated_date: '2026-04-17 15:12'
labels:
  - frontend
  - ui
  - drills
dependencies:
  - TASK-007.01
references:
  - README.md
parent_task_id: TASK-007
ordinal: 72000
---
## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
1. Build the basic conjugation drill component as a single-input interface. The component displays:

- The **pronoun** (e.g. "je")
- The **infinitive** (e.g. "etre")
- The **tense** (e.g. "Present")

2. The SingleInputDrill will take an inifitive and a tense as props
3. A pronoun will be chosen at random for the user to answer
4. The user sees a single text input where they type the correct conjugated verb phrase.
5. On submit, the component shows whether the answer is correct or incorrect:
   1. If the answer is wrong then the input box will be shown in an error state with a micro animation, the user supplied answer will show as strikethrough followed by the correct answer
   2. If the answer is correct the input box will be shown in a success state
6. Then another pronoun is chosent at random.
7. There is no scoring, just a simple answer/question structure

This is the simplest drill case — one prompt, one input, one answer.
<!-- SECTION:DESCRIPTION:END -->

## Implementation Plan

### File Structure
```
frontend/src/features/drills/
├── components/
│   ├── SingleInputDrill.tsx      # Main component
│   ├── DrillDisplay.tsx          # Displays prompt info
│   ├── AnswerInput.tsx           # Text input component
│   ├── ResultFeedback.tsx        # Shows correct/incorrect feedback
│   └── index.ts                  # Component exports
├── hooks/
│   └── useDrill.ts               # Custom hook for drill logic
└── styles/                       # Component-specific styles
```

### Component Breakdown

**SingleInputDrill.tsx (Main Component)**
- Manages state for user answers and submission
- Coordinates between sub-components
- Handles drill progression and scoring

**DrillDisplay.tsx**
- Displays the current prompt (pronoun, infinitive, tense)
- Clean, readable presentation of drill information

**AnswerInput.tsx**
- Text input for user's conjugated verb answer
- Form submission handling
- Input validation

**ResultFeedback.tsx**
- Shows correct/incorrect feedback
- Displays expected answer when wrong
- Next question button

**useDrill.ts (Custom Hook)**
- State management for current drill item
- Answer checking logic
- Score tracking
- Question progression

### Technical Approach

1. **State Management**: Use SolidJS signals for reactive state
2. **Provider Integration**: Consume from DrillProvider interface (TASK-007.01)
3. **Error Handling**: Use Result<T> pattern for robust error handling
4. **Testing**: Comprehensive unit and integration tests
5. **Styling**: Component-specific CSS with consistent design system

### Timeline Estimate
- **Total**: ~1 week
- **Step 1**: File structure and skeletons (1 day)
- **Step 2**: Core logic and provider integration (2 days)
- **Step 3**: UI components implementation (2 days)
- **Step 4**: Testing and quality assurance (1 day)
- **Step 5**: Review and refinements (1 day)

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The component displays the pronoun, infinitive, and tense for the current drill prompt.
- [ ] #2 A single text input accepts the user's answer (verb phrase).
- [ ] #3 On submit, the answer is compared to the expected value and correct/incorrect feedback is shown.
- [ ] #4 If incorrect, the correct answer is displayed.
- [ ] #5 The component consumes data from the provider interface defined in TASK-007.01.
- [ ] #6 Component handles error states gracefully (e.g., invalid drill data).
- [ ] #8 User can progress to next question after answering.
- [ ] #9 All UI elements are accessible (keyboard navigation, screen reader support).
- [ ] #10 Comprehensive test coverage (unit and integration tests).
<!-- AC:END -->
