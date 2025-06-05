package admin

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"bufio"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// handler of "/admin/addstaff"
func Admin_add_staff_page(w http.ResponseWriter, r *http.Request) error {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err == nil {
		err = utility.Render_template(w, "./admin/templates/addstaff.html", nil)
		return err
	}
	return err
}

// handler of "/admin/addstaffprocessor"
func Admin_add_staff_processor(w http.ResponseWriter, r *http.Request) error {
	_, err := session.Is_user_authorized(r, []string{"admin"})
	if err == nil {
		err = r.ParseForm()
		if err == nil {
			name := r.PostForm.Get("staffname")
			family := r.PostForm.Get("stafffamily")
			position := r.PostForm.Get("position")
			password := r.PostForm.Get("password")
			err = Create_staff_record(name, family, position, password)
			if err == nil {
				http.Redirect(w, r, "/staff/home", http.StatusFound)
				return err
			}
		}
	}
	return err
}

func Create_staff_record(name string, family string, position string, password string) error {
	random_staffid_postfix := strconv.Itoa(rand.IntN(99999-10000) + 10000)
	var random_staffid string
	if position == "recipient" {
		random_staffid = fmt.Sprintf("r%v", random_staffid_postfix)
	} else if position == "storekeeper" {
		random_staffid = fmt.Sprintf("s%v", random_staffid_postfix)
	} else if position == "admin" {
		random_staffid = fmt.Sprintf("a%v", random_staffid_postfix)
	}
	random_userid := uuid.New().String()
	database, err := databasetool.Connect_to_database()
	if err == nil {
		defer database.Close()
		querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, userid, position, password) VALUES (?, ?, ?, ?, ?, ?)")
		if err == nil {
			defer querry.Close()
			_, err = querry.Exec(name, family, random_staffid, random_userid, position, password)
			return err
		}
	}
	return err
}

func Create_admin_user() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name for Admin user : ")
	name, err := reader.ReadString('\n')
	if err == nil {
		name = strings.Replace(name, "\n", "", -1)
		fmt.Println("Family for Admin user : ")
		family, err := reader.ReadString('\n')
		if err == nil {
			family = strings.Replace(family, "\n", "", -1)
			fmt.Println("Password for Admin user : ")
			password, err := reader.ReadString('\n')
			if err == nil {
				password = strings.Replace(password, "\n", "", -1)
				err = Create_staff_record(name, family, "admin", password)
				return err
			}
		}
	}
	return err
}
