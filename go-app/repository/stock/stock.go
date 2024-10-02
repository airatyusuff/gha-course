package stock

import (
	"acme/model"
)

type StockRepository interface {
	GetStock() ([]model.Stock, error)
	AddStock(s model.Stock) (id int, err error)
}
