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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recived request")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintf(w, "Response Melalui Middleware Dulu")
	})

	middleware := LogMiddleware{Handler: router}

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)
	response := recorder.Result()

	buytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Response Melalui Middleware Dulu", string(buytes))

}
