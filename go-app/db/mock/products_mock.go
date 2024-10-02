package mock

import (
	"acme/model"
)

type MockProductRepository struct {
	MockGetProducts func() ([]model.Product, error)
	MockAddProduct  func(p model.Product) (int, error)
}

func (m *MockProductRepository) GetProducts() ([]model.Product, error) {
	return m.MockGetProducts()
}

func (m *MockProductRepository) AddProduct(p model.Product) (int, error) {
	return m.MockAddProduct(p)
}
