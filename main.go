package main

import (
	"app/routes"
	"fmt"
)

func main() {
	fmt.Println("SERVING AT PORT 2020")

	routes.Routes()
}
