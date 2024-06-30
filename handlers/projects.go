package handlers

import (
	"net/http"

	"github.com/svbrodersen/mySite/utils"
	"github.com/svbrodersen/mySite/views/templates"
)

func loggedInProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Projects(true, true).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Projects(false, true).Render(r.Context(), w)
	}
}

func defaultProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Projects(true, false).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Projects(false, false).Render(r.Context(), w)
	}
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		defaultProjects(w, r)
		return
	}
	tknStr := cookie.Value
	err = utils.AuthVerifyToken(tknStr)
	// Verification failed
	if err != nil {
		defaultProjects(w, r)
		return
	}
	// Only admin logged in
	loggedInProjects(w, r)
}
