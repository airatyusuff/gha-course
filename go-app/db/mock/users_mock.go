package mock

import (
	"acme/model"
)

type MockUserRepository struct {
	MockGetUsers   func() ([]model.User, error)
	MockGetUser    func(id int) (model.User, error)
	MockAddUser    func(user model.User) (int, error)
	MockUpdateUser func(id int, user model.User) error
	MockDeleteUser func(id int) error
}

func (m *MockUserRepository) GetUsers() ([]model.User, error) {
	return m.MockGetUsers()
}

func (m *MockUserRepository) GetUser(id int) (model.User, error) {
	return m.MockGetUser(id)
}

func (m *MockUserRepository) AddUser(user model.User) (int, error) {
	return m.MockAddUser(user)
}

func (m *MockUserRepository) UpdateUser(id int, user model.User) error {
	return m.MockUpdateUser(id, user)
}

func (m *MockUserRepository) DeleteUser(id int) error {
	return m.MockDeleteUser(id)
}
