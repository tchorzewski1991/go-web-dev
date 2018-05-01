package main

import (
	"text/template"
	"log"
	"os"
)

func main() {

	// Template loading is the two way process. First step is always
	// to parse specific file and the second is to execute/inject loaded
	// template into something that implements io.Writer interface. In
	// case below this is a simple os.Stdout.
	tpl, err := template.ParseFiles("template-1.gohtml")
	if err != nil { log.Fatalln(err) }

	err = tpl.Execute(os.Stdout, nil)
	if err != nil { log.Fatalln(err) }

	// We can specify any Writer we want. Great example will be to
	// use os package and create whole new 'index.html' file and
	// use Execute() function to load/inject parsed template into that file.
	file, err := os.Create("index.html")
	if err != nil { log.Fatalln(err) }

	err = tpl.Execute(file, nil)
	if err != nil { log.Fatalln(err) }

	// We can also load more than one template at once. This is
	// very common way to working with templates in web development
	// with Go. Specifying which template to execute could be achieved
	// with ExecuteTemplate() function.
	tpl, err = tpl.ParseFiles("template-2.gohtml", "template-3.gohtml")
	if err != nil { log.Fatalln(err) }

	err = tpl.ExecuteTemplate(os.Stdout, "template-2.gohtml", nil)
	if err != nil { log.Fatalln(err) }

	// Choosing which template to parse is ok, but it is not the most
	// appreciated way to resolve issue of parsing all templates. This
	// list could be really, really long and typing names of all files
	// could be redundant. We always should use ParseGlob() instead.

	tpl, err = template.ParseGlob("templates/*.gohtml")
	if err != nil { log.Fatalln(err) }

	err = tpl.ExecuteTemplate(os.Stdout, "template-5.gohtml", nil)
	if err != nil { log.Fatalln(err) }
}
