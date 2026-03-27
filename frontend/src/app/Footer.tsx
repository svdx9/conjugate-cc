import { Component } from 'solid-js';
import { A } from '@solidjs/router';

const Footer: Component = () => {
  return (
    <footer class="border-t border-border bg-surface py-8 transition-colors dark:border-border-dark dark:bg-surface-dark">
      <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
        <div class="flex justify-center gap-8">
          <A
            href="/contact"
            class="text-sm font-medium transition-colors hover:text-highlight"
          >
            Contact
          </A>
          <A
            href="/cookie-policy"
            class="text-sm font-medium transition-colors hover:text-highlight"
          >
            Cookie Policy
          </A>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
