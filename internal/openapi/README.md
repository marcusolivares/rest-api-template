# internal/openapi

Provides HTTP handlers for serving the OpenAPI specification and Redoc UI,
without any external UI framework dependency.

## How it works

| Handler        | Route                | What it serves                                          |
|----------------|----------------------|---------------------------------------------------------|
| `SpecHandler`  | `GET /openapi.json`  | Raw OpenAPI 3.0 JSON (read from embedded FS at startup) |
| `UIHandler`    | `GET /redoc`         | Redoc HTML page (Redoc loaded from unpkg CDN)           |

The Redoc page loads the spec from `/openapi.json` in the browser, so no
proxy or server-side configuration is required.

## Enabling the UI

Set `OPENAPI_UI_ENABLE=true` in environment variables (or `.env`):

```env
OPENAPI_UI_ENABLE=true
OPENAPI_UI_PATH=/redoc   # optional, this is the default
```

Then open <http://localhost:8080/redoc> in your browser.

## Replacing the bundled spec

Run `make openapi` whenever the source YAML spec changes.  The updated
`resources/gen/openapi/openapi.json` is embedded at compile time via
`resources/gen/embed.go`.
