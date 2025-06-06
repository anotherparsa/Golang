package staff

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

// handler of "/" and "/staff/home" and "/staff" and "/admin/home" and "/admin"
func Staff_home_page(w http.ResponseWriter, r *http.Request) {
	user, err := session.Is_user_authorized(r, []string{"admin", "recipient", "storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	if user.Position == "recipient" {
		err = utility.Render_template(w, "./staff/templates/recipient.html", nil)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
	} else if user.Position == "storekeeper" {
		err = utility.Render_template(w, "./staff/templates/warehouse.html", nil)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
	} else if user.Position == "admin" {
		err = utility.Render_template(w, "./admin/templates/admin.html", nil)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
	}
}

func Get_staff_by(condition string, condition_value string) (model.Staff, error) {
	staff_instance := model.Staff{}
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return staff_instance, err
	}
	defer database.Close()
	row := database.QueryRow(fmt.Sprintf("SELECT * FROM staff WHERE %v=?", condition), condition_value)
	err = row.Scan(&staff_instance.Id, &staff_instance.Name, &staff_instance.Family, &staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position, &staff_instance.Password)
	if err != nil {
		return staff_instance, err
	}
	return staff_instance, nil
}

func Edit_staff_record(id string, name string, family string, random_staffid string, position string, password string) error {
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return err
	}
	defer database.Close()
	querry, err := database.Prepare("UPDATE staff SET name=?, family=?, staffid=?, position=?, password=? WHERE id=?")
	if err != nil {
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(name, family, random_staffid, position, password, id)
	if err != nil {
		return err
	}
	return nil
}
