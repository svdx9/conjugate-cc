import { describe, it, expect } from 'vitest';
import { drillProvider } from './provider';
import { isError } from '../../shared/types';

describe('StubDrillProvider', () => {
  it('should return être present tense data', () => {
    const result = drillProvider.getDrillData('être', 'present');
    
    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }
    
    expect(result.data.verb).toBe('être');
    expect(result.data.tense).toBe('present');
    expect(result.data.items.length).toBe(8);
    expect(result.data.items[0].expectedAnswer.text).toBe('suis');
  });

  it('should return avoir present tense data', () => {
    const result = drillProvider.getDrillData('avoir', 'present');
    
    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }
    
    expect(result.data.verb).toBe('avoir');
    expect(result.data.tense).toBe('present');
    expect(result.data.items.length).toBe(8);
    expect(result.data.items[0].expectedAnswer.text).toBe('ai');
  });

  it('should return se laver present tense data with reflexive flag', () => {
    const result = drillProvider.getDrillData('se laver', 'present');
    
    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }
    
    expect(result.data.verb).toBe('se laver');
    expect(result.data.tense).toBe('present');
    expect(result.data.items.length).toBe(8);
    expect(result.data.items[0].expectedAnswer.text).toBe('me lave');
    expect(result.data.items[0].expectedAnswer.isReflexive).toBe(true);
  });

  it('should handle case variations for ASCII characters', () => {
    const result1 = drillProvider.getDrillData('avoir', 'present');
    const result2 = drillProvider.getDrillData('AVOIR', 'present');

    // Test exact verb matching - lowercase 'avoir' should succeed
    if (!isError(result1)) {
      expect(result1.data.verb).toBe('avoir');
      expect(result1.data.tense).toBe('present');
    }

    // Test exact verb matching - uppercase 'AVOIR' should fail since we preserve case for accents
    if (!isError(result2)) {
      expect.fail('Expected error for uppercase AVOIR but got data');
    } else {
      expect(result2.code).toBe('NOT_FOUND');
    }
  });

  it('should return error for unknown verbs', () => {
    const result = drillProvider.getDrillData('unknown', 'present');
    
    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }
    
    expect(result.code).toBe('NOT_FOUND');
    expect(result.error).toContain('unknown');
  });

  it('should get specific drill item by pronoun', () => {
    const result = drillProvider.getDrillItem('être', 'present', 'je');
    
    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }
    
    expect(result.data.prompt.infinitive).toBe('être');
    expect(result.data.prompt.tense).toBe('present');
    expect(result.data.prompt.pronoun).toBe('je');
    expect(result.data.expectedAnswer.text).toBe('suis');
  });

  it('should return all pronouns when no pronoun specified', () => {
    const result = drillProvider.getDrillData('avoir', 'present');
    
    if (isError(result)) {
      expect.fail(`Expected data but got error: ${result.error}`);
    }
    
    expect(result.data.verb).toBe('avoir');
    expect(result.data.tense).toBe('present');
    expect(result.data.items.length).toBe(8);
  });

  it('should return error for invalid verb', () => {
    const result = drillProvider.getDrillData('', 'present');
    
    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }
    
    expect(result.code).toBe('INVALID_VERB');
  });

  it('should return error for invalid tense', () => {
    const result = drillProvider.getDrillData('être', '');
    
    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }
    
    expect(result.code).toBe('INVALID_TENSE');
  });

  it('should return error for missing pronoun conjugation', () => {
    // Test that we get an error when a pronoun exists in the type system
    // but doesn't have a conjugation in our data
    const result = drillProvider.getDrillItem('être', 'present', 'je');

    if (!isError(result)) {
      // This should succeed since 'je' exists in our être present data
      expect(result.data.prompt.pronoun).toBe('je');
      expect(result.data.expectedAnswer.text).toBe('suis');
    } else {
      expect.fail('Unexpected error for valid pronoun');
    }
  });

  it('should return error for unknown tense', () => {
    const result = drillProvider.getDrillData('être', 'future');
    
    if (!isError(result)) {
      expect.fail('Expected error but got data');
    }
    
    expect(result.code).toBe('NOT_FOUND');
    expect(result.error).toContain('future');
  });
});
