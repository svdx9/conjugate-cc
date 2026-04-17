import { Component, For } from 'solid-js';
import { A } from '@solidjs/router';
import { toggleDarkMode, isDarkMode } from './darkMode';

const Navigation: Component = () => {
  const navLinks = [
    { label: 'Home', href: '/' },
    { label: 'Drills', href: '/drills' },
    { label: 'Verbs', href: '/verbs' },
    { label: 'Help', href: '/help' },
  ];

  return (
    <header class="bg-background sticky top-0 z-50 transition-colors">
      <nav class="px-4 sm:px-6 lg:px-8">
        <div class="border-border flex h-16 items-center justify-between border-b px-2">
          <A href="/" class="text-foreground text-2xl font-bold transition-colors hover:opacity-80">
            conjugate.cc
          </A>

          <div class="flex items-center gap-4">
            <ul class="flex gap-0">
              <For each={navLinks}>
                {(link) => (
                  <li>
                    <A
                      href={link.href}
                      class="text-foreground hover:text-highlight inline-flex h-10 items-center px-4 text-sm font-medium transition-colors"
                    >
                      {link.label}
                    </A>
                  </li>
                )}
              </For>
            </ul>
            <button
              onClick={toggleDarkMode}
              class="text-foreground text-lg transition-colors hover:opacity-80"
              aria-label="Toggle dark mode"
            >
              {isDarkMode() ? '○' : '◐'}
            </button>
          </div>
        </div>
      </nav>
    </header>
  );
};

export default Navigation;
