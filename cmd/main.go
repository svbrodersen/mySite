package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/svbrodersen/mySite/handlers"
	"github.com/svbrodersen/mySite/views/templates"
)

func main() {
	component := templates.Hello("Simon")
	http.Handle("/", handlers.indexHandler)
	fmt.Println("Listening on :3333")
	http.ListenAndServe(":3333", nil)
}
