import { DrillData, DrillItem, Pronoun, Tense } from './types';
import { Result, error, success } from '../../shared/types';

export const validPronouns: Pronoun[] = [
  'je',
  'tu',
  'il',
  'elle',
  'on',
  'nous',
  'vous',
  'ils',
  'elles',
];

const validTenses: Tense[] = ['présent', 'imparfait', 'passé_composé', 'futur'];

export interface DrillProvider {
  getDrillData(verb: string, tense: Tense): Result<DrillData>;
  getDrillItem(verb: string, tense: Tense, pronoun: Pronoun): Result<DrillItem>;
  getAvailableTenses(verb: string): string[];
  getAvailableVerbs(): string[];
}

type ConjugationEntry = [Pronoun, string, Partial<{ isReflexive: boolean }>?];

function buildDrillData(verb: string, tense: Tense, entries: ConjugationEntry[]): DrillData {
  return {
    verb,
    tense,
    items: entries.map(([pronoun, text, extra]) => ({
      prompt: { infinitive: verb, tense, pronoun },
      expectedAnswer: { text, ...extra },
    })),
  };
}

const verbData: Record<string, Partial<Record<Tense, DrillData>>> = {
  être: {
    présent: buildDrillData('être', 'présent', [
      ['je', 'suis'],
      ['tu', 'es'],
      ['il', 'est'],
      ['on', 'est'],
      ['nous', 'sommes'],
      ['vous', 'êtes'],
      ['ils', 'sont'],
    ]),
  },
  avoir: {
    présent: buildDrillData('avoir', 'présent', [
      ['je', 'ai'],
      ['tu', 'as'],
      ['il', 'a'],
      ['nous', 'avons'],
      ['vous', 'avez'],
      ['ils', 'ont'],
    ]),
  },
  'se laver': {
    présent: buildDrillData('se laver', 'présent', [
      ['je', 'me lave', { isReflexive: true }],
      ['tu', 'te laves', { isReflexive: true }],
      ['il', 'se lave', { isReflexive: true }],
      ['nous', 'nous lavons', { isReflexive: true }],
      ['vous', 'vous lavez', { isReflexive: true }],
      ['ils', 'se lavent', { isReflexive: true }],
    ]),
  },
};

function validateTense(tense: string): Result<Tense> {
  if (!validTenses.includes(tense as Tense)) {
    return error(`Invalid tense: "${tense}" is not valid`, 'INVALID_TENSE');
  }
  return success(tense as Tense);
}

function validatePronoun(pronoun: string): Result<Pronoun> {
  if (!validPronouns.includes(pronoun as Pronoun)) {
    return error(`Invalid pronoun: "${pronoun}" is not valid`, 'INVALID_PRONOUN');
  }
  return success(pronoun as Pronoun);
}

function normalizeVerb(verb: string): Result<string> {
  if (!verb || !verb.trim()) {
    return error('Invalid verb: must be non-empty', 'INVALID_VERB');
  }
  return success(verb.trim().toLowerCase());
}

function synthesizePronouns(items: DrillItem[]): DrillItem[] {
  const basePronouns = new Set(items.map((item) => item.prompt.pronoun));
  const synthesized = [...items];

  const ilItem = items.find((i) => i.prompt.pronoun === 'il');
  if (ilItem) {
    if (!basePronouns.has('elle')) {
      synthesized.push({
        prompt: { ...ilItem.prompt, pronoun: 'elle' },
        expectedAnswer: { ...ilItem.expectedAnswer },
      });
    }
    if (!basePronouns.has('on')) {
      synthesized.push({
        prompt: { ...ilItem.prompt, pronoun: 'on' },
        expectedAnswer: { ...ilItem.expectedAnswer },
      });
    }
  }

  const ilsItem = items.find((i) => i.prompt.pronoun === 'ils');
  if (ilsItem && !basePronouns.has('elles')) {
    synthesized.push({
      prompt: { ...ilsItem.prompt, pronoun: 'elles' },
      expectedAnswer: { ...ilsItem.expectedAnswer },
    });
  }

  return synthesized;
}

class StubDrillProvider implements DrillProvider {
  getDrillData(verb: string, tense: Tense): Result<DrillData> {
    const tenseResult = validateTense(tense);
    if (!tenseResult.ok) return tenseResult;

    const verbResult = normalizeVerb(verb);
    if (!verbResult.ok) {
      return error(`Invalid verb: "${verb}" not valid`, 'INVALID_VERB');
    }
    const normalizedVerb = verbResult.data;

    const tenseData = verbData[normalizedVerb];
    if (!tenseData) {
      return error(`Verb "${normalizedVerb}" not found`, 'NOT_FOUND', {
        availableVerbs: Object.keys(verbData),
      });
    }

    const drillData = tenseData[tense];
    if (!drillData) {
      return error(`Tense "${tense}" not found for "${normalizedVerb}"`, 'NOT_FOUND', {
        availableTenses: Object.keys(tenseData),
      });
    }

    return success({
      ...drillData,
      items: synthesizePronouns(drillData.items),
    });
  }

  getDrillItem(verb: string, tense: Tense, pronoun: Pronoun): Result<DrillItem> {
    const pronounResult = validatePronoun(pronoun);
    if (!pronounResult.ok) return pronounResult;

    const drillResult = this.getDrillData(verb, tense);
    if (!drillResult.ok) return drillResult;

    const item = drillResult.data.items.find((i) => i.prompt.pronoun === pronounResult.data);
    if (!item) {
      return error(
        `No conjugation for "${drillResult.data.verb}" "${tense}" "${pronoun}"`,
        'NOT_FOUND',
      );
    }
    return success(item);
  }

  getAvailableTenses(verb: string): string[] {
    const verbResult = normalizeVerb(verb);
    if (!verbResult.ok) return [];
    const data = verbData[verbResult.data];
    return data ? Object.keys(data) : [];
  }

  getAvailableVerbs(): string[] {
    return Object.keys(verbData);
  }
}

export const drillProvider: DrillProvider = new StubDrillProvider();
