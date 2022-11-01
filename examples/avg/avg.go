package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cborgdb/cborg-go-driver/pkg/cborg"
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

	// Init "numbers" Collection
	numbers := client.Database("cborgdb").Collection("numbers")

	// First Insert
	firstStart := time.Now()
	numbers.InsertOne(0)
	firstEnd := time.Since(firstStart)

	// min = Max = AVG
	min, max, avg := firstEnd, firstEnd, firstEnd

	// Insert items
	for i := 1; i < 1000; i++ {
		// Insert
		start := time.Now()
		numbers.InsertOne(i)
		end := time.Since(start)

		// Get min, max
		if end < min {
			min = end
		}
		if end > max {
			max = end
		}

		// Bad avg compute :)
		v := (int(avg.Nanoseconds())*(i-1) + int(end.Nanoseconds())) / i
		s := string(strconv.Itoa(v) + string("ns"))
		avg, _ = time.ParseDuration(s)
	}

	// Print min, max, avg
	fmt.Printf("min: %v\n", min)
	fmt.Printf("max: %v\n", max)
	fmt.Printf("avg: %v\n", avg)
}
