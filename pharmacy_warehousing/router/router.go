package router

import (
	"PharmacyWarehousing/admin"
	"PharmacyWarehousing/drugs"
	"PharmacyWarehousing/login"
	"PharmacyWarehousing/session"
	"PharmacyWarehousing/staff"
	"net/http"
	"strings"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	UrlPath := r.URL.Path

	if UrlPath == "/home" || UrlPath == "/" {
		staff.Staff_home_page_handler(w, r)
	} else if strings.HasPrefix(UrlPath, "/login") {
		if UrlPath == "/login" {
			login.LoginPageHandler(w, r)
		} else {
			login.LoginHandler(w, r)
		}
	} else if strings.HasPrefix(UrlPath, "/admin") {
		if UrlPath == "/admin" || UrlPath == "/admin/home" {
			admin.Admin_home_page_handler(w, r)
		} else if UrlPath == "/admin/addstaff" {
			admin.Admin_add_staff_page_handler(w, r)
		} else if UrlPath == "/admin/addstaffprocess" {
			admin.Admin_add_staff_processor(w, r)
		}
	} else if strings.HasPrefix(UrlPath, "/drug") {
		if UrlPath == "/drug/adddrug" {
			drugs.Create_drug_page_handler(w, r)
		} else if UrlPath == "/drug/adddrugprocessor" {
			drugs.Create_drug_processor(w, r)
		}
	} else if strings.HasPrefix(UrlPath, "/cookie") {
		if UrlPath == "/cookie/set" {
			session.Set_cookie(w, r)
		} else if UrlPath == "/cookie/read" {
			session.Read_cookie(w, r)
		} else if UrlPath == "/cookie/delete" {
			session.Delete_cookie(w, r)
		}
	}
}
