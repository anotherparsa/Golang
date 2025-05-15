package model

import (
	"PharmacyWarehousing/databasetool"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Staff struct {
	Id       string
	Name     string
	Family   string
	Staffid  string
	Userid   string
	Position string
	Password string
}

type Drug struct {
	Id      string
	Name    string
	Drugid  string
	Company string
	Price   string
	Stock   string
}

// staff
func Create_staff(name string, family string, staffid string, userid string, position string, password string) error {
	database, err := databasetool.Connect_to_database()

	if err != nil {
		return err
	}

	defer database.Close()

	uuid := uuid.New().String()

	querry, err := database.Prepare("INSERT INTO staff (name, family, staffid, userid, position, password) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer querry.Close()

	_, err = querry.Exec(name, family, staffid, uuid, position, password)

	if err != nil {
		return err
	}

	return nil
}

func Read_all_staff() ([]Staff, error) {
	staff_array := []Staff{}

	database, err := databasetool.Connect_to_database()

	if err != nil {
		return staff_array, err
	}

	defer database.Close()

	rows, err := database.Query("SELECT staffid, userid, position FROM staff")

	if err != nil {
		return staff_array, err
	}

	defer rows.Close()

	staff_instance := Staff{}

	for rows.Next() {
		err = rows.Scan(&staff_instance.Staffid, &staff_instance.Userid, &staff_instance.Position)

		if err != nil {
			return staff_array, err
		}

		staff_array = append(staff_array, staff_instance)
	}

	return staff_array, nil
}

// drugs
func Create_drug(name string, drugid string, company string, price string, stock string) {
	database, err := databasetool.Connect_to_database()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO drug (name, drugid, company, prices, stock) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Printf("Failed to prepare the querry : %v\n", err)
	}

	defer querry.Close()

	_, err = querry.Exec(name, drugid, company, price, stock)

	if err != nil {
		fmt.Printf("Failed to execute the querry : %v\n", err)
	}
}

func Read_all_drug() []Drug {
	database, err := databasetool.Connect_to_database()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	rows, err := database.Query("SELECT * FROM drug")

	if err != nil {
		fmt.Printf("Failed to querry the database : %v\n", err)
	}

	defer rows.Close()

	drug_instance := Drug{}
	drug_array := []Drug{}

	for rows.Next() {
		err = rows.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)

		if err != nil {
			fmt.Printf("Failed to scan rows : %v\n", err)
		}

		drug_array = append(drug_array, drug_instance)
	}

	return drug_array
}

// session
func Create_session(userid string, sessionid string) {
	database, err := databasetool.Connect_to_database()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO session (userid, sessionid) VALUES (?, ?)")

	if err != nil {
		fmt.Printf("Failed to prepare the querry : %v\n", err)
	}
	defer querry.Close()

	_, err = querry.Exec(userid, sessionid)

	if err != nil {
		fmt.Printf("Failed to execute the querry : %v\n", err)
	}
}

func Check_if_cookie_exists(r *http.Request, cookiename string) (string bool) {
	_, err := r.Cookie(cookiename)
	return err == nil

}
