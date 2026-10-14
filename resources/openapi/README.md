# resources/openapi

OpenAPI 3.0 source specification for the Calculator Service API.

## Structure

```
openapi.yml                    ← root spec; references schema files via $ref
schema/
  common/
    error.yaml                 ← ErrorResponse — used by all 4xx/5xx responses
  healthcheck/
    healthcheck.yaml           ← HealthcheckResponse
  calculator/
    add.yaml                   ← AddRequest, AddData, AddResponse
    multiply.yaml              ← MultiplyRequest, MultiplyData, MultiplyResponse
```

## Workflow

The source YAML files are **authoring** files only.  They are bundled into a
single `resources/gen/openapi/openapi.json` by running:

```bash
make openapi
```

This requires [Redocly CLI](https://redocly.com/docs/cli/) (the Makefile will
install it automatically via npm if missing).
