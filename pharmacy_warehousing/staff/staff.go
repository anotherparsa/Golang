package staff

import (
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
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
