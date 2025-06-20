package login

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"net/http"

	"github.com/google/uuid"
)

func Login_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"storekeeper", "admin", "recipient"})
	if err == nil {
		http.Redirect(w, r, "/staff/home", http.StatusFound)
	} else {
		err := utility.Render_template(w, "./login/templates/login.html", nil)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
	}
}

func Login_processor(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"storekeeper", "admin", "recipient"})
	if err == nil {
		http.Redirect(w, r, "/staff/home", http.StatusFound)
	} else {
		err := r.ParseForm()
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
		staffid := r.PostForm.Get("staffid")
		password := r.PostForm.Get("password")
		userid, err := Authenticate_user(staffid, password)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
		new_uuid := uuid.New().String()
		err = session.Set_session(w, new_uuid, userid)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
		http.Redirect(w, r, "/staff/home", http.StatusFound)
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
	if err != nil {
		return userid, err
	}
	return userid, nil
}
