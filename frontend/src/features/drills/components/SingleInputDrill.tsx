import { Component, Show } from 'solid-js';
import { useDrill } from '../hooks/useDrill';
import DrillDisplay from './DrillDisplay';
import AnswerInput from './AnswerInput';

interface SingleInputDrillProps {
  verb: string;
  tense: string;
}

const SingleInputDrill: Component<SingleInputDrillProps> = (props) => {
  const [state, actions] = useDrill(
    () => props.verb,
    () => props.tense,
  );

  const canSubmit = () =>
    state.userAnswer().trim().length > 0 && state.answerState() === 'unanswered';
  const canGoNext = () => state.answerState() !== 'unanswered';
  const showNext = () => state.answerState() !== 'unanswered';

  return (
    <div class="space-y-6">
      <Show when={state.isLoading()}>
        <div class="text-muted-foreground py-8 text-center">Loading drill...</div>
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

          <AnswerInput
            value={state.userAnswer()}
            onInput={actions.setUserAnswer}
            disabled={state.answerState() !== 'unanswered'}
            pronoun={state.currentItem()?.prompt.pronoun}
            answerState={state.answerState()}
            correctAnswer={state.currentItem()?.expectedAnswer.text}
          />

          <div class="mt-4 flex justify-end">
            <button
              type="button"
              onClick={() => (showNext() ? actions.nextQuestion() : actions.submitAnswer())}
              disabled={showNext() ? !canGoNext() : !canSubmit()}
              class="bg-primary text-primary-foreground inline-flex h-10 items-center rounded-[var(--radius)] px-8 text-sm font-medium transition-colors hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-50"
            >
              {showNext() ? 'Next' : 'Submit'}
            </button>
          </div>
        </div>
      </Show>
    </div>
  );
};

export default SingleInputDrill;
