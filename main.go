package main

import (
	"app/mailer"
	"app/routes"
	"fmt"
)

func main() {
	fmt.Println("SERVING AT PORT 2020")
	mailer.Mailer()

	routes.Routes()
}
