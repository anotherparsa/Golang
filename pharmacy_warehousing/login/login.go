package login

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Login_page(w http.ResponseWriter, r *http.Request) {
	if model.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/home", http.StatusFound)
	} else {
		utility.Render_template(w, "./login/templates/login.html")
	}
}

func Login_processor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	staffid := r.PostForm.Get("staffid")
	password := r.PostForm.Get("password")
	userid, err := Authenticate_user(staffid, password)

	if err != nil {
		fmt.Printf("Failed to authenticate : %v\n", err)
		http.Redirect(w, r, "/staff/login", 302)
	} else {
		new_uuid := uuid.New().String()
		session.Set_session(w, new_uuid, userid)
		http.Redirect(w, r, "/staff/home", 302)
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

	staffinstance := model.Staff{}

	err = row.Scan(&staffinstance.Userid)

	if err != nil {
		return userid, err
	}

	return staffinstance.Userid, err
}
