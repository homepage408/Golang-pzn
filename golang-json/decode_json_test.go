package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{"FirstName":"Teguh","LastName":"Setiawan","Age":20,"Married":false,"Alamat":{"Desa":"Kalapasawit","Dusun":"Kalapagada","Kecamatan":"Lakbok","Kabupaten":"Ciamis"}}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println("Customer FirstName: " + customer.FirstName)
}
