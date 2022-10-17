package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJsonDecode(t *testing.T) {
	jsonString := `{"id":"00001","name":"Mackbook Pro M1","price":20000000}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "00001",
		"name":  "Mackbook Pro M1 Encode",
		"price": 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
