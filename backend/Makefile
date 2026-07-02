ifneq (,$(wildcard .env))
include .env
export
endif

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

tidy:
	go mod tidy

test:
	go test ./...

migrate-up:
	cat migrations/0*.sql | psql "$(DATABASE_URL)"

migrate-down:
	psql "$(DATABASE_URL)" -f migrations/down.sql

seed:
	psql "$(DATABASE_URL)" -f migrations/seed.sql

refresh: migrate-down migrate-up seed

.PHONY: run build tidy test migrate-up migrate-down seed refresh
