package router

import (
	"github.com/svbrodersen/mySite/database"
	"github.com/svbrodersen/mySite/handlers"
	"net/http"
)

func Route(mux *http.ServeMux) {

	accessor, err := database.NewSqliteAccessor()
	if err != nil {
		panic("Unable to open database")
	}
	proj := handlers.ProjectsDatabase{Db: accessor}
	mux.HandleFunc("/", handlers.AboutHandler)
	mux.HandleFunc("/projects", proj.ProjectsHandler)
	mux.HandleFunc("/contact", handlers.ContactHandler)
	mux.HandleFunc("/cv", handlers.CvHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/projects/*", proj.ProjectDetailHandler)
	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
