package handlers

import (
	"github.com/svbrodersen/mySite/views/templates"
	"net/http"
)

func CvHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Cv(true).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Cv(false).Render(r.Context(), w)
	}
}
