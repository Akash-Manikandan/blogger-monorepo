.PHONY: dev dev-fe dev-be install generate-proto

# Default target
dev: generate-proto
	@make -j2 dev-fe dev-be

# Start frontend development server
dev-fe:
	@echo "Starting frontend development server..."
	cd blogger-app && bun run dev

# Start backend development server
dev-be:
	@echo "Starting backend development server..."
	cd blogger-service && nodemon --exec go run cmd/api/main.go --signal SIGTERM

# Install dependencies
install:
	@echo "Installing frontend dependencies..."
	cd blogger-app && npm install
	@echo "Installing backend dependencies..."
	cd blogger-service && go mod download

# Generate protobuf code
generate-proto:
	@echo "Generating protobuf code..."
	./scripts/generate-grpc.sh

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -rf blogger-app/protobuf/generated
	rm -rf blogger-service/proto

# Help target
help:
	@echo "Available targets:"
	@echo "  dev          - Start both frontend and backend development servers"
	@echo "  dev-fe       - Start frontend development server"
	@echo "  dev-be       - Start backend development server"
	@echo "  install      - Install dependencies for both frontend and backend"
	@echo "  generate-proto - Generate protobuf code"
	@echo "  clean        - Clean generated files"
	@echo "  help         - Show this help message"
