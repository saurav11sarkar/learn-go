package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saurav11sarkar/student-api/database"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Search for product
	for _, product := range database.ProductDataList {
		if product.Id == id {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(product); err != nil {
				http.Error(w, "Error encoding product", http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
