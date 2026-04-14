import { Component } from 'solid-js';
import PageShell from '../../shared/PageShell';
import { SingleInputDrill } from './components';

const DrillsPage: Component = () => {
  return (
    <PageShell>
      <div class="mb-6">
        <span class="inline-flex items-center rounded-full bg-foreground/5 px-3 py-1.5 text-xs font-medium uppercase tracking-widest text-muted-foreground">
          Quick Drill
        </span>
      </div>

      <h1 class="mb-16 text-5xl font-bold leading-tight tracking-tight sm:text-6xl">
        Practice French
        <br />
        verb conjugations
      </h1>

      <SingleInputDrill verb="être" tense="présent" />
    </PageShell>
  );
};

export default DrillsPage;