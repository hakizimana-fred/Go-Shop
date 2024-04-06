package main

import (
	"ecommerce-project/seeder"
	"fmt"
)

func main() {
	err := seeder.SeedData()
	if err != nil {
		fmt.Println("Error seeding data:", err)
	}

}
