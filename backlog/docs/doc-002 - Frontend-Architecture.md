---
id: doc-002
title: Frontend Architecture
type: other
created_date: '2026-04-17 09:09'
updated_date: '2026-04-17 13:32'
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

---

## Drills Feature Architecture

### 2. System Layer Diagram

```
┌─────────────────────────────────────────────┐
│              Data Source Layer              │
│                                             │
│  VerbSetProvider                            │
│  (defines which verb+tense pairs to drill)  │
└──────────────────┬──────────────────────────┘
                   │ verb+tense pairs
                   ▼
┌─────────────────────────────────────────────┐
│           Orchestration Layer               │
│                                             │
│  DrillTestBuilder ◀── DrillBehaviour        │
│  (sequencing,         (random/sequential)   │
│   scoring)                                  │
└──────────────────┬──────────────────────────┘
                   │ requests conjugation for current verb+tense
                   ▼
┌─────────────────────────────────────────────┐
│           Conjugation Engine                │
│  (pure domain — no UI, no framework)        │
│                                             │
│  verb + tense → all pronoun conjugations    │
└──────────────────┬──────────────────────────┘
                   │ DrillData
                   ▼
┌─────────────────────────────────────────────┐
│              UI Layer                       │
│                                             │
│  SingleInputDrill    VerbDrill              │
│  (1 pronoun,         (6 stacked inputs,     │
│   immediate feedback) batch submit)         │
└─────────────────────────────────────────────┘
```

### 3. UI Layer

Two drill modes that receive `DrillData` from the orchestration layer.

**SingleInputDrill** (exists — `components/SingleInputDrill.tsx`)
- Props: `verb`, `tense`, `pronoun`
- Single text input for one pronoun at a time
- Immediate correct/incorrect feedback on submit
- Uses existing `AnswerInput` component

**VerbDrill** (new — `components/VerbDrill.tsx`)
- Props: `verb`, `tense`
- Displays all 6 standard pronoun rows (je, tu, il/elle, nous, vous, ils/elles) with 6 input fields stacked vertically
- **Batch submit**: user fills all 6 fields, clicks one Submit button
- Results shown per-row after submission
- Reuses existing `AnswerInput` component (one instance per pronoun row)
- New hook `useVerbDrill` manages 6 answer signals and batch correctness state

### 4. DrillTestBuilder (Orchestration)

Not a UI component. Manages a drill session end-to-end.

**Responsibilities:**
- Holds the current verb set (received from VerbSetProvider)
- Selects the next verb+tense pair (delegated to DrillBehaviour)
- Requests conjugation data from the Conjugation Engine for the selected pair
- Tracks session-level scoring: correct/total across multiple verbs, per-verb results
- Signals "drill complete" when the set is exhausted (sequential mode) or after N rounds

**Key types:**
- `SessionScore`: `{ correct: number; total: number; results: VerbDrillResult[] }`
- `VerbDrillResult`: `{ verb: string; tense: Tense; perPronoun: { pronoun: Pronoun; correct: boolean }[] }`

### 5. DrillBehaviour (Selection Strategy)

A strategy object consumed by DrillTestBuilder. Determines how the next verb+tense pair is selected from the set.

- **`random`** — picks randomly from the set
- **`sequential`** — iterates in order, index 0 to N; returns null when exhausted

Signature: `(setSize: number, state: BehaviourState) => number | null`

### 6. VerbSetProvider (Data Source)

Defines **what** verbs/tenses are available for drilling. Separate from the Conjugation Engine (which knows **how** to conjugate).

**Interface:**
- Returns an array of `{ verb: string; tense: Tense }` pairs

**Implementations:**
- `StubVerbSetProvider` — hardcoded array (for development)
- Future: `ApiVerbSetProvider` — fetches user-created verb sets from the backend (e.g. "reflexive verbs in futur proche")

### 7. Conjugation Engine

Pure domain module. No UI imports, no SolidJS, no framework code. Given a verb infinitive and a tense, produces all pronoun conjugations.

**Replaces:** the hardcoded `verbData` map in current `StubDrillProvider`

**Must handle:**
- All tenses (présent, imparfait, passé composé, futur, and future additions)
- Reflexive verbs (reflexive pronoun changes with subject, elision before vowels)
- Compound tenses (auxiliary selection: avoir vs être; past participle agreement)
- Stem-changing verbs (e.g. appeler: j'appelle vs nous appelons)
- Irregular verbs (être, avoir, aller, faire, etc. — stored as exception data)
- Negation (ne...pas wrapping, correct placement with compound tenses and reflexives)
- **Multiple valid conjugations** (e.g. s'asseoir: "je m'assieds" / "je m'assois")

**Multiple valid forms:** handled via `ExpectedAnswer.alternates?: string[]`. The canonical answer goes in `text`, alternatives in `alternates`. Answer-checking accepts any valid form.

**Pronoun synthesis:** continues the existing pattern from `synthesizePronouns` — derives elle/on from il and elles from ils, since most conjugation references list only 6 canonical forms.

### 8. What Stays / Changes / Is New

**Stays (unchanged):**
- `types.ts` — all existing types remain (additive change only: `alternates` field)
- `shared/types.ts` — `Result<T>` type
- `components/AnswerInput.tsx` — reusable for both drill modes
- `components/DrillDisplay.tsx` — reusable for verb+tense header display
- `components/SingleInputDrill.tsx` — existing component, works as-is
- `DrillProvider` interface in `provider.ts`

**Changes:**
- `provider.ts` — `StubDrillProvider` moves to a dev/test role; the Conjugation Engine becomes the production implementation of `DrillProvider`
- `hooks/useDrill.ts` — accepts a `DrillProvider` via parameter or SolidJS context instead of importing the singleton directly
- `DrillsPage.tsx` — wires up VerbSetProvider + DrillTestBuilder + chosen drill mode

**New:**
- `components/VerbDrill.tsx` — 6-input batch-submit drill component
- `hooks/useVerbDrill.ts` — hook for VerbDrill (6 answer signals, batch submission)
- `verb-set-provider.ts` — VerbSetProvider interface + StubVerbSetProvider
- `drill-behaviour.ts` — random and sequential behaviour functions
- `drill-test-builder.ts` — session orchestrator
- `engine/` directory — conjugation engine (pure domain module)

**Type addition:**
- `ExpectedAnswer.alternates?: string[]` — for verbs with multiple valid conjugations
- Answer-checking logic extended to check against both `text` and `alternates`
