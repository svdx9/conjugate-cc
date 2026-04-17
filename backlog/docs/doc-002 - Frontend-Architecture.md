---
id: doc-002
title: Frontend Architecture
type: other
created_date: '2026-04-17 09:09'
updated_date: '2026-04-17 09:10'
---
# Frontend Architecture

## Enforceable Design Decisions

### 1. DrillProvider API Contract

**Status**: Active  
**Applies To**: `src/features/drills/provider.ts`

The `DrillProvider` interface accepts raw `string` parameters, with all validation performed internally within the provider implementation.

#### Rationale

- **Simpler public API**: Callers pass strings without needing to import or use domain types (`Tense`, `Pronoun`)
- **Encapsulated validation**: All validation logic lives in one place (the provider), making it easier to maintain and test
- **Consistent error handling**: Errors are returned as `Result` types with consistent error codes

#### Interface Contract

```typescript
export interface DrillProvider {
  getDrillData(verb: string, tense: string): Result<DrillData>;
  getDrillItem(verb: string, tense: string, pronoun: string): Result<DrillItem>;
  getAvailableTenses(verb: string): string[];
  getAvailableVerbs(): string[];
}
```

#### Validation Rules (Internal)

| Parameter | Validation | Error Code |
|-----------|------------|------------|
| `verb` | Non-empty string, normalized to lowercase | `INVALID_VERB` |
| `tense` | Must be one of: `'présent'`, `'imparfait'`, `'passé_composé'`, `'futur'` | `INVALID_TENSE` |
| `pronoun` | Must be one of: `'je'`, `'tu'`, `'il'`, `'elle'`, `'on'`, `'nous'`, `'vous'`, `'ils'`, `'elles'` | `INVALID_PRONOUN` |
| `verb` + `tense` combination | Must exist in provider's data | `NOT_FOUND` |

#### Enforcement

- Interface signature enforces `string` parameters (not domain types)
- All implementations must return `Result<T>` with appropriate error codes
- Unit tests verify validation behavior in `provider.test.ts`
