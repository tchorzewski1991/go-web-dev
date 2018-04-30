package main

import "fmt"

func main() {
	str := `This string will be injected into the paragraph!`

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Go web dev!</title>
		</head>
		<body>
			<p>` + str + `</p>
		</body>
	</html>
	`
	fmt.Println(tpl)
}
