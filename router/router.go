package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Product is a struct that defines data sent over HTTP methods
type Product struct {
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

// AllProducts is a slice of all data sent over HTTP methods
// -> Mockup database
type AllProducts []Product

// ProductHandler is our custom router struct that handles different HTTP methods
type ProductHandler struct {
	sync.Mutex
	products AllProducts
}

// NewProductHandler is constructor
func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		products: AllProducts{
			Product{"Bananas", 15.00},
			Product{"Nutela", 35.00},
			Product{"Kiwi", 20.00},
			Product{"Milk", 10.00},
		},
	}
}

// Making our own HANDLER func with all of the http methods needed
func (ph *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ph.Get(w, r)
	case "POST":
		ph.Post(w, r)
	case "PUT", "PATCH":
		ph.Put(w, r)
	case "DELETE":
		ph.Delete(w, r)
	default:
		respondERROR(w, http.StatusMethodNotAllowed, "...Ooops, You called an invalid HTTP method... (Allowed only GET, POST, PUT and DELETE")
	}
}

// Making generic response with JSON data
func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Unable to send data ...", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// Making generic error response
func respondERROR(w http.ResponseWriter, statusCode int, errMsg string) {
	respondJSON(w, statusCode, map[string]string{"error": errMsg})
}

// Get function fetches data from DB (mockup struct)
func (ph *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from GET method")
}

// Post function sends/writes to data in DB
func (ph *ProductHandler) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from POST method")
}

// Put function updates data in DB
func (ph *ProductHandler) Put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from PUT method")
}

// Delete function deletes data from DB
func (ph *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from DELETE method")
}
