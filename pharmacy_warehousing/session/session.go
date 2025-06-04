package session

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"fmt"
	"net/http"
)

func Set_session(w http.ResponseWriter, sessionid string, userid string) error {
	//setting the cookie on the browser
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionid",
		Value: sessionid,
		Path:  "/",
	})
	//calling the function to create the session record in the database
	err := Create_session_record(userid, sessionid)
	return err
}

func Create_session_record(userid string, sessionid string) error {
	//connecting to the database
	database, err := databasetool.Connect_to_database()
	//error in connecting to the database
	if err != nil {
		return err
	}
	defer database.Close()
	//preparing the querry
	querry, err := database.Prepare("INSERT INTO session (userid, sessionid) VALUES (?, ?)")
	//error in preparing the querry
	if err != nil {
		return err
	}
	defer querry.Close()
	//executing the querry
	_, err = querry.Exec(userid, sessionid)
	//error in executing the querry
	if err != nil {
		return err
	}
	return nil
}

func Delete_session_record(sessionid string) error {
	//connecting to the database
	database, err := databasetool.Connect_to_database()
	//error in connecting to the database
	if err != nil {
		return err
	}
	defer database.Close()
	//preparing the querry
	querry, err := database.Prepare("DELETE FROM session WHERE sessionid=?")
	//error in preparing the querry
	if err != nil {
		return err
	}
	defer querry.Close()
	//executing the querry
	_, err = querry.Exec(sessionid)
	if err != nil {
		return err
	}
	return nil
}
func Check_if_cookie_exists(r *http.Request, cookiename string) bool {
	_, err := r.Cookie(cookiename)
	return err == nil
}

func Is_user_authorized(r *http.Request, authorized_positions []string) (model.Staff, error) {
	staff_instance := model.Staff{}
	var userid string
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		return staff_instance, err
	}
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return staff_instance, err
	}
	defer database.Close()
	querry := "SELECT userid FROM session WHERE sessionid=?"
	row := database.QueryRow(querry, cookie.Value)
	err = row.Scan(&userid)
	if err != nil {
		return staff_instance, err
	}
	querry = "SELECT * FROM staff WHERE userid=?"
	row = database.QueryRow(querry, userid)
	err = row.Scan(&staff_instance.Id, &staff_instance.Name, &staff_instance.Family, &staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position, &staff_instance.Password)
	//error in scanning the row
	if err != nil {
		return staff_instance, err
	}
	for _, v := range authorized_positions {
		if staff_instance.Position == v {
			return staff_instance, err
		}
	}
	return staff_instance, err
}

// handler of "/staff/logout"
func User_logout(w http.ResponseWriter, r *http.Request) {
	if !Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	} else {
		cookie, err := r.Cookie("sessionid")
		if err != nil {
			fmt.Printf("Error 25: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
		err = Delete_session_record(cookie.Value)
		if err != nil {
			fmt.Printf("Error 26: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
		http.SetCookie(w, &http.Cookie{
			Name:   "sessionid",
			MaxAge: -1,
			Path:   "/",
		})
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
}
