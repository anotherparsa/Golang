package staff

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"net/http"
)

type DataToSend struct {
	Staff interface{}
}

// handler of "/" and "/staff/home" and "/staff" and "/admin/home" and "/admin"
func Staff_home_page(w http.ResponseWriter, r *http.Request) error {
	user, err := session.Is_user_authorized(r, []string{"admin", "recipient", "storekeeper"})
	if err == nil {
		if user.Position == "recipient" {
			err = utility.Render_template(w, "./staff/templates/recipient.html", nil)
			return err
		} else if user.Position == "storekeeper" {
			err = utility.Render_template(w, "./staff/templates/warehouse.html", nil)
			return err
		} else if user.Position == "admin" {
			err = utility.Render_template(w, "./admin/templates/admin.html", nil)
			return err
		}
	}
	return err
}

// handler of  "/admin/allstaff"
func All_staff_page(w http.ResponseWriter, r *http.Request) error {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err == nil {
		staff_array, err := All_staff()
		if err == nil {
			data := DataToSend{Staff: staff_array}
			err = utility.Render_template(w, "./admin/templates/allstaff.html", data)
			return err
		}
	}
	return err
}

func All_staff() ([]model.Staff, error) {
	staff_instance := model.Staff{}
	staff_array := []model.Staff{}
	database, err := databasetool.Connect_to_database()
	if err == nil {
		defer database.Close()
		rows, err := database.Query("SELECT * FROM staff")
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&staff_instance.Id, &staff_instance.Name, &staff_instance.Family, &staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position, &staff_instance.Password)
				if err == nil {
					staff_array = append(staff_array, staff_instance)
				}
			}
		}
	}
	return staff_array, err
}
