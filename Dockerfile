# Build stage
FROM golang:1.23.4 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o pharma .

# Final stage
FROM ubuntu:22.04
RUN apt-get update && \
    apt-get install -y ca-certificates tzdata netcat-openbsd && \
    ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/pharma .

# Copy wait-for-db script
COPY wait-for-db.sh .
RUN chmod +x wait-for-db.sh

# Copy env file (if needed)
COPY .env .env

ENV TZ=Asia/Kolkata

# Start with wait-for-db.sh
CMD ["sh", "wait-for-db.sh"]
