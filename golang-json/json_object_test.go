package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	Married   bool      `json:"married"`
	Hobbies   []string  `json:"hobbies"`
	Alamat    Alamat    `json:"alamat"`
	Addreses  []Address `json:"addreses"`
}

type Alamat struct {
	Desa      string `json:"desa"`
	Dusun     string `json:"dusun"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
}

type Address struct {
	Street     string `json:"street"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Teguh",
		LastName:  "Setiawan",
		Age:       20,
		Married:   false,
		Alamat: Alamat{
			Desa:      "Kalapasawit",
			Dusun:     "Kalapagada",
			Kecamatan: "Lakbok",
			Kabupaten: "Ciamis",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
