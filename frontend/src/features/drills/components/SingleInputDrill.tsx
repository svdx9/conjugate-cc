import { Component, For, Show, createMemo } from 'solid-js';
import { useDrill } from '../hooks/useDrill';
import DrillDisplay from './DrillDisplay';
import AnswerInput from './AnswerInput';

interface SingleInputDrillProps {
  verb: string;
  tense: string;
}

const SingleInputDrill: Component<SingleInputDrillProps> = (props) => {
  const [state, actions] = useDrill(() => props.verb, () => props.tense);

  const totalQuestions = createMemo(() => state.drillData()?.items.length ?? 0);

  // Current question number (1-indexed) within this drill
  // Uses currentIndex from the drill state, which represents the current item position
  const questionNumber = createMemo(() => state.currentItem() ? state.currentIndex() + 1 : 0);

  return (
    <div class="space-y-6">
      <Show when={state.isLoading()}>
        <div class="py-8 text-center text-muted-foreground">Loading drill...</div>
      </Show>

      <Show when={state.error()}>
        <div class="rounded-[var(--radius-md)] border border-destructive bg-destructive/10 px-4 py-3">
          <p class="font-medium text-destructive-foreground">Error: {state.error()}</p>
        </div>
      </Show>

      <Show when={!state.isLoading() && !state.error() && !state.isComplete() && state.currentItem()}>
        {/* Exercise card */}
        <div class="rounded-[var(--radius-xl)] border border-border bg-secondary/50 p-8 dark:bg-secondary/30">
          <div class="mb-8">
            <DrillDisplay item={state.currentItem()!} />
          </div>

          <AnswerInput
            value={state.userAnswer()}
            onInput={actions.setUserAnswer}
            onSubmit={actions.submitAnswer}
            onReset={actions.nextQuestion}
            disabled={false}
            pronoun={state.currentItem()?.prompt.pronoun}
            answerState={state.answerState()}
            correctAnswer={state.currentItem()?.expectedAnswer.text}
          />
        </div>

        {/* Progress indicator */}
        <Show when={totalQuestions() > 0}>
          <div class="border-t border-border pt-6">
            <div class="flex items-center justify-between text-sm text-muted-foreground">
              <span>
                Question {questionNumber()} of {totalQuestions()}
              </span>
              <div class="flex gap-1.5">
                <For each={Array.from({ length: totalQuestions() })}>
                  {(_, i) => (
                    <div
                      class={`h-1.5 w-7 rounded-full transition-all ${
                        i() < state.currentIndex()
                          ? 'bg-highlight'
                          : i() === state.currentIndex()
                            ? 'bg-foreground/20'
                            : 'bg-border'
                      }`}
                    />
                  )}
                </For>
              </div>
            </div>
          </div>
        </Show>
      </Show>

      <Show when={!state.isLoading() && !state.error() && state.isComplete()}>
        <div class="rounded-[var(--radius-xl)] border border-border bg-secondary/50 p-8 dark:bg-secondary/30">
          <h2 class="text-xl font-bold text-foreground">Drill Complete!</h2>
          <p class="mt-2 text-lg text-muted-foreground">
            Your score: {state.score().correct} / {state.score().total}
          </p>
          <button
            type="button"
            onClick={actions.resetDrill}
            class="mt-6 inline-flex h-10 items-center rounded-[var(--radius)] bg-primary px-6 text-sm font-medium text-primary-foreground transition-colors hover:opacity-90"
          >
            Start Over
          </button>
        </div>
      </Show>
    </div>
  );
};

export default SingleInputDrill;
