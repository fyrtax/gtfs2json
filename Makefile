# Project name
PROJECT_NAME := gtfs2json

# Version - This will be overridden by the Git tag in the CI workflow
VERSION := dev

# Build directory
BUILD_DIR := build

# Target platforms
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 linux/arm64

# Standard compilation flags - We'll ensure main.Version is set
LDFLAGS_BASE := -w -s # Basic flags for smaller binaries: -s (strip symbol table), -w (omit DWARF symbol table)

# Dynamic LDFLAGS based on VERSION
LDFLAGS := -ldflags "-X main.Version=$(VERSION) $(LDFLAGS_BASE)"
STATIC_LDFLAGS := -ldflags "-X main.Version=$(VERSION) $(LDFLAGS_BASE) -extldflags '-static'"


.PHONY: all clean build release cli static-cli

all: clean build

clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)

build: cli

cli:
	@echo "Building CLI for current platform..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_cli .

static-cli:
	@echo "Building statically linked CLI for current platform..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build $(STATIC_LDFLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_cli .

release:
	@echo "Building release binaries for all platforms with Version: $(VERSION)..."
	@mkdir -p $(BUILD_DIR)/release
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -f1 -d'/'); \
		ARCH=$$(echo $$platform | cut -f2 -d'/'); \
		output_name=$(PROJECT_NAME)_$(VERSION)_$$OS-$$ARCH; \
		if [ "$$OS" = "windows" ]; then \
			output_name="$${output_name}.exe"; \
		fi; \
		echo "Building for $$OS/$$ARCH -> $(BUILD_DIR)/release/$$output_name"; \
		GOOS=$$OS GOARCH=$$ARCH CGO_ENABLED=0 go build $(STATIC_LDFLAGS) -o $(BUILD_DIR)/release/$$output_name .; \
	done
	@echo "Release binaries built in $(BUILD_DIR)/release/"

# Target to display available commands
help:
	@echo "Available commands:"
	@echo "  make build           - Compiles the project CLI for the current platform"
	@echo "  make cli             - Compiles only the CLI for the current platform"
	@echo "  make static-cli      - Compiles a statically linked CLI for the current platform"
	@echo "  make release         - Creates statically linked releases for all platforms (binaries in build/release/)"
	@echo "  make clean           - Removes the build directory"
	@echo "  make all             - Executes clean, then build"
	@echo "  make help            - Displays this help message"