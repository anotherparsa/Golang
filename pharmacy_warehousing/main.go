package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the product page")
	})
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
