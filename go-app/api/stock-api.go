package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type StockAPI struct {
	stockService *service.StockService
}

func NewStockAPI(s *service.StockService) *StockAPI {
	return &StockAPI{stockService: s}
}

func (api *StockAPI) GetAllStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET method: /api/stock")

	result, err := api.stockService.GetAllStock()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (api *StockAPI) AddNewMovieStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST method: /api/stock")

	var m model.Stock
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := api.stockService.AddNewStock(m)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Movie stock added successfully: %d", id)
}
