package login

import (
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/staff"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Login_page(w http.ResponseWriter, r *http.Request) {
	utility.Render_template(w, "./login/templates/login.html")
}

func Login_processor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	staff, err := staff.Get_staff(r.PostForm.Get("staffid"), r.PostForm.Get("password"))

	if err != nil {
		fmt.Printf("Failed to get the staff : %v\n", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		new_uuid := uuid.New().String()
		session.Set_session(w, new_uuid, staff.Userid)
		http.Redirect(w, r, "/staff/home", 302)
	}

}
