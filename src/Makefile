# Don't ask me what is going in this Makefile
include .env
MIGRATION_PATH = ./migrations

.PHONY:fmt vet build
fmt:
		go fmt ./...

vet:	fmt
		go vet ./...

test:	fmt
	@go test -v ./...

build:
	@mkdir -p bin
	if [ -f "bin/${BINARY}" ]; then \
		rm bin/${BINARY}; \
		echo "Deleted old binary"; \
	fi
	@echo "Building binary..."
	go build -o bin/${BINARY} main.go

run: build
	@echo "Starting API..."
	@bin/${BINARY}

deps:
	go mod download

tidy:
	go mod tidy

create_migrations:
	sqlx migrate add -r $(filter-out $@,$(MAKECMDGOALS))

# This line prevents make from trying to interpret the arguments as targets
%:
	@:
migrate-up:
	sqlx migrate run --database-url "${DB_URL}"
migrate-down:
	sqlx migrate revert --database-url "${DB_URL}"
