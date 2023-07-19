package controllers

import (
	"TestTaskHaloLab/internal/api/validator"
	"TestTaskHaloLab/internal/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type RegionService interface {
	TemperatureMinimumInsideRegion(coordinates *dtos.Coordinates) (*dtos.Temperature, error)
	TemperatureMaximumInsideRegion(coordinates *dtos.Coordinates) (*dtos.Temperature, error)
}

type RegionController struct {
	RegionService RegionService
}

func NewRegionController(regionService RegionService) *RegionController {
	return &RegionController{RegionService: regionService}
}

// TemperatureMinimumInsideRegion
//
//	@Summary		Get Minimum temperature inside region
//	@Description	Current minimum temperature inside the region (a volume represented by the range of coordinates)
//	@Tags			region
//	@Produce		json
//	@Param			xMin	query		string					true	"Minimum X coordinate"
//	@Param			xMax	query		string					true	"Maximum X coordinate"
//	@Param			yMin	query		string					true	"Minimum Y coordinate"
//	@Param			yMax	query		string					true	"Maximum Y coordinate"
//	@Param			zMin	query		string					true	"Minimum Z coordinate"
//	@Param			zMax	query		string					true	"Maximum Z coordinate"
//	@Success		200		{object}	dtos.Temperature		"Minimum temperature in volume"
//	@Failure		400		{object}	dtos.MessageResponse	"Failed to return minimum temperature"
//	@Failure		422		{object}	dtos.MessageResponse	"Wrong coordinates"
//	@Failure		500		{object}	dtos.MessageResponse	"Internal Server Error"
//	@Router			/region/temperature/min [get]
func (rc *RegionController) TemperatureMinimumInsideRegion(c *fiber.Ctx) error {
	q := new(dtos.CoordinatesString)
	if err := c.QueryParser(q); err != nil {
		return err
	}
	coordinates, err := validator.ConvertAndValidateCoordinatesQueries(q)
	if err != nil {
		return err
	}

	temperature, err := rc.RegionService.TemperatureMinimumInsideRegion(coordinates)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(temperature)
}

// TemperatureMaximumInsideRegion
//
//	@Summary		Get Maximum temperature inside region
//	@Description	Current maximum temperature inside the region (a volume represented by the range of coordinates)
//	@Tags			region
//	@Produce		json
//	@Param			xMin	query		string					true	"Minimum X coordinate"
//	@Param			xMax	query		string					true	"Maximum X coordinate"
//	@Param			yMin	query		string					true	"Minimum Y coordinate"
//	@Param			yMax	query		string					true	"Maximum Y coordinate"
//	@Param			zMin	query		string					true	"Minimum Z coordinate"
//	@Param			zMax	query		string					true	"Maximum Z coordinate"
//	@Success		200		{object}	dtos.Temperature		"Maximum temperature in volume"
//	@Failure		400		{object}	dtos.MessageResponse	"Failed to return maximum temperature"
//	@Failure		422		{object}	dtos.MessageResponse	"Wrong coordinates"
//	@Failure		500		{object}	dtos.MessageResponse	"Internal Server Error"
//	@Router			/region/temperature/max [get]
func (rc *RegionController) TemperatureMaximumInsideRegion(c *fiber.Ctx) error {
	q := new(dtos.CoordinatesString)
	if err := c.QueryParser(q); err != nil {
		return err
	}
	coordinates, err := validator.ConvertAndValidateCoordinatesQueries(q)
	if err != nil {
		return err
	}

	temperature, err := rc.RegionService.TemperatureMaximumInsideRegion(coordinates)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(temperature)
}
