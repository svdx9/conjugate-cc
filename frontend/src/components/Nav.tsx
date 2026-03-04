export function Nav() {
  return (
    <nav class="border-b border-line px-4 py-3 sm:px-8 sm:py-4" aria-label="Main navigation">
      <ul class="flex flex-wrap items-center gap-x-6 gap-y-2 list-none m-0 p-0">
        <li>
          <a
            href="/"
            class="font-sans font-bold text-base text-ink no-underline hover:text-sage focus-visible:ring-2 focus-visible:ring-sage focus-visible:ring-offset-1 focus-visible:ring-offset-bg transition-colors"
          >
            conjugate.cc
          </a>
        </li>
        <li>
          <a
            href="/drills"
            class="font-sans text-sm text-muted no-underline hover:text-ink focus-visible:ring-2 focus-visible:ring-sage focus-visible:ring-offset-1 focus-visible:ring-offset-bg transition-colors"
          >
            Drills
          </a>
        </li>
        <li>
          <a
            href="/verbs"
            class="font-sans text-sm text-muted no-underline hover:text-ink focus-visible:ring-2 focus-visible:ring-sage focus-visible:ring-offset-1 focus-visible:ring-offset-bg transition-colors"
          >
            Verbs
          </a>
        </li>
      </ul>
    </nav>
  );
}
