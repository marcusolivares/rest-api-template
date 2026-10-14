# internal/rest/middleware

Custom chi-compatible HTTP middlewares.

| File          | Middleware        | HTTP status | Purpose                                                         |
|---------------|-------------------|-------------|-----------------------------------------------------------------|
| `apikey.go`   | `APIKey`          | 401         | Validates the `api-key` request header                          |
| `logger.go`   | `ZerologLogger`   | —           | Logs method, path, status code, and latency for every request   |
| `recover.go`  | `PanicRecovery`   | 500         | Catches unhandled panics and returns a JSON 500 error body      |

## Registration order (in `internal/rest/server/server.go`)

```go
router.Use(
    chimw.RequestID,       // inject X-Request-Id
    middleware.ZerologLogger(logger),
    middleware.PanicRecovery(logger),
)
// API-key validation is applied per route group, not globally:
router.Route("/calculator/v1", func(r chi.Router) {
    r.Use(middleware.APIKey(cfg.APIKey))
    // ...
})
```
