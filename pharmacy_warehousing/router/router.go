package router

import (
	"PharmacyWarehousing/admin"
	"PharmacyWarehousing/home"
	"PharmacyWarehousing/login"
	"PharmacyWarehousing/products"
	"net/http"
	"strings"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	UrlPath := r.URL.Path

	if UrlPath == "/home" || UrlPath == "/" {
		home.HomePageHandler(w, r)
	} else if strings.HasPrefix(UrlPath, "/product") {
		if UrlPath == "/products" {
			products.ShowProducts(w, r)
		} else {
			products.ShowProduct(w, r)
		}
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
	}
}
