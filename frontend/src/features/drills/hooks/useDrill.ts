import { createSignal, createEffect } from 'solid-js';
import { DrillItem, DrillData } from '../types';
import { drillProvider } from '../provider';
import { checkAnswer } from '../checkAnswer';

export type AnswerState = 'unanswered' | 'correct' | 'incorrect';

export interface DrillState {
  currentItem: () => DrillItem | null;
  userAnswer: () => string;
  answerState: () => AnswerState;
  isLoading: () => boolean;
  error: () => string | null;
}

export interface DrillActions {
  setUserAnswer: (answer: string) => void;
  submitAnswer: () => void;
  nextQuestion: () => void;
}

export function useDrill(verb: () => string, tense: () => string): [DrillState, DrillActions] {
  const [drillData, setDrillData] = createSignal<DrillData | null>(null);
  const [currentItem, setCurrentItem] = createSignal<DrillItem | null>(null);
  const [userAnswer, setUserAnswer] = createSignal('');
  const [answerState, setAnswerState] = createSignal<AnswerState>('unanswered');
  const [isLoading, setIsLoading] = createSignal(true);
  const [error, setError] = createSignal<string | null>(null);

  const loadDrillData = () => {
    const v = verb();
    const t = tense();
    setIsLoading(true);
    setError(null);
    const result = drillProvider.getDrillData(v, t);
    if (result.ok) {
      setDrillData(result.data);
      pickRandomItem();
    } else {
      setError(result.error);
    }
    setIsLoading(false);
  };

  const pickRandomItem = () => {
    const data = drillData();
    if (!data || data.items.length === 0) return;
    const candidates =
      data.items.length > 1 ? data.items.filter((i) => i !== currentItem()) : data.items;
    setCurrentItem(candidates[Math.floor(Math.random() * candidates.length)]);
  };

  const submitAnswer = () => {
    const item = currentItem();
    if (!item || answerState() !== 'unanswered') return;
    const isCorrect = checkAnswer(userAnswer(), item.expectedAnswer.text);
    setAnswerState(isCorrect ? 'correct' : 'incorrect');
  };

  const nextQuestion = () => {
    setUserAnswer('');
    setAnswerState('unanswered');
    pickRandomItem();
  };

  createEffect(() => {
    loadDrillData();
  });

  const state: DrillState = {
    currentItem,
    userAnswer,
    answerState,
    isLoading,
    error,
  };

  const actions: DrillActions = {
    setUserAnswer,
    submitAnswer,
    nextQuestion,
  };

  return [state, actions];
}
