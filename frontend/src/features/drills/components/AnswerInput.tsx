import { Component, Show, createEffect, onMount } from 'solid-js';

const VOWELS = 'aeiouAEIOU';
const VOWEL_LIKE_CONSONANTS = 'hH';

const canElide = (userInput: string): boolean => {
  const first = userInput.trim()[0];
  if (!first) return false;
  return VOWELS.includes(first) || VOWEL_LIKE_CONSONANTS.includes(first);
};

const getElidedPronoun = (pronoun: string, userInput: string): string => {
  if (pronoun !== 'je') return pronoun;
  if (canElide(userInput)) return "j'";
  return pronoun;
};

interface AnswerInputProps {
  value: string;
  onInput: (value: string) => void;
  onSubmit: () => void;
  onReset?: () => void;
  disabled: boolean;
  pronoun?: string;
  answerState?: 'unanswered' | 'correct' | 'incorrect';
  correctAnswer?: string;
}

const AnswerInput: Component<AnswerInputProps> = (props) => {
  let inputRef: HTMLInputElement | undefined;

  createEffect(() => {
    if (props.answerState === 'unanswered' && inputRef) {
      inputRef.focus();
    }
  });

  onMount(() => {
    if (inputRef) {
      inputRef.focus();
    }
  });

  const displayPronoun = () => {
    if (!props.pronoun) return '';
    return getElidedPronoun(props.pronoun, props.value);
  };

  return (
    <div class="space-y-4">
      {/* Unified pronoun + input row */}
      <div class="flex items-stretch overflow-hidden rounded-[var(--radius-lg)] border border-border bg-background transition-colors focus-within:border-foreground/30">
        {props.pronoun && (
          <div class="flex items-center border-r border-border bg-muted px-5 text-lg text-foreground" aria-hidden="true">
            {displayPronoun()}
          </div>
        )}
        <input
          ref={inputRef}
          type="text"
          value={props.value}
          onInput={(e) => {
            const value = e.currentTarget.value;
            // Limit to 50 characters (reasonable max for French verb conjugations)
            if (value.length <= 50) {
              props.onInput(value);
            }
          }}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && !props.disabled) {
              if (props.answerState !== 'unanswered' && props.onReset) {
                props.onReset();
              } else {
                props.onSubmit();
              }
            }
          }}
          disabled={props.disabled}
          placeholder="Type the conjugated verb..."
          class="flex-1 bg-transparent px-5 py-4 text-lg text-foreground outline-none placeholder:text-muted-foreground"
          maxLength={50}
          aria-label={`Conjugate ${props.pronoun || 'verb'} in the current tense`}
        />
      </div>

      <Show when={props.answerState === 'correct'}>
        <div class="rounded-[var(--radius-md)] border border-green-500 bg-green-50 px-4 py-3 dark:bg-green-900/20" role="alert" aria-live="polite">
          <p class="font-medium text-green-700 dark:text-green-400">Correct!</p>
        </div>
      </Show>

      <Show when={props.answerState === 'incorrect'}>
        <div class="rounded-[var(--radius-md)] border border-destructive bg-destructive/10 px-4 py-3" role="alert" aria-live="polite">
          <p class="font-medium text-destructive-foreground">Incorrect</p>
          <Show when={props.correctAnswer}>
            <p class="mt-1 text-sm text-muted-foreground">
              Correct answer: <span class="font-medium text-foreground">{props.correctAnswer}</span>
            </p>
          </Show>
        </div>
      </Show>

      {/* Submit / Next button — right-aligned */}
      <div class="flex justify-end">
        <button
          type="button"
          onClick={() => {
            if (props.answerState !== 'unanswered' && props.onReset) {
              props.onReset();
            } else {
              props.onSubmit();
            }
          }}
          disabled={props.disabled || (props.answerState === 'unanswered' && !props.value.trim())}
          class="inline-flex h-10 items-center rounded-[var(--radius)] bg-primary px-8 text-sm font-medium text-primary-foreground transition-colors hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-50"
        >
          {props.answerState !== 'unanswered' ? 'Next' : 'Submit'}
        </button>
      </div>
    </div>
  );
};

export default AnswerInput;