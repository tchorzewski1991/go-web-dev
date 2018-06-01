package main

import (
	"net/http"
	"fmt"
	"strconv"
)

var visits int = 0

func main() {
	http.HandleFunc("/", root)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	visits += 1

	c, err := req.Cookie("visits")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "visits",
			Value: strconv.Itoa(visits),
		})
	}

	c.Value = strconv.Itoa(visits)

	http.SetCookie(w, c)

	fmt.Fprintln(w, "Total visits: ", visits)
}
