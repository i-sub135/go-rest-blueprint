# Simple Makefile for Go Blueprint

BINARY_NAME=go-blueprint
MAIN_PATH=./main.go

.PHONY: deps build run dev

# Load dependencies and tidy modules
deps:
	@echo "Loading dependencies..."
	go mod download
	go mod tidy

# Build the application
build:
	@echo "Building application..."
	go build -o build/$(BINARY_NAME) $(MAIN_PATH)

# Run the application
run:
	@echo "Running application..."
	go run $(MAIN_PATH)

# Development with hot reload
dev:
	@echo "Starting hot reload..."
	find . -name "*.go" | entr -r go run $(MAIN_PATH)