package routers

import (
	"TestTaskHaloLab/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewSensorRouter(c *controllers.SensorController) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/:codeName/temperature/average", c.AverageTemperatureInInterval)
	}
}
