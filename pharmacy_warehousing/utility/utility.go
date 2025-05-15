package utility

import (
	"log"
	"net/http"
	"text/template"
)

func Render_template(w http.ResponseWriter, path string) {
	template, err := template.ParseFiles(path)

	if err != nil {
		log.Print(err)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = template.Execute(w, nil)

	if err != nil {
		log.Print(err)
		return
	}
}

