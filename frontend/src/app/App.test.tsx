import { render, screen } from "@solidjs/testing-library";

import { App } from "./App";

describe("App", () => {
  it("renders the bootstrapped application shell", () => {
    render(() => <App />);

    expect(screen.getByRole("heading", { name: "conjugate.cc" })).toBeDefined();
    expect(
      screen.getByText("SolidJS frontend bootstrap is ready for MVP work."),
    ).toBeDefined();
    expect(screen.getByRole("button", { name: "Drills coming soon" })).toBeDefined();
  });
});
