package main

import (
	_ "breeding-rest/docs" // Import the generated docs
	"breeding-rest/internal/routes"
	"breeding-rest/internal/utils"

	"os"
)

// @title Pepper Analytics API
// @version 1.0
// @description A Go REST API that calculates F2 generation probabilities in plant breeding genetics. Calculate inheritance patterns and expected phenotype distributions based on parent plant genotypes. Features Swagger documentation and handles multiple genetic traits with dominant/recessive expressions.
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
