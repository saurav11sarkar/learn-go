package cmd

import (
	"fmt"
	"net/http"

	"github.com/saurav11sarkar/student-api/handlers"
	"github.com/saurav11sarkar/student-api/routes"
)

func Server() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.PostProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	globalRoute := routes.GlobalMiddleware(mux)
	fmt.Println("🚀 Server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", globalRoute); err != nil {
		fmt.Println("Server error:", err)
		return
	}
}
