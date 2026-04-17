import { Component, JSX } from 'solid-js';
import { Route } from '@solidjs/router';
import LandingPage from '../features/landing/LandingPage';
import DrillsPage from '../features/drills/DrillsPage';
import QuickDrillPage from '../features/drills/QuickDrillPage';
import FullDrillPage from '../features/drills/FullDrillPage';
import VerbsPage from '../features/verbs/VerbsPage';
import HelpPage from '../features/help/HelpPage';
import ContactPage from '../features/contact/ContactPage';
import CookiePolicyPage from '../features/legal/CookiePolicyPage';
import Navigation from './Navigation';
import Footer from './Footer';
import { useBackendStatus } from '../hooks/useBackendStatus';

const Layout: Component<{ children?: JSX.Element }> = (props) => {
  const backend = useBackendStatus();
  return (
    <div class="bg-background text-foreground flex min-h-screen flex-col transition-colors">
      <Navigation />
      <main class="flex-1">{props.children}</main>
      <Footer backend={backend} />
    </div>
  );
};

const App: Component = () => {
  return (
    <Route path="/" component={Layout}>
      <Route path="/" component={LandingPage} />
      <Route path="/drills" component={DrillsPage} />
      <Route path="/drills/quick" component={QuickDrillPage} />
      <Route path="/drills/full" component={FullDrillPage} />
      <Route path="/verbs" component={VerbsPage} />
      <Route path="/help" component={HelpPage} />
      <Route path="/contact" component={ContactPage} />
      <Route path="/cookie-policy" component={CookiePolicyPage} />
    </Route>
  );
};

export default App;
