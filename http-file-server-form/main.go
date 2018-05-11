package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", postForm)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func postForm(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>` + v)
}


