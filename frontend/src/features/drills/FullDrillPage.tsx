import { Component } from 'solid-js';
import { useParams } from '@solidjs/router';
import PageShell from '../../shared/PageShell';
import { VerbDrill } from './components';

const FullDrillPage: Component = () => {
  const params = useParams<{ verb: string; tense: string }>();
  return (
    <PageShell>
      <div class="mb-6">
        <span class="bg-foreground/5 text-muted-foreground inline-flex items-center rounded-full px-3 py-1.5 text-xs font-medium tracking-widest uppercase">
          Full Drill
        </span>
      </div>

      <h1 class="mb-16 text-5xl leading-tight font-bold tracking-tight sm:text-6xl">
        Practice French
        <br />
        verb conjugations
      </h1>

      <VerbDrill verb={decodeURIComponent(params.verb)} tense={decodeURIComponent(params.tense)} />
    </PageShell>
  );
};

export default FullDrillPage;
