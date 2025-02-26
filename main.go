package main

import (
	"example-api/config"
	"example-api/model"
	"example-api/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// Connect to the database
	config.ConnectDatabase()

	// Auto-migrate models
	config.DB.AutoMigrate(&model.Role{})

	// setup routes
	r := routes.SetupRoutes()

	port := os.Getenv("PORT")

	r.Run(":" + port)

}
