package main

import (
	"html/template"
	"strings"
	"net"
	"log"
	"bufio"
	"fmt"
	"bytes"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.New("").Funcs(tplHelpers).ParseGlob(`templates/*`))
}

var tplHelpers = template.FuncMap{
	"capitalize": func (s string) string { return strings.Title(s) },
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil { log.Fatalln(err) }
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil { log.Fatalln(err) }
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}
