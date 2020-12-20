package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/vmware/vending/external/db"
	"github.com/vmware/vending/internal/constants"
	"github.com/vmware/vending/internal/product"
	"github.com/vmware/vending/internal/user"
	"io"
	"net/http"
)

func ok(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{ \"status\": \"OK\"}")
}

func init() {
	db.InitDB()
}

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(constants.HealthCheckUrl, ok)
	//router.Use(login.LoggingMiddleware)

	//Users related operations
	var u user.User
	router.HandleFunc(constants.GetAllUser, u.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc(constants.CreateUser, u.CreateUser).Methods(http.MethodPost)
	router.HandleFunc(constants.GetUserByUsername, u.GetUserForUsername).Methods(http.MethodGet)
	router.HandleFunc(constants.DeleteUserByUsername, u.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc(constants.Login, u.LoginHandler).Methods(http.MethodPost)

	//Users related operations
	var p product.Product
	router.HandleFunc(constants.GetAllProduct, p.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc(constants.CreateProduct, p.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc(constants.GetUserByUsername, p.GetProductByName).Methods(http.MethodGet)
	router.HandleFunc(constants.DeleteProductByName, p.DeleteProduct).Methods(http.MethodDelete)



	log.Println("Starting Webserver...")
	handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(":8080", router); err != nil && err != http.ErrServerClosed {
		log.Fatal(fmt.Sprintf("Application startup failed with error %s", err.Error()))
	}
}


