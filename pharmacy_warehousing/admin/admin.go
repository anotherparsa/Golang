package admin

import (
	"PharmacyWarehousing/staff"
	"PharmacyWarehousing/utility"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func Admin_home_page_handler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "./admin/templates/home.html")
}

func Admin_add_staff_page_handler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "./admin/templates/addstaff.html")
}

func Admin_add_staff_processor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	random_staffid := strconv.Itoa(rand.Intn(99999))
	position := r.PostForm.Get("position")

	if position == "recipe" {
		random_staffid = fmt.Sprintf("r%v", random_staffid)
	} else if position == "storekeeper" {
		random_staffid = fmt.Sprintf("s%v", random_staffid)
	}

	staff.Create_staff(w, r, r.PostForm.Get("name"), r.PostForm.Get("family"), random_staffid, position, r.PostForm.Get("password"))
	http.Redirect(w, r, "/admin/home", 302)
}
