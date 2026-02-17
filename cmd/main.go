package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"thai-monitor-platform/backend/internal/config"
	"thai-monitor-platform/backend/internal/routes"
)

func main() {
	// Load .env file from current directory
	godotenv.Load(".env")

	cfg := config.Load()

	// Validate TMD token
	if cfg.TMDToken == "" {
		log.Println("⚠️  WARNING: TOKEN_DATA_TMD is not set. Weather API will not work.")
	}

	app := fiber.New(fiber.Config{
		AppName: "Thai Monitor API v1.0",
	})

	// Setup routes
	routes.Setup(app, cfg)

	// Start server
	port := cfg.ServerPort
	if port == "" {
		port = "3001"
	}

	log.Printf("🚀 Thai Monitor Backend starting on port %s", port)
	log.Printf("📡 API available at http://localhost:%s/api", port)

	if err := app.Listen(":" + port); err != nil {
		log.Printf("❌ Failed to start server: %v", err)
		os.Exit(1)
	}
}
