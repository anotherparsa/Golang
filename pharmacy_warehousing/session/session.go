package session

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/utility"
	"errors"
	"fmt"
	"net/http"
	"slices"
)

func Set_session(w http.ResponseWriter, sessionid string, userid string) error {
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionid",
		Value: sessionid,
		Path:  "/",
	})
	err := Create_session_record(userid, sessionid)
	return err
}

func Create_session_record(userid string, sessionid string) error {
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

func Delete_session_record(condition string, condition_value string) error {
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return err
	}
	defer database.Close()
	querry, err := database.Prepare(fmt.Sprintf("DELETE FROM session WHERE %v=?", condition))
	if err != nil {
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(condition_value)
	if err != nil {
		return err
	}
	return nil
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
	if err != nil {
		return staff_instance, err
	}
	if slices.Contains(authorized_positions, staff_instance.Position) {
		return staff_instance, nil
	} else {
		return staff_instance, errors.New("")
	}
}

// handler of "/staff/logout"
func User_logout(w http.ResponseWriter, r *http.Request) {
	_, err := Is_user_authorized(r, []string{"admin", "recipient", "storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	err = Delete_session_record("sessionid", cookie.Value)
	if err != nil {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "sessionid",
		MaxAge: -1,
		Path:   "/",
	})
	http.Redirect(w, r, "/staff/login", http.StatusFound)

}
