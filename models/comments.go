package models

import (
	"app/structs"
	"cloud.google.com/go/firestore"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"math/rand"
)

func CreateComment(key string, com *structs.EventComment) (error, *structs.EventComment) {
	err1 := errors.New("event data does not exist")
	co_val := "comxxx-" + fmt.Sprintln(rand.Intn(9999999))
	commentid := "com-" + fmt.Sprintln(rand.Intn(99999))

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

	_, err = client.Collection("comments").Doc(co_val).Create(ctx, map[string]interface{}{
		"User":      com.User,
		"Comment":   com.Comment,
		"Time":      firestore.ServerTimestamp,
		"Ref":       docRef,
		"Commentid": commentid,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err, com
	}
	return err, com
}

func DeleteComment(key string) error {
	_, err := po.Collection("comments").Doc(key).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return err
}

func GetallComments() ([]structs.EventComment, error) {
	iter := client.Collection("comments").Documents(ctx)

	// Define a slice to store the results
	var comments []structs.EventComment

	// Iterate over the documents and append them to the slice
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		// Create a variable of type Student
		var comment structs.EventComment

		// Unmarshal Firestore document into the struct
		if err := doc.DataTo(&comment); err != nil {
			log.Fatalf("Error unmarshaling Firestore document: %v", err)
		}

		// Append the struct to the slice
		comments = append(comments, comment)
	}

	return comments, nil
	//// Now 'students' contains the data from Firestore in the form of a slice of structs
	//fmt.Printf("Number of students: %d\n", len(students))
	//for _, s := range students {
	//	fmt.Printf("Name: %s, Age: %d, Major: %s\n", s.Name, s.Age, s.Major)
	//}
}
