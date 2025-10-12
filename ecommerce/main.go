package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product

func rootRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is running")
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	handle(w, r)
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(products)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var newProduct Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		newProduct.ID = len(products) + 1
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(newProduct)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	app := http.NewServeMux()

	app.Handle("GET /", http.HandlerFunc(rootRoutes))
	app.Handle("GET /products", http.HandlerFunc(handleProducts))
	app.Handle("POST /products", http.HandlerFunc(handleProducts))

	fmt.Println("server is running on port http://localhost:8080")

	err := http.ListenAndServe(":8080", app)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		return
	}

}

func init() {
	products = []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 20.99},
		{ID: 3, Name: "Product 3", Price: 30.99},
		{ID: 4, Name: "Product 4", Price: 40.99},
		{ID: 5, Name: "Product 5", Price: 50.99},
		{ID: 6, Name: "Product 6", Price: 60.99},
		{ID: 7, Name: "Product 7", Price: 70.99},
		{ID: 8, Name: "Product 8", Price: 80.99},
		{ID: 9, Name: "Product 9", Price: 90.99},
		{ID: 10, Name: "Product 10", Price: 100.99},
	}
}
