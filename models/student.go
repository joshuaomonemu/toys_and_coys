package models

import (
	database "app/database"
	"context"
	"encoding/json"
	"log"
)

type Students struct {
	Matno      string `json:"matno"`
	Name       string `json:"name"`
	Level      string `json:"level"`
	Department string `json:"department"`
	Password   string `json:"password"`
}

var client = database.CreateClient().Collection("students")
var ctx = context.Background()
var stu *Students
func CreateStudent() bool {
	//Cleaning data for student registration
	_, err := client.Doc("m.18EEE12811").Create(ctx, map[string]interface{}{
		"department": "mechanical",
		"level":      "500",
		"name":       "spades michael",
		"password":   "1234dog",
		"matno": "m.18EEE12811",
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

func ReadStudent(key string) (error,[]byte) {
	data, err := client.Doc(key).Get(ctx)
	if err != nil {
		return err, nil
	}
	m := data.Data()
	matno := m["matno"].(string)
	department := m["department"].(string)
	level := m["level"].(string)
	name := m["name"].(string)
	password := m["password"].(string)


	payload := &Students{
		Matno:       matno,
		Name:    name,
		Level: level,
		Department: department,
		Password: password,
	}
	bs, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Error converting to JSON")
	}
	return nil, bs
}

func DeleteStudent(key string) bool{
	_, err := client.Doc(key).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return true
	}else {
		return false
	}
}