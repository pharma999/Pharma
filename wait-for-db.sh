#!/bin/sh
set -e

host="pharma-db"
port=5432

echo "⏳ Waiting for database to be ready at $host:$port..."

while ! nc -z "$host" "$port"; do
  echo "Database not ready - waiting..."
  sleep 2
done

echo "✅ Database is up!"
echo "🔎 DB_CONN string:"
echo "$DB_CONN"

# Start the app using absolute path
exec /app/main
