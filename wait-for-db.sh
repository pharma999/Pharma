#!/bin/sh

# wait-for-db.sh

set -e

host="pharma-db"
port=5432

echo "⏳ Waiting for database to be ready at $host:$port..."

# Wait until PostgreSQL is ready
while ! nc -z "$host" "$port"; do
  echo "Database not ready - waiting..."
  sleep 2
done

echo "✅ Database is up!"

# Debug: Print the DB_CONN string
echo "🔎 DB_CONN string:"
echo "$DB_CONN"

# Start the app
echo "🚀 Starting the app..."
exec sh -c "./main"
