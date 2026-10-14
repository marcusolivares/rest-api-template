CREATE TABLE IF NOT EXISTS calculator_transactions (
    transaction_uuid UUID           NOT NULL,
    value1           NUMERIC(16, 2) NOT NULL,
    value2           NUMERIC(16, 2) NOT NULL,
    operation        TEXT           NOT NULL,   -- 'ADD' | 'MULTIPLY'
    result           NUMERIC(16, 2) NOT NULL,
    created_at       TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    created_by       TEXT           NOT NULL DEFAULT 'admin',
    deleted          BOOLEAN        NOT NULL DEFAULT FALSE,
    deleted_at       TIMESTAMPTZ             DEFAULT NULL,
    updated_at       TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_by       TEXT           NOT NULL DEFAULT 'admin',
    CONSTRAINT calculator_transactions_pkey PRIMARY KEY (transaction_uuid)
);

COMMENT ON TABLE  calculator_transactions             IS 'Audit log of every calculator computation.';
COMMENT ON COLUMN calculator_transactions.operation   IS 'Allowed values: ADD, MULTIPLY';
