package user

import (
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/vmware/vending/external/middleware"
	"io/ioutil"
	"net/http"
)

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

var LoginMap = make(map[string]bool)

func (u User) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Content-Type", "*")

	
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Content-Type", "*")

	
	var users Users
	if err := users.GetAllUsers(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusOK, users)
}

func (u *User) GetUserForUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Content-Type", "*")

	
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Content-Type", "*")

	
	vars := mux.Vars(r)
	username, _ := vars["username"]
	u.Username = username
	if err := u.DeleteByUsername(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusNoContent, nil)
}

func (u *User) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Content-Type", "*")

	
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res ResponseResult
	if err != nil {
		//log.Fatal(err)
	}
	isAuthenticated := user.Authenticate()

	if !isAuthenticated {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	user.Token = tokenString
	user.Password = ""
	LoginMap[user.Username] = true

	err = json.NewEncoder(w).Encode(user)

}
