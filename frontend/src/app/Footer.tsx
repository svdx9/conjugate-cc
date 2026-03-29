import { Component, Show } from 'solid-js';
import { A } from '@solidjs/router';
import { backendAvailable, gitSha, buildTime } from '../store/backend';

const Footer: Component = () => {
  return (
    <footer class="border-border bg-surface dark:border-border-dark dark:bg-surface-dark border-t py-8 transition-colors">
      <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
        <div class="flex w-full items-center justify-between">
          <div class="flex gap-8">
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
          <Show when={import.meta.env.DEV}>
            <div class="flex items-center gap-2 font-mono text-xs text-gray-500 dark:text-gray-400">
              <span
                class="h-2 w-2 rounded-full"
                classList={{
                  'bg-green-500': backendAvailable(),
                  'bg-red-500': !backendAvailable(),
                }}
              />
              <span>{backendAvailable() ? `${gitSha()} ${buildTime()}` : ''}</span>
            </div>
          </Show>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
