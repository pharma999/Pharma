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

# Install necessary packages: CA certs, tzdata, netcat
RUN apt-get update && \
    apt-get install -y ca-certificates tzdata netcat-openbsd && \
    ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env file
COPY .env .env

# Copy wait-for-db script
COPY wait-for-db.sh .

# Set environment variable for time zone
ENV TZ=Asia/Kolkata

# Run the app via wait-for-db
CMD ["sh", "wait-for-db.sh"]
