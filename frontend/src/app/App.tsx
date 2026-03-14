import { Component } from 'solid-js';

const App: Component = () => {
  return (
    <div class="min-h-screen bg-gray-50 font-sans text-gray-900">
      <header class="border-b border-gray-200 bg-white px-6 py-4 shadow-sm">
        <h1 class="text-2xl font-bold text-indigo-600">Conjugate.cc</h1>
      </header>
      <main class="mx-auto max-w-4xl p-8 text-center">
        <h2 class="mb-4 text-4xl font-extrabold">Master Your Verb Conjugations</h2>
        <p class="mb-8 text-xl text-gray-600">
          The most effective way to practice and memorize verb patterns.
        </p>
        <div class="inline-block rounded-lg border border-gray-100 bg-white p-6 shadow-md">
          <p class="text-lg font-medium">Bootstrap Success!</p>
          <p class="text-gray-500">SolidJS + Tailwind CSS v4 is running.</p>
        </div>
      </main>
    </div>
  );
};

export default App;
