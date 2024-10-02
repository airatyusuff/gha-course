package main

import (
	"acme/api"
	"acme/config"
	"acme/db/postgres"
	"acme/repository/movie"
	"acme/repository/product"
	"acme/repository/stock"
	"acme/repository/user"
	"acme/service"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// config := config.LoadDatabaseConfig(".env.inmemory")
	config := config.LoadDatabaseConfig()
	db, dbErr := initialiseDatabase(config)

	if dbErr != nil {
		fmt.Println("Error initializing the database:", dbErr)
	}
	defer db.DB.Close()

	// initialise user services and api layer
	userRepo := user.NewPostgresUserRepository(db.DB)
	userService := service.NewUserService(userRepo)
	userAPI := api.NewUserAPI(userService)

	// initialise products repository here
	productRepo := product.NewPostgresProductRepository(db.DB)
	productService := service.NewProductService(productRepo)
	productAPI := api.NewProductAPI(productService)

	movieRepo := movie.NewPostgresMovieRepository(db.DB)
	movieService := service.NewMovieService(movieRepo)
	movieAPI := api.NewMovieAPI(movieService)

	stockRepo := stock.NewPostgresStockRepository(db.DB)
	stockService := service.NewStockService(stockRepo)
	stockAPI := api.NewStockAPI(stockService)

	//using a multiplexer instead
	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", userAPI.GetUsers)
	router.HandleFunc("POST /api/users", userAPI.CreateUser)
	router.HandleFunc("GET /api/users/{id}", userAPI.GetSingleUser)
	router.HandleFunc("DELETE /api/users/{id}", userAPI.DeleteUser)
	router.HandleFunc("PUT /api/users/{id}", userAPI.UpdateUser)
	router.HandleFunc("GET /api/products", productAPI.GetProducts)
	router.HandleFunc("POST /api/products", productAPI.AddNewProduct)
	router.HandleFunc("GET /api/movies", movieAPI.GetMovies)
	router.HandleFunc("POST /api/movies", movieAPI.AddNewMovie)
	router.HandleFunc("GET /api/stock", stockAPI.GetAllStock)
	router.HandleFunc("POST /api/stock", stockAPI.AddNewMovieStock)

	fmt.Println("server on port 8080")
	err := http.ListenAndServe(":8080", CorsMiddleware(router))

	if err != nil {
		fmt.Println("error starting server", err)
	}
}

func initialiseDatabase(config config.DatabaseConfig) (*postgres.Postgres, error) {
	switch config.Type {
	case "postgres":
		connectionString := fmt.Sprintf(
			"dbname=%s user=%s password=%s host=%s sslmode=%s",
			config.DBName, config.User, config.Password, config.Host, config.SSLMode,
		)
		db, err := postgres.PostgresConnection(connectionString)
		if err != nil {
			panic(err)
		}
		return db, nil
	// case "inmemory":
	// 	return user.NewInMemoryUserRepository(), nil
	default:
		return nil, fmt.Errorf("unsupported db type: %s", config.Type)
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

// func simple_main() {
// 	http.HandleFunc("/", rootHandler)
// 	http.HandleFunc("/api/users", getUsers)

// 	fmt.Println("server on port 8080")
// 	err := http.ListenAndServe(":8080", nil)

// 	if err != nil {
// 		fmt.Println("error starting server", err)
// 	}
// }
