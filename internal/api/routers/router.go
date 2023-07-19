package routers

import (
	_ "TestTaskHaloLab/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App,
	groupRouter func(router fiber.Router),
	regionRouter func(router fiber.Router),
	sensorRouter func(router fiber.Router)) {

	api := app.Group("")

	//	Use 'swag init' to generate new /docs files, details: https://github.com/gofiber/swagger#usage
	api.Get("/docs/*", swagger.HandlerDefault)
	// Redirect to docs
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/docs/")
	})

	api.Route("/group", groupRouter)

	api.Route("/region", regionRouter)

	api.Route("/sensor", sensorRouter)
}
