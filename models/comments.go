package models

import (
	"app/structs"
	"cloud.google.com/go/firestore"
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func CreateComment(key string, com *structs.EventComment) (error, *structs.EventComment) {
	err1 := errors.New("event data does not exist")

	query := client.Collection("events").Where("Event_id", "==", key)
	docRef := client.Collection("events").Doc(key).ID

	// Retrieve the documents matching the query
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("Error querying Firestore: %v", err)
	}
	if len(docs) <= 0 {
		return err1, nil
	}

	co_val := "com-" + fmt.Sprintln(rand.Intn(999999)) + key
	_, err = client.Collection("comments").Doc(co_val).Create(ctx, map[string]interface{}{
		"User":    com.User,
		"Comment": com.Comment,
		"Time":    firestore.ServerTimestamp,
		"Ref":     docRef,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err, com
	}
	return err, com
}
