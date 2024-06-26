package main

import (
	"fmt"
	"github.com/svbrodersen/mySite/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	fmt.Println("Listening on :3333")
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		print("Error happend %s", err)
	}
}
