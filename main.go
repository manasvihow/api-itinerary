package main

import (
	"example/go-v1/handlers"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "example/go-v1/docs" 
)

// @title           Itinerary API
// @version         1.0
// @description     An API for generating travel itineraries and returning them as a PDF.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	_ = os.Mkdir("output", 0755)

	router := gin.Default()

	// API endpoint
	router.POST("/api/itinerary", handlers.GenerateItinerary)

	// Swagger UI endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
