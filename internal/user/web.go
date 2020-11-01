package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vmware/vending/external/middleware"
	"net/http"
)

func (u User) CreateUser(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var user User
	if err := decoder.Decode(&user); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Save(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusCreated, user)
}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	if err := users.GetAllUsers(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusOK, users)
}

func (u *User) GetUserForUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, _ := vars["username"]
	u.Username = username
	if err := u.FetchByUsername(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusOK, u)
}

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, _ := vars["username"]
	u.Username = username
	if err := u.DeleteByUsername(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusNoContent, nil)
}