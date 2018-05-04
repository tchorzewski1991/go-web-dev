package main

import (
	"net/http"
	"fmt"
)

type customHandler int

func (c customHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w,"Hello from custom handler!")
}

func main() {
	var c customHandler

	// Responsibility of http.ListenAndServe() function is to take
	// care of all incoming requests and handle them with provided
	// Handler type. Handler type is an interface. If we inspect that
	// interface we can notice it's matter of implementing ServeHTTP()
	// method which takes ResponseWriter and a pointer to the Request.
	// Expanding any other type with that method means it becomes a new
	// Handler as well. ListenAndServe() expects Handler and any incoming
	// requests from the web will be handled by that Handler (customHandler
	// in our case)
	http.ListenAndServe(":8080", c)
}
