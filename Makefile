# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@./tailwindcss -i cmd/web/static/css/input.css -o cmd/web/static/css/output.css
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

templ:
	@templ generate -watch -proxy="http://localhost:8080/"

tailwind:
	@./tailwindcss -i cmd/web/static/css/input.css -o cmd/web/css/static/output.css --watch

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
