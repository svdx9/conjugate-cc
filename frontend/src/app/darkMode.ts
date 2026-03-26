import { createSignal } from 'solid-js';

const getInitialDarkMode = () => {
  if (typeof window === 'undefined') return false;
  try {
    return localStorage.getItem('darkMode') === 'true';
  } catch {
    return false;
  }
};

const [isDarkMode, setIsDarkMode] = createSignal(getInitialDarkMode());

export const toggleDarkMode = () => {
  const newValue = !isDarkMode();
  setIsDarkMode(newValue);
  try {
    localStorage.setItem('darkMode', String(newValue));
  } catch {
    // localStorage might not be available in tests
  }
};

export { isDarkMode };
