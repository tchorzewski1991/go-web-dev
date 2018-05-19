package main

import (
	"github.com/tchorzewski1991/go-web-dev/exercises/http-link-parser"
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {

	// Our link.Parse() function expect to take as an argument
	// a new reader. New readers can be created easily with
	// strings.NewReader() constructor.

	f, err := ioutil.ReadFile("example.html")

	r := strings.NewReader(string(f))

	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
