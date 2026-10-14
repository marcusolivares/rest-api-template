# cmd/http

Entry point for the Calculator HTTP server.

## What it does

1. Calls `config.LoadFromEnv()` to parse all environment variables.
2. Opens a PostgreSQL connection pool via `pgstore.NewPool`.
3. Delegates to `server.Serve` which sets up routes and blocks until a
   shutdown signal is received.

## Run locally

```bash
# Copy the example env file and edit as needed
cp .env.example .env

# Source environment variables (or use `direnv`)
set -a; source .env; set +a

# Run
go run ./cmd/http
```

## Build

```bash
go build -o calculator-service ./cmd/http
./calculator-service
```
