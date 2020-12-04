package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	// Using defaultServeMux
	log.Fatal(http.ListenAndServe(":5500", nil))
}

// HomeHandler func handles the root
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hi this is home route</h>")
}
