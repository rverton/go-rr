sqlc/generate:
	cd internal/database && sqlc generate

frontend/watch:
	cd frontend && npm run dev

.PHONY: frontend/build
frontend/build:
	cd frontend && npm run build

run:
	go run ./cmd/gorr/
