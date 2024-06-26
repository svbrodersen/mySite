package handlers

import (
	"github.com/svbrodersen/mySite/views/templates"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.Hello().Render(r.Context(), w)
}
