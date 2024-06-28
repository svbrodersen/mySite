package handlers

import (
	"github.com/svbrodersen/mySite/views/templates"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.AboutMe(true).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.AboutMe(false).Render(r.Context(), w)
	}
}
