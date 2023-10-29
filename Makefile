.PHONY: help
help:
	@echo "-Комп перезагружал?\n-Да!\n-Клаву протирал?\n-Да!\n-Тогда не знаю в чем проблема..."

.PHONY: up
up:
	docker compose up --build

.PHONY: up_local_infra
up_local_infra:
	docker compose -f docker-compose.dev.yml up --build

.PHONY: migrate
migrate:
	cd migrations && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./migration .
	cd migrations && ./migration

.PHONY: up_local_app
up_local_app:
	cd app && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/main ./cmd/app/main.go
	cd app && ./build/main -host=localhost -port=6379

.PHONY: lint
lint:
	cd app && golangci-lint run --config ./build/.golangci.yml ./...

.PHONY: test
test:
	cd app && go test ./...

.PHONY: test_race
test_race:
	cd app && go test -race -short ./...

.PHONY: generate
generate:
	cd app && go generate ./...
