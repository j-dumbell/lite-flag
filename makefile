export DB_USER=myuser
export DB_PASSWORD=mypassword
export DB_NAME=postgres
export DB_HOST=localhost
export DB_PORT=5432
export CIPHER_KEY=abcdefghijklmnopqrstuvwx

start:
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) go run cmd/api/main.go

bootstrap:
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) CIPHER_KEY=$(CIPHER_KEY) go run cmd/bootstrap/main.go

test:
	go test ./...
