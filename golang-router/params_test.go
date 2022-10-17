package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/julienschmidt/httprouter"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/product/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	buytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 2", string(buytes))

}
