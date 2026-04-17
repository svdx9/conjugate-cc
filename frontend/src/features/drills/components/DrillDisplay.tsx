import { Component } from 'solid-js';
import { DrillItem } from '../types';

interface DrillDisplayProps {
  item: DrillItem;
}

const DrillDisplay: Component<DrillDisplayProps> = (props) => {
  return (
    <div>
      <div class="text-muted-foreground mb-3 text-sm font-medium tracking-widest uppercase">
        Conjugate
      </div>
      <div class="flex flex-wrap items-baseline gap-2 text-2xl leading-snug font-normal tracking-tight sm:text-3xl">
        <span class="text-foreground font-semibold">{props.item.prompt.infinitive}</span>
        <span class="text-muted-foreground">in the</span>
        <span class="text-foreground font-semibold">{props.item.prompt.tense}</span>
        <span class="text-muted-foreground">tense</span>
      </div>
    </div>
  );
};

export default DrillDisplay;
