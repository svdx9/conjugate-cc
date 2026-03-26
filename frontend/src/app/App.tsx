import { Component, JSX, createMemo } from 'solid-js';
import { Route } from '@solidjs/router';
import LandingPage from '../features/landing/LandingPage';
import DrillsPage from '../features/drills/DrillsPage';
import VerbsPage from '../features/verbs/VerbsPage';
import HelpPage from '../features/help/HelpPage';
import ContactPage from '../features/contact/ContactPage';
import CookiePolicyPage from '../features/legal/CookiePolicyPage';
import Navigation from './Navigation';
import Footer from './Footer';
import { isDarkMode } from './darkMode';

const Layout: Component<{ children?: JSX.Element }> = (props) => {
  const bgColor = createMemo(() => (isDarkMode() ? '#111111' : '#ffffff'));
  const textColor = createMemo(() => (isDarkMode() ? '#ffffff' : '#000000'));

  return (
    <div
      class="flex min-h-screen flex-col transition-colors"
      style={{
        'background-color': bgColor(),
        color: textColor(),
      }}
    >
      <Navigation />
      <main class="flex-1">{props.children}</main>
      <Footer />
    </div>
  );
};

const App: Component = () => {
  return (
    <>
      <Route
        path="/"
        component={() => (
          <Layout>
            <LandingPage />
          </Layout>
        )}
      />
      <Route
        path="/drills"
        component={() => (
          <Layout>
            <DrillsPage />
          </Layout>
        )}
      />
      <Route
        path="/verbs"
        component={() => (
          <Layout>
            <VerbsPage />
          </Layout>
        )}
      />
      <Route
        path="/help"
        component={() => (
          <Layout>
            <HelpPage />
          </Layout>
        )}
      />
      <Route
        path="/contact"
        component={() => (
          <Layout>
            <ContactPage />
          </Layout>
        )}
      />
      <Route
        path="/cookie-policy"
        component={() => (
          <Layout>
            <CookiePolicyPage />
          </Layout>
        )}
      />
    </>
  );
};

export default App;
