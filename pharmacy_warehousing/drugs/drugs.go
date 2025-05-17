package drugs

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

func Create_drug_page(w http.ResponseWriter, r *http.Request) {
	//checking if the user is authorized, which means both session and their position
	err := session.Is_user_authorized(r, "storekeeper")
	//error in authorizing the user
	if err != nil {
		fmt.Printf("Error 10: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	err = utility.Render_template(w, "./drugs/templates/adddrug.html")
	//error in rendering the template
	if err != nil {
		fmt.Printf("Error 11: %v\n", err)
		http.Redirect(w, r, "/error", http.StatusFound)
	}
}

func Create_drug_processor(w http.ResponseWriter, r *http.Request) {
	//checking if the user is authorized, which means both session and their position
	err := session.Is_user_authorized(r, "storekeeper")
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
	drug_name := r.PostForm.Get("name")
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
	querry, err := database.Prepare("INSERT INTO drug (name, drugid, company, prices, stock) VALUES (?, ?, ?, ?, ?)")
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
