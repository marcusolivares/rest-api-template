# resources/pgsql/migrations

SQL migration files managed by [golang-migrate](https://github.com/golang-migrate/migrate).

## Naming convention

```
YYYYMMDDHHMMSS_<short_description>.up.sql
YYYYMMDDHHMMSS_<short_description>.down.sql
```

Example:
```
20260521000001_create_calculator_transactions.up.sql
20260521000001_create_calculator_transactions.down.sql
```

## File format

Each migration consists of two files:

**`*.up.sql`** — forward migration SQL:
```sql
CREATE TABLE IF NOT EXISTS example (...);
```

**`*.down.sql`** — rollback SQL (required; leave empty if tables must not be dropped):
```sql
-- REQUIRED, BUT LEAVE IT EMPTY. WE SHOULD NOT DROP THE TABLES.
```

## Running migrations locally

```bash
# Install golang-migrate with PostgreSQL support (once)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

export DATABASE_URL="postgres://postgres:postgres@localhost:5432/calculator?sslmode=disable"

migrate -path ./resources/pgsql/migrations -database "$DATABASE_URL" up      # apply all pending
migrate -path ./resources/pgsql/migrations -database "$DATABASE_URL" version  # show current version
```
