package service

import (
	"acme/model"
	"acme/repository/stock"
	"errors"
	"fmt"
)

type StockService struct {
	repository stock.StockRepository
}

func NewStockService(repo stock.StockRepository) *StockService {
	return &StockService{repository: repo}
}

func (s *StockService) GetAllStock() ([]model.Stock, error) {
	results, err := s.repository.GetStock()

	if err != nil {
		fmt.Println("error getting movie stock from db:", err)
		return nil, errors.New("there was an error getting the movies stock from the database")
	}

	return results, nil
}

func (s *StockService) AddNewStock(m model.Stock) (int, error) {
	id, err := s.repository.AddStock(m)

	if err != nil {
		fmt.Println("error adding new stock to db:", err)
		return -1, errors.New("there was an error adding new movie stock to the database")
	}

	return id, nil
}
