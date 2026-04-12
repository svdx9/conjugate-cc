// Drill data provider interface and stub implementation
// This DrillProvider interface defines the contract that both the stub implementation
// and the future conjugation engine must implement

import { DrillData, DrillItem, Pronoun, Tense } from './types';
import { Result, error, success } from '../../shared/types';

export const validPronouns: Pronoun[] = ['je', 'tu', 'il', 'elle', 'on', 'nous', 'vous', 'ils', 'elles'];

// List of valid tenses - not an exclusive list, future tense support may expand
const validTenses: Tense[] = ['présent', 'imparfait', 'passé_composé', 'futur'];

export interface DrillProvider {
  /**
   * Get drill data for a specific verb and tense
   * @param verb - The infinitive form of the verb (e.g., "être", "avoir")
   * @param tense - The tense to conjugate in (e.g., "présent", "imparfait")
   * @returns Result<DrillData> with either data or error
   */
  getDrillData(verb: string, tense: Tense): Result<DrillData>;

  /**
   * Get a specific drill item for a verb, tense, and pronoun
   * @param verb - The infinitive form of the verb (e.g., "être", "avoir")
   * @param tense - The tense to conjugate in (e.g., "présent", "imparfait")
   * @param pronoun - Pronoun to get specific conjugation (e.g., "je", "tu")
   * @returns Result<DrillItem> with either data or error
   */
  getDrillItem(verb: string, tense: Tense, pronoun: Pronoun): Result<DrillItem>;
}

// Verb conjugation data stored as JSON object
const verbData: Record<string, Record<string, DrillData>> = {
  être: {
    présent: {
      verb: 'être',
      tense: 'présent' as Tense,
      items: [
        {
          prompt: { infinitive: 'être', tense: 'présent' as Tense, pronoun: 'je' },
          expectedAnswer: { text: 'suis' },
        },
        {
          prompt: { infinitive: 'être', tense: 'présent' as Tense, pronoun: 'tu' },
          expectedAnswer: { text: 'es' },
        },
        {
          prompt: { infinitive: 'être', tense: 'présent' as Tense, pronoun: 'il' },
          expectedAnswer: { text: 'est' },
        },
        {
          prompt: { infinitive: 'être', tense: 'présent' as Tense, pronoun: 'on' },
          expectedAnswer: { text: 'est' },
        },
        {
          prompt: { infinitive: 'être', tense: 'présent' as Tense, pronoun: 'nous' },
          expectedAnswer: { text: 'sommes' },
        },
        {
          prompt: { infinitive: 'être', tense: 'présent' as Tense, pronoun: 'vous' },
          expectedAnswer: { text: 'êtes' },
        },
        {
          prompt: { infinitive: 'être', tense: 'présent' as Tense, pronoun: 'ils' },
          expectedAnswer: { text: 'sont' },
        },
      ],
    },
  },
  avoir: {
    présent: {
      verb: 'avoir',
      tense: 'présent' as Tense,
      items: [
        {
          prompt: { infinitive: 'avoir', tense: 'présent' as Tense, pronoun: 'je' },
          expectedAnswer: { text: 'ai' },
        },
        {
          prompt: { infinitive: 'avoir', tense: 'présent' as Tense, pronoun: 'tu' },
          expectedAnswer: { text: 'as' },
        },
        {
          prompt: { infinitive: 'avoir', tense: 'présent' as Tense, pronoun: 'il' },
          expectedAnswer: { text: 'a' },
        },
        {
          prompt: { infinitive: 'avoir', tense: 'présent' as Tense, pronoun: 'nous' },
          expectedAnswer: { text: 'avons' },
        },
        {
          prompt: { infinitive: 'avoir', tense: 'présent' as Tense, pronoun: 'vous' },
          expectedAnswer: { text: 'avez' },
        },
        {
          prompt: { infinitive: 'avoir', tense: 'présent' as Tense, pronoun: 'ils' },
          expectedAnswer: { text: 'ont' },
        },
      ],
    },
  },
  'se laver': {
    présent: {
      verb: 'se laver',
      tense: 'présent' as Tense,
      items: [
        {
          prompt: { infinitive: 'se laver', tense: 'présent' as Tense, pronoun: 'je' },
          expectedAnswer: { text: 'me lave', isReflexive: true },
        },
        {
          prompt: { infinitive: 'se laver', tense: 'présent' as Tense, pronoun: 'tu' },
          expectedAnswer: { text: 'te laves', isReflexive: true },
        },
        {
          prompt: { infinitive: 'se laver', tense: 'présent' as Tense, pronoun: 'il' },
          expectedAnswer: { text: 'se lave', isReflexive: true },
        },
        {
          prompt: { infinitive: 'se laver', tense: 'présent' as Tense, pronoun: 'nous' },
          expectedAnswer: { text: 'nous lavons', isReflexive: true },
        },
        {
          prompt: { infinitive: 'se laver', tense: 'présent' as Tense, pronoun: 'vous' },
          expectedAnswer: { text: 'vous lavez', isReflexive: true },
        },
        {
          prompt: { infinitive: 'se laver', tense: 'présent' as Tense, pronoun: 'ils' },
          expectedAnswer: { text: 'se lavent', isReflexive: true },
        },
      ],
    },
  },
};

