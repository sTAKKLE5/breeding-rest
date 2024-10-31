package routes

import (
	"breeding-rest/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	switch mode {

	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")
	router.Static("/css", "./static/css")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := router.Group("/api/v1")
	v1.POST("/trait-probability", handlers.CalculateTraitProbability)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
