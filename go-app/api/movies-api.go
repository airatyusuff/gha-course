package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type MovieAPI struct {
	movieService *service.MovieService
}

func NewMovieAPI(movieService *service.MovieService) *MovieAPI {
	return &MovieAPI{movieService: movieService}
}

func (api *MovieAPI) GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET method: /api/movies")

	result, err := api.movieService.GetMovies()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (api *MovieAPI) AddNewMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST method: /api/movies")

	var m model.Movie
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := api.movieService.AddMovie(m)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Movie added successfully: %d", id)
}
