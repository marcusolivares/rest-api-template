# internal/config

Parses application configuration from environment variables using [caarlos0/env v11](https://github.com/caarlos0/env).

## Files

| File        | Purpose                                              |
|-------------|------------------------------------------------------|
| `config.go` | Root `Env` struct, `LoadFromEnv()`, logger builder   |
| `http.go`   | `EnvHTTP` — HTTP server timeouts and bind address    |
| `pgsql.go`  | `EnvPostgreSQL` — PostgreSQL connection pool settings |

## Usage

```go
cfg := config.LoadFromEnv()
logger := cfg.Logger()
```

All fields have `envDefault` values so the service starts with sane defaults
without requiring a `.env` file (see `.env.example` at the repo root).
