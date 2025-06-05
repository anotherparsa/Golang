package databasetool

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "testuser"
	db_password = "jwegoqj398329rhwegojldgjsldgkj"
	db_address  = "127.0.0.1"
	db_port     = "3306"
	db_name     = "pharmacywarehouse"
)

func Connect_to_database() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name))
}
