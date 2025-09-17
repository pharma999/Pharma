# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary (replace 'server' with your app name if needed)
RUN go build -o server .

# Stage 2: Ubuntu-based runtime
FROM ubuntu:22.04

# Set working directory
WORKDIR /app

# Install minimal dependencies (like CA certificates)
RUN apt-get update && apt-get install -y \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

# Copy the Go binary from builder
COPY --from=builder /app/server .

# Expose app port
EXPOSE 8080

# Run the app
CMD ["./server"]

