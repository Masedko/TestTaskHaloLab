package controllers

import (
	"TestTaskHaloLab/internal/api/validator"
	"TestTaskHaloLab/internal/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type GroupService interface {
	TemperatureAverage(groupName string) (*dtos.Temperature, error)
	TransparencyAverage(groupName string) (*dtos.Transparency, error)
	Species(groupName string) (*[]dtos.Fish, error)
	SpeciesTopN(groupName string, N int, timeInterval *dtos.TimeInterval) (*[]dtos.Fish, error)
}

type GroupController struct {
	GroupService GroupService
}

func NewGroupController(groupService GroupService) *GroupController {
	return &GroupController{GroupService: groupService}
}

// TemperatureAverage
//
//	@Summary		Calculate temperature average for a group
//	@Description	Calculates the temperature average for the specified group.
//	@Tags			group
//	@Produce		json
//	@Param			groupName	path		string					true	"Name of the group '{greekLetter}'"
//	@Success		200			{object}	dtos.Temperature		"Temperature"
//	@Failure		400			{object}	dtos.MessageResponse	"Failed to return temperature average"
//	@Failure		422			{object}	dtos.MessageResponse	"Wrong groupName format"
//	@Failure		500			{object}	dtos.MessageResponse	"Internal Server Error"
//	@Router			/group/{groupName}/temperature/average [get]
func (gc *GroupController) TemperatureAverage(c *fiber.Ctx) error {
	groupNameString := c.Params("groupName")
	groupName, err := validator.ValidateGroupName(groupNameString)
	if err != nil {
		return err
	}
	temperature, err := gc.GroupService.TemperatureAverage(groupName)
	return c.Status(fiber.StatusOK).JSON(temperature)
}

// TransparencyAverage
//
//	@Summary		Calculate transparency average for a group
//	@Description	Calculates the transparency average for the specified group.
//	@Tags			group
//	@Produce		json
//	@Param			groupName	path		string					true	"Name of the group '{greekLetter}'"
//	@Success		200			{object}	dtos.Transparency		"Transparency"
//	@Failure		400			{object}	dtos.MessageResponse	"Failed to return transparency average"
//	@Failure		422			{object}	dtos.MessageResponse	"Wrong groupName format"
//	@Failure		500			{object}	dtos.MessageResponse	"Internal Server Error"
//	@Router			/group/{groupName}/transparency/average [get]
func (gc *GroupController) TransparencyAverage(c *fiber.Ctx) error {
	groupNameString := c.Params("groupName")
	groupName, err := validator.ValidateGroupName(groupNameString)
	if err != nil {
		return err
	}
	transparency, err := gc.GroupService.TransparencyAverage(groupName)
	return c.Status(fiber.StatusOK).JSON(transparency)
}

// Species
//
//	@Summary		Get Species inside group
//	@Description	Full list of species (with counts) currently detected inside the group
//	@Tags			group
//	@Produce		json
//	@Param			groupName	path		string					true	"Name of the group '{greekLetter}'"
//	@Success		200			{object}	[]dtos.Fish				"Fish species with count"
//	@Failure		400			{object}	dtos.MessageResponse	"Failed to return fish species"
//	@Failure		422			{object}	dtos.MessageResponse	"Wrong groupName format"
//	@Failure		500			{object}	dtos.MessageResponse	"Internal Server Error"
//	@Router			/group/{groupName}/species [get]
func (gc *GroupController) Species(c *fiber.Ctx) error {
	groupNameString := c.Params("groupName")
	groupName, err := validator.ValidateGroupName(groupNameString)
	if err != nil {
		return err
	}
	fish, err := gc.GroupService.Species(groupName)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fish)
}

// SpeciesTopN
//
//	@Summary		Get Top N Species inside group
//	@Description	List of top N species (with counts) currently detected inside the group
//	@Tags			group
//	@Produce		json
//	@Param			groupName		path		string					true	"Name of the group '{greekLetter}'"
//	@Param			N				path		int						true	"Top N"
//	@Param			fromDateTime	query		string					false	"Start Date Time"
//	@Param			untilDateTime	query		string					false	"End Date Time"
//	@Success		200				{object}	[]dtos.Fish				"Fish species with count"
//	@Failure		400				{object}	dtos.MessageResponse	"Failed to return fish species"
//	@Failure		422				{object}	dtos.MessageResponse	"Wrong groupName format"
//	@Failure		500				{object}	dtos.MessageResponse	"Internal Server Error"
//	@Router			/group/{groupName}/species/top/{N} [get]
func (gc *GroupController) SpeciesTopN(c *fiber.Ctx) error {
	groupNameString := c.Params("groupName")
	groupName, err := validator.ValidateGroupName(groupNameString)
	if err != nil {
		return err
	}
	N, err := c.ParamsInt("N")
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
	fish, err := gc.GroupService.SpeciesTopN(groupName, N, timeInterval)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fish)
}
