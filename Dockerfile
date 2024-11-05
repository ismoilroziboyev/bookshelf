# Build stage
FROM golang:1.23.2-alpine3.19 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o bookshelf cmd/main.go


# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/bookshelf .
COPY --from=builder /app/migrations migrations/

ENTRYPOINT [ "/app/bookshelf" ]
