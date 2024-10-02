package main

import (
	"acme/api"
	"acme/db/mock"
	"acme/model"
	"acme/service"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRootHandler(t *testing.T) {
	//Arrange
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	testResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	// Act
	handler.ServeHTTP(testResponse, req)

	//Assert
	if status := testResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello Airah"
	if testResponse.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", testResponse.Body.String(), expected)
	}
}

// using Mocks
func TestGetUsersHandlerWithMock(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	testResponse := httptest.NewRecorder()
	expected := []model.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Terry"},
	}
	mockRepo := &mock.MockUserRepository{
		MockGetUsers: func() ([]model.User, error) {
			return expected, nil
		},
	}
	userService := service.NewUserService(mockRepo)
	userAPI := api.NewUserAPI(userService)

	handler := http.HandlerFunc(userAPI.GetUsers)
	handler.ServeHTTP(testResponse, req)
	if status := testResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual []model.User
	if err := json.Unmarshal(testResponse.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetProductsHandlerWithMock(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	testResponse := httptest.NewRecorder()
	expected := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10, StockCount: 5},
	}
	mockRepo := &mock.MockProductRepository{
		MockGetProducts: func() ([]model.Product, error) {
			return expected, nil
		},
	}
	productService := service.NewProductService(mockRepo)
	productAPI := api.NewProductAPI(productService)

	handler := http.HandlerFunc(productAPI.GetProducts)
	handler.ServeHTTP(testResponse, req)
	if status := testResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual []model.Product
	if err := json.Unmarshal(testResponse.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

// integration tests
func TestRootHandlerWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer server.Close()

	response, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	defer response.Body.Close()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello Airah"
	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(bodyBytes) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(bodyBytes), expected)
	}
}

// Uses in-memory DB

// func TestGetUsersHandler(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/users", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	testResponse := httptest.NewRecorder()
// 	handler := http.HandlerFunc(userAPI.GetUsers)
// 	expected := []model.User{
// 		{ID: 1, Name: "User 1"},
// 		{ID: 2, Name: "User 2"},
// 		{ID: 3, Name: "User 3"},
// 	}
// 	if err != nil {
// 		t.Fatalf("Failed to marshal expected JSON: %v", err)
// 	}
// 	// Act
// 	handler.ServeHTTP(testResponse, req)
// 	// Assert
// 	if status := testResponse.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// 	var actual []model.User
// 	if err := json.Unmarshal(testResponse.Body.Bytes(), &actual); err != nil {
// 		t.Fatalf("Failed to unmarshal response body: %v", err)
// 	}
// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
// 	}
// }

// func TestGetUsersHandlerWithServer(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(userAPI.GetUsers))
// 	defer server.Close()
// 	response, err := http.Get(server.URL + "/api/users")
// 	if err != nil {
// 		t.Fatalf("Failed to send GET request: %v", err)
// 	}
// 	defer response.Body.Close()
// 	if status := response.StatusCode; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// 	body, errRead := io.ReadAll(response.Body)
// 	if errRead != nil {
// 		t.Fatalf("Error reading response body: %v", err)
// 	}
// 	expected := []model.User{
// 		{ID: 1, Name: "User 1"},
// 		{ID: 2, Name: "User 2"},
// 		{ID: 3, Name: "User 3"},
// 	}
// 	var actual []model.User
// 	if err := json.Unmarshal(body, &actual); err != nil {
// 		t.Fatalf("Failed to unmarshal response body: %v", err)
// 	}
// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
// 	}
// }
