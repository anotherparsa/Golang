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
	//checking if user already logged
	if session.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/home", http.StatusFound)
	} else {
		err := utility.Render_template(w, "./login/templates/login.html")
		if err != nil {
			fmt.Printf("Error 20: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
	}
}

func Login_processor(w http.ResponseWriter, r *http.Request) {
	//checking if user already logged
	if session.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/home", http.StatusFound)
	} else {
		//parsing the form
		r.ParseForm()
		//gerring data from the form
		staffid := r.PostForm.Get("staffid")
		password := r.PostForm.Get("password")
		//authentticating the user with staff id and returning its userid
		userid, err := Authenticate_user(staffid, password)
		if err != nil {
			//authentication failed
			fmt.Printf("Error 21: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		} else {
			//authentication was successful and a session will be set
			new_uuid := uuid.New().String()
			err = session.Set_session(w, new_uuid, userid)
			if err != nil {
				fmt.Printf("Error 22: %v\n", err)
				http.Redirect(w, r, "/error", http.StatusFound)
			}
			http.Redirect(w, r, "/staff/home", http.StatusFound)
		}
	}

}

func Authenticate_user(staffid string, password string) (string, error) {
	var userid string
	//connecting to the database
	database, err := databasetool.Connect_to_database()
	//error in connecting to the database
	if err != nil {
		return userid, err
	}
	defer database.Close()
	//getting the row
	querry := "SELECT userid FROM staff WHERE staffid=? AND password=?"
	row := database.QueryRow(querry, staffid, password)
	err = row.Scan(&userid)
	//error in scanning the row
	if err != nil {
		return userid, err
	}
	return userid, err
}
