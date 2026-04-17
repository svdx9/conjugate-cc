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
  let nextButtonRef: HTMLButtonElement | undefined;

  createEffect(() => {
    if (props.answerState === 'unanswered' && inputRef) {
      inputRef.focus();
    } else if (
      (props.answerState === 'correct' || props.answerState === 'incorrect') &&
      nextButtonRef
    ) {
      setTimeout(() => nextButtonRef.focus(), 0);
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
        class="bg-background flex items-stretch overflow-hidden rounded-[var(--radius-lg)] border transition-all"
        classList={{
          'border-border focus-within:border-foreground/30': props.answerState === 'unanswered',
          'border-success/30 bg-success/5 shadow-sm': props.answerState === 'correct',
          'border-red-500/30 bg-red-400/5 animate-shake shadow-sm':
            props.answerState === 'incorrect',
        }}
      >
        {props.pronoun && (
          <div
            class="flex items-center border-r px-6 text-lg transition-colors"
            classList={{
              'border-border bg-muted text-foreground':
                props.answerState === 'unanswered' || !props.answerState,
              'border-success/20 bg-success/[0.08] text-success': props.answerState === 'correct',
              'border-red-500/20 bg-red-500/[0.08] text-foreground':
                props.answerState === 'incorrect',
            }}
            aria-hidden="true"
          >
            {displayPronoun()}
          </div>
        )}
        <Show
          when={props.answerState === 'incorrect'}
          fallback={
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
              class="text-foreground placeholder:text-muted-foreground flex-1 bg-transparent px-6 py-5 text-lg outline-none"
              maxLength={50}
              aria-label={`Conjugate ${props.pronoun || 'verb'} in the current tense`}
            />
          }
        >
          <div class="flex-1 px-6 py-5 text-lg" aria-live="polite">
            <s class="text-muted-foreground">{props.value}</s>
            <span class="text-foreground ml-2 font-medium">{props.correctAnswer}</span>
          </div>
        </Show>
      </div>

      <div class="flex justify-end">
        <button
          ref={nextButtonRef}
          type="button"
          onClick={() => {
            if (props.answerState !== 'unanswered' && props.onReset) {
              props.onReset();
            } else {
              props.onSubmit();
            }
          }}
          disabled={props.disabled || (props.answerState === 'unanswered' && !props.value.trim())}
          class="bg-primary text-primary-foreground inline-flex h-10 items-center rounded-[var(--radius)] px-8 text-sm font-medium transition-colors hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-50"
        >
          {props.answerState !== 'unanswered' ? 'Next' : 'Submit'}
        </button>
      </div>
    </div>
  );
};

export default AnswerInput;
