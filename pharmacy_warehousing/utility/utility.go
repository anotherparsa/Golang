package utility

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/google/uuid"
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

func Error_handler(w http.ResponseWriter, errortext string) {
	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	_, _ = file.WriteString(fmt.Sprintf("%v - %v\n", line_number, errortext))
	line_number++
	Render_template(w, "./utility/templates/errorpage.html", nil)
}

func Generate_staffid_userid(position string) (string, string) {
	var random_staffid string
	random_staffid_postfix := strconv.Itoa(rand.IntN(89999) + 10000)
	if position == "recipient" {
		random_staffid = fmt.Sprintf("r%v", random_staffid_postfix)
	} else if position == "storekeeper" {
		random_staffid = fmt.Sprintf("s%v", random_staffid_postfix)
	} else if position == "admin" {
		random_staffid = fmt.Sprintf("a%v", random_staffid_postfix)
	}
	return random_staffid, uuid.New().String()
}
