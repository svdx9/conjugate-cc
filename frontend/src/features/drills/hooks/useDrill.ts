import { createSignal, createMemo, createEffect, batch } from 'solid-js';
import { DrillItem, DrillData, Tense } from '../types';
import { drillProvider } from '../provider';

export type AnswerState = 'unanswered' | 'correct' | 'incorrect';

function fisherYatesShuffle<T>(array: T[]): T[] {
  const result = [...array];
  for (let i = result.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [result[i], result[j]] = [result[j], result[i]];
  }
  return result;
}

export interface DrillState {
  currentItem: () => DrillItem | null;
  currentIndex: () => number;
  drillData: () => DrillData | null;
  userAnswer: () => string;
  answerState: () => AnswerState;
  score: () => { correct: number; total: number };
  isLoading: () => boolean;
  error: () => string | null;
  isComplete: () => boolean;
  availableTenses: () => string[];
}

export interface DrillActions {
  setUserAnswer: (answer: string) => void;
  submitAnswer: () => void;
  nextQuestion: () => void;
  resetDrill: () => void;
  setVerb: (verb: string) => void;
  setTense: (tense: string) => void;
}

export function useDrill(verb: () => string, tense: () => string): [DrillState, DrillActions] {
  const [currentVerb, setCurrentVerb] = createSignal(verb());
  const [currentTense, setCurrentTense] = createSignal(tense());
  const [drillData, setDrillData] = createSignal<DrillData | null>(null);
  const [currentIndex, setCurrentIndex] = createSignal(0);
  const [userAnswer, setUserAnswer] = createSignal('');
  const [answerState, setAnswerState] = createSignal<AnswerState>('unanswered');
  const [score, setScore] = createSignal({ correct: 0, total: 0 });
  const [isLoading, setIsLoading] = createSignal(true);
  const [error, setError] = createSignal<string | null>(null);

  const currentItem = createMemo(() => {
    const data = drillData();
    if (!data) return null;
    const idx = currentIndex();
    return data.items[idx] ?? null;
  });

  const isComplete = createMemo(() => {
    const data = drillData();
    if (!data) return false;
    return currentIndex() >= data.items.length;
  });

  const loadDrillData = (v: string, t: string) => {
    setIsLoading(true);
    setError(null);
    const result = drillProvider.getDrillData(v, t as Tense);
    if (!result.ok && result.code === 'INVALID_TENSE') {
      setError(`Invalid tense "${t}"`);
      setIsLoading(false);
      return;
    }
    if (result.ok) {
      const shuffled = fisherYatesShuffle(result.data.items);
      setDrillData({ ...result.data, items: shuffled });
      setCurrentIndex(0);
      setScore({ correct: 0, total: 0 });
      setUserAnswer('');
      setAnswerState('unanswered');
    } else {
      if (result.code === 'NOT_FOUND' && result.details?.availableTenses) {
        const availableTenses = result.details.availableTenses as string[];
        if (availableTenses.length > 0) {
          loadDrillData(v, availableTenses[0]);
          setIsLoading(false);
          return;
        }
      }
      setError(`No drill data available for "${v}" - "${t}"`);
    }
    setIsLoading(false);
  };

  const setVerb = (newVerb: string) => {
    setCurrentVerb(newVerb);
    loadDrillData(newVerb, currentTense());
  };

  const setTense = (newTense: string) => {
    const available = drillProvider.getAvailableTenses(currentVerb());
    if (!available.includes(newTense)) {
      if (available.length > 0) {
        setCurrentTense(available[0]);
        loadDrillData(currentVerb(), available[0]);
      }
      return;
    }
    setCurrentTense(newTense);
    loadDrillData(currentVerb(), newTense);
  };

  createEffect(() => {
    const v = verb();
    const t = tense();
    batch(() => {
      setCurrentVerb(v);
      setCurrentTense(t);
      loadDrillData(v, t);
    });
  });

  const checkAnswer = (userAns: string, expected: string): boolean => {
    const normalizedUser = userAns.trim().toLowerCase();
    const normalizedExpected = expected.trim().toLowerCase();
    return normalizedUser === normalizedExpected;
  };

  const submitAnswer = () => {
    const item = currentItem();
    if (!item) return;
    if (answerState() !== 'unanswered') return;

    const isCorrect = checkAnswer(userAnswer(), item.expectedAnswer.text);

    setAnswerState(isCorrect ? 'correct' : 'incorrect');
    setScore((prev) => ({
      correct: prev.correct + (isCorrect ? 1 : 0),
      total: prev.total + 1,
    }));
  };

  const nextQuestion = () => {
    setCurrentIndex((idx) => idx + 1);
    setUserAnswer('');
    setAnswerState('unanswered');
  };

  const resetDrill = () => {
    loadDrillData(currentVerb(), currentTense());
  };

  const state: DrillState = {
    currentItem: currentItem,
    currentIndex: currentIndex,
    drillData: drillData,
    userAnswer: userAnswer,
    answerState: answerState,
    score: score,
    isLoading: isLoading,
    error: error,
    isComplete: isComplete,
    availableTenses: () => drillProvider.getAvailableTenses(currentVerb()),
  };

  const actions: DrillActions = {
    setUserAnswer,
    submitAnswer,
    nextQuestion,
    resetDrill,
    setVerb,
    setTense,
  };

  return [state, actions];
}