package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

func main() {
	http.HandleFunc("/", generateUUID)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// Notice use of HttpOnly attribute. What does it mean in terms
// of security is that, you can't access this cookie from
// Javascript layer. We can access it only through HTTP Protocol.
func generateUUID(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")

	if err != nil {
		id, _  := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value:  id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	fmt.Println(cookie)
}
