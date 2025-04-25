package router

import (
	"PharmacyWarehousing/home"
	"PharmacyWarehousing/login"
	"fmt"
	"net/http"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		home.HomePageHandler(w, r)
	case "/login":
		login.LoginPageHandler(w, r)
	case "/home":
		home.HomePageHandler(w, r)
	default:
		fmt.Fprintf(w, "This is 404 page")
	}
}
