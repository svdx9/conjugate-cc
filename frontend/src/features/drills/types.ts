// Drill data types for conjugation practice

export type Pronoun = 'je' | 'tu' | 'il' | 'elle' | 'nous' | 'vous' | 'ils' | 'elles';

export type Tense = 'présent' | 'imparfait' | 'passé_composé' | 'futur';

export type DrillPrompt = {
  infinitive: string;
  tense: Tense;
  pronoun: Pronoun;
};

export type ExpectedAnswer = {
  text: string;
  isReflexive?: boolean;
  isAuxiliary?: boolean;
  isParticiple?: boolean;
};

export type DrillItem = {
  prompt: DrillPrompt;
  expectedAnswer: ExpectedAnswer;
};

export type DrillData = {
  verb: string;
  tense: Tense;
  items: DrillItem[];
};
