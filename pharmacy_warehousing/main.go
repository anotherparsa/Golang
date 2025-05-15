package main

import (
	"PharmacyWarehousing/admin"
	"PharmacyWarehousing/router"
	"fmt"
	"net/http"
)

func main() {
	//creating admin user
	admin.Create_admin_user()
	//calling routiner
	http.HandleFunc("/", router.Routing)
	//serving static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//running the server
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
