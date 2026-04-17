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
            onSubmit={actions.submitAnswer}
            onReset={actions.nextQuestion}
            disabled={false}
            pronoun={state.currentItem()?.prompt.pronoun}
            answerState={state.answerState()}
            correctAnswer={state.currentItem()?.expectedAnswer.text}
          />
        </div>
      </Show>
    </div>
  );
};

export default SingleInputDrill;
