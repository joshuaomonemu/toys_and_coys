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

func CreateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var event *structs.EventComment
	var cpayload *structs.CommentPayload

	//Getting data
	// Using json.Unmarshal
	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal([]byte(body), &event)
	if err != nil {
		panic(err)
	}
	//Sending data over to modelling page to carry out account creation and return a bool response on completion
	err, _ = models.CreateComment(id, event)
	if err != nil {
		cpayload = &structs.CommentPayload{
			Success: false,
			Errors:  err.Error(),
			Message: "Error Occurred when commenting",
		}
		fmt.Println(cpayload)
	} else {
		cpayload = &structs.CommentPayload{
			Success: true,
			Message: "Commented on event",
		}
	}
	jsn, err := json.Marshal(cpayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	var cpayload *structs.CommentPayload

	params := mux.Vars(r)
	id := params["id"]

	err := models.DeleteComment(id)
	if err != nil {
		cpayload = &structs.CommentPayload{
			Success: false,
			Errors:  err.Error(),
			Message: "This comment could not be deleted",
		}
	} else {
		cpayload = &structs.CommentPayload{
			Success: true,
			Message: "This comment has been deleted",
		}
	}

	jsn, err := json.Marshal(cpayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}

func GetallComments(w http.ResponseWriter, r *http.Request) {
	CT := r.Header.Get("Content-Type")
	fmt.Println(CT)

	var cpayload *structs.CommentPayload

	data, err := models.GetallComments()
	if err != nil {
		cpayload = &structs.CommentPayload{
			Success: false,
			Errors:  err.Error(),
			Message: "Couldn't trace any comments",
		}
	} else {
		cpayload = &structs.CommentPayload{
			Success: true,
			Data:    data,
			Message: "This comment has been deleted",
		}
	}

	jsn, err := json.Marshal(cpayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}
