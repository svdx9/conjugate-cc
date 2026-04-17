import { Component, For, Show } from 'solid-js';
import { useVerbDrill } from '../hooks/useVerbDrill';
import DrillDisplay from './DrillDisplay';
import AnswerInput from './AnswerInput';
import { DrillData } from '../types';

const pronounOrder = ['je', 'tu', 'il/elle', 'nous', 'vous', 'ils/elles'] as const;

interface VerbDrillProps {
  drillData: DrillData;
}

const VerbDrill: Component<VerbDrillProps> = (props) => {
  const [state, actions] = useVerbDrill(() => props.drillData);

  const currentItem = () => props.drillData.items[0] || null;

  return (
    <div class="space-y-6">
      <Show when={!props.drillData}>
        <div class="text-muted-foreground py-8 text-center">No drill data available.</div>
      </Show>

      <Show when={props.drillData}>
        <div class="border-border bg-secondary/50 dark:bg-secondary/30 rounded-[var(--radius-xl)] border p-8">
          <div class="mb-8">
            <Show when={currentItem()}>
              <DrillDisplay item={currentItem()} />
            </Show>
          </div>

          <div class="space-y-4">
            <For each={pronounOrder}>
              {(pronounKey) => {
                const actualPronoun =
                  pronounKey === 'il/elle' ? 'il' : pronounKey === 'ils/elles' ? 'ils' : pronounKey;
                const item = () =>
                  props.drillData.items.find((i) => i.prompt.pronoun === actualPronoun);
                const correctAnswer = () => {
                  const it = item();
                  if (!it) return '';
                  if (pronounKey === 'il/elle') {
                    const il = props.drillData.items.find((i) => i.prompt.pronoun === 'il');
                    const elle = props.drillData.items.find((i) => i.prompt.pronoun === 'elle');
                    return il?.expectedAnswer.text || elle?.expectedAnswer.text || '';
                  }
                  if (pronounKey === 'ils/elles') {
                    const ils = props.drillData.items.find((i) => i.prompt.pronoun === 'ils');
                    const elles = props.drillData.items.find((i) => i.prompt.pronoun === 'elles');
                    return ils?.expectedAnswer.text || elles?.expectedAnswer.text || '';
                  }
                  return it.expectedAnswer.text;
                };

                return (
                  <div class="flex items-center gap-4">
                    {/* <div class="w-20 shrink-0 text-right">
                      <span class="text-lg font-medium text-foreground">
                        {getPronounLabel(pronounKey)}
                      </span>
                    </div> */}
                    <div class="flex-1">
                      <AnswerInput
                        value={state.userAnswers()[actualPronoun] || ''}
                        onInput={(val) => actions.setUserAnswer(actualPronoun, val)}
                        disabled={state.isSubmitted()}
                        pronoun={actualPronoun}
                        answerState={state.answerStates()[actualPronoun]}
                        correctAnswer={correctAnswer()}
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
            <Show
              when={!state.isSubmitted()}
              fallback={
                <button
                  type="button"
                  onClick={actions.reset}
                  class="bg-primary text-primary-foreground inline-flex h-10 items-center rounded-[var(--radius)] px-8 text-sm font-medium transition-colors hover:opacity-90"
                >
                  Try Again
                </button>
              }
            >
              <button
                type="button"
                onClick={actions.submitBatch}
                disabled={Object.values(state.userAnswers()).every((a) => !a.trim())}
                class="bg-primary text-primary-foreground inline-flex h-10 items-center rounded-[var(--radius)] px-8 text-sm font-medium transition-colors hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-50"
              >
                Submit
              </button>
            </Show>
          </div>
        </div>
      </Show>
    </div>
  );
};

export default VerbDrill;
