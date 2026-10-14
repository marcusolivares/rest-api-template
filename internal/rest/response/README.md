# internal/rest/response

Provides a standard JSON response envelope used by all HTTP handlers.

## Response shape

```json
{
  "code": 200,
  "message": "OK",
  "data": { "...": "..." }
}
```

For errors the `data` field is omitted:

```json
{
  "code": 401,
  "message": "unauthorized: invalid or missing api-key"
}
```

## API

| Function     | When to use                                        |
|--------------|----------------------------------------------------|
| `WriteJSON`  | 2xx success responses; wraps a data payload        |
| `WriteError` | 4xx / 5xx error responses; carries a plain message |
