.PHONY: dev dev-fe dev-be install generate-proto

# Default target
dev: generate-proto
	@make -j2 dev-fe dev-be

# Start app development server
dev-fe:
	@echo "Starting App development server..."
	cd blogger-app && bun run dev

# Start service development server
dev-be:
	@echo "Starting service development server..."
	cd blogger-service && nodemon --exec go run cmd/api/main.go --signal SIGTERM

# Install dependencies
install:
	@echo "Installing app dependencies..."
	cd blogger-app && bun install
	@echo "Installing service dependencies..."
	cd blogger-service && go mod download

# Generate protobuf code
generate-proto:
	@echo "Generating protobuf code..."
	cd protobuf && buf generate

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -rf blogger-app/protobuf/generated
	rm -rf blogger-service/proto

# Help target
help:
	@echo "Available targets:"
	@echo "  dev          - Start both app and service development servers"
	@echo "  dev-fe       - Start app development server"
	@echo "  dev-be       - Start service development server"
	@echo "  install      - Install dependencies for both app and service"
	@echo "  generate-proto - Generate protobuf code"
	@echo "  clean        - Clean generated files"
	@echo "  help         - Show this help message"
