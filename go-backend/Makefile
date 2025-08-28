.PHONY: run build test clean deps

# Default target
all: deps build

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o workout-builder-api main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f workout-builder-api

# Run with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	air

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Show help
help:
	@echo "Available commands:"
	@echo "  run          - Run the application"
	@echo "  build        - Build the application"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  dev          - Run with hot reload (requires air)"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  deps         - Install dependencies"
	@echo "  install-tools - Install development tools"
	@echo "  help         - Show this help"
