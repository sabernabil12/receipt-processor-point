# Variables
BINARY_NAME = receipt-processor-point
SRC_DIR = ./cmd/server

# Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) $(SRC_DIR)

# Run the server
run: build
	@echo "Starting the server..."
	./bin/$(BINARY_NAME)

# Install necessary tools
install-tools:
	@echo "Installing tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang/mock/mockgen@latest
	@echo "Tools installed."

# Run linter
lint:
	@echo "Linting code..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...

# Run tests
test:
	@echo "Running tests..."
	go test -v internal/services/*

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin/
