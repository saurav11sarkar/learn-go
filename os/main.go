package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TProduct struct {
	ID          int
	Title       string
	Description string
	Price       float64
	Image       string
}

var productList = []TProduct{
	{ID: 1, Title: "Product 1", Description: "Description 1", Price: 10.99, Image: "image1.jpg"},
	{ID: 2, Title: "Product 2", Description: "Description 2", Price: 20.99, Image: "image2.jpg"},
	{ID: 3, Title: "Product 3", Description: "Description 3", Price: 30.99, Image: "image3.jpg"},
	{ID: 4, Title: "Product 4", Description: "Description 4", Price: 40.99, Image: "image4.jpg"},
	{ID: 5, Title: "Product 5", Description: "Description 5", Price: 50.99, Image: "image5.jpg"},
}

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(productList)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/product", handleGetProduct)

	fmt.Println("Server is running on port http://localhost:5000")

	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println(err)
		return
	}

}
