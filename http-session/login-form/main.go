package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
	"errors"
)

var tpl *template.Template

type User struct {
	Email    string
	Password string
}

var mockSID, _   = uuid.NewV4()
var mockEmail    = "test@email.com"
var mockPassword = "password"

type dbSessions map[string]string
var  sessions = dbSessions{}
func (dbS dbSessions) NotFound() error {
	return errors.New("Session not found")
}

type dbUsers map[string]User
var  users = dbUsers{}
func (dbU dbUsers) NotFound() error {
	return errors.New("User not found")
}

func mockUser() {
	users[mockEmail] = User{ Email: mockEmail, Password: mockPassword}
}

func mockSession() {
	sessions[mockSID.String()] = mockEmail
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	mockSession()
	mockUser()

	fmt.Print("Mock sessionID: ", mockSID)
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	redirectTo(w, req, "/dashboard")
}

func dashboard(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		redirectTo(w, req, "/login")
		return
	}

	u, err := getUser(req)
	if err != nil {
		redirectTo(w, req, "/login")
		return
	}

	tpl.ExecuteTemplate(w, "dashboard.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		redirectTo(w, req, "/dashboard")
		return
	}

	if req.Method == http.MethodPost {
		eml := req.FormValue("email")

		sID, _ := uuid.NewV4()

		c := &http.Cookie{
			Name:     "session-ID",
			Value:    sID.String(),
			HttpOnly: true,
		}

		http.SetCookie(w, c)

		sessions[sID.String()] = eml

		redirectTo(w, req, "/dashboard")
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session-ID")
	if err == http.ErrNoCookie {
		redirectTo(w, req, "/login")
		return
	}

	c.MaxAge = -1

	http.SetCookie(w, c)
	redirectTo(w, req, "/dashboard")
}
