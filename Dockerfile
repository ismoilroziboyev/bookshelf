# Build stage
FROM golang:1.21-alpine3.20 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o bookshelf cmd/main.go


# Run stage
FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/bookshelf .
COPY --from=builder /app/migrations migrations/

ENTRYPOINT [ "/app/bookshelf" ]
