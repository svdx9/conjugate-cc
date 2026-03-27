import { Component } from 'solid-js';
import PageShell from '../../shared/PageShell';

const LandingPage: Component = () => {
  return (
    <PageShell>
      <div class="space-y-6 text-center sm:space-y-8">
        <h1 class="text-4xl font-bold tracking-tight transition-colors sm:text-5xl lg:text-6xl">
          Master French Verb Conjugations
        </h1>

        <p class="mx-auto max-w-2xl text-lg transition-colors sm:text-xl">
          The most effective way to practice and memorize verb patterns. Learn at your own pace with
          our intelligent drilling system.
        </p>
      </div>
    </PageShell>
  );
};

export default LandingPage;
