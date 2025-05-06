package admin

import (
	"PharmacyWarehousing/staff"
	"PharmacyWarehousing/utility"
	"net/http"
)

func Admin_home_page_handler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "./admin/templates/home.html")
}

func Admin_add_staff_page_handler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "./admin/templates/addstaff.html")
}

func Admin_add_staff_processor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	staff.Create_staff(r.PostForm.Get("name"), r.PostForm.Get("family"), "0", r.PostForm.Get("position"), r.PostForm.Get("password"))
	http.Redirect(w, r, "/admin/home", 302)
}
