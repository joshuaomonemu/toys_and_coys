package api

import (
	"app/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

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
	fmt.Println(id)
	resp := models.DeleteStudent(id)
	if resp == true {
		w.Header().Set("Deleted", "True")
	} else {
		w.Header().Set("Deleted", "False")
	}
}
