package controllers

import (
	"fmt"
	"github.com/user_registration_and_login_with_go/config"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("app/views/users/new.html", "app/views/users/registered.html"))

func makeHandler(fn func(r http.ResponseWriter, w *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new.html", nil)
}

func StartWebServer() error {
	http.HandleFunc("/new/", makeHandler(newHandler))
	http.HandleFunc("/save/", makeHandler(userSignUp))
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), nil)
}
