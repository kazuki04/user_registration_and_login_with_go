package controllers

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/user_registration_and_login_with_go/app/models"
	"golang.org/x/crypto/bcrypt"
)

func userSignUp(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var user models.User
	// json.NewDecoder(r.Body).Decode(&user)
	user := &models.User{Name: r.FormValue("name"), Email: r.FormValue("email"), Password: r.FormValue("password")}
	fmt.Println(r.FormValue("password"))
	fmt.Println(user.Password)
	user.Password = getHash([]byte(user.Password))
	fmt.Println(user.Password)
	user.Save()
	// if user.Save().Error != nil {
	// 	return
	// }
	templates.ExecuteTemplate(w, "registered.html", user)
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
