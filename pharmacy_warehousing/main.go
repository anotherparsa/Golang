package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the product page")
	})

	http.HandleFunc("/login", loginPageHandler)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	template_render(w, "templates/login.html")
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func template_render(w http.ResponseWriter, path string) {
	template, err := template.ParseFiles(path)

	if err != nil {
		log.Print(err)
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		log.Print(err)
		return
	}
}
