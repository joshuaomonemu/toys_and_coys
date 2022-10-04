package api

import (
	"app/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var p *models.Students

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	//Getting data from filled forms to sedn across to the database
	matno := r.FormValue("matno")
	fullname := r.FormValue("fullname")
	department := r.FormValue("department")
	password := r.FormValue("password")
	level := r.FormValue("level")
	course_list := r.FormValue("courses")

	//Sending data over to modelling page to carry out account creation and return a bool response on completion
	resp := models.CreateStudent(matno, fullname, department, password, level, course_list)

	//Sending response to response header
	if resp == true {
		w.Header().Set("Created", "true")
	} else {
		w.Header().Set("Created", "False")
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	//Sanitizing incoming string to prevent error when reading
	id = strings.Trim(id, "{}")
	fmt.Println(id)
	resp := models.DeleteStudent(id)
	if resp == true {
		w.Header().Set("Deleted", "True")
	} else {
		w.Header().Set("Deleted", "False")
	}
}

func ReadStudent(w http.ResponseWriter, r *http.Request){
	//Sanitizing incoming string to prevent error when reading and
	params := mux.Vars(r)
	id := params["id"]
	id = strings.Trim(id, "{}")
	fmt.Println(id)

	//Sourcing User Details from Database
	_, new_sample := models.ReadStudent(id)
	err := json.Unmarshal(new_sample, &p)
	if err != nil{
		log.Fatal("Error unmarshalling JSON")
	}
	dept, matno, name, level := string(p.Department), string(p.Matno), string(p.Name), string(p.Level)

	//Sending details as header values
	w.Header().Set("Department", dept)
	w.Header().Set("Matriculation Number", matno)
	w.Header().Set("Name", name)
	w.Header().Set("Level", level)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	id = strings.Trim(id, "{}")
	fmt.Println(id)


	//Sourcing data from request header
	dept := w.Header().Get("dept")
	matno := w.Header().Get("matno")
	name := w.Header().Get("name")
	level := w.Header().Get("level")

	//Data Schema for updating students information
	packed := &models.Students{
		Level: level,
		Department: dept,
		Matno: matno,
		Name: name,
	}
	//Getting response from server to check if update was successful
	resp := models.UpdateStudent(id, packed)
	if resp == true{
	w.Header().Set("update", "complete")
}else{
	w.Header().Set("update", "failed")
	}
}


func LoginStudent(w http.ResponseWriter, r *http.Request){
	matno := r.FormValue("matno");
	password := r.FormValue("password")

	resp := models.LoginStudent(matno, password)
	if resp == true{
		w.Header().Set("login", "successful")
	}else{
		w.Header().Set("login", "failed")
	}
}
