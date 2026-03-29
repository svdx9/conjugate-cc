.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  generate           Generate code from OpenAPI specs (backend + frontend)"
	@echo "  frontend-generate  Generate frontend TypeScript types from OpenAPI"
	@echo "  frontend-dev       Start frontend dev server"
	@echo "  backend-dev       Start backend dev server"
	@echo "  dev-all           Start both frontend and backend dev servers"
	@echo "  test              Run all tests"
	@echo "  lint              Run all linters"

frontend-generate:
	cd frontend && npm run generate

generate: backend/generate frontend-generate

frontend-dev:
	cd frontend && npm run dev

backend-dev:
	cd backend && make build && cd bin && ./server

dev-all:
	make -j2 backend-dev frontend-dev

test:
	cd backend && make test
	cd frontend && npm run test

lint:
	cd backend && make lint
	cd frontend && npm run lint
