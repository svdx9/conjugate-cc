import { Component, Show } from 'solid-js';
import { A } from '@solidjs/router';
import { BackendStatus } from '../hooks/useBackendStatus';

interface FooterProps {
  backend: BackendStatus;
}

const Footer: Component<FooterProps> = (props) => {
  return (
    <footer class="bg-background py-8 transition-colors">
      <div class="px-4 sm:px-6 lg:px-8">
        <div class="flex w-full items-center justify-between px-2">
          <Show when={import.meta.env.DEV}>
            <div class="text-muted-foreground flex items-center gap-2 font-mono text-xs">
              <span
                class="h-2 w-2 rounded-full"
                classList={{
                  'bg-green-500': props.backend.available(),
                  'bg-red-500': !props.backend.available(),
                }}
              />
              <span>
                {props.backend.available()
                  ? `${props.backend.gitSha()} ${props.backend.buildTime()}`
                  : ''}
              </span>
            </div>
          </Show>
          <div class="flex gap-8">
            <A
              href="/contact"
              class="text-foreground hover:text-highlight text-sm font-medium transition-colors"
            >
              Contact
            </A>
            <A
              href="/cookie-policy"
              class="text-foreground hover:text-highlight text-sm font-medium transition-colors"
            >
              Cookie Policy
            </A>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
