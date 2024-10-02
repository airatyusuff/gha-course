-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stock (
        stock_id SERIAL PRIMARY KEY,
        movie_id INTEGER REFERENCES movies (movie_id),
        copies INTEGER NOT NULL DEFAULT 0
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock;
-- +goose StatementEnd