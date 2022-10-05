package models

import (
	database "app/database"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"strings"
)

type Students struct {
	Matno      string `json:"matno"`
	Name       string `json:"name"`
	Level      string `json:"level"`
	Department string `json:"department"`
	Password   string `json:"password"`
	Registry string `json:"registry"`
}

type Classes struct{
	Courses []string
	Class_id []string
}

var client = database.CreateClient().Collection("students")
var registryClient = database.CreateClient().Collection("registry")
var course_list = database.CreateClient().Collection("courses")
var ctx = context.Background()
//var registry *firestore.DocumentRef
var stu *Students

//Function to create students

func CreateStudent(matno, fullname, department, password, level, course_list string) bool {
	//Cleaning data for student registration
	course_li := strings.SplitAfter(course_list, ",")
	_, err := client.Doc(matno).Create(ctx, map[string]interface{}{
		"department": department,
		"level":      level,
		"name":       fullname,
		"password":   password,
		"matno":      matno,
		"course_list": course_li,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

// Function to read students information
func ReadStudent(key string) (error, []byte) {
	data, err := client.Doc(key).Get(ctx)
	m := data.Data()
	registry := m["registry"].(*firestore.DocumentRef)

	if err == nil {
		registry_data, registry_err := registryClient.Doc(registry.ID).Get(ctx)
		registryData := registry_data.Data()
		fmt.Println(registryData)

		if registry_err == nil{
			student := &Students{
				Name: m["name"].(string),
				Matno: m["matno"].(string),
				Department: registryData["name"].(string),
				Level: "400",
			}
			payload, _err := json.Marshal(student)
			if _err != nil{
				fmt.Println("Error processing student")
			}
			return nil, payload
		} else {
			return registry_err, nil
		}
	} else {
		return err, nil
	}
}

// Function to delete student
func DeleteStudent(key string) bool {
	_, err := client.Doc(key).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return true
	} else {
		return false
	}
}

// Function to update student details
func UpdateStudent(key string, students *Students) bool {
	_, err := client.Doc(key).Set(ctx, map[string]interface{}{
		"department": students.Department,
		"level":      students.Level,
		"name":       students.Name,
		"matno":      students.Matno,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

func LoginStudent(key, password string)bool {
	var x map[string]interface{}
	query := client.Where("matno", "==", key).Where("password", "==", password).Documents(ctx)
	for{
		doc, err := query.Next()
		if err == iterator.Done{
			break
		}
		//Ignoring error handling on testing phase
		//if err != nil{
		//	fmt.Println(err)
		//}
		x = doc.Data()

		//fmt.Println(doc.Data())
	}
	if len(x) == 0{
		return false
	}else{
		return true
	}
}

func GetAllCourses() ([]string, []string){
	c := &Classes{
		Courses: []string{},
		Class_id: []string{},
	}
	//Please do not touch anything it doesn't even make sense to me
	var course_code string
	var class_iden string
	var course_outline []string
	var class_id []string
	//

	iter := course_list.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		m := doc.Data()


		course_code = m["course_code"].(string)
		c.Courses = append(c.Courses, course_code)
		course_outline = c.Courses

		class_iden = m["class_id"].(string)
		c.Class_id = append(c.Class_id, class_iden)
		class_id = c.Class_id
	}
	return course_outline, class_id
}
//func GetAllCourses(){
//
//}
