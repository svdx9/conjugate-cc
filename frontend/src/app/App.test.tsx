import { render, screen } from '@solidjs/testing-library';
import { describe, it, expect } from 'vitest';
import App from './App';

describe('App', () => {
  it('renders the application shell', () => {
    render(() => <App />);

    expect(screen.getByText('Conjugate.cc')).toBeInTheDocument();
    expect(screen.getByText('Master Your Verb Conjugations')).toBeInTheDocument();
    expect(screen.getByText('Bootstrap Success!')).toBeInTheDocument();
  });
});
