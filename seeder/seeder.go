package seeder

import (
	"context"
	"ecommerce-project/database"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Price  float64            `json:"price" bson:"price"`
	Rating float64            `json:"rating" bson:"rating"`
	Image  string             `json:"image" bson:"image"`
}

func SeedData() error {
	client, err := database.MongoConnect()

	if err != nil {
		return err
	}

	products := []Product{
		{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Rating: 4.5, Image: "https://images.pexels.com/photos/18105/pexels-photo.jpg?auto=compress&cs=tinysrgb&w=600"},
		{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Rating: 4.8, Image: "https://images.pexels.com/photos/788946/pexels-photo-788946.jpeg?auto=compress&cs=tinysrgb&w=600"},
		{ID: primitive.NewObjectID(), Name: "Headphones", Price: 99.99, Rating: 4.2, Image: "https://images.pexels.com/photos/1649771/pexels-photo-1649771.jpeg?auto=compress&cs=tinysrgb&w=600"},
	}

	var productsInterface []interface{}

	collection := client.Database(database.Database).Collection(string(database.ProductsCollection))

	for _, p := range products {
		productsInterface = append(productsInterface, p)
	}

	collection.DeleteMany(context.Background(), productsInterface)
	_, err = collection.InsertMany(context.Background(), productsInterface)

	if err != nil {
		return err
	}
	fmt.Println("Successfully seeded data")
	return nil

}
