package main

import (
	"PharmacyWarehousing/router"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", router.RoutingHandler)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
