package drugs

import (
	"PharmacyWarehousing/model"
	"PharmacyWarehousing/utility"
	"net/http"
)

func Create_drug_page(w http.ResponseWriter, r *http.Request) {
	utility.Render_template(w, "./drugs/templates/adddrug.html")
}

func Create_drug_processor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	model.Create_drug(r.PostForm.Get("name"), r.PostForm.Get("drugid"), r.PostForm.Get("company"), r.PostForm.Get("price"), r.PostForm.Get("stock"))
	http.Redirect(w, r, "/admin/home", http.StatusFound)
}
