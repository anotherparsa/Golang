package home

import (
	"PharmacyWarehousing/utility"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "./staff/templates/home.html")
}
