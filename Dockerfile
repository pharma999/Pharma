# Stage 1: Build Go binary
FROM golang:1.23.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Runtime image
FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y ca-certificates tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary
COPY --from=builder /app/main .

COPY .env .

# Start application
CMD ["./main"]
