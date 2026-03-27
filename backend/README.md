# Conjugate.cc Backend

The backend is a Go service providing the conjugation engine and API.

## Repository Layout (relative to /backend)

- `cmd/`: Application entrypoints.
- `internal/`: Private library code.
- `sql/`: SQL queries and migrations.
- `tools/`: Local development tools.

## Prerequisites

- Go v1.25
- Make

## Getting Started

### Installation

Install development tools:

```bash
make tools-install
```

### Development

- `make test`: Run all tests.
- `make lint`: Run the linter.
- `make format`: Format the code using `gofumpt`.
- `make build`: Build the server binary.

The project uses `air` for live-reloading during development. Once configured, you can run:

```bash
./tools/bin/air
```

### API Generation

The project uses `oapi-codegen` for generating Go server/client code from OpenAPI specifications.
Specifications are located in `docs/schema/`.
