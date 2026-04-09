// Drill data provider interface and stub implementation
// This DrillProvider interface defines the contract that both the stub implementation
// and the future conjugation engine must implement

import { DrillData, DrillItem, Pronoun, Tense } from './types';

export type DrillError = {
  error: string;
  code: 'INVALID_VERB' | 'INVALID_TENSE' | 'INVALID_PRONOUN' | 'NOT_FOUND';
  details?: Record<string, unknown>;
};

export type DrillResult<T> = T | DrillError;

export interface DrillProvider {
  /**
   * Get drill data for a specific verb and tense
   * @param verb - The infinitive form of the verb (e.g., "être", "avoir")
   * @param tense - The tense to conjugate in (e.g., "present", "imparfait")
   * @returns DrillData containing all conjugations for the specified verb/tense combination, or DrillError if not found
   */
  getDrillData(verb: string, tense: string): DrillResult<DrillData>;
  
  /**
   * Get a specific drill item for a verb, tense, and pronoun
   * @param verb - The infinitive form of the verb (e.g., "être", "avoir")
   * @param tense - The tense to conjugate in (e.g., "present", "imparfait")
   * @param pronoun - Pronoun to get specific conjugation (e.g., "je", "tu")
   * @returns DrillItem for the specified verb/tense/pronoun combination, or DrillError if not found
   */
  getDrillItem(verb: string, tense: string, pronoun: Pronoun): DrillResult<DrillItem>;
}

// Verb conjugation data stored as JSON object
const verbData: Record<string, Record<string, DrillData>> = {
  être: {
    present: {
      verb: 'être',
      tense: 'present' as Tense,
      items: [
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'je' }, expectedAnswer: { text: 'suis' } },
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'tu' }, expectedAnswer: { text: 'es' } },
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'il' }, expectedAnswer: { text: 'est' } },
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'elle' }, expectedAnswer: { text: 'est' } },
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'nous' }, expectedAnswer: { text: 'sommes' } },
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'vous' }, expectedAnswer: { text: 'êtes' } },
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'ils' }, expectedAnswer: { text: 'sont' } },
        { prompt: { infinitive: 'être', tense: 'present' as Tense, pronoun: 'elles' }, expectedAnswer: { text: 'sont' } },
      ]
    }
  },
  avoir: {
    present: {
      verb: 'avoir',
      tense: 'present' as Tense,
      items: [
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'je' }, expectedAnswer: { text: 'ai' } },
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'tu' }, expectedAnswer: { text: 'as' } },
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'il' }, expectedAnswer: { text: 'a' } },
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'elle' }, expectedAnswer: { text: 'a' } },
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'nous' }, expectedAnswer: { text: 'avons' } },
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'vous' }, expectedAnswer: { text: 'avez' } },
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'ils' }, expectedAnswer: { text: 'ont' } },
        { prompt: { infinitive: 'avoir', tense: 'present' as Tense, pronoun: 'elles' }, expectedAnswer: { text: 'ont' } },
      ]
    }
  },
  'se laver': {
    present: {
      verb: 'se laver',
      tense: 'present' as Tense,
      items: [
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'je' }, expectedAnswer: { text: 'me lave', isReflexive: true } },
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'tu' }, expectedAnswer: { text: 'te laves', isReflexive: true } },
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'il' }, expectedAnswer: { text: 'se lave', isReflexive: true } },
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'elle' }, expectedAnswer: { text: 'se lave', isReflexive: true } },
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'nous' }, expectedAnswer: { text: 'nous lavons', isReflexive: true } },
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'vous' }, expectedAnswer: { text: 'vous lavez', isReflexive: true } },
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'ils' }, expectedAnswer: { text: 'se lavent', isReflexive: true } },
        { prompt: { infinitive: 'se laver', tense: 'present' as Tense, pronoun: 'elles' }, expectedAnswer: { text: 'se lavent', isReflexive: true } },
      ]
    }
  }
};

// Stub implementation with hardcoded conjugation data
// This will be replaced by a real conjugation engine in the future
class StubDrillProvider implements DrillProvider {
  
  getDrillData(verb: string, tense: string): DrillResult<DrillData> {
    // Validate inputs
    if (!verb || typeof verb !== 'string' || verb.trim() === '') {
      return {
        error: 'Invalid verb: verb must be a non-empty string',
        code: 'INVALID_VERB'
      };
    }
    
    if (!tense || typeof tense !== 'string' || tense.trim() === '') {
      return {
        error: 'Invalid tense: tense must be a non-empty string',
        code: 'INVALID_TENSE'
      };
    }
    
    // Normalize verb and tense for lookup
    // Note: We trim but don't lowercase to preserve accents like "être"
    const normalizedVerb = verb.trim();
    const normalizedTense = tense.toLowerCase().trim();
    
    // Check if verb exists in our data (case-sensitive for accents)
    if (!verbData[normalizedVerb]) {
      return {
        error: `Verb "${normalizedVerb}" not found in conjugation data`,
        code: 'NOT_FOUND',
        details: { availableVerbs: Object.keys(verbData) }
      };
    }
    
    // Check if tense exists for this verb
    if (!verbData[normalizedVerb][normalizedTense]) {
      return {
        error: `Tense "${normalizedTense}" not found for verb "${normalizedVerb}"`,
        code: 'NOT_FOUND',
        details: { availableTenses: Object.keys(verbData[normalizedVerb]) }
      };
    }
    
    // Return the found data
    return verbData[normalizedVerb][normalizedTense];
  }
  
  getDrillItem(verb: string, tense: string, pronoun: Pronoun): DrillResult<DrillItem> {
    // Validate inputs
    if (!verb || typeof verb !== 'string' || verb.trim() === '') {
      return {
        error: 'Invalid verb: verb must be a non-empty string',
        code: 'INVALID_VERB'
      };
    }
    
    if (!tense || typeof tense !== 'string' || tense.trim() === '') {
      return {
        error: 'Invalid tense: tense must be a non-empty string',
        code: 'INVALID_TENSE'
      };
    }
    
    // Normalize inputs
    // Note: We trim but don't lowercase verbs to preserve accents like "être"
    const normalizedVerb = verb.trim();
    const normalizedTense = tense.toLowerCase().trim();
    const normalizedPronoun = pronoun.toLowerCase().trim() as Pronoun;
    
    // Get the drill data for this verb/tense combination
    const drillDataResult = this.getDrillData(normalizedVerb, normalizedTense);
    
    // If getDrillData returned an error, propagate it
    if ('error' in drillDataResult) {
      return drillDataResult;
    }
    
    // Find the specific item for the pronoun
    const item = drillDataResult.items.find(item => item.prompt.pronoun === normalizedPronoun);
    if (item) {
      return item;
    } else {
      return {
        error: `No conjugation found for verb "${normalizedVerb}" in tense "${normalizedTense}" with pronoun "${normalizedPronoun}"`,
        code: 'NOT_FOUND'
      };
    }
  }
}

// Export singleton instance
// To switch to a real conjugation engine, replace StubDrillProvider with the engine implementation
export const drillProvider: DrillProvider = new StubDrillProvider();