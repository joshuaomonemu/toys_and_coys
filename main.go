package main

import (
	"app/routes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Project starts here")
	fmt.Println(os.Getenv("PORT"))
	routes.Routes()
}
