package service

import (
	"acme/model"
	"acme/repository/movie"
	"errors"
	"fmt"
)

type MovieService struct {
	repository movie.MoviesRepository
}

func NewMovieService(repo movie.MoviesRepository) *MovieService {
	return &MovieService{repository: repo}
}

func (s *MovieService) GetMovies() ([]model.Movie, error) {
	products, err := s.repository.GetMovies()

	if err != nil {
		fmt.Println("error getting movies from db:", err)
		return nil, errors.New("there was an error getting the movies from the database")
	}

	return products, nil
}

func (s *MovieService) AddMovie(m model.Movie) (int, error) {
	id, err := s.repository.AddMovie(m)

	if err != nil {
		fmt.Println("error adding new movie to db:", err)
		return -1, errors.New("there was an error adding new movie to the database")
	}

	return id, nil
}
