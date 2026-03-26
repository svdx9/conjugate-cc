import { Component, createMemo } from 'solid-js';
import { isDarkMode } from '../../app/darkMode';

const DrillsPage: Component = () => {
  const textColor = createMemo(() => (isDarkMode() ? '#ffffff' : '#000000'));

  return (
    <main class="pt-16 sm:pt-20 lg:pt-24">
      <div class="mx-auto max-w-6xl px-4 py-12 sm:px-6 sm:py-16 lg:px-8 lg:py-20">
        <h1 class="text-4xl font-bold" style={{ color: textColor() }}>
          Drills
        </h1>
        <p class="mt-4 text-lg" style={{ color: textColor() }}>
          Interactive drill experience coming soon.
        </p>
      </div>
    </main>
  );
};

export default DrillsPage;
