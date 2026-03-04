import { render, screen } from "@solidjs/testing-library";
import { App } from "./App";

describe("App", () => {
  beforeEach(() => {
    render(() => <App />);
  });

  it("renders the brand nav link", () => {
    expect(screen.getByRole("link", { name: "conjugate.cc" })).toBeInTheDocument();
  });

  it("renders Drills and Verbs nav links", () => {
    expect(screen.getByRole("link", { name: "Drills" })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Verbs" })).toBeInTheDocument();
  });

  it("renders the hero heading", () => {
    expect(screen.getByRole("heading", { level: 1 })).toHaveTextContent("conjugate.cc");
  });

  it("renders the Start Drilling CTA", () => {
    expect(screen.getByRole("link", { name: "Start" })).toBeInTheDocument();
  });
});
