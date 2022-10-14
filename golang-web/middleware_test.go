package main

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHttp(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("before")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After")
}

// func TestMiddleware(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Println("Handler Executed")
// 		fmt.Fprint(writer, "Hello Middleware")
// 	})

// 	logMiddleware := new(LogMiddleware)
// 	logMiddleware.Handler = mux

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: logMiddleware,
// 	}
// }
