-- +goose Up
-- +goose StatementBegin
CREATE TYPE account_status_enum AS ENUM ('active', 'inactive', 'suspended');

CREATE TABLE user_account (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(24) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,  
    is_email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    account_status account_status_enum DEFAULT 'active'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_account;
DROP TYPE account_status_enum;
-- +goose StatementEnd
