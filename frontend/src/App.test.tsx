import { render, screen } from "@solidjs/testing-library";
import { App } from "./App";

describe("App", () => {
  it("renders the brand nav link", () => {
    render(() => <App />);
    expect(screen.getByRole("link", { name: "conjugate.cc" })).toBeInTheDocument();
  });

  it("renders Drills and Verbs nav links", () => {
    render(() => <App />);
    expect(screen.getByRole("link", { name: "Drills" })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Verbs" })).toBeInTheDocument();
  });

  it("renders the hero heading", () => {
    render(() => <App />);
    expect(screen.getByRole("heading", { level: 1 })).toHaveTextContent("conjugate.cc");
  });

  it("renders the Start Drilling CTA", () => {
    render(() => <App />);
    expect(screen.getByRole("link", { name: "Start Drilling" })).toBeInTheDocument();
  });
});
