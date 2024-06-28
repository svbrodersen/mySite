package handlers

import (
	"github.com/svbrodersen/mySite/views/templates"
	"net/http"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Projects(true).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Projects(false).Render(r.Context(), w)
	}
}
