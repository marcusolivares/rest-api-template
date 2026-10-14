# internal/rest/handler/calculator

HTTP handlers for the Calculator API (`/calculator/v1`).

## Endpoints

| File          | Method | Route                          | Description                   |
|---------------|--------|--------------------------------|-------------------------------|
| `add.go`      | POST   | `/calculator/v1/add`           | Adds `value1` and `value2`    |
| `multiply.go` | POST   | `/calculator/v1/multiply`      | Multiplies `value1` by `value2` |

## Request / Response

**Request body** (same for both endpoints):

```json
{ "value1": 25.5, "value2": 10.0 }
```

**Success — 200**:

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "transaction-uuid": "550e8400-e29b-41d4-a716-446655440000",
    "value1": 25.5,
    "value2": 10.0,
    "operation": "ADD",
    "result": 35.5
  }
}
```

## HTTP Status Codes

| Status | Condition                                              |
|--------|--------------------------------------------------------|
| 200    | Successful computation and persisted transaction       |
| 400    | Either operand is outside `[-1000, 1000]`              |
| 401    | `api-key` header missing or incorrect                  |
| 500    | Unexpected storage error                               |

The `api-key` header is validated by the `middleware.APIKey` middleware registered
at the route group level in `internal/rest/server/server.go`.
