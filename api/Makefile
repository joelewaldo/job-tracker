BINARY_NAME := server
CMD_PATH := ./cmd/server
BIN_DIR := ./bin
MIGRATIONS_DIR := migrations
DATABASE_URL ?= postgres://user:pass@localhost:4001/db?sslmode=disable
MIGRATE_CMD := migrate

# Check if migrate command exists, otherwise use go migrate
ifeq (, $(shell which migrate))
    MIGRATE_CMD := go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
endif

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_PATH)

run:
	@echo "Running $(BINARY_NAME)..."
	@$(BIN_DIR)/$(BINARY_NAME)

dev: build run

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

test:
	@echo "Running tests..."
	@go test ./...

# Migrations
migrate-up:
	@echo "Running migrations up..."
	@$(MIGRATE_CMD) -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" up

migrate-down:
	@echo "Rolling back last migration..."
	@$(MIGRATE_CMD) -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" down

migrate-create:
	@echo "Creating new migration..."
	@if [ -z "$(NAME)" ]; then echo "Please provide NAME=<migration_name>"; exit 1; fi
	@$(MIGRATE_CMD) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

.PHONY: all build run dev clean test migrate-up migrate-down migrate-create
