package main

import (
	"log"
	"os"

	"shipping-api/db"
	"shipping-api/logger"
	"shipping-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using defaults")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	if err := logger.Init(); err != nil {
		log.Printf("Failed to initialize logger: %v", err)
	}
	defer logger.Close()

	logger.Info("Server starting", map[string]interface{}{
		"port": port,
		"env":  env,
	})

	db.ConnectDB()

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	log.Printf("🚀 Server starting on port %s (env: %s)...", port, env)
	logger.Info("Server started successfully", map[string]interface{}{
		"port": port,
		"env":  env,
	})

	router.Run(":" + port)
}
