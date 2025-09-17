# Use correct Go version
FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Final minimal image
FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]

