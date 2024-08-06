# Project name
PROJECT_NAME := gtfs2json

# Version
VERSION := 1.0.0

# Main Go files
CLI_MAIN := cmd/cli/main.go
WEBSERVER_MAIN := cmd/webserver/main.go

# Build directory
BUILD_DIR := build

# Target platforms
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 linux/arm64

# Standard compilation flags
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

# Static compilation flags
STATIC_FLAGS := -ldflags "-X main.Version=$(VERSION) -w -extldflags '-static'"

.PHONY: all clean build release cli webserver static-cli static-webserver static-release

all: clean build

clean:
	rm -rf $(BUILD_DIR)

build: cli webserver

cli:
	go build -o $(BUILD_DIR)/$(PROJECT_NAME)_cli $(LDFLAGS) $(CLI_MAIN)

webserver:
	go build -o $(BUILD_DIR)/$(PROJECT_NAME)_webserver $(LDFLAGS) $(WEBSERVER_MAIN)

static-cli:
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(PROJECT_NAME)_cli $(STATIC_FLAGS) $(CLI_MAIN)

static-webserver:
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(PROJECT_NAME)_webserver $(STATIC_FLAGS) $(WEBSERVER_MAIN)

release:
	@mkdir -p $(BUILD_DIR)/release
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -f1 -d'/'); \
		ARCH=$$(echo $$platform | cut -f2 -d'/'); \
		cli_output_name=$(PROJECT_NAME)_cli; \
		webserver_output_name=$(PROJECT_NAME)_webserver; \
		if [ "$$OS" = "windows" ]; then \
			cli_output_name="$$cli_output_name.exe"; \
			webserver_output_name="$$webserver_output_name.exe"; \
		fi; \
		echo "Building for $$OS/$$ARCH..."; \
		GOOS=$$OS GOARCH=$$ARCH CGO_ENABLED=0 go build $(STATIC_FLAGS) -o $(BUILD_DIR)/release/$(PROJECT_NAME)_$(VERSION)_$$OS-$$ARCH/$$cli_output_name $(CLI_MAIN); \
		GOOS=$$OS GOARCH=$$ARCH CGO_ENABLED=0 go build $(STATIC_FLAGS) -o $(BUILD_DIR)/release/$(PROJECT_NAME)_$(VERSION)_$$OS-$$ARCH/$$webserver_output_name $(WEBSERVER_MAIN); \
		cd $(BUILD_DIR)/release && zip -r $(PROJECT_NAME)_$(VERSION)_$$OS-$$ARCH.zip $(PROJECT_NAME)_$(VERSION)_$$OS-$$ARCH; \
		cd ../.. ; \
	done

static-release: release

# Target to display available commands
help:
	@echo "Available commands:"
	@echo "  make build           - Compiles the project (both CLI and Webserver) for the current platform"
	@echo "  make cli             - Compiles only the CLI for the current platform"
	@echo "  make webserver       - Compiles only the Webserver for the current platform"
	@echo "  make static-cli      - Compiles a statically linked CLI for the current platform"
	@echo "  make static-webserver - Compiles a statically linked Webserver for the current platform"
	@echo "  make release         - Creates statically linked releases for all platforms"
	@echo "  make clean           - Removes the build directory"
	@echo "  make all             - Executes clean, then build"
	@echo "  make help            - Displays this help message"