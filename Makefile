build:
	@go build -o bin/ cmd/api/main.go

run: build
	@./bin/main

test:
	@go test -v ./...