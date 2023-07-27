package api

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var event *models.Events

func CreateEvent(w http.ResponseWriter, r *http.Request) {
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
	resp := models.CreateEvent(event)

	//Sending response to response header
	if resp == true {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(303)
	}
}
