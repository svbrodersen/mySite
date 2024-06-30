package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/svbrodersen/mySite/utils"
	"github.com/svbrodersen/mySite/views/templates"
)

type User struct {
	username string
	password string
}

var admin_username string = os.Getenv("ADMIN_USERNAME")
var admin_password string = os.Getenv("ADMIN_PASSWORD")

func LoginPost(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		templates.Login("fail").Render(r.Context(), w)
		return
	}
	if (user.username != admin_username) || (user.password != admin_password) {
		w.WriteHeader(http.StatusUnauthorized)
		templates.Login("fail").Render(r.Context(), w)
		return
	}

	tokenString, expiration_time, err := utils.AuthCreateToken(user.username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		templates.Login("fail").Render(r.Context(), w)
		return
	}
	http.SetCookie(
		w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expiration_time,
		},
	)
	templates.Login("success").Render(r.Context(), w)
}

func LoginGet(w http.ResponseWriter, r *http.Request) {
	templates.Login("default").Render(r.Context(), w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		LoginPost(w, r)
	} else if r.Method == "GET" {
		LoginGet(w, r)
	}
}
