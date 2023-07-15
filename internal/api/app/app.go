package app

import (
	"TestTaskHaloLab/configuration"
	"TestTaskHaloLab/internal/data/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

// Run
// @title			Test Task for Halo Lab
// @version		1.0
// @description	API for set of underwater sensors
// @host			localhost:8080
// @BasePath		/
func Run(c *configuration.EnvConfigModel) {

	db := database.ConnectDB(c)

	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})

	//	Logger middleware for logging HTTP request/response details
	app.Use(logger.New())

	//	CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	routers.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
