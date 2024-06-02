# builder
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bootstrap cmd/bootstrap/main.go

# runner
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/bootstrap .
EXPOSE 8080
CMD ["./bootstrap"]
