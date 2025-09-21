#!/bin/sh

set -e

echo "Waiting for database at $DB_CONN ..."

# Extract host and port from DB_CONN for the health check
DB_HOST=$(echo $DB_CONN | sed -n 's/.*host=\([^ ]*\).*/\1/p')
DB_PORT=$(echo $DB_CONN | sed -n 's/.*port=\([^ ]*\).*/\1/p')

until nc -z "$DB_HOST" "$DB_PORT"; do
  echo "Database not ready - waiting..."
  sleep 2
done

echo "Database is up - starting app!"
# Start your application (adjust the binary name if needed)
exec ./pharma
