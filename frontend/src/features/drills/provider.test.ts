import { describe, it, expect } from 'vitest';
import { drillProvider, validPronouns } from './provider';
import { isError } from '../../shared/types';

const numPronouns = validPronouns.length;

describe('StubDrillProvider', () => {
  it('should return être present tense data', () => {
    const result = drillProvider.getDrillData('être', 'présent');

    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }

    expect(result.data.verb).toBe('être');
    expect(result.data.tense).toBe('présent');
    expect(result.data.items.length).toBe(numPronouns);
    expect(result.data.items[0].expectedAnswer.text).toBe('suis');
  });

  it('should return avoir present tense data', () => {
    const result = drillProvider.getDrillData('avoir', 'présent');

    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }

    expect(result.data.verb).toBe('avoir');
    expect(result.data.tense).toBe('présent');
    expect(result.data.items.length).toBe(numPronouns);
    expect(result.data.items[0].expectedAnswer.text).toBe('ai');
  });

  it('should return se laver present tense data with reflexive flag', () => {
    const result = drillProvider.getDrillData('se laver', 'présent');

    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }

    expect(result.data.verb).toBe('se laver');
    expect(result.data.tense).toBe('présent');
    expect(result.data.items.length).toBe(numPronouns);
    expect(result.data.items[0].expectedAnswer.text).toBe('me lave');
    expect(result.data.items[0].expectedAnswer.isReflexive).toBe(true);
  });

  it('should handle case variations for ASCII characters', () => {
    const result1 = drillProvider.getDrillData('avoir', 'présent');
    const result2 = drillProvider.getDrillData('AVOIR', 'présent');

    // Test exact verb matching - lowercase 'avoir' should succeed
    if (!isError(result1)) {
      expect(result1.data.verb).toBe('avoir');
      expect(result1.data.tense).toBe('présent');
    }

    // Test case variations - 'AVOIR' (uppercase) should now succeed
    // since toLowerCase() correctly handles all characters including ASCII
    if (!isError(result2)) {
      expect(result2.data.verb).toBe('avoir');
    } else {
      expect.fail('Expected success for uppercase AVOIR but got error');
    }
  });

  it('should handle case variations including accented characters', () => {
    const result = drillProvider.getDrillData('ÊTRE', 'présent');

    if (!isError(result)) {
      expect(result.data.verb).toBe('être');
    } else {
      expect.fail('Expected success for uppercase ÊTRE but got error');
    }
  });

  it('should return error for unknown verbs', () => {
    const result = drillProvider.getDrillData('unknown', 'présent');

    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }

    expect(result.code).toBe('NOT_FOUND');
    expect(result.error).toContain('unknown');
  });

  it('should get specific drill item by pronoun', () => {
    const result = drillProvider.getDrillItem('être', 'présent', 'je');

    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }

    expect(result.data.prompt.infinitive).toBe('être');
    expect(result.data.prompt.tense).toBe('présent');
    expect(result.data.prompt.pronoun).toBe('je');
    expect(result.data.expectedAnswer.text).toBe('suis');
  });

  it('should return all pronouns when no pronoun specified', () => {
    const result = drillProvider.getDrillData('avoir', 'présent');

    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }

    expect(result.data.verb).toBe('avoir');
    expect(result.data.tense).toBe('présent');
    expect(result.data.items.length).toBe(numPronouns);
  });

  it('should return error for invalid verb', () => {
    const result = drillProvider.getDrillData('', 'présent');

    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }

    expect(result.code).toBe('INVALID_VERB');
  });

  it('should return error for invalid tense', () => {
    const result = drillProvider.getDrillData('être', 'invalide');

    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }

    expect(result.code).toBe('INVALID_TENSE');
  });

  it('should return item for valid pronoun', () => {
    // Test that we get an item when a valid pronoun exists in our data
    const result = drillProvider.getDrillItem('être', 'présent', 'je');

    if (!isError(result)) {
      // This should succeed since 'je' exists in our être present data
      expect(result.data.prompt.pronoun).toBe('je');
      expect(result.data.expectedAnswer.text).toBe('suis');
    } else {
      expect.fail('Unexpected error for valid pronoun');
    }
  });

  it('should return error for unknown tense', () => {
    const result = drillProvider.getDrillData('être', 'futur');

    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }

    expect(result.code).toBe('NOT_FOUND');
    expect(result.error).toContain('futur');
  });
});
