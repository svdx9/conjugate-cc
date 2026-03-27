import { Component } from 'solid-js';
import { A } from '@solidjs/router';

const Footer: Component = () => {
  return (
    <footer class="border-border bg-surface dark:border-border-dark dark:bg-surface-dark border-t py-8 transition-colors">
      <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
        <div class="flex justify-center gap-8">
          <A href="/contact" class="hover:text-highlight text-sm font-medium transition-colors">
            Contact
          </A>
          <A
            href="/cookie-policy"
            class="hover:text-highlight text-sm font-medium transition-colors"
          >
            Cookie Policy
          </A>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
