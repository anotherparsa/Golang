package utility

import (
	"net/http"
	"text/template"
)

func Render_template(w http.ResponseWriter, path string, data interface{}) error {
	template, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	err = template.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
