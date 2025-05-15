package session

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"fmt"
	"net/http"
)

func Set_session(w http.ResponseWriter, sessionid string, userid string) {

	http.SetCookie(w, &http.Cookie{
		Name:  "sessionid",
		Value: sessionid,
		Path:  "/",
	})
	Create_session(userid, sessionid)
}

func User_with_sessionid(sessionid string) (model.Staff, error) {
	database, err := databasetool.Connect_to_database()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry := "SELECT userid FROM session WHERE sessionid=?"

	row := database.QueryRow(querry, sessionid)
	var userid string

	err = row.Scan(&userid)

	if err != nil {
		fmt.Printf("Failed to get the user id : %v\n", err)
	}

	querry = "SELECT staffid, userid, position FROM staff WHERE userid=?"
	row = database.QueryRow(querry, userid)
	staff_instance := model.Staff{}
	err = row.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)

	if err != nil {
		fmt.Printf("Failed to scan the row : %v\n", err)
	}

	return staff_instance, err

}

func Is_user_authorized(sessionid string, position string) (model.Staff, error) {
	user, err := User_with_sessionid(sessionid)
	if err != nil {
		fmt.Printf("Failed to get the user : %v\n", err)
		return user, err
	}
	if user.Position == position {
		return user, nil
	} else {
		return user, err
	}

}

// session
func Create_session(userid string, sessionid string) {
	database, err := databasetool.Connect_to_database()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO session (userid, sessionid) VALUES (?, ?)")

	if err != nil {
		fmt.Printf("Failed to prepare the querry : %v\n", err)
	}
	defer querry.Close()

	_, err = querry.Exec(userid, sessionid)

	if err != nil {
		fmt.Printf("Failed to execute the querry : %v\n", err)
	}
}

func Check_if_cookie_exists(r *http.Request, cookiename string) (string bool) {
	_, err := r.Cookie(cookiename)
	return err == nil

}
