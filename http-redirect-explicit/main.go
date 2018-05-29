package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/explicit", explicitRedirect)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Go to /explicit to fire explicit redirect")
}

// We can initiate explicit redirect by setting up `Location` header and
// using http.StatusSeeOther status code (303). Notice fact, that 303 SeeOther
// will always use method GET to initiate redirect.
func explicitRedirect(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Location", "https://github.com/tchorzewski1991")
	w.WriteHeader(http.StatusSeeOther)
}