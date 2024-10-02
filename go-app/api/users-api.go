package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type UserAPI struct {
	userService *service.UserService
}

func NewUserAPI(userService *service.UserService) *UserAPI {
	return &UserAPI{userService: userService}
}

func (api *UserAPI) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET method: /api/users")

	users, err := api.userService.GetUsers()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (api *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST method: /api/users")

	newUser, err := decodeUser(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := api.userService.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully: %d", id)
}

func (api *UserAPI) GetSingleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET method: /api/users/{userId}")

	userId := r.PathValue("id")
	idNum, err := parseId(userId)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fetchedUser, serviceErr := api.userService.GetSingleUser(idNum)
	if serviceErr != nil {
		http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(fetchedUser)
}

func (api *UserAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE method: /api/users/{userId}")

	userId := r.PathValue("id")
	idNum, err := parseId(userId)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	serviceErr := api.userService.DeleteUser(idNum)
	if serviceErr != nil {
		http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User with ID %d deleted successfully", idNum)
}

func (api *UserAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT method: /api/users/{userId}")

	userId := r.PathValue("id")
	idNum, err := parseId(userId)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var updatedDetails model.User
	updatedDetails, decodeErr := decodeUser(r.Body)
	if decodeErr != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	serviceErr := api.userService.UpdateUser(idNum, updatedDetails)
	if serviceErr != nil {
		http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User updated successfully: %d", idNum)
}

func parseId(idStr string) (idNum int, err error) {
	idNum, err = strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return 0, err
	}

	return idNum, nil

}

func decodeUser(body io.ReadCloser) (user model.User, err error) {
	err = json.NewDecoder(body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return model.User{}, err
	}

	return user, nil
}
