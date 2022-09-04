package api

import (
	"app/models"
	"net/http"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	res := models.CreateStudent()
	if res = true {
		w.Header().Set("Created", "Done")
	}
}
