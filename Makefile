build:
	@go build -o bin/ECOM cmd/main.go

test:build
	@go test -v ./...

run:build
	@./bin/ECOM

migration:
	@migrate create -ext sql -dir ./cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run ./cmd/migrate/main.go up

migrate-down:
	@go run ./cmd/migrate/main.go down