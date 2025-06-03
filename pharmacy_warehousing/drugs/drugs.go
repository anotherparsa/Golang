package drugs

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

type dataToSend struct {
	Drugs []model.Drug
}

func Create_drug_page(w http.ResponseWriter, r *http.Request) {
	//checking if the user is authorized, which means both session and their position
	err := session.Is_user_authorized(r, []string{"storekeeper"})
	//error in authorizing the user
	if err != nil {
		fmt.Printf("Error 10: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	err = utility.Render_template(w, "./drugs/templates/adddrug.html", nil)
	//error in rendering the template
	if err != nil {
		fmt.Printf("Error 11: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}

}

func Create_drug_processor(w http.ResponseWriter, r *http.Request) {
	//checking if the user is authorized, which means both session and their position
	err := session.Is_user_authorized(r, []string{"storekeeper"})
	//error in authorizing the user
	if err != nil {
		fmt.Printf("Error 13: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	//parsing the form
	err = r.ParseForm()
	//error in parsing the form
	if err != nil {
		fmt.Printf("Error 14: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	//getting data from form
	drug_name := r.PostForm.Get("drugname")
	drug_id := r.PostForm.Get("drugid")
	drug_company := r.PostForm.Get("company")
	drug_price := r.PostForm.Get("price")
	drug_stock := r.PostForm.Get("stock")
	//creating drug record in the database
	err = Create_drug_record(drug_name, drug_id, drug_company, drug_price, drug_stock)
	//error in creating drug record in database
	if err != nil {
		fmt.Printf("Error 15: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	//redircting to the storekeeper home
	http.Redirect(w, r, "/staff/home", http.StatusFound)

}

func Create_drug_record(drugname string, drugid string, company string, price string, stock string) error {
	//connecting to the database
	database, err := databasetool.Connect_to_database()
	//error in connecting to the database
	if err != nil {
		return err
	}
	defer database.Close()
	//preparing the querry
	querry, err := database.Prepare("INSERT INTO drug (name, drugid, company, price, stock) VALUES (?, ?, ?, ?, ?)")
	//error in preparing the querry
	if err != nil {
		return err
	}
	defer querry.Close()
	//executing the querry
	_, err = querry.Exec(drugname, drugid, company, price, stock)
	//error in executing the querry
	if err != nil {
		return err
	}
	return err
}
func All_drugs() ([]model.Drug, error) {
	drug_instance := model.Drug{}
	drug_array := []model.Drug{}
	database, err := databasetool.Connect_to_database()
	if err != nil {
		fmt.Printf("Error 41 %v\n", err)
		return drug_array, err
	}
	defer database.Close()
	rows, err := database.Query("SELECT * FROM drug")
	if err != nil {
		fmt.Printf("Error 42 %v\n", err)
		return drug_array, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
		if err != nil {
			fmt.Printf("failed to scan the row %v\n", err)
			continue
		}
		drug_array = append(drug_array, drug_instance)
	}

	if rows.Err() != nil {
		return drug_array, err
	}
	return drug_array, err
}

func All_drugs_page(w http.ResponseWriter, r *http.Request) {
	err := session.Is_user_authorized(r, []string{"storekeeper, recipient"})
	if err != nil {
		fmt.Printf("Error 40 %v\n", err)
		http.Redirect(w, r, "/staff/login", http.StatusFound)
	}
	Drugs_array, err := All_drugs()
	data := dataToSend{Drugs: Drugs_array}
	if err != nil {
		fmt.Printf("Error 54 %v\n", err)
	}
	fmt.Println(Drugs_array)
	err = utility.Render_template(w, "./drugs/templates/alldrugs.html", data)
	if err != nil {
		fmt.Printf("Error 55 %v\n", err)
	}
}

/*
	func All_drugs_page(w http.ResponseWriter, r *http.Request) {
		err := session.Is_user_authorized(r, []string{"storekeeper"})
		if err != nil {
			fmt.Printf("Error 40%v\n", err)
			http.Redirect(w, r, "/staff/login", http.StatusFound)
		}
		//connecting to the database
		database, err := databasetool.Connect_to_database()
		//failed to connect to the database
		if err != nil {
			fmt.Printf("Error 41 %v\n", err)
		}
		defer database.Close()
		//preparing the querry
		queery, err := database.Prepare("SELECT * FROM drug")
		//error in preparing the querry
		if err != nil {
			fmt.Printf("Error 42 %v\n", err)
		}
		defer queery.Close()
		_, err = database.Exec(queery)

		drug_array := []model.Drug{}
		err = utility.Render_template(w, "./drugs/templates/alldrugs.html", nil)

}
*/
func Search_for_drug(drugname string) (model.Drug, error) {
	found_drug := model.Drug{}
	//connecting to the database
	database, err := databasetool.Connect_to_database()
	//failed to connect to the database
	if err != nil {
		return found_drug, err
	}
	defer database.Close()
	querry := "SELECT FROM drug WHERE name=?"
	row := database.QueryRow(querry, drugname)
	//scanning the row
	err = row.Scan(&found_drug.Id, &found_drug.Name, &found_drug.Drugid, &found_drug.Company, &found_drug.Price, &found_drug.Stock)
	//error in scanning the row
	if err != nil {
		return found_drug, err
	}
	return found_drug, err

}
