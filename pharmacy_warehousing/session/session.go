package session

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"errors"
	"net/http"
)

//sets a session with sessionid and userid passed to it and returns an error
func Set_session(w http.ResponseWriter, sessionid string, userid string) error {
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionid",
		Value: sessionid,
		Path:  "/",
	})
	err := Create_session(userid, sessionid)
	return err
}

//creates session record in the database with userid and sessionid
func Create_session(userid string, sessionid string) error {
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return err
	}
	defer database.Close()
	querry, err := database.Prepare("INSERT INTO session (userid, sessionid) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(userid, sessionid)
	if err != nil {
		return err
	}
	return nil
}

//returns the user associated with that session id and returns with an instance of the staff and error
func User_with_sessionid(sessionid string) (model.Staff, error) {
	database, err := databasetool.Connect_to_database()
	staff_instance := model.Staff{}
	if err != nil {
		return staff_instance, err
	}
	defer database.Close()
	//gets the userid associated with that sessionid in session table
	querry := "SELECT userid FROM session WHERE sessionid=?"
	row := database.QueryRow(querry, sessionid)
	var userid string
	err = row.Scan(&userid)

	if err != nil {
		return staff_instance, err
	}
	//gets the user associated with that userid from staff table
	querry = "SELECT staffid, userid, position FROM staff WHERE userid=?"
	row = database.QueryRow(querry, userid)
	err = row.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)
	if err != nil {
		return staff_instance, err
	}
	return staff_instance, err
}

//checks if a specific cookie exist on the browser or not and returns a boolian
func Check_if_cookie_exists(r *http.Request, cookiename string) bool {
	_, err := r.Cookie(cookiename)
	return err == nil
}

//checks if the user is authorized to access that path or not and returns an error
func Is_user_authorized(r *http.Request, position string) error {
	//gets the sessionid cookie
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		return err
	}
	//gets the user with that sessionid
	user, err := User_with_sessionid(cookie.Value)
	if err != nil {
		return err
	}
	//checks if the user's id is the same as is it should be
	if user.Position == position {
		return nil
	} else {
		return errors.New("user not authorized")
	}
}
