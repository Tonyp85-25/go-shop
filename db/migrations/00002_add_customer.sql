-- +goose Up
CREATE TABLE customers(
    id SERIAL PRIMARY KEY,
    public_id UUID NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    first_name VARCHAR (100) NOT NULL,
    last_name VARCHAR (100) NOT NULL,
    phone VARCHAR (20),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_customers_email ON customers (email);
CREATE INDEX idx_customers_deleted_at ON customers (deleted_at);

-- +goose Down
DROP TABLE IF EXISTS customers;
DROP INDEX IF EXISTS idx_customers_deleted_at;
DROP INDEX IF EXISTS idx_customers_email;
