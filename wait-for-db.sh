#!/bin/sh
set -e
#
echo "Waiting for database at $DB_HOST:$DB_PORT..."
#
## Loop until the DB is ready
until nc -z "$DB_HOST" "$DB_PORT"; do
  echo "Database not ready - waiting..."
  sleep 2
done

echo "Database is up! Starting application..."
exec ./main
