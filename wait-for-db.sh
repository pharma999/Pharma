# #!/bin/sh

# # Default values if environment variables are not set
# DB_HOST=${DB_HOST}
# DB_PORT=${DB_PORT}
# DB_USER=${DB_USER}
# DB_PASSWORD=${DB_PASSWORD}
# DB_NAME=${DB_NAME}

# # MongoDB defaults
# MONGO_HOST=${MONGO_HOST}
# MONGO_PORT=${MONGO_PORT}
# MONGO_DBNAME=${MONGO_DBNAME}


# # Construct connection string
# DB_CONN="host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable TimeZone=Asia/Kolkata"

# echo "⏳ Waiting for database to be ready at ${DB_HOST}:${DB_PORT}..."

# # Wait until the database is accepting connections
# while ! nc -z ${DB_HOST} ${DB_PORT}; do
#   sleep 2
#   echo "Database not ready - waiting..."
# done

# echo "✅ Database is up!"
# echo "🔎 DB_CONN string: $DB_CONN"

# echo "⏳ Waiting for MongoDB at ${MONGO_HOST}:${MONGO_PORT}..."
# while ! nc -z ${MONGO_HOST} ${MONGO_PORT}; do
#   sleep 2
#   echo "MongoDB not ready - waiting..."
# done
# echo "✅ MongoDB is up!"
# echo "🔎 Mongo connection: mongodb://${MONGO_HOST}:${MONGO_PORT}/${MONGO_DBNAME}"

# # Run the Go application
# exec /app/main
