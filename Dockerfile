# ---- Build Stage ----
FROM golang:1.23.4 AS builder

WORKDIR /app

# Download Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary
RUN go build -o main .

# ---- Final Minimal Image ----
FROM ubuntu:22.04

# Install CA certificates and tzdata for time zones
RUN apt-get update && \
    apt-get install -y ca-certificates tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env file into image
COPY .env .env

# Set environment variable for time zone
ENV TZ=Asia/Kolkata

# Run the app
CMD ["./main"]

