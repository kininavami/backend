package session

import (
	"github.com/vmware/vending/external/middleware"
	"net/http"
)

// login authenticates the user
func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = authenticate(w, r); err != nil {
		middleware.RespondError(w, http.StatusUnauthorized, err)
	}

	//user := &User{
	//	Username:      username,
	//	Authenticated: true,
	//}
	//
	//session.Values["user"] = user
	//
	err = session.Save(r, w)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	http.Redirect(w, r, "/secret", http.StatusFound)
}