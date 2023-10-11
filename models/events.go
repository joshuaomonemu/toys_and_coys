package models

import (
	"app/database"
	"app/structs"
	"fmt"
	"log"
	"math/rand"
)

var po = database.CreateClient()
var m map[string]interface{}

func CreateEvent(event *structs.Events) (error, string) {
	post_val := "post-" + fmt.Sprintln(rand.Intn(999999))

	res, err1 := po.Collection("events").Doc(post_val).Create(ctx, event)
	if err1 != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err1)
		return err1, ""
	}
	return err1, res.UpdateTime.GoString()
}

// Function to read events
func ReadEvent(key string) (error, map[string]interface{}) {
	data, err := po.Collection("events").Doc(key).Get(ctx)
	m = data.Data()
	if err != nil {
		err1 := err
		return err1, m
	} else {
		var err1 error
		return err1, m
	}
}

// Function to delete an event
func DeleteEvent(key string) error {
	_, err := po.Collection("events").Doc(key).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return err
}
