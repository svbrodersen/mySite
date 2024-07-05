package main

import (
	"fmt"
	"github.com/svbrodersen/mySite/router"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router.Route(mux)
	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		print("Error happend %s", err)
	}
}
