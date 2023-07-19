package app

import (
	"TestTaskHaloLab/configuration"
	"TestTaskHaloLab/internal/api/controllers"
	"TestTaskHaloLab/internal/api/controllers/generateControllers"
	"TestTaskHaloLab/internal/api/middleware"
	"TestTaskHaloLab/internal/api/routers"
	"TestTaskHaloLab/internal/core/services"
	"TestTaskHaloLab/internal/core/services/generateServices"
	"TestTaskHaloLab/internal/data/database"
	"TestTaskHaloLab/internal/data/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"math/rand"
	"time"
)

// Run
//
//	@title			Test Task for Halo Lab
//	@version		1.0
//	@description	API for set of underwater sensors
//	@host			localhost:8080
//	@BasePath		/
func Run(c *configuration.EnvConfigModel) {
	db := database.ConnectDB(c)
	redisCache := database.ConnectRedis(c)

	sensorRepository := repository.NewSensorRepository(db)
	sensorDataRepository := repository.NewSensorDataRepository(db)
	fishSpeciesRepository := repository.NewFishSpeciesRepository(db)

	rand.Seed(time.Now().UnixNano())
	// GenerateSensors
	sensorGenerateService := generateServices.NewSensorGenerateService(sensorRepository)
	sensorGenerateController := generateControllers.NewSensorGenerateController(sensorGenerateService)
	sensorGenerateController.GenerateSensors()

	// GenerateSensorData
	sensorDataGenerateService := generateServices.NewSensorDataGenerateService(sensorRepository, sensorDataRepository, fishSpeciesRepository)
	sensorDataGenerateController := generateControllers.NewSensorDataGenerateController(sensorDataGenerateService)
	sensorDataGenerateController.GenerateSensorData()

	groupService := services.NewGroupService(redisCache, sensorDataRepository, fishSpeciesRepository)
	regionService := services.NewRegionService(sensorRepository, sensorDataRepository)
	sensorService := services.NewSensorService(sensorDataRepository)

	groupController := controllers.NewGroupController(groupService)
	regionController := controllers.NewRegionController(regionService)
	sensorController := controllers.NewSensorController(sensorService)

	groupRouter := routers.NewGroupRouter(groupController)
	regionRouter := routers.NewRegionRouter(regionController)
	sensorRouter := routers.NewSensorRouter(sensorController)

	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})

	//	Logger middleware for logging HTTP request/response details
	app.Use(logger.New())

	//	CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost",
		AllowHeaders: "GET",
	}))

	routers.SetupRoutes(app, groupRouter, regionRouter, sensorRouter)

	log.Fatal(app.Listen(":8080"))

}
