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

func TestNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Boleh Akses")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintf(w, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	buytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Gak Boleh Akses", string(buytes))

}
