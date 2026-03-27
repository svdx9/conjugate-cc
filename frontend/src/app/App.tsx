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
    <div class="bg-surface text-text-primary dark:bg-surface-dark dark:text-text-primary-dark flex min-h-screen flex-col transition-colors">
      <Navigation />
      <main class="flex-1">{props.children}</main>
      <Footer />
    </div>
  );
};

const App: Component = () => {
  return (
    <Route path="/" component={Layout}>
      <Route path="/" component={LandingPage} />
      <Route path="/drills" component={DrillsPage} />
      <Route path="/verbs" component={VerbsPage} />
      <Route path="/help" component={HelpPage} />
      <Route path="/contact" component={ContactPage} />
      <Route path="/cookie-policy" component={CookiePolicyPage} />
    </Route>
  );
};

export default App;
