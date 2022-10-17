package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "00001",
		Name:     "Mackbook Pro M1",
		ImageUrl: "http://example.com/iamge.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}

func TestJsonTagDecode(t *testing.T) {
	// golang decode no matter uppercase or lowercase
	jsonString := `{"id":"00001","name":"Mackbook Pro M1","image_url":"http://example.com/iamge.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
