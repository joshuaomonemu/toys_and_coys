package models

import (
	database "app/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type Users struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	D_o_b        string `json:"d_o_b"`
	Username     string `json:"username"`
}

var client = database.CreateClient().Collection("users")
var ctx = context.Background()

// var registry *firestore.DocumentRef
var usr *Users
var payload []byte

//Function to create students

func CreateUser(usr *Users) bool {
	//Cleaning data for user registration
	//course_li := strings.SplitAfter(course_list, ",")
	_, err := client.Doc(usr.Username).Create(ctx, usr)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

// Function to read user information
func ReadUser(key string) (error, []byte) {
	data, err := client.Doc(key).Get(ctx)
	m := data.Data()

	if err == nil {
		user := &Users{
			Firstname:    m["firstname"].(string),
			Lastname:     m["lastname"].(string),
			Username:     m["username"].(string),
			Email:        m["email"].(string),
			Phone_number: m["phone_number"].(string),
		}
		payload, err = json.Marshal(user)
		if err != nil {
			fmt.Println("Error processing user details")
		}
		return err, nil
	} else {
		return nil, payload
	}
}

//// Function to delete user
//func DeleteUser(key string) bool {
//	_, err := client.Doc(key).Delete(ctx)
//	if err != nil {
//		// Handle any errors in an appropriate way, such as returning them.
//		log.Printf("An error has occurred: %s", err)
//		return true
//	} else {
//		return false
//	}
//}
//
//// Function to update user details
//func UpdateUser(key string, students *Students) bool {
//	_, err := client.Doc(key).Set(ctx, map[string]interface{}{
//		"department": students.Department,
//		"level":      students.Level,
//		"name":       students.Name,
//		"matno":      students.Matno,
//	})
//	if err != nil {
//		log.Printf("An error has occurred: %s", err)
//		return false
//	}
//	return true
//}
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
