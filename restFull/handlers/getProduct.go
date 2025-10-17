package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/saurav11sarkar/student-api/database"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(database.ProductDataList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
