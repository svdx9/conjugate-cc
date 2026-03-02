# conjugate-cc

`conjugate-cc` is being built as a Go backend plus SolidJS frontend repository. This task establishes the shared scaffold and root developer workflow so follow-on tasks can add the first runnable services without reorganizing the project.

## Repository Layout

The repository is organized around separate backend and frontend work areas:

```text
.
|-- backend/
|   |-- cmd/api/
|   |-- internal/
|   |   |-- api/v1/
|   |   |-- config/
|   |   |-- http/
|   |   `-- status/
|   `-- tools/
|-- docs/schema/v1/
|-- frontend/
|   `-- src/
`-- Makefile
```

- `backend/` holds the Go service, repo-local backend tooling, and generated API code.
- `docs/schema/v1/` holds the schema-first OpenAPI contract and generation config for the backend.
- `frontend/` will hold the SolidJS application.
- `Makefile` provides the shared backend workflow from the repository root so contributors do not need to remember backend-specific command lines.

## Local Prerequisites

Install these tools locally before working on the project:

- Go for backend development.
- Node.js and `npm` for frontend development.
- `make` for the root backend workflow targets.

## Backend Workflow

Run backend commands from the repository root:

```sh
make tools-install
make generate
make format
make lint
make test
make build
make debug-build
make backend-run
```

Notes:

- `make tools-install` installs backend formatter, linter, code generator, and hot-reload binaries into `backend/tools/bin/`.
- `make generate` regenerates `backend/internal/api/v1/api.gen.go` from `docs/schema/v1/conjugate.yaml`.
- `make build`, `make debug-build`, and `make backend-run` pass the current git SHA and UTC build time into the backend entrypoint as `main.serviceGitSHA` and `main.serviceBuildTime`.
- `make backend-run` starts the API once, while `make backend-dev` runs the same backend through `air` for local hot reload.

The initial backend endpoints are:

- `GET /v1/health`
- `GET /v1/build-info`

Local backend verification example:

```sh
make backend-run
curl http://localhost:8080/v1/health
curl http://localhost:8080/v1/build-info
```

## Frontend Workflow

The frontend application lives under `frontend/`. Contributors should work from that directory:

```sh
cd frontend
npm install
npm run dev
npm run test
npm run build
```

Notes:

- `npm run dev` starts the Vite development server for the SolidJS app shell introduced in `TASK-1.3`.
- `npm run test` verifies the first render path by rendering the app shell with Vitest and `@solidjs/testing-library`.
- `npm run build` confirms the frontend bootstrap produces a production bundle.
- `TASK-1.3` bootstraps the reusable frontend shell only; `TASK-1.4` will replace the placeholder content with the MVP front page.

## Verification

The shared setup, backend bootstrap, and frontend bootstrap are considered verified when:

- `backend/`, `backend/cmd/`, `backend/internal/`, `backend/tools/`, `frontend/`, and `frontend/src/` exist in the repository.
- The root `Makefile` exposes `test`, `lint`, `build`, `debug-build`, `format`, `generate`, `backend-run`, and `backend-dev`.
- `frontend/package.json` exposes `dev`, `test`, and `build` scripts that match the documented npm workflow.
- The documented commands above match the repository layout, npm scripts, and the `Makefile` recipes.
