package main

import (
	"PharmacyWarehousing/router"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", router.Routing)

	//serving static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
