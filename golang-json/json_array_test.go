package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Teguh",
		LastName:  "Setiawan",
		Age:       20,
		Married:   false,
		Hobbies:   []string{"Game", "Reading", "Coding"},
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

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Teguh","LastName":"Setiawan","Age":20,"Married":false,"Hobbies":["Game","Reading","Coding"],"Alamat":{"Desa":"Kalapasawit","Dusun":"Kalapagada","Kecamatan":"Lakbok","Kabupaten":"Ciamis"}}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println("Customer FirstName: " + customer.FirstName)
	fmt.Println("Customer Hoobies: ", customer.Hobbies)
}

func TestJsonArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Teguh",
		LastName:  "Setiawan",
		Age:       20,
		Married:   false,
		Hobbies:   []string{"Game", "Reading", "Coding"},
		Alamat: Alamat{
			Desa:      "Kalapasawit",
			Dusun:     "Kalapagada",
			Kecamatan: "Lakbok",
			Kabupaten: "Ciamis",
		},
		Addreses: []Address{
			{
				Street:     "Jalan Jalan",
				Country:    "Indonesia",
				PostalCode: "7474748",
			},
			{
				Street:     "Jalan Jalan Lagi",
				Country:    "Indonesia",
				PostalCode: "776655",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Teguh","LastName":"Setiawan","Age":20,"Married":false,"Hobbies":["Game","Reading","Coding"],"Alamat":{"Desa":"Kalapasawit","Dusun":"Kalapagada","Kecamatan":"Lakbok","Kabupaten":"Ciamis"},"Addreses":[{"Street":"Jalan Jalan","Country":"Indonesia","PostalCode":"7474748"},{"Street":"Jalan Jalan Lagi","Country":"Indonesia","PostalCode":"776655"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println("Customer FirstName: " + customer.FirstName)
	fmt.Println("Customer Hoobies: ", customer.Hobbies)
	fmt.Println("Customer Addreses: ", customer.Addreses)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Jalan","Country":"Indonesia","PostalCode":"7474748"},{"Street":"Jalan Jalan Lagi","Country":"Indonesia","PostalCode":"776655"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	json.Unmarshal(jsonBytes, addresses)

	fmt.Println("Addresses", addresses)
}

func TestOnlyJsonArrayComplexEncode(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Jalan",
			Country:    "Indonesia",
			PostalCode: "7474748",
		},
		{
			Street:     "Jalan Jalan Lagi",
			Country:    "Indonesia",
			PostalCode: "776655",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
