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
	err, resp := models.CreateComment(id, event)
	if err != nil {
		cpayload = &structs.CommentPayload{
			Succeeded: false,
			Errors:    err,
			Message:   "Error Occurred when commenting",
		}
		fmt.Println(cpayload)
	} else {
		cpayload = &structs.CommentPayload{
			Succeeded: true,
			Data: structs.EventComment{
				User:    resp.User,
				Comment: resp.Comment,
				Time:    resp.Time,
			},
			Errors:  nil,
			Message: "Commented on event",
		}
	}
	jsn, err := json.Marshal(cpayload)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))
}
