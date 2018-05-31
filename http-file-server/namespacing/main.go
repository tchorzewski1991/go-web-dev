package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)

	// When using http.FileServer() - http.Dir(".") combination, basically
	// we publicly expose every file from root location. Sometimes it is not
	// expected behavior and very common practise is to use namespaces for
	// separating or grouping files in some logical sense. This is what we
	// can achieve with http.StripPrefix(). This function allows for URI
	// manipulations in quite common and digestible way.

	http.Handle("/resources/", func() http.Handler {
		return http.StripPrefix("/resources", http.FileServer(http.Dir("./assets")))
	}())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<img src="/resources/image.jpg">`)
}
