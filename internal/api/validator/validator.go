package validator

import (
	"TestTaskHaloLab/internal/core/dtos"
	"TestTaskHaloLab/internal/data/static"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

var validate = validator.New()

func ValidateStruct[T any](payload T) error {
	err := validate.Struct(payload)
	return err
}

// ValidateCodeName
// Parse URL parameter <groupName>. And split groupName on Name and Index and put in dtos.Group struct
func ValidateCodeName(codeName string) (*dtos.Group, error) {
	codeNameSlice := strings.Split(codeName, "%20")
	if len(codeNameSlice) != 2 {
		return nil, fiber.NewError(fiber.StatusBadRequest,
			"Wrong format of codeName. Expected '{greekLetter} {int}'")
	}
	groupName := codeNameSlice[0]
	if !contains(static.GreekLetters, groupName) {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity,
			fmt.Sprintf("Invalid groupName given: %s. Expected greek letter", groupName))
	}
	groupIndex, err := strconv.Atoi(codeNameSlice[1])
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity,
			fmt.Sprintf("Invalid groupIndex given: %s. Expected int index", codeNameSlice[1]))
	}
	return &dtos.Group{Name: groupName, Index: groupIndex}, nil
}

func ValidateGroupName(groupName string) (string, error) {
	if !contains(static.GreekLetters, groupName) {
		return "", fiber.NewError(fiber.StatusUnprocessableEntity,
			fmt.Sprintf("Invalid groupName given: %s. Expected greek letter", groupName))
	}
	return groupName, nil
}

// contains
// Check if slice of strings contains string
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ConvertAndValidateCoordinatesQueries(queries *dtos.CoordinatesString) (*dtos.Coordinates, error) {
	xMin, err := strconv.ParseFloat(queries.XMin, 64)
	if err != nil {
		return &dtos.Coordinates{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid coordinates")
	}
	xMax, err := strconv.ParseFloat(queries.XMax, 64)
	if err != nil {
		return &dtos.Coordinates{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid coordinates")
	}
	yMin, err := strconv.ParseFloat(queries.YMin, 64)
	if err != nil {
		return &dtos.Coordinates{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid coordinates")
	}
	yMax, err := strconv.ParseFloat(queries.YMax, 64)
	if err != nil {
		return &dtos.Coordinates{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid coordinates")
	}
	zMin, err := strconv.ParseFloat(queries.ZMin, 64)
	if err != nil {
		return &dtos.Coordinates{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid coordinates")
	}
	zMax, err := strconv.ParseFloat(queries.ZMax, 64)
	if err != nil {
		return &dtos.Coordinates{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid coordinates")
	}
	coordinates := dtos.Coordinates{
		XMin: xMin, XMax: xMax,
		YMin: yMin, YMax: yMax,
		ZMin: zMin, ZMax: zMax}
	if coordinates.XMin < static.XMin || coordinates.XMax > static.XMax ||
		coordinates.YMin < static.YMin || coordinates.YMax > static.YMax ||
		coordinates.ZMin < static.ZMin || coordinates.ZMax > static.ZMax ||
		coordinates.XMin >= coordinates.XMax || coordinates.YMin >= coordinates.YMax || coordinates.ZMin >= coordinates.ZMax {
		return &dtos.Coordinates{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid coordinates")
	}
	return &coordinates, nil
}

func ConvertAndValidateTimeIntervalQueries(queries *dtos.TimeIntervalUnixTimestamp) (*dtos.TimeInterval, error) {
	if queries.StarTime == "" || queries.EndTime == "" {
		return &dtos.TimeInterval{}, nil
	}
	startTime, err := ConvertUnixTimestampToTime(queries.StarTime)
	if err != nil {
		return &dtos.TimeInterval{}, err
	}
	endTime, err := ConvertUnixTimestampToTime(queries.EndTime)
	if err != nil {
		return &dtos.TimeInterval{}, err
	}
	timeInterval := dtos.TimeInterval{StarTime: startTime, EndTime: endTime}
	if endTime.Sub(startTime).Seconds() < 0 {
		return &dtos.TimeInterval{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid time")
	}
	return &timeInterval, nil
}

func ConvertUnixTimestampToTime(t string) (time.Time, error) {
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return time.Time{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid time(Cannot parse timestamp)")
	}
	return time.Unix(i, 0), nil
}
