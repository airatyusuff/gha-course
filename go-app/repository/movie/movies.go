package movie

import (
	"acme/model"
)

type MoviesRepository interface {
	GetMovies() ([]model.Movie, error)
	AddMovie(p model.Movie) (id int, err error)
}
