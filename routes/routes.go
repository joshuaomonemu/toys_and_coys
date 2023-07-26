package routes

import (
	"app/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Starting route points
var r = mux.NewRouter()

func Routes() {
	//Starting Server and running on port 2020

	// Api for Users
	//r.HandleFunc("/api/user/get{id}", api.GetUser).Methods("GET")
	port := "2020"

	r.HandleFunc("/api/user/create", api.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/delete/{id}", api.DeleteUser).Methods("GET")
	r.HandleFunc("/api/user/get/{id}", api.ReadUser).Methods("GET")
	//r.HandleFunc("/api/user/login", api.LoginUser).Methods("POST")
	//r.HandleFunc("/api/user/update/{id}", api.UpdateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
