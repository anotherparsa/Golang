package utility

import (
	"html/template"
	"log"
	"net/http"
)

func TemplateRendering(w http.ResponseWriter, path string) {
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
