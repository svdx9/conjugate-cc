import { Component, JSX } from 'solid-js';

const PageShell: Component<{ children?: JSX.Element }> = (props) => {
  return (
    <div class="pt-16 sm:pt-20 lg:pt-24">
      <div class="mx-auto max-w-6xl px-4 py-12 sm:px-6 sm:py-16 lg:px-8 lg:py-20">
        {props.children}
      </div>
    </div>
  );
};

export default PageShell;
