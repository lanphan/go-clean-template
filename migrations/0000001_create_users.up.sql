CREATE TABLE IF NOT EXISTS users
(
    id             uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_address VARCHAR(42) NOT NULL,
    reward         int NOT NULL,
    claim_status   boolean NOT NULL,
    created_at     TIMESTAMP NOT NULL
);

CREATE INDEX idx_users_address ON users(wallet_address);
