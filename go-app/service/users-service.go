package service

import (
	"acme/model"
	"acme/repository/user"
	"errors"
	"fmt"
)

type UserService struct {
	repository user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	users, err := s.repository.GetUsers()

	if err != nil {
		fmt.Println("error getting users from db:", err)
		return nil, errors.New("there was an error getting the users from the database")
	}

	return users, nil
}

func (s *UserService) CreateUser(newUser model.User) (int, error) {
	id, err := s.repository.AddUser(newUser)

	if err != nil {
		fmt.Println("error adding new user to db:", err)
		return -1, errors.New("there was an error adding new user to the database")
	}

	return id, nil
}

func (s *UserService) GetSingleUser(id int) (model.User, error) {
	fetchedUser, err := s.repository.GetUser(id)
	if err != nil {
		fmt.Println("error fetching user from db:", err)
		return model.User{}, errors.New("error fetching user")
	}

	return fetchedUser, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)

	if err != nil {
		fmt.Println("error deleting user:", err)
		return errors.New("error deleting user")
	}

	return nil
}

func (s *UserService) UpdateUser(id int, body model.User) error {
	err := s.repository.UpdateUser(id, body)

	if err != nil {
		fmt.Println("error updating user:", err)
		return errors.New("error updating user")
	}

	return nil
}
