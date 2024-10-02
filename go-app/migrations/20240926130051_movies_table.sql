-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movies (
        movie_id SERIAL PRIMARY KEY,
        title VARCHAR(100) NOT NULL,
        year NUMERIC NOT NULL,
        genre VARCHAR(50) NOT NULL,
        director VARCHAR(100) NOT NULL,
        rating INTEGER
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movies;
-- +goose StatementEnd