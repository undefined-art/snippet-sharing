BINARY_NAME=gin-rest

build:
	@echo "Building binary..."
	go build -o $(BINARY_NAME) ./cmd

run: build
	@echo "Running app..."
	./$(BINARY_NAME)

test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)

fmt:
	@echo "Formatting code..."
	go fmt ./...

lint:
	@echo "Linting code..."
	golangci-lint run

.PHONY: build run test clean fmt lint
