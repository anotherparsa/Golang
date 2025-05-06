package databasetool

import (
	"PharmacyWarehousing/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "testuser"
	db_password = "rrc3498urc38r9j999m8j"
	db_address  = "127.0.0.1"
	db_port     = "3306"
	db_name     = "pharmacywarehouse"
)

func connect() (*sql.DB, error) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name))

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	return database, err
}

func Create_staff(name string, family string, staffid string, position string, password string) {
	database, err := connect()

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
		fmt.Printf("Failed to execute the querry : %v\n", err)
	}
}

func Read_all_staff() []model.Staff {
	database, err := connect()

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
		err := rows.Scan(&staff_instance.Id, &staff_instance.Name, &staff_instance.Family, &staff_instance.Staffid, &staff_instance.Position, &staff_instance.Password)

		if err != nil {
			fmt.Printf("Failed to scan the row : %v\n", err)
			continue
		}

		staff_array = append(staff_array, staff_instance)
	}

	if rows.Err() != nil {
		fmt.Printf("Failed during iteration : %v\n", err)
	}

	return staff_array
}