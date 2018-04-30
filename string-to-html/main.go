package main

import "fmt"

func main() {
	injected := `This string will be injected into HTML paragraph`

	template := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Go web dev</title>
		</head>
		<body>
			<p>` + injected + `<p>
		</body>
	</html>
	`

	fmt.Println(template)
}
