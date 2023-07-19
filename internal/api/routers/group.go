package routers

import (
	"TestTaskHaloLab/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewGroupRouter(c *controllers.GroupController) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/:groupName/temperature/average", c.TemperatureAverage)
		router.Get("/:groupName/transparency/average", c.TransparencyAverage)
		router.Get("/:groupName/species", c.Species)
		router.Get("/:groupName/species/top/:N", c.SpeciesTopN)
	}
}
