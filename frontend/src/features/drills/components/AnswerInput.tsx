import { Component, Show, createEffect, onMount } from 'solid-js';

const VOWELS = 'aeiouàâäéèêëïîôùûüAEIOUÀÂÄÉÈÊËÏÎÔÙÛÜ';
const VOWEL_LIKE_CONSONANTS = 'hH';
const ELISION_WORDS = ['y', 'Y'];

const canElide = (userInput: string): boolean => {
  const trimmed = userInput.trim();
  if (!trimmed) return false;
  const first = trimmed[0];
  if (VOWELS.includes(first) || VOWEL_LIKE_CONSONANTS.includes(first)) {
    return true;
  }
  const firstWord = trimmed.split(/\s+/)[0];
  return ELISION_WORDS.includes(firstWord);
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
      <div
        class="flex items-stretch overflow-hidden rounded-[var(--radius-lg)] border bg-background transition-all"
        classList={{
          'border-border focus-within:border-foreground/30': props.answerState === 'unanswered',
          'border-success/30 bg-success/5 shadow-sm': props.answerState === 'correct',
          'border-red-500/30 bg-red-400/5 animate-shake shadow-sm': props.answerState === 'incorrect',
        }}
      >
        {props.pronoun && (
          <div
            class="flex items-center border-r px-6 text-lg transition-colors"
            classList={{
              'border-border bg-muted text-foreground': props.answerState === 'unanswered' || !props.answerState,
              'border-success/20 bg-success/[0.08] text-success': props.answerState === 'correct',
              'border-red-500/20 bg-red-500/[0.08] text-foreground': props.answerState === 'incorrect',
            }}
            aria-hidden="true"
          >
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
          class="flex-1 bg-transparent px-6 py-5 text-lg text-foreground outline-none placeholder:text-muted-foreground"
          maxLength={50}
          aria-label={`Conjugate ${props.pronoun || 'verb'} in the current tense`}
        />
      </div>

      <Show when={props.answerState === 'correct'}>
        <div class="rounded-[var(--radius-md)] border border-success bg-success/10 px-4 py-3 dark:bg-success-foreground/10" role="alert" aria-live="polite">
          <p class="font-medium text-success dark:text-success-foreground">Correct!</p>
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
