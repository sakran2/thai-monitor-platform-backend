package routes

import (
	"thai-monitor-platform/backend/internal/config"
	"thai-monitor-platform/backend/internal/handler"
	"thai-monitor-platform/backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App, cfg *config.Config) {
	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "thai-monitor-backend",
		})
	})

	// API group
	api := app.Group("/api")

	// Weather routes
	weatherSvc := service.NewWeatherService(cfg)
	weatherHandler := handler.NewWeatherHandler(weatherSvc)

	weather := api.Group("/weather")
	weather.Get("/forecast", weatherHandler.GetForecast)
	weather.Get("/provinces", weatherHandler.GetProvinces)

	// Earthquake routes
	earthquakeHandler := handler.NewEarthquakeHandler(cfg)
	earthquake := api.Group("/earthquake")
	earthquake.Get("/latest", earthquakeHandler.GetLatest)
}
