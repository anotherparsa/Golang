package utility

import (
	"http/template"
	"log"
	"net/http"
)

func TemplateRendering(w http.ResponseWriter, path string) {
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
