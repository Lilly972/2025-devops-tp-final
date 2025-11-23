# Backend - Christmas Gift List API

A RESTful API server built with Go, Chi router, and PostgreSQL for managing Christmas gift lists.

## Tech Stack

- **Go** 1.23+
- **Chi** - Lightweight HTTP router
- **pgx** - PostgreSQL driver and toolkit
- **PostgreSQL** 15+ - Database

## Prerequisites

- Go 1.23 or higher
- PostgreSQL 15 or higher (local installation or Docker)
- Environment variables configured (see below)

## Getting Started

### 1. Install Dependencies

```bash
go mod download
```

### 2. Configure Environment

Create a `.env` file in the `backend/` directory or set these environment variables:

```bash
DATABASE_URL=postgres://postgres:postgres@localhost:5432/christmas_gifts?sslmode=disable
PORT=8080
```

**Environment Variables:**

- `DATABASE_URL` - PostgreSQL connection string
  - Format: `postgres://[user]:[password]@[host]:[port]/[database]?sslmode=[mode]`
  - Default: `postgres://postgres:postgres@localhost:5432/christmas_gifts?sslmode=disable`
- `PORT` - Server port (default: `8080`)

### 3. Setup Database

Use docker or local database

### 4. Run Migrations

Execute the migration script to create the database schema:

```bash
./scripts/migrate.sh
```

This will apply all migrations in the `migrations/` directory.

### 5. Seed Database (Optional)

Populate the database with sample data:

```bash
go run scripts/seed.go
```

### 6. Start the Server

```bash
go run cmd/server/main.go
```

The API server will start at `http://localhost:8080`.

## Development

### Available Commands

```bash
# Run the server
go run cmd/server/main.go

# Build production binary
go build -o bin/server cmd/server/main.go

# Run the binary
./bin/server

# Run tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run tests with verbose output
go test ./... -v

# Run specific test
go test -run TestName ./...

# Check for code issues
go vet ./...

# Format code
go fmt ./...
```

### Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── database/
│   │   └── database.go       # Database connection setup
│   ├── handlers/
│   │   ├── gift.go           # Gift endpoints
│   │   └── person.go         # Person endpoints
│   ├── models/
│   │   ├── models.go         # Data models
│   │   └── models_test.go    # Model tests
│   └── repository/
│       ├── gift.go           # Gift database operations
│       └── person.go         # Person database operations
├── migrations/
│   ├── 001_initial_schema.up.sql    # Schema creation
│   └── 001_initial_schema.down.sql  # Schema rollback
├── scripts/
│   ├── migrate.sh            # Migration script
│   └── seed.go               # Database seeding
├── go.mod                    # Go module definition
└── README.md                 # This file
```

## API Endpoints

See [API Documentation](../docs/api.md) for detailed endpoint information.

### Quick Reference

**People:**

- `GET /api/people` - List all people
- `POST /api/people` - Create a person
- `GET /api/people/:id` - Get person details
- `DELETE /api/people/:id` - Delete a person

**Gifts:**

- `GET /api/people/:personId/gifts` - List gifts for a person
- `POST /api/people/:personId/gifts` - Create a gift for a person
- `PATCH /api/gifts/:id/select` - Mark gift as selected
- `DELETE /api/gifts/:id` - Delete a gift

## Database

The application uses PostgreSQL with the following schema:

- **people** - Gift recipients
- **gifts** - Gift ideas for each person

See [Database Documentation](../docs/database.md) for detailed schema information.

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test ./... -cover

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Building for Production

```bash
# Build binary
go build -o bin/server cmd/server/main.go

# Run production binary
./bin/server
```

The production binary can be used without golang installed.

For production deployment, ensure:

- PostgreSQL is properly configured and secured
- Environment variables are set correctly
- Migrations have been run

## CORS Configuration

If you have any CORS related issue, you may need to adjust CORS settings in `cmd/server/main.go` to restrict allowed origins.

## Troubleshooting

### Database Connection Issues

- Verify PostgreSQL is running: `pg_isready`
- Check connection string format
- Ensure database exists: `psql -l`
- Verify credentials and permissions

### Migration Issues

- Ensure PostgreSQL is accessible
- Check migration files syntax
- Verify `DATABASE_URL` environment variable

### Port Already in Use

Change the `PORT` environment variable or stop the process using port 8080:

```bash
# Find process
lsof -i :8080

# Kill process
kill <PID>
```

## Contributing

When contributing to the backend:

1. Follow Go best practices and idioms
2. Run `go fmt ./...` before committing
3. Run `go vet ./...` to check for issues
4. Add tests for new functionality
5. Update API documentation if endpoints change

## License

See root project README for license information.
