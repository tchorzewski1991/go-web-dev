package main

import (
	"text/template"
	"os"
)

var tpl *template.Template

type Person struct {
	fname string
	lname string
	admin bool
}

func (p Person) FullName() string {
	return p.fname + ` ` + p.lname
}

func (p Person) IsAdmin() bool {
	return p.admin == true
}

func init() {
	tpl = template.Must(template.ParseGlob(`templates/*`))
}

func main() {
	person := Person{"Joe", "Doe", true}
	tpl.ExecuteTemplate(os.Stdout, `index.gohtml`, person)
}
