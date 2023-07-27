package api

import (
	"app/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var usr *models.Users

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
		w.WriteHeader(303)
	}
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	//Sanitizing incoming string to prevent error when reading and
	params := mux.Vars(r)
	id := params["id"]
	id = strings.Trim(id, "{}")

	//Sourcing User Details from Database
	_, payload := models.ReadUser(id)
	jsn, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

//func DeleteUser(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id := params["id"]
//	//Sanitizing incoming string to prevent error when reading
//	id = strings.Trim(id, "{}")
//	fmt.Println(id)
//	resp := models.DeleteUser(id)
//	if resp == true {
//		w.Header().Set("Deleted", "True")
//	} else {
//		w.Header().Set("Deleted", "False")
//	}
//}
//

//func UpdateUser(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id := params["id"]
//	id = strings.Trim(id, "{}")
//	fmt.Println(id)
//
//	//Sourcing data from request header
//	dept := w.Header().Get("dept")
//	matno := w.Header().Get("matno")
//	name := w.Header().Get("name")
//	level := w.Header().Get("level")
//
//	//Data Schema for updating students information
//	packed := &models.Students{
//		Level:      level,
//		Department: dept,
//		Matno:      matno,
//		Name:       name,
//	}
//	//Getting response from server to check if update was successful
//	resp := models.UpdateUser(id, packed)
//	if resp == true {
//		w.Header().Set("update", "complete")
//	} else {
//		w.Header().Set("update", "failed")
//	}
//}
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
