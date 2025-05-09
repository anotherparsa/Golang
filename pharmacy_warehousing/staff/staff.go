package staff

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Staff_home_page_handler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")

	if err != nil {
		fmt.Printf("Failed to read the cookie : %v\n", err)
	}
	fmt.Println("We've reached home")
	fmt.Println(cookie.Value)
	user, _ := session.User_with_sessionid(cookie.Value)
	fmt.Printf("Use position is %v\n", user)
	if user.Position == "recipe" {
		utility.TemplateRendering(w, "./staff/templates/recipient.html")
		fmt.Println("We've reached reci home")
	} else if user.Position == "storekeeper" {
		fmt.Println("We've reached ware home")
		utility.TemplateRendering(w, "./staff/templates/warehouse.html")
	}
}

func Create_staff(name string, family string, staffid string, position string, password string) {
	database, err := databasetool.Connect()
	new_uuid := uuid.New().String()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, userid, position, password) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Printf("Failed to prepare the querry : %v\n", err)
	}

	defer querry.Close()

	_, err = querry.Exec(name, family, staffid, new_uuid, position, password)

	if err != nil {
		fmt.Printf("Failed to execute the queey : %v\n", err)
	}
}

func Read_all_staff() []model.Staff {
	database, err := databasetool.Connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	rows, err := database.Query("SELECT * FROM staff")

	if err != nil {
		fmt.Printf("Failed to querry the database : %v\n", err)
	}

	defer rows.Close()

	staff_instance := model.Staff{}
	staff_array := []model.Staff{}

	for rows.Next() {
		err = rows.Scan(&staff_instance.Id, &staff_instance.Name, &staff_instance.Family, &staff_instance.Staffid, &staff_instance.Position, &staff_instance.Password)

		if err != nil {
			fmt.Printf("Failed to scan rows : %v\n", err)
			continue
		}

		staff_array = append(staff_array, staff_instance)
	}

	return staff_array
}

func Get_staff(staffid string, password string) (model.Staff, error) {
	database, err := databasetool.Connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry := "SELECT staffid, userid, position FROM staff WHERE staffid=? AND password=?"

	row := database.QueryRow(querry, staffid, password)

	staff_instance := model.Staff{}

	err = row.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)

	if err != nil {
		fmt.Printf("Failed to scan the row : %v\n", err)
	}

	return staff_instance, err
}
