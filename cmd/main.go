package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/svbrodersen/mySite/views/templates"
)

func main() {
	component := templates.Hello("Simon")
	http.Handle("/", templ.Handler(component))
	fmt.Println("Listening on :3333")
	http.ListenAndServe(":3333", nil)
}
