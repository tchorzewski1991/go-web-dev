package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/expire", expireCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `<a href="/set">Set Cookie</a>`)
}

func setCookie(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "token",
	})
	fmt.Fprint(w, `<a href="/read">Read Cookie</a>`)
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprint(w, `<a href="/expire">Expire</a>`)
}

// Notice that, to remove cookie you need to setup MaxAge attribute
// accordingly to value specified in docs. -1 will inform browser
// about need to expire selected cookie.
func expireCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)

}