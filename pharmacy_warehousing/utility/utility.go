package utility

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var line_number = 0

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

func Error_handler(w http.ResponseWriter, errortext string, sourceoferror string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		defer file.Close()
		_, _ = file.WriteString(fmt.Sprintf("%v - %v => %v\n", line_number, errortext, sourceoferror))
	}
	line_number++
	Render_template(w, "./utility/templates/errorpage.html", nil)
}
