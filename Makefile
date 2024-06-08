# Simple Makefile for a Go project

# Build the application
all: build

run: docker air


build:
	@echo "Building..."
	@templ generate
	@go build -o main cmd/api/main.go

# Generate templ
templ: 
	@templ generate
	
# Run Air
air: 
	@echo "Application is running..."
	@air
# Run the application




# Run docker compose
docker:
	@docker compose up -d 

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
