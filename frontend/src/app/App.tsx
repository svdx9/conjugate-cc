import { Component, JSX } from 'solid-js';
import { Route } from '@solidjs/router';
import LandingPage from '../features/landing/LandingPage';
import DrillsPage from '../features/drills/DrillsPage';
import VerbsPage from '../features/verbs/VerbsPage';
import HelpPage from '../features/help/HelpPage';
import ContactPage from '../features/contact/ContactPage';
import CookiePolicyPage from '../features/legal/CookiePolicyPage';
import Navigation from './Navigation';
import Footer from './Footer';

const Layout: Component<{ children?: JSX.Element }> = (props) => {
  return (
    <div class="flex min-h-screen flex-col bg-surface text-text-primary transition-colors dark:bg-surface-dark dark:text-text-primary-dark">
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
