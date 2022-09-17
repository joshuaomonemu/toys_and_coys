package routes

import (
	"app/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// Starting route points
var r = mux.NewRouter()

func Routes() {
	//r := mux.NewRouter()

	// fs := http.FileServer(http.Dir("./view/assets/"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// http.Handle("/", r)
	// r.HandleFunc("/", handlers.Index)

	// Api for User
	//r.HandleFunc("/api/user/get", api.GetUser)

	// r.HandleFunc("/api/user/update", api.UpdateUser)
	// r.HandleFunc("/api/user/delete", api.DeleteUser)
	// r.HandleFunc("/api/user/addfriend", api.Newfriend)

	// // Api for messages
	// r.HandleFunc("/api/message/create", api.CreateMessage)
	// r.HandleFunc("/api/chat/get/", api.CreateMessage)

	// r.HandleFunc("/ws", util.WsEndpoint)

	//Starting Server and running on port 2020

	// Api for User
	//r.HandleFunc("/api/user/get{id}", api.GetUser).Methods("GET")
	port :=  os.Getenv("PORT")
	r.HandleFunc("/api/student/create", api.CreateStudent).Methods("POST")
	r.HandleFunc("/api/student/delete/{id}", api.DeleteStudent).Methods("GET")
	r.HandleFunc("/api/student/get/{id}", api.ReadStudent).Methods("GET")
	r.HandleFunc("/api/student/update/{id}", api.UpdateStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
