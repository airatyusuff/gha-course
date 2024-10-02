package movie

import (
	"acme/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresMovieRepository struct {
	DB *sqlx.DB
}

func NewPostgresMovieRepository(db *sqlx.DB) *PostgresMovieRepository {
	return &PostgresMovieRepository{DB: db}
}

func (repo *PostgresMovieRepository) GetMovies() ([]model.Movie, error) {
	movies := []model.Movie{}

	err := sqlx.Select(repo.DB, &movies, "SELECT * FROM movies")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.Movie{}, errors.New("could not get all movies")
	}

	return movies, nil
}

func (repo *PostgresMovieRepository) AddMovie(m model.Movie) (id int, err error) {
	err = repo.DB.QueryRow(
		"INSERT INTO movies (title, year, genre, director, rating) VALUES ($1,$2,$3,$4,$5) RETURNING movie_id",
		m.Title, m.Year, m.Genre, m.Director, m.Rating).Scan(&id)

	if err != nil {
		fmt.Println("error inserting movie into the database:", err)
		return 0, errors.New("could not add new movie")
	}

	return id, nil
}
