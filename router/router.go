package router

import (
	"github.com/svbrodersen/mySite/handlers"
	"net/http"
)

func Route(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.AboutHandler)
	mux.HandleFunc("/projects", handlers.ProjectsHandler)
	mux.HandleFunc("/contact", handlers.ContactHandler)
	mux.HandleFunc("/cv", handlers.CvHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
