import { Component, For, Show, createEffect } from 'solid-js';
import { useVerbDrill } from '../hooks/useVerbDrill';
import DrillDisplay from './DrillDisplay';
import AnswerInput from './AnswerInput';

const pronounOrder = ['je', 'tu', 'il/elle', 'nous', 'vous', 'ils/elles'] as const;

interface VerbDrillProps {
  verb: string;
  tense: string;
}

const VerbDrill: Component<VerbDrillProps> = (props) => {
  const [state, actions] = useVerbDrill(
    () => props.verb,
    () => props.tense,
  );

  let buttonRef: HTMLButtonElement | undefined;

  createEffect(() => {
    if (state.isSubmitted() && buttonRef) {
      buttonRef.focus();
    }
  });

  return (
    <div class="space-y-6">
      <Show when={state.isLoading()}>
        <div class="text-muted-foreground py-8 text-center">Loading...</div>
      </Show>

      <Show when={state.error()}>
        <div class="border-destructive bg-destructive/10 rounded-[var(--radius-md)] border px-4 py-3">
          <p class="text-destructive-foreground font-medium">Error: {state.error()}</p>
        </div>
      </Show>

      <Show when={!state.isLoading() && !state.error() && state.currentItem()}>
        <div class="border-border bg-secondary/50 dark:bg-secondary/30 rounded-[var(--radius-xl)] border p-8">
          <div class="mb-8">
            <DrillDisplay item={state.currentItem()!} />
          </div>

          <div class="space-y-4">
            <For each={pronounOrder}>
              {(pronounKey) => {
                const actualPronoun =
                  pronounKey === 'il/elle' ? 'il' : pronounKey === 'ils/elles' ? 'ils' : pronounKey;

                return (
                  <div class="flex items-center gap-4">
                    <div class="flex-1">
                      <AnswerInput
                        value={state.userAnswers()[actualPronoun] || ''}
                        onInput={(val) => actions.setUserAnswer(actualPronoun, val)}
                        disabled={state.isSubmitted()}
                        pronoun={actualPronoun}
                        answerState={state.answerStates()[actualPronoun]}
                        correctAnswer={state.correctAnswers()[actualPronoun] || ''}
                        autoFocus={pronounKey === 'je'}
                      />
                    </div>
                  </div>
                );
              }}
            </For>
          </div>

          <Show when={state.isSubmitted()}>
            <div class="bg-muted/50 mt-6 rounded-[var(--radius-md)] border px-4 py-3">
              <p class="text-center font-medium">
                Score: {state.correctCount()} / {state.totalCount()}
              </p>
            </div>
          </Show>

          <div class="mt-4 flex justify-end">
            <button
              ref={buttonRef}
              type="button"
              onClick={() => (state.isSubmitted() ? actions.reset() : actions.submitBatch())}
              disabled={
                !state.isSubmitted() && Object.values(state.userAnswers()).every((a) => !a.trim())
              }
              class="bg-primary text-primary-foreground inline-flex h-10 items-center rounded-[var(--radius)] px-8 text-sm font-medium transition-colors hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-50"
            >
              {state.isSubmitted() ? 'Next' : 'Submit'}
            </button>
          </div>
        </div>
      </Show>
    </div>
  );
};

export default VerbDrill;
