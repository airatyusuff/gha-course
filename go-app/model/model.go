package model

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Product struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Price      float64 `json:"price" db:"price"`
	StockCount int     `json:"stock_count" db:"stock_count"`
}

type Movie struct {
	ID       int    `json:"movie_id" db:"movie_id"`
	Title    string `json:"title" db:"title"`
	Year     int    `json:"year" db:"year"`
	Genre    string `json:"genre" db:"genre"`
	Director string `json:"director" db:"director"`
	Rating   int    `json:"rating" db:"rating"`
}

type Stock struct {
	ID      int `json:"stock_id" db:"stock_id"`
	Copies  int `json:"copies" db:"copies"`
	MovieID int `json:"movie_id" db:"movie_id"`
}

type ExternalProduct struct {
	Label    string  `json:"label"`
	Received string  `json:"received"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Cost     float64 `json:"cost"`
	Paid     bool    `json:"paid"`
	DatePaid *string `json:"datePaid"`
	Img      string  `json:"img"`
}
