package main

import (
	"text/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"strings"
)

var tpl *template.Template

var tplFuncMap = template.FuncMap{ "capitalize": capitalize }

func capitalize(s string) string { return strings.Title(s) }

func init() {
	tpl = template.Must(template.New("").Funcs(tplFuncMap).ParseGlob("templates/*"))
}


func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/users/:user", about)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params)  {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "user.gohtml", p.ByName("user"))
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
