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

func Staff_home_page_handler(w http.ResponseWriter, r *http.Request) {
	if utility.Is_user_logged(r) {
		cookie, err := r.Cookie("sessionid")
		if err == nil {
			user, err := session.User_with_sessionid(cookie.Value)
			if err == nil {
				if user.Position == "recipient" {
					utility.TemplateRendering(w, "./staff/templates/recipient.html")
				} else if user.Position == "storekeeper" {
					utility.TemplateRendering(w, "./staff/templates/warehouse.html")
				} else {
					fmt.Printf("Unauthorized user \n")
				}
			} else {
				fmt.Printf("failed to get the user with the sessionid %v\n", err)
			}
		} else {
			fmt.Printf("Failed to get the cookie : %v\n", err)
		}
	} else {
		fmt.Printf("User is not logged in!\n")
		http.Redirect(w, r, "/login", 302)
	}
}

func Create_staff(w http.ResponseWriter, r *http.Request, name string, family string, staffid string, position string, password string) error {
	if !utility.Is_user_logged() {
		http.Redirect(w, r, "/login", http.StatusFound)
		return session.User_not_logged_in_error
	}

	database, err := databasetool.Connect()

	if err != nil {
		return databasetool.Database_connection_failure
	}

	defer database.Close()

	uuid := uuid.New().String()

	querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, userid, position, password) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return databasetool.Querry_prepration_failure
	}

	defer querry.Close()

	_, err = querry.Exec(name, family, staffid, uuid, position, password)

	if err != nil {
		return databasetool.Querry_execution_failure
	}

	return nil
}

func Read_all_staff() ([]model.Staff, error) {
	staff_array := []model.Staff{}

	if !utility.Is_user_logged() {
		http.Redirect(w, r, "/login", http.StatusFound)
		return staff_array, session.User_not_logged_in_error
	}

	database, err := databasetool.Connect()

	if err != nil {
		return staff_array, databasetool.Database_connection_failure
	}

	defer database.Close()

	rows, err := database.Query("SELECT staffid, userid, position FROM staff")
	
	if err != nil {
		return staff_array, databasetool.Querry_execution_failure
	}

	defer rows.Close()

	staff_instance := model.Staff{}

	for rows.Next() {
		err = rows.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)
		
		if err != nil {
			return staff_array, databasetool.Scanning_rows_failure
		}

		staff_array = append(staff_array, staff_instance)
	}

	return staff_array, nil
}

func Get_staff(staffid string, password string) (model.Staff, error) {
	if !utility.Is_user_logged(){
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	
	database, err := databasetool.Connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry := "SELECT staffid, userid, position FROM staff WHERE staffid=? AND password=?"

	row := database.QueryRow(querry, staffid, password)

	staff_instance := model.Staff{}

	err = row.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)

	if err != nil {
		fmt.Printf("Failed to scan the row : %v\n", err)
	}

	return staff_instance, err
}
