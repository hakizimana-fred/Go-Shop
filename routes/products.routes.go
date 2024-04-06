package routes

import (
	"fmt"
	"net/http"
)

func InitRoutes() {
	router := http.NewServeMux()

	router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Getting products")
	})

}
