#!/bin/sh
set -e

# Default values if not provided
DB_HOST=${DB_HOST:-db}
DB_PORT=${DB_PORT:-5432}

echo "Waiting for database at $DB_HOST:$DB_PORT ..."

# Wait until DB is ready
until nc -z "$DB_HOST" "$DB_PORT"; do
  echo "Database not ready - waiting..."
  sleep 2
done

echo "Database is up! Starting the app..."
exec ./main
