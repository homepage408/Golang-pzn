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

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Ketemu")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/halllo", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	buytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Gak Ketemu", string(buytes))

}
