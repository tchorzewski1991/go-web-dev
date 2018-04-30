package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main() {
	injected := "This string will be injected into HTML"

	template := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Hello World!</title>
			</head>
			<body>
				<h1>` + injected + `</h1>
			</body>
		</html>
	`)

	file, err := os.Create("index.html")

	if err != nil {
		log.Fatal("There was an error while trying to create a file...", err)
	}

	defer file.Close()

	io.Copy(file, strings.NewReader(template))
}
