.PHONY: build clean

# Variables (can be overridden)
BINARY_NAME ?= bloggo
FRONTEND_DIR ?= frontend
BACKEND_DIR ?= backend
OUTPUT_DIR ?= build
EMBED_DIR = $(BACKEND_DIR)/frontend

# Cross-compilation variables
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 1
CC ?=
LDFLAGS ?= -s -w

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
	@echo "🔧 Building backend (GOOS=$(GOOS) GOARCH=$(GOARCH))..."
	@cd $(BACKEND_DIR) && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(if $(CC),CC=$(CC),) \
		go build -ldflags="$(LDFLAGS)" -o ../$(OUTPUT_DIR)/$(BINARY_NAME) cli/main.go
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
