package main

import "net/http"

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session-ID")

	if err != nil {
		return false
	}

	eml   := sessions[c.Value]
	_, ok := users[eml]
	return ok
}

func getUser(req *http.Request) (User, error) {
	var user User

	c, err := req.Cookie("session-ID")
	if err != nil {
		return user, err
	}

	eml, ok := sessions[c.Value]
	if !ok {
		return user, sessions.NotFound()
	}

	user, ok = users[eml]
	if !ok {
		return user, users.NotFound()
	}

	return user, nil
}

func redirectTo(w http.ResponseWriter, req *http.Request, loc string) {
	http.Redirect(w, req, loc, http.StatusSeeOther)
}
