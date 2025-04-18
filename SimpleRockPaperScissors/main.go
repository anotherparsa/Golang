package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	fmt.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("index.html")

	if err != nil {
		log.Println(err)
		return
	}

	err = template.Execute(w, nil)
	
	if err != nil {
		log.Println(err)
		return
	}
}
