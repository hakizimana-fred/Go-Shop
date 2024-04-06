package controllers

import (
	"fmt"
	"net/http"
)

func ProductsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProducts(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all products from db")
}
