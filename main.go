package main

import (
	_ "breeding-rest/docs" // Import the generated docs
	"breeding-rest/internal/routes"
	"breeding-rest/internal/utils"

	"os"
)

// @title Pepper Analytics API
// @version 1.0
// @description The Pepper Analytics AI platform is a comprehensive system designed to manage and analyze data related to plant breeding projects. It provides a RESTful API for various operations, including managing species, varieties, plants, fertilizers, observations, seed banks, breeding projects, and notes. The platform uses the Gin framework for routing and middleware, and it integrates Swagger for API documentation. The system ensures data integrity through transaction middleware and supports CRUD operations for all major entities.
// @host localhost:8080
// @BasePath /api/v1
// @contact.name: -
// @contact.email: -
// @license.name: MIT
// @license.url: https://opensource.org/licenses/MIT
// @schemes http https

func init() {
	// Load environment variables
	utils.LoadEnvVars()
}

func main() {
	r := routes.SetupRouter()
	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
