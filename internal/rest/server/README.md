# internal/rest/server

Constructs the chi router, registers routes and middlewares, and runs the HTTP
server with graceful shutdown.
Each HTTP endpoint must be in `lower-kebab-case` format.

## Entry point

```go
server.Serve(cfg, logger, pool)
```

`Serve` blocks until a `SIGINT` or `SIGTERM` is received, then drains in-flight
requests up to `HTTP_SERVER_SHUTDOWN_TIMEOUT` before returning.

## Route table

| Method | Path                       | Auth required | Handler                          |
|--------|----------------------------|---------------|----------------------------------|
| GET    | `/healthcheck`             | No            | `handler/healthcheck.Get`        |
| POST   | `/calculator/v1/add`       | `api-key`     | `handler/calculator.Add`         |
| POST   | `/calculator/v1/multiply`  | `api-key`     | `handler/calculator.Multiply`    |

## Dependencies

| Package                      | Purpose                                      |
|------------------------------|----------------------------------------------|
| `internal/rest/handler/`     | HTTP handler factories, grouped by feature   |
| `internal/rest/middleware/`  | Reusable chi-compatible middleware functions |
