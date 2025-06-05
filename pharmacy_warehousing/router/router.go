package router

import (
	"PharmacyWarehousing/admin"
	"PharmacyWarehousing/drugs"
	"PharmacyWarehousing/login"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/staff"
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
	"strings"
)

func Routing(w http.ResponseWriter, r *http.Request) {
	var err error
	url_path := r.URL.Path
	if url_path == "/" {
		if err = staff.Staff_home_page(w, r); err != nil {
			utility.Error_handler(w, err.Error(), "staff")
		}
	} else if strings.HasPrefix(url_path, "/staff") {
		if url_path == "/staff/home" || url_path == "/staff" {
			if err = staff.Staff_home_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "staff")
			}
		} else if url_path == "/staff/login" {
			if err = login.Login_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "staff")
			}
		} else if url_path == "/staff/loginprocessor" {
			if err = login.Login_processor(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "staff")
			}
		} else if url_path == "/staff/logout" {
			if err = session.User_logout(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "staff")
			}
		} else {
			fmt.Fprintf(w, "404 page not found")
		}
	} else if strings.HasPrefix(url_path, "/admin") {
		if url_path == "/admin/home" || url_path == "/admin" {
			if err = staff.Staff_home_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "admin")
			}
		} else if url_path == "/admin/addstaff" {
			if err = admin.Admin_add_staff_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "admin")
			}
		} else if url_path == "/admin/addstaffprocessor" {
			if err = admin.Admin_add_staff_processor(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "admin")
			}
		} else if url_path == "/admin/allstaff" {
			if err = staff.All_staff_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "admin")
			}
		} else {
			fmt.Fprintf(w, "404 page not found")
		}
	} else if strings.HasPrefix(url_path, "/drug") {
		if url_path == "/drug/alldrugs" || url_path == "/drug" {
			if err = drugs.All_drugs_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "drug")
			}
		} else if url_path == "/drug/adddrug" {
			if err = drugs.Create_drug_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "drug")
			}
		} else if url_path == "/drug/adddrugprocessor" {
			if err = drugs.Create_drug_processor(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "drug")
			}
		} else if url_path == "/drug/searchdrug" {
			if err = drugs.Search_result_page(w, r); err != nil {
				utility.Error_handler(w, err.Error(), "drug")
			}
		} else {
			fmt.Fprintf(w, "404 page nout fount")
		}
	}
}
