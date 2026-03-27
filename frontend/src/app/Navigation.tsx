import { Component, For } from 'solid-js';
import { A } from '@solidjs/router';
import { isDarkMode, toggleDarkMode } from './darkMode';

const Navigation: Component = () => {
  const navLinks = [
    { label: 'Home', href: '/' },
    { label: 'Drills', href: '/drills' },
    { label: 'Verbs', href: '/verbs' },
    { label: 'Contact', href: '/contact' },
    { label: 'Help', href: '/help' },
  ];

  return (
    <header
      class="sticky top-0 z-50 transition-colors"
      style={{
        'background-color': isDarkMode() ? '#111111' : '#ffffff',
        'color-scheme': isDarkMode() ? 'dark' : 'light',
      }}
    >
      <nav class="px-4 sm:px-6 lg:px-8">
        <div
          class="flex px-2 h-16 items-center justify-between"
          style={{
            'border-bottom': 'light-dark(1px solid #00000026, 1px solid #ffffff26)',
          }}
        >

          <A
            href="/"
            class="text-2xl font-bold transition-colors"
            style={{ color: isDarkMode() ? '#ffffff' : '#000000' }}
          >
            conjugate.cc
          </A>

          <div class="flex items-center gap-4">
            <ul class="flex gap-0">
              <For each={navLinks}>
                {(link) => (
                  <li>
                    <A
                      href={link.href}
                      class="inline-flex h-10 items-center px-4 text-sm font-medium transition-colors"
                      style={{ color: isDarkMode() ? '#ffffff' : '#000000' }}
                    >
                      {link.label}
                    </A>
                  </li>
                )}
              </For>
            </ul>
            <button
              onClick={toggleDarkMode}
              class="text-lg transition-colors"
              style={{ color: isDarkMode() ? '#ffffff' : '#000000' }}
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
