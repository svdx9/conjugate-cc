import { Nav } from "../components/Nav";

export function AppShell() {
  return (
    <div class="app-shell">
      <Nav />
      <main class="hero">
        <div class="hero__content">
          <p class="hero__eyebrow">Conjugation drill application</p>
          <h1 class="hero__title">conjugate.cc</h1>
          <p class="hero__description">
            Master verb conjugations through focused, repeating drills. Build
            confidence and fluency one form at a time.
          </p>
          <a href="/drills" class="hero__cta">
            Start Drilling
          </a>
        </div>
      </main>
    </div>
  );
}
