# Simple Makefile for Go Blueprint

BINARY_NAME=go-blueprint
MAIN_PATH=./main.go
VERSION=$(shell cat version)

.PHONY: deps build run dev tag

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

# Create git tag from version file
tag:
	@echo "Creating git tag v$(VERSION)..."
	git tag v$(VERSION)
	git push origin v$(VERSION)
	@echo "Tag v$(VERSION) created and pushed"