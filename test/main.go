package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Staff struct {
	Id       string
	Name     string
	Family   string
	Position string
}

const (
	dbUser     = "testuser"
	dbPassword = "rrc3498urc38r9j999m8j"
	dbAddress  = "127.0.0.1"
	dbPort     = "3306"
	dbName     = "pharmacywarehouse"
)

func main() {
	create("n1", "f1", "p1")
	create("n2", "f2", "p2")
	create("n3", "f3", "p3")
	create("n4", "f4", "p4")

	for i, v := range readAll() {
		fmt.Printf("Instance %v the array is %v \n", i, v)
	}
}

func connectToDatabase() (*sql.DB, error) {
	dataBase, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName))

	if err != nil {
		fmt.Printf("Error in connecting to the database %v\n", err)
	}

	return dataBase, err

}

// crud
// create
func create(name string, family string, position string) {
	dataBase, err := connectToDatabase()

	if err != nil {
		fmt.Printf("Failed to connect to the database %v\n", err)
	}

	defer dataBase.Close()

	querry, err := dataBase.Prepare("INSERT INTO staff (name, family, position) VALUES (?, ?, ?)")

	if err != nil {
		fmt.Printf("Error in preparing the querry %v\n", err)
	}

	defer querry.Close()

	_, err = querry.Exec(name, family, position)

	if err != nil {
		fmt.Printf("Error executing the querry %v\n", err)
	}
}

// read
func readAll() []Staff {
	dataBase, err := connectToDatabase()

	if err != nil {
		fmt.Printf("Failed to connect to the database %v\n", err)
	}

	defer dataBase.Close()

	rows, err := dataBase.Query("SELECT * FROM staff")

	if err != nil {
		fmt.Printf("Error in querrying %v\n", err)
	}

	defer rows.Close()

	staffsArray := []Staff{}
	staffInstance := Staff{}
	for rows.Next() {
		err = rows.Scan(&staffInstance.Id, &staffInstance.Name, &staffInstance.Family, &staffInstance.Position)

		if err != nil {
			fmt.Printf("Error in scanning rows %v\n", err)
			continue
		}
		staffsArray = append(staffsArray, staffInstance)
	}

	if rows.Err() != nil {
		fmt.Printf("Error during the iteration %v\n", rows.Err())
	}

	return staffsArray
}
