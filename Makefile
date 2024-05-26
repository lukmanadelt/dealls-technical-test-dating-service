dependencies:
	@echo "Running dating service dependencies..."
	docker-compose -f docker-compose-dependencies.yaml up

build-origin:
	@echo "Building dating service..."
	@go build -o ./dealls-technical-test-dating-service ./cmd/app/...
	@echo "Finish building dating service"

run-origin:
	@echo "Running dating service..."
	@./dealls-technical-test-dating-service

build:
	@echo "Building dating service..."
	docker-compose build

run:
	@echo "Running dating service...."
	docker-compose up

test:
	@echo "Running tests..."
	@go test ./...
	@echo "Finish running tests"

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run
	@echo "Finish running golangci-lint"
