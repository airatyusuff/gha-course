package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//using a multiplexer instead
	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)

	fmt.Println("server on port 8080")
	err := http.ListenAndServe(":8080", CorsMiddleware(router))

	if err != nil {
		fmt.Println("error starting server", err)
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving route /")
	io.WriteString(w, "Hello Airah")
}
