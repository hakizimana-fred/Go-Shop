package controllers

import (
	"context"
	"ecommerce-project/database"
	"ecommerce-project/seeder"
	"ecommerce-project/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProductsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProducts(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProducts(w http.ResponseWriter, _ *http.Request) error {
	client, err := database.MongoConnect()
	var products []*seeder.Product

	if err != nil {
		return err
	}
	collection := client.Database(database.Database).Collection(string(database.ProductsCollection))

	cur, err := collection.Find(context.Background(), bson.D{
		primitive.E{},
	})

	for cur.Next(context.Background()) {
		var p seeder.Product
		err := cur.Decode(&p)

		if err != nil {
			return nil
		}
		products = append(products, &p)
	}

	return utils.WriteJson(w, http.StatusOK, products)
}
