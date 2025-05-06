package drugs

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

func Create_drug_page_handler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "./drugs/templates/adddrug.html")
}

func Create_drug_processor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Create_drug(r.PostForm.Get("name"), r.PostForm.Get("drugid"), r.PostForm.Get("company"), r.PostForm.Get("price"), r.PostForm.Get("stock"))
	http.Redirect(w, r, "/admin/home", 302)
}

func Create_drug(name string, drugid string, company string, price string, stock string) {
	database, err := databasetool.Connect()

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

func Read_all_drug() []model.Drug {
	database, err := databasetool.Connect()

	if err != nil {
		fmt.Printf("Failed to connect to the database : %v\n", err)
	}

	defer database.Close()

	rows, err := database.Query("SELECT * FROM drug")

	if err != nil {
		fmt.Printf("Failed to querry the database : %v\n", err)
	}

	defer rows.Close()

	drug_instance := model.Drug{}
	drug_array := []model.Drug{}

	for rows.Next() {
		err = rows.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)

		if err != nil {
			fmt.Printf("Failed to scan rows : %v\n", err)
		}

		drug_array = append(drug_array, drug_instance)
	}

	return drug_array
}
