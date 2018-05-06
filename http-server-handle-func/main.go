package main

import (
	"text/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func about(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func home(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "home.gohtml", nil)
}

func main() {

	// We don't need to define our own Handler type. If it's enough
	// to use DefaultServeMux (which probably is) only thing we
	// need to do is to use built-in http HandleFunc(). It basically
	// takes care on registering and handling new 'routes' for our
	// ServeMux.
	http.HandleFunc("/about", about)
	http.HandleFunc("/home", home)

	// When we provide nil as a handler for ListenAndServe() function
	// Golang will use something called DefaultServeMux. ServeMux is an HTTP
	// request multiplexer. It matches the URL of each incoming request
	// against a list of registered patterns and calls the handler for
	// the pattern that most closely matches the URL.
	http.ListenAndServe(":8080", nil)
}
