package staff

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

type DataToSend struct {
	Staff []model.Staff
}

// handler of "/" and "/staff/home" and "/staff" and "/admin/home" and "/admin"
func Staff_home_page(w http.ResponseWriter, r *http.Request) {
	user, err := session.Is_user_authorized(r, []string{"admin", "recipient", "storekeeper"})
	if err != nil {
		fmt.Printf("Error Staff 1: %v\n", err)
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
	if user.Position == "recipient" {
		err = utility.Render_template(w, "./staff/templates/recipient.html", nil)
		if err != nil {
			fmt.Printf("Error Staff 2: %v\n", err)
		}
	} else if user.Position == "storekeeper" {
		err = utility.Render_template(w, "./staff/templates/warehouse.html", nil)
		if err != nil {
			fmt.Printf("Error Staff 3: %v\n", err)
		}
	} else if user.Position == "admin" {
		err = utility.Render_template(w, "./admin/templates/admin.html", nil)
		if err != nil {
			fmt.Printf("Error Staff 4: %v\n", err)
		}
	} else {
		fmt.Printf("Error Staff 4: %v\n", err)
	}
}

//handler of  "/admin/allstaff"
func All_staff_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err != nil {
		fmt.Printf("Error Staff 5: %v\n", err)
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
	staff_array, err := All_staff()
	if err != nil {
		fmt.Printf("Error Staff 6: %v\n", err)
	}
	data := DataToSend{Staff: staff_array}
	err = utility.Render_template(w, "./admin/templates/allstaff.html", data)
	if err != nil {
		fmt.Printf("Error Staff 7: %v\n", err)
	}
}

func All_staff() ([]model.Staff, error) {
	staff_instance := model.Staff{}
	staff_array := []model.Staff{}
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return staff_array, err
	}
	defer database.Close()
	rows, err := database.Query("SELECT * FROM staff")
	if err != nil {
		return staff_array, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&staff_instance.Id, &staff_instance.Name, &staff_instance.Family, &staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position, &staff_instance.Password)
		if err != nil {
			continue
		}
		staff_array = append(staff_array, staff_instance)
	}
	if rows.Err() != nil {
		return staff_array, err
	}
	return staff_array, err
}
