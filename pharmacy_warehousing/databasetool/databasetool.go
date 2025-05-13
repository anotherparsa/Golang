package databasetool

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Database_connection_failure = errors.New("Failed to connect to the database")
var Querry_prepration_failure = errors.New("Failed to prepare the queery")
var Querry_execution_failure = errors.New("Failed to execute the querry")
var Scanning_rows_failure = errors.New("Failed to scan the rows")

const (
	db_user     = "testuser"
	db_password = "rrc3498urc38r9j999m8j"
	db_address  = "127.0.0.1"
	db_port     = "3306"
	db_name     = "pharmacywarehouse"
)

func Connect_to_database() (*sql.DB, error) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name))

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	return database, err
}
