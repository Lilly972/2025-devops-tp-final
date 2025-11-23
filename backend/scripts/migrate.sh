#!/bin/bash
set -e

DB_URL="${DATABASE_URL:-postgres://postgres:postgres@localhost:5432/christmas_gifts?sslmode=disable}"

echo "Running migrations..."

# Install golang-migrate if not already installed
if ! command -v migrate &> /dev/null; then
    echo "Installing golang-migrate..."
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
fi

migrate -path ./migrations -database "$DB_URL" up

echo "Migrations completed successfully!"
