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
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var upayload *structs.UserPayload
	var usr *structs.Users

	//Getting data
	// Using json.Unmarshal
	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal([]byte(body), &usr)
	if err != nil {
		panic(err)
	}
	//Sending data over to modelling page to carry out account creation and return a bool response on completion
	err, resp := models.CreateUser(usr)

	if err != nil {
		upayload = &structs.UserPayload{
			Success: false,
			Errors:  err.Error(),
			Message: "User account could not be created",
		}
	} else {
		upayload = &structs.UserPayload{
			Success: true,
			Data: structs.Users{
				Firstname:    usr.Firstname,
				Lastname:     usr.Lastname,
				Username:     usr.Username,
				Email:        usr.Email,
				Phone_number: usr.Phone_number,
				D_o_b:        usr.D_o_b,
			},
			Message: "User account created at " + resp,
		}
	}
	jsn, err := json.Marshal(upayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))

}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	var upayload *structs.UserPayload
	params := mux.Vars(r)
	id := params["id"]

	//Sourcing User Details from Database
	err, m := models.ReadUser(id)

	if err != nil {
		upayload = &structs.UserPayload{
			Success: false,
			Errors:  err.Error(),
			Message: "User account could not be read",
		}

	} else {
		upayload = &structs.UserPayload{
			Success: true,
			Data: structs.Users{
				Firstname:    m["Firstname"].(string),
				Lastname:     m["Lastname"].(string),
				Username:     m["Username"].(string),
				Email:        m["Email"].(string),
				Phone_number: m["Phone_number"].(string),
				D_o_b:        m["D_o_b"].(string),
			},
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
	var upayload *structs.UserPayload
	params := mux.Vars(r)
	id := params["id"]

	err := models.DeleteUser(id)
	if err != nil {
		upayload = &structs.UserPayload{
			Success: false,
			Errors:  err.Error(),
			Message: "User account could not be deleted",
		}
	} else {
		upayload = &structs.UserPayload{
			Success: true,
			Message: "User account deleted",
		}
	}

	jsn, err := json.Marshal(upayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var usr *structs.Users
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the bank")
	request_type := r.Method
	fmt.Println(request_type)
}
