BINARY_PATH=./bin/server
SERVER_PATH=./cmd/server/main.go
LAMBDA_BINARY_PATH=./bin/lambdaserver
LAMBDA_SERVER_PATH=./cmd/lambdaserver/main.go
DB_CREATE_PATH=./cmd/database/create/main.go
DB_MIGRATE_PATH=./cmd/database/migrate/main.go
DB_SEED_PATH=./cmd/seed/main.go
DB_DROP_PATH=./cmd/database/drop/main.go
TESTDB_CREATE_PATH=./cmd/testdatabase/create/main.go
TESTDB_MIGRATE_PATH=./cmd/testdatabase/migrate/main.go
TESTDB_DROP_PATH=./cmd/testdatabase/drop/main.go

createdb:
	@echo "Creating database..."
	@go run ${DB_CREATE_PATH}

migratedb:
	@echo "Migrating database..."
	@go run ${DB_MIGRATE_PATH}

seeddb:
	@echo "Seeding database..."
	@go run ${DB_SEED_PATH}

dropdb:
	@echo "Dropping database..."
	@go run ${DB_DROP_PATH}

createtestdb:
	@echo "Creating test database..."
	@go run ${TESTDB_CREATE_PATH}

migratetestdb:
	@echo "Migrating test database..."
	@go run ${TESTDB_MIGRATE_PATH}

droptestdb:
	@echo "Dropping test database..."
	@go run ${TESTDB_DROP_PATH}

build:
	@echo "Building server at ${BINARY_PATH}..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}

buildlambda:
	@echo "Building AWS Lambda server at ${LAMBDA_BINARY_PATH}..."
	@go build -o ${LAMBDA_BINARY_PATH} ${LAMBDA_SERVER_PATH}

run:
	@echo "Running server..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}
	@eval ${BINARY_PATH}

runlambda:
	@echo "Running AWS Lambda server..."
	@go build -o ${LAMBDA_BINARY_PATH} ${LAMBDA_SERVER_PATH}
	@eval ${LAMBDA_BINARY_PATH}

fmt:
	@echo "Formatting source files..."
	@go fmt ./...

lint:
	@echo "Linting source files..."
	@golangci-lint run

lintfix:
	@echo "Linting and automatically fixing source files..."
	@golangci-lint run --fix

clean:
	@echo "Removing object files and cached files..."
	@go clean
	@rm ${BINARY_PATH}

test:
	@echo "Running tests..."
	@go test ./...
