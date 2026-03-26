import { render, screen } from '@solidjs/testing-library';
import { describe, it, expect } from 'vitest';
import { Router, Route } from '@solidjs/router';
import LandingPage from './LandingPage';

import { JSX } from 'solid-js';

const TestWrapper = (props: { children: JSX.Element }) => (
  <Router base="/">
    <Route path="/" component={() => props.children} />
  </Router>
);

describe('LandingPage', () => {
  it('displays the hero section with main heading', () => {
    render(() => (
      <TestWrapper>
        <LandingPage />
      </TestWrapper>
    ));

    expect(screen.getByText(/Master French Verb/)).toBeInTheDocument();
  });

  it('includes descriptive messaging', () => {
    render(() => (
      <TestWrapper>
        <LandingPage />
      </TestWrapper>
    ));

    expect(screen.getByText(/most effective way to practice/)).toBeInTheDocument();
  });
});
