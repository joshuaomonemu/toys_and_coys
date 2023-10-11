package routes

import (
	"app/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Starting route points
var r = mux.NewRouter()
var port string

func Routes() {
	//Starting Server and running on port 2020

	// Api for User
	//r.HandleFunc("/controller/user/get{id}", controller.GetUser).Methods("GET")
	port = "2020"

	//Endpoints and Route points for users
	r.HandleFunc("/user/create", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/delete/{id}", controller.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/get/{id}", controller.ReadUser).Methods("GET")
	r.HandleFunc("/user/update/{id}", controller.UpdateUser).Methods("PATCH")

	//Route points and endpoints for events
	r.HandleFunc("/events/create", controller.CreateEvent).Methods("POST")
	r.HandleFunc("/events/delete/{id}", controller.DeleteEvent).Methods("DELETE")
	r.HandleFunc("/events/get/{id}", controller.ReadEvent).Methods("GET")

	r.HandleFunc("/events/comment/{id}", controller.CreateComment).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
