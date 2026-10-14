# internal/rest/handler

HTTP handler functions, grouped by domain.

Each sub-package corresponds to one resource or feature area and contains only `http.HandlerFunc` factories — no routing or middleware logic.

Each field in request/response body, and query parameter must be written as `lower-kebab-case` format.

| Sub-package      | Routes prefix          |
|------------------|------------------------|
| `healthcheck/`   | `/healthcheck`         |
| `calculator/`    | `/calculator/v1`       |
