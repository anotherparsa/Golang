package databasetool

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

func Connect() (*sql.DB, error) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name))

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	return database, err
}
