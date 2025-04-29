package products

import (
	"fmt"
	"net/http"
	"strings"
)

func ShowProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is for all of the products")
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is product number %v ", strings.TrimPrefix(r.URL.Path, "/product/"))
}
