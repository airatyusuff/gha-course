package service

import (
	"acme/model"
	"acme/repository/product"
	"errors"
	"fmt"
)

type ProductService struct {
	repository product.ProductRepository
}

func NewProductService(repo product.ProductRepository) *ProductService {
	return &ProductService{repository: repo}
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	products, err := s.repository.GetProducts()

	if err != nil {
		fmt.Println("error getting products from db:", err)
		return nil, errors.New("there was an error getting the products from the database")
	}

	return products, nil
}

func (s *ProductService) AddProduct(p model.Product) (int, error) {
	id, err := s.repository.AddProduct(p)

	if err != nil {
		fmt.Println("error adding new product to db:", err)
		return -1, errors.New("there was an error adding new product to the database")
	}

	return id, nil
}
