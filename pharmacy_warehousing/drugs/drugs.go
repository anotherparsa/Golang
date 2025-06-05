package drugs

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

type DataToSend struct {
	User interface{}
	Drug interface{}
}

// handler of "/drug/adddrug" {
func Create_drug_page(w http.ResponseWriter, r *http.Request) error {
	_, err := session.Is_user_authorized(r, []string{"storekeeper"})
	if err == nil {
		err = utility.Render_template(w, "./drugs/templates/adddrug.html", nil)
		return err
	}
	return err
}

// handler of "/drug/adddrugprocessor"
func Create_drug_processor(w http.ResponseWriter, r *http.Request) error {
	_, err := session.Is_user_authorized(r, []string{"storekeeper"})
	if err == nil {
		err = r.ParseForm()
		if err == nil {
			drug_name := r.PostForm.Get("drugname")
			drug_id := r.PostForm.Get("drugid")
			drug_company := r.PostForm.Get("company")
			drug_price := r.PostForm.Get("price")
			drug_stock := r.PostForm.Get("stock")
			err = Create_drug_record(drug_name, drug_id, drug_company, drug_price, drug_stock)
			if err == nil {
				http.Redirect(w, r, "/staff/home", http.StatusFound)
				return nil
			}
		}
	}
	return err
}

func Create_drug_record(drugname string, drugid string, company string, price string, stock string) error {
	database, err := databasetool.Connect_to_database()
	if err == nil {
		defer database.Close()
		querry, err := database.Prepare("INSERT INTO drug (name, drugid, company, price, stock) VALUES (?, ?, ?, ?, ?)")
		if err == nil {
			defer querry.Close()
			_, err = querry.Exec(drugname, drugid, company, price, stock)
			return err
		}
	}
	return err
}

// handler of "/drug/alldrugs" and "/drug"
func All_drugs_page(w http.ResponseWriter, r *http.Request) error {
	User, err := session.Is_user_authorized(r, []string{"recipient", "storekeeper"})
	if err == nil {
		Drugs_array, err := All_drugs()
		if err == nil {
			data := DataToSend{Drug: Drugs_array, User: User}
			err = utility.Render_template(w, "./drugs/templates/alldrugs.html", data)
			return err
		}
	}
	return err

}

func All_drugs() ([]model.Drug, error) {
	drug_array := []model.Drug{}
	database, err := databasetool.Connect_to_database()
	if err == nil {
		defer database.Close()
		rows, err := database.Query("SELECT * FROM drug")
		if err == nil {
			defer rows.Close()
			drug_instance := model.Drug{}
			for rows.Next() {
				err := rows.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
				if err == nil {
					drug_array = append(drug_array, drug_instance)
				}
			}
			return drug_array, err
		} else {
			//failed to querry the database
			fmt.Printf("Error Drugs 13: %v\n", err)
			return drug_array, err
		}
	} else {
		//failed to connect to the database
		fmt.Printf("Error Drugs 12: %v\n", err)
		return drug_array, err
	}

}

// handler of "/drug/searchdrug"
func Search_result_page(w http.ResponseWriter, r *http.Request) error {
	user, err := session.Is_user_authorized(r, []string{"recipient", "storekeeper"})
	if err == nil {
		drug_name := r.URL.Query().Get("drugname")
		fmt.Println(drug_name)
		result_drug, err := Find_drug(drug_name)
		if err == nil {
			data := DataToSend{User: user, Drug: result_drug}
			utility.Render_template(w, "./staff/templates/searchresult.html", data)
		} else {
			//failed to find the drug
			fmt.Printf("Error Drugs 9: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
	} else {
		//user is not authorized
		fmt.Printf("Error Drugs 9: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	return nil
}

func Find_drug(drugname string) (model.Drug, error) {
	drug_instance := model.Drug{}
	database, err := databasetool.Connect_to_database()
	if err == nil {
		defer database.Close()
		row := database.QueryRow("SELECT * FROM drug WHERE name=?", drugname)
		err = row.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
		return drug_instance, err
	} else {
		return drug_instance, err
	}

}
