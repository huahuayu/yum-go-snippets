package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

// For this code to run, you will need this package:
// go get github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewRandom() // version 4 uuid
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
