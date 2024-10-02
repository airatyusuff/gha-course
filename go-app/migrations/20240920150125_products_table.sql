-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) NOT NULL,
        price NUMERIC NOT NULL,
        stock_count INTEGER NOT NULL DEFAULT 0
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd