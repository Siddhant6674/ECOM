build:
	@go build -o bin/ECOM cmd/main.go

test:build
	@go test -v ./...

run:build
	@./bin/ECOM

