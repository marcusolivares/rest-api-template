# internal/store/pg

PostgreSQL data-access layer using [jackc/pgx v5](https://github.com/jackc/pgx).

## Files

| File                  | Purpose                                                        |
|-----------------------|----------------------------------------------------------------|
| `pool.go`             | `NewPool` — creates and health-checks a `pgxpool.Pool`         |
| `calculator_tx.go`    | `CalculatorTxStore` — CRUD for `calculator_transactions` table |

## Connection

The pool is configured via `pg.Config` which is populated from `internal/config.EnvPostgreSQL`.  
The DSN format used is:

```
postgres://user:pass@host:port/dbname?sslmode=disable&search_path=public
```
