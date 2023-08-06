package models

import (
	"app/database"
	"app/structs"
	"fmt"
	"log"
	"math/rand"
)

var po = database.CreateClient()
var err_arr []string
var m map[string]interface{}
var epayload *structs.EventPayload

func CreateEvent(event *structs.Events) bool {
	post_val := "post-" + fmt.Sprintln(rand.Intn(999999))

	_, err := po.Collection("events").Doc(post_val).Create(ctx, event)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

// Function to read events
func ReadEvent(key string) *structs.EventPayload {
	data, err := po.Collection("events").Doc(key).Get(ctx)
	m = data.Data()
	if err != nil {
		err1 := err.Error()
		err_arr = append(err_arr, err1)

		epayload = &structs.EventPayload{
			Succeeded: false,
			Errors:    []string(err_arr),
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
				Comments: structs.EventComment{
					User:    m["Comments"].(map[string]interface{})["User"].(string),
					Comment: m["Comments"].(map[string]interface{})["Comment"].(string),
					Time:    m["Comments"].(map[string]interface{})["Time"].(string),
				},
			},
			Errors: []string(err_arr),
		}
	}
	return epayload
}

// Function to delete an event
func DeleteEvent(key string) bool {
	_, err := po.Collection("events").Doc(key).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return true
	} else {
		return false
	}
}
