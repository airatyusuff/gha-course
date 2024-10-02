package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProductAPI struct {
	productService *service.ProductService
}

func NewProductAPI(productService *service.ProductService) *ProductAPI {
	return &ProductAPI{productService: productService}
}

func (api *ProductAPI) GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET method: /api/products")
	products, err := api.productService.GetProducts()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (api *ProductAPI) AddNewProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST method: /api/products")

	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := api.productService.AddProduct(newProduct)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product added successfully: %d", id)
}

func (api *ProductAPI) GetProductsFromExternal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET method: /api/products")
	products, err := fetchProductsFromExternal()
	if err != nil {
		fmt.Println("Error fetching products, returning dummy data:", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dummyProducts)
		return
	}

	formattedProducts := make([]map[string]interface{}, len(products))
	for i, product := range products {
		formattedProducts[i] = map[string]interface{}{
			"id":       i + 1,
			"label":    product.Label,
			"cost":     product.Cost,
			"received": product.Received,
			"type":     product.Type,
			"quantity": product.Quantity,
			"paid":     product.Paid,
			"datePaid": product.DatePaid,
			"img":      product.Img,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formattedProducts)
}

func fetchProductsFromExternal() ([]model.ExternalProduct, error) {
	resp, err := http.Get("https://acme2inventory.azurewebsites.net/api/inventory")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var products []model.ExternalProduct
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

var dummyProducts = []map[string]interface{}{
	{
		"label":    "Anchovy Fillets",
		"received": "2023-07-20T00:00:00",
		"type":     "Consumer Services",
		"quantity": 493,
		"cost":     11523.90,
		"paid":     false,
		"datePaid": nil,
		"id":       1,
		"img":      "https://picsum.photos/id/23/3887/4899",
	},
	{
		"label":    "Anchovy In Oil",
		"received": "2022-09-08T00:00:00",
		"type":     "Consumer Services",
		"quantity": 778,
		"cost":     9228.09,
		"paid":     true,
		"datePaid": nil,
		"id":       2,
		"img":      "https://picsum.photos/id/20/3670/2462",
	},
	{
		"label":    "Appetiser - Bought",
		"received": "2022-03-20T00:00:00",
		"type":     "Consumer Services",
		"quantity": 822,
		"cost":     2611.30,
		"paid":     true,
		"datePaid": nil,
		"id":       3,
		"img":      "https://picsum.photos/id/22/4434/3729",
	},
	{
		"label":    "Appetizer - Asian Shrimp Roll",
		"received": "2023-06-11T00:00:00",
		"type":     "Energy",
		"quantity": 750,
		"cost":     2912.21,
		"paid":     false,
		"datePaid": nil,
		"id":       4,
		"img":      "https://picsum.photos/id/16/2500/1667",
	},
	{
		"label":    "Appetizer - Assorted Box",
		"received": "2021-01-28T00:00:00",
		"type":     "Health Care",
		"quantity": 269,
		"cost":     11120.81,
		"paid":     true,
		"datePaid": nil,
		"id":       5,
		"img":      "https://picsum.photos/id/22/4434/3729",
	},
	{
		"label":    "Appetizer - Chicken Satay",
		"received": "2020-03-08T00:00:00",
		"type":     "n/a",
		"quantity": 141,
		"cost":     5212.90,
		"paid":     false,
		"datePaid": nil,
		"id":       6,
		"img":      "https://picsum.photos/id/23/3887/4899",
	},
	{
		"label":    "Appetizer - Crab And Brie",
		"received": "2021-05-07T00:00:00",
		"type":     "Technology",
		"quantity": 362,
		"cost":     1375.75,
		"paid":     true,
		"datePaid": nil,
		"id":       7,
		"img":      "https://picsum.photos/id/20/3670/2462",
	},
	{
		"label":    "Appetizer - Lobster Phyllo Roll",
		"received": "2021-03-16T00:00:00",
		"type":     "Finance",
		"quantity": 175,
		"cost":     8413.00,
		"paid":     false,
		"datePaid": nil,
		"id":       8,
		"img":      "https://picsum.photos/id/16/2500/1667",
	},
	{
		"label":    "Appetizer - Mango Chevre",
		"received": "2020-09-10T00:00:00",
		"type":     "Health Care",
		"quantity": 256,
		"cost":     9317.12,
		"paid":     true,
		"datePaid": nil,
		"id":       9,
		"img":      "https://picsum.photos/id/22/4434/3729",
	},
	{
		"label":    "Appetizer - Seafood Assortment",
		"received": "2021-03-19T00:00:00",
		"type":     "n/a",
		"id":       10,
		"quantity": 287,
		"cost":     8388.80,
		"paid":     false,
		"datePaid": nil,
		"img":      "https://picsum.photos/id/23/3887/4899",
	},
	{
		"cost":     2018.36,
		"datePaid": nil,
		"id":       11,
		"label":    "Apple - Delicious  Golden",
		"paid":     true,
		"quantity": 620,
		"received": "2021-01-26T00:00:00",
		"type":     "n/a",
		"img":      "https://picsum.photos/id/22/4434/3729",
	},
	{
		"cost":     1309.43,
		"datePaid": nil,
		"id":       12,
		"label":    "Apple - Delicious  Red",
		"paid":     false,
		"quantity": 250,
		"received": "2021-04-14T00:00:00",
		"type":     "n/a",
		"img":      "https://picsum.photos/id/16/2500/1667",
	},
	{
		"cost":     6237.19,
		"datePaid": nil,
		"id":       13,
		"label":    "Arctic Char - Fillets",
		"paid":     false,
		"quantity": 664,
		"received": "2022-06-29T00:00:00",
		"type":     "n/a",
		"img":      "https://picsum.photos/id/23/3887/4899",
	},
	{
		"cost":     10080,
		"datePaid": nil,
		"id":       14,
		"label":    "Arrowroot",
		"paid":     true,
		"quantity": 59,
		"received": "2019-06-25T00:00:00",
		"type":     "Consumer Services",
		"img":      "https://picsum.photos/id/16/2500/1667",
	},
	{
		"cost":     11741.01,
		"datePaid": nil,
		"id":       15,
		"label":    "Asparagus - Green  Fresh",
		"paid":     true,
		"quantity": 790,
		"received": "2019-05-10T00:00:00",
		"type":     "Consumer Services",
		"img":      "https://picsum.photos/id/22/4434/3729",
	},
	{
		"cost":     4773.47,
		"datePaid": nil,
		"id":       16,
		"label":    "Aspic - Amber",
		"paid":     false,
		"quantity": 127,
		"received": "2022-11-13T00:00:00",
		"type":     "n/a",
	},
	{
		"cost":     3748.04,
		"datePaid": nil,
		"id":       17,
		"label":    "Aspic - Clear",
		"paid":     true,
		"quantity": 394,
		"received": "2019-04-30T00:00:00",
		"type":     "n/a",
	},
	{
		"cost":     7259.32,
		"datePaid": nil,
		"id":       18,
		"label":    "Aspic - Light",
		"paid":     true,
		"quantity": 294,
		"received": "2024-01-01T00:00:00",
		"type":     "n/a",
	},
	{
		"cost":     94.33,
		"datePaid": nil,
		"id":       19,
		"label":    "Bacardi Breezer - Strawberry",
		"paid":     false,
		"quantity": 506,
		"received": "2020-09-30T00:00:00",
		"type":     "Capital Goods",
	},
	{
		"cost":     7131.69,
		"datePaid": nil,
		"id":       20,
		"label":    "Bacardi Raspberry",
		"paid":     false,
		"quantity": 614,
		"received": "2022-06-11T00:00:00",
		"type":     "Energy",
	},
	{
		"cost":     4326,
		"datePaid": nil,
		"id":       21,
		"label":    "Bacon Strip Precooked",
		"paid":     true,
		"quantity": 481,
		"received": "2023-01-19T00:00:00",
		"type":     "n/a",
	},
	{
		"cost":     6481.4,
		"datePaid": nil,
		"id":       22,
		"label":    "Bag - Clear 7 Lb",
		"paid":     true,
		"quantity": 9,
		"received": "2024-03-26T00:00:00",
		"type":     "Health Care",
	},
	{
		"cost":     10589.1,
		"datePaid": nil,
		"id":       23,
		"label":    "Bagel - 12 Grain Preslice",
		"paid":     true,
		"quantity": 519,
		"received": "2023-08-12T00:00:00",
		"type":     "n/a",
	},
	{
		"cost":     11189.9,
		"datePaid": nil,
		"id":       24,
		"label":    "Bagel - Ched Chs Presliced",
		"paid":     true,
		"quantity": 930,
		"received": "2021-02-18T00:00:00",
		"type":     "Consumer Services",
	},
	{
		"cost":     8884.63,
		"datePaid": nil,
		"id":       25,
		"label":    "Bagel - Plain",
		"paid":     true,
		"quantity": 890,
		"received": "2023-01-29T00:00:00",
		"type":     "Consumer Services",
	},
	{
		"cost":     337.07,
		"datePaid": nil,
		"id":       26,
		"label":    "Bagelers",
		"paid":     true,
		"quantity": 444,
		"received": "2020-02-07T00:00:00",
		"type":     "Consumer Services",
	},
	{
		"cost":     6466.6,
		"datePaid": nil,
		"id":       27,
		"label":    "Bagelers - Cinn / Brown Sugar",
		"paid":     false,
		"quantity": 629,
		"received": "2022-01-30T00:00:00",
		"type":     "Energy",
	},
	{
		"cost":     10294.11,
		"datePaid": nil,
		"id":       28,
		"label":    "Baking Powder",
		"paid":     false,
		"quantity": 654,
		"received": "2022-01-05T00:00:00",
		"type":     "Finance",
	},
	{
		"cost":     3596.39,
		"datePaid": nil,
		"id":       29,
		"label":    "Bar - Granola Trail Mix Fruit Nut",
		"paid":     false,
		"quantity": 402,
		"received": "2019-01-29T00:00:00",
		"type":     "Technology",
	},
	{
		"cost":     9587.83,
		"datePaid": nil,
		"id":       30,
		"label":    "Bar - Sweet And Salty Chocolate",
		"paid":     true,
		"quantity": 987,
		"received": "2022-12-31T00:00:00",
		"type":     "Technology",
	},
	{
		"cost":     2362.45,
		"datePaid": nil,
		"id":       31,
		"label":    "Bar Mix - Lime",
		"paid":     false,
		"quantity": 585,
		"received": "2024-09-20T00:00:00",
		"type":     "Health Care",
	},
	{
		"cost":     4191.35,
		"datePaid": nil,
		"id":       32,
		"label":    "Bar Nature Valley",
		"paid":     false,
		"quantity": 885,
		"received": "2021-09-07T00:00:00",
		"type":     "Finance",
	},
	{
		"cost":     7580.57,
		"datePaid": nil,
		"id":       33,
		"label":    "Basil - Fresh",
		"paid":     true,
		"quantity": 391,
		"received": "2024-04-11T00:00:00",
		"type":     "Health Care",
	},
	{
		"cost":     2792.49,
		"datePaid": nil,
		"id":       34,
		"label":    "Basil - Seedlings Cookstown",
		"paid":     true,
		"quantity": 368,
		"received": "2022-01-25T00:00:00",
		"type":     "Capital Goods",
	},
	{
		"cost":     1416.14,
		"datePaid": nil,
		"id":       35,
		"label":    "Bay Leaf Ground",
		"paid":     true,
		"quantity": 451,
		"received": "2021-10-02T00:00:00",
		"type":     "Consumer Non-Durables",
	},
	{
		"cost":     11000.8,
		"datePaid": nil,
		"id":       36,
		"label":    "Beans - Black Bean  Preserved",
		"paid":     true,
		"quantity": 192,
		"received": "2019-01-01T00:00:00",
		"type":     "Consumer Services",
	},
	{
		"cost":     10522.08,
		"datePaid": nil,
		"id":       37,
		"label":    "Beans - Butter Lrg Lima",
		"paid":     true,
		"quantity": 399,
		"received": "2023-02-16T00:00:00",
		"type":     "Consumer Services",
	},
}
