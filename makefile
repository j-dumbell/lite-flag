export DB_USER=myuser
export DB_PASSWORD=mypassword
export DB_NAME=postgres
export DB_HOST=localhost
export DB_PORT=5432

start-api:
	cd api && docker compose up -d postgres && \
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) go run cmd/api/main.go

bootstrap:
	cd api && docker compose up -d postgres && \
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) go run cmd/bootstrap/main.go

build-api:
	cd api && go build -o bin/api cmd/api/main.go

build-bootstrap:
	cd api && go build -o bin/bootstrap cmd/bootstrap/main.go

docker-build-api:
	docker build -f api/cmd/api/Dockerfile -t jdumbell92/lite-flag:api-0.1 ./api

docker-build-bootstrap:
	docker build -f api/cmd/bootstrap/Dockerfile -t jdumbell92/lite-flag:bootstrap-0.1 ./api

docker-run-bootstrap:
	cd api && docker compose up postgres bootstrap

docker-run-api:
	cd api && docker compose up postgres api

test:
	cd api && go test ./...

lint:
	cd api && golangci-lint run

oapi-gen:
	cd api && oapi-codegen --config=oapi-codegen.yaml ../oapi/spec.yaml

gen-clients:
	openapi-generator generate -i ./oapi/spec.yaml -g typescript-axios -o ./clients/typescript && \
	openapi-generator generate -i ./oapi/spec.yaml -g go -o ./clients/go
