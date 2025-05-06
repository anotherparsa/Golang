package main

import (
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

type Staff struct {
	Id       string
	Name     string
	Family   string
	Staffid  string
	Position string
	password string
}

func main() {
	insert_staff("testname1", "testfamily1", 1, "testpos1", "testpassword1")
	insert_staff("testname2", "testfamily2", 2, "testpos2", "testpassword2")
	insert_staff("testname3", "testfamily3", 3, "testpos3", "testpassword3")
	insert_staff("testname4", "testfamily4", 4, "testpos4", "testpassword4")

	for i, v := range read_all_staff() {
		fmt.Printf("The index is %v and the value is %v\n", i, v)
	}

	delete_staff(2)
	delete_staff(1)

	for i, v := range read_all_staff() {
		fmt.Printf("The index is %v and the value is %v\n", i, v)
	}

	update_staff("staffid", "3", "staffid", "55")

}

func connect() (*sql.DB, error) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name))

	if err != nil {
		fmt.Printf("Error connecting to the database %v\n", err)
	}

	return database, err
}

func insert_staff(name string, family string, staffid int, positions string, password string) {
	database, err := connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, position, password) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Printf("Error preparing the querry: %v\n", err)
	}

	defer querry.Close()

	_, err = querry.Exec(name, family, staffid, positions, password)

	if err != nil {
		fmt.Printf("Error executing the querry: %v\n", err)
	}
}

func read_all_staff() []Staff {
	database, err := connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	defer database.Close()

	rows, err := database.Query("SELECT * FROM staff")

	if err != nil {
		fmt.Printf("Error in querrying %v\n", err)
	}

	defer rows.Close()

	staffInstance := Staff{}
	staffArray := []Staff{}

	for rows.Next() {
		err = rows.Scan(&staffInstance.Id, &staffInstance.Name, &staffInstance.Family, &staffInstance.Staffid, &staffInstance.Position, &staffInstance.password)

		if err != nil {
			fmt.Printf("Error in scanning the rows %v\n", err)
			continue
		}

		staffArray = append(staffArray, staffInstance)
	}

	if rows.Err() != nil {
		fmt.Printf("Error during the iteration %v\n", rows.Err())
	}

	return staffArray

}

func delete_staff(staffid int) {
	database, err := connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("DELETE FROM staff WHERE staffid=?")

	if err != nil {
		fmt.Printf("Error in preparing the querry: %v\n", err)
	}

	defer querry.Close()

	_, err = querry.Exec(staffid)

	if err != nil {
		fmt.Printf("Error executing the querry: %v\n", err)
	}
}

func update_staff(condition string, conditionValue string, updatecolumn string, updatedvalue string) {
	database, err := connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare(fmt.Sprintf("UPDATE staff SET %s=? WHERE %s=?", updatecolumn, condition))

	if err != nil {
		fmt.Printf("Error in preparing the querry: %v", err)
	}

	defer querry.Close()

	_, err = querry.Exec(updatedvalue, conditionValue)

	if err != nil {
		fmt.Printf("Error in executing the querry: %v\n", err)
	}

}
