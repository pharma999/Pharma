#!/bin/sh

# Default values if environment variables are not set
DB_HOST=${DB_HOST:-db}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME=${DB_NAME:-healthcare}

# Construct connection string
DB_CONN="host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable TimeZone=Asia/Kolkata"

echo "⏳ Waiting for database to be ready at ${DB_HOST}:${DB_PORT}..."

# Wait until the database is accepting connections
while ! nc -z ${DB_HOST} ${DB_PORT}; do
  sleep 2
  echo "Database not ready - waiting..."
done

echo "✅ Database is up!"
echo "🔎 DB_CONN string: $DB_CONN"

# Run the Go application
exec /app/main
