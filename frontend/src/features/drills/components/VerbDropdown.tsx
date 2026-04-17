import { Component, createSignal, Show, For, onMount, onCleanup } from 'solid-js';

interface VerbDropdownProps {
  value: string;
  options: string[];
  onChange: (value: string) => void;
}

const VerbDropdown: Component<VerbDropdownProps> = (props) => {
  const [isOpen, setIsOpen] = createSignal(false);
  const [highlightedIndex, setHighlightedIndex] = createSignal(0);
  let containerRef: HTMLDivElement | undefined;

  const currentIndex = () => {
    const idx = props.options.indexOf(props.value);
    return idx >= 0 ? idx : 0;
  };

  const handleKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'Escape') {
      setIsOpen(false);
      return;
    }

    if (!isOpen()) {
      if (e.key === 'Enter' || e.key === ' ' || e.key === 'ArrowDown') {
        setIsOpen(true);
        setHighlightedIndex(currentIndex());
        e.preventDefault();
      }
      return;
    }

    switch (e.key) {
      case 'ArrowDown':
        e.preventDefault();
        setHighlightedIndex((i) => (i + 1) % props.options.length);
        break;
      case 'ArrowUp':
        e.preventDefault();
        setHighlightedIndex((i) => (i - 1 + props.options.length) % props.options.length);
        break;
      case 'Enter':
      case ' ':
        e.preventDefault();
        props.onChange(props.options[highlightedIndex()]);
        setIsOpen(false);
        break;
    }
  };

  const handleClickOutside = (e: MouseEvent) => {
    if (containerRef && !containerRef.contains(e.target as Node)) {
      setIsOpen(false);
    }
  };

  const handleOptionClick = (option: string) => {
    props.onChange(option);
    setIsOpen(false);
  };

  onMount(() => {
    document.addEventListener('click', handleClickOutside);
  });

  onCleanup(() => {
    document.removeEventListener('click', handleClickOutside);
  });

  return (
    <div ref={containerRef} class="relative inline-block">
      <button
        type="button"
        onClick={() => setIsOpen(!isOpen())}
        onKeyDown={handleKeyDown}
        class="border-muted-foreground text-foreground flex items-center gap-1 rounded-none border-b border-dotted px-1 py-0.5 font-bold hover:border-solid focus:outline-none"
        aria-haspopup="listbox"
        aria-expanded={isOpen()}
        aria-activedescendant={isOpen() ? `option-${highlightedIndex()}` : undefined}
      >
        <span>{props.value}</span>
        <span class="text-muted-foreground text-xs">▼</span>
      </button>

      <Show when={isOpen()}>
        <div
          class="border-border bg-card absolute top-full left-0 z-10 mt-1 min-w-32 rounded-[var(--radius-md)] border"
          role="listbox"
        >
          <For each={props.options}>
            {(option, index) => (
              <button
                type="button"
                role="option"
                id={`option-${index()}`}
                aria-selected={option === props.value}
                onClick={(e) => {
                  e.stopPropagation();
                  handleOptionClick(option);
                }}
                onMouseEnter={() => setHighlightedIndex(index())}
                classList={{
                  'w-full px-3 py-2 text-left text-foreground hover:bg-accent focus:bg-accent': true,
                  'bg-accent': index() === highlightedIndex(),
                }}
              >
                {option}
              </button>
            )}
          </For>
        </div>
      </Show>
    </div>
  );
};

export default VerbDropdown;
