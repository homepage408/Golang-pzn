package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}

func TestEncoder(t *testing.T) {
	logJson("Teguh")
	logJson(33)
	logJson(true)
	logJson([]string{"Teguh", "Setiawan"})

}
