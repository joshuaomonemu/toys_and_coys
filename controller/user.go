package controller

import (
	"app/models"
	"app/structs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var usr *structs.Users
var upayload *structs.UserPayload

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Getting data
	// Using json.Unmarshal
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), &usr)
	if err != nil {
		panic(err)
	}
	//Sending data over to modelling page to carry out account creation and return a bool response on completion
	resp := models.CreateUser(usr)

	//Sending response to response header
	if resp == true {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(301)
	}
}

func ReadUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	//Sourcing User Details from Database
	err, m := models.ReadUser(id)

	if err != nil {
		upayload = &structs.UserPayload{
			Succeeded: false,
			Errors:    err,
		}

	} else {
		upayload = &structs.UserPayload{
			Succeeded: true,
			Data: structs.Users{
				Firstname:    m["firstname"].(string),
				Lastname:     m["lastname"].(string),
				Username:     m["username"].(string),
				Email:        m["email"].(string),
				Phone_number: m["phone_number"].(string),
				D_o_b:        m["d_o_b"].(string),
			},
			Errors:  nil,
			Message: "User data printed successfully",
		}
	}

	jsn, err := json.Marshal(upayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	//Sanitizing incoming string to prevent error when reading
	id = strings.Trim(id, "{}")
	fmt.Println(id)
	resp := models.DeleteUser(id)
	if resp == true {
		w.WriteHeader(200)
		w.Header().Set("Deleted", "True")
	} else {
		w.WriteHeader(303)
		w.Header().Set("Deleted", "False")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), &usr)
	if err != nil {
		panic(err)
	}

	resp := models.UpdateUser(id, usr)

	if resp == true {
		w.Header().Set("update", "complete")
	} else {
		w.Header().Set("update", "failed")
	}
}

//
//func LoginUser(w http.ResponseWriter, r *http.Request) {
//	username := r.FormValue("username")
//	password := r.FormValue("password")
//
//	resp := models.LoginUser(username, password)
//	if resp == true {
//		w.Header().Set("login", "successful")
//	} else {
//		w.Header().Set("login", "failed")
//	}
//}
