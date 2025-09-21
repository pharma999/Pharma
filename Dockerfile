# Stage 1: Build Go binary for Linux
FROM golang:1.23.4 AS builder
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary for Linux
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Final runtime image
FROM ubuntu:22.04
RUN apt-get update && \
    apt-get install -y ca-certificates tzdata netcat-openbsd && \
    ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary and wait script from builder
COPY --from=builder /app/main .
COPY wait-for-db.sh .

# Make them executable
RUN chmod +x main wait-for-db.sh

# Run wait script
CMD ["sh", "wait-for-db.sh"]
