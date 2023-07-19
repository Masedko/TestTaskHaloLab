package controllers

import (
	"TestTaskHaloLab/internal/api/validator"
	"TestTaskHaloLab/internal/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type SensorService interface {
	AverageTemperatureInInterval(group *dtos.Group, timeInterval *dtos.TimeInterval) (*dtos.Temperature, error)
}

type SensorController struct {
	SensorService SensorService
}

func NewSensorController(sensorService SensorService) *SensorController {
	return &SensorController{SensorService: sensorService}
}

// AverageTemperatureInInterval
//
//	@Summary		Get average temperature inside region for particular sensor
//	@Description	Average temperature detected by a particular sensor between the specified date/time pairs (UNIX timestamps)
//	@Tags			sensor
//	@Produce		json
//	@Param			fromDateTime	query		string					true	"Start Date Time"
//	@Param			untilDateTime	query		string					true	"End Date Time"
//	@Param			codeName		path		string					true	"Name of the group '{greekLetter} {int}'"
//	@Success		200				{object}	dtos.Temperature		"Average temperature"
//	@Failure		400				{object}	dtos.MessageResponse	"Failed to return average temperature"
//	@Failure		422				{object}	dtos.MessageResponse	"Wrong UNIX time"
//	@Failure		500				{object}	dtos.MessageResponse	"Internal Server Error"
//	@Router			/sensor/{codeName}/temperature/average [get]
func (sc *SensorController) AverageTemperatureInInterval(c *fiber.Ctx) error {
	codeNameString := c.Params("codeName")
	group, err := validator.ValidateCodeName(codeNameString)
	if err != nil {
		return err
	}

	q := new(dtos.TimeIntervalUnixTimestamp)
	if err := c.QueryParser(q); err != nil {
		return err
	}
	timeInterval, err := validator.ConvertAndValidateTimeIntervalQueries(q)
	if err != nil {
		return err
	}

	temperature, err := sc.SensorService.AverageTemperatureInInterval(group, timeInterval)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(temperature)
}
