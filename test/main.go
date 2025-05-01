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
	Position string
}

func main() {
	insert_data("testn1", "testf1", "testp1")
	insert_data("testn2", "testf2", "testp2")
	insert_data("testn3", "testf3", "testp3")
	insert_data("testn4", "testf4", "testp4")
}

func connect() *sql.DB {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name))

	if err != nil {
		fmt.Printf("Error connecting database %v\n", err)
	}

	return database
}

func insert_data(name string, family string, position string) {
	database := connect()

	querry, _ := database.Prepare("INSERT INTO `staff` (name, family, position) VALUES (?, ?, ?)")

	_, err := querry.Exec(name, family, position)

	if err != nil {
		fmt.Printf("Error executing the query %v\n", err)
	}
}
