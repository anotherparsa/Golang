package router

import (
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
	}
}
