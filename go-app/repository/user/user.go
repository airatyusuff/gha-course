package user

import (
	"acme/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	AddUser(user model.User) (id int, err error)
	GetUser(id int) (model.User, error)
	UpdateUser(id int, user model.User) error
	DeleteUser(id int) error
}
