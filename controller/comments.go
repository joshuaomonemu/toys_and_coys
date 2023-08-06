package controller

import (
	"app/models"
	"app/structs"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var event *structs.Events

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	//Getting data
	// Using json.Unmarshal
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), &event)
	if err != nil {
		panic(err)
	}
	//Sending data over to modelling page to carry out account creation and return a bool response on completion
	resp := models.CreateEvent(event)

	//Sending response to response header
	if resp == true {
		w.WriteHeader(200)
		w.Header().Add("success", "true")
	} else {
		w.WriteHeader(303)
		w.Header().Add("success", "false")
	}
}

func ReadEvent(w http.ResponseWriter, r *http.Request) {
	//Sanitizing incoming string to prevent error when reading and
	params := mux.Vars(r)
	id := params["id"]

	//Sourcing User Details from Database
	payload := models.ReadEvent(id)
	jsn, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	id = strings.Trim(id, "{}")
	resp := models.DeleteEvent(id)
	if resp == true {
		w.WriteHeader(200)
		w.Header().Set("Deleted", "True")
	} else {
		w.WriteHeader(417)
		w.Header().Set("Deleted", "False")
	}
}
