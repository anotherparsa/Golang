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

func Create_session(userid string, sessionid string) {
	database, err := databasetool.Connect()

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

func User_with_sessionid(sessionid string) (model.Staff, error) {
	database, err := databasetool.Connect()
	fmt.Printf("Session id is %v \n", sessionid)
	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry := "SELECT userid FROM session WHERE sessionid=?"
	fmt.Println(querry)
	row := database.QueryRow(querry, sessionid)
	var userid string
	fmt.Printf("user id is %v\n", userid)
	err = row.Scan(&userid)

	if err != nil {
		fmt.Printf("Failed to get the user id : %v\n", err)
	}

	querry = "SELECT staffid, userid, position FROM staff WHERE userid=?"

	row = database.QueryRow(querry, userid)

	staff_instance := model.Staff{}
	fmt.Printf("Session id is %v \n", sessionid)
	err = row.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)

	fmt.Println("HEre ")
	fmt.Println(staff_instance)
	return staff_instance, err

}
