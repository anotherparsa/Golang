package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	fmt.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
