package main

import (
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	cwd, _ := os.Getwd()
	//fmt.Println("filepath is:------------",filepath.Join( cwd, "./template/login.gtpl" ))
	//t, err := template.ParseFiles()
	tpl = template.Must(template.ParseGlob(filepath.Join(cwd, "./templates/*")))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewRandom() // version 4 uuid
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV
