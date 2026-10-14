# resources/gen

Generated artefacts package.  **Do not edit files in this package manually.**

## Files

| File                   | How produced             | Purpose                                       |
|------------------------|--------------------------|-----------------------------------------------|
| `embed.go`             | Hand-written (once)      | Embeds `openapi/openapi.json` into the binary |
| `openapi/openapi.json` | `make openapi` (Redocly) | Bundled, single-file OpenAPI 3.0 JSON spec    |

## Regenerating

```bash
make openapi
```

This runs:
1. `redocly bundle resources/openapi/openapi.yml → resources/gen/openapi/openapi.json`

Commit the regenerated file alongside any spec changes.
