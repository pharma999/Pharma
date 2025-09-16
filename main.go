package main

import (
	"demo/database"
	"demo/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/twilio/twilio-go/client"
)

func main() {
	fmt.Println("Welcome to the demo api...")

	database.ConnectDB()
	// Setup router
	r := gin.Default()
	routes.SetupRouter(r)

	r.Run(":8080")

}
