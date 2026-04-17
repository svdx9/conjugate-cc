import { Component, createSignal, onMount, Show } from 'solid-js';
import PageShell from '../../shared/PageShell';
import { VerbDrill } from './components';
import { DrillData } from './types';
import { drillProvider } from './provider';

const FullDrillPage: Component = () => {
  const [drillData, setDrillData] = createSignal<DrillData | null>(null);
  const [loading, setLoading] = createSignal(true);
  const [error, setError] = createSignal<string | null>(null);

  onMount(() => {
    const result = drillProvider.getDrillData('être', 'présent');
    if (result.ok) {
      setDrillData(result.data);
    } else {
      setError(result.error);
    }
    setLoading(false);
  });

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

      <Show when={loading()}>
        <div class="text-muted-foreground py-8 text-center">Loading...</div>
      </Show>

      <Show when={error()}>
        <div class="border-destructive bg-destructive/10 rounded-[var(--radius-md)] border px-4 py-3">
          <p class="text-destructive-foreground font-medium">Error: {error()}</p>
        </div>
      </Show>

      <Show when={!loading() && !error() && drillData()}>
        <VerbDrill drillData={drillData()!} />
      </Show>
    </PageShell>
  );
};

export default FullDrillPage;
