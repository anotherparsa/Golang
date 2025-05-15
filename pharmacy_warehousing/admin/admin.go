package admin

import (
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/utility"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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
	random_staffid := strconv.Itoa(rand.Intn(99999))
	position := r.PostForm.Get("position")
	random_userid := uuid.New().String()
	password := r.PostForm.Get("password")

	if position == "recipient" {
		random_staffid = fmt.Sprintf("r%v", random_staffid)
	} else if position == "storekeeper" {
		random_staffid = fmt.Sprintf("s%v", random_staffid)
	}

	err = model.Create_staff(name, family, random_staffid, random_userid, position, password)

	if err != nil {
		fmt.Printf("Failed to create the staff : %v\n", err)
	}

	http.Redirect(w, r, "/admin/home", http.StatusFound)
}
