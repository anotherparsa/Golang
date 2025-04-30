package databasetool

import (
	"PharmacyWarehousing/model"
	"database/sql"
	"fmt"
)

const (
	username = "testuser"
	password = "rrc3498urc38r9j999m8j"
	hostname = "127.0.0.1"
	port     = "3306"
	database = "pharmacywarehouse"
)

func connect() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	return sql.Open("mysql", dataSourceName)
}

func SelectAllStaff() []model.Staff {
	dataBase, err := connect()

	if err != nil {
		fmt.Printf("Error connecting to the database %v\n", err)
	}

	rows, err := dataBase.Query("SELECT * FROM staff")

	if err != nil {
		fmt.Printf("Error querying database %v\n", err)
	}

	staffArray := []model.Staff{}
	staffInstance := model.Staff{}

	for rows.Next() {
		err = rows.Scan(&staffInstance.Id, &staffInstance.Name, &staffInstance.Family, &staffInstance.Position)
		if err != nil {
			fmt.Printf("Error scanning rows %v\n", err)
			continue
		}
		staffArray = append(staffArray, staffInstance)
	}

	return staffArray
}

func AddStaff(name string, family string, position string) error {
	// Use a prepared statement to prevent SQL injection
	querry := "INSERT INTO staff (name, family, position) VALUES (?, ?, ?);"
	dataBase, err := connect()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	defer dataBase.Close() // Ensure the connection is closed

	stmt, err := dataBase.Prepare(querry)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close() // Ensure the statement is closed

	// Execute the statement
	_, err = stmt.Exec(name, family, position)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}

	return nil // Return nil if everything went well
}
