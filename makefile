APP_NAME=snippet-sharing
IMAGE_NAME=snippet-sharing-image
CONTAINER_NAME=snippet-sharing-container
PORT=8080

.PHONY: build run clean docker-build docker-run docker-clean

build:
	go mod tidy
	go build -ldflags="-w -s" -o $(APP_NAME) ./cmd/main.go

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run --rm -it -p $(PORT):$(PORT) -v $(PWD)/data:/app/data --name $(CONTAINER_NAME) $(IMAGE_NAME)

docker-clean:
	docker rm -f $(CONTAINER_NAME) || true
	docker rmi $(IMAGE_NAME) || true