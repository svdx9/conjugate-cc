import { render, screen } from '@solidjs/testing-library';
import { describe, it, expect } from 'vitest';
import Navigation from './Navigation';
import LandingPage from '../features/landing/LandingPage';
import { Router, Route } from '@solidjs/router';

describe('App', () => {
  it('navigation renders the app name and links', () => {
    render(() => (
      <Router base="/">
        <Route path="/" component={() => <Navigation />} />
      </Router>
    ));

    expect(screen.getByText('conjugate.cc')).toBeInTheDocument();
    expect(screen.getByText('Home')).toBeInTheDocument();
    expect(screen.getByText('Drills')).toBeInTheDocument();
    expect(screen.getByText('Verbs')).toBeInTheDocument();
  });

  it('landing page displays hero content', () => {
    render(() => (
      <Router base="/">
        <Route path="/" component={LandingPage} />
      </Router>
    ));

    expect(screen.getByText(/Master French Verb/)).toBeInTheDocument();
    expect(screen.getByText(/most effective way to practice/)).toBeInTheDocument();
  });
});
