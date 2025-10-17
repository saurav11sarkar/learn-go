package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/saurav11sarkar/student-api/database"
)

func PostProduct(w http.ResponseWriter, r *http.Request) {
	var product database.ProductData
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product.Id = len(database.ProductDataList) + 1
	database.ProductDataList = append(database.ProductDataList, product)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
