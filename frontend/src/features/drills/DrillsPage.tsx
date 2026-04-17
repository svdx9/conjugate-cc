import { Component } from 'solid-js';
import { A } from '@solidjs/router';
import PageShell from '../../shared/PageShell';

const DrillsPage: Component = () => {
  return (
    <PageShell>
      <div class="mb-6">
        <span class="bg-foreground/5 text-muted-foreground inline-flex items-center rounded-full px-3 py-1.5 text-xs font-medium tracking-widest uppercase">
          Drills
        </span>
      </div>

      <h1 class="mb-16 text-5xl leading-tight font-bold tracking-tight sm:text-6xl">
        Practice French
        <br />
        verb conjugations
      </h1>

      <div class="grid gap-4 sm:grid-cols-2">
        <A
          href="/drills/quick"
          class="border-border bg-secondary/50 hover:bg-secondary/70 block rounded-[var(--radius-xl)] border p-6 transition-colors"
        >
          <h2 class="mb-2 text-xl font-semibold">Quick Drill</h2>
          <p class="text-muted-foreground">Single pronoun, random selection. Fast practice.</p>
        </A>

        <A
          href="/drills/full"
          class="border-border bg-secondary/50 hover:bg-secondary/70 block rounded-[var(--radius-xl)] border p-6 transition-colors"
        >
          <h2 class="mb-2 text-xl font-semibold">Full Drill</h2>
          <p class="text-muted-foreground">
            All 6 pronouns, batch submission. Comprehensive practice.
          </p>
        </A>
      </div>
    </PageShell>
  );
};

export default DrillsPage;
