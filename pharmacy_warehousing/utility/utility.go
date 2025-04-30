package utility

import (
	"PharmacyWarehousing/model"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var DataBase *sql.DB

const (
	username = "testuser"
	password = "rrc3498urc38r9j999m8j"
	hostname = "127.0.0.1"
	port     = "3306"
	database = "pharmacywarehouse"
)

// opening a connection
func connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	return sql.Open("mysql", dsn)
}

func SelectAllStaff() []model.Staff {
	db, err := connect()
	if err != nil {
		fmt.Println("Error connecting to database: 1", err)
		return nil // Return an empty slice on error
	}

	rows, err := db.Query("SELECT * FROM staff")
	if err != nil {
		fmt.Println("Error querying database: 1", err)
		return nil // Return an empty slice on error
	}

	staffs := []model.Staff{}

	for rows.Next() {
		staff := model.Staff{}
		err = rows.Scan(&staff.Id, &staff.Name, &staff.Family, &staff.Position)
		if err != nil {
			fmt.Println("Error scanning row: 1", err)
			continue // Skip this row but continue with the next
		}
		staffs = append(staffs, staff)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		fmt.Println("Error during row iteration:", err)
	}

	return staffs
}

func Print_staff() {
	staff := SelectAllStaff()
	for i, v := range staff {
		fmt.Printf("%v and %v", i, v)
	}
}

func TemplateRendering(w http.ResponseWriter, path string) {
	template, err := template.ParseFiles(path)

	if err != nil {
		log.Print(err)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = template.Execute(w, nil)

	if err != nil {
		log.Print(err)
		return
	}
}
