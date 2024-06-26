package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", roothandler)
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
}
