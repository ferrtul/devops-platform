APP_NAME=devops-platform
APP_DIR=app

.PHONY: run build test fmt vet lint docker-build docker-up docker-down clean

run:
	cd $(APP_DIR) && go run .

build:
	cd $(APP_DIR) && go build -o ../bin/$(APP_NAME) .

test:
	cd $(APP_DIR) && go test ./...

fmt:
	cd $(APP_DIR) && gofmt -w .

vet:
	cd $(APP_DIR) && go vet ./...

docker-build:
	docker compose build

docker-up:
	docker compose up -d

docker-down:
	docker compose down

clean:
	rm -rf bin