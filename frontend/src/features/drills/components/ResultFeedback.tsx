import { Component, Show } from 'solid-js';
import { AnswerState } from '../hooks/useDrill';

interface ResultFeedbackProps {
  answerState: AnswerState;
  correctAnswer: string;
  isLast: boolean;
  onNext: () => void;
  onRestart: () => void;
}

const ResultFeedback: Component<ResultFeedbackProps> = (props) => {
  return (
    <div class="space-y-4">
      <Show when={props.answerState === 'correct'}>
        <div class="rounded-[var(--radius-md)] border border-green-500 bg-green-50 px-4 py-3 dark:bg-green-900/20">
          <p class="font-medium text-green-700 dark:text-green-400">Correct!</p>
        </div>
      </Show>

      <Show when={props.answerState === 'incorrect'}>
        <div class="rounded-[var(--radius-md)] border border-destructive bg-destructive/10 px-4 py-3">
          <p class="font-medium text-destructive-foreground">Incorrect</p>
          <p class="mt-1 text-sm text-muted-foreground">
            Correct answer: <span class="font-medium text-foreground">{props.correctAnswer}</span>
          </p>
        </div>
      </Show>

      <div class="flex gap-3">
        <Show when={!props.isLast}>
          <button
            type="button"
            onClick={() => props.onNext()}
            class="inline-flex h-10 items-center bg-primary px-6 text-sm font-medium text-primary-foreground rounded-[var(--radius)] transition-colors hover:opacity-90"
          >
            Next Question
          </button>
        </Show>
        <Show when={props.isLast}>
          <button
            type="button"
            onClick={() => props.onRestart()}
            class="inline-flex h-10 items-center bg-primary px-6 text-sm font-medium text-primary-foreground rounded-[var(--radius)] transition-colors hover:opacity-90"
          >
            Start Over
          </button>
        </Show>
      </div>
    </div>
  );
};

export default ResultFeedback;