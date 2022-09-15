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
	resp := models.CreateStudent()
	w.Header().Set("Page", "Create")
	if resp == true {
		w.Header().Set("Created", "Done")
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