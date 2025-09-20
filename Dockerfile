# ---- Build Stage ----
FROM golang:1.23.4 AS builder

WORKDIR /app

# Download Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build Go binary
RUN go build -o main .

# ---- Final Minimal Image ----
FROM ubuntu:22.04

# Install dependencies including netcat for wait-for-db.sh
RUN apt-get update && \
    apt-get install -y ca-certificates tzdata netcat-openbsd && \
    ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary and wait script
COPY --from=builder /app/main .
COPY wait-for-db.sh .

# Copy environment file
COPY .env .env

# Set time zone
ENV TZ=Asia/Kolkata

# Make script executable
RUN chmod +x wait-for-db.sh

# Start the app using wait-for-db
CMD ["sh", "wait-for-db.sh"]
