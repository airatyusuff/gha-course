package stock

import (
	"acme/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresStockRepository struct {
	DB *sqlx.DB
}

func NewPostgresStockRepository(db *sqlx.DB) *PostgresStockRepository {
	return &PostgresStockRepository{DB: db}
}

func (repo *PostgresStockRepository) GetStock() ([]model.Stock, error) {
	result := []model.Stock{}

	err := sqlx.Select(repo.DB, &result, "SELECT * FROM stock")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.Stock{}, errors.New("could not get all")
	}

	return result, nil
}

func (repo *PostgresStockRepository) AddStock(p model.Stock) (id int, err error) {
	err = repo.DB.QueryRow(
		"INSERT INTO stock (movie_id, copies) VALUES ($1,$2) RETURNING stock_id",
		p.MovieID, p.Copies).Scan(&id)

	if err != nil {
		fmt.Println("error inserting movie stock into the database:", err)
		return 0, errors.New("could not add new movie stock")
	}

	return id, nil
}
