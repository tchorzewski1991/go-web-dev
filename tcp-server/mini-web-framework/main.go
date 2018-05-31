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

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	fields := strings.Fields(ln)
	method := fields[0]
	url    := fields[1]

	if method == `GET` && url == `/` {
		index(conn)
	}
	if method == `GET` && url == `/about` {
		about(conn)
	}
}

func index(conn net.Conn) {
	tpl = tpl.Lookup("index.gohtml")

	var tplBuffer bytes.Buffer

	if err := tpl.Execute(&tplBuffer, struct {
		Home  bool
		About bool
	}{true, false }); err != nil {
		log.Fatalln(err)
	}

	body := tplBuffer.String()

	fmt.Fprint (conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint (conn, "Content-Type: text/html\r\n")
	fmt.Fprint (conn, "\r\n")
	fmt.Fprint (conn, body)

}

func about(conn net.Conn) {
	tpl = tpl.Lookup("index.gohtml")

	var tplBuffer bytes.Buffer

	if err := tpl.Execute(&tplBuffer, struct {
		Home  bool
		About bool
	}{false, true}); err != nil {
		log.Fatalln(err)
	}

	body := tplBuffer.String()

	fmt.Fprint (conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint (conn, "Content-Type: text/html\r\n")
	fmt.Fprint (conn, "\r\n")
	fmt.Fprint (conn, body)
}
