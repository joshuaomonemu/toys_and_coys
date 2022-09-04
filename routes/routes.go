package routes

import (
	"net/http"
)

//Starting route points
func Routes() {
	// r := mux.NewRouter()

	// fs := http.FileServer(http.Dir("./view/assets/"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// http.Handle("/", r)
	// r.HandleFunc("/", handlers.Index)

	// Api for User
	//r.HandleFunc("/api/user/get", api.GetUser)
	//r.HandleFunc("/api/user/create", api.CreateUser)
	// r.HandleFunc("/api/user/update", api.UpdateUser)
	// r.HandleFunc("/api/user/delete", api.DeleteUser)
	// r.HandleFunc("/api/user/addfriend", api.Newfriend)

	// // Api for messages
	// r.HandleFunc("/api/message/create", api.CreateMessage)
	// r.HandleFunc("/api/chat/get/", api.CreateMessage)

	// r.HandleFunc("/ws", util.WsEndpoint)

	//Starting Server and running on port 2020
	http.ListenAndServe(":2020", nil)
}
