run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

tidy:
	go mod tidy

migrate:
	cat migrations/*.sql | psql "$(DATABASE_URL)"

test:
	go test ./...
