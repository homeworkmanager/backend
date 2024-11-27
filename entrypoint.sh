#!/bin/bash
set -e


echo "Running migrations..."
goose -dir ./internal/db/migrations postgres "user=$POSTGRES_USER password=$POSTGRES_PASSWORD database=$POSTGRES_DB_NAME host=$POSTGRES_HOST port=$POSTGRES_PORT" up

echo "Starting application..."
exec "$@"
