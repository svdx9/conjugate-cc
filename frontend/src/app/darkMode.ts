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
  } catch {
    // localStorage might not be available in tests
  }
};

export { isDarkMode };
