package main

import (
	"app/models"
	"app/routes"
	"fmt"
)

var p *models.Students

func main() {
	fmt.Println("Project starts here")
	//models.RegStudent()
	routes.Routes()
}
