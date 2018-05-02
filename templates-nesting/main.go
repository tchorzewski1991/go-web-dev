package main

import (
	"text/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(`templates/*`))
}

type Person struct {
	fname string
	lname string
	Admin bool
}

func (p Person) FullName() string {
	return p.fname + ` ` + p.lname
}

func main() {
	p1 := Person{"Joe", "Doe", true}
	p2 := Person{"Boe", "Moe", false}

	people := []Person{p1, p2}

	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", people)
}
