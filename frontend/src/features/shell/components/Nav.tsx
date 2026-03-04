export function Nav() {
  return (
    <nav class="nav" aria-label="Main navigation">
      <ul class="nav__list">
        <li class="nav__item">
          <a href="/" class="nav__link nav__link--brand">
            conjugate.cc
          </a>
        </li>
        <li class="nav__item">
          <a href="/drills" class="nav__link">
            Drills
          </a>
        </li>
        <li class="nav__item">
          <a href="/verbs" class="nav__link">
            Verbs
          </a>
        </li>
      </ul>
    </nav>
  );
}
