import { Component, createMemo } from 'solid-js';
import { A } from '@solidjs/router';
import { isDarkMode } from './darkMode';

const Footer: Component = () => {
  const textColor = createMemo(() => (isDarkMode() ? '#ffffff' : '#000000'));

  return (
    <footer
      class="border-t py-8 transition-colors"
      style={{
        'background-color': isDarkMode() ? '#111111' : '#ffffff',
        'border-color': isDarkMode() ? '#ffffff26' : '#000000',
        'color-scheme': isDarkMode() ? 'dark' : 'light',
      }}
    >
      <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
        <div class="flex justify-center gap-8">
          <A
            href="/contact"
            class="text-sm font-medium transition-colors"
            style={{ color: textColor() }}
            onMouseEnter={(e) => {
              (e.currentTarget as HTMLAnchorElement).style.color = '#ebc61d';
            }}
            onMouseLeave={(e) => {
              (e.currentTarget as HTMLAnchorElement).style.color = textColor();
            }}
          >
            Contact
          </A>
          <A
            href="/cookie-policy"
            class="text-sm font-medium transition-colors"
            style={{ color: textColor() }}
            onMouseEnter={(e) => {
              (e.currentTarget as HTMLAnchorElement).style.color = '#ebc61d';
            }}
            onMouseLeave={(e) => {
              (e.currentTarget as HTMLAnchorElement).style.color = textColor();
            }}
          >
            Cookie Policy
          </A>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
