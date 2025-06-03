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
	db_name     = "test"
)

func main() {

	//insert_testone(fmt.Sprintf("%v Data in col11", i), fmt.Sprintf("%v Data in col12", i), fmt.Sprintf("%v Data in col13", i), fmt.Sprintf("%v Data in col14", i))
	//insert_testtwo(fmt.Sprintf("%v Data in col21", i), fmt.Sprintf("%v Data in col22", i), fmt.Sprintf("%v Data in col23", i), fmt.Sprintf("%v Data in col24", i))
	err := insert_testone("1 data", "second data", "dfdfs", "dsf")
	fmt.Println(err)
}

func Connect_to_database() (*sql.DB, error) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name))
	if err != nil {
		return database, err
	}
	return database, err
}

func insert_testone(col11 string, col12 string, col13 string, col14 string) error {
	database, err := Connect_to_database()
	if err != nil {
		fmt.Println("Here 1")
		return err
	}
	defer database.Close()
	querry, err := database.Prepare("INSERT INTO testone (col11, col12, col14, col14) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Here 2")
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(col11, col12, col13, col14)
	if err != nil {
		return err
	}
	return err
}

func insert_testtwo(col21 string, col22 string, col23 string, col24 string) {
	database, err := Connect_to_database()

	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO testtwo (col21, col22, col23, col24) VALUES (?, ?, ?, ?)")

	if err != nil {
		fmt.Printf("Error preparing the querry 2: %v\n", err)
		return // Exit the function if there's an error
	}

	defer querry.Close()

	_, err = querry.Exec(col21, col22, col23, col24)

	if err != nil {
		fmt.Printf("Error executing the querry: %v\n", err)
		return // Exit the function if there's an error
	}
}

/*func read_all_staff()  {
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

	//return staffArray

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
*/
