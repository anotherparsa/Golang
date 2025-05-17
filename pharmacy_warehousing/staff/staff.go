package staff

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func Staff_home_page(w http.ResponseWriter, r *http.Request) {
	//checking to see if user already logged in or not
	if !session.Check_if_cookie_exists(r, "sessionid") {
		http.Redirect(w, r, "/staff/login", http.StatusFound)
		return
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
				err = utility.Render_template(w, "./staff/templates/recipient.html")
				if err != nil {
					fmt.Printf("Error1 : %v\n", err)
				}
			} else if user.Position == "storekeeper" {
				err = utility.Render_template(w, "./staff/templates/warehouse.html")
				if err != nil {
					fmt.Printf("Error2 : %v\n", err)
				}
			} else if user.Position == "admin" {
				err = utility.Render_template(w, "./admin/templates/admin.html")
				if err != nil {
					fmt.Printf("Error 3: %v\n", err)
				}
			} else {
				fmt.Printf("Unauthorized user \n")
			}
		} else {
			fmt.Printf("Error 4: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
	} else {
		fmt.Printf("Error 5: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
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
