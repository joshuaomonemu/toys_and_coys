package main

import (
	"app/models"
	"encoding/json"
	"fmt"
)

var p *models.Students

func main() {
	fmt.Println("Project starts here")
	//models.RegStudent()
	_, new_sample := models.ReadStudent("m.18EEE12811")
	json.Unmarshal(new_sample, &p)
	dept := string(p.Department)
	matno := string(p.Matno)
	name := string(p.Name)
	level := string(p.Level)
	fmt.Println(dept, matno, name, level)
}
