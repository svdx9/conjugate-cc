import { Nav } from "../components/Nav";

export function HomePage() {
  return (
    <div class="flex flex-col min-h-screen bg-bg font-sans text-ink">
      <Nav />
      <main class="flex-1 flex items-center justify-center px-4 py-12 sm:py-20">
        <div class="max-w-[36rem] w-full text-center animate-rise">
          <p class="text-eyebrow font-bold uppercase tracking-eyebrow text-muted mb-2">
            Conjugation drill application
          </p>
          <h1 class="font-serif text-hero sm:text-hero-lg tracking-hero-title text-ink mt-1.5 mb-4 leading-tight">
            conjugate.cc
          </h1>
          <p class="text-base text-muted max-w-sm mx-auto mb-8">
            Master verb conjugations through focused, repeating drills. Build
            confidence and fluency one verb at a time.
          </p>
          <a
            href="/drills"
            class="font-sans font-medium text-base text-sage no-underline transition-transform duration-200 ease-out hover:-translate-y-px hover:scale-[1.03]"
          >
            Start
          </a>
        </div>
      </main>
    </div>
  );
}
