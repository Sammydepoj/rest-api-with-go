APP_NAME=rest-api
BIN_NAME =gobin
BUILD_DIR=./bin
GO_FILES=$(shell find . -name "*.go" -not -path ".vendor/*")

run:
	@echo "Running the server"
	@go run main.go || true
deps: 
	@echo "Installing dependencies"
	@go mod tidy
fmt:
	@echo "Formatting the code"
	@go fmt ./...
build:
	@echo "Building the project"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BIN_NAME) main.go
	@echo "Build Complete: $(BUILD_DIR)"
clean:
	@echo "Cleaning up"
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned up complete"
stop:
	@echo "Stopping the server"
	@pkill -f "go run main.go" ||  echo "Server not running"
migrate-up:
	@echo "Running migrations"
	@dbmate up
	@echo "Migrations applied"
migrate-down:
	@echo "Rolling back migrations"
	@dbmate down
	@echo "Migrations rolled back"
help:
	@echo "Available commands:"
	@echo "make run - Run the server"
	@echo "make deps - Install dependencies"
	@echo "make fmt - Format the code"
	@echo "make build - Build the project"
	@echo "make clean - Clean up the project"
	@echo "make stop - Stop the server"
	@echo "make migrate-up - Run migrations"
	@echo "make migrate-down - Rollback migrations"
	@echo "make help - Show this help message"


