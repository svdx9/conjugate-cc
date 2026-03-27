import { createSignal } from 'solid-js';

const getInitialDarkMode = () => {
  if (typeof window === 'undefined') return false;
  try {
    return localStorage.getItem('darkMode') === 'true';
  } catch (e) {
    if (import.meta.env.DEV) {
      console.warn('Failed to access localStorage for dark mode preference:', e);
    }
    return false;
  }
};

const [isDarkMode, setIsDarkMode] = createSignal(getInitialDarkMode());

const applyDarkClass = (dark: boolean) => {
  if (typeof document === 'undefined') return;
  document.documentElement.classList.toggle('dark', dark);
};

applyDarkClass(getInitialDarkMode());

export const toggleDarkMode = () => {
  const newValue = !isDarkMode();
  setIsDarkMode(newValue);
  applyDarkClass(newValue);
  try {
    localStorage.setItem('darkMode', String(newValue));
  } catch (e) {
    // localStorage might not be available in tests
    if (import.meta.env.DEV) {
      console.warn('Failed to access localStorage for dark mode preference:', e);
    }
  }
};

export { isDarkMode };
