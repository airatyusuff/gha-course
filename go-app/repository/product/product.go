package product

import (
	"acme/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	AddProduct(p model.Product) (id int, err error)
}
