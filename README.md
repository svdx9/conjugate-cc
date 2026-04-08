# Conjugate.cc

A language learning application for practicing verb conjugations.

## Repository Structure

- `backend/`: Go backend service.
- `frontend/`: SolidJS frontend application.
- `docs/`: Documentation and API specifications.
- `backlog/`: Project task management.

## Prerequisites

- **Go**: v1.25.7
- **Node.js**: v20+
- **npm**: (or pnpm/yarn)
- **Make**: (GNU Make)
- **Docker**: (Optional, for database)

## Local Development

### Quick Start

Start both the backend and frontend with a single command:

```bash
make dev-all
```

This starts:
- Backend (Go) on http://localhost:8080
- Frontend (SolidJS + Vite) on http://localhost:3000

Open http://localhost:3000 in your browser. The footer shows a green indicator when connected to the backend, along with the git SHA and build time.

### Verification

Verify the backend is running:

```bash
curl http://localhost:8080/api/v1/status
# Expected: {"status":"ok"}
```

Check the frontend footer - it should show a green indicator when the backend is running.

### Backend

The backend uses a `Makefile` for common tasks. You can run these from the `backend/` directory:

- `make test`: Run backend tests.
- `make lint`: Run backend linter (installs `golangci-lint` to `backend/tools/bin`).
- `make build`: Build backend binary to `backend/bin/server`.
- `make debug-build`: Build backend binary with debug symbols.
- `make format`: Format backend code (installs `gofumpt` to `backend/tools/bin`).

For more details, see [backend/README.md](backend/README.md).

### Frontend

The frontend is a SolidJS application using Vite.

```bash
cd frontend
npm install
npm run dev
```

For more details, see [frontend/README.md](frontend/README.md).
