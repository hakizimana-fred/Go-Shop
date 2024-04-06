package main

import (
	"ecommerce-project/controllers"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/products", controllers.ProductsController)

	http.ListenAndServe(":4000", router)
}
