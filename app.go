package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/sreekrishnaraja/golang-sample-api/config"
	"github.com/sreekrishnaraja/golang-sample-api/controllers"
	. "github.com/sreekrishnaraja/golang-sample-api/dao"
)

var config = Config{}
var dao = UsersDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.ListAllUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", controllers.FindUser).Methods("GET")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
