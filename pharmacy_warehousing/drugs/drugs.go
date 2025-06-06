package drugs

import (
	"PharmacyWarehousing/databasetool"
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
	"strings"
)

type DataToSend struct {
	User interface{}
	Drug interface{}
}

// handler of "/drug/adddrug" {
func Create_drug_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	err = utility.Render_template(w, "./drugs/templates/adddrug.html", nil)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
}

// handler of "/drug/adddrugprocessor"
func Create_drug_processor(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	err = r.ParseForm()
	if err != nil {
		utility.Error_handler(w, err.Error())
		return

	}
	drug_name := r.PostForm.Get("drugname")
	drug_id := r.PostForm.Get("drugid")
	drug_company := r.PostForm.Get("company")
	drug_price := r.PostForm.Get("price")
	drug_stock := r.PostForm.Get("stock")
	err = Create_drug_record(drug_name, drug_id, drug_company, drug_price, drug_stock)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return

	}
	http.Redirect(w, r, "/staff/home", http.StatusFound)

}

func Edit_drug_page(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	drug, err := Get_drug_by("id", strings.TrimPrefix(r.URL.Path, "/staff/editdrug/"))
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	data := DataToSend{Drug: drug}
	err = utility.Render_template(w, "./drugs/templates/editdrug.html", data)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
}

func Edit_drug_processor(w http.ResponseWriter, r *http.Request) {
	_, err := session.Is_user_authorized(r, []string{"storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	err = r.ParseForm()
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	name := r.PostForm.Get("drugname")
	drugid := r.PostForm.Get("drugid")
	company := r.PostForm.Get("company")
	price := r.PostForm.Get("price")
	stock := r.PostForm.Get("stock")
	id := r.PostForm.Get("id")
	err = Edit_staff_record(id, name, drugid, company, price, stock)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	http.Redirect(w, r, "/drug/alldrugs", http.StatusFound)
}

func Get_drug_by(condition string, condition_value string) (model.Drug, error) {
	drug_instance := model.Drug{}
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return drug_instance, err
	}
	defer database.Close()
	row := database.QueryRow(fmt.Sprintf("SELECT * FROM drug WHERE %v=?", condition), condition_value)
	err = row.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
	if err != nil {
		return drug_instance, err
	}
	return drug_instance, nil
}

func Edit_staff_record(id string, name string, drugid string, company string, price string, stock string) error {
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return err
	}
	defer database.Close()
	querry, err := database.Prepare("UPDATE drug SET name=?, drugid=?, company=?, price=?, stock=?")
	if err != nil {
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(name, drugid, company, price, stock)
	if err != nil {
		return err
	}
	return nil
}

func Create_drug_record(drugname string, drugid string, company string, price string, stock string) error {
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return err
	}
	defer database.Close()
	querry, err := database.Prepare("INSERT INTO drug (name, drugid, company, price, stock) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer querry.Close()
	_, err = querry.Exec(drugname, drugid, company, price, stock)
	if err != nil {
		return err
	}
	return nil
}

// handler of "/drug/alldrugs" and "/drug"
func All_drugs_page(w http.ResponseWriter, r *http.Request) {
	User, err := session.Is_user_authorized(r, []string{"recipient", "storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	Drugs_array, err := All_drugs()
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	data := DataToSend{Drug: Drugs_array, User: User}
	err = utility.Render_template(w, "./drugs/templates/alldrugs.html", data)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
}

func All_drugs() ([]model.Drug, error) {
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
		err := rows.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
		if err == nil {
			drug_array = append(drug_array, drug_instance)
		}
	}
	return drug_array, rows.Err()
}

// handler of "/drug/searchdrug"
func Search_result_page(w http.ResponseWriter, r *http.Request) {
	user, err := session.Is_user_authorized(r, []string{"recipient", "storekeeper"})
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	drug_name := r.URL.Query().Get("drugname")
	fmt.Println(drug_name)
	result_drug, err := Find_drug(drug_name)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
	data := DataToSend{User: user, Drug: result_drug}
	err = utility.Render_template(w, "./staff/templates/searchresult.html", data)
	if err != nil {
		utility.Error_handler(w, err.Error())
		return
	}
}

func Find_drug(drugname string) (model.Drug, error) {
	drug_instance := model.Drug{}
	database, err := databasetool.Connect_to_database()
	if err != nil {
		return drug_instance, err
	}
	defer database.Close()
	row := database.QueryRow("SELECT * FROM drug WHERE name=?", drugname)
	err = row.Scan(&drug_instance.Id, &drug_instance.Name, &drug_instance.Drugid, &drug_instance.Company, &drug_instance.Price, &drug_instance.Stock)
	if err != nil {
		return drug_instance, err
	}
	return drug_instance, row.Err()
}
