BINARY_PATH=./bin/server/server.out
SERVER_PATH=./cmd/server/main.go
DB_CREATE_PATH=./cmd/database/create/main.go

createdb:
	@echo "Creating database..."
	@go run ${DB_CREATE_PATH}

build:
	@echo "Building server at ${BINARY_PATH}..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}

run:
	@echo "Running server..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}
	@eval ${BINARY_PATH}

fmt:
	@echo "Formatting source files..."
	@go fmt ./...

clean:
	@echo "Removing object files and cached files..."
	@go clean
	@rm ${BINARY_PATH}
