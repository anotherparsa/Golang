package admin

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type DataToSend struct {
	Staff interface{}
}

// handler of "/admin/addstaff"
func Admin_add_staff_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	err = utility.Render_template(w, "./admin/templates/addstaff.html", nil)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
}

// handler of "/admin/addstaffprocessor"
func Admin_add_staff_processor(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	err = r.ParseForm()
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	name := r.PostForm.Get("staffname")
	family := r.PostForm.Get("stafffamily")
	position := r.PostForm.Get("position")
	password := r.PostForm.Get("password")
	random_staffid, random_userid := utility.Generate_staffid_userid(position)
	err = Create_staff_record(name, family, random_staffid, random_userid, position, password)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	data := DataToSend{Staff: model.Staff{
		Name:     name,
		Family:   family,
		Staffid:  random_staffid,
		Position: position,
		Password: password,
	}}
	err = utility.Render_template(w, "./admin/templates/staffcreation.html", data)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
}

// handler of "admin/editstaff/?""
func Admin_edit_staff_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	user, err := Get_staff_by("id", strings.TrimPrefix(r.URL.Path, "/admin/editstaff/"))
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	data := DataToSend{Staff: user}
	err = utility.Render_template(w, "./admin/templates/editstaffpage.html", data)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
}

func Create_staff_record(name string, family string, random_staffid string, random_userid string, position string, password string) error {
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return err
	}
	defer database.Close()
	querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, userid, position, password) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(name, family, random_staffid, random_userid, position, password)
	if err != nil {
		return err
	}
	return nil
}

func Create_admin_user() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name for Admin user : ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println("Family for Admin user : ")
	family, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	family = strings.Replace(family, "\n", "", -1)
	fmt.Println("Password for Admin user : ")
	password, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	password = strings.Replace(password, "\n", "", -1)
	random_staffid, random_userid := utility.Generate_staffid_userid("admin")
	err = Create_staff_record(name, family, random_staffid, random_userid, "admin", password)
	if err != nil {
		return err
	}
	fmt.Printf("Staff id : %v\n", random_staffid)
	return nil
}

// handler of  "/admin/allstaff"
func All_staff_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	staff_array, err := All_staff()
	if err != nil {
		fmt.Println("We've reached here2")
		utility.Error_handler(w, err.Error())
		return
	}
	data := DataToSend{Staff: staff_array}
	err = utility.Render_template(w, "./admin/templates/allstaff.html", data)
	if err != nil {
		fmt.Println("We've reached here3")
		utility.Error_handler(w, err.Error())
		return
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
		if err == nil {
			staff_array = append(staff_array, staff_instance)
		}
	}
	return staff_array, rows.Err()
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

func Admin_edit_staff_processor(w http.ResponseWriter, r *http.Request) {
	user, err := session.Is_user_authorized(r, []string{"admin"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	err = r.ParseForm()
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	name := r.PostForm.Get("name")
	familt := r.PostForm.Get("family")
	staffid := r.PostForm.Get("staffid")
	previous_staffid := staffid
	position := r.PostForm.Get("position")
	password := r.PostForm.Get("password")
	id := r.PostForm.Get("id")
	if staffid[0] != position[0] {
		staffid, _ = utility.Generate_staffid_userid(position)
	}
	err = Edit_staff_record(id, name, familt, staffid, position, password)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	if previous_staffid == user.Staffid {
		err = session.User_logout(w, r)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
	} else {
		http.Redirect(w, r, "/admin/allstaff", http.StatusFound)
	}
}

func Delete_staff_record(w http.ResponseWriter, r *http.Request) {
	user, err := session.Is_user_authorized(r, []string{"admin"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	staffsid := strings.TrimPrefix(r.URL.Path, "/admin/deletestaff/")
	database, err := databasetool.Connect_to_database()
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	defer database.Close()
	querry, err := database.Prepare("DELETE FROM staff WHERE id=?")
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	defer querry.Close()
	_, err = querry.Exec(staffsid)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	if staffsid == user.Id {
		err = session.User_logout(w, r)
		if err != nil {
			utility.Error_handler(w, err.Error())
			return
		}
	}
	http.Redirect(w, r, "/admin/allstaff", http.StatusFound)
}
