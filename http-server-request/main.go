package main

import (
	"net/http"
	"html/template"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type customHandler int

func (c customHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil { fmt.Fprint(w, err) }

	data := struct {
		Params map[string][]string
		Header http.Header
	}{
		req.Form,
		req.Header,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main()  {
	var c customHandler
	http.ListenAndServe(":8080", c)
}
