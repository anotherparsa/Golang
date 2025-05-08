package session

import (
	"fmt"
	"net/http"
)

func Set_cookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "loggeduser",
		Value: "testuser",
		Path:  "/",
	})
}

func Read_cookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("loggeduser")

	if err != nil {
		fmt.Printf("Failed to read the cookie : %v\n", err)
	}

	fmt.Fprintf(w, "Your cookie name is %v and you cookie value is %v\n", cookie.Name, cookie.Value)
}
