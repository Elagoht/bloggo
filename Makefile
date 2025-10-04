.PHONY: build clean dev-backend dev-frontend help build-linux-amd64 build-linux-arm64

# Variables
BINARY_NAME=bloggo
FRONTEND_DIR=frontend
BACKEND_DIR=backend
OUTPUT_DIR=build
EMBED_DIR=$(BACKEND_DIR)/internal/embed

# Build the full-stack application (frontend + backend in single binary)
build:
	@echo "🏗️  Building Bloggo Full-Stack Application..."
	@echo ""
	@$(MAKE) build-frontend
	@$(MAKE) embed-frontend
	@$(MAKE) build-backend
	@echo ""
	@echo "✅ Build complete! Binary: $(OUTPUT_DIR)/$(BINARY_NAME)"
	@echo "🚀 Run with: cd $(OUTPUT_DIR) && ./$(BINARY_NAME)"

# Build frontend
build-frontend:
	@echo "📦 Installing frontend dependencies..."
	@cd $(FRONTEND_DIR) && npm ci
	@echo "📦 Building frontend..."
	@cd $(FRONTEND_DIR) && npm run build

# Embed frontend into backend
embed-frontend:
	@echo "📋 Copying frontend to backend..."
	@rm -rf $(EMBED_DIR)/dist
	@mkdir -p $(EMBED_DIR)
	@cp -r $(FRONTEND_DIR)/dist $(EMBED_DIR)/

# Build backend with embedded frontend
build-backend:
	@echo "📁 Creating output directory..."
	@mkdir -p $(OUTPUT_DIR)
	@echo "🔧 Building backend..."
	@cd $(BACKEND_DIR) && go build -o ../$(OUTPUT_DIR)/$(BINARY_NAME) cli/main.go
	@echo "📋 Copying configuration files..."
	@cp $(BACKEND_DIR)/.env.example $(OUTPUT_DIR)/.env.example
	@if [ -f $(BACKEND_DIR)/.env ]; then cp $(BACKEND_DIR)/.env $(OUTPUT_DIR)/.env; fi

# Build for Linux AMD64
build-linux-amd64:
	@echo "🐧 Building for Linux AMD64..."
	@$(MAKE) build-frontend
	@$(MAKE) embed-frontend
	@mkdir -p $(OUTPUT_DIR)
	@cd $(BACKEND_DIR) && GOOS=linux GOARCH=amd64 go build -o ../$(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64 cli/main.go
	@cp $(BACKEND_DIR)/.env.example $(OUTPUT_DIR)/.env.example
	@echo "✅ Linux AMD64 build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64"

# Build for Linux ARM64
build-linux-arm64:
	@echo "🐧 Building for Linux ARM64..."
	@$(MAKE) build-frontend
	@$(MAKE) embed-frontend
	@mkdir -p $(OUTPUT_DIR)
	@cd $(BACKEND_DIR) && GOOS=linux GOARCH=arm64 go build -o ../$(OUTPUT_DIR)/$(BINARY_NAME)-linux-arm64 cli/main.go
	@cp $(BACKEND_DIR)/.env.example $(OUTPUT_DIR)/.env.example
	@echo "✅ Linux ARM64 build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-linux-arm64"

# Build all Linux binaries
build-all-linux: build-linux-amd64 build-linux-arm64
	@echo "✅ All Linux builds complete"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(FRONTEND_DIR)/dist
	@rm -rf $(EMBED_DIR)/dist
	@rm -rf $(OUTPUT_DIR)
	@echo "✅ Clean complete"

# Run backend in development mode (without embedded frontend)
dev-backend:
	@echo "🔧 Starting backend in development mode..."
	@cd $(BACKEND_DIR) && go run cli/main.go

# Run frontend in development mode
dev-frontend:
	@echo "⚛️  Starting frontend in development mode..."
	@cd $(FRONTEND_DIR) && npm run dev

# Show help
help:
	@echo "Bloggo Build Commands:"
	@echo ""
	@echo "  make build              - Build full-stack application (single binary)"
	@echo "  make build-linux-amd64  - Build for Linux AMD64"
	@echo "  make build-linux-arm64  - Build for Linux ARM64"
	@echo "  make build-all-linux    - Build all Linux binaries"
	@echo "  make clean              - Clean build artifacts"
	@echo "  make dev-backend        - Run backend in development mode"
	@echo "  make dev-frontend       - Run frontend in development mode"
	@echo "  make help               - Show this help message"
