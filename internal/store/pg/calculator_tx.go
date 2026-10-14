package pg

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// CalculatorTxStore handles persistence for the calculator_transactions table.
type CalculatorTxStore struct {
	pool *pgxpool.Pool
}

// NewCalculatorTxStore returns a new CalculatorTxStore backed by pool.
func NewCalculatorTxStore(pool *pgxpool.Pool) *CalculatorTxStore {
	return &CalculatorTxStore{pool: pool}
}

// CreateTxParams carries the inputs required to insert a new transaction row.
type CreateTxParams struct {
	Value1    float64
	Value2    float64
	Operation string // "ADD" | "MULTIPLY"
	Result    float64
}

// Create inserts a new calculator_transactions row.
// The transaction-uuid is generated here and returned to the caller.
// created_by and updated_by are hard-coded to 'admin' as per business rules.
func (s *CalculatorTxStore) Create(ctx context.Context, p CreateTxParams) (uuid.UUID, error) {
	txUUID := uuid.New()

	_, err := s.pool.Exec(ctx, `
		INSERT INTO calculator_transactions (
			transaction_uuid, value1, value2, operation, result,
			created_by, updated_by
		) VALUES (
			$1, $2, $3, $4, $5,
			'admin', 'admin'
		)
	`, txUUID, p.Value1, p.Value2, p.Operation, p.Result)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CalculatorTxStore.Create: %w", err)
	}

	return txUUID, nil
}
