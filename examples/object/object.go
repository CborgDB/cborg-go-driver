package main

import (
	"os"

	"github.com/cborgdb/cborg-go-driver/cborg"
)

type User struct {
	pseudo    string
	firstName string
	lastName  string
	age       uint
}

type Car struct {
	marque string
	model  string
	year   uint
	km     uint
}

func main() {
	client := cborg.NewClient("127.0.0.1", "30000")
	err := client.Connect()
	if err != nil {
		os.Exit(1)
	}
	defer client.Disconnect()

	// Create "cborgdb" database
	resultDatabase, _ := client.CreateDatabase("cborgdb")

	// Create "users" and "cars" collections
	resultUsers, _ := resultDatabase.Database.CreateCollection("users")
	resultCars, _ := resultDatabase.Database.CreateCollection("cars")

	// Get collections from results
	users := resultUsers.Collection
	cars := resultCars.Collection

	// Insert some user items
	users.InsertOne(User{"adilon", "Adil", "Benhlal", 30})
	users.InsertOne(User{"bob", "Bob", "Obo", 28})
	users.InsertOne(User{"alice", "Alice", "Ecila", 2})
	users.FindOne(map[string]interface{}{"firstName": "Adil", "age": 30})

	// Insert some car items
	cars.InsertOne(Car{"Mercedes", "Classe c", 2018, 300_000})
	cars.InsertOne(Car{"Mercedes", "GLE", 2021, 50_000})
	cars.FindOne(map[string]interface{}{"model": "Classe c", "year": 2018})
}
