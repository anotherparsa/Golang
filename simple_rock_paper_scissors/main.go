package main

import (
	"SimpleRockPaperScissors/rps"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/play", play_round)
	http.HandleFunc("/", home_page_handler)
	fmt.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func play_round(w http.ResponseWriter, r *http.Request) {
	round := rps.Play_round(w, r)
	out, err := json.MarshalIndent(round, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(out)
}

func home_page_handler(w http.ResponseWriter, r *http.Request) {
	render_template(w, "index.html")
}

func render_template(w http.ResponseWriter, page string) {
	template, err := template.ParseFiles(page)

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
