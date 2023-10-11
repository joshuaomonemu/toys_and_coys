package models

import (
	"app/structs"
	"cloud.google.com/go/firestore"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"math/rand"
)

func CreateComment(key string, com *structs.EventComment) (error, *structs.EventComment) {

	iter := client.Collection("events").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(doc)
	}

	docRef := po.Collection("events").Doc(key)
	fmt.Println(docRef)

	co_val := "com-" + fmt.Sprintln(rand.Intn(999999)) + key
	_, err := po.Collection("comments").Doc(co_val).Create(ctx, map[string]interface{}{
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
