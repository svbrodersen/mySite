package handlers

import (
	"github.com/svbrodersen/mySite/views/templates"
	"net/http"
	"time"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Contact(true).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Contact(false).Render(r.Context(), w)
	} else if r.Method == "POST" {
		time.Sleep(2 * time.Second)
	}
}
