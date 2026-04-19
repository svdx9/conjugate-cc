import { createSignal, createEffect } from 'solid-js';
import { DrillData, DrillItem } from '../types';
import { drillProvider } from '../provider';
import { checkAnswer } from '../checkAnswer';
import { AnswerState } from './useDrill';

export interface VerbDrillState {
  currentItem: () => DrillItem | null;
  userAnswers: () => Record<string, string>;
  answerStates: () => Record<string, AnswerState>;
  correctAnswers: () => Record<string, string>;
  isSubmitted: () => boolean;
  isLoading: () => boolean;
  error: () => string | null;
  correctCount: () => number;
  totalCount: () => number;
}

export interface VerbDrillActions {
  setUserAnswer: (pronoun: string, answer: string) => void;
  submitBatch: () => void;
  reset: () => void;
}

export function useVerbDrill(
  verb: () => string,
  tense: () => string,
): [VerbDrillState, VerbDrillActions] {
  const corePronouns = ['je', 'tu', 'il', 'nous', 'vous', 'ils'];

  const [drillData, setDrillData] = createSignal<DrillData | null>(null);
  const [isLoading, setIsLoading] = createSignal(true);
  const [error, setError] = createSignal<string | null>(null);

  const loadDrillData = () => {
    setIsLoading(true);
    setError(null);
    const result = drillProvider.getDrillData(verb(), tense());
    if (result.ok) {
      setDrillData(result.data);
    } else {
      setError(result.error);
    }
    setIsLoading(false);
  };

  createEffect(() => {
    loadDrillData();
  });

  const buildEmptyAnswers = (): Record<string, string> => {
    const answers: Record<string, string> = {};
    corePronouns.forEach((p) => {
      answers[p] = '';
    });
    return answers;
  };

  const buildEmptyStates = (): Record<string, AnswerState> => {
    const states: Record<string, AnswerState> = {};
    corePronouns.forEach((p) => {
      states[p] = 'unanswered';
    });
    return states;
  };

  const [userAnswers, setUserAnswers] = createSignal<Record<string, string>>(buildEmptyAnswers());
  const [answerStates, setAnswerStates] =
    createSignal<Record<string, AnswerState>>(buildEmptyStates());
  const [isSubmitted, setIsSubmitted] = createSignal(false);
  const [correctCount, setCorrectCount] = createSignal(0);

  const setUserAnswer = (pronoun: string, answer: string) => {
    if (isSubmitted()) return;
    setUserAnswers((prev) => ({ ...prev, [pronoun]: answer }));
  };

  const submitBatch = () => {
    const data = drillData();
    if (!data || isSubmitted()) return;

    const newStates: Record<string, AnswerState> = {};
    let correct = 0;

    corePronouns.forEach((p) => {
      const item = data.items.find((i) => i.prompt.pronoun === p);
      if (!item) return;
      const userAns = userAnswers()[p] || '';
      const isCorrect = checkAnswer(userAns, item.expectedAnswer.text);
      newStates[p] = isCorrect ? 'correct' : 'incorrect';
      if (isCorrect) correct++;
    });

    setAnswerStates(newStates);
    setCorrectCount(correct);
    setIsSubmitted(true);
  };

  const reset = () => {
    setUserAnswers(buildEmptyAnswers());
    setAnswerStates(buildEmptyStates());
    setIsSubmitted(false);
    setCorrectCount(0);
  };

  const state: VerbDrillState = {
    currentItem: () => {
      const data = drillData();
      if (!data || data.items.length === 0) return null;
      return data.items[0];
    },
    userAnswers,
    answerStates,
    correctAnswers: () => {
      const data = drillData();
      if (!data) return {};
      const answers: Record<string, string> = {};
      corePronouns.forEach((p) => {
        const item = data.items.find((i) => i.prompt.pronoun === p);
        if (item) answers[p] = item.expectedAnswer.text;
      });
      return answers;
    },
    isSubmitted,
    isLoading,
    error,
    correctCount,
    totalCount: () => corePronouns.length,
  };

  const actions: VerbDrillActions = {
    setUserAnswer,
    submitBatch,
    reset,
  };

  return [state, actions];
}
