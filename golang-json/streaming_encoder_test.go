package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Teguh",
		LastName:  "Setiawan",
		Age:       20,
		Married:   true,
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}
