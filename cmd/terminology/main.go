package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere-platform/zarish-terminology-server/internal/data"
	"github.com/zarishsphere-platform/zarish-terminology-server/internal/database"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default to 8081 for Terminology
	}

	// Connect to Database
	database.Connect()

	// Load Data
	dataDir := "../../zarish-fhir-data"
	if envDir := os.Getenv("DATA_DIR"); envDir != "" {
		dataDir = envDir
	}
	data.LoadTerminologyData(dataDir)


	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"service": "zarish-terminology-server",
		})
	})

	log.Printf("Starting Zarish Sphere Terminology Server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
