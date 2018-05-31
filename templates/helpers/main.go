package main

import (
	"text/template"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
	Job   string
	Age   int
}

func (p Person) FullName() string  {
	return p.fname + ` ` + p.lname
}

// We are allowed to attach any helper function to template.
// Note that functions must be added in slightly different manner
// to make it work as expected. Additional functions must be
// added through FuncMap aggregate and right before template
// parsing.
var helperFunc = template.FuncMap{
	"capitalize": capitalize,
	"downcase"  : downcase,
}

func capitalize(s string) string {
	return strings.Title(s)
}

func downcase(s string) string  {
	return strings.ToLower(s)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(helperFunc).ParseGlob(`templates/*`))
}

func main()  {
	person := Person{"joe", "doe", "PROGRAMMER", 42}
	tpl.ExecuteTemplate(os.Stdout, `index.gohtml`, person)
}
