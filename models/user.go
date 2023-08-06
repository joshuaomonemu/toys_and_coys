package models

import (
	database "app/database"
	"app/structs"
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

var client = database.CreateClient()
var ctx = context.Background()

// var registry *firestore.DocumentRef
var usr *structs.Users

//Function to create students

func CreateUser(usr *structs.Users) bool {
	//Cleaning data for user registration
	//course_li := strings.SplitAfter(course_list, ",")
	_, err := client.Collection("users").Doc(usr.Username).Create(ctx, usr)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

// Function to read user information
func ReadUser(key string) (error, map[string]interface{}) {
	data, err := client.Collection("users").Doc(key).Get(ctx)
	m = data.Data()
	if err != nil {
		err1 := err
		return err1, m
	} else {
		var err1 error
		return err1, m
	}
}

// Function to delete user
func DeleteUser(key string) bool {
	_, err := client.Collection("users").Doc(key).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return true
	} else {
		return false
	}
}

// Function to update user details
func UpdateUser(key string, usr *structs.Users) bool {
	_, err := client.Doc(key).Set(ctx, map[string]interface{}{
		"phone_number": usr.Phone_number,
		"email":        usr.Email,
		"firstname":    usr.Firstname,
		"lastname":     usr.Lastname,
	}, firestore.MergeAll)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

//
//func LoginUser(key, password string) bool {
//	var x map[string]interface{}
//	query := client.Where("matno", "==", key).Where("password", "==", password).Documents(ctx)
//	for {
//		doc, err := query.Next()
//		if err == iterator.Done {
//			break
//		}
//		//Ignoring error handling on testing phase
//		//if err != nil{
//		//	fmt.Println(err)
//		//}
//		x = doc.Data()
//
//		//fmt.Println(doc.Data())
//	}
//	if len(x) == 0 {
//		return false
//	} else {
//		return true
//	}
//}
