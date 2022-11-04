package main

import (
	"os"

	"github.com/cborgdb/cborg-go-driver/cborg"
)

func main() {
	client := cborg.NewClient("127.0.0.1", "30000")
	err := client.Connect()
	if err != nil {
		os.Exit(1)
	}
	defer client.Disconnect()

	// Create Database
	result, _ := client.CreateDatabase("cborgdb")
	// Create Collection using previous result
	result.Database.CreateCollection("numbers")

	// Init "number" Collection
	numbers := client.Database("cborgdb").Collection("numbers")

	// Insert items
	for i := 0; i < 1000; i++ {
		numbers.InsertOne(i)
	}
}
