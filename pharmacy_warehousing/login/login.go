package login

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Login_page(w http.ResponseWriter, r *http.Request) {
	if session.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/home", http.StatusFound)
	} else {
		err := utility.Render_template(w, "./login/templates/login.html", nil)
		if err != nil {
			fmt.Printf("Error 20: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
	}
}

func Login_processor(w http.ResponseWriter, r *http.Request) {
	if session.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/home", http.StatusFound)
	} else {
		err := r.ParseForm()
		if err == nil {
			staffid := r.PostForm.Get("staffid")
			password := r.PostForm.Get("password")
			userid, err := Authenticate_user(staffid, password)
			if err == nil {
				new_uuid := uuid.New().String()
				err = session.Set_session(w, new_uuid, userid)
				if err == nil {
					http.Redirect(w, r, "/staff/home", http.StatusFound)
				} else {
					fmt.Printf("Error 20: %v\n", err)
					http.Redirect(w, r, "/error", http.StatusFound)
				}
			} else {
				fmt.Printf("Error 20: %v\n", err)
				http.Redirect(w, r, "/error", http.StatusFound)
			}
		} else {
			fmt.Printf("Error 20: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
	}
}

func Authenticate_user(staffid string, password string) (string, error) {
	var userid string
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return userid, err
	}
	defer database.Close()
	querry := "SELECT userid FROM staff WHERE staffid=? AND password=?"
	row := database.QueryRow(querry, staffid, password)
	err = row.Scan(&userid)
	return userid, err
}
