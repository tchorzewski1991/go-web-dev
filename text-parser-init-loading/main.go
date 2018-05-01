package main

import (
	"text/template"
	"os"
	"log"
)

// We declare tpl on package scope to make it accessible easily form
// every location within the package. tpl variable is a container for
// all parsed templates.
var tpl *template.Template

// init() function runs only, when our program is starting.
// Loading templates in that way makes your program much more performant,
// as init() function will be executed only once. Notice usage of Must()
// helper function to throw panic() in case of an error due to global
// templates parsing.
func init() {
	tpl = template.Must(template.ParseGlob(`templates/*.gohtml`))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, `template-1.gohtml`, nil)
	if err != nil { log.Fatalln(err) }

	err = tpl.ExecuteTemplate(os.Stdout, `template-2.gohtml`, nil)
	if err != nil { log.Fatalln(err) }
}
