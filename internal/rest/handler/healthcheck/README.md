# internal/rest/handler/healthcheck

HTTP handler for the health check endpoint (`/healthcheck`).

## Endpoints

| File               | Method | Route           | Auth | Description                   |
|--------------------|--------|-----------------|------|-------------------------------|
| `healthcheck.go`   | GET    | `/healthcheck`  | No   | Returns service health status |

## Response

**Success — 200:**

```json
{"status":"ok"}
```

The response is intentionally a plain JSON object — no envelope wrapper — so that
infrastructure probes (load balancers, Kubernetes liveness/readiness checks) can
parse it without knowledge of the application's `{code, message, data}` envelope.

## HTTP Status Codes

| Status | Condition          |
|--------|--------------------|
| 200    | Service is healthy |

No authentication is required. This endpoint is registered outside the `api-key`
middleware group in `internal/rest/server/server.go`.
