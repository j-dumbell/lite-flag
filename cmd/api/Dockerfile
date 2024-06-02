# builder
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api cmd/api/main.go

# runner
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/api .
EXPOSE 8080
CMD ["./api"]