// Stub implementation with hardcoded conjugation data
// This will be replaced by a real conjugation engine in the future
class StubDrillProvider implements DrillProvider {
  /**
   * Validates verb and tense inputs.
   * Returns Result<{normalizedVerb: string, normalizedTense: Tense}> on success,
   * or error if validation fails.
   */
  private validateInputs(
    verb: string,
    tense: Tense,
  ): Result<{ normalizedVerb: string; normalizedTense: Tense }> {
    // Validate inputs
    if (!verb || typeof verb !== 'string' || verb.trim() === '') {
      return error('Invalid verb: verb must be a non-empty string', 'INVALID_VERB');
    }

    // Normalize verb for lookup
    // toLowerCase() correctly handles accented characters (e.g., "ÊTRE" → "être")
    const normalizedVerb = verb.trim().toLowerCase();

    return success({ normalizedVerb, normalizedTense: tense });
  }

  /**
   * Validates tense input.
   * @param tense - The tense to validate
   * @returns Result<Tense> with the tense or error
   */
  private validateTense(tense: string): Result<Tense> {
    if (!validTenses.includes(tense as Tense)) {
      return error(`Invalid tense: "${tense}" is not a valid tense`, 'INVALID_TENSE');
    }
    return success(tense as Tense);
  }

  getDrillData(verb: string, tense: Tense): Result<DrillData> {
    // Validate tense at runtime
    const tenseValidation = this.validateTense(tense);
    if (!tenseValidation.ok) {
      return tenseValidation;
    }

    const validationResult = this.validateInputs(verb, tense);
    if (!validationResult.ok) {
      return validationResult;
    }

    const { normalizedVerb, normalizedTense } = validationResult.data;

    // Check if verb exists in our data
    if (!verbData[normalizedVerb]) {
      return error(`Verb "${normalizedVerb}" not found in conjugation data`, 'NOT_FOUND', {
        availableVerbs: Object.keys(verbData),
      });
    }

    // Check if tense exists for this verb
    if (!verbData[normalizedVerb][normalizedTense]) {
      return error(
        `Tense "${normalizedTense}" not found for verb "${normalizedVerb}"`,
        'NOT_FOUND',
        { availableTenses: Object.keys(verbData[normalizedVerb]) },
      );
    }

    const baseData = verbData[normalizedVerb][normalizedTense];

    const basePronouns = new Set(baseData.items.map((item) => item.prompt.pronoun));

    const synthesizedItems = [...baseData.items];

    for (const item of baseData.items) {
      if (item.prompt.pronoun === 'il') {
        if (!basePronouns.has('elle')) {
          synthesizedItems.push({
            prompt: { ...item.prompt, pronoun: 'elle' as Pronoun },
            expectedAnswer: { ...item.expectedAnswer },
          });
        }
        if (!basePronouns.has('on')) {
          synthesizedItems.push({
            prompt: { ...item.prompt, pronoun: 'on' as Pronoun },
            expectedAnswer: { ...item.expectedAnswer },
          });
        }
      } else if (item.prompt.pronoun === 'ils') {
        if (!basePronouns.has('elles')) {
          synthesizedItems.push({
            prompt: { ...item.prompt, pronoun: 'elles' as Pronoun },
            expectedAnswer: { ...item.expectedAnswer },
          });
        }
      }
    }

    return success({
      ...baseData,
      items: synthesizedItems,
    });
  }

  /**
   * Validates pronoun input and normalizes it
   * @param pronoun - The pronoun to validate
   * @returns Result<Pronoun> with normalized pronoun or error
   */
  private validatePronoun(pronoun: Pronoun): Result<Pronoun> {
    const normalizedPronoun = pronoun.toLowerCase().trim();

    // Validate pronoun at runtime to ensure type safety
    if (!validPronouns.includes(normalizedPronoun as Pronoun)) {
      return error(
        `Invalid pronoun: "${normalizedPronoun}" is not a valid pronoun`,
        'INVALID_INPUT',
      );
    }

    return success(normalizedPronoun as Pronoun);
  }

  getDrillItem(verb: string, tense: Tense, pronoun: Pronoun): Result<DrillItem> {
    const drillDataResult = this.getDrillData(verb, tense);
    if (!drillDataResult.ok) {
      return drillDataResult;
    }

    const pronounValidation = this.validatePronoun(pronoun);
    if (!pronounValidation.ok) {
      return pronounValidation;
    }

    const normalizedPronoun = pronounValidation.data;
    const item = drillDataResult.data.items.find(
      (i) => i.prompt.pronoun === normalizedPronoun,
    );

    if (item) {
      return success(item);
    }

    return error(
      `No conjugation found for verb "${drillDataResult.data.verb}" in tense "${drillDataResult.data.tense}" with pronoun "${normalizedPronoun}"`,
      'NOT_FOUND',
    );
  }
}

// Export singleton instance
// To switch to a real conjugation engine, replace StubDrillProvider with the engine implementation
export const drillProvider: DrillProvider = new StubDrillProvider();
