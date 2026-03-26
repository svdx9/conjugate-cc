import { Component, createMemo } from 'solid-js';
import { isDarkMode } from '../../app/darkMode';

const LandingPage: Component = () => {
  const textColor = createMemo(() => (isDarkMode() ? '#ffffff' : '#000000'));

  return (
    <main class="pt-16 sm:pt-20 lg:pt-24">
      <section class="mx-auto max-w-6xl px-4 py-12 sm:px-6 sm:py-16 lg:px-8 lg:py-20">
        <div class="space-y-6 text-center sm:space-y-8">
          <h1
            class="text-4xl font-bold tracking-tight transition-colors sm:text-5xl lg:text-6xl"
            style={{ color: textColor() }}
          >
            Master French Verb Conjugations
          </h1>

          <p
            class="mx-auto max-w-2xl text-lg transition-colors sm:text-xl"
            style={{ color: textColor() }}
          >
            The most effective way to practice and memorize verb patterns. Learn at your own pace with our
            intelligent drilling system.
          </p>
        </div>
      </section>
    </main>
  );
};

export default LandingPage;
