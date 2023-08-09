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
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {

	var event *structs.Events
	var epayload *structs.EventPayload

	//Getting data
	// Using json.Unmarshal
	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal([]byte(body), &event)
	if err != nil {
		panic(err)
	}
	//Sending data over to modelling page to carry out account creation and return a bool response on completion
	err1, resp := models.CreateEvent(event)

	if err1 != nil {
		epayload = &structs.EventPayload{
			Succeeded: false,
			Errors:    err1,
			Message:   "Event could not be created",
		}
	} else {
		epayload = &structs.EventPayload{
			Succeeded: true,
			Data: structs.Events{
				User:    event.User,
				Content: event.Content,
				Text:    event.Text,
				Time:    event.Time,
				Likes:   event.Likes,
			},
			Errors:  nil,
			Message: "User account created at " + resp,
		}
	}
	jsn, err := json.Marshal(epayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

func ReadEvent(w http.ResponseWriter, r *http.Request) {
	var epayload *structs.EventPayload

	params := mux.Vars(r)
	id := params["id"]

	err, m := models.ReadEvent(id)

	if err != nil {
		epayload = &structs.EventPayload{
			Succeeded: false,
			Errors:    err,
		}

	} else {
		epayload = &structs.EventPayload{
			Succeeded: true,
			Data: structs.Events{
				User:    m["User"].(string),
				Content: m["Content"].(string),
				Text:    m["Text"].(string),
				Time:    m["Time"].(string),
				Likes:   m["Likes"].(int64),
			},
			Errors: nil,
		}
	}
	jsn, err := json.Marshal(epayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var epayload *structs.EventPayload
	params := mux.Vars(r)
	id := params["id"]

	err := models.DeleteEvent(id)
	if err != nil {
		epayload = &structs.EventPayload{
			Succeeded: false,
			Errors:    err,
			Message:   "This event could not be deleted",
		}
	} else {
		epayload = &structs.EventPayload{
			Succeeded: true,
			Errors:    err,
			Message:   "This event has been deleted",
		}
	}

	jsn, err := json.Marshal(epayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}
