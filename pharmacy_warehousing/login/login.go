package login

import (
	"PharmacyWarehousing/utility"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "templates/login.html")
}
