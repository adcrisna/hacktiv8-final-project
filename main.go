package main

import (
	"final-project/routes"
	"fmt"
)

func main() {
	router := routes.RouterServer()
	err := router.Run(":3000")
	if err != nil {
		fmt.Println("Error Saat Connect Ke DB", err)
	}
}
