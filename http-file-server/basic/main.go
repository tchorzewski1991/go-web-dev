package main

import (
	"net/http"
	"io"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/image", image)
	http.ListenAndServe(":8080", nil)
}

func image(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, `<img src="image.jpg">`)
}