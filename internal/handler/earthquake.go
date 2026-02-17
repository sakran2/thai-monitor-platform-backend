package handler

import (
	"thai-monitor-platform/backend/internal/collector"
	"thai-monitor-platform/backend/internal/config"

	"github.com/gofiber/fiber/v2"
)

type EarthquakeHandler struct {
	Config *config.Config
}

func NewEarthquakeHandler(cfg *config.Config) *EarthquakeHandler {
	return &EarthquakeHandler{Config: cfg}
}

func (h *EarthquakeHandler) GetLatest(c *fiber.Ctx) error {
	earthquakes, err := collector.FetchLatestEarthquakes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": earthquakes,
	})
}
