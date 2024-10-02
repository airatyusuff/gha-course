package user

import (
	"acme/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" //postgres driver for the database/sql package in Go
)

type PostgresUserRepository struct {
	DB *sqlx.DB // DB is a struct type in the sql package
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (repo *PostgresUserRepository) GetUsers() ([]model.User, error) {
	users := []model.User{}

	err := sqlx.Select(repo.DB, &users, "SELECT * FROM users")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.User{}, errors.New("could not get all users")
	}

	return users, nil
}

func (repo *PostgresUserRepository) AddUser(user model.User) (id int, err error) {
	err = repo.DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&id)
	if err != nil {
		fmt.Println("error inserting user into the database:", err)
		return 0, errors.New("could not add new user")
	}

	return id, nil
}

func (repo *PostgresUserRepository) GetUser(id int) (model.User, error) {
	var user model.User

	err := repo.DB.QueryRow("SELECT id, name FROM users WHERE id =$1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return model.User{}, errors.New("user id not found")
	}

	return user, nil
}

func (repo *PostgresUserRepository) DeleteUser(id int) error {
	_, err := repo.DB.Exec("DELETE FROM users WHERE id =$1", id)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return errors.New("could not delete user with no id found")
	}

	return nil
}

func (repo *PostgresUserRepository) UpdateUser(id int, body model.User) error {
	_, err := repo.DB.Exec("UPDATE users SET name = ($1) WHERE id =$2", body.Name, id)

	if err != nil {
		fmt.Println("Error querying the database:", err)
		return errors.New("could not update user")
	}

	return nil
}
