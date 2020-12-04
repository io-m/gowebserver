package main

import (
	"fmt"
	"log"
	"net/http"

	r "github.com/io-m/gowebserver/router"
)

func main() {
	ph := r.NewProductHandler()
	http.Handle("/products", ph)
	http.Handle("/products/", ph)
	// Using defaultServeMux
	log.Fatal(http.ListenAndServe(":5500", nil))
}

// HomeHandler func handles the root
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hi this is home route</h>")
}
