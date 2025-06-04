package drugs

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"database/sql"
	"fmt"
	"net/http"
)

type DataToSend struct {
	User interface{}
	Drug interface{}
}

// handler of "/drug/adddrug" {
func Create_drug_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"storekeeper"})
	if err == nil {
		err = utility.Render_template(w, "./drugs/templates/adddrug.html", nil)
		if err != nil {
			fmt.Printf("Error Drugs 2: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
	} else {
		fmt.Printf("Error Drugs 1: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
}

// handler of "/drug/adddrugprocessor"
func Create_drug_processor(w http.ResponseWriter, r *http.Request) {
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
			} else {
				fmt.Printf("Error Drugs 5: %v\n", err)
				http.Redirect(w, r, "/error", http.StatusFound)
			}
		} else {
			fmt.Printf("Error Drugs 4: %v\n", err)
			http.Redirect(w, r, "/error", http.StatusFound)
		}
	} else {
		fmt.Printf("Error Drugs 3: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
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
		} else {
			return err
		}
	} else {
		return err
	}
}

// handler of "/drug/alldrugs" and "/drug"
func All_drugs_page(w http.ResponseWriter, r *http.Request) {
	User, err := session.Is_user_authorized(r, []string{"recipient", "storekeeper"})
	if err == nil {
		Drugs_array, err := All_drugs()
		if err == nil {
			data := DataToSend{Drug: Drugs_array, User: User}
			err = utility.Render_template(w, "./drugs/templates/alldrugs.html", data)
			if err != nil {
				fmt.Printf("Error Drugs 9: %v\n", err)
				http.Redirect(w, r, "/error", http.StatusFound)
			}
		}
	}
}

func All_drugs() ([]model.Drug, error) {
	drug_instance := model.Drug{}
	drug_array := []model.Drug{}
	database, err := databasetool.Connect_to_database()
	if err == nil {
		defer database.Close()
		rows, err := database.Query("SELECT * FROM drug")
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
				if err == nil {
					drug_array = append(drug_array, drug_instance)
				} else {
					fmt.Printf("Error Drugs 14: %v\n", err)
					continue
				}
			}
		} else {
			fmt.Printf("Error Drugs 13: %v\n", err)
			return drug_array, err
		}
	} else {
		fmt.Printf("Error Drugs 12: %v\n", err)
		return drug_array, err
	}
	return drug_array, err
}

// handler of //
func Search_result_page(w http.ResponseWriter, r *http.Request) {
	user, err := session.Is_user_authorized(r, []string{"recipient", "storekeeper"})
	if err == nil {
		drug_name := r.URL.Query().Get("drugname")
		fmt.Println(drug_name)
		result_drug, err := Find_drug(drug_name)
		if err == nil {
			fmt.Println("Hey")
			data := DataToSend{User: user, Drug: result_drug}
			fmt.Println(data)
			utility.Render_template(w, "./staff/templates/searchresult.html", data)
		}
	}
}

func Find_drug(drugname string) (model.Drug, error) {
	// Create a zero-value Drug instance
	drug_instance := model.Drug{}

	// Connect to the database
	database, err := databasetool.Connect_to_database()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return drug_instance, err
	}
	defer database.Close() // Ensure the database is closed when done

	// Query for the drug
	row := database.QueryRow("SELECT * FROM drug WHERE name=?", drugname)

	// Scan the result into drug_instance
	err = row.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)

	// Check for errors during scanning
	if err != nil {
		if err == sql.ErrNoRows {
			// Drug not found
			return drug_instance, fmt.Errorf("drug not found: %s", drugname)
		}
		// Log other types of errors
		fmt.Printf("Error scanning drug: %v\n", err)
		return drug_instance, err
	}

	// Return the found drug instance
	return drug_instance, nil
}
