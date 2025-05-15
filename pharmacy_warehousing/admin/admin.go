package admin

import (
	"PharmacyWarehousing/staff"
	"PharmacyWarehousing/utility"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Admin_home_page(w http.ResponseWriter, r *http.Request) {
	utility.Render_template(w, "./admin/templates/home.html")
}

func Admin_add_staff_page(w http.ResponseWriter, r *http.Request) {
	utility.Render_template(w, "./admin/templates/addstaff.html")
}

func Admin_add_staff_processor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Printf("Failed to parse the form : %v\n", err)
	}

	name := r.PostForm.Get("name")
	family := r.PostForm.Get("family")
	position := r.PostForm.Get("position")
	password := r.PostForm.Get("password")

	err = staff.Create_staff_record(name, family, position, password)

	if err != nil {
		fmt.Printf("Failed to create the staff : %v\n", err)
	}

	http.Redirect(w, r, "/admin/home", http.StatusFound)
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

	staff.Create_staff_record(name, family, "admin", password)

	return err
}
