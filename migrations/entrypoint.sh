#!/bin/bash

set -e

DBSTRING="user=$POSTGRES_USER password=$POSTGRES_PASSWORD database=$POSTGRES_DB_NAME host=$POSTGRES_HOST port=$POSTGRES_PORT sslmode=$POSTGRES_SSLMODE"

echo "Running migrations..."

goose postgres "$DBSTRING" up
