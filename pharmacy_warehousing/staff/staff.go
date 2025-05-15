package staff

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Staff_home_page(w http.ResponseWriter, r *http.Request) {
	//checking to see if user already logged in or not
	if session.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
	//user is logged in
	//getting the cookie
	cookie, err := r.Cookie("sessionid")
	if err == nil {
		//getting the user associated with that sessionid
		user, err := session.User_with_sessionid(cookie.Value)
		if err == nil {
			//showing different home pages according to the positions
			if user.Position == "recipient" {
				utility.Render_template(w, "./staff/templates/recipient.html")
			} else if user.Position == "storekeeper" {
				utility.Render_template(w, "./staff/templates/warehouse.html")
			} else {
				fmt.Printf("Unauthorized user \n")
			}
		} else {
			fmt.Printf("failed to get the user with the sessionid %v\n", err)
		}
	} else {
		fmt.Printf("Failed to get the cookie : %v\n", err)
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
}

// staff
func Create_staff(name string, family string, staffid string, userid string, position string, password string) error {
	database, err := databasetool.Connect_to_database()

	if err != nil {
		return err
	}

	defer database.Close()

	uuid := uuid.New().String()

	querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, userid, position, password) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer querry.Close()

	_, err = querry.Exec(name, family, staffid, uuid, position, password)

	if err != nil {
		return err
	}

	return nil
}

func Read_all_staff() ([]model.Staff, error) {
	staff_array := []model.Staff{}

	database, err := databasetool.Connect_to_database()

	if err != nil {
		return staff_array, err
	}

	defer database.Close()

	rows, err := database.Query("SELECT staffid, userid, position FROM staff")

	if err != nil {
		return staff_array, err
	}

	defer rows.Close()

	staff_instance := model.Staff{}

	for rows.Next() {
		err = rows.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)

		if err != nil {
			return staff_array, err
		}

		staff_array = append(staff_array, staff_instance)
	}

	return staff_array, nil
}
