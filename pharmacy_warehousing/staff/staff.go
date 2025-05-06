package staff

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

func Staff_home_page_handler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "./staff/templates/home.html")
}

func Create_staff(name string, family string, staffid string, position string, password string) {
	database, err := databasetool.Connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, position, password) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Printf("Failed to prepare the querry : %v\n", err)
	}

	defer querry.Close()

	_, err = querry.Exec(name, family, staffid, position, password)

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
