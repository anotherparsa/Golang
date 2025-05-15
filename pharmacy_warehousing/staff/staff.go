package staff

import (
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

func Staff_home_page(w http.ResponseWriter, r *http.Request) {
	//checking to see if user already logged in or not
	if !model.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
	//user is logged in
	//getting the cookie
	cookie, err := r.Cookie("sessionid")
	if err == nil {
		//getting the user associated with that sessionid
		user, err := session.User_with_sessionid(cookie.Value)
		if err == nil {
			//showing different home pages according to the positions
			if user.Position == "recipient" {
				utility.Render_template(w, "./staff/templates/recipient.html")
			} else if user.Position == "storekeeper" {
				utility.Render_template(w, "./staff/templates/warehouse.html")
			} else {
				fmt.Printf("Unauthorized user \n")
			}
		} else {
			fmt.Printf("failed to get the user with the sessionid %v\n", err)
		}
	} else {
		fmt.Printf("Failed to get the cookie : %v\n", err)
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
}
