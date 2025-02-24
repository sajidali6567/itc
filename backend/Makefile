APP_NAME = itc
BIN_DIR = bin
SRC_DIR = .
GO_FILES = $(shell find $(SRC_DIR) -type f -name '*.go')

.PHONY: all test build clean run

all: test build

test:
	@echo "Running unit tests..."
	@go test ./... -v
	@echo "All tests passed!"

build: test
	@echo "Building the application..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(SRC_DIR)/main.go
	@echo "Build successful! Binary is in $(BIN_DIR)/$(APP_NAME)"

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)
	@echo "Cleanup complete!"

