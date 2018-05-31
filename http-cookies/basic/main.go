package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// Notice permanent redirection with usage of headers
// location value. Nice fact about HTTP 301 MovedPermanently
// status is fact, that it will be cached by browser to prevent
// from additional requests to the application server.
func root(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Location", "/read")
	w.WriteHeader(http.StatusMovedPermanently)
}

// Using explicit return statement on error handling is
// essential in example below. We want to terminate
// execution process on that point.
func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Cookie found: ", c.Value)
}

func set(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie-name",
		Value: "cookie-value",
	})

	fmt.Fprintln(w, "Cookie written")
	fmt.Fprintln(w, "Check it at: dev tools / application / cookies")
}

