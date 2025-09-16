# Abdal Better PassList - Makefile
# Author: Ebrahim Shafiei (EbraSha)
# Email: Prof.Shafiei@Gmail.com

.PHONY: build run clean test help install deps

# Application name
APP_NAME = abdal-better-passlist
BUILD_DIR = build
MAIN_FILE = main.go

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod

# Build flags
LDFLAGS = -ldflags "-X main.version=1.0.0 -X main.buildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')"

# Default target
all: deps build

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Build the application
build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "✅ Build completed: $(BUILD_DIR)/$(APP_NAME)"

# Build for Windows
build-windows:
	@echo "🔨 Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME).exe $(MAIN_FILE)
	@echo "✅ Windows build completed: $(BUILD_DIR)/$(APP_NAME).exe"

# Build for Linux
build-linux:
	@echo "🔨 Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux $(MAIN_FILE)
	@echo "✅ Linux build completed: $(BUILD_DIR)/$(APP_NAME)-linux"

# Build for macOS
build-macos:
	@echo "🔨 Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-macos $(MAIN_FILE)
	@echo "✅ macOS build completed: $(BUILD_DIR)/$(APP_NAME)-macos"

# Build all platforms
build-all: build-windows build-linux build-macos
	@echo "✅ All platform builds completed"

# Run the application in interactive mode
run:
	@echo "🚀 Running $(APP_NAME) in interactive mode..."
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	./$(BUILD_DIR)/$(APP_NAME) --interactive

# Run with specific country
run-country:
	@echo "🚀 Running $(APP_NAME) with country selection..."
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	./$(BUILD_DIR)/$(APP_NAME) --country ir --complexity sensitive --workers 4

# Run tests
test:
	@echo "🧪 Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "🧪 Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "✅ Clean completed"

# Format code
fmt:
	@echo "🎨 Formatting code..."
	$(GOCMD) fmt ./...

# Lint code
lint:
	@echo "🔍 Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Install the application
install: build
	@echo "📦 Installing $(APP_NAME)..."
	@if [ -f $(BUILD_DIR)/$(APP_NAME) ]; then \
		sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/; \
		echo "✅ $(APP_NAME) installed to /usr/local/bin/"; \
	else \
		echo "❌ Build not found. Run 'make build' first."; \
	fi

# Uninstall the application
uninstall:
	@echo "🗑️  Uninstalling $(APP_NAME)..."
	@if [ -f /usr/local/bin/$(APP_NAME) ]; then \
		sudo rm /usr/local/bin/$(APP_NAME); \
		echo "✅ $(APP_NAME) uninstalled"; \
	else \
		echo "⚠️  $(APP_NAME) not found in /usr/local/bin/"; \
	fi

# Show help
help:
	@echo "🔧 Abdal Better PassList - Available Commands:"
	@echo ""
	@echo "📦 Dependencies:"
	@echo "  deps          Install Go dependencies"
	@echo ""
	@echo "🔨 Building:"
	@echo "  build         Build for current platform"
	@echo "  build-windows Build for Windows"
	@echo "  build-linux   Build for Linux"
	@echo "  build-macos   Build for macOS"
	@echo "  build-all     Build for all platforms"
	@echo ""
	@echo "🚀 Running:"
	@echo "  run           Run in interactive mode"
	@echo "  run-country   Run with Iran country preset"
	@echo ""
	@echo "🧪 Testing:"
	@echo "  test          Run tests"
	@echo "  test-coverage Run tests with coverage report"
	@echo ""
	@echo "🎨 Code Quality:"
	@echo "  fmt           Format code"
	@echo "  lint          Lint code"
	@echo ""
	@echo "📦 Installation:"
	@echo "  install       Install to /usr/local/bin/"
	@echo "  uninstall     Remove from /usr/local/bin/"
	@echo ""
	@echo "🧹 Maintenance:"
	@echo "  clean         Clean build artifacts"
	@echo "  help          Show this help message"
	@echo ""
	@echo "👨‍💻 Developer: Ebrahim Shafiei (EbraSha)"
	@echo "📧 Email: Prof.Shafiei@Gmail.com"
