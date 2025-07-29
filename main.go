package main

import (
	"example/go-v1/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the output directory if it doesn't exist
	_ = os.Mkdir("output", 0755)

	router := gin.Default()

	router.POST("/api/itinerary", handlers.GenerateItinerary)

	router.Run(":8080")
}
