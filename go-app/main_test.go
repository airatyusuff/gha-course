package main

import (
	"net/http"
	"net/http/httptest"
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
