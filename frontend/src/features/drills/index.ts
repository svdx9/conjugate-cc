// Main exports for the drills feature

export type { Pronoun, Tense, DrillPrompt, ExpectedAnswer, DrillItem, DrillData } from './types';

export { type DrillProvider, drillProvider, validPronouns } from './provider';

export { SingleInputDrill, DrillDisplay, AnswerInput, ResultFeedback } from './components';
export { useDrill } from './hooks/useDrill';
