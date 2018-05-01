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
	Age   int
}

func (p Person) FullName() string {
	return p.fname + " " + p.lname
}

func main() {
	p1 := Person{"Joe", "Doe", 20}
	p2 := Person{"Boe", "Moe", 21}

	people := []Person{p1, p2}

	tpl.ExecuteTemplate(os.Stdout, `index.gohtml`, people)
}
