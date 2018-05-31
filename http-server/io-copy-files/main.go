package main

import (
	"net/http"
	"io"
	"os"
	"log"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/image.jpg", handleImage)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/image.jpg">`)
}

func handleImage(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open(`image.jpg`)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		log.Fatalln(err)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}
