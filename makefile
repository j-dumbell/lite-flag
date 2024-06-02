export DB_USER=myuser
export DB_PASSWORD=mypassword
export DB_NAME=postgres
export DB_HOST=localhost
export DB_PORT=5432

start-api:
	docker compose up -d postgres && \
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) go run cmd/api/main.go


bootstrap:
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) go run cmd/bootstrap/main.go

build-api:
	go build -o bin/api cmd/api/main.go

build-bootstrap:
	go build -o bin/bootstrap cmd/bootstrap/main.go

docker-build-api:
	docker build -f cmd/api/Dockerfile -t jdumbell92/lite-flag:api-0.1 .

docker-build-bootstrap:
	docker build -f cmd/bootstrap/Dockerfile -t jdumbell92/lite-flag:bootstrap-0.1 .

docker-run-bootstrap:
	docker compose up postgres bootstrap

docker-run-api:
	docker compose up postgres api

test:
	go test ./...
