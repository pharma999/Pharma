#!/bin/sh
set -e
until nc -z "$DB_HOST" "$DB_PORT"; do
  echo "Waiting for database..."
  sleep 2
done
exec ./main

