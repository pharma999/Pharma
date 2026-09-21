#!/usr/bin/env bash
set -Eeuo pipefail

# Wait for a TCP service before starting the application.
# Override these values through the container environment when needed.
DB_HOST="${DB_HOST:-pharma-db}"
DB_PORT="${DB_PORT:-5432}"
DB_TIMEOUT="${DB_TIMEOUT:-60}"

wait_for_service() {
  local host="$1"
  local port="$2"
  local name="$3"
  local deadline=$((SECONDS + DB_TIMEOUT))

  echo "Waiting for ${name} at ${host}:${port}..."
  while ! (echo >"/dev/tcp/${host}/${port}") >/dev/null 2>&1; do
    if (( SECONDS >= deadline )); then
      echo "Timed out after ${DB_TIMEOUT}s waiting for ${name} at ${host}:${port}" >&2
      exit 1
    fi
    sleep 2
  done
  echo "${name} is ready."
}

wait_for_service "$DB_HOST" "$DB_PORT" "PostgreSQL"

# Replace this with the MongoDB health check if MongoDB is required before startup:
# wait_for_service "${MONGO_HOST:-mongo}" "${MONGO_PORT:-27017}" "MongoDB"

exec /app/main "$@"
