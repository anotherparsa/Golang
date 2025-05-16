package drugs

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

func Create_drug_page(w http.ResponseWriter, r *http.Request) {
	err := session.Is_user_authorized(r, "storekeeper")

	if err != nil {
		fmt.Printf("Error : %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	err = utility.Render_template(w, "./drugs/template/adddrug.html")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}

}

func Create_drug_processor(w http.ResponseWriter, r *http.Request) {
	err := session.Is_user_authorized(r, "storekeeper")
	//user is not authorized
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	err = r.ParseForm()
	//error in parsing the form
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	drug_name := r.PostForm.Get("name")
	drug_id := r.PostForm.Get("drugid")
	drug_company := r.PostForm.Get("company")
	drug_price := r.PostForm.Get("price")
	drug_stock := r.PostForm.Get("stock")
	err = Create_drug(drug_name, drug_id, drug_company, drug_price, drug_stock)
	//error in creating drug record in database
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	http.Redirect(w, r, "/admin/home", http.StatusFound)

}

// drugs
func Create_drug(name string, drugid string, company string, price string, stock string) error {
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return err
	}
	defer database.Close()
	querry, err := database.Prepare("INSERT INTO drug (name, drugid, company, prices, stock) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(name, drugid, company, price, stock)
	if err != nil {
		return err
	}
	return err
}

func Read_all_drug() ([]model.Drug, error) {
	drug_array := []model.Drug{}
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return drug_array, err
	}
	defer database.Close()
	rows, err := database.Query("SELECT * FROM drug")
	if err != nil {
		return drug_array, err
	}
	defer rows.Close()
	drug_instance := model.Drug{}
	for rows.Next() {
		err = rows.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
		if err != nil {
			return drug_array, err
		}
		drug_array = append(drug_array, drug_instance)
	}
	return drug_array, err
}
