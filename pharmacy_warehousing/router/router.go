package router

import (
	"PharmacyWarehousing/admin"
	"PharmacyWarehousing/drugs"
	"PharmacyWarehousing/login"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/staff"
	"fmt"
	"net/http"
	"strings"
)

func Routing(w http.ResponseWriter, r *http.Request) {
	//routes are divided mainly into 3 divisions
	//admin, staff, drug
	url_path := r.URL.Path
	if url_path == "/" {
		staff.Staff_home_page(w, r)
	} else if strings.HasPrefix(url_path, "/staff") {
		if url_path == "/staff/home" || url_path == "/staff" {
			staff.Staff_home_page(w, r)
		} else if url_path == "/staff/login" {
			login.Login_page(w, r)
		} else if url_path == "/staff/loginprocessor" {
			login.Login_processor(w, r)
		} else if url_path == "/staff/logout" {
			session.User_logout(w, r)
		} else {
			fmt.Fprintf(w, "404 page not found")
		}
	} else if strings.HasPrefix(url_path, "/admin") {
		if url_path == "/admin/home" || url_path == "/admin" {
			staff.Staff_home_page(w, r)
		} else if url_path == "/admin/addstaff" {
			admin.Admin_add_staff_page(w, r)
		} else if url_path == "/admin/addstaffprocessor" {
			admin.Admin_add_staff_processor(w, r)
		} else {
			fmt.Fprintf(w, "404 page not found")
		}
	} else if strings.HasPrefix(url_path, "/drug") {
		if url_path == "/drug/alldrugs" || url_path == "/drug" {
			//
		} else if url_path == "/drug/adddrug" {
			drugs.Create_drug_page(w, r)
		} else if url_path == "/drug/adddrugprocessor" {
			drugs.Create_drug_processor(w, r)
		} else {
			fmt.Fprintf(w, "404 page nout fount")
		}
	} else if strings.HasPrefix(url_path, "/error") {
		fmt.Fprintf(w, "This is error page")
	}
}
