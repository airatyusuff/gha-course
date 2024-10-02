package product

import (
	"acme/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresProductRepository struct {
	DB *sqlx.DB
}

func NewPostgresProductRepository(db *sqlx.DB) *PostgresProductRepository {
	return &PostgresProductRepository{DB: db}
}

func (repo *PostgresProductRepository) GetProducts() ([]model.Product, error) {
	products := []model.Product{}

	err := sqlx.Select(repo.DB, &products, "SELECT * FROM products")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.Product{}, errors.New("could not get all products")
	}

	return products, nil
}

func (repo *PostgresProductRepository) AddProduct(p model.Product) (id int, err error) {
	err = repo.DB.QueryRow(
		"INSERT INTO products (name, price, stock_count) VALUES ($1,$2,$3) RETURNING id",
		p.Name, p.Price, p.StockCount).Scan(&id)

	if err != nil {
		fmt.Println("error inserting product into the database:", err)
		return 0, errors.New("could not add new product")
	}

	return id, nil
}
