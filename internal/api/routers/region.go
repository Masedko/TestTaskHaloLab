package routers

import (
	"TestTaskHaloLab/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewRegionRouter(c *controllers.RegionController) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/temperature/min", c.TemperatureMinimumInsideRegion)
		router.Get("/temperature/max", c.TemperatureMaximumInsideRegion)
	}
}
