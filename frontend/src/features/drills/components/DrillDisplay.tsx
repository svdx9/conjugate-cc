import { Component } from 'solid-js';
import { DrillItem } from '../types';

interface DrillDisplayProps {
  item: DrillItem;
}

const DrillDisplay: Component<DrillDisplayProps> = (props) => {
  return (
    <div>
      <div class="mb-3 text-sm font-medium uppercase tracking-widest text-muted-foreground">
        Conjugate
      </div>
      <div class="flex flex-wrap items-baseline gap-2 text-2xl font-normal leading-snug tracking-tight sm:text-3xl">
        <span class="font-semibold text-foreground">{props.item.prompt.infinitive}</span>
        <span class="text-muted-foreground">in the</span>
        <span class="font-semibold text-foreground">{props.item.prompt.tense}</span>
        <span class="text-muted-foreground">tense</span>
      </div>
    </div>
  );
};

export default DrillDisplay;
