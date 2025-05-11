package utility

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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

func Is_user_logged(r *http.Request) bool {
	_, err := r.Cookie("sessionid")

	if err != nil {
		fmt.Printf("Failed to get the cookie : %v\n", err)
		return false
	}

	return true

}
