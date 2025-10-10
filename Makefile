.PHONY: build clean

# Variables
BINARY_NAME=bloggo
FRONTEND_DIR=frontend
BACKEND_DIR=backend
OUTPUT_DIR=build
EMBED_DIR=$(BACKEND_DIR)/frontend

# Build the full-stack application (frontend + backend in single binary)
build:
	@echo "🏗️  Building Bloggo Full-Stack Application..."
	@echo ""
	@echo "📦 Installing frontend dependencies..."
	@cd $(FRONTEND_DIR) && npm ci
	@echo "📦 Building frontend..."
	@cd $(FRONTEND_DIR) && npm run build
	@echo "📋 Copying frontend to backend..."
	@rm -rf $(EMBED_DIR)/dist
	@mkdir -p $(EMBED_DIR)
	@cp -r $(FRONTEND_DIR)/dist $(EMBED_DIR)/
	@echo "📁 Creating output directory..."
	@mkdir -p $(OUTPUT_DIR)
	@echo "🔧 Building backend..."
	@cd $(BACKEND_DIR) && go build -o ../$(OUTPUT_DIR)/$(BINARY_NAME) cli/main.go
	@echo "📋 Copying configuration files..."
	@cp $(BACKEND_DIR)/.env.example $(OUTPUT_DIR)/.env.example
	@if [ -f $(BACKEND_DIR)/.env ]; then cp $(BACKEND_DIR)/.env $(OUTPUT_DIR)/.env; fi
	@echo ""
	@echo "✅ Build complete! Binary: $(OUTPUT_DIR)/$(BINARY_NAME)"
	@echo "🚀 Run with: cd $(OUTPUT_DIR) && ./$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(FRONTEND_DIR)/dist
	@rm -rf $(EMBED_DIR)/dist
	@rm -rf $(OUTPUT_DIR)
	@echo "✅ Clean complete"
