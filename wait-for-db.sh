#!/bin/sh
set -e

# Default values if env variables are missing
DB_HOST=${DB_HOST:-db}
DB_PORT=${DB_PORT:-5432}

echo "Waiting for database at $DB_HOST:$DB_PORT..."

# Loop until database port is open
while ! nc -z "$DB_HOST" "$DB_PORT"; do
  echo "Database not ready - waiting..."
  sleep 2
done

echo "Database is up - starting app!"
exec ./main
