package staff

import (
	"PharmacyWarehousing/databasetool"
	"fmt"
)

func Create_drug(name string, drugid string, company string, price string, stock string) {
	database, err := databasetool.Connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	querry, err := database.Prepare("INSERT INTO drug (name, drugid, company, price, stock) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Printf("Failed to prepare the querry : %v\n", err)
	}

	defer querry.Close()

	_, err = querry.Exec(name, drugid, company, price, stock)

	if err != nil {
		fmt.Printf("Failed to execute the querry : %v\n", err)
	}
}
