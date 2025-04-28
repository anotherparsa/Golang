package login

import (
	"PharmacyWarehousing/utility"
	"fmt"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	utility.TemplateRendering(w, "templates/login.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "%v %v", r.PostForm.Get("username"), r.PostForm.Get("password"))
}
